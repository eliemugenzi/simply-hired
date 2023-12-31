package service

import (
	"github.com/eliemugenzi/simply-hired/db/models"
	"github.com/eliemugenzi/simply-hired/dto"
	repository "github.com/eliemugenzi/simply-hired/repositories"
	"github.com/mashingan/smapping"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthService interface {
	Register(userDto dto.User) (*gorm.DB, models.User)
	VerifyCredential(email string, password string) (bool, uint64)
}

type authService struct {
	authRepo repository.AuthRepo
}

func NewAuthService(authRepo repository.AuthRepo) *authService {
	return &authService{
		authRepo: authRepo,
	}
}

func (service *authService) Register(userDto dto.User) (*gorm.DB, models.User) {
   userModel := models.User{}
   err := smapping.FillStruct(&userModel, smapping.MapFields(&userDto))
   if err != nil {
	panic(err)
   }
   return service.authRepo.Register(userModel)
}

func (service *authService) VerifyCredential(email string, password string) (bool, uint64) {
   result, user := service.authRepo.FindByEmail(email)
   if result.Error == nil && user.ID != 0 {
	return comparePassword([]byte(user.Password), []byte(password)), uint64(user.ID)
   }

   return false, 0
}


func comparePassword(hashedPass []byte, plainPass []byte) bool {
	err := bcrypt.CompareHashAndPassword(hashedPass, plainPass)

	return err == nil
}


