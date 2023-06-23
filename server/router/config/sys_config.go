package config

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type SysConfigRouter struct {
}

// InitSysConfigRouter 初始化 SysConfig 路由信息
func (s *SysConfigRouter) InitSysConfigRouter(Router *gin.RouterGroup) {
	sysConfigRouter := Router.Group("sysConfig").Use(middleware.OperationRecord())
	sysConfigRouterWithoutRecord := Router.Group("sysConfig")
	var sysConfigApi = v1.ApiGroupApp.ConfigApiGroup.SysConfigApi
	{
		sysConfigRouter.POST("createSysConfig", sysConfigApi.CreateSysConfig)             // 新建SysConfig
		sysConfigRouter.DELETE("deleteSysConfig", sysConfigApi.DeleteSysConfig)           // 删除SysConfig
		sysConfigRouter.DELETE("deleteSysConfigByIds", sysConfigApi.DeleteSysConfigByIds) // 批量删除SysConfig
		sysConfigRouter.PUT("updateSysConfig", sysConfigApi.UpdateSysConfig)              // 更新SysConfig
	}
	{
		sysConfigRouterWithoutRecord.GET("findSysConfig", sysConfigApi.FindSysConfig)       // 根据ID获取SysConfig
		sysConfigRouterWithoutRecord.GET("getSysConfigList", sysConfigApi.GetSysConfigList) // 获取SysConfig列表

		sysConfigRouterWithoutRecord.GET("getExchangeRate", sysConfigApi.GetExchangeRate) // 获取汇率
	}
}
