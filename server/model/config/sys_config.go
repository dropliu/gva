// 自动生成模板SysConfig
package config

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// SysConfig 结构体
type SysConfig struct {
	global.GVA_MODEL
	Name    string `json:"name" form:"name" gorm:"column:name;comment:名称;size:255;"`
	Key     string `json:"key" form:"key" gorm:"column:key_name;comment:键;size:255;"`
	Value   string `json:"value" form:"value" gorm:"column:value;comment:值;size:255;"`
	Comment string `json:"comment" form:"comment" gorm:"column:comment;comment:描述;size:255;"`
}

// TableName SysConfig 表名
func (SysConfig) TableName() string {
	return "sys_config"
}
