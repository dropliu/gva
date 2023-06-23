package channel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayChannelRouter struct {
}

// InitPayChannelRouter 初始化 PayChannel 路由信息
func (s *PayChannelRouter) InitPayChannelRouter(Router *gin.RouterGroup) {
	payChannelRouter := Router.Group("payChannel").Use(middleware.OperationRecord())
	payChannelRouterWithoutRecord := Router.Group("payChannel")
	var payChannelApi = v1.ApiGroupApp.ChannelApiGroup.PayChannelApi
	{
		payChannelRouter.POST("createPayChannel", payChannelApi.CreatePayChannel)   // 新建PayChannel
		payChannelRouter.DELETE("deletePayChannel", payChannelApi.DeletePayChannel) // 删除PayChannel
		payChannelRouter.DELETE("deletePayChannelByIds", payChannelApi.DeletePayChannelByIds) // 批量删除PayChannel
		payChannelRouter.PUT("updatePayChannel", payChannelApi.UpdatePayChannel)    // 更新PayChannel
	}
	{
		payChannelRouterWithoutRecord.GET("findPayChannel", payChannelApi.FindPayChannel)        // 根据ID获取PayChannel
		payChannelRouterWithoutRecord.GET("getPayChannelList", payChannelApi.GetPayChannelList)  // 获取PayChannel列表
	}
}
