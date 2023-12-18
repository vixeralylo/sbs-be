package dto

type RequestContainer struct {
	OrderDate string `json:"order_date"`
	InvoiceNo string `json:"invoice_no"`
	Sku       string `json:"sku"`
	Qty       int    `json:"qty"`
	Price     int    `json:"Price"`
}

type RequestSo struct {
	MarketplaceId string `json:"marketplace_id" validate:"required"`
	StartDate     string `json:"start_date" validate:"required"`
	EndDate       string `json:"end_date" validate:"required"`
}
