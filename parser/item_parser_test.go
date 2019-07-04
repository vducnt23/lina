package parser

import (
	"testing"
)

func TestItemParserImpl_ParseEbookPage(t *testing.T) {
	itemParser := NewItemParser()

	pageURL := "https://www.amazon.com/gp/product/B07HZ2Q1MK?storeType=ebooks&pf_rd_t=40901&pd_rd_i=B07HZ2Q1MK&pf_rd_m=A2R2RITDJNW1Q6&pageType=STOREFRONT&pf_rd_p=f6df1b3c-b603-43b2-a799-fe9150612dc1&pf_rd_r=0R2JTYBVPA9MJ7ATTYTN&pd_rd_wg=TZyw1&pf_rd_s=merchandised-search-3&ref_=dbs_f_ebk_rwt_wigo_HCcsf_ms3_kmw_ae642e4b-ad2c-4298-8dd0-2f96f1b&pd_rd_w=CCI1S&pf_rd_i=154606011&pd_rd_r=fa54b7ce-bee5-4883-9e3e-605a98db33ca"

	ebook, err := itemParser.ParseEbookPage(pageURL)
	if err != nil {
		t.Fatal(err)
	}

	if ebook.Title != "City of Girls: A Novel" {
		t.Error("Title must be: ", "City of Girls: A Novel")
		t.Fail()
	}

	if ebook.Author != "Elizabeth Gilbert" {
		t.Error("Author must be: ", "Elizabeth Gilbert", ", actual: ", ebook.Author)
		t.Fail()
	}

	if ebook.Price != 13.0 {
		t.Error("Price must be: ", 13.0, ", actual: ", ebook.Price)
		t.Fail()
	}

	if ebook.CoverImageURL != "https://images-na.ssl-images-amazon.com/images/I/51xcfAfbZzL._SY346_.jpg" {
		t.Error("CoverImageURL must be: ", "https://images-na.ssl-images-amazon.com/images/I/51xcfAfbZzL._SY346_.jpg",
			", actual: ", ebook.CoverImageURL)
		t.Fail()
	}

	if ebook.ASIN != "B07HZ2Q1MK" {
		t.Error("ASIN must be: ", "B07HZ2Q1MK", ", actual: ", ebook.ASIN)
		t.Fail()
	}
}

func TestItemParserImpl_ParseTShirtPage(t *testing.T) {
	itemParser := NewItemParser()

	tshirt, err := itemParser.ParseTShirtPage("https://www.amazon.com/Vacay-Shirt-Funny-Family-Vacation/dp/B07RRFJ16Z/ref=sr_1_3?keywords=family+t+shirt&qid=1561906472&s=apparel&sr=1-3")
	if err != nil {
		t.Fatal(err)
	}

	if tshirt.Title != "Vacay Mode T Shirt Funny Family Vacation Gift Men Women" {
		t.Error("Product title must be: ", "Vacay Mode T Shirt Funny Family Vacation Gift Men Women")
		t.Fail()
	}

	if tshirt.Price != 14.99 {
		t.Error("Procut price must be: ", 14.99)
		t.Fail()
	}
}
