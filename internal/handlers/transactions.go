package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
	"vericred/internal/db"
	"vericred/internal/models"
)


type EtherscanResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		BlockHash        string `json:"blockHash"`
		BlockNumber      string `json:"blockNumber"`
		From             string `json:"from"`
		Gas              string `json:"gas"`
		GasPrice         string `json:"gasPrice"`
		Hash             string `json:"hash"`
		Input            string `json:"input"`
		Nonce            string `json:"nonce"`
		To               string `json:"to"`
		TransactionIndex string `json:"transactionIndex"`
		Value            string `json:"value"`
		V                string `json:"v"`
		R                string `json:"r"`
		S                string `json:"s"`
	} `json:"result"`
}

type BlockResponse struct {
	Jsonrpc string `json:"jsonrpc"`
	ID      int    `json:"id"`
	Result  struct {
		Timestamp string `json:"timestamp"`
	} `json:"result"`
}

func SetTransactionInfo(w http.ResponseWriter, r *http.Request) {
	var body map[string]any

	json.NewDecoder(r.Body).Decode(&body)

	apiKey := "9SV725QFIZ8NPSAKJQ43IGTBHKG8P6CE55"          // too lazy to set in env variable 
	txHash := body["transaction_hash"].(string)

	url := fmt.Sprintf("https://api-sepolia.etherscan.io/api?module=proxy&action=eth_getTransactionByHash&txhash=%s&apikey=%s", txHash, apiKey)

	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error making request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	bdy, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	var response EtherscanResponse
	err = json.Unmarshal(bdy, &response)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		os.Exit(1)
	}
	
	TimeStamp, err := getBlockTimestamp(apiKey, response.Result.BlockNumber)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	var trnx = models.Transaction{
		TxHash: response.Result.Hash,
		BlockNumber: response.Result.BlockNumber,
		From: response.Result.From,
		To: response.Result.To,
		ValueETH: response.Result.Value,
		Gas: response.Result.Gas,
		GasPrice: response.Result.GasPrice,
		Timestamp: TimeStamp,
	}

	createResult := db.DB.Create(&trnx)
	
	if createResult.Error != nil {
		http.Error(w, "Some error when creating transaction db.", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Transaction successfully recorded."))

}


func getBlockTimestamp(apiKey, blockNumber string) (time.Time, error) {
	url := fmt.Sprintf("https://api-sepolia.etherscan.io/api?module=proxy&action=eth_getBlockByNumber&tag=%s&boolean=true&apikey=%s", blockNumber, apiKey)
	resp, err := http.Get(url)
	if err != nil {
		return time.Time{}, err
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)

	var blockResp BlockResponse
	if err := json.Unmarshal(body, &blockResp); err != nil {
		return time.Time{}, err
	}

	timestampHex := blockResp.Result.Timestamp
	timestampInt, err := strconv.ParseInt(timestampHex[2:], 16, 64)
	if err != nil {
		return time.Time{}, err
	}

	return time.Unix(timestampInt, 0), nil
}

func ShowAllTransactions(w http.ResponseWriter, r *http.Request) {
	var trnxs []models.Transaction
	result := db.DB.Find(&trnxs)
	if result.Error != nil {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "trnx not found"})
		return
	}
	json.NewEncoder(w).Encode(trnxs)
}