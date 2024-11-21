package handler

import (
	"elasticsearch/feature/user"
	"elasticsearch/feature/user/dtos"
	"encoding/json"
	"net/http"
)


type controller struct {
	service user.Usecase
}

func New(service user.Usecase) user.Handler {
	return &controller{
		service: service,
	}
}


func (ctl *controller) CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var user dtos.InputUser

    // Decode JSON body to struct
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    // Validate input (example: check if Name is empty)
    if user.Name == "" {
        http.Error(w, "Name is required", http.StatusBadRequest)
        return
    }

    // Call the service to create the user
    result, err := ctl.service.CreateUser(user)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    // Return success response
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "User created successfully",
        "user":    result,
    })
}

func (ctl *controller) GetUserDetails(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID parameter is required", http.StatusBadRequest)
        return
    }

    user, err := ctl.service.GetUserByID(id)
    if err != nil {
        http.Error(w, err.Error(), http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func (ctl *controller) SearchUsers(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    query := r.URL.Query().Get("q")

    users, err := ctl.service.SearchUsers(query)
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "users": users,
        "count": len(users),
    })
}

func (ctl *controller) DeleteUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodDelete {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID parameter is required", http.StatusBadRequest)
        return
    }

    if err := ctl.service.DeleteUser(id); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{
        "message": "User deleted successfully",
    })
}

func (ctl *controller) UpdateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPut {
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID parameter is required", http.StatusBadRequest)
        return
    }

    var user user.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    if err := ctl.service.UpdateUser(id, user); err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "User updated successfully",
        "user":    user,
    })
}