package handler

import (
	"crowdfundex/helper"
	"crowdfundex/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	var input user.RegistrationUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.APIResponse("Register account unsuccessfully!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Register account unsuccessfully!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	formatter := user.FormatUser(newUser, "tokentokentokentokentokentoken")
	response := helper.APIResponse("Account hasbeen registered!", http.StatusOK, "success", formatter)

	// cara tracing
	// fmt.Println("%v", response)
	c.JSON(http.StatusOK, response)
	//tangkap input dari user
	//map input dari user ke struct RegistrationUserInput
	//struct di atas ke parsing sebagai parameter
}
