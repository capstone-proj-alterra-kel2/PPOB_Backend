package users

import (
	"PPOB_BACKEND/app/aws"
	"PPOB_BACKEND/app/middlewares"
	"PPOB_BACKEND/businesses/users"
	"PPOB_BACKEND/controllers"
	"PPOB_BACKEND/controllers/users/request"
	"PPOB_BACKEND/controllers/users/response"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(userUC users.Usecase) *UserController {
	return &UserController{
		userUsecase: userUC,
	}
}

func (ctrl *UserController) GetAll(c echo.Context) error {
	usersData := ctrl.userUsecase.GetAll()

	users := []response.User{}

	for _, user := range usersData {
		if user.RoleName != "admin" && user.RoleName != "superadmin" {
			users = append(users, response.FromDomain(user))
		}
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all users", users)
}

func (ctrl *UserController) CreateUser(c echo.Context) error {
	var result string
	input := request.User{}
	image, _ := c.FormFile("image")
	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}
	user, err := ctrl.userUsecase.Register(input.ToDomain())

	isEmailDuplicate := strings.Contains(fmt.Sprint(err), "users_email_key")
	isNumberDuplicate := strings.Contains(fmt.Sprint(err), "users_phone_number_key")

	if isEmailDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if isNumberDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "Number already registered")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "user created", response.FromDomain(user))
}

func (ctrl *UserController) GetAllAdmin(c echo.Context) error {
	usersData := ctrl.userUsecase.GetAll()

	admins := []response.User{}

	for _, user := range usersData {
		if user.RoleName != "user" && user.RoleName != "superadmin" {
			admins = append(admins, response.FromDomain(user))
		}
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all Admin", admins)
}

func (ctrl *UserController) CreateAdmin(c echo.Context) error {
	var result string
	input := request.User{}
	image, _ := c.FormFile("image")
	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}
	user, err := ctrl.userUsecase.CreateAdmin(input.ToDomain())

	isEmailDuplicate := strings.Contains(fmt.Sprint(err), "users_email_key")
	isNumberDuplicate := strings.Contains(fmt.Sprint(err), "users_phone_number_key")
	if isEmailDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if isNumberDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "Number already registered")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "admin created", response.FromDomain(user))
}

func (ctrl *UserController) UpdateDataUser(c echo.Context) error {
	var result string
	idUser := c.Param("user_id")
	image, _ := c.FormFile("image")
	role := ctrl.userUsecase.Profile(idUser).RoleName
	input := request.UpdateData{}
	if role != "admin" && role != "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "cant update admin & superadmin")
	}
	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	updatedData, err := ctrl.userUsecase.UpdateData(idUser, input.ToDomain())
	if err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", err.Error())
	}

	if updatedData.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "data user not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "data user updated", response.FromDomain(updatedData))
}

func (ctrl *UserController) UpdateDataAdmin(c echo.Context) error {
	var result string
	idUser := c.Param("user_id")
	image, _ := c.FormFile("image")
	input := request.UpdateData{}
	role := ctrl.userUsecase.Profile(idUser).RoleName
	if role != "user" && role != "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "cant update user & superadmin")
	}
	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	updatedData, err := ctrl.userUsecase.UpdateData(idUser, input.ToDomain())
	if err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", err.Error())
	}

	if updatedData.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "data user not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "data user updated", response.FromDomain(updatedData))
}

func (ctrl *UserController) Register(c echo.Context) error {
	var result string
	input := request.User{}
	image, _ := c.FormFile("image")
	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}
	user, err := ctrl.userUsecase.Register(input.ToDomain())

	isEmailDuplicate := strings.Contains(fmt.Sprint(err), "users_email_key")
	isNumberDuplicate := strings.Contains(fmt.Sprint(err), "users_phone_number_key")

	if isEmailDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if isNumberDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "Number already registered")
	}

	return controllers.NewResponse(c, http.StatusCreated, "success", "user registered", response.FromDomain(user))
}

func (ctrl *UserController) Login(c echo.Context) error {
	userInput := request.UserLogin{}

	if err := c.Bind(&userInput); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	if err := userInput.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	token := ctrl.userUsecase.Login(userInput.ToDomain())
	if token == "" {
		return controllers.NewResponseFail(c, http.StatusUnauthorized, "failed", "invalid email or password")
	}
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}

func (ctrl *UserController) Logout(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)

	if isListed := middlewares.CheckToken(user.Raw); !isListed {
		return controllers.NewResponseFail(c, http.StatusUnauthorized, "failed", "invalid token")
	}
	middlewares.Logout(user.Raw)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "logout succes",
	})
}

func (ctrl *UserController) Profile(c echo.Context) error {
	idUser := middlewares.GetUserID(c)
	user := ctrl.userUsecase.Profile(idUser)

	if user.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "cannot load profile")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "profile loaded", response.FromDomain(user))
}

func (ctrl *UserController) UpdatePassword(c echo.Context) error {
	idUser := middlewares.GetUserID(c)
	input := request.UpdatePassword{}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	isPasswordUpdated := ctrl.userUsecase.UpdatePassword(idUser, input.ToDomain())
	if !isPasswordUpdated {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "wrong old password")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "password updated", "")
}

func (ctrl *UserController) UpdateData(c echo.Context) error {
	var result string
	idUser := middlewares.GetUserID(c)
	image, _ := c.FormFile("image")
	input := request.UpdateData{}

	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	updatedData, err := ctrl.userUsecase.UpdateData(idUser, input.ToDomain())
	if err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", err.Error())
	}

	if updatedData.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "data user not found")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "data user updated", response.FromDomain(updatedData))
}

func (ctrl *UserController) UpdateImage(c echo.Context) error {
	var result string
	idUser := middlewares.GetUserID(c)
	input := request.UpdateImage{}
	image, _ := c.FormFile("image")
	image.Filename = time.Now().String() + ".png"
	if image != nil {
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}
	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}
	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}
	user, err := ctrl.userUsecase.UpdateImage(idUser, input.ToDomain())
	if err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", err.Error())
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "image updated", response.FromDomain(user))
}
