package entity

type SbsProduct struct {
	Sku         string  `json:"sku" gorm:"column:sku;primary_key;"`
	ProductName string  `json:"product_name" gorm:"column:product_name;"`
	Stock       string  `json:"stock" gorm:"column:stock;"`
	Hpp         int     `json:"hpp" gorm:"column:hpp;"`
	Price       float64 `json:"price" gorm:"column:price;"`
	BrandId     string  `json:"brand_id" gorm:"column:brand_id;"`
	SupplierId  string  `json:"supplier_id" gorm:"column:supplier_id;"`
	Seq         string  `json:"seq" gorm:"column:seq;"`
	Gross       float64 `json:"gross"`
	Admin       float64 `json:"admin"`
	Ongkir      float64 `json:"ongkir"`
	CleanMargin float64 `json:"clean_margin"`
	Pct         float64 `json:"pct"`
}

type SbsSalesOrder struct {
	MarketPlaceId    string  `json:"marketplace_id" gorm:"column:marketplace_id;primary_key;"`
	OrderDate        string  `json:"order_date" gorm:"column:order_date;primary_key;"`
	InvoiceNo        string  `json:"invoice_no" gorm:"column:invoice_no;primary_key;"`
	Sku              string  `json:"sku" gorm:"column:sku;primary_key;"`
	ProductName      string  `json:"product_name" gorm:"column:product_name;"`
	Qty              int     `json:"qty" gorm:"column:qty;"`
	SalesPrice       int     `json:"sales_price" gorm:"column:sales_price;"`
	TotalPrice       int     `json:"total_price" gorm:"column:total_price;"`
	Hpp              int     `json:"hpp" gorm:"column:hpp;"`
	GrossMargin      int     `json:"gross_margin" gorm:"column:gross_margin;"`
	PowerMerchantFee float32 `json:"power_merchant_fee" gorm:"column:power_merchant_fee;"`
	OngkirFee        float32 `json:"ongkir_fee" gorm:"column:ongkir_fee;"`
	CleanMargin      float32 `json:"clean_margin" gorm:"column:clean_margin;"`
}

type SbsPurchaseOrder struct {
	PoNumber    string  `json:"po_number" gorm:"column:po_number;primary_key;"`
	PoDate      string  `json:"po_date" gorm:"column:po_date;primary_key;"`
	Sku         string  `json:"sku" gorm:"column:sku;primary_key;"`
	ProductName string  `json:"product_name" gorm:"column:product_name;"`
	Qty         int     `json:"qty" gorm:"column:qty;"`
	Price       float64 `json:"price" gorm:"column:price;"`
	Discount    float64 `json:"discount" gorm:"column:discount;"`
	Ppn         float64 `json:"ppn" gorm:"column:ppn;"`
	TotalPrice  float64 `json:"total_price" gorm:"column:total_price;"`
}
