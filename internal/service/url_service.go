package service

import (
	"github.com/ekideno/gosnip/internal/domain"
	"github.com/ekideno/gosnip/internal/util"
)

type URLService struct {
	URLRepo domain.URLRepository
}

func NewURLService(urlRepo domain.URLRepository) *URLService {
	return &URLService{URLRepo: urlRepo}
}

func (s *URLService) CreateRandom(original string) (*domain.URL, error) {
	var url domain.URL
	if uniqueURL, err := s.URLRepo.CheckUniqueURL(original); err != nil {
		return nil, err
	} else if !uniqueURL {
		foundURL, err := s.URLRepo.GetByURL(original) // TODO
		return foundURL, err
	}

	url.Original = original
	url.Slug = util.GenerateRandomString()

	err := s.URLRepo.Create(&url)
	return &url, err
}

func (s *URLService) CreateSpecial(url string, slug string) error {
	var snipURL domain.URL
	snipURL.Original = url
	snipURL.Slug = slug
	err := s.URLRepo.Create(&snipURL)
	return err
}

func (s *URLService) GetBySlug(slug string) (string, error) {
	snipURL, err := s.URLRepo.GetBySlug(slug)
	if err != nil {
		return "", err
	}
	return snipURL.Original, nil
}
