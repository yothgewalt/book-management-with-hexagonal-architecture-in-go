package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core/service"
	"net/http"
)

type MemberHandler struct {
	memberService service.MemberService
}

func NewMemberHandler(memberService service.MemberService) MemberHandler {
	return MemberHandler{memberService: memberService}
}

func (m MemberHandler) AuthMember() gin.HandlerFunc {
	return func(c *gin.Context) {
		var member service.AuthMemberRequester
		if err := c.ShouldBindJSON(&member); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
			return
		}

		response, err := m.memberService.AuthMember(member)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, response)
	}
}

func (m MemberHandler) NewMember() gin.HandlerFunc {
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

func (m MemberHandler) ReadMembers() gin.HandlerFunc {
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

func (m MemberHandler) ReadMemberById() gin.HandlerFunc {
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

func (m MemberHandler) DeleteMemberById() gin.HandlerFunc {
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
			"message": "the member has been deleted sort by uuid.",
		})
	}
}
