package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"github.com/yongyuth-chuankhuntod/book-management-with-hexagonal-architecture-in-go/internal/core/service"
	"net/http"
	"strconv"
	"strings"
)

type BookHandler struct {
	bookService service.BookService
}

func NewBookHandler(bookService service.BookService) BookHandler {
	return BookHandler{bookService: bookService}
}

func (b BookHandler) NewBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := []byte("MY_SIGNATURE_FOR_JWT")
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return signature, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["signature"])
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		var newBook service.NewBookRequester
		if err := c.ShouldBindJSON(&newBook); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err,
			})
		}

		response, err := b.bookService.NewBook(newBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusCreated, response)
	}
}

func (b BookHandler) ReadBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := []byte("MY_SIGNATURE_FOR_JWT")
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return signature, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["signature"])
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		responses, err := b.bookService.ReadBooks()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusFound, responses)
	}
}

func (b BookHandler) ReadBookById() gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := []byte("MY_SIGNATURE_FOR_JWT")
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return signature, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["signature"])
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		id := c.Param("id")
		conv, _ := strconv.ParseUint(id, 10, 32)

		response, err := b.bookService.ReadBookById(uint(conv))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusFound, response)
	}
}

func (b BookHandler) DeleteBookById() gin.HandlerFunc {
	return func(c *gin.Context) {
		signature := []byte("MY_SIGNATURE_FOR_JWT")
		tokenString := c.Request.Header.Get("Authorization")
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}

			return signature, nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": err,
			})
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			fmt.Println(claims["signature"])
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		id := c.Param("id")
		conv, _ := strconv.ParseUint(id, 10, 32)

		err = b.bookService.DeleteBookById(uint(conv))
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err,
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "the member has been deleted sort by id.",
		})
	}
}
