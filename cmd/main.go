package main

import (
	"github.com/gin-gonic/gin"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core/database"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core/port"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core/service"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/handler"
	"gorm.io/gorm"
	"log"
)

func main() {
	var err error

	core.Database, err = database.Connect("configure.yaml", &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatalf("Failed to connect the database cause %v", err)
	}

	err = core.Database.AutoMigrate(&port.Member{})
	if err != nil {
		log.Fatalf("Failed to auto migrate the model cause %v", err)
	}

	memberRepository := port.NewMemberRepository(core.Database)
	memberService := service.NewMemberService(memberRepository)
	memberHandler := handler.NewMemberHandler(memberService)

	bookRepository := port.NewBookRepository(core.Database)

	_ = bookRepository

	r := gin.Default()

	api := r.Group("/api", func(c *gin.Context) { c.Next() })
	api.POST("/v1/auth/member/login", memberHandler.AuthMember())
	api.POST("/v1/create/member", memberHandler.NewMember())
	api.GET("/v1/members", memberHandler.ReadMembers())
	api.GET("/v1/members/:uuid", memberHandler.ReadMemberById())
	api.DELETE("/v1/delete/member/:uuid", memberHandler.DeleteMemberById())

	err = r.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
