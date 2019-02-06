package api

import (
	"context"
	"encoding/json"
	"errors"
	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strconv"
)

var errBadRoute = errors.New("bad route")
var errBadRequestBody = errors.New("error reading request body")


func MakeServiceHandlers() http.Handler {

	opts := []kithttp.ServerOption{
		kithttp.ServerErrorEncoder(encodeError),
	}
	r := mux.NewRouter()

	// swagger:route GET /pool/index pool emptyReq
	// responses:
	// 200: PoolData

	poolGetIndexHandler := kithttp.NewServer(
		poolGetIndexEndpoint(),
		decodeEmptyRequest,
		encodeResponse, opts...,
	)
	r.Handle("/api/pool/index", poolGetIndexHandler).Methods("GET")

	// swagger:route GET /pool/incomeHistory pool emptyReq
	// responses:
	// 200: IncomeHistory

	poolGetIncomeHistoryHandler := kithttp.NewServer(
		poolGetIncomeHistoryEndpoint(),
		decodeEmptyRequest,
		encodeResponse, opts...,
	)
	r.Handle("/api/pool/incomeHistory", poolGetIncomeHistoryHandler).Methods("GET")

	// swagger:route GET /block/count24h block emptyReq
	// responses:
	// 200: BlockCount
	blockGetBlockCountHandler := kithttp.NewServer(
		blockGetBlockCountEndpoint(),
		decodeEmptyRequest,
		encodeResponse,
	)
	r.Handle("/api/block/count24h", blockGetBlockCountHandler).Methods("GET")

	// swagger:route GET /miner/featureIncome miner emptyReq
	// responses:
	// 200: FutureIncome
	minerGetFutureIncomeHandler := kithttp.NewServer(
		minerGetFutureIncomeEndpoint(),
		decodeEmptyRequest,
		encodeResponse,
	)
	r.Handle("/api/miner/featureIncome", minerGetFutureIncomeHandler).Methods("GET")

	// swagger:operation GET /miner/{walletId}/billInfo miner emptyReq
	//
	// Returns wallet Bill Info
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/BillInfo"
	minerGetBillInfoHandler := kithttp.NewServer(
		minerGetBillInfoEndpoint(),
		decodeWalletIdRequest,
		encodeResponse,
	)
	r.Handle("/api/miner/{walletId}/billInfo", minerGetBillInfoHandler).Methods("GET")

	// swagger:operation GET /miner/{walletId}/bill miner emptyReq
	//
	// Returns wallet bill
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/Bill"
	minerGetBillHandler := kithttp.NewServer(
		minerGetBillEndpoint(),
		decodeWalletIdRequest,
		encodeResponse,
	)
	r.Handle("/api/miner/{walletId}/bill", minerGetBillHandler).Methods("GET")

	// swagger:operation GET /miner/{walletId}/shares miner emptyReq
	//
	// Returns wallet shares
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/Shares"
	minerGetSharesHandler := kithttp.NewServer(
		minerGetSharesEndpoint(),
		decodeWalletIdRequest,
		encodeResponse,
	)
	r.Handle("/api/miner/{walletId}/shares", minerGetSharesHandler).Methods("GET")

	// swagger:operation GET /miner/{walletId}/hashrate miner emptyReq
	//
	// Returns wallet hashrate
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/Hashrate"
	minerGetHashrateHandler := kithttp.NewServer(
		minerGetHashrateEndpoint(),
		decodeWalletIdRequest,
		encodeResponse,
	)
	r.Handle("/api/miner/{walletId}/hashrate", minerGetHashrateHandler).Methods("GET")

	// swagger:operation GET /miner/{walletId}/workers/counts miner emptyReq
	//
	// Returns wallet counts
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/WorkerCount"
	minerGetCountHistoryHandler := kithttp.NewServer(
		minerGetCountHistoryEndpoint(),
		decodeWalletIdRequest,
		encodeResponse,
	)
	r.Handle("/api/miner/{walletId}/workers/counts", minerGetCountHistoryHandler).Methods("GET")

	// swagger:operation GET /miner/ETH/{walletId} miner emptyReq
	//
	// Returns wallet data
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/WalletInfo"
	getCoinWalletData := kithttp.NewServer(
		getWalletEndpoint(),
		decodeWalletIdRequest,
		encodeResponse,
	)
	r.Handle("/api/miner/ETH/{walletId}", getCoinWalletData).Methods("GET")

	// swagger:operation GET /miner/ETH/{walletId}/{workerId} miner emptyReq
	//
	// Returns wallet and worker data
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// - name: workerId
	//   in: path
	//   description: Wallets's worker
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/..."
	getCoinWalletWorkerData := kithttp.NewServer(
		getWalletWorkerEndpoint(),
		decodeWalletIdWorkerIdRequest,
		encodeResponse,
	)
	r.Handle("/api/miner/ETH/{walletId}/{workerId}", getCoinWalletWorkerData).Methods("GET")

	// swagger:operation GET /page/miner?value={walletId} page emptyReq
	//
	// Returns wallet counts
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/MinerWorker"
	minerGetMinerHandler := kithttp.NewServer(
		minerGetMinerEndpoint(),
		decodeWalletIdFromParam,
		encodeResponse,
	)
	r.Handle("/api/page/miner", minerGetMinerHandler).Methods("GET")

	// swagger:route GET /user/userInfo user emptyReq
	// responses:
	// 200: UserInfo
	userGetUserInfoHandler := kithttp.NewServer(
		userGetUserInfoEndpoint(),
		decodeEmptyRequest,
		encodeResponse,
	)
	r.Handle("/api/user/userInfo", userGetUserInfoHandler).Methods("GET")


	//TODO don't need this anymore
	userRegisterHandler := kithttp.NewServer(
		userRegisterEndpoint(),
		decodeUserRequest,
		encodeResponse,
	)
	r.Handle("/api/register/{username}/{password}", userRegisterHandler).Methods("GET")


	//TODO don't need this anymore
	userLoginHandler := kithttp.NewServer(
		userLoginEndpoint(),
		decodeUserRequest,
		encodeResponse,
	)
	r.Handle("/api/login/{username}/{password}", userLoginHandler).Methods("GET")


	// swagger:operation GET /api/private/{fid} users emptyReq
	//
	// Returns wallet counts
	// ---
	// parameters:
	// - name: fid
	//   in: path
	//   description: User's FID
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/UserWallet"
	userWalletGetHandler := kithttp.NewServer(
		getUserWalletEndpoint(),
		decodeUserFidRequest,
		encodeResponse,
	)
	r.Handle("/api/private/{fid}", userWalletGetHandler).Methods("GET")

	// swagger:operation POST /api/private/wallets
	//
	// Add new wallet
	// ---
	// parameters:
	// - name: fid
	//   in: body
	//   description: User's FID
	//   required: true
	//   type: integer
	// - name: wallet
	//   in: body
	//   description: User's wallet
	//   required: true
	//   type: string
	// - name: coin
	//   in: body
	//   description: Wallet's coin
	//   required: true
	//   type: string
	// responses:
	//   "201":
	//     description: Created

	userWalletPostHandler := kithttp.NewServer(
		postUserWalletEndpoint(),
		decodePostUserWalletRequest,
		encodeResponse,
	)
	r.Handle("/api/private/wallets", userWalletPostHandler).Methods("POST")

	// swagger:operation GET /miner/statistic/worker/{workerId}
	//
	// Returns worker's statistic
	// ---
	// parameters:
	// - name: workerId
	//   in: path
	//   description: User's worker
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/..."

	workersStatisticHandler := kithttp.NewServer(
		getWorkersStatisticEndpoint(),
		decodeWorkerIdRequest,
		encodeResponse,
	)
	r.Handle("/api/private/statistic/worker/{workerId}", workersStatisticHandler).Methods("GET")

	// swagger:operation GET /miner/statistic/wallet/{walletId}
	//
	// Returns wallet's statistic
	// ---
	// parameters:
	// - name: walletId
	//   in: path
	//   description: User's wallet
	//   required: true
	//   type: string
	// responses:
	//   "200":
	//     "$ref": "#/responses/..."

	walletsStatisticHandler := kithttp.NewServer(
		getWalletsStatisticEndpoint(),
		decodeWalletIdRequest,
		encodeResponse,
	)
	r.Handle("/api/private/statistic/wallet/{walletId}", walletsStatisticHandler).Methods("GET")

	return r
}

