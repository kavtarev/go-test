package main

import (
	"encoding/json"
	"fmt"

	"net/http"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type hui struct {
	Name string
	Age  *string
}

func main() {
	var ttt hui
	age := "dfdsf"

	ttt.Name = "99999"
	ttt.Age = &age

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		res, err := json.Marshal(ttt)
		if err != nil {
			fmt.Println(1111111)

		}
		w.Write(res)
	})

	err := http.ListenAndServe(":3003", mux)
	if err != nil {
		panic(err)
	}

	fmt.Println("should not be there")
}
