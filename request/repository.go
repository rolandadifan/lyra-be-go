package request

import "gorm.io/gorm"

type Repository interface {
	Save(request Request) (Request, error)
	Find() ([]Request, error)
	Delete(ID int) (Request, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(request Request) (Request, error) {
	err := r.db.Create(&request).Error
	if err != nil {
		return request, err
	}
	return request, nil
}

func (r *repository) Find() ([]Request, error) {
	var request []Request
	err := r.db.Find(&request).Error
	if err != nil {
		return request, err
	}
	return request, nil
}

func (r *repository) Delete(ID int) (Request, error) {
	var request Request
	err := r.db.Where("id = ?", ID).Delete(&request).Error
	if err != nil {
		return request, err
	}
	return request, nil
}
