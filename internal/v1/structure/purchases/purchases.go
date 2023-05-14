package purchase

import (
	"time"

	model "eirc.app/internal/v1/structure"
)

// Table struct is database table struct
type Table struct {
	// 編號
	Purchase_ID string `gorm:"primaryKey;column:purchase_id;uuid_generate_v4()type:UUID;" json:"purchase_id,omitempty"`
	// 申請人
	Applicant string `gorm:"column:applicant;type:VARCHAR;" json:"applicant,omitempty"`
	// 公司名稱
	Company_name string `gorm:"column:company_name;type:VARCHAR;" json:"company_name,omitempty"`
	// 請購部門
	Requisition_department string `gorm:"column:requisition_department;type:VARCHAR;" json:"requisition_department,omitempty"`
	// 品名
	Product string `gorm:"column:product;type:VARCHAR;" json:"product,omitempty"`
	// 請購數量
	Requisition_quantity int64 `gorm:"column:requisition_quantity;type:INT4;" json:"requisition_quantity,omitempty"`
	// 單價
	Price int64 `gorm:"column:price;type:INT4;" json:"price,omitempty"`
	//流水號
	Position string `gorm:"->;column:position;type:VARCHAR;" json:"position,omitempty"`
	// 是否刪除
	Is_Deleted bool `gorm:"column:is_deleted;type:bool;false" json:"is_deleted,omitempty"`
	// 創建時間
	Created_At time.Time `gorm:"column:created_at;type:TIMESTAMP;" json:"created_at"`
	// 創建者
	Created_By string `gorm:"column:created_by;type:VARCHAR;" json:"created_by,omitempty"`
	// 更新時間
	Updated_At *time.Time `gorm:"column:updated_at;type:TIMESTAMP;" json:"updated_at,omitempty"`
	// 更新者
	Updated_By *string `gorm:"column:updated_by;type:VARCHAR;" json:"updated_by,omitempty"`
}

// Base struct is corresponding to table structure file
type Base struct {
	// 編號
	Purchase_ID string `json:"purchase_id,omitempty"`
	// 申請人
	Applicant string `json:"applicant,omitempty"`
	// 公司名稱
	Company_name string `json:"company_name,omitempty"`
	// 請購部門
	Requisition_department string `json:"requisition_department,omitempty"`
	// 品名
	Product string `json:"product,omitempty"`
	// 請購數量
	Requisition_quantity int64 `json:"requisition_quantity,omitempty"`
	// 單價
	Price int64 `json:"price,omitempty"`
	//流水號
	Position string `json:"position,omitempty"`
	// 是否刪除
	Is_Deleted bool `json:"is_deleted,omitempty"`
	// 創建時間
	Created_At time.Time `json:"created_at"`
	// 創建者
	Created_By string `json:"created_by,omitempty"`
	// 更新時間
	Updated_At *time.Time `json:"updated_at,omitempty"`
	// 更新者
	Updated_By *string `json:"updated_by,omitempty"`
}

// Single return structure file
type Single struct {
	// 編號
	Purchase_ID string `json:"purchase_id,omitempty"`
	// 申請人
	Applicant string `json:"applicant,omitempty"`
	// 公司名稱
	Company_name string `json:"company_name,omitempty"`
	// 請購部門
	Requisition_department string `json:"requisition_department,omitempty"`
	// 品名
	Product string `json:"product,omitempty"`
	// 請購數量
	Requisition_quantity int64 `json:"requisition_quantity,omitempty"`
	// 單價
	Price int64 `json:"price,omitempty"`
	//流水號
	Position string `json:"position,omitempty"`
	// 創建時間
	Created_At time.Time `json:"created_at"`
	// 創建者
	Created_By string `json:"created_by,omitempty"`
	// 更新時間
	Updated_At *time.Time `json:"updated_at,omitempty"`
	// 更新者
	Updated_By *string `json:"updated_by,omitempty"`
}

// Created struct is used to create
type Created struct {
	// 申請人
	Applicant string `json:"applicant" binding:"required" validate:"required"`
	// 公司名稱
	Company_name string `json:"company_name,omitempty" binding:"required" validate:"required"`
	// 請購部門
	Requisition_department string `json:"requisition_department" binding:"required" validate:"required"`
	// 品名
	Product string `json:"product" binding:"required" validate:"required"`
	// 請購數量
	Requisition_quantity int64 `json:"requisition_quantity" binding:"required" validate:"required"`
	// 單價
	Price int64 `json:"price" binding:"required" validate:"required"`
	// 創建者
	Created_By string `json:"created_by" swaggerignore:"true"`
}

// Field is structure file for search
type Field struct {
	// 編號
	Purchase_ID string `json:"purchase_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 申請人
	Applicant *string `json:"applicant,omitempty" form:"applicant"`
	// 公司名稱
	Company_name *string `json:"company_name,omitempty" form:"company_name"`
	// 請購部門
	Requisition_department *string `json:"requisition_department,omitempty" form:"requisition_department"`
	// 品名
	Product *string `json:"product,omitempty" form:"product"`
	// 請購數量
	Requisition_quantity *int64 `json:"requisition_quantity,omitempty" form:"requisition_quantity"`
	// 單價
	Price *int64 `json:"price,omitempty" form:"price"`
	// 是否刪除
	Is_Deleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// Fields is the searched structure file (including pagination)
type Fields struct {
	Field
	model.InPage
}

// List is multiple return structure files
type List struct {
	Purchase []*struct {
		// 編號
		Purchase_ID string `json:"purchase_id,omitempty"`
		// 申請人
		Applicant string `json:"applicant,omitempty"`
		// 公司名稱
		Company_Name string `json:"company_name,omitempty"`
		// 請購部門
		Requisition_department string `json:"requisition_department,omitempty"`
		// 品名
		Product string `json:"product,omitempty"`
		// 請購數量
		Requisition_quantity int64 `json:"requisition_quantity,omitempty"`
		// 單價
		Price int64 `json:"price,omitempty"`
		//流水號
		Position string `json:"position,omitempty"`
		// 創建者
		Created_By string `json:"created_by,omitempty"`
	} `json:"purchases"`
	model.OutPage
}

// Updated struct is used to update
type Updated struct {
	// 編號
	Purchase_ID string `json:"purchase_id,omitempty" binding:"omitempty,uuid4" swaggerignore:"true"`
	// 公司名稱
	Company_Name *string `json:"company_name,omitempty"`
	// 申請人
	Applicant *string `json:"applicant,omitempty"`
	// 請購部門
	Requisition_department *string `json:"requisition_department,omitempty"`
	// 品名
	Product *string `json:"product,omitempty"`
	// 請購數量
	Requisition_quantity *int64 `json:"requisition_quantity,omitempty"`
	// 單價
	Price *int64 `json:"price,omitempty"`
	// 更新者
	Updated_By *string `json:"updated_by,omitempty" swaggerignore:"true"`
	// 是否刪除
	Is_Deleted *bool `json:"is_deleted,omitempty" swaggerignore:"true"`
}

// TableName sets the insert table name for this struct type
func (a *Table) TableName() string {
	return "purchases"
}
