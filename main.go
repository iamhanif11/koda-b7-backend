package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func main() {
	router := gin.Default()

	//login
	router.POST("/login", func(ctx *gin.Context) {
		var l Login
		if err := ctx.ShouldBindWith(&l, binding.JSON); err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "Error",
				Data:    nil,
				Success: false,
				Error:   "Internal Server Error",
			})
			return
		}
		var userLoggedIn User
		for _, u := range users {
			if u.Email == l.Email && u.Password == l.Password {
				userLoggedIn = u
				break
			}
			if u.Email != l.Email && u.Password != l.Password {
				ctx.JSON(http.StatusUnauthorized, Response{
					Message: "Invalid Email Or Password",
					Data:    nil,
					Success: false,
					Error:   "Unauthorized",
				})
				return
			}
		}
		ctx.JSON(http.StatusOK, Response{
			Message: "Login Succesfull",
			Data:    userLoggedIn,
			Success: true,
			Error:   "",
		})
	})

	router.POST("/register", func(ctx *gin.Context) {
		var l Login
		if err := ctx.ShouldBindWith(&l, binding.JSON); err != nil {
			ctx.JSON(http.StatusInternalServerError, Response{
				Message: "Error",
				Data:    nil,
				Success: false,
				Error:   "Internal Server Error",
			})
			return
		}

		newUser := User{
			Email:    l.Email,
			Password: l.Password,
		}
		users = append(users, newUser)

		log.Println(users)
		ctx.JSON(http.StatusOK, Response{
			Message: "Register Succesfully",
			Data:    newUser,
			Success: true,
			Error:   "",
		})
	})

	router.Run("localhost:8080")
}

var users []User

type User struct {
	Email    string
	Password string
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Response struct {
	Message string
	Data    any
	Success bool
	Error   string
}
