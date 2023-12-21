package dto

type RequestContainer struct {
	OrderDate string `json:"order_date"`
	InvoiceNo string `json:"invoice_no"`
	Sku       string `json:"sku"`
	Qty       int    `json:"qty"`
	Price     int    `json:"Price"`
	IsPayment bool   `json:"is_payment"`
}

type RequestSo struct {
	MarketplaceId string `json:"marketplace_id" validate:"required"`
	StartDate     string `json:"start_date" validate:"required"`
	EndDate       string `json:"end_date" validate:"required"`
}

type RequestPo struct {
	PoNumber    string  `json:"po_number" validate:"required"`
	PoDate      string  `json:"po_date" validate:"required"`
	Sku         string  `json:"sku" validate:"required"`
	ProductName string  `json:"product_name"`
	Qty         int     `json:"qty"`
	Price       float64 `json:"price"`
	Discount    float64 `json:"discount"`
	Ppn         float64 `json:"ppn"`
	TotalPrice  float64 `json:"total_price"`
}
