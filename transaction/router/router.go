package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"transaction/controllers/alive"
	"transaction/controllers/transaction"
)

// Package classification Documentation of transaction Service
//
// Documentation for transaction Service
//
// Schemes: https
// BasePath: /
// Version: 1.0.0
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// securityDefinitions:
//   BasicAuth:
//     type: basic
//     name: Authorization
//     in: header
//
// swagger:meta
func router(ginEngine *gin.Engine) {
	ginEngine.Use(cors.Default())

	ginEngine.GET("/alive", alive.Alive)
	ginEngine.PUT("/transaction", transaction.Update)
}
