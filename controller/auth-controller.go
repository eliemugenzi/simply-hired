package controller

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/eliemugenzi/simply-hired/dto"
	"github.com/eliemugenzi/simply-hired/serializer"
	service "github.com/eliemugenzi/simply-hired/services"
	"github.com/eliemugenzi/simply-hired/utils"
	"github.com/eliemugenzi/simply-hired/utils/logger"
	"github.com/gin-gonic/gin"
)

type AuthController interface {
	Login(context *gin.Context)
	Register(context *gin.Context)
	VerifyToken(context *gin.Context)
	RefreshToken(context *gin.Context)
}

type authController struct {
	authService service.AuthService
	jwtService service.JwtService
	logger *logger.Logger
}

func NewAuthController(authService service.AuthService, jwtService service.JwtService, logger *logger.Logger) *authController {
   return &authController{
	authService: authService,
	jwtService: jwtService,
	logger: logger,
   }
}

func (controller *authController) Login(context *gin.Context) {
	var loginDto dto.Login
	err := context.ShouldBindJSON(&loginDto)
	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		controller.logger.Error().Err(err).Msg("")
		return
	}

	isValidCredential, userId := controller.authService.VerifyCredential(loginDto.Email, loginDto.Password)

	if isValidCredential {
		tokenPair := controller.jwtService.GenerateTokenPair(userId)
		context.JSON(
			http.StatusOK, 
			utils.GetResponse(http.StatusOK, "Login is successful...", tokenPair),
		)

		return
	}

	context.JSON(
		http.StatusBadRequest,
		 utils.GetResponse(http.StatusBadRequest, "Invalid credentials", nil),
		)

		controller.logger.Error().Msg("Invalid credentials")
}

func (controller * authController) Register(context *gin.Context) {
    var userDto dto.User
	err := context.ShouldBindJSON(&userDto)

	if err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		controller.logger.Error().Err(err).Msg("")
		return
	}

	existingUser := controller.authService.FindUserByEmail(userDto.Email)

	if existingUser.ID != 0 {
       context.JSON(
		http.StatusConflict,
		utils.GetResponse(
			http.StatusConflict,
			"A user with this email already exists",
			nil,
		),
	   )

	   controller.logger.Error().Msg("A user with this email already exists")

	   return
	}

	result, user := controller.authService.Register(userDto)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, result.Error.Error(), nil))
		controller.logger.Error().Err(err).Msg("")
		return
	}

	userSerializer := serializer.UserSerializer{
		User: user,
	}

	context.JSON(
		http.StatusCreated,
		 utils.GetResponse(
			http.StatusCreated,
			 "A user has been successfully created",
	    userSerializer.Response()),
	)
}

func (controller *authController) VerifyToken(context *gin.Context) {
	tokenDto := dto.Token{}

	if err := context.ShouldBindJSON(&tokenDto); err != nil {
		context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
		controller.logger.Error().Err(err).Msg("")
		return
	}

	token, _ := utils.ValidateToken(tokenDto.Token)

	if token == nil || !token.Valid {
		context.AbortWithStatusJSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, "Invalid Token", nil))
		controller.logger.Error().Msg("Invalid Token")
		return
	}

	context.JSON(http.StatusOK, utils.GetResponse(http.StatusOK, "The token is valid", gin.H {
		"is_valid": true,
	}))
}


func (controller *authController) RefreshToken(context *gin.Context) {
   tokenDto := dto.Token{}
   if err := context.ShouldBindJSON(&tokenDto); err != nil {
	  context.JSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
	  controller.logger.Error().Err(err).Msg("")
	  return
   }
   token, err := utils.ValidateToken(tokenDto.Token)
   if token == nil || !token.Valid {
	context.AbortWithStatusJSON(http.StatusBadRequest, utils.GetResponse(http.StatusBadRequest, err.Error(), nil))
	controller.logger.Error().Err(err).Msg("")
	return
   }

   if claims, ok := token.Claims.(jwt.MapClaims); ok {
	context.JSON(
		http.StatusOK,
		utils.GetResponse(
			http.StatusOK, 
			"Token pair ready",
			controller.jwtService.GenerateTokenPair(claims["user_id"])),
		 )
   } else {
	context.AbortWithStatusJSON(
		http.StatusBadRequest,
		utils.GetResponse(
			http.StatusBadRequest,
			"Failed to claim a token",
			nil,
		),
	)

	controller.logger.Error().Msg("Failed to claim a token")
   }
}