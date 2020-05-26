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

type AuthMock struct{}

func (m *AuthMock) CreateSession(ctx context.Context, req *SessionRequest) (*Session, error) {
	res :=
		&Session{
			AccessToken: "veritatis",
		}
	return res, nil
}

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
			Response: "repellat",
		}
	return res, nil
}
func (m *WalletMock) DropSession(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "rem",
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
					Direction:  0,
					Timestamp:  173,
					Type:       1,
					Obfuscated: false,
				},
				&TxRecord{
					Direction:  1,
					Timestamp:  909,
					Type:       0,
					Obfuscated: true,
				},
				&TxRecord{
					Direction:  1,
					Timestamp:  487,
					Type:       3,
					Obfuscated: true,
				},
				&TxRecord{
					Direction:  1,
					Timestamp:  852,
					Type:       4,
					Obfuscated: false,
				},
				&TxRecord{
					Direction:  0,
					Timestamp:  877,
					Type:       0,
					Obfuscated: false,
				},
				&TxRecord{
					Direction:  1,
					Timestamp:  573,
					Type:       5,
					Obfuscated: true,
				},
				&TxRecord{
					Direction:  1,
					Timestamp:  454,
					Type:       2,
					Obfuscated: false,
				},
				&TxRecord{
					Direction:  0,
					Timestamp:  557,
					Type:       0,
					Obfuscated: true,
				},
			},
		}
	return res, nil
}

type MempoolMock struct{}

func (m *MempoolMock) GetUnconfirmedBalance(ctx context.Context, req *GetUnconfirmedBalanceRequest) (*BalanceResponse, error) {
	res :=
		&BalanceResponse{}
	return res, nil
}
func (m *MempoolMock) SelectTx(ctx context.Context, req *SelectRequest) (*SelectResponse, error) {
	res :=
		&SelectResponse{
			Result: []*Tx{
				&Tx{
					Type: 7,
					Id:   "e68507a8-0205-4954-9283-66822da6c818",
				},
				&Tx{
					Type: 5,
					Id:   "25956993-ccb6-4852-ad7e-8e6bec467722",
				},
				&Tx{
					Type: 3,
					Id:   "9c422570-44c0-49c0-a9e9-824ac8364f3e",
				},
				&Tx{
					Type: 4,
					Id:   "11243a3f-3831-4913-b873-8884f0664a80",
				},
				&Tx{
					Type: 7,
					Id:   "452eae28-1560-4552-b449-f51aa8d4189d",
				},
				&Tx{
					Type: 4,
					Id:   "489259fa-422f-4e70-8ddf-f6deed857d20",
				},
				&Tx{
					Type: 2,
					Id:   "c2c71ad4-eafa-4812-bded-443cf10d773e",
				},
				&Tx{
					Type: 3,
					Id:   "10c3a5f0-fc07-4fac-8242-4ee871249df2",
				},
			},
		}
	return res, nil
}

type ChainMock struct{}

func (m *ChainMock) RebuildChain(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "unde",
		}
	return res, nil
}
func (m *ChainMock) GetSyncProgress(ctx context.Context, req *EmptyRequest) (*SyncProgressResponse, error) {
	res :=
		&SyncProgressResponse{
			Progress: 394.5703,
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

type ProvisionerMock struct{}

func (m *ProvisionerMock) AutomateStakes(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "et",
		}
	return res, nil
}

type BlockGeneratorMock struct{}

func (m *BlockGeneratorMock) AutomateBids(ctx context.Context, req *EmptyRequest) (*GenericResponse, error) {
	res :=
		&GenericResponse{
			Response: "reiciendis",
		}
	return res, nil
}
