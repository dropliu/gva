package middleware

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"gorm.io/gorm"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
	"github.com/flipped-aurora/gin-vue-admin/server/service"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

const (
	HeaderToken = "x-token"
)

var jwtService = service.ServiceGroupApp.SystemServiceGroup.JwtService

func JWTAuth(c *gin.Context) {
	// 我们这里jwt鉴权取头部信息 x-token 登录时回返回token信息 这里前端需要把token存储到cookie或者本地localStorage中 不过需要跟后端协商过期时间 可以约定刷新令牌或者重新登录
	token := c.Request.Header.Get("x-token")
	if token == "" {
		// TODO: 没找到token，尝试走请求接口签名的逻辑

		response.FailWithDetailed(gin.H{"reload": true}, "未登录或非法访问", c)
		c.Abort()
		return
	}
	if jwtService.IsBlacklist(token) {
		response.FailWithDetailed(gin.H{"reload": true}, "您的帐户异地登陆或令牌失效", c)
		c.Abort()
		return
	}
	j := utils.NewJWT()

	// parseToken 解析token包含的信息
	claims, err := j.ParseToken(token)
	if err != nil {
		if errors.Is(err, utils.TokenExpired) {
			response.FailWithDetailed(gin.H{"reload": true}, "授权已过期", c)
			c.Abort()
			return
		}
		response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
		c.Abort()
		return
	}

	// 已登录用户被管理员禁用 需要使该用户的jwt失效 此处比较消耗性能 如果需要 请自行打开
	// 用户被删除的逻辑 需要优化 此处比较消耗性能 如果需要 请自行打开

	//if user, err := userService.FindUserByUuid(claims.UUID.String()); err != nil || user.Enable == 2 {
	//	_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: token})
	//	response.FailWithDetailed(gin.H{"reload": true}, err.Error(), c)
	//	c.Abort()
	//}
	if claims.ExpiresAt-time.Now().Unix() < claims.BufferTime {
		dr, _ := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
		claims.ExpiresAt = time.Now().Add(dr).Unix()
		newToken, _ := j.CreateTokenByOldToken(token, *claims)
		newClaims, _ := j.ParseToken(newToken)
		c.Header("new-token", newToken)
		c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt, 10))
		if global.GVA_CONFIG.System.UseMultipoint {
			RedisJwtToken, err := jwtService.GetRedisJWT(newClaims.Username)
			if err != nil {
				global.GVA_LOG.Error("get redis jwt failed", zap.Error(err))
			} else { // 当之前的取成功时才进行拉黑操作
				_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
			}
			// 无论如何都要记录当前的活跃状态
			_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
		}
	}
	c.Set("claims", claims)
	// c.Next()
}

func UserAuth(c *gin.Context) {
	token := c.Request.Header.Get("x-token")
	if len(token) > 0 {
		// 有token时，优先处理token
		JWTAuth(c)
	} else {
		SigAuth(c)
	}
}

const (
	QueryNonce  = "nonce"
	QueryTS     = "ts"
	QuerySig    = "sig"
	QueryApiKey = "apiKey"

	SigValidDur = 10 // 默认10s
)

// SigAuth 签名校验
func SigAuth(c *gin.Context) {
	secret := validSigParam(c)
	if len(secret) == 0 {
		c.Abort()
		response.FailWithMessage("参数异常", c)
		return
	}

	body, err := RequestSig(c.Request, secret)
	if err != nil { // 没通过校验
		c.Abort()
		response.FailWithMessage("签名不一致", c)
		return
	}

	// 将body放置到gin的上下文当中
	c.Set(gin.BodyBytesKey, body)

	// TODO: 生成一个
}

func validSigParam(c *gin.Context) []byte {
	var (
		apiKey   = c.Query(QueryApiKey)
		nonceStr = c.Query(QueryNonce)
		tsStr    = c.Query(QueryTS)
		sig      = c.Query(QuerySig)
	)

	// 1. 参数校验
	if len(apiKey) == 0 || len(nonceStr) == 0 ||
		len(tsStr) == 0 || len(sig) == 0 {
		return nil
	}

	nonce, err := strconv.Atoi(nonceStr)
	if err != nil || nonce <= 0 {
		return nil
	}

	ts, err := strconv.ParseInt(tsStr, 10, 0)
	if err != nil || ts <= 0 {
		return nil
	}

	now := time.Now().UTC().Unix()
	if int64(nonce) == ts || now-ts > SigValidDur {
		return nil
	}

	// 拿到uuid对应的api_secret
	secret, err := LoadApiSecret(apiKey)
	if err != nil {
		fmt.Println("loadApiSeret", err.Error())
		return nil
	}
	return secret
}

