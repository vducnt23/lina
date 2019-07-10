package main

import (
	"fmt"
	"github.com/ducnt114/lina/repositories"
	"github.com/ducnt114/lina/worker"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocolly/colly"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"log"
	"os"
	"strings"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	dborm, err := gorm.Open("mysql", fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8&parseTime=True",
		os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_PORT"), os.Getenv("MYSQL_DBNAME")))
	if err != nil {
		panic(err)
	}
	repositoryProvider := repositories.NewRepoProvider(dborm)

	q := make(chan string)

	crawlerWorker := worker.NewEbookWorker(q, repositoryProvider.GetEbookRepo())
	crawlerWorker.Start()

	c := colly.NewCollector()

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if strings.HasPrefix(e.Attr("href"), "/gp/product/") {
			e.Request.Visit(e.Attr("href"))
		}
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting", r.URL.String())
		q <- r.URL.String()
	})

	c.Visit(os.Getenv("INIT_URL"))
	close(q)
}
