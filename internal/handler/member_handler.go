package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
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
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusCreated, response)
	}
}

func (m memberHandler) ReadMembers() gin.HandlerFunc {
	return func(c *gin.Context) {
		members, err := m.memberService.ReadMembers()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusFound, members)
	}
}

func (m memberHandler) ReadMemberById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("uuid")

		convertUUID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		member, err := m.memberService.ReadMemberById(convertUUID)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusFound, member)
	}
}

func (m memberHandler) DeleteMemberById() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("uuid")

		convertUUID, err := uuid.Parse(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}

		if err := m.memberService.DeleteMemberById(convertUUID); err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "the member has been deleted (soft) sort by uuid.",
		})
	}
}
