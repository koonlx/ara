package handler

import (
    "encoding/json"
    "ara/mulgyeol/model"
    "ara/mulgyeol/service"
    "net/http"
    "github.com/gorilla/mux"
)

type UserHandler struct {
    service service.UserService
}

func NewUserHandler(service service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

// 사용자 데이터를 받아 Kafka로 보냅니다.
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := h.service.SendUserData(&user); err != nil {
        http.Error(w, "Failed to send data", http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write([]byte("User data sent to Kafka"))
}

// 라우터 설정
func (h *UserHandler) SetupRouter() *mux.Router {
    router := mux.NewRouter()
    router.HandleFunc("/user", h.CreateUser).Methods("POST")
    return router
}
