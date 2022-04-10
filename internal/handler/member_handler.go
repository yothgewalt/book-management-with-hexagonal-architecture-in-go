package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core/service"
	"net/http"
)

type memberHandler struct {
	memberService service.MemberService
}

func NewMemberHandler(memberService service.MemberService) memberHandler {
	return memberHandler{memberService: memberService}
}

func (m memberHandler) NewMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var member service.NewMemberRequester
		if err := c.ShouldBindJSON(&member); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		response, err := m.memberService.NewMember(member)
		if err != nil {
			c.JSON(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusCreated, response)
	}
}
