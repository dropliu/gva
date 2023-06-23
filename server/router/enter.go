package router

import (
	"github.com/flipped-aurora/gin-vue-admin/server/router/agent"
	"github.com/flipped-aurora/gin-vue-admin/server/router/channel"
	"github.com/flipped-aurora/gin-vue-admin/server/router/config"
	"github.com/flipped-aurora/gin-vue-admin/server/router/example"
	"github.com/flipped-aurora/gin-vue-admin/server/router/system"
)

type RouterGroup struct {
	System  system.RouterGroup
	Example example.RouterGroup
	Channel channel.RouterGroup
	Agent   agent.RouterGroup
	Config  config.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
