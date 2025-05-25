package service

import (
	"errors"
	"fmt"
	"github.com/ekideno/gosnip/internal/domain"
	"github.com/ekideno/gosnip/internal/util"
	"gorm.io/gorm"
)

type URLService struct {
	URLRepo domain.URLRepository
}

func NewURLService(urlRepo domain.URLRepository) *URLService {
	return &URLService{URLRepo: urlRepo}
}
func (s *URLService) Create(original string, slug string) (*domain.URL, error) {
	existing, err := s.URLRepo.GetByURL(original)
	if err == nil {
		if slug != "" && slug != existing.Slug {
			conflict, err := s.URLRepo.GetBySlug(slug)
			if err == nil && conflict != nil {
				return nil, fmt.Errorf("code '%s' is already taken", slug)
			}
			if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, err
			}
			existing.Slug = slug
			if err := s.URLRepo.Update(existing); err != nil {
				return nil, err
			}
		}
		return existing, nil
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if slug == "" {
		for {
			slug = util.GenerateRandomString()
			if _, err := s.URLRepo.GetBySlug(slug); errors.Is(err, gorm.ErrRecordNotFound) {
				break
			}
		}
	} else {
		_, err := s.URLRepo.GetBySlug(slug)
		if err == nil {
			return nil, fmt.Errorf("code '%s' is already taken", slug)
		}
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	url := &domain.URL{
		Original: original,
		Slug:     slug,
	}

	if err := s.URLRepo.Create(url); err != nil {
		return nil, err
	}

	return url, nil
}

func (s *URLService) GetBySlug(slug string) (string, error) {
	snipURL, err := s.URLRepo.GetBySlug(slug)
	if err != nil {
		return "", err
	}
	return snipURL.Original, nil
}
