package doc_json

import (
	"encoding/json"
	"fmt"
	"testing"
)

// Product 商品信息
type Product1 struct {
	Name      string
	ProductID int64
	Number    int
	Price     float64
	IsOnSale  bool
}

func Test_Normal(t *testing.T) {

	p := &Product1{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	//结果{"Name":"Xiao mi 6","ProductID":1,"Number":10000,"Price":2499,"IsOnSale":true}
}

// Product _
type Product2 struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"-"` // 表示不进行序列化
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func Test_Self(t *testing.T) {

	p := &Product2{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))

	//结果{"name":"Xiao mi 6","number":10000,"price":2499,"is_on_sale":"false"}
}

// Product _
type Product3 struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,omitempty"`
	Number    int     `json:"number"`
	Price     float64 `json:"price"`
	IsOnSale  bool    `json:"is_on_sale,omitempty"`
}

func Test_omitempty(t *testing.T) {

	p := &Product3{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = false
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 0

	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	// 结果{"name":"Xiao mi 6","number":10000,"price":2499}
}

// Product 商品信息
// Product _
type Product4 struct {
	Name      string  `json:"name"`
	ProductID int64   `json:"product_id,string"`
	Number    int     `json:"number,string"`
	Price     float64 `json:"price,string"`
	IsOnSale  bool    `json:"is_on_sale,string"`
}

func Test_ShiftType(t *testing.T) {
	p := &Product4{}
	p.Name = "Xiao mi 6"
	p.IsOnSale = true
	p.Number = 10000
	p.Price = 2499.00
	p.ProductID = 1
	data, _ := json.Marshal(p)
	fmt.Println(string(data))
	//结果{"name":"Xiao mi 6","product_id":"1","number":"10000","price":"2499","is_on_sale":"true"}
}
