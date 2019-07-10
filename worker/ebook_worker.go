package worker

import (
	"github.com/ducnt114/lina/parser"
	"github.com/ducnt114/lina/repositories"
	"github.com/jinzhu/gorm"
	"log"
)

type ebookWorker struct {
	q          <-chan string
	itemParser parser.ItemParser
	ebookRepo  repositories.EbookRepository
}

func NewEbookWorker(q <-chan string, ebookRepo repositories.EbookRepository) BaseWorker {
	ebookWorker := &ebookWorker{
		q:          q,
		itemParser: parser.NewItemParser(),
		ebookRepo:  ebookRepo,
	}
	return ebookWorker
}

func (w *ebookWorker) Start() {
	go func() {
		for {
			pageURL, ok := <-w.q
			if ok {
				err := w.processURL(pageURL)
				if err != nil {
					log.Printf("Error when processURL: %v, detail: %v", pageURL, err)
				}
			} else {
				return
			}

		}
	}()
}

func (w *ebookWorker) processURL(pageURL string) error {
	ebook, err := w.itemParser.ParseEbookPage(pageURL)
	if err != nil {
		log.Printf("Error when ParseEbookPage, pageURL: %v, detail: %v", pageURL, err)
		return err
	}
	oldEbook, err := w.ebookRepo.FindByASIN(ebook.ASIN)
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Error when FindByASIN, detail: ", err)
		return err
	}
	if oldEbook.ID != 0 {
		log.Printf("Ebook found with ASIN: %v. Skip this page.", ebook.ASIN)
		return nil
	}
	err = w.ebookRepo.Save(ebook)
	if err != nil {
		log.Printf("Error when save book with ASIN: %v, detail: %v", ebook.ASIN, err)
		return err
	}

	return nil
}
