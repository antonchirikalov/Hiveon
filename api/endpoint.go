package api

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"hiveon-api/model"
	. "hiveon-api/service"
	. "hiveon-api/utils"
)

var (
	poolService   PoolService
	blockService  BlockService
	minerService  MinerService
	walletService WalletService
	userService   UserService
)

func init() {
	poolService = NewPoolService()
	blockService = NewBlockService()
	minerService = NewMinerService()
	walletService = NewWalletService()
	userService = NewUserService()
}

type WalletRequest struct {
	walletId string
}

type UserFIDRequest struct {
	fid int
}

type UserWalletPostRequest struct {
	Fid int
	Wallet string
	Coin string
}

type WalletWorkerRequest struct {
	walletId string
	workerId string
}

type WorkerRequest struct {
	workerId string
}

type UserRequest struct {
	Username string
	Password string
}

func poolGetIndexEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, _ interface{}) (response interface{}, err error) {
		return poolService.GetIndex(), nil
	}
}

func poolGetIncomeHistoryEndpoint() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (response interface{}, err error) {
		return poolService.GetIncomeHistory(), nil
	}
}

func blockGetBlockCountEndpoint() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (response interface{}, err error) {
		return blockService.GetBlockCount(), nil
	}
}

func minerGetFutureIncomeEndpoint() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (response interface{}, err error) {
		return minerService.GetFutureIncome(), nil
	}
}

func minerGetBillInfoEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletRequest)
		return minerService.GetBillInfo(FormatWalletID(req.walletId)), nil
	}
}

func minerGetBillEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletRequest)
		return minerService.GetBill(FormatWalletID(req.walletId)), nil
	}
}

func minerGetSharesEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletRequest)
		return minerService.GetShares(FormatWalletID(req.walletId), ""), nil
	}
}

func minerGetMinerEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletRequest)
		return minerService.GetMiner(FormatWalletID(req.walletId), ""), nil
	}
}

func minerGetHashrateEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletRequest)
		return minerService.GetHashrate(FormatWalletID(req.walletId), ""), nil
	}
}

func minerGetCountHistoryEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletRequest)
		return minerService.GetCountHistory(FormatWalletID(req.walletId)), nil
	}
}

//stub
func userGetUserInfoEndpoint() endpoint.Endpoint {
	return func(_ context.Context, _ interface{}) (response interface{}, err error) {
		userInfo := model.UserInfo{Code: 200}
		return userInfo, nil
	}
}

//auth
func userRegisterEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)
		return Register(req.Username, req.Password)
	}
}

func userLoginEndpoint() endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserRequest)
		return Login(req.Username, req.Password)
	}
}

func getWalletEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletRequest)
		return walletService.GetWalletInfo(FormatWalletID(req.walletId)), nil
	}
}

func getWalletWorkerEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletWorkerRequest)
		return walletService.GetWalletWorkerInfo(FormatWalletID(req.walletId), req.workerId), nil
	}
}

func getUserWalletEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserFIDRequest)
		return userService.GetUserWallet(req.fid), nil
	}
}

func postUserWalletEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(UserWalletPostRequest)
		userService.SaveUserWallet(req.Fid, req.Wallet, req.Coin)
		userInfo := model.UserInfo{Code: 201}
		return userInfo, nil
	}
}

func getWorkersStatisticEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WorkerRequest)
		res := minerService.CalcWorkersStat("", req.workerId)
		return res,nil
	}
}

func getWalletsStatisticEndpoint() endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (response interface{}, err error) {
		req := request.(WalletRequest)
		res := minerService.CalcWorkersStat(req.walletId, "")
		return res,nil
	}
}
