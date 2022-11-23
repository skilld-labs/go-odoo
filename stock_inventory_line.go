package odoo

import (
	"fmt"
)

// StockInventoryLine represents stock.inventory.line model.
type StockInventoryLine struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	InventoryId         *Many2One  `xmlrpc:"inventory_id,omptempty"`
	InventoryLocationId *Many2One  `xmlrpc:"inventory_location_id,omptempty"`
	LocationId          *Many2One  `xmlrpc:"location_id,omptempty"`
	LocationName        *String    `xmlrpc:"location_name,omptempty"`
	PackageId           *Many2One  `xmlrpc:"package_id,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	ProdLotId           *Many2One  `xmlrpc:"prod_lot_id,omptempty"`
	ProdlotName         *String    `xmlrpc:"prodlot_name,omptempty"`
	ProductCode         *String    `xmlrpc:"product_code,omptempty"`
	ProductId           *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductName         *String    `xmlrpc:"product_name,omptempty"`
	ProductQty          *Float     `xmlrpc:"product_qty,omptempty"`
	ProductUomId        *Many2One  `xmlrpc:"product_uom_id,omptempty"`
	State               *Selection `xmlrpc:"state,omptempty"`
	TheoreticalQty      *Float     `xmlrpc:"theoretical_qty,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockInventoryLines represents array of stock.inventory.line model.
type StockInventoryLines []StockInventoryLine

// StockInventoryLineModel is the odoo model name.
const StockInventoryLineModel = "stock.inventory.line"

// Many2One convert StockInventoryLine to *Many2One.
func (sil *StockInventoryLine) Many2One() *Many2One {
	return NewMany2One(sil.Id.Get(), "")
}

// CreateStockInventoryLine creates a new stock.inventory.line model and returns its id.
func (c *Client) CreateStockInventoryLine(sil *StockInventoryLine) (int64, error) {
	return c.Create(StockInventoryLineModel, sil)
}

// UpdateStockInventoryLine updates an existing stock.inventory.line record.
func (c *Client) UpdateStockInventoryLine(sil *StockInventoryLine) error {
	return c.UpdateStockInventoryLines([]int64{sil.Id.Get()}, sil)
}

// UpdateStockInventoryLines updates existing stock.inventory.line records.
// All records (represented by ids) will be updated by sil values.
func (c *Client) UpdateStockInventoryLines(ids []int64, sil *StockInventoryLine) error {
	return c.Update(StockInventoryLineModel, ids, sil)
}

// DeleteStockInventoryLine deletes an existing stock.inventory.line record.
func (c *Client) DeleteStockInventoryLine(id int64) error {
	return c.DeleteStockInventoryLines([]int64{id})
}

// DeleteStockInventoryLines deletes existing stock.inventory.line records.
func (c *Client) DeleteStockInventoryLines(ids []int64) error {
	return c.Delete(StockInventoryLineModel, ids)
}

// GetStockInventoryLine gets stock.inventory.line existing record.
func (c *Client) GetStockInventoryLine(id int64) (*StockInventoryLine, error) {
	sils, err := c.GetStockInventoryLines([]int64{id})
	if err != nil {
		return nil, err
	}
	if sils != nil && len(*sils) > 0 {
		return &((*sils)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.inventory.line not found", id)
}

// GetStockInventoryLines gets stock.inventory.line existing records.
func (c *Client) GetStockInventoryLines(ids []int64) (*StockInventoryLines, error) {
	sils := &StockInventoryLines{}
	if err := c.Read(StockInventoryLineModel, ids, nil, sils); err != nil {
		return nil, err
	}
	return sils, nil
}

// FindStockInventoryLine finds stock.inventory.line record by querying it with criteria.
func (c *Client) FindStockInventoryLine(criteria *Criteria) (*StockInventoryLine, error) {
	sils := &StockInventoryLines{}
	if err := c.SearchRead(StockInventoryLineModel, criteria, NewOptions().Limit(1), sils); err != nil {
		return nil, err
	}
	if sils != nil && len(*sils) > 0 {
		return &((*sils)[0]), nil
	}
	return nil, fmt.Errorf("stock.inventory.line was not found with criteria %v", criteria)
}

// FindStockInventoryLines finds stock.inventory.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryLines(criteria *Criteria, options *Options) (*StockInventoryLines, error) {
	sils := &StockInventoryLines{}
	if err := c.SearchRead(StockInventoryLineModel, criteria, options, sils); err != nil {
		return nil, err
	}
	return sils, nil
}

// FindStockInventoryLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockInventoryLineModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockInventoryLineId finds record id by querying it with criteria.
func (c *Client) FindStockInventoryLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockInventoryLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.inventory.line was not found with criteria %v and options %v", criteria, options)
}
