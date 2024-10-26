package handlers

import (
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
)

// RegisterStockRoutes: 주식 관련 경로를 등록합니다.
func RegisterStockRoutes(router *mux.Router) {
    stockRouter := router.PathPrefix("/stocks").Subrouter()
    stockRouter.HandleFunc("", GetStocks).Methods("GET")
    stockRouter.HandleFunc("", CreateStock).Methods("POST")
    stockRouter.HandleFunc("/{id}", DeleteStock).Methods("DELETE")
}

// GET 핸들러: 주식 목록 조회
func GetStocks(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Fetching all stocks")
}

// POST 핸들러: 새로운 주식 추가
func CreateStock(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintln(w, "Creating a new stock")
}

// DELETE 핸들러: 주식 삭제
func DeleteStock(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]
    fmt.Fprintf(w, "Deleting stock with ID: %s\n", id)
}
