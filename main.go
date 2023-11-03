// Using this tutorial to create an API connected with my database via GORM: https://blog.logrocket.com/rest-api-golang-gin-gorm/
// GORM Documentation: https://gorm.io/docs/models.html
// Using swagger and Gin: https://medium.com/@kumar16.pawan/integrating-swagger-with-gin-framework-in-go-f8d4883f4833
// Other articles to read:
// - https://dave.cheney.net/2014/09/14/go-list-your-swiss-army-knife
// - https://golang.org/cmd/go/#hdr-List_packages

/* OUTROS ITENS A ESTUDAR baseados na leitura do artigo https://about.gitlab.com/blog/2017/11/27/go-tools-and-gitlab-how-to-do-continuous-integration-like-a-boss/
 - Linter: https://staticcheck.dev/
 - Testes unitários: 
      - https://blog.jetbrains.com/go/2022/11/22/comprehensive-guide-to-testing-in-go/
      - https://medium.com/deliveryherotechhub/how-to-write-unit-test-in-go-1df2b98ad510
 - Data Race: https://golang.org/doc/articles/race_detector.html
 - Memory Sanitizer com Clang: https://clang.llvm.org/docs/MemorySanitizer.html
 - Code Coverage:
 - Makefile
 - 
*/

package main

import (
	"github.com/gin-gonic/gin"
	"opus127.com.br/models"
	"opus127.com.br/controllers"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "opus127.com.br/docs"
)

// @title RestFUL API usando Go, Gin, Swagger e GORM + SQLServer
// @version 1.0
// @description Essa é uma API de testes criada para testes usando a lingagem go

// @contact.name Opus127
// @contact.url http://www.opus127.com.br
// @contact.email suporte@opus127.com.br

// @license.name MIT
// @license.url https://opensource.org/licenses/MIT

// @host localhost:8080
// @BasePath /
// @query.collection.format multi
func main() {
	r := gin.Default()
	models.ConnectDatabase()
	
	r.GET("/customers", controllers.FindAllCustomers)
	r.GET("/customers/:id", controllers.FindCustomer)
	r.POST("/customers", controllers.CreateCustomer)
	r.PATCH("/customers/:id", controllers.UpdateCustomer)
	r.DELETE("/customers/:id", controllers.DeleteCustomer)

	r.GET("/clientes", controllers.FindAllClientes)
	r.GET("/clientes/:id", controllers.FindCliente)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run()
}
