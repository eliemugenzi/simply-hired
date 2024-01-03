package repository

import (
	"fmt"

	"github.com/eliemugenzi/simply-hired/db/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type AuthRepo interface {
	Register(user models.User) (*gorm.DB, models.User)
	FindByEmail(email string) (*gorm.DB, models.User)
	FindById(id uint) (*gorm.DB, models.User)
}

type authRepo struct {
	db *gorm.DB
}


func NewAuthRepo(db *gorm.DB) *authRepo {
	return &authRepo{
		db: db,
	}
}

func (repo *authRepo) Register(user models.User) (*gorm.DB, models.User) {
	user.Password = hashAndSalt([]byte(user.Password))
	userResult := repo.db.Create(&user)
	return userResult, user
}

func (repo *authRepo) FindByEmail(email string) (*gorm.DB, models.User) {
	user := models.User{}
	userResult := repo.db.Where("email = ?", email).Take(&user)
	fmt.Println("EMAIL BY USER", user)
	return userResult, user
}

func (repo *authRepo) FindById(id uint) (*gorm.DB, models.User) {
	user := models.User{}
	userResult := repo.db.Where("id = ?", id).Take(&user)

	return userResult, user
}

func hashAndSalt(password []byte) string {
	hash, err := bcrypt.GenerateFromPassword(password, bcrypt.MinCost)

	if err != nil {
		panic("Failed to hash password")
	}

	return string(hash)
}