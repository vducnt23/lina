package parser

import (
	"errors"
	"github.com/ducnt114/lina/models"
	"log"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// ItemParserService --
type ItemParser interface {
	ParseEbookPage(pageURL string) (*models.Ebook, error)
	ParseTShirtPage(pageURL string) (*models.Tshirt, error)
}

type itemParserImpl struct {
}

// NewItemParserService --
func NewItemParser() ItemParser {
	return &itemParserImpl{}
}

func (i *itemParserImpl) ParseEbookPage(pageURL string) (*models.Ebook, error) {
	startTime := time.Now()
	ebook := &models.Ebook{}

	httpClient := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, pageURL, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error when send get request, detail: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("Status code error: ", resp.StatusCode)
		return nil, errors.New("Bad status code received")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#ebooksProductTitle").Each(func(i int, s *goquery.Selection) {
		ebook.Title = strings.TrimSpace(s.Text())
	})

	doc.Find("a.a-link-normal.contributorNameID").Each(func(i int, s *goquery.Selection) {
		ebook.Author = strings.TrimSpace(s.Text())
	})

	doc.Find("span.a-size-base.a-color-price.a-color-price").Each(func(i int, s *goquery.Selection) {
		priceString := strings.TrimSpace(s.Text())
		items := strings.Split(priceString, "$")

		price, err := strconv.ParseFloat(items[1], 64)
		if err != nil {
			log.Println("Parse price fail, detail: ", err)
			return
		}
		ebook.Price = math.Ceil(price)
	})

	doc.Find("#ebooksImgBlkFront").Each(func(i int, s *goquery.Selection) {
		imageSrc, exist := s.Attr("src")
		if exist {
			ebook.CoverImageURL = imageSrc
		}
	})

	doc.Find("#productDetailsTable").Each(func(i int, s *goquery.Selection) {
		detailTableText := s.Text()
		items := strings.Split(detailTableText, "ASIN:")
		if len(items) == 2 {
			items1 := strings.Split(items[1], "\n")
			if len(items1) >= 1 {
				ebook.ASIN = strings.TrimSpace(items1[0])
			}
		}
	})

	endTime := time.Now()
	processTime := endTime.Unix() - startTime.Unix()

	log.Println("Process time: ", processTime)

	return ebook, nil
}

func (i *itemParserImpl) ParseTShirtPage(pageURL string) (*models.Tshirt, error) {
	tshirt := &models.Tshirt{}

	httpClient := &http.Client{}

	req, err := http.NewRequest(http.MethodGet, pageURL, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.169 Safari/537.36")

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Println("Error when send get request, detail: ", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		log.Println("Status code error: ", resp.StatusCode)
		return nil, errors.New("Bad status code received")
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("#productTitle").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		tshirt.Title = strings.TrimSpace(s.Text())
	})

	doc.Find("#priceblock_ourprice").Each(func(i int, s *goquery.Selection) {
		// For each item found, get the band and title
		priceString := s.Text()
		items := strings.Split(priceString, "$")

		price, err := strconv.ParseFloat(items[1], 64)
		if err != nil {
			log.Println("Parse price fail, detail: ", err)
			return
		}
		tshirt.Price = price
	})

	return tshirt, nil
}
