package main

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	Db "github.com/zombieleet/codesphere-test-todo-app/internal/infrastructure/db"
	todoController "github.com/zombieleet/codesphere-test-todo-app/internal/todo/controllers"
)

func main() {

	Db.OpenConnection()
	mux := http.NewServeMux()

	catchAll := http.HandlerFunc(func(rw http.ResponseWriter, rq *http.Request) {

		if rq.Body == nil || rq.ContentLength == 0 {
			mux.ServeHTTP(rw, rq)
			return
		}

		var data interface{}

		err := json.NewDecoder(rq.Body).Decode(&data)

		if err != nil {
			rw.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(rw, err.Error())
			return
		}

		ctx := context.WithValue(rq.Context(), struct{}{}, data)
		rq = rq.WithContext(ctx)

		mux.ServeHTTP(rw, rq)

	})

	mux.Handle("/todo/", http.StripPrefix("/todo", todoController.InitTodoController()))

	server := &http.Server{
		Addr:    ":1337",
		Handler: catchAll,
	}

	fmt.Println("Listenings on port :1337")
	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}

}
