package main

import (
	"log"

	"github.com/Akhildas-ts/pkg/auth"
	"github.com/Akhildas-ts/pkg/config"
	"github.com/Akhildas-ts/pkg/order"
	"github.com/Akhildas-ts/pkg/product"
	"github.com/gin-gonic/gin"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	r := gin.Default()

	authSvc := *auth.RegisterRoutes(r, &c)
	product.RegisterRoutes(r, &c, &authSvc)
	order.RegisterRoutes(r, &c, &authSvc)

	r.Run(c.Port)
}
