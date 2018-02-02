package main

import (
	"github.com/paddyquinn/smartcar-api/dao"
	"github.com/paddyquinn/smartcar-api/handler"
)

func main() {
	hdlr := &handler.Handler{DAO: &dao.DAO{}}
	router := hdlr.SetUpRouter()
	router.Run()
}
