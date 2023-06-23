package v1

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/agent"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/channel"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/config"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/example"
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup  system.ApiGroup
	ExampleApiGroup example.ApiGroup
	ChannelApiGroup channel.ApiGroup
	AgentApiGroup   agent.ApiGroup
	ConfigApiGroup  config.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
