package app

import (
	"context"

	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"github.com/NpoolPlatform/sphinx-service/pkg/db/ent/transaction"
	"github.com/gogo/status"
	"google.golang.org/grpc/codes"
)

// 创建账号
func CreateWallet(ctx context.Context, coinName, uuid string) (account *trading.CreateWalletResponse, err error) {
	account = nil
	_, err = crud.CoinName2CoinID(coinName)
	if err != nil {
		return
	}
	address, err := LetCreateWallet(coinName, uuid)
	if err != nil {
		return
	}
	account = &trading.CreateWalletResponse{
		Info: &trading.EntAccount{
			CoinName: coinName,
			Address:  address,
		},
	}
	return
}

// 余额查询
func GetWalletBalance(ctx context.Context, in *trading.GetWalletBalanceRequest) (resp *trading.GetWalletBalanceResponse, err error) {
	balance, err := LetGetWalletBalance(in.Info.CoinName, in.Info.Address)
	if err != nil {
		return
	}
	resp = &trading.GetWalletBalanceResponse{
		Info: &trading.EntAccount{
			CoinName: in.Info.CoinName,
			Address:  in.Info.Address,
		},
		AmountFloat64: balance,
	}
	return
}

// 转账 / 提现
func CreateTransaction(ctx context.Context, in *trading.CreateTransactionRequest) (resp *trading.CreateTransactionResponse, err error) {
	// preset
	resp = &trading.CreateTransactionResponse{
		Info: in.Info,
	}
	_, err = crud.CheckRecordIfExistTransaction(in)
	if err != nil {
		return
	}
	// check uuid signature
	if in.UUIDSignature == "forbidden" {
		err = status.Error(codes.Canceled, "user signature invalid")
		return
	}
	// if amount > xxx, needManualReview => true, and etc.
	needManualReview := true
	// convert type
	txType := transaction.TypeUnknown
	if in.Info.InsiteTxType == "withdraw" {
		txType = transaction.TypeWithdraw
	} else if in.Info.InsiteTxType == "recharge" {
		txType = transaction.TypeRecharge
	} else if in.Info.InsiteTxType == "payment" {
		txType = transaction.TypePayment
	}
	// insert sql record
	info, err := crud.CreateRecordTransaction(in, needManualReview, txType)
	if err != nil {
		return
	}
	/*
		// rabbitmq notify
		// err = server.PublishNotificationTransactionCreate(&message.NotificationTransaction{
		// 	ID:                  info.ID,
		// 	TransactionIDInsite: info.TransactionIDInsite,
		// })
		// if err != nil {
		// 	err = status.Errorf(codes.Internal, "publishing notification error, check rabbitmq: %v", err)
		// } else {
		// 	resp.Info = "success"
		// }
	*/
	// grpc call approval service
	err = LetApproveTransaction(info)
	if err != nil {
		err = status.Errorf(codes.Internal, "cannot notify transaction approval service, error: %v", err)
		return
	}
	// done
	return resp, err
}

// 交易状态查询
func GetTransaction(ctx context.Context, in *trading.GetTransactionRequest) (resp *trading.GetTransactionResponse, err error) {
	// MARK 交给钱包代理
	return
}
