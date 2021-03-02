package subscribe

import "gorm.io/gorm"

type Repository interface {
	Save(subscribe Subscribe) (Subscribe, error)
	FindAll() ([]Subscribe, error)
	Delete(ID int) (Subscribe, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(subscribe Subscribe) (Subscribe, error) {
	err := r.db.Create(&subscribe).Error
	if err != nil {
		return subscribe, err
	}
	return subscribe, nil
}

func (r *repository) FindAll() ([]Subscribe, error) {
	var subs []Subscribe
	err := r.db.Find(&subs).Error
	if err != nil {
		return subs, err
	}
	return subs, nil
}

func (r *repository) Delete(ID int) (Subscribe, error) {
	var subscribe Subscribe
	err := r.db.Where("id = ?", ID).Delete(&subscribe).Error
	if err != nil {
		return subscribe, err
	}
	return subscribe, nil
}
