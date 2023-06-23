package agent

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayWithdrawalCashRouter struct {
}

// InitPayWithdrawalCashRouter 初始化 PayWithdrawalCash 路由信息
func (s *PayWithdrawalCashRouter) InitPayWithdrawalCashRouter(Router *gin.RouterGroup) {
	payWithdrawalCashRouter := Router.Group("payWithdrawalCash").Use(middleware.OperationRecord())
	payWithdrawalCashRouterWithoutRecord := Router.Group("payWithdrawalCash")
	var payWithdrawalCashApi = v1.ApiGroupApp.AgentApiGroup.PayWithdrawalCashApi
	{
		payWithdrawalCashRouter.POST("createPayWithdrawalCash", payWithdrawalCashApi.CreatePayWithdrawalCash)             // 新建PayWithdrawalCash
		payWithdrawalCashRouter.DELETE("deletePayWithdrawalCash", payWithdrawalCashApi.DeletePayWithdrawalCash)           // 删除PayWithdrawalCash
		payWithdrawalCashRouter.DELETE("deletePayWithdrawalCashByIds", payWithdrawalCashApi.DeletePayWithdrawalCashByIds) // 批量删除PayWithdrawalCash
		payWithdrawalCashRouter.PUT("updatePayWithdrawalCash", payWithdrawalCashApi.UpdatePayWithdrawalCash)              // 更新PayWithdrawalCash
	}
	{
		payWithdrawalCashRouterWithoutRecord.GET("findPayWithdrawalCash", payWithdrawalCashApi.FindPayWithdrawalCash)       // 根据ID获取PayWithdrawalCash
		payWithdrawalCashRouterWithoutRecord.GET("getPayWithdrawalCashList", payWithdrawalCashApi.GetPayWithdrawalCashList) // 获取PayWithdrawalCash列表
		payWithdrawalCashRouterWithoutRecord.GET("getAgentPayWithdrawalCashList", payWithdrawalCashApi.GetAgentPayWithdrawalCashList)
	}
}
