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
	"strconv"
	"strings"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/morkid/paginate"
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
	pg := paginate.New()
	size, _ := strconv.Atoi(c.QueryParam("size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	search := c.QueryParam("search")
	sort := c.QueryParam("sort")

	usersData, userDomain := ctrl.userUsecase.GetAll(size, page, sort, search)
	users := []response.User{}
	for _, user := range userDomain {
		users = append(users, response.FromDomain(user))
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all users", pg.Response(usersData, c.Request(), &users))
}

func (ctrl *UserController) CreateUser(c echo.Context) error {
	var result string
	input := request.User{}
	image, _ := c.FormFile("image")

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	email, phone := ctrl.userUsecase.CheckDuplicateUser(input.Email, input.PhoneNumber)
	if email && phone {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email & password already registered")
	}
	if email {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if phone {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "phone already registered")
	}

	if image != nil {
		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
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
	pg := paginate.New()
	size, _ := strconv.Atoi(c.QueryParam("size"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	search := c.QueryParam("search")
	sort := c.QueryParam("sort")
	usersData, userDomain := ctrl.userUsecase.GetAllAdmin(size, page, sort, search)
	users := []response.User{}
	for _, user := range userDomain {
		users = append(users, response.FromDomain(user))
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "all Admin", pg.Response(usersData, c.Request(), &users))
}

func (ctrl *UserController) CreateAdmin(c echo.Context) error {
	var result string
	input := request.User{}
	image, _ := c.FormFile("image")

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	email, phone := ctrl.userUsecase.CheckDuplicateUser(input.Email, input.PhoneNumber)
	if email && phone {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email & password already registered")
	}
	if email {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if phone {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "phone already registered")
	}

	if image != nil {
		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
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

func (ctrl *UserController) DetailAdmin(c echo.Context) error {
	idUser := c.Param("user_id")
	user := ctrl.userUsecase.Profile(idUser)
	if user.RoleName == "user" || user.RoleName == "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "prevent getting detail user & superadmin")
	}
	if user.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "cannot get detail admin")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "get detail admin", response.FromDomain(user))
}

func (ctrl *UserController) DetailUser(c echo.Context) error {
	idUser := c.Param("user_id")
	user := ctrl.userUsecase.Profile(idUser)
	if user.RoleName == "admin" || user.RoleName == "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "prevent getting detail admin & superadmin")
	}
	if user.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "cannot get detail user")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "get detail user", response.FromDomain(user))
}

func (ctrl *UserController) DeleteUser(c echo.Context) error {
	idUser := c.Param("user_id")
	role := ctrl.userUsecase.Profile(idUser).RoleName
	if role == "admin" || role == "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "cant delete admin & superadmin")
	}
	if isSuccess := ctrl.userUsecase.DeleteUser(idUser); !isSuccess {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "cannot delete user not found")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "user deleted", "")
}

func (ctrl *UserController) DeleteAdmin(c echo.Context) error {
	idUser := c.Param("user_id")
	role := ctrl.userUsecase.Profile(idUser).RoleName
	if role == "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "cant delete superadmin")
	}
	if isSuccess := ctrl.userUsecase.DeleteUser(idUser); !isSuccess {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "cannot delete user not found")
	}
	return controllers.NewResponse(c, http.StatusOK, "success", "user deleted", "")
}

func (ctrl *UserController) UpdateDataUser(c echo.Context) error {
	var result string
	idUser := c.Param("user_id")
	image, _ := c.FormFile("image")
	role := ctrl.userUsecase.Profile(idUser).RoleName
	input := request.UpdateData{}

	if role == "admin" || role == "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "cant update admin & superadmin")
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	if image != nil {
		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}

	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	updatedData, err := ctrl.userUsecase.UpdateData(idUser, input.ToDomain())

	if updatedData.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "data user not found")
	}

	isEmailDuplicate := strings.Contains(fmt.Sprint(err), "users_email_key")
	isNumberDuplicate := strings.Contains(fmt.Sprint(err), "users_phone_number_key")

	if isEmailDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if isNumberDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "Number already registered")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "data user updated", response.FromDomain(updatedData))
}

func (ctrl *UserController) UpdateDataAdmin(c echo.Context) error {
	var result string
	idUser := c.Param("user_id")
	image, _ := c.FormFile("image")
	input := request.UpdateData{}
	role := ctrl.userUsecase.Profile(idUser).RoleName
	if role == "user" || role == "superadmin" {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "cant update superadmin")
	}

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	if image != nil {
		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
	}

	if err := input.Validate(); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "validation failed")
	}

	updatedData, err := ctrl.userUsecase.UpdateData(idUser, input.ToDomain())

	if updatedData.ID == 0 {
		return controllers.NewResponseFail(c, http.StatusNotFound, "failed", "data user not found")
	}

	isEmailDuplicate := strings.Contains(fmt.Sprint(err), "users_email_key")
	isNumberDuplicate := strings.Contains(fmt.Sprint(err), "users_phone_number_key")

	if isEmailDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if isNumberDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "Number already registered")
	}

	return controllers.NewResponse(c, http.StatusOK, "success", "data user updated", response.FromDomain(updatedData))
}

func (ctrl *UserController) Register(c echo.Context) error {
	var result string
	input := request.User{}
	image, _ := c.FormFile("image")

	if err := c.Bind(&input); err != nil {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "invalid request")
	}

	email, phone := ctrl.userUsecase.CheckDuplicateUser(input.Email, input.PhoneNumber)
	if email && phone {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email & password already registered")
	}
	if email {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if phone {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "phone already registered")
	}

	if image != nil {
		image.Filename = time.Now().String() + ".png"
		src, _ := image.Open()
		defer src.Close()
		result, _ = aws.UploadToS3(c, "profile/", image.Filename, src)
		input.Image = result
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

	if image != nil {
		image.Filename = time.Now().String() + ".png"
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
	isEmailDuplicate := strings.Contains(fmt.Sprint(err), "users_email_key")
	isNumberDuplicate := strings.Contains(fmt.Sprint(err), "users_phone_number_key")

	if isEmailDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "email already registered")
	}
	if isNumberDuplicate {
		return controllers.NewResponseFail(c, http.StatusBadRequest, "failed", "Number already registered")
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

	if image != nil {
		image.Filename = time.Now().String() + ".png"
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
