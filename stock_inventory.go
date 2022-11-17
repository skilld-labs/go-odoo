package odoo

import (
	"fmt"
)

// StockInventory represents stock.inventory model.
type StockInventory struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	AccountingDate *Time      `xmlrpc:"accounting_date,omitempty"`
	CategoryId     *Many2One  `xmlrpc:"category_id,omitempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	Date           *Time      `xmlrpc:"date,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Exhausted      *Bool      `xmlrpc:"exhausted,omitempty"`
	Filter         *Selection `xmlrpc:"filter,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	LineIds        *Relation  `xmlrpc:"line_ids,omitempty"`
	LocationId     *Many2One  `xmlrpc:"location_id,omitempty"`
	LotId          *Many2One  `xmlrpc:"lot_id,omitempty"`
	MoveIds        *Relation  `xmlrpc:"move_ids,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	PackageId      *Many2One  `xmlrpc:"package_id,omitempty"`
	PartnerId      *Many2One  `xmlrpc:"partner_id,omitempty"`
	ProductId      *Many2One  `xmlrpc:"product_id,omitempty"`
	State          *Selection `xmlrpc:"state,omitempty"`
	TotalQty       *Float     `xmlrpc:"total_qty,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockInventorys represents array of stock.inventory model.
type StockInventorys []StockInventory

// StockInventoryModel is the odoo model name.
const StockInventoryModel = "stock.inventory"

// Many2One convert StockInventory to *Many2One.
func (si *StockInventory) Many2One() *Many2One {
	return NewMany2One(si.Id.Get(), "")
}

// CreateStockInventory creates a new stock.inventory model and returns its id.
func (c *Client) CreateStockInventory(si *StockInventory) (int64, error) {
	return c.Create(StockInventoryModel, si)
}

// UpdateStockInventory updates an existing stock.inventory record.
func (c *Client) UpdateStockInventory(si *StockInventory) error {
	return c.UpdateStockInventorys([]int64{si.Id.Get()}, si)
}

// UpdateStockInventorys updates existing stock.inventory records.
// All records (represented by ids) will be updated by si values.
func (c *Client) UpdateStockInventorys(ids []int64, si *StockInventory) error {
	return c.Update(StockInventoryModel, ids, si)
}

// DeleteStockInventory deletes an existing stock.inventory record.
func (c *Client) DeleteStockInventory(id int64) error {
	return c.DeleteStockInventorys([]int64{id})
}

// DeleteStockInventorys deletes existing stock.inventory records.
func (c *Client) DeleteStockInventorys(ids []int64) error {
	return c.Delete(StockInventoryModel, ids)
}

// GetStockInventory gets stock.inventory existing record.
func (c *Client) GetStockInventory(id int64) (*StockInventory, error) {
	sis, err := c.GetStockInventorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if sis != nil && len(*sis) > 0 {
		return &((*sis)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.inventory not found", id)
}

// GetStockInventorys gets stock.inventory existing records.
func (c *Client) GetStockInventorys(ids []int64) (*StockInventorys, error) {
	sis := &StockInventorys{}
	if err := c.Read(StockInventoryModel, ids, nil, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindStockInventory finds stock.inventory record by querying it with criteria.
func (c *Client) FindStockInventory(criteria *Criteria) (*StockInventory, error) {
	sis := &StockInventorys{}
	if err := c.SearchRead(StockInventoryModel, criteria, NewOptions().Limit(1), sis); err != nil {
		return nil, err
	}
	if sis != nil && len(*sis) > 0 {
		return &((*sis)[0]), nil
	}
	return nil, fmt.Errorf("no stock.inventory was found with criteria %v", criteria)
}

// FindStockInventorys finds stock.inventory records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventorys(criteria *Criteria, options *Options) (*StockInventorys, error) {
	sis := &StockInventorys{}
	if err := c.SearchRead(StockInventoryModel, criteria, options, sis); err != nil {
		return nil, err
	}
	return sis, nil
}

// FindStockInventoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockInventoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockInventoryId finds record id by querying it with criteria.
func (c *Client) FindStockInventoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockInventoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no stock.inventory was found with criteria %v and options %v", criteria, options)
}
