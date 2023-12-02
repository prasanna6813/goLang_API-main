package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type User struct {
	Name string
	Age  int
}

type errorResp struct {
	StatusCode int
	Message    string
}

var (
	users = map[string]User{}
)

func main() {
	http.HandleFunc("/createUser", addUser)
	http.HandleFunc("/users", getUsers)
	fmt.Println("server started")
	log.Fatalf("Server not started, err: %v\n", http.ListenAndServe(":8000", nil))
}

// create/Add new record of user in map

func addUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}
	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		errr := errorResp{
			StatusCode: http.StatusBadRequest,
			Message:    "payload is not valid, error: " + err.Error(),
		}
		json.NewEncoder(w).Encode(errr)
		return
	}
	users[user.Name] = user
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(user)
	fmt.Println("users are: ", users)
	return
}

// return user record of user in map

func getUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		return
	}
	err := json.NewEncoder(w).Encode(&users)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		errr := errorResp{
			StatusCode: http.StatusInternalServerError,
			Message:    "payload is not created, error: " + err.Error(),
		}
		json.NewEncoder(w).Encode(errr)
		return
	}

	w.WriteHeader(http.StatusOK)

	return
}
