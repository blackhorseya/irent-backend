package user

// ArrearsRecord declare arrears record
type ArrearsRecord struct {
	OrderNo     string `json:"order_no"`
	TotalAmount int    `json:"total_amount"`
}

type Arrears struct {
	Records     []*ArrearsRecord `json:"records"`
	TotalAmount int              `json:"total_amount"`
}