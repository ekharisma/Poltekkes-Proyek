package model

import (
	"github.com/ekharisma/poltekkes-webservice/entity"
	"gorm.io/gorm"
	"log"
)

type IUserModel interface {
	GetById(id uint) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Create(userModel *entity.User) error
}

type UserModel struct {
	db *gorm.DB
}

func CreateUserModel(db *gorm.DB) IUserModel {
	o := UserModel{db: db}
	if err := o.db.AutoMigrate(&entity.User{}); err != nil {
		log.Panicln("Error migrating user entity. Reason : ", err.Error())
	}
	return &o
}

func (u UserModel) GetById(id uint) (*entity.User, error) {
	user := &entity.User{}
	err := u.db.First(user, id).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u UserModel) GetAll() ([]*entity.User, error) {
	var users []*entity.User
	err := u.db.Find(&users).Error
	if err != nil {
		log.Panicln("Error can't get all user data. Reason : ", err.Error())
		return nil, err
	}
	return users, nil
}

func (u UserModel) Create(users *entity.User) error {
	if err := u.db.Create(users).Error; err != nil {
		return err
	}
	return nil
}
