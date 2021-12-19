// 自动生成模板Contact
package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Contact 结构体
// 如果含有time.Time 请自行import time包
type Contact struct {
      global.GVA_MODEL
      Name  string `json:"name" form:"name" gorm:"column:name;comment:;"`
      Phone  string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
      Wechat  string `json:"wechat" form:"wechat" gorm:"column:wechat;comment:;"`
}


// TableName Contact 表名
func (Contact) TableName() string {
  return "contact"
}

