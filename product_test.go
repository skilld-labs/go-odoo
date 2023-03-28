package odoo_test

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	tc "github.com/testcontainers/testcontainers-go/modules/compose"
	"github.com/xiatechs/XFuze/testutil"

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

func TestClient_CreateProduct(t *testing.T) {
	compose, err := tc.NewDockerCompose("testdata/docker-compose.yml")
	assert.NoError(t, err, "NewDockerComposeAPI()")

	t.Cleanup(func() {
		assert.NoError(t, compose.Down(
			context.Background(),
			tc.RemoveOrphans(true),
			tc.RemoveImagesLocal,
		), "compose.Down()")
	})

	ctx, cancel := context.WithCancel(context.Background())
	t.Cleanup(cancel)

	assert.NoError(t, compose.Up(ctx, tc.Wait(true)), "compose.Up()")

	// The Odoo instance can take a while to boot up, keep trying every 5 seconds for 1min30
	var client *odoo.Client
	assert.NoError(t, testutil.WaitUntil(3*time.Second, 90*time.Second, func() (bool, error) {
		client, err = odoo.Connect(
			"http://localhost:8091",
			"xiatech_test",
			"admin",
			"admin",
		)
		if err != nil {
			t.Log(err.Error())
			return false, nil
		}
		t.Log("successfully connected to Odoo")
		return true, nil
	}))

	// NOTE: Name, Type, Category ID, Tracking, Unit of Measure ID & Unit of Measure Product ID are mandatory params.
	_, err = client.CreateProduct(
		odoo.ProductTemplate{
			Name:          uuid.New().String(),
			Type:          "product",
			ResponsibleId: true,
			CategId:       8,
			UomId:         1,
			UomPoId:       1,
			Tracking:      "none",
		})
	assert.NoError(t, err)
}
