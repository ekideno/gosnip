package domain

type URL struct {
	ID       uint   `gorm:"primaryKey;autoIncrement"`
	Slug     string `gorm:"uniqueIndex;size:100;not null"`
	Original string `gorm:"not null"`
}

type URLRepository interface {
	Create(url *URL) error
	GetBySlug(slug string) (*URL, error)
	Delete(slug string) error
	CheckUniqueURL(url string) (bool, error)
	GetByURL(original string) (*URL, error)
	Update(url *URL) error
}

type URLService interface {
	Create(url string, slug string) (*URL, error)
	GetBySlug(slug string) (string, error)
}
