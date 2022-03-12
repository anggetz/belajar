package main

import (
	"belajar/golang/repository"
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/user", User)

	fmt.Println("service running on port", ":8085")
	err := http.ListenAndServe(":8085", nil)
	if err != nil {
		panic(err)
	}
}

func User(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-type", "application/json")

	users, err := repository.NewUserRepository().Get()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	byteArray, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(byteArray)
}
