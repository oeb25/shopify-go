package shopify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

type Shopify struct {
	ShopName string
	APIkey   string
	Password string
}

func (s *Shopify) BuildUrl() string {
	return "https://" + s.APIkey + ":" + s.Password + "@" + s.ShopName + ".myshopify.com"
}

func (s *Shopify) Get(url string) (*http.Response, error) {
	return http.Get(s.BuildUrl() + "/admin/products.json")
}

func (s *Shopify) GetInto(url string, v interface{}) error {
	resp, err := s.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(body, v)
	if err != nil {
		return err
	}

	return nil
}

type Image struct {
	Src string `json:"src"`
}

type Option struct {
	Name string `json:"name"`
}

type Variant struct {
	Title string `json:"title"`
}

type Product struct {
	BodyHTML       string    `json:"body_html"`
	CreatedAt      time.Time `json:"created_at"`
	Handle         string    `json:"handle"`
	ID             int       `json:"id"`
	Image          Image     `json:"image"`
	Images         []Image   `json:"images"`
	Options        []Option  `json:"images"`
	ProductType    string    `json:"product_type"`
	PublishedAt    time.Time `json:"published_at"`
	PublishedScope string    `json:"product_scope"`
	Tags           string    `json:"tags"`
	Title          string    `json:"title"`
	Variants       []Variant `json:"variants"`
	Vendor         string    `json:"vendor"`
}

func (p *Product) GetMetafields(shop *Shopify) (map[string]interface{}, error) {
	metafields := make(map[string]interface{})

	err := shop.GetInto("/admin/products/"+strconv.Itoa(p.ID)+"/metafields.json", &metafields)
	if err != nil {
		return nil, err
	}

	return metafields, nil
}

func (s *Shopify) GetProducts() ([]Product, error) {
	var dd map[string][]Product
	err := s.GetInto("/admin/products.json", &dd)
	if err != nil {
		return nil, err
	}

	return dd["products"], nil
}
