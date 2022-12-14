package main

import (
	"crowdfundex/handler"
	"crowdfundex/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfundex?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	// test update file name avatar
	// userService.SaveAvatar(1, "images/1-profile.png")

	// input := user.LoginUserInput{
	// 	Email:    "hawani@gmail.com",
	// 	Password: "plain123",
	// }
	// user, err := userService.LoginUser(i  nput)

	// userByEmail, err := userRepository.FindByEmail("thekaizar@gmail.com")

	// if err != nil {
	// 	fmt.Println("Terjadi kesalahan")
	// 	log.Fatal(err.Error())
	// }

	// fmt.Println(user.Name)
	// fmt.Println(user.Occupation)

	// if userByEmail.ID == 0 {
	// 	fmt.Println("User tidak ditemukan!")
	// } else {
	// 	fmt.Println(userByEmail.Name)
	// }

	// userHandler := handler.NewUserHandler(userService)
	router := gin.Default()

	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)

	router.Run()
	// userInput := user.RegistrationUserInput{}
	// userInput.Name = "Hawani Abdurrahman"
	// userInput.Email = "hawani@gmail.com"
	// userInput.Occupation = "CEO"
	// userInput.Password = "plain123"

	// userService.RegisterUser(userInput)

	// user := user.User{
	// 	Name: "Dinda Realita",
	// }

	// userRepository.Save(user)

	// input: dari user
	// handler: mapping input dari user -> struct input
	// service: melakukan mapping dari struct input ke struct user
	// repository: save butuh struct user (bukan struct input)
	// db

	// fmt.Println("Database connected")

	// var users []user.User

	// db.Find(&users)

	// for _, values := range users {
	// 	fmt.Println(values.Name)
	// 	fmt.Println(values.Role)
	// 	fmt.Println("=============")
	// }

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()
}

// func handler(c *gin.Context) {
// 	dsn := "root:@tcp(127.0.0.1:3306)/db_crowdfundex?charset=utf8mb4&parseTime=True&loc=Local"
// 	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	fmt.Println("Database connected")

// 	var users []user.User

// 	db.Find(&users)
// 	c.JSON(http.StatusOK, users)
// }
