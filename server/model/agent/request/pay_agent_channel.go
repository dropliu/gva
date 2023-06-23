package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/agent"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type PayAgentChannelSearch struct{
    agent.PayAgentChannel
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
