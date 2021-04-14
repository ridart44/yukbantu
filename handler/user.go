package handler

// The handler is used to handling the request; this handler is used to handling user related request
import (
	"net/http"
	"yukbantu/helper"
	"yukbantu/user"

	"github.com/gin-gonic/gin"
)

type userHandler struct {
	userService user.Service
}

func NewUserHandler(userService user.Service) *userHandler {
	return &userHandler{userService}
}

func (h *userHandler) RegisterUser(c *gin.Context) {
	// catch input from user
	var input user.RegisterUserInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Account Registration Failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	newUser, err := h.userService.RegisterUser(input)
	if err != nil {
		response := helper.APIResponse("Account Registration Failed!", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	formatter := user.FormatUser(newUser, "this is the token")

	response := helper.APIResponse("Account has been registered", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	// map input from user to struct RegisterUserInput
	// pass the struct as a service parameter
}

func (h *userHandler) Login(c *gin.Context) {
	// user input (email & password)
	// handler catch the input
	var input user.LoginInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}

		response := helper.APIResponse("Login Failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	loggedinUser, err := h.userService.Login(input)
	if err != nil {
		errorMessage := gin.H{"errors": err.Error()}
		response := helper.APIResponse("Login Failed!", http.StatusUnprocessableEntity, "error", errorMessage)
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	formatter := user.FormatUser(loggedinUser, "tokentokentoken")
	response := helper.APIResponse("Login Successful!", http.StatusOK, "success", formatter)
	c.JSON(http.StatusOK, response)
	// handler mapping user input to input struct
	// struct input passing the service
	// service search the relevant email using user repository
	// matching the password
}
