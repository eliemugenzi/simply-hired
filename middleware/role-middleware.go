package middleware

import (
	"net/http"

	repository "github.com/eliemugenzi/simply-hired/repositories"
	"github.com/eliemugenzi/simply-hired/utils"
	"github.com/gin-gonic/gin"
)


type RoleMiddleware interface {
	AuthorizeRole(role string) gin.HandlerFunc
}

type roleMiddleware struct {
	authRepo repository.AuthRepo
}

func NewRoleMiddleware(authRepo repository.AuthRepo) *roleMiddleware {
	return &roleMiddleware{
		authRepo: authRepo,
	}
}

func (middleware *roleMiddleware) AuthorizeRole(role string) gin.HandlerFunc {

	return func(ctx *gin.Context) {
		tokenString := utils.GetTokenString(ctx)

		if tokenString == "" {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				 utils.GetResponse(http.StatusUnauthorized, "Unauthorized access", nil))

				 return
		}

		token, err := utils.ValidateToken(tokenString)

		if token == nil || !token.Valid {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetResponse(http.StatusUnauthorized, err.Error(), nil))

			return
		}

		userId, err := utils.GetUserIdFromToken(tokenString)

		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, utils.GetResponse(http.StatusUnauthorized, err.Error(), nil))
		}

		_, foundUser :=middleware.authRepo.FindById(userId)

		if(foundUser.Role != role) {
			ctx.AbortWithStatusJSON(http.StatusForbidden, utils.GetResponse(http.StatusForbidden, "Forbidden access", nil))

			return
		}

		ctx.Set("user_id", userId)


	}
}