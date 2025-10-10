package odoo

// StockInventoryAdjustmentName represents stock.inventory.adjustment.name model.
type StockInventoryAdjustmentName struct {
	CreateDate              *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	InventoryAdjustmentName *String   `xmlrpc:"inventory_adjustment_name,omitempty"`
	QuantIds                *Relation `xmlrpc:"quant_ids,omitempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockInventoryAdjustmentNames represents array of stock.inventory.adjustment.name model.
type StockInventoryAdjustmentNames []StockInventoryAdjustmentName

// StockInventoryAdjustmentNameModel is the odoo model name.
const StockInventoryAdjustmentNameModel = "stock.inventory.adjustment.name"

// Many2One convert StockInventoryAdjustmentName to *Many2One.
func (sian *StockInventoryAdjustmentName) Many2One() *Many2One {
	return NewMany2One(sian.Id.Get(), "")
}

// CreateStockInventoryAdjustmentName creates a new stock.inventory.adjustment.name model and returns its id.
func (c *Client) CreateStockInventoryAdjustmentName(sian *StockInventoryAdjustmentName) (int64, error) {
	ids, err := c.CreateStockInventoryAdjustmentNames([]*StockInventoryAdjustmentName{sian})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockInventoryAdjustmentName creates a new stock.inventory.adjustment.name model and returns its id.
func (c *Client) CreateStockInventoryAdjustmentNames(sians []*StockInventoryAdjustmentName) ([]int64, error) {
	var vv []interface{}
	for _, v := range sians {
		vv = append(vv, v)
	}
	return c.Create(StockInventoryAdjustmentNameModel, vv, nil)
}

// UpdateStockInventoryAdjustmentName updates an existing stock.inventory.adjustment.name record.
func (c *Client) UpdateStockInventoryAdjustmentName(sian *StockInventoryAdjustmentName) error {
	return c.UpdateStockInventoryAdjustmentNames([]int64{sian.Id.Get()}, sian)
}

// UpdateStockInventoryAdjustmentNames updates existing stock.inventory.adjustment.name records.
// All records (represented by ids) will be updated by sian values.
func (c *Client) UpdateStockInventoryAdjustmentNames(ids []int64, sian *StockInventoryAdjustmentName) error {
	return c.Update(StockInventoryAdjustmentNameModel, ids, sian, nil)
}

// DeleteStockInventoryAdjustmentName deletes an existing stock.inventory.adjustment.name record.
func (c *Client) DeleteStockInventoryAdjustmentName(id int64) error {
	return c.DeleteStockInventoryAdjustmentNames([]int64{id})
}

// DeleteStockInventoryAdjustmentNames deletes existing stock.inventory.adjustment.name records.
func (c *Client) DeleteStockInventoryAdjustmentNames(ids []int64) error {
	return c.Delete(StockInventoryAdjustmentNameModel, ids)
}

// GetStockInventoryAdjustmentName gets stock.inventory.adjustment.name existing record.
func (c *Client) GetStockInventoryAdjustmentName(id int64) (*StockInventoryAdjustmentName, error) {
	sians, err := c.GetStockInventoryAdjustmentNames([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sians)[0]), nil
}

// GetStockInventoryAdjustmentNames gets stock.inventory.adjustment.name existing records.
func (c *Client) GetStockInventoryAdjustmentNames(ids []int64) (*StockInventoryAdjustmentNames, error) {
	sians := &StockInventoryAdjustmentNames{}
	if err := c.Read(StockInventoryAdjustmentNameModel, ids, nil, sians); err != nil {
		return nil, err
	}
	return sians, nil
}

// FindStockInventoryAdjustmentName finds stock.inventory.adjustment.name record by querying it with criteria.
func (c *Client) FindStockInventoryAdjustmentName(criteria *Criteria) (*StockInventoryAdjustmentName, error) {
	sians := &StockInventoryAdjustmentNames{}
	if err := c.SearchRead(StockInventoryAdjustmentNameModel, criteria, NewOptions().Limit(1), sians); err != nil {
		return nil, err
	}
	return &((*sians)[0]), nil
}

// FindStockInventoryAdjustmentNames finds stock.inventory.adjustment.name records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryAdjustmentNames(criteria *Criteria, options *Options) (*StockInventoryAdjustmentNames, error) {
	sians := &StockInventoryAdjustmentNames{}
	if err := c.SearchRead(StockInventoryAdjustmentNameModel, criteria, options, sians); err != nil {
		return nil, err
	}
	return sians, nil
}

// FindStockInventoryAdjustmentNameIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockInventoryAdjustmentNameIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockInventoryAdjustmentNameModel, criteria, options)
}

// FindStockInventoryAdjustmentNameId finds record id by querying it with criteria.
func (c *Client) FindStockInventoryAdjustmentNameId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockInventoryAdjustmentNameModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
