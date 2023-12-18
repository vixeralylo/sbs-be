package entity

const (
	TABLE_PRODUCT     = "sbs_product"
	TABLE_SALES_ORDER = "sbs_sales_order"
)

type Tabler interface {
	TableName() string
}

func (SbsProduct) TableName() string {
	return TABLE_PRODUCT
}

func (SbsSalesOrder) TableName() string {
	return TABLE_SALES_ORDER
}
