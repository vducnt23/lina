package repositories

import (
	"github.com/ducnt114/lina/models"
	"github.com/jinzhu/gorm"
)

type EbookRepository interface {
	Save(ebook *models.Ebook) error
}

type ebookRepoImpl struct {
	db *gorm.DB
}

func NewEbookRepository(db *gorm.DB) EbookRepository {
	return &ebookRepoImpl{
		db: db,
	}
}

func (e *ebookRepoImpl) Save(ebook *models.Ebook) error {
	return e.db.Save(ebook).Error
}
