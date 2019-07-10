package repositories

import (
	"github.com/ducnt114/lina/models"
	"github.com/jinzhu/gorm"
)

type EbookRepository interface {
	Save(ebook *models.Ebook) error
	FindByASIN(asin string) (*models.Ebook, error)
}

type ebookRepoImpl struct {
	db *gorm.DB
}

func newEbookRepository(db *gorm.DB) EbookRepository {
	return &ebookRepoImpl{
		db: db,
	}
}

func (e *ebookRepoImpl) Save(ebook *models.Ebook) error {
	return e.db.Save(ebook).Error
}

func (e *ebookRepoImpl) FindByASIN(asin string) (*models.Ebook, error) {
	ebook := &models.Ebook{}
	err := e.db.Model(&models.Ebook{}).Where("asin = ?", asin).First(ebook).Error
	return ebook, err
}
