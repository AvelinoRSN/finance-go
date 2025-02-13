package transaction

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/AvelinoRSN/finance-go/model/transaction"
	util "github.com/AvelinoRSN/finance-go/uitl"
)

func GetTransactions(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))


	var transactions = transaction.Transtactions{
		transaction.Transaction{
			Title:     "Salario",
			Amount:    6000.6,
			Type:      0,
			CreatedAt: util.StringTotime("2019-02-12T09:00:00"),
		},
	}

	_ = json.NewEncoder(w).Encode(transactions)
}

func CreateATransaction(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Success"))

	var resp = transaction.Transtactions{}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	_ = json.Unmarshal(body, &resp)

	fmt.Print(resp)
}