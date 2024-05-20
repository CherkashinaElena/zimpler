package zimpler

import (
	"bufio"
	"errors"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
	"os"
	"strconv"
)

// Customer describes customer input data
type Customer struct {
	Name  string
	Candy string
	Eaten int
}

// CustomerFetcher fetches customer data from source
type CustomerFetcher interface {
	Fetch(source string) ([]*Customer, error)
}

type HttpCustomerFetcher struct {
}

// Fetch fetches customer data from web page via http
func (h *HttpCustomerFetcher) Fetch(source string) ([]*Customer, error) {
	resp, err := http.Get(source)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New("can't fetch content of the page")
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	return findCustomers(doc), nil
}

func NewHttpCustomerFetcher() *HttpCustomerFetcher {
	return &HttpCustomerFetcher{}
}

type FileCustomerFetcher struct {
}

// Fetch fetches customer data from a local html file via http
func (f *FileCustomerFetcher) Fetch(filePath string) ([]*Customer, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(bufio.NewReader(file))
	if err != nil {
		return nil, err
	}

	return findCustomers(doc), nil
}

func NewFileCustomerFetcher() *FileCustomerFetcher {
	return &FileCustomerFetcher{}
}

func findCustomers(doc *goquery.Document) []*Customer {
	var customers []*Customer
	doc.Find("section[id=candystore-customers] table[id=\"top.customers\"] > tbody > tr").Each(func(i int, trElem *goquery.Selection) {
		customerData := make([]string, 3)
		trElem.Find("td").Each(func(i int, tdElem *goquery.Selection) {
			if i > 2 {
				log.Printf("skipping unnecessary data for %v\n", customerData[0])
				return
			}

			customerData[i] = tdElem.Text()
		})

		eaten, err := strconv.Atoi(customerData[2])
		if err != nil {
			log.Printf("unexpected data, skipping the customer %v\n", customerData[0])
			return
		}

		customers = append(customers, &Customer{
			Name:  customerData[0],
			Candy: customerData[1],
			Eaten: eaten,
		})
	})

	return customers
}
