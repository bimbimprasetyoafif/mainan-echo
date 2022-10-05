package main

import (
	"github.com/coba/config"
	"github.com/coba/databases"
	v1 "github.com/coba/routes/v1"
)

func main() {
	config.InitConfig()

	databases.InitDatabase()

	e, trace := v1.InitRoute()
	defer trace.Close()
	//
	e.Start(config.Cfg.APIPort)

	//http.HandleFunc("/", HandlerA)

	//http.Handle("/", http.HandlerFunc( HandlerA ) )

	//http.Handle("/A", MiddlewareA(http.HandlerFunc(HandlerA)))
	//http.Handle("/B", MiddlewareB(http.HandlerFunc(HandlerB)))
	//
	//http.ListenAndServe(":8080", nil)
}
