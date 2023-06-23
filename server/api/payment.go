package api

import (
	"fmt"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils/starkbank"
	"github.com/gin-gonic/gin"
	"github.com/starkinfra/core-go/starkcore/user/organization"
	"go.uber.org/zap"
)

func CreatePayment(c *gin.Context) {
	var req service.Payment
	c.ShouldBindJSON(&req)
	if len(req.Name) == 0 || len(req.Identifier) == 0 {
		response.FailWithMessage("参数异常", c)
		return
	}

	now := time.Now()
	req.CreatedAt = now
	req.UpdatedAt = now

	if err := service.CreatePayment(req); err != nil {
		global.GVA_LOG.Error("createPaymetn", zap.Error(err))
		response.FailWithMessage("创建失败，请重试", c)
		return
	}

	response.Ok(c)
}

func UpdatePayment(c *gin.Context) {
	var req service.Payment
	c.ShouldBindJSON(&req)
	if req.ID <= 0 || len(req.Name) == 0 || len(req.Identifier) == 0 {
		response.FailWithMessage("参数异常", c)
		return
	}

	req.UpdatedAt = time.Now()
	if err := service.UpdatePayment(req); err != nil {
		response.FailWithMessage("更新失败，请重试", c)
		return
	}

	response.Ok(c)
}

func DeletePayment(c *gin.Context) {
	id, err := strconv.ParseInt(c.Query("id"), 10, 0)
	if err != nil && id <= 0 {
		response.FailWithMessage("参数异常", c)
		return
	}

	var req = service.Payment{
		ID: id,
	}
	if err := service.DeletePayment(req); err != nil {
		response.FailWithMessage("更新失败，请重试", c)
		return
	}

	response.Ok(c)
}

func GetAllPayment(c *gin.Context) {
	list, err := service.ListAllPayment()
	if err != nil {
		response.FailWithMessage("服务异常，请重试", c)
		return
	}
	result := response.PageResult{
		List:     &list,
		Total:    int64(len(list)),
		Page:     1,
		PageSize: 10,
	}

	response.OkWithData(&result, c)
}

type BrcodeResponse struct {
	URL        string `json:"url"`
	Expiration int    `json:"expiration"`
	CreatedAt  string `json:"creatd_at"`
}

// GetBrcode 获取代理专属的二维码
func GetBrcode(c *gin.Context) {
	uuid := c.Query("apiKey")
	if len(uuid) == 0 {
		response.FailWithMessage("user not exists", c)
		return
	}

	// 查找用户
	user, err := service.ServiceGroupApp.SystemServiceGroup.
		UserService.FindUserByUuid(uuid)
	if err != nil {
		response.FailWithMessage("user not exists", c)
		return
	}

	pay, ok := global.FindPayment(user.Payment)
	if !ok {
		response.FailWithMessage("payment not exists", c)
		return
	}

	// 暂时只支持starkbank
	if pay.ID != "starkbank" {
		response.FailWithMessage("unsupported payment", c)
		return
	}

	// 构建参数
	var usr = organization.Organization{
		Id:          pay.OrgID,
		PrivateKey:  pay.PrivateKey,
		Environment: pay.Env,
		WorkspaceId: pay.WorkspaceID,
	}

	var expiration = 3600

	tags := []string{
		fmt.Sprintf("payment=%s", pay.ID),
		fmt.Sprintf("uuid=%s", uuid),
		fmt.Sprintf("createdAt=%s", time.Now().Format(time.RFC3339)),
	}
	brcode, err := starkbank.CreateDynamicBrcode(usr, 4000, expiration, tags...)
	if err != nil {
		global.GVA_LOG.Error("GetBrcode", zap.Error(err))
		response.FailWithMessage("create brcode failed, try again please", c)
		return
	}

	var result = BrcodeResponse{
		URL:        brcode.PictureUrl,
		Expiration: expiration,
		CreatedAt:  brcode.Created.Format(time.RFC3339),
	}

	response.OkWithData(&result, c)
	return
}
