// 自动生成模板PayAgentChannel
package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PayAgentChannel 结构体
type PayAgentChannel struct {
	global.GVA_MODEL
	Name              string `json:"name" form:"name" gorm:"column:name;comment:名称;size:255;"`
	Identifier        string `json:"identifier" form:"identifier" gorm:"column:identifier;comment:渠道唯一标识;size:255;"`
	Uuid              string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户uuid;size:191;"`
	Comment           string `json:"comment" form:"comment" gorm:"column:comment;comment:描述;size:255;"`
	Status            *int   `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	ChannelIdentifier string `json:"channelIdentifier" form:"userIdentifier" gorm:"column:channel_identifier;comment:渠道唯一标识;size:70;"`
}

// TableName PayAgentChannel 表名
func (PayAgentChannel) TableName() string {
	return "pay_agent_channel"
}
