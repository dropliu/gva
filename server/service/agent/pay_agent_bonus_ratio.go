package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/agent"
	agentReq "github.com/flipped-aurora/gin-vue-admin/server/model/agent/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PayAgentBonusRatioService struct {
}

// CreatePayAgentBonusRatio 创建PayAgentBonusRatio记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentBonusRatioService *PayAgentBonusRatioService) CreatePayAgentBonusRatio(payAgentBonusRatio *agent.PayAgentBonusRatio) (err error) {
	err = global.GVA_DB.Debug().Model(&agent.PayAgentBonusRatio{}).Where("uuid = ?", payAgentBonusRatio.Uuid).Assign(agent.PayAgentBonusRatio{Value: payAgentBonusRatio.Value}).FirstOrCreate(payAgentBonusRatio).Error
	return err
}

// DeletePayAgentBonusRatio 删除PayAgentBonusRatio记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentBonusRatioService *PayAgentBonusRatioService) DeletePayAgentBonusRatio(payAgentBonusRatio agent.PayAgentBonusRatio) (err error) {
	err = global.GVA_DB.Delete(&payAgentBonusRatio).Error
	return err
}

// DeletePayAgentBonusRatioByIds 批量删除PayAgentBonusRatio记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentBonusRatioService *PayAgentBonusRatioService) DeletePayAgentBonusRatioByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]agent.PayAgentBonusRatio{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePayAgentBonusRatio 更新PayAgentBonusRatio记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentBonusRatioService *PayAgentBonusRatioService) UpdatePayAgentBonusRatio(payAgentBonusRatio agent.PayAgentBonusRatio) (err error) {
	err = global.GVA_DB.Save(&payAgentBonusRatio).Error
	return err
}

// GetPayAgentBonusRatio 根据id获取PayAgentBonusRatio记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentBonusRatioService *PayAgentBonusRatioService) GetPayAgentBonusRatio(id uint) (payAgentBonusRatio agent.PayAgentBonusRatio, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&payAgentBonusRatio).Error
	return
}

// GetPayAgentBonusRatioInfoList 分页获取PayAgentBonusRatio记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentBonusRatioService *PayAgentBonusRatioService) GetPayAgentBonusRatioInfoList(info agentReq.PayAgentBonusRatioSearch) (list []agent.PayAgentBonusRatio, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&agent.PayAgentBonusRatio{})
	var payAgentBonusRatios []agent.PayAgentBonusRatio
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&payAgentBonusRatios).Error
	return payAgentBonusRatios, total, err
}

func (payAgentBonusRatioService *PayAgentBonusRatioService) GetPayAgentBonusRatioByUuid(uuid string) (payAgentBonusRatio agent.PayAgentBonusRatio, err error) {
	err = global.GVA_DB.Where("uuid = ?", uuid).First(&payAgentBonusRatio).Error
	return
}
