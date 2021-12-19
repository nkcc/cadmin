package autocode

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
    "github.com/flipped-aurora/gin-vue-admin/server/model/autocode"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
    autocodeReq "github.com/flipped-aurora/gin-vue-admin/server/model/autocode/request"
    "github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
    "github.com/flipped-aurora/gin-vue-admin/server/service"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
)

type ContactApi struct {
}

var contactService = service.ServiceGroupApp.AutoCodeServiceGroup.ContactService


// CreateContact 创建Contact
// @Tags Contact
// @Summary 创建Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Contact true "创建Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contact/createContact [post]
func (contactApi *ContactApi) CreateContact(c *gin.Context) {
	var contact autocode.Contact
	_ = c.ShouldBindJSON(&contact)
	if err := contactService.CreateContact(contact); err != nil {
        global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteContact 删除Contact
// @Tags Contact
// @Summary 删除Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Contact true "删除Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /contact/deleteContact [delete]
func (contactApi *ContactApi) DeleteContact(c *gin.Context) {
	var contact autocode.Contact
	_ = c.ShouldBindJSON(&contact)
	if err := contactService.DeleteContact(contact); err != nil {
        global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteContactByIds 批量删除Contact
// @Tags Contact
// @Summary 批量删除Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /contact/deleteContactByIds [delete]
func (contactApi *ContactApi) DeleteContactByIds(c *gin.Context) {
	var IDS request.IdsReq
    _ = c.ShouldBindJSON(&IDS)
	if err := contactService.DeleteContactByIds(IDS); err != nil {
        global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateContact 更新Contact
// @Tags Contact
// @Summary 更新Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body autocode.Contact true "更新Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /contact/updateContact [put]
func (contactApi *ContactApi) UpdateContact(c *gin.Context) {
	var contact autocode.Contact
	_ = c.ShouldBindJSON(&contact)
	if err := contactService.UpdateContact(contact); err != nil {
        global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindContact 用id查询Contact
// @Tags Contact
// @Summary 用id查询Contact
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocode.Contact true "用id查询Contact"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /contact/findContact [get]
func (contactApi *ContactApi) FindContact(c *gin.Context) {
	var contact autocode.Contact
	_ = c.ShouldBindQuery(&contact)
	if err, recontact := contactService.GetContact(contact.ID); err != nil {
        global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recontact": recontact}, c)
	}
}

// GetContactList 分页获取Contact列表
// @Tags Contact
// @Summary 分页获取Contact列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.ContactSearch true "分页获取Contact列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /contact/getContactList [get]
func (contactApi *ContactApi) GetContactList(c *gin.Context) {
	var pageInfo autocodeReq.ContactSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := contactService.GetContactInfoList(pageInfo); err != nil {
	    global.GVA_LOG.Error("获取失败!", zap.Error(err))
        response.FailWithMessage("获取失败", c)
    } else {
        response.OkWithDetailed(response.PageResult{
            List:     list,
            Total:    total,
            Page:     pageInfo.Page,
            PageSize: pageInfo.PageSize,
        }, "获取成功", c)
    }
}
