// 自动生成模板PayAgentBonusRatio
package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	
)

// PayAgentBonusRatio 结构体
type PayAgentBonusRatio struct {
      global.GVA_MODEL
      Uuid  string `json:"uuid" form:"uuid" gorm:"column:uuid;comment:用户id;size:191;"`
      Value  *float64 `json:"value" form:"value" gorm:"column:value;comment:;size:22;"`
}


// TableName PayAgentBonusRatio 表名
func (PayAgentBonusRatio) TableName() string {
  return "pay_agent_bonus_ratio"
}

