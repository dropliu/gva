// 自动生成模板PayChannel
package channel

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PayChannel 结构体
type PayChannel struct {
	global.GVA_MODEL
	CreateTime        *time.Time `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
	UpdateTime        *time.Time `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`
	DeleteTime        *time.Time `json:"deleteTime" form:"deleteTime" gorm:"column:delete_time;comment:删除时间;"`
	Name              string     `json:"name" form:"name" gorm:"column:name;comment:渠道名称;size:255;"`
	Identifier        string     `json:"identifier" form:"identifier" gorm:"column:identifier;comment:渠道唯一标识;size:255;"`
	Comment           string     `json:"comment" form:"comment" gorm:"column:comment;comment:描述;size:255;"`
	Status            *int       `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	ChannelIdentifier string     `json:"channelIdentifier" form:"userIdentifier" gorm:"column:channel_identifier;comment:渠道唯一标识;size:70;"`
}

// TableName PayChannel 表名
func (PayChannel) TableName() string {
	return "pay_channel"
}
