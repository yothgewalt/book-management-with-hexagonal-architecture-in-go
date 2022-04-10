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

	core.Database.AutoMigrate(&port.Member{})

	memberRepository := port.NewMemberRepository(core.Database)
	memberService := service.NewMemberService(memberRepository)
	memberHandler := handler.NewMemberHandler(memberService)

	r := gin.Default()

	api := r.Group("/api", func(c *gin.Context) { c.Next() })
	api.POST("/v1/create/member", memberHandler.NewMember())

	r.Run(":3000")
}
