package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autoCodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
)

type ContactService struct {
}

// CreateContact 创建Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService) CreateContact(contact autocode.Contact) (err error) {
	err = global.GVA_DB.Create(&contact).Error
	return err
}

// DeleteContact 删除Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService)DeleteContact(contact autocode.Contact) (err error) {
	err = global.GVA_DB.Delete(&contact).Error
	return err
}

// DeleteContactByIds 批量删除Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService)DeleteContactByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]autocode.Contact{},"id in ?",ids.Ids).Error
	return err
}

// UpdateContact 更新Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService)UpdateContact(contact autocode.Contact) (err error) {
	err = global.GVA_DB.Save(&contact).Error
	return err
}

// GetContact 根据id获取Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService)GetContact(id uint) (err error, contact autocode.Contact) {
	err = global.GVA_DB.Where("id = ?", id).First(&contact).Error
	return
}

// GetContactInfoList 分页获取Contact记录
// Author [piexlmax](https://github.com/piexlmax)
func (contactService *ContactService)GetContactInfoList(info autoCodeReq.ContactSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
    // 创建db
	db := global.GVA_DB.Model(&autocode.Contact{})
    var contacts []autocode.Contact
    // 如果有条件搜索 下方会自动创建搜索语句
    if info.Name != "" {
        db = db.Where("name LIKE ?","%"+ info.Name+"%")
    }
    if info.Phone != "" {
        db = db.Where("phone LIKE ?","%"+ info.Phone+"%")
    }
    if info.Wechat != "" {
        db = db.Where("wechat LIKE ?","%"+ info.Wechat+"%")
    }
	err = db.Count(&total).Error
	if err!=nil {
    	return
    }
	err = db.Limit(limit).Offset(offset).Find(&contacts).Error
	return err, contacts, total
}
