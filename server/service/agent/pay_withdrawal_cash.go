package agent

import (
	"strings"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	agentReq "github.com/flipped-aurora/gin-vue-admin/server/model/agent/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/withdrawalcash"
)

type PayWithdrawalCashService struct {
}

// CreatePayWithdrawalCash 创建PayWithdrawalCash记录
// Author [piexlmax](https://github.com/piexlmax)
func (payWithdrawalCashService *PayWithdrawalCashService) CreatePayWithdrawalCash(payWithdrawalCash *withdrawalcash.PayWithdrawalCash) (err error) {
	err = global.GVA_DB.Debug().Create(payWithdrawalCash).Error
	return err
}

// DeletePayWithdrawalCash 删除PayWithdrawalCash记录
// Author [piexlmax](https://github.com/piexlmax)
func (payWithdrawalCashService *PayWithdrawalCashService) DeletePayWithdrawalCash(payWithdrawalCash withdrawalcash.PayWithdrawalCash) (err error) {
	err = global.GVA_DB.Delete(&payWithdrawalCash).Error
	return err
}

// DeletePayWithdrawalCashByIds 批量删除PayWithdrawalCash记录
// Author [piexlmax](https://github.com/piexlmax)
func (payWithdrawalCashService *PayWithdrawalCashService) DeletePayWithdrawalCashByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]withdrawalcash.PayWithdrawalCash{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePayWithdrawalCash 更新PayWithdrawalCash记录
// Author [piexlmax](https://github.com/piexlmax)
func (payWithdrawalCashService *PayWithdrawalCashService) UpdatePayWithdrawalCash(payWithdrawalCash withdrawalcash.PayWithdrawalCash) (err error) {
	err = global.GVA_DB.Save(&payWithdrawalCash).Error
	return err
}

// GetPayWithdrawalCash 根据id获取PayWithdrawalCash记录
// Author [piexlmax](https://github.com/piexlmax)
func (payWithdrawalCashService *PayWithdrawalCashService) GetPayWithdrawalCash(id uint) (payWithdrawalCash withdrawalcash.PayWithdrawalCash, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&payWithdrawalCash).Error
	return
}

// GetPayWithdrawalCashInfoList 分页获取PayWithdrawalCash记录
// Author [piexlmax](https://github.com/piexlmax)
func (payWithdrawalCashService *PayWithdrawalCashService) GetPayWithdrawalCashInfoList(info agentReq.PayWithdrawalCashSearch, uuid string) (list []withdrawalcash.PayWithdrawalCash, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&withdrawalcash.PayWithdrawalCash{})
	var payWithdrawalCashs []withdrawalcash.PayWithdrawalCash
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if uuid != "" {
		db = db.Where("uuid = ?", uuid)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Preload("User").Find(&payWithdrawalCashs).Error
	return payWithdrawalCashs, total, err
}

// 获取金额
func (payWithdrawalCashService *PayWithdrawalCashService) GetPayWithdrawalCashStatic(uuid string) (paidNum float64, num float64, err error) {
	db := global.GVA_DB.Model(&withdrawalcash.PayWithdrawalCash{})
	if uuid != "" {
		db = db.Where("uuid = ?", uuid)
	}
	err = db.Debug().Select("SUM(money)").Where("status = ?", 1).Row().Scan(&paidNum)
	if err != nil {
		if strings.Contains(err.Error(), "Scan error on column index 0,") {
			err = nil
		}
	}

	err = db.Debug().Select("SUM(money)").Row().Scan(&num)
	if err != nil {
		if strings.Contains(err.Error(), "Scan error on column index 0,") {
			err = nil
		}
	}
	return
}
