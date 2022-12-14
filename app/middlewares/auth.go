package middlewares

import (
	"net/http"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var whitelist []string = make([]string, 200)

type JWTCustomClaims struct {
	ID uint `json:"id"`
	jwt.StandardClaims
	HasRole string `json:"has_role"`
}

type ConfigJWT struct {
	SecretJWT      string
	ExpireDuration int
}

func (cj *ConfigJWT) Init() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &JWTCustomClaims{},
		SigningKey: []byte(cj.SecretJWT),
	}
}

// GenerateToken perform generating token and exp from userID
func (cj *ConfigJWT) GenerateToken(userID uint, role string) string {

	claims := JWTCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpireDuration))).Unix(),
		},
		role,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	listedToken, _ := token.SignedString([]byte(cj.SecretJWT))

	whitelist = append(whitelist, listedToken)
	return listedToken
}

// GetUser perform claims user
func GetUser(c echo.Context) *JWTCustomClaims {
	user := c.Get("user").(*jwt.Token)

	if isListed := CheckToken(user.Raw); !isListed {
		return nil
	}
	claims := user.Claims.(*JWTCustomClaims)
	return claims
}

// GetUserID perform get id user from JWT
func GetUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	userID := claims.ID
	idUser := strconv.FormatUint(uint64(userID), 10)
	return idUser
}

func CheckStatusToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userID := GetUser(c)

		if userID == nil {
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"message": "invalid token",
			})
		}
		return next(c)
	}
}

// IsSuperAdmin perform athorized only Superadmin can access
func IsSuperAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		HasRole := claims.HasRole

		if HasRole == "superadmin" {
			return next(c)
		}
		return echo.ErrForbidden
	}
}

// IsAdmin perform athorized only Superadmin & admin can access
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		HasRole := claims.HasRole

		if HasRole == "admin" || HasRole == "superadmin" {
			return next(c)
		}
		return echo.ErrForbidden
	}
}

// CheckToken perform checking the token in whitelist
func CheckToken(token string) bool {
	for _, listedToken := range whitelist {
		if listedToken == token {
			return true
		}
	}
	return false
}

// Logout perform deleting token in whitelist
func Logout(token string) bool {
	for i, listedToken := range whitelist {
		if listedToken == token {
			whitelist = append(whitelist[:i], whitelist[i+1:]...)
		}
	}
	return true
}
