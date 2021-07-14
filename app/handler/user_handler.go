package handler

import (
	"log"
	"net/http"
	"wx/app/model"
	"wx/app/zerror"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
)

type UserHandler struct {
	UserService  model.UserService
	TokenService model.TokenService
}

// Me handler calls services for getting
func (h *UserHandler) Me(c *gin.Context) {

	var id int64 = 1

	u, err := h.UserService.Get(c, id)

	if err != nil {
		log.Printf("Unable to find user: %v\n%v", id, err)
		e := zerror.NewNotFound("user", cast.ToString(id))

		c.JSON(e.Status(), gin.H{
			"error": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": u,
	})

}

type signupReq struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,gte=6,lte=30"`
}

// Signup handler
func (h *UserHandler) Signup(c *gin.Context) {
	var req signupReq
	if err := c.Bind(&req); err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(zerror.Status(err), gin.H{
			"error": err,
		})
	}

	u := &model.User{
		Email:    req.Email,
		Password: req.Password,
	}
	err := h.UserService.Signup(c, u)
	if err != nil {
		log.Printf("Failed to sign up user: %v\n", err.Error())
		c.JSON(zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	// create token pair as strings
	tokens, err := h.TokenService.NewPairFromUser(c, u, "")

	if err != nil {
		log.Printf("Failed to create tokens for user: %v\n", err.Error())
		c.JSON(zerror.Status(err), gin.H{
			"error": err,
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"tokens": tokens,
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
