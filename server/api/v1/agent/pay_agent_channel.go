package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/agent"
	agentReq "github.com/flipped-aurora/gin-vue-admin/server/model/agent/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PayAgentChannelApi struct {
}

var payAgentChannelService = service.ServiceGroupApp.AgentServiceGroup.PayAgentChannelService

var payChannelService = service.ServiceGroupApp.ChannelServiceGroup.PayChannelService

// CreatePayAgentChannel 创建PayAgentChannel
// @Tags PayAgentChannel
// @Summary 创建PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayAgentChannel true "创建PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentChannel/createPayAgentChannel [post]
func (payAgentChannelApi *PayAgentChannelApi) CreatePayAgentChannel(c *gin.Context) {
	var payAgentChannel agent.PayAgentChannel
	err := c.ShouldBindJSON(&payAgentChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	_, err = payChannelService.GetPayChannelInfoByUserIdentifier(payAgentChannel.ChannelIdentifier, payAgentChannel.Identifier)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := payAgentChannelService.CreatePayAgentChannel(&payAgentChannel); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayAgentChannel 删除PayAgentChannel
// @Tags PayAgentChannel
// @Summary 删除PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayAgentChannel true "删除PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payAgentChannel/deletePayAgentChannel [delete]
func (payAgentChannelApi *PayAgentChannelApi) DeletePayAgentChannel(c *gin.Context) {
	var payAgentChannel agent.PayAgentChannel
	err := c.ShouldBindJSON(&payAgentChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payAgentChannelService.DeletePayAgentChannel(payAgentChannel); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayAgentChannelByIds 批量删除PayAgentChannel
// @Tags PayAgentChannel
// @Summary 批量删除PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payAgentChannel/deletePayAgentChannelByIds [delete]
func (payAgentChannelApi *PayAgentChannelApi) DeletePayAgentChannelByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payAgentChannelService.DeletePayAgentChannelByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayAgentChannel 更新PayAgentChannel
// @Tags PayAgentChannel
// @Summary 更新PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayAgentChannel true "更新PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payAgentChannel/updatePayAgentChannel [put]
func (payAgentChannelApi *PayAgentChannelApi) UpdatePayAgentChannel(c *gin.Context) {
	var payAgentChannel agent.PayAgentChannel
	err := c.ShouldBindJSON(&payAgentChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	_, err = payChannelService.GetPayChannelInfoByUserIdentifier(payAgentChannel.ChannelIdentifier, payAgentChannel.Identifier)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := payAgentChannelService.UpdatePayAgentChannel(payAgentChannel); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayAgentChannel 用id查询PayAgentChannel
// @Tags PayAgentChannel
// @Summary 用id查询PayAgentChannel
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query agent.PayAgentChannel true "用id查询PayAgentChannel"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payAgentChannel/findPayAgentChannel [get]
func (payAgentChannelApi *PayAgentChannelApi) FindPayAgentChannel(c *gin.Context) {
	var payAgentChannel agent.PayAgentChannel
	err := c.ShouldBindQuery(&payAgentChannel)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repayAgentChannel, err := payAgentChannelService.GetPayAgentChannel(payAgentChannel.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayAgentChannel": repayAgentChannel}, c)
	}
}

// GetPayAgentChannelList 分页获取PayAgentChannel列表
// @Tags PayAgentChannel
// @Summary 分页获取PayAgentChannel列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query agentReq.PayAgentChannelSearch true "分页获取PayAgentChannel列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentChannel/getPayAgentChannelList [get]
func (payAgentChannelApi *PayAgentChannelApi) GetPayAgentChannelList(c *gin.Context) {
	var pageInfo agentReq.PayAgentChannelSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := payAgentChannelService.GetPayAgentChannelInfoList(pageInfo); err != nil {
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
