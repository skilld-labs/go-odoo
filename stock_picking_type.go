package odoo

import (
	"fmt"
)

// StockPickingType represents stock.picking.type model.
type StockPickingType struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty"`
	Active                 *Bool      `xmlrpc:"active,omitempty"`
	BarcodeNomenclatureId  *Many2One  `xmlrpc:"barcode_nomenclature_id,omitempty"`
	Code                   *Selection `xmlrpc:"code,omitempty"`
	Color                  *Int       `xmlrpc:"color,omitempty"`
	CountPicking           *Int       `xmlrpc:"count_picking,omitempty"`
	CountPickingBackorders *Int       `xmlrpc:"count_picking_backorders,omitempty"`
	CountPickingDraft      *Int       `xmlrpc:"count_picking_draft,omitempty"`
	CountPickingLate       *Int       `xmlrpc:"count_picking_late,omitempty"`
	CountPickingReady      *Int       `xmlrpc:"count_picking_ready,omitempty"`
	CountPickingWaiting    *Int       `xmlrpc:"count_picking_waiting,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	DefaultLocationDestId  *Many2One  `xmlrpc:"default_location_dest_id,omitempty"`
	DefaultLocationSrcId   *Many2One  `xmlrpc:"default_location_src_id,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	LastDonePicking        *String    `xmlrpc:"last_done_picking,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	RatePickingBackorders  *Int       `xmlrpc:"rate_picking_backorders,omitempty"`
	RatePickingLate        *Int       `xmlrpc:"rate_picking_late,omitempty"`
	ReturnPickingTypeId    *Many2One  `xmlrpc:"return_picking_type_id,omitempty"`
	Sequence               *Int       `xmlrpc:"sequence,omitempty"`
	SequenceId             *Many2One  `xmlrpc:"sequence_id,omitempty"`
	ShowEntirePacks        *Bool      `xmlrpc:"show_entire_packs,omitempty"`
	ShowOperations         *Bool      `xmlrpc:"show_operations,omitempty"`
	ShowReserved           *Bool      `xmlrpc:"show_reserved,omitempty"`
	UseCreateLots          *Bool      `xmlrpc:"use_create_lots,omitempty"`
	UseExistingLots        *Bool      `xmlrpc:"use_existing_lots,omitempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
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
	return nil, fmt.Errorf("no stock.picking.type was found with criteria %v", criteria)
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
	return -1, fmt.Errorf("no stock.picking.type was found with criteria %v and options %v", criteria, options)
}
