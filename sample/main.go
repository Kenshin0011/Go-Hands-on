package main

// Hello World表示する用
import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Hello, World!")
	result, err := doSomething("input")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Result:", result)
}

func doSomething(input string) (string, error) {
	if input == "" {
		return "", errors.New("input is empty")
	}
	return "result", nil
}

// API叩く用
// import (
// 	"encoding/json"
// 	"errors"
// 	"net/http"
// )

// func main() {
// 	http.HandleFunc("/api", apiHandler)
// 	http.ListenAndServe(":8080", nil)
// }

// func apiHandler(w http.ResponseWriter, r *http.Request) {
// 	result, err := doSomething("input")
// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	json.NewEncoder(w).Encode(result)
// }

// func doSomething(input string) (string, error) {
// 	if input == "" {
// 		return "", errors.New("input is empty")
// 	}
// var slice []int
// slice = append(slice, 1, 2, 3)
// slice := []int{1, 2, 3}
// 	return "result", nil
// }
