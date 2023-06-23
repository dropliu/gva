package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/agent"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    agentReq "github.com/flipped-aurora/gin-vue-admin/server/model/agent/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type PayAgentBonusRatioApi struct {
}

var payAgentBonusRatioService = service.ServiceGroupApp.AgentServiceGroup.PayAgentBonusRatioService


// CreatePayAgentBonusRatio 创建PayAgentBonusRatio
// @Tags PayAgentBonusRatio
// @Summary 创建PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayAgentBonusRatio true "创建PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentBonusRatio/createPayAgentBonusRatio [post]
func (payAgentBonusRatioApi *PayAgentBonusRatioApi) CreatePayAgentBonusRatio(c *gin.Context) {
	var payAgentBonusRatio agent.PayAgentBonusRatio
	err := c.ShouldBindJSON(&payAgentBonusRatio)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payAgentBonusRatioService.CreatePayAgentBonusRatio(&payAgentBonusRatio); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayAgentBonusRatio 删除PayAgentBonusRatio
// @Tags PayAgentBonusRatio
// @Summary 删除PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayAgentBonusRatio true "删除PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payAgentBonusRatio/deletePayAgentBonusRatio [delete]
func (payAgentBonusRatioApi *PayAgentBonusRatioApi) DeletePayAgentBonusRatio(c *gin.Context) {
	var payAgentBonusRatio agent.PayAgentBonusRatio
	err := c.ShouldBindJSON(&payAgentBonusRatio)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payAgentBonusRatioService.DeletePayAgentBonusRatio(payAgentBonusRatio); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayAgentBonusRatioByIds 批量删除PayAgentBonusRatio
// @Tags PayAgentBonusRatio
// @Summary 批量删除PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payAgentBonusRatio/deletePayAgentBonusRatioByIds [delete]
func (payAgentBonusRatioApi *PayAgentBonusRatioApi) DeletePayAgentBonusRatioByIds(c *gin.Context) {
	var IDS request.IdsReq
    err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payAgentBonusRatioService.DeletePayAgentBonusRatioByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayAgentBonusRatio 更新PayAgentBonusRatio
// @Tags PayAgentBonusRatio
// @Summary 更新PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayAgentBonusRatio true "更新PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payAgentBonusRatio/updatePayAgentBonusRatio [put]
func (payAgentBonusRatioApi *PayAgentBonusRatioApi) UpdatePayAgentBonusRatio(c *gin.Context) {
	var payAgentBonusRatio agent.PayAgentBonusRatio
	err := c.ShouldBindJSON(&payAgentBonusRatio)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payAgentBonusRatioService.UpdatePayAgentBonusRatio(payAgentBonusRatio); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayAgentBonusRatio 用id查询PayAgentBonusRatio
// @Tags PayAgentBonusRatio
// @Summary 用id查询PayAgentBonusRatio
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query agent.PayAgentBonusRatio true "用id查询PayAgentBonusRatio"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payAgentBonusRatio/findPayAgentBonusRatio [get]
func (payAgentBonusRatioApi *PayAgentBonusRatioApi) FindPayAgentBonusRatio(c *gin.Context) {
	var payAgentBonusRatio agent.PayAgentBonusRatio
	err := c.ShouldBindQuery(&payAgentBonusRatio)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repayAgentBonusRatio, err := payAgentBonusRatioService.GetPayAgentBonusRatio(payAgentBonusRatio.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayAgentBonusRatio": repayAgentBonusRatio}, c)
	}
}

// GetPayAgentBonusRatioList 分页获取PayAgentBonusRatio列表
// @Tags PayAgentBonusRatio
// @Summary 分页获取PayAgentBonusRatio列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query agentReq.PayAgentBonusRatioSearch true "分页获取PayAgentBonusRatio列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payAgentBonusRatio/getPayAgentBonusRatioList [get]
func (payAgentBonusRatioApi *PayAgentBonusRatioApi) GetPayAgentBonusRatioList(c *gin.Context) {
	var pageInfo agentReq.PayAgentBonusRatioSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := payAgentBonusRatioService.GetPayAgentBonusRatioInfoList(pageInfo); err != nil {
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
