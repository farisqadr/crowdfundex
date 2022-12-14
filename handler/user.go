package handler

import (
	"crowdfundex/helper"
	"crowdfundex/user"
	"fmt"
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

func (h *userHandler) CheckEmailAvailability(c *gin.Context) {
	// ada input dari user
	// maping input jadi user input struct
	// parsing user email input struct ke service
	// service memanggil repository mengecheck ketersediaan email
	// repository check db
	var input user.CheckEmailInput

	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	isEmailAvailable, err := h.userService.IsEmailAvailability(input)

	if err != nil {
		errorMessage := gin.H{"errors": "Server error"}

		response := helper.APIResponse("Email checking failed", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	data := gin.H{
		"is_available": isEmailAvailable,
	}

	metaMessage := "Email has been registered"

	if isEmailAvailable {
		metaMessage = "Email is available"
	}

	response := helper.APIResponse(metaMessage, http.StatusOK, "success", data)
	c.JSON(http.StatusUnprocessableEntity, response)
	return

}

func (h *userHandler) UploadAvatar(c *gin.Context) {
	// input dari user
	// sinpan gambarnya di folder "images/"
	// di service kita panggil rep0
	// JWT (sementara hardcode, seakan2 user yg login ID = 1)
	// repo ambil data user yg ID 1
	// repo update data user simpan lokasi file

	file, err := c.FormFile("avatar")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed Upload Avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	// harusnya dapet dari jwt
	userID := 1
	// tambahan manipulasi text untuk menghindari duplikat
	path := fmt.Sprintf("images/%d-%s", userID, file.Filename)

	err = c.SaveUploadedFile(file, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed Upload Avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	_, err = h.userService.SaveAvatar(userID, path)
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.APIResponse("Failed Upload Avatar", http.StatusBadRequest, "error", data)

		c.JSON(http.StatusBadRequest, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.APIResponse("Success Upload Avatar", http.StatusOK, "success", data)

	c.JSON(http.StatusOK, response)
}
