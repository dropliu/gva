package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type PayAgentBonusRatioRouter struct {
}

// InitPayAgentBonusRatioRouter 初始化 PayAgentBonusRatio 路由信息
func (s *PayAgentBonusRatioRouter) InitPayAgentBonusRatioRouter(Router *gin.RouterGroup) {
	payAgentBonusRatioRouter := Router.Group("payAgentBonusRatio").Use(middleware.OperationRecord())
	payAgentBonusRatioRouterWithoutRecord := Router.Group("payAgentBonusRatio")
	var payAgentBonusRatioApi = v1.ApiGroupApp.AgentApiGroup.PayAgentBonusRatioApi
	{
		payAgentBonusRatioRouter.POST("createPayAgentBonusRatio", payAgentBonusRatioApi.CreatePayAgentBonusRatio)   // 新建PayAgentBonusRatio
		payAgentBonusRatioRouter.DELETE("deletePayAgentBonusRatio", payAgentBonusRatioApi.DeletePayAgentBonusRatio) // 删除PayAgentBonusRatio
		payAgentBonusRatioRouter.DELETE("deletePayAgentBonusRatioByIds", payAgentBonusRatioApi.DeletePayAgentBonusRatioByIds) // 批量删除PayAgentBonusRatio
		payAgentBonusRatioRouter.PUT("updatePayAgentBonusRatio", payAgentBonusRatioApi.UpdatePayAgentBonusRatio)    // 更新PayAgentBonusRatio
	}
	{
		payAgentBonusRatioRouterWithoutRecord.GET("findPayAgentBonusRatio", payAgentBonusRatioApi.FindPayAgentBonusRatio)        // 根据ID获取PayAgentBonusRatio
		payAgentBonusRatioRouterWithoutRecord.GET("getPayAgentBonusRatioList", payAgentBonusRatioApi.GetPayAgentBonusRatioList)  // 获取PayAgentBonusRatio列表
	}
}
