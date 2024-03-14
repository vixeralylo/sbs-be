package entity

const (
	TABLE_PRODUCT        = "sbs_product"
	TABLE_SALES_ORDER    = "sbs_sales_order"
	TABLE_PURCHASE_ORDER = "sbs_purchase_order"
	TABLE_COST           = "sbs_cost"
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

func (SbsPurchaseOrder) TableName() string {
	return TABLE_PURCHASE_ORDER
}

func (SbsCost) TableName() string {
	return TABLE_COST
}
