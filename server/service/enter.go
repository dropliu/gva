package service

import (
	"github.com/flipped-aurora/gin-vue-admin/server/service/agent"
	"github.com/flipped-aurora/gin-vue-admin/server/service/channel"
	"github.com/flipped-aurora/gin-vue-admin/server/service/config"
	"github.com/flipped-aurora/gin-vue-admin/server/service/example"
	"github.com/flipped-aurora/gin-vue-admin/server/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup  system.ServiceGroup
	ExampleServiceGroup example.ServiceGroup
	ChannelServiceGroup channel.ServiceGroup
	AgentServiceGroup   agent.ServiceGroup
	ConfigServiceGroup  config.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
