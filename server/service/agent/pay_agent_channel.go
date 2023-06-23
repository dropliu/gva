package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/agent"
	agentReq "github.com/flipped-aurora/gin-vue-admin/server/model/agent/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PayAgentChannelService struct {
}

// CreatePayAgentChannel 创建PayAgentChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentChannelService *PayAgentChannelService) CreatePayAgentChannel(payAgentChannel *agent.PayAgentChannel) (err error) {
	err = global.GVA_DB.Create(payAgentChannel).Error
	return err
}

// DeletePayAgentChannel 删除PayAgentChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentChannelService *PayAgentChannelService) DeletePayAgentChannel(payAgentChannel agent.PayAgentChannel) (err error) {
	err = global.GVA_DB.Delete(&payAgentChannel).Error
	return err
}

// DeletePayAgentChannelByIds 批量删除PayAgentChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentChannelService *PayAgentChannelService) DeletePayAgentChannelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]agent.PayAgentChannel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePayAgentChannel 更新PayAgentChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentChannelService *PayAgentChannelService) UpdatePayAgentChannel(payAgentChannel agent.PayAgentChannel) (err error) {
	err = global.GVA_DB.Save(&payAgentChannel).Error
	return err
}

// GetPayAgentChannel 根据id获取PayAgentChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentChannelService *PayAgentChannelService) GetPayAgentChannel(id uint) (payAgentChannel agent.PayAgentChannel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&payAgentChannel).Error
	return
}

// GetPayAgentChannelInfoList 分页获取PayAgentChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payAgentChannelService *PayAgentChannelService) GetPayAgentChannelInfoList(info agentReq.PayAgentChannelSearch) (list []agent.PayAgentChannel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&agent.PayAgentChannel{})
	var payAgentChannels []agent.PayAgentChannel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.Uuid != "" {
		db = db.Where("uuid = ?", info.Uuid)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&payAgentChannels).Error
	return payAgentChannels, total, err
}
