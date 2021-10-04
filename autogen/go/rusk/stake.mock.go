// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: stake.proto

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

type StakeServiceMock struct{}

func (m *StakeServiceMock) NewStake(ctx context.Context, req *StakeTransactionRequest) (*Transaction, error) {
	res :=
		&Transaction{
			Version: 842,
			Type:    578,
		}
	return res, nil
}
func (m *StakeServiceMock) FindStake(ctx context.Context, req *FindStakeRequest) (*FindStakeResponse, error) {
	res :=
		&FindStakeResponse{
			Stakes: []*Stake{
				&Stake{},
				&Stake{},
				&Stake{},
				&Stake{},
				&Stake{},
				&Stake{},
				&Stake{},
				&Stake{},
			},
		}
	return res, nil
}
