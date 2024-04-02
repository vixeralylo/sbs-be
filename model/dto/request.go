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
	SoNumber      string `json:"so_number"`
	IsNotPayment  string `json:"is_not_payment"`
}

type RequestPo struct {
	PoNumber     string  `json:"po_number"`
	PoDate       string  `json:"po_date"`
	Sku          string  `json:"sku"`
	StartDate    string  `json:"start_date"`
	EndDate      string  `json:"end_date"`
	ProductName  string  `json:"product_name"`
	Qty          int     `json:"qty"`
	Price        float64 `json:"price"`
	Discount     float64 `json:"discount"`
	Ppn          float64 `json:"ppn"`
	TotalPrice   float64 `json:"total_price"`
	IsNotPayment string  `json:"is_not_payment"`
}

type RequestSoManual struct {
	MarketplaceId string `json:"marketplace_id"`
	OrderDate     string `json:"order_date"`
	Sku           string `json:"sku"`
	Qty           string `json:"qty"`
	Price         string `json:"Price"`
	SoNumber      string `json:"so_number"`
	TotalPrice    string `json:"total_price"`
}

type RequestCost struct {
	CostType      string `json:"cost_type"`
	CostName      string `json:"cost_name"`
	Qty           string `json:"qty"`
	Price         string `json:"price"`
	AddedPrice    string `json:"added_price"`
	TotalPrice    string `json:"total_price"`
	InvoiceNo     string `json:"invoice_no"`
	MarketplaceId string `json:"marketplace_id"`
	Date          string `json:"date"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
}
