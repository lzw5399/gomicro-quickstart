/**
 * @Author: lzw5399
 * @Date: 2020/8/8 23:07
 * @Desc: product结构体
 */
package model

import "strconv"

type Product struct {
	Id   int
	Name string
}

func NewProduct(id int, name string) *Product {
	return &Product{
		Id:   id,
		Name: name,
	}
}

func NewProductList(count int) []*Product {
	products := make([]*Product, 0)
	for i := 0; i < count; i++ {
		products = append(products, NewProduct(i+1, "productName"+strconv.Itoa(i+1)))
	}

	return products
}
