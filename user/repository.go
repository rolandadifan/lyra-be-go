package user

import "gorm.io/gorm"

type Repository interface {
	Register(user User) (User, error)
	FindByEmail(email string) (User, error)
	UserToken(token Token) (Token, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Register(user User) (User, error) {
	err := r.db.Create(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	err := r.db.Where("email = ?", email).Find(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) UserToken(token Token) (Token, error) {
	err := r.db.Create(&token).Error
	if err != nil {
		return token, err
	}
	return token, nil
}
