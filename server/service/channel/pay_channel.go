package channel

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/channel"
	channelReq "github.com/flipped-aurora/gin-vue-admin/server/model/channel/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
)

type PayChannelService struct {
}

// CreatePayChannel 创建PayChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payChannelService *PayChannelService) CreatePayChannel(payChannel *channel.PayChannel) (err error) {
	err = global.GVA_DB.Create(payChannel).Error
	return err
}

// DeletePayChannel 删除PayChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payChannelService *PayChannelService) DeletePayChannel(payChannel channel.PayChannel) (err error) {
	err = global.GVA_DB.Delete(&payChannel).Error
	return err
}

// DeletePayChannelByIds 批量删除PayChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payChannelService *PayChannelService) DeletePayChannelByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]channel.PayChannel{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePayChannel 更新PayChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payChannelService *PayChannelService) UpdatePayChannel(payChannel channel.PayChannel) (err error) {
	err = global.GVA_DB.Save(&payChannel).Error
	return err
}

// GetPayChannel 根据id获取PayChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payChannelService *PayChannelService) GetPayChannel(id uint) (payChannel channel.PayChannel, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&payChannel).Error
	return
}

// GetPayChannelInfoList 分页获取PayChannel记录
// Author [piexlmax](https://github.com/piexlmax)
func (payChannelService *PayChannelService) GetPayChannelInfoList(info channelReq.PayChannelSearch) (list []channel.PayChannel, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&channel.PayChannel{})
	var payChannels []channel.PayChannel
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&payChannels).Error
	return payChannels, total, err
}

// 获取通道
func (payChannelService *PayChannelService) GetPayChannelInfoByUserIdentifier(channelIdentifier, identifier string) (payChannel channel.PayChannel, err error) {
	err = global.GVA_DB.Where("channel_identifier = ? AND identifier = ?", channelIdentifier, identifier).First(&payChannel).Error
	return
}
