package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayAgentChannelRouter struct {
}

// InitPayAgentChannelRouter 初始化 PayAgentChannel 路由信息
func (s *PayAgentChannelRouter) InitPayAgentChannelRouter(Router *gin.RouterGroup) {
	payAgentChannelRouter := Router.Group("payAgentChannel").Use(middleware.OperationRecord())
	payAgentChannelRouterWithoutRecord := Router.Group("payAgentChannel")
	var payAgentChannelApi = v1.ApiGroupApp.AgentApiGroup.PayAgentChannelApi
	{
		payAgentChannelRouter.POST("createPayAgentChannel", payAgentChannelApi.CreatePayAgentChannel)   // 新建PayAgentChannel
		payAgentChannelRouter.DELETE("deletePayAgentChannel", payAgentChannelApi.DeletePayAgentChannel) // 删除PayAgentChannel
		payAgentChannelRouter.DELETE("deletePayAgentChannelByIds", payAgentChannelApi.DeletePayAgentChannelByIds) // 批量删除PayAgentChannel
		payAgentChannelRouter.PUT("updatePayAgentChannel", payAgentChannelApi.UpdatePayAgentChannel)    // 更新PayAgentChannel
	}
	{
		payAgentChannelRouterWithoutRecord.GET("findPayAgentChannel", payAgentChannelApi.FindPayAgentChannel)        // 根据ID获取PayAgentChannel
		payAgentChannelRouterWithoutRecord.GET("getPayAgentChannelList", payAgentChannelApi.GetPayAgentChannelList)  // 获取PayAgentChannel列表
	}
}
