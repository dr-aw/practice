package httpHandler

import (
	"github.com/dr-aw/practice/internal/app/database"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorm.io/gorm"
	"log"
	"net/http"
	"time"
)

var jwtKey = []byte("your_secret_key")
var db *gorm.DB

// StartServer запускает HTTP сервер
func StartServer(gormdb *gorm.DB) {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	db = gormdb

	// renderer
	renderer, err := initializeTemplates()
	if err != nil {
		log.Fatal("error initialize renderer: ", err)
	}
	e.Renderer = renderer

	// endpoints
	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "base.html", map[string]interface{}{
			"Title": "Home",
		})
	})

	e.GET("/register", func(c echo.Context) error {
		return c.Render(http.StatusOK, "register.html", map[string]interface{}{
			"Title": "Register",
		})
	})
	e.POST("/register", register)

	e.GET("/login", func(c echo.Context) error {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"Title": "Login",
		})
	})
	e.POST("/login", login)
	//e.GET("/list", listProducts)

	e.Logger.Fatal(e.Start(":8080"))
}

// register обрабатывает регистрацию пользователей
func register(c echo.Context) error {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.Bind(&user); err != nil {
		return c.String(http.StatusBadRequest, "Invalid input")
	}

	err := database.AddUser(db, user.Username, user.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "Could not register user")
	}

	return c.Render(http.StatusOK, "register.html", map[string]interface{}{
		"Title":   "Registration Successful",
		"Message": "User registered successfully",
	})
}

func login(c echo.Context) error {
	username, password, ok := c.Request().BasicAuth()
	if !ok {
		return echo.ErrUnauthorized
	}

	if err := database.AuthUser(db, username, password); err != nil {
		return echo.ErrUnauthorized
	}

	// JWT-generation
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(jwtKey) // Подпись токена
	if err != nil {
		return echo.ErrInternalServerError // Если не удалось подписать токен
	}

	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"Title": "Login Successful",
		"Token": signedToken,
	})
}

// listProducts выводит список всех продуктов
//func listProducts(c echo.Context) error {
//	products, err := database.GetAllProducts()
//	if err != nil {
//		return c.String(http.StatusInternalServerError, "Could not retrieve products")
//	}
//
//	return c.JSON(http.StatusOK, products)
//}
