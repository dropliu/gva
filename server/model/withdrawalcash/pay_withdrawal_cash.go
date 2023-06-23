// 自动生成模板PayWithdrawalCash
package withdrawalcash

import (
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

// PayWithdrawalCash 结构体
type PayWithdrawalCash struct {
	global.GVA_MODEL
	CreateTime   *time.Time     `json:"createTime" form:"createTime" gorm:"column:create_time;comment:创建时间;"`
	UpdateTime   *time.Time     `json:"updateTime" form:"updateTime" gorm:"column:update_time;comment:更新时间;"`
	DeleteTime   *time.Time     `json:"deleteTime" form:"deleteTime" gorm:"column:delete_time;comment:删除时间;"`
	Status       *int           `json:"status" form:"status" gorm:"column:status;comment:状态;size:10;"`
	CashType     *int           `json:"cashType" form:"cashType" gorm:"column:cash_type;comment:提现类型;size:10;"`
	Account      string         `json:"account" form:"account" gorm:"column:account;comment:账号信息;size:255;"`
	Comment      string         `json:"comment" form:"comment" gorm:"column:comment;comment:提现备注;size:255;"`
	OpenpixKey   string         `json:"openpixKey" form:"openpixKey" gorm:"column:openpix_key;comment:openfix_key;size:255;"`
	Uuid         string         `json:"uuid"`
	Fee          float64        `json:"fee"`
	ExchangeRate string         `json:"exchangeRate" gorm:"column:exchange_rate"`
	User         system.SysUser `json:"user" gorm:"foreignkey:UUID;references:Uuid""`
	Money        *float64       `json:"money" gorm:"column:money"`
	MoneySum     *float64       `json:"moneySum" gorm:"column:money_sum"`
	MoneyEnd     *float64       `json:"moneyEnd" gorm:"column:money_end"`
}

// TableName PayWithdrawalCash 表名
func (PayWithdrawalCash) TableName() string {
	return "pay_withdrawal_cash"
}
