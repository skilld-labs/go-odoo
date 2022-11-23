package odoo

import (
	"fmt"
)

// StockPickingType represents stock.picking.type model.
type StockPickingType struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	BarcodeNomenclatureId  *Many2One  `xmlrpc:"barcode_nomenclature_id,omptempty"`
	Code                   *Selection `xmlrpc:"code,omptempty"`
	Color                  *Int       `xmlrpc:"color,omptempty"`
	CountPicking           *Int       `xmlrpc:"count_picking,omptempty"`
	CountPickingBackorders *Int       `xmlrpc:"count_picking_backorders,omptempty"`
	CountPickingDraft      *Int       `xmlrpc:"count_picking_draft,omptempty"`
	CountPickingLate       *Int       `xmlrpc:"count_picking_late,omptempty"`
	CountPickingReady      *Int       `xmlrpc:"count_picking_ready,omptempty"`
	CountPickingWaiting    *Int       `xmlrpc:"count_picking_waiting,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	DefaultLocationDestId  *Many2One  `xmlrpc:"default_location_dest_id,omptempty"`
	DefaultLocationSrcId   *Many2One  `xmlrpc:"default_location_src_id,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	LastDonePicking        *String    `xmlrpc:"last_done_picking,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	RatePickingBackorders  *Int       `xmlrpc:"rate_picking_backorders,omptempty"`
	RatePickingLate        *Int       `xmlrpc:"rate_picking_late,omptempty"`
	ReturnPickingTypeId    *Many2One  `xmlrpc:"return_picking_type_id,omptempty"`
	Sequence               *Int       `xmlrpc:"sequence,omptempty"`
	SequenceId             *Many2One  `xmlrpc:"sequence_id,omptempty"`
	ShowEntirePacks        *Bool      `xmlrpc:"show_entire_packs,omptempty"`
	ShowOperations         *Bool      `xmlrpc:"show_operations,omptempty"`
	ShowReserved           *Bool      `xmlrpc:"show_reserved,omptempty"`
	UseCreateLots          *Bool      `xmlrpc:"use_create_lots,omptempty"`
	UseExistingLots        *Bool      `xmlrpc:"use_existing_lots,omptempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// StockPickingTypes represents array of stock.picking.type model.
type StockPickingTypes []StockPickingType

// StockPickingTypeModel is the odoo model name.
const StockPickingTypeModel = "stock.picking.type"

// Many2One convert StockPickingType to *Many2One.
func (spt *StockPickingType) Many2One() *Many2One {
	return NewMany2One(spt.Id.Get(), "")
}

// CreateStockPickingType creates a new stock.picking.type model and returns its id.
func (c *Client) CreateStockPickingType(spt *StockPickingType) (int64, error) {
	return c.Create(StockPickingTypeModel, spt)
}

// UpdateStockPickingType updates an existing stock.picking.type record.
func (c *Client) UpdateStockPickingType(spt *StockPickingType) error {
	return c.UpdateStockPickingTypes([]int64{spt.Id.Get()}, spt)
}

// UpdateStockPickingTypes updates existing stock.picking.type records.
// All records (represented by ids) will be updated by spt values.
func (c *Client) UpdateStockPickingTypes(ids []int64, spt *StockPickingType) error {
	return c.Update(StockPickingTypeModel, ids, spt)
}

// DeleteStockPickingType deletes an existing stock.picking.type record.
func (c *Client) DeleteStockPickingType(id int64) error {
	return c.DeleteStockPickingTypes([]int64{id})
}

// DeleteStockPickingTypes deletes existing stock.picking.type records.
func (c *Client) DeleteStockPickingTypes(ids []int64) error {
	return c.Delete(StockPickingTypeModel, ids)
}

// GetStockPickingType gets stock.picking.type existing record.
func (c *Client) GetStockPickingType(id int64) (*StockPickingType, error) {
	spts, err := c.GetStockPickingTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if spts != nil && len(*spts) > 0 {
		return &((*spts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of stock.picking.type not found", id)
}

// GetStockPickingTypes gets stock.picking.type existing records.
func (c *Client) GetStockPickingTypes(ids []int64) (*StockPickingTypes, error) {
	spts := &StockPickingTypes{}
	if err := c.Read(StockPickingTypeModel, ids, nil, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindStockPickingType finds stock.picking.type record by querying it with criteria.
func (c *Client) FindStockPickingType(criteria *Criteria) (*StockPickingType, error) {
	spts := &StockPickingTypes{}
	if err := c.SearchRead(StockPickingTypeModel, criteria, NewOptions().Limit(1), spts); err != nil {
		return nil, err
	}
	if spts != nil && len(*spts) > 0 {
		return &((*spts)[0]), nil
	}
	return nil, fmt.Errorf("stock.picking.type was not found with criteria %v", criteria)
}

// FindStockPickingTypes finds stock.picking.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingTypes(criteria *Criteria, options *Options) (*StockPickingTypes, error) {
	spts := &StockPickingTypes{}
	if err := c.SearchRead(StockPickingTypeModel, criteria, options, spts); err != nil {
		return nil, err
	}
	return spts, nil
}

// FindStockPickingTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockPickingTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(StockPickingTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindStockPickingTypeId finds record id by querying it with criteria.
func (c *Client) FindStockPickingTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockPickingTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("stock.picking.type was not found with criteria %v and options %v", criteria, options)
}
