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
			Response: "qui",
		}
	return res, nil
}
func (m *WalletMock) GetWalletStatus(ctx context.Context, req *EmptyRequest) (*WalletStatusResponse, error) {
	res :=
		&WalletStatusResponse{
			Loaded: true,
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
					Direction: 0,
					Timestamp: 694,
					Type:      2,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 896,
					Type:      0,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 698,
					Type:      5,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 245,
					Type:      1,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 494,
					Type:      2,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 493,
					Type:      1,
				},
				&TxRecord{
					Direction: 0,
					Timestamp: 769,
					Type:      4,
				},
				&TxRecord{
					Direction: 1,
					Timestamp: 113,
					Type:      2,
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
					Type: 3,
					Id:   "ffdbd8f7-547f-4b62-8c94-a397d0b4ed54",
				},
				&Tx{
					Type: 2,
					Id:   "9e1b0cd5-96ca-4a73-9752-ed0f097dcc95",
				},
				&Tx{
					Type: 1,
					Id:   "0736b0ad-638c-45f2-80ba-fa39c7bd3605",
				},
				&Tx{
					Type: 0,
					Id:   "ec763cde-d118-4002-9166-0d8ca4fa7a50",
				},
				&Tx{
					Type: 0,
					Id:   "1cd6c388-eb4d-49a3-abfd-f144066b9fb3",
				},
				&Tx{
					Type: 0,
					Id:   "b105988f-c0cf-4483-bfcb-a001e86bad30",
				},
				&Tx{
					Type: 4,
					Id:   "d9a57dfa-ade7-498a-a788-776da0dc08c4",
				},
				&Tx{
					Type: 0,
					Id:   "08e9378f-1678-4569-8d93-ea00e727bad3",
				},
			},
		}
	return res, nil
}

type ChainMock struct{}

func (m *ChainMock) RebuildChain(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "omnis",
		}
	return res, nil
}
func (m *ChainMock) GetSyncProgress(ctx context.Context, req *EmptyRequest) (*SyncProgressResponse, error) {
	res :=
		&SyncProgressResponse{
			Progress: 800.4834,
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
			Response: "est",
		}
	return res, nil
}
