package odoo

// StockBackorderConfirmation represents stock.backorder.confirmation model.
type StockBackorderConfirmation struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PickIds     *Relation `xmlrpc:"pick_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockBackorderConfirmations represents array of stock.backorder.confirmation model.
type StockBackorderConfirmations []StockBackorderConfirmation

// StockBackorderConfirmationModel is the odoo model name.
const StockBackorderConfirmationModel = "stock.backorder.confirmation"

// Many2One convert StockBackorderConfirmation to *Many2One.
func (sbc *StockBackorderConfirmation) Many2One() *Many2One {
	return NewMany2One(sbc.Id.Get(), "")
}

// CreateStockBackorderConfirmation creates a new stock.backorder.confirmation model and returns its id.
func (c *Client) CreateStockBackorderConfirmation(sbc *StockBackorderConfirmation) (int64, error) {
	ids, err := c.CreateStockBackorderConfirmations([]*StockBackorderConfirmation{sbc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockBackorderConfirmation creates a new stock.backorder.confirmation model and returns its id.
func (c *Client) CreateStockBackorderConfirmations(sbcs []*StockBackorderConfirmation) ([]int64, error) {
	var vv []interface{}
	for _, v := range sbcs {
		vv = append(vv, v)
	}
	return c.Create(StockBackorderConfirmationModel, vv, nil)
}

// UpdateStockBackorderConfirmation updates an existing stock.backorder.confirmation record.
func (c *Client) UpdateStockBackorderConfirmation(sbc *StockBackorderConfirmation) error {
	return c.UpdateStockBackorderConfirmations([]int64{sbc.Id.Get()}, sbc)
}

// UpdateStockBackorderConfirmations updates existing stock.backorder.confirmation records.
// All records (represented by ids) will be updated by sbc values.
func (c *Client) UpdateStockBackorderConfirmations(ids []int64, sbc *StockBackorderConfirmation) error {
	return c.Update(StockBackorderConfirmationModel, ids, sbc, nil)
}

// DeleteStockBackorderConfirmation deletes an existing stock.backorder.confirmation record.
func (c *Client) DeleteStockBackorderConfirmation(id int64) error {
	return c.DeleteStockBackorderConfirmations([]int64{id})
}

// DeleteStockBackorderConfirmations deletes existing stock.backorder.confirmation records.
func (c *Client) DeleteStockBackorderConfirmations(ids []int64) error {
	return c.Delete(StockBackorderConfirmationModel, ids)
}

// GetStockBackorderConfirmation gets stock.backorder.confirmation existing record.
func (c *Client) GetStockBackorderConfirmation(id int64) (*StockBackorderConfirmation, error) {
	sbcs, err := c.GetStockBackorderConfirmations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sbcs)[0]), nil
}

// GetStockBackorderConfirmations gets stock.backorder.confirmation existing records.
func (c *Client) GetStockBackorderConfirmations(ids []int64) (*StockBackorderConfirmations, error) {
	sbcs := &StockBackorderConfirmations{}
	if err := c.Read(StockBackorderConfirmationModel, ids, nil, sbcs); err != nil {
		return nil, err
	}
	return sbcs, nil
}

// FindStockBackorderConfirmation finds stock.backorder.confirmation record by querying it with criteria.
func (c *Client) FindStockBackorderConfirmation(criteria *Criteria) (*StockBackorderConfirmation, error) {
	sbcs := &StockBackorderConfirmations{}
	if err := c.SearchRead(StockBackorderConfirmationModel, criteria, NewOptions().Limit(1), sbcs); err != nil {
		return nil, err
	}
	return &((*sbcs)[0]), nil
}

// FindStockBackorderConfirmations finds stock.backorder.confirmation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBackorderConfirmations(criteria *Criteria, options *Options) (*StockBackorderConfirmations, error) {
	sbcs := &StockBackorderConfirmations{}
	if err := c.SearchRead(StockBackorderConfirmationModel, criteria, options, sbcs); err != nil {
		return nil, err
	}
	return sbcs, nil
}

// FindStockBackorderConfirmationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockBackorderConfirmationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockBackorderConfirmationModel, criteria, options)
}

// FindStockBackorderConfirmationId finds record id by querying it with criteria.
func (c *Client) FindStockBackorderConfirmationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockBackorderConfirmationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
