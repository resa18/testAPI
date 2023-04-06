package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"
)

type Responses struct {
	Status string                 `json:"status"`
	Data   map[string]interface{} `json:"data"`
}
type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}

var customer_xid = "ea0212d3-abd6-406f-8c67-868e814a2436"
var token = "cb04f9f26632ad602f14acef21c58f58f6fe5fb55a"
var walletEnable = map[string]interface{}{
	"id":       "6ef31ed3-f396-4b6c-8049-674ddede1b16",
	"owned_by": "c4d7d61f-b702-44a8-af97-5dbdafa96551",
	"status":   "enabled",
	"balance":  0,
}

var walletBalance = map[string]interface{}{
	"id":       "c4d7d61f-b702-44a8-af97-5dbdafa96551",
	"owned_by": "6ef31975-67b0-421a-9493-667569d89556",
	"status":   "enabled",
	"balance":  0,
}

var deposit = map[string]interface{}{
	"id":           "ea0212d3-abd6-406f-8c67-868e814a2433",
	"deposited_by": "526ea8b2-428e-403b-b9fd-f10972e0d6fe",
	"status":       "success",
	"amount":       0,
	"reference_id": "f4cee01f-9188-4a29-aa9a-cb7fb97d8e0a",
}

var withdraw = map[string]interface{}{
	"id":           "ea0212d3-abd6-406f-8c67-868e814a2433",
	"withdrawn_by": "526ea8b2-428e-403b-b9fd-f10972e0d6fe",
	"status":       "success",
}

var disabledWallet = map[string]interface{}{
	"id":       "6ef31ed3-f396-4b6c-8049-674ddede1b16",
	"owned_by": "526ea8b2-428e-403b-b9fd-f10972e0d6fe",
	"status":   "disabled",
	"balance":  0,
}

func routes() {
	http.HandleFunc("/api/v1/init", getToken)
	http.HandleFunc("/api/v1/wallet", getWallet)
	http.HandleFunc("/api/v1/wallet/transactions", getTransactions)
	http.HandleFunc("/api/v1/wallet/deposits", getDeposit)
	http.HandleFunc("/api/v1/wallet/withdrawals", getWithdrawals)

}

func main() {

	routes()

	fmt.Println("Server running on port :8000")
	http.ListenAndServe(":8000", nil)
}
func errorMethod(status string, message string) (err ErrorResponse) {
	err = ErrorResponse{
		Status:  status,
		Message: message,
	}
	return err
}

func getToken(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if r.Method != "POST" {
		w.WriteHeader(http.StatusForbidden)
		response := errorMethod("error", "Method not allowed")
		json.NewEncoder(w).Encode(response)
		return
	}

	if r.FormValue("customer_xid") == customer_xid {
		datas := make(map[string]interface{})
		datas["token"] = token
		response := Responses{}
		response.Status = "success"
		response.Data = datas
		json.NewEncoder(w).Encode(response)
		return
	} else {
		w.WriteHeader(http.StatusExpectationFailed)
		response := errorMethod("fail", "customer_xid doesn't match")
		json.NewEncoder(w).Encode(response)
		return
	}

}

func getWallet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var body map[string]interface{}
	if r.Method == "POST" {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Token") {
			w.WriteHeader(http.StatusExpectationFailed)
			response := errorMethod("fail", "invalid token")
			json.NewEncoder(w).Encode(response)
			return
		}
		walletEnable["enabled_at"] = time.Now()
		response := Responses{}
		response.Status = "success"
		response.Data = map[string]interface{}{
			"wallet": walletEnable,
		}
		json.NewEncoder(w).Encode(response)
		return
	} else if r.Method == "GET" {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Token") {
			w.WriteHeader(http.StatusExpectationFailed)
			response := errorMethod("fail", "invalid token")
			json.NewEncoder(w).Encode(response)
			return
		}
		walletBalance["enabled_at"] = time.Now()
		response := Responses{}
		response.Status = "success"
		response.Data = map[string]interface{}{
			"wallet": walletBalance,
		}
		json.NewEncoder(w).Encode(response)
		return
	} else if r.Method == "PATCH" {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Token") {
			w.WriteHeader(http.StatusExpectationFailed)
			response := errorMethod("fail", "invalid token")
			json.NewEncoder(w).Encode(response)
			return
		}
		if r.FormValue("is_disabled") == "true" {
			disabledWallet["disabled_at"] = time.Now()
			response := Responses{}
			response.Status = "success"
			response.Data = map[string]interface{}{
				"wallet": disabledWallet,
			}
			json.NewEncoder(w).Encode(response)
			return
		}
	} else {
		w.WriteHeader(http.StatusForbidden)
		response := errorMethod("error", "Method not allowed")
		json.NewEncoder(w).Encode(response)
		return
	}

}

func getTransactions(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// var body map[string]interface{}
	if r.Method == "GET" {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Token") {
			w.WriteHeader(http.StatusExpectationFailed)
			response := errorMethod("fail", "invalid token")
			json.NewEncoder(w).Encode(response)
			return
		}
		response := Responses{}
		response.Status = "fail"
		response.Data = map[string]interface{}{
			"error": "Wallet disabled",
		}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
		response := errorMethod("error", "Method not allowed")
		json.NewEncoder(w).Encode(response)
		return
	}

}

func getDeposit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Token") {
			w.WriteHeader(http.StatusExpectationFailed)
			response := errorMethod("fail", "invalid token")
			json.NewEncoder(w).Encode(response)
			return
		}
		deposit["deposit_at"] = time.Now()
		response := Responses{}
		response.Status = "success"
		response.Data = map[string]interface{}{
			"deposit": deposit,
		}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
		response := errorMethod("error", "Only POST method accepted")
		json.NewEncoder(w).Encode(response)
		return
	}
}

func getWithdrawals(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method == "POST" {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Token") {
			w.WriteHeader(http.StatusExpectationFailed)
			response := errorMethod("fail", "invalid token")
			json.NewEncoder(w).Encode(response)
			return
		}
		withdraw["withdrawn_at"] = time.Now()
		withdraw["amount"] = r.FormValue("amount")
		withdraw["reference_id"] = r.FormValue("reference_id")
		response := Responses{}
		response.Status = "success"
		response.Data = map[string]interface{}{
			"withdrawal": withdraw,
		}
		json.NewEncoder(w).Encode(response)
		return
	} else {
		w.WriteHeader(http.StatusForbidden)
		response := errorMethod("error", "Only POST method accepted")
		json.NewEncoder(w).Encode(response)
		return
	}
}
