// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: transfer.proto

package rusk

import (
	fmt "fmt"
	math "math"
	proto "github.com/gogo/protobuf/proto"
	context "context"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type TransferMock struct{}

func (m *TransferMock) NewTransfer(ctx context.Context, req *TransferTransactionRequest) (*Transaction, error) {
	res :=
		&Transaction{
			Version: 255,
			Type:    469,
		}
	return res, nil
}
