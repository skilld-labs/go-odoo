package odoo_test

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/xiatechs/go-odoo/v2"
)

func TestClient_Unmarshal_Products(t *testing.T) {
	// products.json is what we get after we fetch data from Odoo and serialize it as JSON.
	productsJSON, err := os.ReadFile("testdata/products.json")
	assert.NoError(t, err)

	var products []odoo.Product
	err = json.Unmarshal(productsJSON, &products)
	assert.NoError(t, err)
}
