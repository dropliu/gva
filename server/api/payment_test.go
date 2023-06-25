package api

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"net/url"
	"sort"
	"strconv"
	"testing"
	"time"
)

func TestPayment(t *testing.T) {

}

// GET http://localhost:8888/payment/brcode?ts=3213&nonce=23121&apiKey=3123

type Payment struct {
	ID          string `yaml:"id"`
	Env         string `yaml:"env"`
	OrgID       string `yaml:"org_id"`
	WorkspaceID string `yaml:"workspace_id"`
	PrivateKey  string `yaml:"private_key"`
}

func main() {
	var url = "http://localhost:8888"
	ts := time.Now().UTC().Unix()
	nonce := rand.Int()
	apiKey := "2bcf71c6-80e4-4178-bc3d-fc8fcfa7c92b"
	apiSecret := "09e3e0b49a6b4c19b4080e773f42c896"
	url = url + "/payment/brcode?ts=" + strconv.FormatInt(ts, 10) +
		"&nonce=" + strconv.Itoa(nonce) + "&apiKey=" + apiKey

	// var body io.Reader
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("newRequest", err.Error())
		return
	}

	// req.URL.Query().Add("ts", strconv.FormatInt(ts, 10))
	sig := RequestSig(req, nil, []byte(apiSecret))
	req.URL.RawQuery = req.URL.RawQuery + "&sig=" + sig

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println("Do", err.Error())
		return
	}

	data, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("body", string(data))
	var br BrcodeResponse
	if err := json.Unmarshal(data, &br); err != nil {
		fmt.Println("unmarsahl", err.Error())
		return
	}
}

// type BrcodeResponse struct {
// 	URL        string `json:"url"`
// 	Expiration int    `json:"expiration"`
// 	CreatedAt  string `json:"creatd_at"`
// }

const (
	QuerySig = "sig"
)

func RequestSig(req *http.Request, body, secret []byte) string {
	method := req.Method
	// scheme := req.URL.Scheme
	host := req.URL.Host
	uri := req.URL.Path

	// 先检查query参数
	var queries, _ = url.ParseQuery(req.URL.RawQuery)
	var total = len(method) + len(uri)
	var keys = make([]string, 0, len(queries))
	for key, vals := range queries {
		if key != QuerySig {
			keys = append(keys, key)
			total += len(key) + len(vals[0]) + 1
		}
	}

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
			// key = url.QueryEscape(key)
			// val = url.QueryEscape(val)
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

	fmt.Println("req", b.String(), nsig)
	return nsig
}
