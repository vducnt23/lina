package repositories

import "github.com/jinzhu/gorm"

type RepoProvider interface {
	GetEbookRepo() EbookRepository
}

type repoProviderImpl struct {
	ebookRepo EbookRepository
}

func NewRepoProvider(db *gorm.DB) RepoProvider {
	provider := &repoProviderImpl{}

	provider.ebookRepo = newEbookRepository(db)

	return provider
}

func (e *repoProviderImpl) GetEbookRepo() EbookRepository {
	return e.ebookRepo
}
