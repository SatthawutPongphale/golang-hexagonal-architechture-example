package http

import (
	userrepo "golang-hexagonal-architechture-example/internal/adapter/db"
	"golang-hexagonal-architechture-example/internal/application"
	"golang-hexagonal-architechture-example/internal/domain"
	"golang-hexagonal-architechture-example/pkg/db"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, database *db.PostgresDatabase) {
	userRepository := userrepo.NewUserRepository(database)
	userService := application.NewUserService(userRepository)

	userGroup := r.Group("/users")
	userGroup.POST("/", func(c *gin.Context) {
		var input domain.UserInput
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		user, err := userService.CreateUser(input)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, user)
	})
}