var emptyRecord interface{} = struct{}{}

func LoadApiSecret(id string) ([]byte, error) {
	val, exist := global.ApiSecretCache.Get(id)
	if val == emptyRecord {
		return nil, nil
	}
	secret, ok := val.([]byte)
	if !exist || !ok || len(secret) == 0 {
		as, err := service.ServiceGroupApp.SystemServiceGroup.UserService.GetApiSecretByUUID(id)
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				global.ApiSecretCache.Set(id, emptyRecord, 30*time.Second)
				err = nil
			}
			return nil, err
		}
		secret = []byte(as)
		global.ApiSecretCache.Set(id, secret, 10*time.Minute)
	}

	return secret, nil
}

// RequestSig (请求签名校验)
// 标准库的http.Request.Body只能读一次，读到的body需要返回给业务逻辑
func RequestSig(req *http.Request, secret []byte) (body []byte, err error) {
	method := req.Method
	// scheme := req.URL.Scheme
	host := req.Host
	uri := req.URL.Path

	// 先检查query参数
	var originalSig string
	var queries, _ = url.ParseQuery(req.URL.RawQuery)
	var total = len(method) + len(host) + len(uri)
	var keys = make([]string, 0, len(queries))
	for key, vals := range queries {
		if key != QuerySig {
			keys = append(keys, key)
			total += len(key) + len(vals[0]) + 1
		} else {
			originalSig = vals[0]
		}
	}

	// 读取body
	body, err = ioutil.ReadAll(req.Body)
	if err != nil {
		return
	}
	// 防止其它地方读数据异常
	req.Body = io.NopCloser(bytes.NewReader(body))

	var b bytes.Buffer
	b.Grow(total)
	// 先写入method和uri
	b.WriteString(method)
	// b.WriteString(scheme)
	// b.WriteString("//")
	b.WriteString(host)
	b.WriteString(uri)
	// 如果有query参数
	if len(keys) > 0 {
		// 从小到大排序
		sort.Strings(keys)
		// 根据排序组装值
		for i, key := range keys {
			var val = queries.Get(key)
			key = url.QueryEscape(key)
			val = url.QueryEscape(val)
			if i == 0 {
				b.WriteByte('?')
			} else {
				b.WriteByte('&')
			}
			b.WriteString(key)
			b.WriteByte('=')
			b.WriteString(val)
		}
	}
	// 如果有body
	if len(body) > 0 {
		b.WriteByte('&')
		b.Write(body)
	}

	mac := hmac.New(sha256.New, secret)
	mac.Write(b.Bytes())
	// hamc256(secret) -> bas64.encode -> queryEscape
	nsig := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	fmt.Println("url", b.String(), "sig", nsig)
	//  校验签名
	if nsig != originalSig {
		return nil, ErrIllegalRequest
	}
	return body, nil
}

var (
	ErrIllegalParam   = errors.New("illegal parameter")
	ErrSiguateExpired = errors.New("")
	ErrIllegalRequest = errors.New("")
)

func getSigParamsFromQuery(vals url.Values) (nonce int, ts int64, sig string, err error) {
	nonceStr := vals.Get(QueryNonce)
	tsStr := vals.Get(QueryTS)
	sig = vals.Get(QuerySig)

	//
	if len(nonceStr) == 0 || len(tsStr) == 0 || len(sig) == 0 {
		err = ErrIllegalParam
		return
	}

	nonce, err = strconv.Atoi(nonceStr)
	if err != nil || nonce <= 0 {
		err = ErrIllegalParam
		return
	}

	ts, err = strconv.ParseInt(tsStr, 10, 0)
	if err != nil || ts <= 0 {
		err = ErrIllegalParam
		return
	}

	if int64(nonce) == ts {
		err = ErrIllegalParam
		return
	}

	return
}
