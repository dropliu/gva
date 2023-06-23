package agent

import (
	"fmt"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/agent"
	agentReq "github.com/flipped-aurora/gin-vue-admin/server/model/agent/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type PayOrderApi struct {
}

var payOrderService = service.ServiceGroupApp.AgentServiceGroup.PayOrderService

// CreatePayOrder 创建PayOrder
// @Tags PayOrder
// @Summary 创建PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayOrder true "创建PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payOrder/createPayOrder [post]
func (payOrderApi *PayOrderApi) CreatePayOrder(c *gin.Context) {
	var payOrder agent.PayOrder
	err := c.ShouldBindJSON(&payOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	fee, err := sysConfigService.CalculationFee(payOrder.ValueIn)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	payOrder.Fee = fee
	val := payOrder.ValueIn - fee
	payOrder.Value = &val
	if err := payOrderService.CreatePayOrder(&payOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeletePayOrder 删除PayOrder
// @Tags PayOrder
// @Summary 删除PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayOrder true "删除PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /payOrder/deletePayOrder [delete]
func (payOrderApi *PayOrderApi) DeletePayOrder(c *gin.Context) {
	var payOrder agent.PayOrder
	err := c.ShouldBindJSON(&payOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payOrderService.DeletePayOrder(payOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeletePayOrderByIds 批量删除PayOrder
// @Tags PayOrder
// @Summary 批量删除PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /payOrder/deletePayOrderByIds [delete]
func (payOrderApi *PayOrderApi) DeletePayOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payOrderService.DeletePayOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdatePayOrder 更新PayOrder
// @Tags PayOrder
// @Summary 更新PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body agent.PayOrder true "更新PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /payOrder/updatePayOrder [put]
func (payOrderApi *PayOrderApi) UpdatePayOrder(c *gin.Context) {
	var payOrder agent.PayOrder
	err := c.ShouldBindJSON(&payOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := payOrderService.UpdatePayOrder(payOrder); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindPayOrder 用id查询PayOrder
// @Tags PayOrder
// @Summary 用id查询PayOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query agent.PayOrder true "用id查询PayOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /payOrder/findPayOrder [get]
func (payOrderApi *PayOrderApi) FindPayOrder(c *gin.Context) {
	var payOrder agent.PayOrder
	err := c.ShouldBindQuery(&payOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if repayOrder, err := payOrderService.GetPayOrder(payOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"repayOrder": repayOrder}, c)
	}
}

// GetPayOrderList 分页获取PayOrder列表
// @Tags PayOrder
// @Summary 分页获取PayOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query agentReq.PayOrderSearch true "分页获取PayOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /payOrder/getPayOrderList [get]
func (payOrderApi *PayOrderApi) GetPayOrderList(c *gin.Context) {
	var pageInfo agentReq.PayOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := payOrderService.GetPayOrderInfoList(pageInfo, ""); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {

		// r, err := payOrderApi.updatePayOrderFee(list)
		// if err != nil {
		// 	global.GVA_LOG.Error("获取失败!", zap.Error(err))
		// 	response.FailWithMessage("获取失败", c)
		// 	return
		// }
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

// 获取代理自己的订单记录
func (payOrderApi *PayOrderApi) GetAgentPayOrderList(c *gin.Context) {
	var pageInfo agentReq.PayOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	uuid := utils.GetUserUuid(c)
	if list, total, err := payOrderService.GetPayOrderInfoList(pageInfo, uuid.String()); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {

		// TODO: 暂时不处理
		// r, err := payOrderApi.updatePayOrderFee(list)
		// if err != nil {
		// 	global.GVA_LOG.Error("获取失败!", zap.Error(err))
		// 	response.FailWithMessage("获取失败", c)
		// 	return
		// }
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (payOrderApi *PayOrderApi) GetAgentPayOrderStaticDay(c *gin.Context) {
	uuid := utils.GetUserUuid(c)

	moneyAll, money, count, err := payOrderService.GetPayOrderStaticDay(uuid.String())
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	bonusRatio, err := payAgentBonusRatioService.GetPayAgentBonusRatioByUuid(uuid.String())
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	money = money * (*bonusRatio.Value) / 100

	m := make(map[string]interface{}, 0)
	m["moneyAll"] = moneyAll
	m["money"] = money
	m["count"] = count
	response.OkWithDetailed(m, "获取成功", c)
}

func (payOrderApi *PayOrderApi) GetPayOrderStaticDay(c *gin.Context) {

	moneyAll, money, count, err := payOrderService.GetPayOrderStaticDay("")
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	commissionRate, err := sysConfigService.GetCommissionRate()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	money = money * commissionRate / 100

	m := make(map[string]interface{}, 0)
	m["moneyAll"] = moneyAll
	m["money"] = money
	m["count"] = count
	response.OkWithDetailed(m, "获取成功", c)
}

func (payOrderApi *PayOrderApi) GetAgentPayOrderStatic(c *gin.Context) {
	uuid := utils.GetUserUuid(c)
	moneyAll, _, err := payOrderService.GetPayOrderStatic(uuid.String())

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	bonusRatio, err := payAgentBonusRatioService.GetPayAgentBonusRatioByUuid(uuid.String())
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	paidVal, val, err := payWithdrawalCashService.GetPayWithdrawalCashStatic(uuid.String())
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	moneyAll = moneyAll * (*bonusRatio.Value) / 100

	m := make(map[string]interface{}, 0)
	//实时余额 = 佣金 - 已经成功提现数
	m["moneyAll"] = moneyAll - paidVal
	//待结算 = 佣金 - 已经发起提现的数值（不管成功）
	m["money"] = moneyAll - val
	response.OkWithDetailed(m, "获取成功", c)

}

func (payOrderApi *PayOrderApi) GetPayOrderStatic(c *gin.Context) {

	moneyAll, _, err := payOrderService.GetPayOrderStatic("")

	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	paidVal, val, err := payWithdrawalCashService.GetPayWithdrawalCashStatic("")
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	commissionRate, err := sysConfigService.GetCommissionRate()
	if err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
		return
	}

	m := make(map[string]interface{}, 0)
	moneyAll = moneyAll * commissionRate / 100

	fmt.Println("moneyAll:", moneyAll, "paidVal:", paidVal)
	//实时余额 = 佣金 - 已经成功提现数
	m["moneyAll"] = moneyAll - paidVal
	//待结算 = 佣金 - 已经发起提现的数值（不管成功）
	m["money"] = moneyAll - val
	response.OkWithDetailed(m, "获取成功", c)

}

// 更新订单手续费
func (payOrderApi *PayOrderApi) updatePayOrderFee(list []agent.PayOrder) (r []agent.PayOrder, err error) {
	for _, item := range list {
		if item.ValueIn == 0 && *item.Value != 0 {
			fee, err := sysConfigService.CalculationFee(item.ValueIn)
			if err != nil {
				return r, err
			}
			item.Fee = fee
			item.ValueIn = *item.Value
			val := *item.Value - fee
			item.Value = &val
			r = append(r, item)

			if err := payOrderService.UpdatePayOrder(item); err != nil {
				global.GVA_LOG.Error("更新失败!", zap.Error(err))
				continue
			}
		}
	}
	return
}
