package config

import (
	"fmt"
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/config"
	configReq "github.com/flipped-aurora/gin-vue-admin/server/model/config/request"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

type SysConfigService struct {
}

// CreateSysConfig 创建SysConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) CreateSysConfig(sysConfig *config.SysConfig) (err error) {
	err = global.GVA_DB.Create(sysConfig).Error
	return err
}

// DeleteSysConfig 删除SysConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) DeleteSysConfig(sysConfig config.SysConfig) (err error) {
	err = global.GVA_DB.Delete(&sysConfig).Error
	return err
}

// DeleteSysConfigByIds 批量删除SysConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) DeleteSysConfigByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]config.SysConfig{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateSysConfig 更新SysConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) UpdateSysConfig(sysConfig config.SysConfig) (err error) {
	err = global.GVA_DB.Save(&sysConfig).Error
	return err
}

// GetSysConfig 根据id获取SysConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) GetSysConfig(id uint) (sysConfig config.SysConfig, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&sysConfig).Error
	return
}

// GetSysConfigInfoList 分页获取SysConfig记录
// Author [piexlmax](https://github.com/piexlmax)
func (sysConfigService *SysConfigService) GetSysConfigInfoList(info configReq.SysConfigSearch) (list []config.SysConfig, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&config.SysConfig{})
	var sysConfigs []config.SysConfig
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&sysConfigs).Error
	return sysConfigs, total, err
}

func (sysConfigService *SysConfigService) GetExchangeRate() (r float64, err error) {
	var sysConfig config.SysConfig
	db := global.GVA_DB.Model(&config.SysConfig{})
	err = db.Where("key_name = ?", "ExchangeRate").First(&sysConfig).Error
	if err != nil {
		return
	}

	if sysConfig.Value == "" {
		r = 1.0
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	if err != nil {
		return
	}

	var floatRateConf config.SysConfig

	db1 := global.GVA_DB.Model(&config.SysConfig{})
	err = db1.Where("key_name = ?", "FloatExchangeRate").First(&floatRateConf).Error
	if err != nil {
		return
	}

	floatRate, _ := strconv.ParseFloat(floatRateConf.Value, 64)

	r += floatRate
	return
}

// 更新浮动汇率
func (sysConfigService *SysConfigService) UpdateFloatExchangeRate() (err error) {
	db := global.GVA_DB.Model(&config.SysConfig{})
	floatExchangeRate, err := utils.GetFloatExchangeRate()
	if err != nil {
		return
	}

	err = db.Where("key_name = ?", "FloatExchangeRate").Update("value", fmt.Sprintf("%f", floatExchangeRate)).Error
	return
}

// 交易金额大于该数值，按百分比收取手续费
func (sysConfigService *SysConfigService) GetPercentageNum() (r float64, err error) {
	var sysConfig config.SysConfig
	db := global.GVA_DB.Model(&config.SysConfig{})
	err = db.Where("key_name = ?", "PercentageNum").First(&sysConfig).Error
	if err != nil {
		return
	}

	if sysConfig.Value == "" {
		r = 50
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 交易金额大于收取百分比值
func (sysConfigService *SysConfigService) GetPercentageFee() (r float64, err error) {
	var sysConfig config.SysConfig
	db := global.GVA_DB.Model(&config.SysConfig{})
	err = db.Where("key_name = ?", "PercentageFee").First(&sysConfig).Error
	if err != nil {
		return
	}

	if sysConfig.Value == "" {
		r = 1.0
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 交易金额小于多少值
func (sysConfigService *SysConfigService) GetFixedValueNum() (r float64, err error) {
	var sysConfig config.SysConfig
	db := global.GVA_DB.Model(&config.SysConfig{})
	err = db.Where("key_name = ?", "FixedValueNum").First(&sysConfig).Error
	if err != nil {
		return
	}

	if sysConfig.Value == "" {
		r = 50
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 交易金额小于固定值手续费用
func (sysConfigService *SysConfigService) GetFixedValueFee() (r float64, err error) {
	var sysConfig config.SysConfig
	db := global.GVA_DB.Model(&config.SysConfig{})
	err = db.Where("key_name = ?", "FixedValueFee").First(&sysConfig).Error
	if err != nil {
		return
	}

	if sysConfig.Value == "" {
		r = 0.5
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 获取平台佣金比例
func (sysConfigService *SysConfigService) GetCommissionRate() (r float64, err error) {
	var sysConfig config.SysConfig
	db := global.GVA_DB.Model(&config.SysConfig{})
	err = db.Where("key_name = ?", "CommissionRate").First(&sysConfig).Error
	if err != nil {
		return
	}

	if sysConfig.Value == "" {
		r = 0.5
		return
	}

	r, err = strconv.ParseFloat(sysConfig.Value, 64)
	return
}

// 计算手续费
func (sysConfigService *SysConfigService) CalculationFee(val float64) (r float64, err error) {
	percentageNum, err := sysConfigService.GetPercentageNum()
	if err != nil {
		return
	}
	if val > percentageNum {
		percentageFee, err := sysConfigService.GetPercentageFee()
		if err != nil {
			return 0, err
		}
		r = val * percentageFee / 100
		return r, nil
	}

	fixedValueNum, err := sysConfigService.GetFixedValueNum()

	if err != nil {
		return
	}

	if val <= fixedValueNum {
		fixedValueFee, err := sysConfigService.GetFixedValueFee()
		if err != nil {
			return 0, err
		}
		return fixedValueFee, nil
	}
	r = 0.5
	return
}
