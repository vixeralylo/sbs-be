package entity

const (
	TABLE_PRODUCT        = "ibg_product"
	TABLE_SALES_ORDER    = "ibg_sales_order"
	TABLE_PURCHASE_ORDER = "ibg_purchase_order"
	TABLE_COST           = "ibg_cost"
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
