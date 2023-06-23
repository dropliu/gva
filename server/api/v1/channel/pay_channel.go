package channel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/channel"
	channelReq "github.com/flipped-aurora/gin-vue-admin/server/model/channel/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PayChannelApi struct {
}

var payChannelService = service.ServiceGroupApp.ChannelServiceGroup.PayChannelService

// CreatePayChannel 创建PayChannel
// @Tags PayChannel
// @Summary 创建PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body channel.PayChannel true "创建PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payChannel/createPayChannel [post]
func (payChannelApi *PayChannelApi) CreatePayChannel(c *gin.Context) {
	var payChannel channel.PayChannel
	err := c.ShouldBindJSON(&payChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payChannelService.CreatePayChannel(&payChannel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayChannel 删除PayChannel
// @Tags PayChannel
// @Summary 删除PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body channel.PayChannel true "删除PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payChannel/deletePayChannel [delete]
func (payChannelApi *PayChannelApi) DeletePayChannel(c *gin.Context) {
	var payChannel channel.PayChannel
	err := c.ShouldBindJSON(&payChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payChannelService.DeletePayChannel(payChannel); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayChannelByIds 批量删除PayChannel
// @Tags PayChannel
// @Summary 批量删除PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payChannel/deletePayChannelByIds [delete]
func (payChannelApi *PayChannelApi) DeletePayChannelByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payChannelService.DeletePayChannelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayChannel 更新PayChannel
// @Tags PayChannel
// @Summary 更新PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body channel.PayChannel true "更新PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payChannel/updatePayChannel [put]
func (payChannelApi *PayChannelApi) UpdatePayChannel(c *gin.Context) {
	var payChannel channel.PayChannel
	err := c.ShouldBindJSON(&payChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payChannelService.UpdatePayChannel(payChannel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayChannel 用id查询PayChannel
// @Tags PayChannel
// @Summary 用id查询PayChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query channel.PayChannel true "用id查询PayChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payChannel/findPayChannel [get]
func (payChannelApi *PayChannelApi) FindPayChannel(c *gin.Context) {
	var payChannel channel.PayChannel
	err := c.ShouldBindQuery(&payChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repayChannel, err := payChannelService.GetPayChannel(payChannel.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayChannel": repayChannel}, c)
	}
}

// GetPayChannelList 分页获取PayChannel列表
// @Tags PayChannel
// @Summary 分页获取PayChannel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query channelReq.PayChannelSearch true "分页获取PayChannel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payChannel/getPayChannelList [get]
func (payChannelApi *PayChannelApi) GetPayChannelList(c *gin.Context) {
	var pageInfo channelReq.PayChannelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := payChannelService.GetPayChannelInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
