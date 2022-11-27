package middlewares

import (
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
	RoleID uint `json:"role_id"`
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
func (cj *ConfigJWT) GenerateToken(userID uint, roleID uint) string {
	claims := JWTCustomClaims{
		userID,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(int64(cj.ExpireDuration))).Unix(),
		},
		roleID,
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

// Get user ID from JWT
func GetUserID(c echo.Context) string {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*JWTCustomClaims)
	userID := claims.ID
	idUser := strconv.FormatUint(uint64(userID ), 10)
	return idUser
}

// IsSuperAdmin perform athorized only Superadmin can access
func IsSuperAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		roleID := claims.RoleID

		if roleID != 3 {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}

// IsAdmin perform athorized only Superadmin & admin can access
func IsAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(*JWTCustomClaims)
		roleID := claims.RoleID

		if roleID == 3 || roleID == 2 {
			return next(c)
		}
		return echo.ErrUnauthorized
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

// logout perform deleting token in whitelist
func Logout(token string) bool {
	for i, listedToken := range whitelist {
		if listedToken == token {
			whitelist = append(whitelist[:i], whitelist[i+1:]...)
		}
	}
	return true
}
