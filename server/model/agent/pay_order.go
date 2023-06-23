// 自动生成模板PayOrder
package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// PayOrder 结构体
type PayOrder struct {
	global.GVA_MODEL
	OrderId           string    `json:"orderId" form:"orderId" gorm:"column:order_id;comment:订单id;size:255;"`
	Channel           string    `json:"channel" form:"channel" gorm:"column:channel;comment:渠道;size:255;"`
	Name              string    `json:"name" form:"name" gorm:"column:name;comment:名称;size:255;"`
	Comment           string    `json:"comment" form:"comment" gorm:"column:comment;comment:描述;size:255;"`
	Value             *float64  `json:"value" form:"value" gorm:"column:value;comment:金额;size:22;"`
	ValueType         string    `json:"valueType" form:"valueType" gorm:"column:value_type;comment:金额单位;size:255;"`
	ChannelIdentifier string    `json:"channelIdentifier" form:"channelIdentifier" gorm:"-"`
	Identifier        string    `json:"identifier" form:"identifier" gorm:"column:identifier;comment:唯一标识符（和代理一一对应）;size:255;"`
	Data              string    `json:"data" form:"data" gorm:"column:data;comment:请求的原始数据;"`
	Status            string    `json:"status" form:"status" gorm:"column:status;comment:订单的状态;size:255;"`
	CashTime          string    `json:"cashTime" gorm:"column:cash_time"`
	Fee               float64   `json:"fee" gorm:"column:fee"`
	ValueIn           float64   `json:"valueIn" gorm:"column:value_in"`
	OrderUser         OrderUser `json:"orderUser" gorm:"-"`
}

type OrderUser struct {
	Username string `json:"userName" gorm:"index;comment:用户登录名"`       // 用户登录名                                                            // 用户登录密码
	NickName string `json:"nickName" gorm:"default:系统用户;comment:用户昵称"` // 用户昵称
}

// TableName PayOrder 表名
func (PayOrder) TableName() string {
	return "pay_order"
}
