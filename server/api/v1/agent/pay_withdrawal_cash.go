package agent

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	agentReq "github.com/flipped-aurora/gin-vue-admin/server/model/agent/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/withdrawalcash"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PayWithdrawalCashApi struct {
}

var payWithdrawalCashService = service.ServiceGroupApp.AgentServiceGroup.PayWithdrawalCashService
var sysConfigService = service.ServiceGroupApp.ConfigServiceGroup.SysConfigService

// CreatePayWithdrawalCash 创建PayWithdrawalCash
// @Tags PayWithdrawalCash
// @Summary 创建PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayWithdrawalCash true "创建PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payWithdrawalCash/createPayWithdrawalCash [post]
func (payWithdrawalCashApi *PayWithdrawalCashApi) CreatePayWithdrawalCash(c *gin.Context) {
	var payWithdrawalCash withdrawalcash.PayWithdrawalCash
	err := c.ShouldBindJSON(&payWithdrawalCash)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uuid := utils.GetUserUuid(c)
	payWithdrawalCash.Uuid = uuid.String()

	list, moneySum, err := payOrderService.GetAgentCashOrder(uuid.String())

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if len(list) == 0 {
		response.FailWithMessage("没有需要提现的订单记录", c)
		return
	}

	payWithdrawalCash.MoneySum = &moneySum

	bonusRatio, err := payAgentBonusRatioService.GetPayAgentBonusRatioByUuid(uuid.String())
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if *bonusRatio.Value <= float64(0) {
		response.FailWithMessage("没有设置分成比例", c)
		return
	}

	money := moneySum * (*bonusRatio.Value) / 100

	payWithdrawalCash.Money = &money

	moneyEnd := money

	if *payWithdrawalCash.CashType == 0 {
		if payWithdrawalCash.ExchangeRate == "" {
			response.FailWithMessage("获取汇率失败", c)
			return
		}

		exchangeRate, _ := strconv.ParseFloat(payWithdrawalCash.ExchangeRate, 64)
		moneyEnd = money / exchangeRate
		payWithdrawalCash.ExchangeRate = fmt.Sprintf("1:%f", exchangeRate)
	}

	payWithdrawalCash.MoneyEnd = &moneyEnd

	err = payOrderService.UpdateAgentCashOrder(list)

	if err != nil {
		global.GVA_LOG.Error("更新订单表失败!", zap.Error(err))
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := payWithdrawalCashService.CreatePayWithdrawalCash(&payWithdrawalCash); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayWithdrawalCash 删除PayWithdrawalCash
// @Tags PayWithdrawalCash
// @Summary 删除PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayWithdrawalCash true "删除PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payWithdrawalCash/deletePayWithdrawalCash [delete]
func (payWithdrawalCashApi *PayWithdrawalCashApi) DeletePayWithdrawalCash(c *gin.Context) {
	var payWithdrawalCash withdrawalcash.PayWithdrawalCash
	err := c.ShouldBindJSON(&payWithdrawalCash)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payWithdrawalCashService.DeletePayWithdrawalCash(payWithdrawalCash); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayWithdrawalCashByIds 批量删除PayWithdrawalCash
// @Tags PayWithdrawalCash
// @Summary 批量删除PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payWithdrawalCash/deletePayWithdrawalCashByIds [delete]
func (payWithdrawalCashApi *PayWithdrawalCashApi) DeletePayWithdrawalCashByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payWithdrawalCashService.DeletePayWithdrawalCashByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayWithdrawalCash 更新PayWithdrawalCash
// @Tags PayWithdrawalCash
// @Summary 更新PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayWithdrawalCash true "更新PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payWithdrawalCash/updatePayWithdrawalCash [put]
func (payWithdrawalCashApi *PayWithdrawalCashApi) UpdatePayWithdrawalCash(c *gin.Context) {
	var payWithdrawalCash withdrawalcash.PayWithdrawalCash
	err := c.ShouldBindJSON(&payWithdrawalCash)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payWithdrawalCashService.UpdatePayWithdrawalCash(payWithdrawalCash); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayWithdrawalCash 用id查询PayWithdrawalCash
// @Tags PayWithdrawalCash
// @Summary 用id查询PayWithdrawalCash
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query agent.PayWithdrawalCash true "用id查询PayWithdrawalCash"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payWithdrawalCash/findPayWithdrawalCash [get]
func (payWithdrawalCashApi *PayWithdrawalCashApi) FindPayWithdrawalCash(c *gin.Context) {
	var payWithdrawalCash withdrawalcash.PayWithdrawalCash
	err := c.ShouldBindQuery(&payWithdrawalCash)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repayWithdrawalCash, err := payWithdrawalCashService.GetPayWithdrawalCash(payWithdrawalCash.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayWithdrawalCash": repayWithdrawalCash}, c)
	}
}

// GetPayWithdrawalCashList 分页获取PayWithdrawalCash列表
// @Tags PayWithdrawalCash
// @Summary 分页获取PayWithdrawalCash列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query agentReq.PayWithdrawalCashSearch true "分页获取PayWithdrawalCash列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payWithdrawalCash/getPayWithdrawalCashList [get]
func (payWithdrawalCashApi *PayWithdrawalCashApi) GetPayWithdrawalCashList(c *gin.Context) {
	var pageInfo agentReq.PayWithdrawalCashSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := payWithdrawalCashService.GetPayWithdrawalCashInfoList(pageInfo, ""); err != nil {
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

// 获取agent的 提现记录
func (payWithdrawalCashApi *PayWithdrawalCashApi) GetAgentPayWithdrawalCashList(c *gin.Context) {
	var pageInfo agentReq.PayWithdrawalCashSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uuid := utils.GetUserUuid(c)
	if list, total, err := payWithdrawalCashService.GetPayWithdrawalCashInfoList(pageInfo, uuid.String()); err != nil {
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
