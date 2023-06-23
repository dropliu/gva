package service

import (
	"errors"
	"fmt"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model"
	"gorm.io/gorm"
)

type Payment struct {
	ID         int64          `json:"id" gorm:"column:id"`
	Name       string         `json:"name" gorm:"column:name"`
	Identifier string         `json:"identifier" gorm:"column:identifier"`
	Desc       string         `json:"desc" gorm:"column:desc"`
	CreatedAt  time.Time      `json:"created_at" gorm:"column:created_at"`
	UpdatedAt  time.Time      `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"` // 删除时间
}

// ListAllPayment 获取所有支付方式
func ListAllPayment() ([]Payment, error) {
	var list []Payment
	return list, global.GVA_DB.Table(model.TablePayment).Find(&list).Error
}

func CreatePayment(req Payment) error {
	db := global.GVA_DB.Table(model.TablePayment).Select("name", "identifier", "desc",
		"created_at", "updated_at")

	// 检查是否有相同记录
	var obj Payment
	err := db.Where("identifier=?", req.Identifier).First(&obj).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return err
	}
	db.Error = nil
	if obj.ID > 0 {
		return fmt.Errorf("重复的标识符:'%s'", obj.Identifier)
	}

	err1 := db.Create(&req).Error
	// fmt.Println("create", err1.Error())
	return err1
}

func UpdatePayment(req Payment) error {
	db := global.GVA_DB.Table(model.TablePayment).Select("name", "identifier", "desc", "updated_at")

	return db.Where("id=?", req.ID).Updates(&req).Error
}

func DeletePayment(req Payment) error {
	db := global.GVA_DB.Table(model.TablePayment).Select("deleted_at")

	return db.Where("id=?", req.ID).Updates(&req).Error
}
