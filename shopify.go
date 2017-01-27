package shopify

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
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

func (s *Shopify) GetProducts() ([]Product, error) {
	resp, err := s.Get("/admin/products.json")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var dd map[string][]Product

	_ = json.Unmarshal(body, &dd)

	return dd["products"], nil
}
