package odoo

import (
	"encoding/json"
	"fmt"
	"time"
)

func example() {
	client, err := Connect(
		"base_url",
		"demo",
		"admin@xiatech.co.uk",
		"test",
	)
	if err != nil {
		panic(err)
	}

	// Get a Product to add to our order
	product, err := client.GetProduct(12)
	if err != nil {
		panic(err)
	}

	var products []Product
	if err := json.Unmarshal(product, &products); err != nil {
		panic(err)
	}

	// create a basic order line with the product
	orderLines := [][]interface{}{
		{0, "", OrderLineItem{ProductId: products[0].ID, ProductUom: products[0].UomId[0]}},
	}

	dateOrder := time.Now()
	saleOrder := SaleOrder{
		Name:             "This is another huge sales order",
		PartnerId:        26,
		DateOrder:        dateOrder.Format("2006-01-02"),
		PricelistId:      1,
		CompanyId:        1,
		OrderLine:        orderLines,
		PartnerInvoiceId: 26,
		DisplayName:      "Big Order",
	}

	var inInterface map[string]interface{}
	inrec, _ := json.Marshal(saleOrder)
	json.Unmarshal(inrec, &inInterface)

	id, err := client.Create(SaleOrderModelID, inInterface)
	if err != nil {
		panic(err)
	}

	fmt.Print(id)
}
