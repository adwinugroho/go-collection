package main

import (
	"github.com/adwinugroho/go-aws-s3/apis"
	"github.com/adwinugroho/go-aws-s3/services"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	serviceInstance := services.NewService()
	apis.Init(e, serviceInstance)
	e.Logger.Fatal(e.Start(":8001"))
}
