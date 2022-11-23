package odoo

import (
	"fmt"
)

// StockInventory represents stock.inventory model.
type StockInventory struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	AccountingDate *Time      `xmlrpc:"accounting_date,omptempty"`
	CategoryId     *Many2One  `xmlrpc:"category_id,omptempty"`
	CompanyId      *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	Date           *Time      `xmlrpc:"date,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Exhausted      *Bool      `xmlrpc:"exhausted,omptempty"`
	Filter         *Selection `xmlrpc:"filter,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	LineIds        *Relation  `xmlrpc:"line_ids,omptempty"`
	LocationId     *Many2One  `xmlrpc:"location_id,omptempty"`
	LotId          *Many2One  `xmlrpc:"lot_id,omptempty"`
	MoveIds        *Relation  `xmlrpc:"move_ids,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	PackageId      *Many2One  `xmlrpc:"package_id,omptempty"`
	PartnerId      *Many2One  `xmlrpc:"partner_id,omptempty"`
	ProductId      *Many2One  `xmlrpc:"product_id,omptempty"`
	State          *Selection `xmlrpc:"state,omptempty"`
	TotalQty       *Float     `xmlrpc:"total_qty,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return nil, fmt.Errorf("stock.inventory was not found with criteria %v", criteria)
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
	return -1, fmt.Errorf("stock.inventory was not found with criteria %v and options %v", criteria, options)
}
