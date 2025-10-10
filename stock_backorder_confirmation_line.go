package odoo

// StockBackorderConfirmationLine represents stock.backorder.confirmation.line model.
type StockBackorderConfirmationLine struct {
	BackorderConfirmationId *Many2One `xmlrpc:"backorder_confirmation_id,omitempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	PickingId               *Many2One `xmlrpc:"picking_id,omitempty"`
	ToBackorder             *Bool     `xmlrpc:"to_backorder,omitempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockBackorderConfirmationLines represents array of stock.backorder.confirmation.line model.
type StockBackorderConfirmationLines []StockBackorderConfirmationLine

// StockBackorderConfirmationLineModel is the odoo model name.
const StockBackorderConfirmationLineModel = "stock.backorder.confirmation.line"

// Many2One convert StockBackorderConfirmationLine to *Many2One.
func (sbcl *StockBackorderConfirmationLine) Many2One() *Many2One {
	return NewMany2One(sbcl.Id.Get(), "")
}

// CreateStockBackorderConfirmationLine creates a new stock.backorder.confirmation.line model and returns its id.
func (c *Client) CreateStockBackorderConfirmationLine(sbcl *StockBackorderConfirmationLine) (int64, error) {
	ids, err := c.CreateStockBackorderConfirmationLines([]*StockBackorderConfirmationLine{sbcl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockBackorderConfirmationLine creates a new stock.backorder.confirmation.line model and returns its id.
func (c *Client) CreateStockBackorderConfirmationLines(sbcls []*StockBackorderConfirmationLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range sbcls {
		vv = append(vv, v)
	}
	return c.Create(StockBackorderConfirmationLineModel, vv, nil)
}

// UpdateStockBackorderConfirmationLine updates an existing stock.backorder.confirmation.line record.
func (c *Client) UpdateStockBackorderConfirmationLine(sbcl *StockBackorderConfirmationLine) error {
	return c.UpdateStockBackorderConfirmationLines([]int64{sbcl.Id.Get()}, sbcl)
}

// UpdateStockBackorderConfirmationLines updates existing stock.backorder.confirmation.line records.
// All records (represented by ids) will be updated by sbcl values.
func (c *Client) UpdateStockBackorderConfirmationLines(ids []int64, sbcl *StockBackorderConfirmationLine) error {
	return c.Update(StockBackorderConfirmationLineModel, ids, sbcl, nil)
}

// DeleteStockBackorderConfirmationLine deletes an existing stock.backorder.confirmation.line record.
func (c *Client) DeleteStockBackorderConfirmationLine(id int64) error {
	return c.DeleteStockBackorderConfirmationLines([]int64{id})
}

// DeleteStockBackorderConfirmationLines deletes existing stock.backorder.confirmation.line records.
func (c *Client) DeleteStockBackorderConfirmationLines(ids []int64) error {
	return c.Delete(StockBackorderConfirmationLineModel, ids)
}

// GetStockBackorderConfirmationLine gets stock.backorder.confirmation.line existing record.
func (c *Client) GetStockBackorderConfirmationLine(id int64) (*StockBackorderConfirmationLine, error) {
	sbcls, err := c.GetStockBackorderConfirmationLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sbcls)[0]), nil
}

// GetStockBackorderConfirmationLines gets stock.backorder.confirmation.line existing records.
func (c *Client) GetStockBackorderConfirmationLines(ids []int64) (*StockBackorderConfirmationLines, error) {
	sbcls := &StockBackorderConfirmationLines{}
	if err := c.Read(StockBackorderConfirmationLineModel, ids, nil, sbcls); err != nil {
		return nil, err
	}
	return sbcls, nil
}

// FindStockBackorderConfirmationLine finds stock.backorder.confirmation.line record by querying it with criteria.
func (c *Client) FindStockBackorderConfirmationLine(criteria *Criteria) (*StockBackorderConfirmationLine, error) {
	sbcls := &StockBackorderConfirmationLines{}
	if err := c.SearchRead(StockBackorderConfirmationLineModel, criteria, NewOptions().Limit(1), sbcls); err != nil {
		return nil, err
	}
	return &((*sbcls)[0]), nil
}

// FindStockBackorderConfirmationLines finds stock.backorder.confirmation.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBackorderConfirmationLines(criteria *Criteria, options *Options) (*StockBackorderConfirmationLines, error) {
	sbcls := &StockBackorderConfirmationLines{}
	if err := c.SearchRead(StockBackorderConfirmationLineModel, criteria, options, sbcls); err != nil {
		return nil, err
	}
	return sbcls, nil
}

// FindStockBackorderConfirmationLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBackorderConfirmationLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockBackorderConfirmationLineModel, criteria, options)
}

// FindStockBackorderConfirmationLineId finds record id by querying it with criteria.
func (c *Client) FindStockBackorderConfirmationLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockBackorderConfirmationLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
