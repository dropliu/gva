package request

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/config"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"time"
)

type SysConfigSearch struct{
    config.SysConfig
    StartCreatedAt *time.Time `json:"startCreatedAt" form:"startCreatedAt"`
    EndCreatedAt   *time.Time `json:"endCreatedAt" form:"endCreatedAt"`
    request.PageInfo
}
