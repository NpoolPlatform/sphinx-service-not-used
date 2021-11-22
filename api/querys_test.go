package api

import (
	"encoding/json"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/NpoolPlatform/go-service-framework/pkg/logger" //nolint
	"github.com/NpoolPlatform/message/npool/coininfo"          //nolint
	"github.com/NpoolPlatform/message/npool/signproxy"
	"github.com/NpoolPlatform/message/npool/trading"
	testinit "github.com/NpoolPlatform/sphinx-service/pkg/test-init"
	resty "github.com/go-resty/resty/v2"
	"github.com/stretchr/testify/assert"
)

var (
	tmpCoinInfo            coininfo.CoinInfo
	tmpAccountInfo         trading.CreateAccountResponse
	tmpTransactionIDInsite string
	testInitAlready        bool
	testHost               string
)

func init() {
	if runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION")); err == nil && runByGithubAction {
		return
	}
	tmpCoinInfo.Enum = 0
	tmpCoinInfo.ID = "6ba7b812-9dad-11d1-80b4-00c04fd430c8"
	tmpCoinInfo.PreSale = false
	tmpCoinInfo.Name = "Unknown"
	tmpCoinInfo.Unit = "DK"
	tmpAccountInfo.CoinName = "Unknown"
	tmpAccountInfo.Uuid = "6ba7b812-9dad-80b4-11d1-00c04fd430c8"
	tmpTransactionIDInsite = "test-tx-6ba7b812-80b4-9dad-11d1"
	testHost = "http://localhost:50160/v1"
}

func TestWholeProcedure(t *testing.T) {
	if runByGithub() {
		return
	}
	var err error
	// test create account
	go MockAccountCreated() // mock success
	err = TestCreateAccount()
	assert.Nil(t, err)
	assert.NotEmpty(t, tmpAccountInfo.Address)
	// test get balance
	go MockAccountBalance() // mock success
	resp, err := TestGetBalance(tmpAccountInfo.Address)
	assert.Nil(t, err)
	assert.NotNil(t, resp)
	assert.Zero(t, resp.AmountFloat64)
	// test create transaction
	// transaction would fail, but err should be nil
	go MockTransactionComplete() // mock success
	err = TestCreateTransaction(tmpAccountInfo.Address, tmpAccountInfo.Address)
	assert.Nil(t, err)
}

func runByGithub() bool {
	var err error
	runByGithubAction, err := strconv.ParseBool(os.Getenv("RUN_BY_GITHUB_ACTION"))
	if err == nil && runByGithubAction {
		return true
	}
	if testInitAlready == false {
		testInitAlready = true
		err = testinit.Init()
		if err != nil {
			logger.Sugar().Errorf("test init failed: %v", err)
		}
	}
	return err == nil
}

func TestCreateAccount() (err error) {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(trading.CreateAccountRequest{
			CoinName: tmpCoinInfo.Name,
			Uuid:     tmpAccountInfo.Uuid,
		}).
		Post(testHost + "/v1/account/register")
	if err != nil {
		return
	}
	expectedReturn := &trading.CreateAccountResponse{}
	err = json.Unmarshal(resp.Body(), expectedReturn)
	if err != nil {
		return
	}
	tmpAccountInfo.Address = expectedReturn.Address
	return
}

func TestCreateTransaction(addressFrom, addressTo string) (err error) {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(trading.CreateTransactionRequest{
			CoinName:            tmpCoinInfo.Name,
			TransactionIdInsite: tmpTransactionIDInsite,
			AddressFrom:         addressFrom,
			AddressTo:           addressTo,
			AmountFloat64:       123456.789,
			Type:                "payment",
			UuidSignature:       "",
			CreatetimeUtc:       time.Now().UTC().Unix(),
		}).
		Post(testHost + "/v1/transaction/create")
	if err != nil {
		return
	}
	expectedReturn := &trading.CreateTransactionResponse{}
	err = json.Unmarshal(resp.Body(), expectedReturn)
	return
}

func TestGetBalance(address string) (expectedReturn *trading.GetBalanceResponse, err error) {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(trading.GetBalanceRequest{
			CoinName:     "Unknown",
			Address:      address,
			TimestampUtc: time.Now().UTC().Unix(),
		}).
		Post(testHost + "/v1/account/balance/get")
	if err != nil {
		return
	}
	expectedReturn = &trading.GetBalanceResponse{}
	err = json.Unmarshal(resp.Body(), expectedReturn)
	return
}

func TestACK(req *trading.ACKRequest) (isOkay bool, err error) {
	cli := resty.New()
	resp, err := cli.R().
		SetHeader("Content-Type", "application/json").
		SetBody(req).
		Post(testHost + "/v1/internal/ack")
	if err != nil {
		return
	}
	expectedReturn := &trading.ACKResponse{}
	err = json.Unmarshal(resp.Body(), expectedReturn)
	if err == nil {
		isOkay = expectedReturn.IsOkay
	}
	return
}

func MockAccountCreated() (isOkay bool) {
	req := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_WalletNew,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: tmpAccountInfo.Uuid + tmpAccountInfo.CoinName,
		TransactionIdChain:  "",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	isOkay, err := TestACK(req)
	if err != nil {
		isOkay = false
		logger.Sugar().Errorf("create account ack error: ", err)
	}
	return
}

func MockAccountBalance() (isOkay bool) {
	req := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_Balance,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: "balance-" + tmpCoinInfo.Name + "-" + "testaddresshere",
		TransactionIdChain:  "",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	isOkay, err := TestACK(req)
	if err != nil {
		isOkay = false
		logger.Sugar().Errorf("account balance ack error: ", err)
	}
	return
}

func MockTransactionComplete() (isOkay bool) {
	req := &trading.ACKRequest{
		TransactionType:     signproxy.TransactionType_PreSign,
		CoinTypeId:          tmpCoinInfo.Enum,
		TransactionIdInsite: tmpTransactionIDInsite,
		TransactionIdChain:  "testchainidhere",
		Address:             "testaddresshere",
		Balance:             0.00,
		IsOkay:              true,
		ErrorMessage:        "",
	}
	isOkay, err := TestACK(req)
	if err != nil {
		isOkay = false
		logger.Sugar().Errorf("account balance ack error: ", err)
	}
	return
}