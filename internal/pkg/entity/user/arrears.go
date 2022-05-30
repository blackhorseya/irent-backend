package user

import (
	"github.com/blackhorseya/irent/pb"
)

// ArrearsRecord declare arrears record
type ArrearsRecord struct {
	OrderNo     string `json:"order_no"`
	TotalAmount int    `json:"total_amount"`
}

// NewArrearsRecordResponse return *pb.ArrearsRecord
func NewArrearsRecordResponse(from *ArrearsRecord) *pb.ArrearsRecord {
	return &pb.ArrearsRecord{
		OrderNo:     from.OrderNo,
		TotalAmount: int32(from.TotalAmount),
	}
}

// Arrears declare arrears struct
type Arrears struct {
	Records     []*ArrearsRecord `json:"records"`
	TotalAmount int              `json:"total_amount"`
}

// NewArrearsResponse return *pb.Arrears
func NewArrearsResponse(from *Arrears) *pb.Arrears {
	var records []*pb.ArrearsRecord
	for _, record := range from.Records {
		records = append(records, NewArrearsRecordResponse(record))
	}

	return &pb.Arrears{
		Records:     records,
		TotalAmount: int32(from.TotalAmount),
	}
}
