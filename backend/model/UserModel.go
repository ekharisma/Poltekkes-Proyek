package model

import (
	"github.com/ekharisma/poltekkes-webservice/entity"
	"gorm.io/gorm"
	"log"
	"time"
)

type IUserModel interface {
	GetById(id uint) (*entity.User, error)
	GetAll() ([]*entity.User, error)
	Create(userModel *entity.User) error
	Update(id int64, user *entity.User) error
	Delete(id int64) error
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

func (u UserModel) Update(id int64, user *entity.User) error {
	dbUser, err := u.GetById(uint(id))
	if err != nil {
		return err
	}
	dbUser.Email = user.Email
	dbUser.Name = user.Name
	dbUser.UpdatedAt = time.Time{}
	if err := u.db.Save(&dbUser).Error; err != nil {
		return err
	}
	return nil
}

func (u UserModel) Delete(id int64) error {
	var user *entity.User
	if err := u.db.Delete(&user, id).Error; err != nil {
		return err
	}
	return nil
}
