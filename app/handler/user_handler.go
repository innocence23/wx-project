package handler

import (
	"log"
	"net/http"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserHandler struct {
	UserService model.UserService
}

// Me handler calls services for getting
func (h *UserHandler) Me(c *gin.Context) {

	uid, _ := uuid.NewRandom()

	u, err := h.UserService.Get(c, uid)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", uid, err)
		e := zerror.NewNotFound("user", uid.String())

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})

}

// Signup handler
func (h *UserHandler) Signup(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signup",
	})
}

// Signin handler
func (h *UserHandler) Signin(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signin",
	})
}

// Signout handler
func (h *UserHandler) Signout(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's signout",
	})
}

// Tokens handler
func (h *UserHandler) Tokens(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's tokens",
	})
}

// Image handler
func (h *UserHandler) Image(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's image",
	})
}

// DeleteImage handler
func (h *UserHandler) DeleteImage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's deleteImage",
	})
}

// Details handler
func (h *UserHandler) Details(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"hello": "it's details",
	})
}
