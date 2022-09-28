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
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Register account unsuccessfully!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
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

func (h *userHandler) Login(c *gin.Context) {
	// user memasukan input email & password
	// input ditangkap handler
	// mapping dari input user ke input user struct
	// input struct parsing service
	// di service mencocokan email dengan bantuan repository user
	// mencocokan password

	var input user.LoginUserInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	loggedinUser, err := h.userService.LoginUser(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}

		response := helper.APIResponse("Login Failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, "tokentokentokentokentokentoken")
	response := helper.APIResponse("Log in sucessfully", http.StatusOK, "success", formatter)

	c.JSON(http.StatusOK, response)

}
