package purchase

import (
	"net/http"

	"eirc.app/internal/pkg/code"
	"eirc.app/internal/pkg/log"
	"eirc.app/internal/pkg/util"
	preset "eirc.app/internal/v1/presenter"
	purchases "eirc.app/internal/v1/structure/purchases"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Created
// @Summary 新增使用者
// @description 新增使用者
// @Tags purchase
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param * body purchases.Created true "新增使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases [post]
func (p *presenter) Created(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	trx := ctx.MustGet("db_trx").(*gorm.DB)
	createBy := util.GenerateUUID()
	input := &purchases.Created{}
	input.Created_By = createBy
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PurchaseResolver.Created(trx, input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// List
// @Summary 條件搜尋使用者
// @description 條件使用者
// @Tags purchase
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param organizationID query string false "組織ID"
// @param purchase query string false "帳號"
// @param chineseName query string false "中文名稱"
// @param roleName query string false "角色名稱"
// @param page query int true "目前頁數,請從1開始帶入"
// @param limit query int true "一次回傳比數,請從1開始帶入,最高上限20"
// @success 200 object code.SuccessfulMessage{body=purchases.List} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases [get]
func (p *presenter) List(ctx *gin.Context) {
	input := &purchases.Fields{}
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	if input.Limit >= preset.DefaultLimit {
		input.Limit = preset.DefaultLimit
	}

	codeMessage := p.PurchaseResolver.List(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// GetByID
// @Summary 取得單一使用者
// @description 取得單一使用者
// @Tags purchase
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param Purchase_ID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=purchases.Single} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases/{Purchase_ID} [get]
func (p *presenter) GetByID(ctx *gin.Context) {
	Purchase_ID := ctx.Param("Purchase_ID")
	input := &purchases.Field{}
	input.Purchase_ID = Purchase_ID
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PurchaseResolver.GetByID(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Delete
// @Summary 刪除單一使用者
// @description 刪除單一使用者
// @Tags purchase
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param Purchase_ID path string true "使用者ID"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases/{Purchase_ID} [delete]
func (p *presenter) Delete(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	Purchase_ID := ctx.Param("purchase_ID")
	input := &purchases.Updated{}
	input.Purchase_ID = Purchase_ID
	input.Updated_By = util.PointerString(updatedBy)
	if err := ctx.ShouldBindQuery(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PurchaseResolver.Deleted(input)
	ctx.JSON(http.StatusOK, codeMessage)
}

// Updated
// @Summary 更新單一使用者
// @description 更新單一使用者
// @Tags purchase
// @version 1.0
// @Accept json
// @produce json
// @param Authorization header string  true "JWE Token"
// @param Purchase_ID path string true "使用者ID"
// @param * body purchases.Updated true "更新使用者"
// @success 200 object code.SuccessfulMessage{body=string} "成功後返回的值"
// @failure 415 object code.ErrorMessage{detailed=string} "必要欄位帶入錯誤"
// @failure 500 object code.ErrorMessage{detailed=string} "伺服器非預期錯誤"
// @Router /v1.0/authority/purchases/{Purchase_ID} [patch]
func (p *presenter) Updated(ctx *gin.Context) {
	// Todo 將UUID改成登入的使用者
	updatedBy := util.GenerateUUID()
	Purchase_ID := ctx.Param("purchase_ID")
	input := &purchases.Updated{}
	input.Purchase_ID = Purchase_ID
	input.Updated_By = util.PointerString(updatedBy)
	if err := ctx.ShouldBindJSON(input); err != nil {
		log.Error(err)
		ctx.JSON(http.StatusOK, code.GetCodeMessage(code.FormatError, err.Error()))

		return
	}

	codeMessage := p.PurchaseResolver.Updated(input)
	ctx.JSON(http.StatusOK, codeMessage)
}
