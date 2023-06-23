package agent

import (
	"fmt"
	"strings"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/agent"
	agentReq "github.com/flipped-aurora/gin-vue-admin/server/model/agent/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/channel"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type PayOrderService struct {
}

// CreatePayOrder 创建PayOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payOrderService *PayOrderService) CreatePayOrder(payOrder *agent.PayOrder) (err error) {
	err = global.GVA_DB.Create(payOrder).Error
	return err
}

// DeletePayOrder 删除PayOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payOrderService *PayOrderService) DeletePayOrder(payOrder agent.PayOrder) (err error) {
	err = global.GVA_DB.Delete(&payOrder).Error
	return err
}

// DeletePayOrderByIds 批量删除PayOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payOrderService *PayOrderService) DeletePayOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]agent.PayOrder{}, "id in ?", ids.Ids).Error
	return err
}

// UpdatePayOrder 更新PayOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payOrderService *PayOrderService) UpdatePayOrder(payOrder agent.PayOrder) (err error) {
	err = global.GVA_DB.Save(&payOrder).Error
	return err
}

// GetPayOrder 根据id获取PayOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payOrderService *PayOrderService) GetPayOrder(id uint) (payOrder agent.PayOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&payOrder).Error
	return
}

// GetPayOrderInfoList 分页获取PayOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (payOrderService *PayOrderService) GetPayOrderInfoList(info agentReq.PayOrderSearch, uuid string) (
	list []agent.PayOrder, total int64, err error) {

	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&agent.PayOrder{})
	var payOrders []agent.PayOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	// 起始时间
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// 根据uuid查询？
	if uuid != "" {
		db = db.Where("identifier in(select identifier from pay_agent_channel where uuid = ?)", uuid)
	}
	// username
	if info.UserName != "" {
		var reqUser system.SysUser
		err = global.GVA_DB.Model(&system.SysUser{}).Where("username = ?", info.UserName).First(&reqUser).Error
		if err != nil {
			if err.Error() == "record not found" {
				err = nil
			}
			return list, total, err
		}
		db = db.Where("identifier in(select identifier from pay_agent_channel where uuid = ?)", reqUser.UUID.String())
	}

	// 查询记录总数
	if err = db.Count(&total).Error; err != nil {
		return
	}

	err = db.Debug().Limit(limit).Offset(offset).Find(&payOrders).Error

	if err != nil {
		return
	}

	mid := make(map[string]string, 0)
	for _, item := range payOrders {
		mid[item.Identifier] = item.Identifier
	}

	ids := make([]string, 0)

	for _, item := range mid {
		ids = append(ids, item)
	}

	db1 := global.GVA_DB.Model(&system.SysUser{}).Where("uuid in (select uuid from pay_agent_channel where identifier in(?))", ids)
	users := make([]system.SysUser, 0)

	err = db1.Find(&users).Error

	if err != nil {
		return
	}

	db2 := global.GVA_DB.Model(&agent.PayAgentChannel{}).Where("identifier in(?)", ids)

	payAgentChannels := make([]agent.PayAgentChannel, 0)

	err = db2.Find(&payAgentChannels).Error

	if err != nil {
		return
	}

	payChannels := make([]channel.PayChannel, 0)
	err = global.GVA_DB.Model(&channel.PayChannel{}).Where("identifier in(?)", ids).Find(&payChannels).Error
	if err != nil {
		return
	}

	mchannel := make(map[string]channel.PayChannel, 0)

	for _, item := range payChannels {
		mchannel[item.Identifier] = item
	}

	muser := make(map[string]system.SysUser, 0)

	for _, item := range users {
		muser[string(item.UUID.String())] = item
	}

	mAgenChanel := make(map[string]system.SysUser, 0)

	for _, item := range payAgentChannels {
		mAgenChanel[item.Identifier] = muser[item.Uuid]
	}

	for _, item := range payOrders {
		user, ok := mAgenChanel[item.Identifier]
		if ok {
			item.OrderUser = agent.OrderUser{
				Username: user.Username,
				NickName: user.NickName,
			}
		}

		channel, ok1 := mchannel[item.Identifier]
		if ok1 {
			item.ChannelIdentifier = channel.ChannelIdentifier
		}

		if item.ValueIn == 0 {

		}

		list = append(list, item)

	}

	return list, total, err
}

// 获取可以提现的订单
func (payOrderService *PayOrderService) GetAgentCashOrder(uuid string) (list []agent.PayOrder, r float64, err error) {
	db := global.GVA_DB.Model(&agent.PayOrder{})
	var payOrders []agent.PayOrder
	db = db.Where("(status = 'paid' OR status = 'OPENPIX:MOVEMENT_CONFIRMED') AND (cash_time = '' OR cash_time is NULL) AND identifier in(select identifier from pay_agent_channel where uuid = ?)", uuid)
	err = db.Find(&payOrders).Error
	if err != nil {
		return
	}

	for _, item := range payOrders {
		r += *item.Value
	}

	return payOrders, r, err
}

// 更新订单变成已经提现
func (payOrderService *PayOrderService) UpdateAgentCashOrder(list []agent.PayOrder) (err error) {
	ids := make([]uint, 0)
	for _, item := range list {
		ids = append(ids, item.ID)
	}
	db := global.GVA_DB.Model(&agent.PayOrder{})
	err = db.Where("id in(?)", ids).Update("cash_time", time.Now().Format("2006-01-02 15:04:05")).Error
	return
}

// 获取当天的
func (payOrderService *PayOrderService) GetPayOrderStaticDay(uuid string) (moneyAll float64, money float64, total int64, err error) {
	db := global.GVA_DB.Model(&agent.PayOrder{})

	nowDayStr := time.Now().Format("2006-01-02")
	starStr := fmt.Sprintf("%s 00:00:00", nowDayStr)

	endStr := fmt.Sprintf("%s 23:59:59", nowDayStr)

	if uuid != "" {
		db = db.Where("identifier in(select identifier from pay_agent_channel where uuid = ?)", uuid)
	}

	var payOrders []agent.PayOrder
	db = db.Where("created_at >= ? AND created_at <= ? and (status = ? or status = ?)", starStr, endStr, "OPENPIX:MOVEMENT_CONFIRMED", "paid")

	err = db.Find(&payOrders).Error

	if err != nil {
		return
	}

	for _, item := range payOrders {
		moneyAll += item.ValueIn
		money += *item.Value
		total++
	}

	return
}

// 统计
func (payOrderService *PayOrderService) GetPayOrderStatic(uuid string) (moneyAll float64, money float64, err error) {
	db := global.GVA_DB.Model(&agent.PayOrder{})
	if uuid != "" {
		db = db.Where("identifier in(select identifier from pay_agent_channel where uuid = ?)", uuid)
	}

	err = db.Debug().Select("SUM(value)").Where("(status = ? or status = ?)", "OPENPIX:MOVEMENT_CONFIRMED", "paid").Row().Scan(&moneyAll)
	if err != nil {
		if strings.Contains(err.Error(), "Scan error on column index 0,") {
			err = nil
		}
		return
	}
	err = db.Debug().Select("SUM(value)").Where("(status = ? or status = ?) AND cash_time = ''", "OPENPIX:MOVEMENT_CONFIRMED", "paid").Row().Scan(&money)
	if err != nil {
		if strings.Contains(err.Error(), "Scan error on column index 0,") {
			err = nil
		}
		return
	}
	return

}
