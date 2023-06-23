package agent

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/gin-gonic/gin"
)

type PayAgentUserRouter struct {
}

func (s *PayAgentUserRouter) InitPayAgentUserRouter(Router *gin.RouterGroup) {
	//	payAgentUserRouter := Router.Group("payAgentUser").Use(middleware.OperationRecord())
	payAgentUserRouterWithoutRecord := Router.Group("payAgentUser")
	var payAgentUserApi = v1.ApiGroupApp.AgentApiGroup.PayAgentUserApi
	{
		payAgentUserRouterWithoutRecord.POST("getAgentUserList", payAgentUserApi.GetAgentUserList)       // 分页获取用户列表
		payAgentUserRouterWithoutRecord.POST("getAgentUserAllList", payAgentUserApi.GetAgentUserAllList) // 分页获取用户列表

		payAgentUserRouterWithoutRecord.POST("addUser", payAgentUserApi.Register)
	}

}
