package main

import (
  "github.com/gin-gonic/gin"
  "github.com/paddyquinn/smartcar-api/dao"
  "github.com/paddyquinn/smartcar-api/handler"
)

func main() {
  router := gin.Default()
  hdlr := &handler.Handler{DAO: &dao.DAO{}}

  vehicles := router.Group("/vehicles/:id")
  {
    vehicles.GET("", hdlr.GetVehicle)
    vehicles.GET("/doors", hdlr.GetDoorSecurity)
    vehicles.GET("/fuel", hdlr.GetFuelRange)
    vehicles.GET("/battery", hdlr.GetBatteryRange)
    vehicles.POST("/engine", hdlr.PushEngineButton)
  }

  router.Run()
}