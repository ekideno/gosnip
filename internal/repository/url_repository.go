package repository

import (
	"github.com/ekideno/gosnip/internal/domain"
	"gorm.io/gorm"
)

type URLRepository struct {
	db *gorm.DB
}

func NewURLRepository(db *gorm.DB) domain.URLRepository {
	return &URLRepository{db: db}
}

func (r URLRepository) Create(url *domain.URL) error {
	err := r.db.Create(url).Error
	return err
}

func (r URLRepository) GetBySlug(slug string) (*domain.URL, error) {
	url := &domain.URL{}
	err := r.db.Where("slug = ?", slug).First(url).Error
	return url, err
}

func (r URLRepository) Delete(slug string) error {
	err := r.db.Where("slug = ?", slug).Delete(&domain.URL{}).Error
	return err
}

func (r *URLRepository) CheckUniqueURL(original string) (bool, error) {
	var url domain.URL
	err := r.db.Where("original = ?", original).First(&url).Error

	if err == gorm.ErrRecordNotFound {
		return true, nil
	}
	if err != nil {
		return false, err
	}
	return false, nil
}

func (r *URLRepository) GetByURL(original string) (*domain.URL, error) {
	var url domain.URL
	err := r.db.Where("original = ?", original).First(&url).Error
	return &url, err
}

func (r *URLRepository) Update(url *domain.URL) error {
	err := r.db.Save(url).Error
	return err
}
