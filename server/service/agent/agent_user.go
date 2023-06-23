package agent

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/system"
)

type PayAgentUserService struct {
}

// 获取我名下的代理
func (self *PayAgentUserService) GetAgentUser(info request.PageInfo, uuid string) (list interface{}, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&system.SysUser{})

	db = db.Where("authority_id = ?", 100)

	if uuid != "" {
		db = db.Where("parent_uid = ?", uuid)
	}

	var userList []system.SysUser
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	err = db.Debug().Limit(limit).Offset(offset).Preload("BonusRatio").Find(&userList).Error
	return userList, total, err
}
