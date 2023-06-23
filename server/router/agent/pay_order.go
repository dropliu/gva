package agent

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayOrderRouter struct {
}

// InitPayOrderRouter 初始化 PayOrder 路由信息
func (s *PayOrderRouter) InitPayOrderRouter(Router *gin.RouterGroup) {
	payOrderRouter := Router.Group("payOrder").Use(middleware.OperationRecord())
	payOrderRouterWithoutRecord := Router.Group("payOrder")
	var payOrderApi = v1.ApiGroupApp.AgentApiGroup.PayOrderApi
	{
		payOrderRouter.POST("createPayOrder", payOrderApi.CreatePayOrder)             // 新建PayOrder
		payOrderRouter.DELETE("deletePayOrder", payOrderApi.DeletePayOrder)           // 删除PayOrder
		payOrderRouter.DELETE("deletePayOrderByIds", payOrderApi.DeletePayOrderByIds) // 批量删除PayOrder
		payOrderRouter.PUT("updatePayOrder", payOrderApi.UpdatePayOrder)              // 更新PayOrder
	}
	{
		payOrderRouterWithoutRecord.GET("findPayOrder", payOrderApi.FindPayOrder)                 // 根据ID获取PayOrder
		payOrderRouterWithoutRecord.GET("getPayOrderList", payOrderApi.GetPayOrderList)           // 获取PayOrder列表
		payOrderRouterWithoutRecord.GET("getAgentPayOrderList", payOrderApi.GetAgentPayOrderList) // 获取PayOrder列表

		payOrderRouterWithoutRecord.GET("getAgentStaticDay", payOrderApi.GetAgentPayOrderStaticDay)
		payOrderRouterWithoutRecord.GET("getAgentStatic", payOrderApi.GetAgentPayOrderStatic)

		payOrderRouterWithoutRecord.GET("getStaticDay", payOrderApi.GetPayOrderStaticDay)
		payOrderRouterWithoutRecord.GET("getStatic", payOrderApi.GetPayOrderStatic)
	}
}