func decodeEmptyRequest(_ context.Context, r *http.Request) (interface{}, error) {
	return r, nil
}


func decodeUserFidRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["fid"])
	if err != nil {
		return err, errBadRoute
	}
	return UserFIDRequest{fid: id}, nil
}


func decodePostUserWalletRequest(_ context.Context, r *http.Request) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var t UserWalletPostRequest
	err := decoder.Decode(&t)
	if err != nil {
		return err, errBadRequestBody
	}
	return t, nil
}

func decodeWalletIdFromParam(_ context.Context, r *http.Request) (interface{}, error) {
	walletId := r.URL.Query().Get("value")
	return WalletRequest{walletId: walletId}, nil
}



func decodeUserRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	username, ok := vars["username"]
	if !ok {
		return nil, errBadRoute
	}
	password, ok := vars["password"]
	if !ok {
		return nil, errBadRoute
	}
	return UserRequest{username, password}, nil
}

func decodeWalletIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	id, ok := vars["walletId"]
	if !ok {
		return nil, errBadRoute
	}
	return WalletRequest{walletId: id}, nil
}

func decodeWalletIdWorkerIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	walletId, ok := vars["walletId"]
	if !ok {
		return nil, errBadRoute
	}
	workerId, ok := vars["workerId"]
	if !ok {
		return nil, errBadRoute
	}
	return WalletWorkerRequest{walletId: walletId, workerId:workerId}, nil
}

func decodeWorkerIdRequest(_ context.Context, r *http.Request) (interface{}, error) {
	vars := mux.Vars(r)
	workerId, ok := vars["workerId"]
	if !ok {
		workerId = "" // all workers
	}
	return WorkerRequest{workerId:workerId}, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
	log.Info(response)
	return json.NewEncoder(w).Encode(response)
}

// encode errors from business-logic
func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	switch err {
	default:
		w.WriteHeader(http.StatusInternalServerError)
	}
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

func AccessControl(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type")
		if r.Method == "OPTIONS" {
			return
		}
		h.ServeHTTP(w, r)
	})
}
