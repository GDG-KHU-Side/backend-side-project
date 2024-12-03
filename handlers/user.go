package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/GDG-KHU-Side/backend-side-project/models"
	pb "github.com/GDG-KHU-Side/backend-side-project/proto"
	"github.com/GDG-KHU-Side/backend-side-project/services"
	"github.com/gorilla/mux"
	"google.golang.org/grpc"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(s *services.UserService) *UserHandler {
	return &UserHandler{service: s}
}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "templates/login.html")
}

func (h *UserHandler) LoginUser(w http.ResponseWriter, r *http.Request) {

	log.Printf("Config: %v", r)

	var loginData models.LoginData

	if err := json.NewDecoder(r.Body).Decode(&loginData); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if loginData.Email == "" || loginData.Password == "" {
		http.Error(w, "Email and Password are required", http.StatusBadRequest)
		return
	}

	logData, err := json.MarshalIndent(loginData, "", "  ") // JSON 변환
	if err != nil {
		log.Fatalf("Failed to marshal loginData: %v", err)
	}
	log.Println("Decoded loginData:", string(logData))

	user, err := h.service.LoginUser(&loginData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	log.Println("Success")
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(user.Name)
	log.Println(user.PhoneNum)
	if user.Name == "" || user.Email == "" || user.Password == "" {
		http.Error(w, "필수 정보를 입력해주세요", http.StatusBadRequest)
		return
	}

	if err := h.service.CreateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	user, err := h.service.GetUser(id)
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if user.ID == 0 || user.Name == "" || user.Email == "" {
		http.Error(w, "User ID, Name, and Email are required", http.StatusBadRequest)
		return
	}

	if err := h.service.UpdateUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["id"], 10, 64)

	if err != nil || id <= 0 {
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	if err := h.service.DeleteUser(id); err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(map[string]string{"message": "User deleted successfully"}); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) LinkRestaurant(w http.ResponseWriter, r *http.Request) {
	var ur models.UserRest

	if err := json.NewDecoder(r.Body).Decode(&ur); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if err := h.service.LinkUserRest(&ur); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(ur); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *services.server) LoginUser(email, password string) (*pb.LoginUserResponse, error) {
	// gRPC 서버 연결
	conn, err := grpc.NewClient("localhost:8088", grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// UserService 클라이언트 생성
	client := pb.NewUserServiceClient(conn)

	// 컨텍스트 설정 (타임아웃 포함)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// Login 요청
	response, err := client.GLoginUser(ctx, &pb.LoginUserRequest{
		Email:    email,
		Password: password,
	})
	if err != nil {
		return nil, err
	}

	return response, nil
}
