package dto

type RequestContainer struct {
	InvoiceNo string `json:"invoice_no"`
	Sku       string `json:"sku"`
	Qty       int    `json:"qty"`
	Price     int    `json:"Price"`
}
