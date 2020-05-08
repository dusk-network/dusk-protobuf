// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: node.proto

package node

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

type WalletMock struct{}

func (m *WalletMock) CreateWallet(ctx context.Context, req *CreateRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) LoadWallet(ctx context.Context, req *LoadRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) CreateFromSeed(ctx context.Context, req *CreateRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) ClearWalletDatabase(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "dolor",
		}
	return res, nil
}
func (m *WalletMock) GetWalletStatus(ctx context.Context, req *EmptyRequest) (*WalletStatusResponse, error) {
	res :=
		&WalletStatusResponse{
			Loaded: false,
		}
	return res, nil
}
func (m *WalletMock) GetAddress(ctx context.Context, req *EmptyRequest) (*LoadResponse, error) {
	res :=
		&LoadResponse{
			Key: &PubKey{},
		}
	return res, nil
}
func (m *WalletMock) GetBalance(ctx context.Context, req *EmptyRequest) (*BalanceResponse, error) {
	res :=
		&BalanceResponse{}
	return res, nil
}
func (m *WalletMock) GetTxHistory(ctx context.Context, req *EmptyRequest) (*TxHistoryResponse, error) {
	res :=
		&TxHistoryResponse{
			Records: []*TxRecord{
				&TxRecord{
					Direction: 1,
					Timestamp: 390,
					Type:      3,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 760,
					Type:      5,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 241,
					Type:      1,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 701,
					Type:      2,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 800,
					Type:      1,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 356,
					Type:      4,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 814,
					Type:      0,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 957,
					Type:      4,
				},
			},
		}
	return res, nil
}

type MempoolMock struct{}

func (m *MempoolMock) GetUnconfirmedBalance(ctx context.Context, req *EmptyRequest) (*BalanceResponse, error) {
	res :=
		&BalanceResponse{}
	return res, nil
}
func (m *MempoolMock) SelectTx(ctx context.Context, req *SelectRequest) (*SelectResponse, error) {
	res :=
		&SelectResponse{
			Result: []*Tx{
				&Tx{
					Type: 4,
					Id:   "6837ac09-13a3-4b03-9391-880a17620962",
				},
				&Tx{
					Type: 4,
					Id:   "51032f90-1d59-4198-94b4-bf12b90536aa",
				},
				&Tx{
					Type: 5,
					Id:   "b7eab07b-bc72-4b6f-bdfc-88b4399f88f4",
				},
				&Tx{
					Type: 2,
					Id:   "15ee51c1-f085-49d2-9055-474cb0b03ae1",
				},
				&Tx{
					Type: 1,
					Id:   "025db802-858a-4387-ae6f-45edfcb29b63",
				},
				&Tx{
					Type: 2,
					Id:   "4adc081a-da50-4915-9a42-c4e0359c4bee",
				},
				&Tx{
					Type: 0,
					Id:   "c83a297b-bd89-4b3c-b87a-7b23b1c39e3a",
				},
				&Tx{
					Type: 0,
					Id:   "8d043903-615b-4817-a95a-e5268177a22f",
				},
			},
		}
	return res, nil
}

type ChainMock struct{}

func (m *ChainMock) RebuildChain(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "tempore",
		}
	return res, nil
}
func (m *ChainMock) GetSyncProgress(ctx context.Context, req *EmptyRequest) (*SyncProgressResponse, error) {
	res :=
		&SyncProgressResponse{
			Progress: 81.6126,
		}
	return res, nil
}

type TransactorMock struct{}

func (m *TransactorMock) CallContract(ctx context.Context, req *CallContractRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Transfer(ctx context.Context, req *TransferRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Bid(ctx context.Context, req *BidRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}
func (m *TransactorMock) Stake(ctx context.Context, req *StakeRequest) (*TransactionResponse, error) {
	res :=
		&TransactionResponse{}
	return res, nil
}

type MaintainerMock struct{}

func (m *MaintainerMock) AutomateConsensusTxs(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "explicabo",
		}
	return res, nil
}
