package app

import (
	"context"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/config"
	"github.com/NpoolPlatform/go-service-framework/pkg/logger"
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading" //nolint
	"github.com/NpoolPlatform/sphinx-service/pkg/crud"
	"golang.org/x/xerrors"
)

// MARK: optimization can be made, use channel (maybe) to improve performance under high qps payload

var (
	mapACK              map[string]*trading.ACKRequest
	ackListenIntervalMs int
	ackListenTimeoutMs  int
)

func init() {
	mapACK = make(map[string]*trading.ACKRequest)
	ackListenIntervalMs = config.GetIntValueWithNameSpace("", "ack_listen_interval_ms")
	ackListenTimeoutMs = config.GetIntValueWithNameSpace("", "ack_listen_timeout_ms")
	if ackListenIntervalMs <= 0 {
		ackListenIntervalMs = 50
	}
	if ackListenTimeoutMs <= 0 {
		ackListenTimeoutMs = 6000
	}
}

func ACK(ctx context.Context, in *trading.ACKRequest) (resp *trading.ACKResponse, err error) {
	resp = &trading.ACKResponse{
		IsOkay: false,
	}
	logger.Sugar().Infof("ack received: %v", in)
	logger.Sugar().Infof("ack type: %v", in.TransactionType)
	if in.TransactionType == signproxy.TransactionType_TransactionNew ||
		in.TransactionType == signproxy.TransactionType_PreSign ||
		in.TransactionType == signproxy.TransactionType_Signature ||
		in.TransactionType == signproxy.TransactionType_Broadcast {
		resp.IsOkay, err = crud.UpdateTransactionStatus(ctx, in)
	} else {
		mapACK[in.TransactionIdInsite] = in
		resp.IsOkay = true
	}
	return
}

func ListenTillSucceeded(transactionIDInsite string) (val *trading.ACKRequest, err error) {
	var ok bool
	ackListenTimeoutMsLoop := ackListenTimeoutMs
	for !ok && ackListenTimeoutMsLoop > 0 {
		time.Sleep(time.Duration(ackListenIntervalMs) * time.Millisecond)
		ackListenTimeoutMsLoop -= ackListenIntervalMs
		val, ok = mapACK[transactionIDInsite]
	}
	if val != nil {
		if !val.IsOkay {
			err = xerrors.New("tx rejected by proxy")
		}
		val = &trading.ACKRequest{
			TransactionType:     val.TransactionType,
			CoinTypeId:          val.CoinTypeId,
			TransactionIdInsite: val.TransactionIdInsite,
			TransactionIdChain:  val.TransactionIdChain,
			Address:             val.Address,
			Balance:             val.Balance,
			IsOkay:              val.IsOkay,
			ErrorMessage:        val.ErrorMessage,
		}
		mapACK[transactionIDInsite] = nil
	} else {
		err = xerrors.New("request timeout, please try again later")
	}
	return
}
