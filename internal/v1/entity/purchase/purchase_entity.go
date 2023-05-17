package purchase

import (
	model "eirc.app/internal/v1/structure/purchases"
)

func (e *entity) Created(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Create(&input).Error

	return err
}

func (e *entity) List(input *model.Fields) (amount int64, output []*model.Table, err error) {
	db := e.db.Model(&model.Table{})

	if input.Applicant != nil {
		db.Where("applicant = ?", input.Applicant)
	}

	if input.Company_name != nil {
		db.Where("company_name like %?%", *input.Company_name)
	}

	if input.Requisition_department != nil {
		db.Where("requisition_department = ?", input.Requisition_department)
	}

	if input.Product != nil {
		db.Where("product = ?", input.Product)
	}

	if input.Requisition_quantity != nil {
		db.Where("requisition_quantity = ?", input.Requisition_quantity)
	}

	if input.Price != nil {
		db.Where("price = ?", input.Price)
	}

	if input.Is_Deleted != nil {
		db.Where("is_deleted = ?", input.Is_Deleted)
	}

	err = db.Count(&amount).Offset(int((input.Page - 1) * input.Limit)).
		Limit(int(input.Limit)).Order("created_at desc").Find(&output).Error

	return amount, output, err
}

func (e *entity) GetByID(input *model.Field) (output *model.Table, err error) {
	db := e.db.Model(&model.Table{}).Where("purchase_id = ?", input.Purchase_ID)
	if input.Is_Deleted != nil {
		db.Where("is_deleted = ?", input.Is_Deleted)
	}

	err = db.First(&output).Error

	return output, err
}

func (e *entity) Deleted(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Delete(&input).Error

	return err
}

func (e *entity) Updated(input *model.Table) (err error) {
	err = e.db.Model(&model.Table{}).Where("purchase_id = ?", input.Purchase_ID).Save(&input).Error

	return err
}
