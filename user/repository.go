package user

import "gorm.io/gorm"

type Repository interface {
	Save(user User) (User, error)
	FindByEmail(email string) (User, error)
	FindById(id uint) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindByEmail(email string) (User, error) {
	var user User
	if err := r.db.Where("email = ?", email).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (r *repository) FindById(id uint) (User, error) {
	var user User
	if err := r.db.Where("id = ?", id).Find(&user).Error; err != nil {
		return user, err
	}
	return user, nil
}
