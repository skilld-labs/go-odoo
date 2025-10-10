package odoo

// StockTrackConfirmation represents stock.track.confirmation model.
type StockTrackConfirmation struct {
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	ProductIds      *Relation `xmlrpc:"product_ids,omitempty"`
	QuantIds        *Relation `xmlrpc:"quant_ids,omitempty"`
	TrackingLineIds *Relation `xmlrpc:"tracking_line_ids,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// StockTrackConfirmations represents array of stock.track.confirmation model.
type StockTrackConfirmations []StockTrackConfirmation

// StockTrackConfirmationModel is the odoo model name.
const StockTrackConfirmationModel = "stock.track.confirmation"

// Many2One convert StockTrackConfirmation to *Many2One.
func (stc *StockTrackConfirmation) Many2One() *Many2One {
	return NewMany2One(stc.Id.Get(), "")
}

// CreateStockTrackConfirmation creates a new stock.track.confirmation model and returns its id.
func (c *Client) CreateStockTrackConfirmation(stc *StockTrackConfirmation) (int64, error) {
	ids, err := c.CreateStockTrackConfirmations([]*StockTrackConfirmation{stc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockTrackConfirmation creates a new stock.track.confirmation model and returns its id.
func (c *Client) CreateStockTrackConfirmations(stcs []*StockTrackConfirmation) ([]int64, error) {
	var vv []interface{}
	for _, v := range stcs {
		vv = append(vv, v)
	}
	return c.Create(StockTrackConfirmationModel, vv, nil)
}

// UpdateStockTrackConfirmation updates an existing stock.track.confirmation record.
func (c *Client) UpdateStockTrackConfirmation(stc *StockTrackConfirmation) error {
	return c.UpdateStockTrackConfirmations([]int64{stc.Id.Get()}, stc)
}

// UpdateStockTrackConfirmations updates existing stock.track.confirmation records.
// All records (represented by ids) will be updated by stc values.
func (c *Client) UpdateStockTrackConfirmations(ids []int64, stc *StockTrackConfirmation) error {
	return c.Update(StockTrackConfirmationModel, ids, stc, nil)
}

// DeleteStockTrackConfirmation deletes an existing stock.track.confirmation record.
func (c *Client) DeleteStockTrackConfirmation(id int64) error {
	return c.DeleteStockTrackConfirmations([]int64{id})
}

// DeleteStockTrackConfirmations deletes existing stock.track.confirmation records.
func (c *Client) DeleteStockTrackConfirmations(ids []int64) error {
	return c.Delete(StockTrackConfirmationModel, ids)
}

// GetStockTrackConfirmation gets stock.track.confirmation existing record.
func (c *Client) GetStockTrackConfirmation(id int64) (*StockTrackConfirmation, error) {
	stcs, err := c.GetStockTrackConfirmations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*stcs)[0]), nil
}

// GetStockTrackConfirmations gets stock.track.confirmation existing records.
func (c *Client) GetStockTrackConfirmations(ids []int64) (*StockTrackConfirmations, error) {
	stcs := &StockTrackConfirmations{}
	if err := c.Read(StockTrackConfirmationModel, ids, nil, stcs); err != nil {
		return nil, err
	}
	return stcs, nil
}

// FindStockTrackConfirmation finds stock.track.confirmation record by querying it with criteria.
func (c *Client) FindStockTrackConfirmation(criteria *Criteria) (*StockTrackConfirmation, error) {
	stcs := &StockTrackConfirmations{}
	if err := c.SearchRead(StockTrackConfirmationModel, criteria, NewOptions().Limit(1), stcs); err != nil {
		return nil, err
	}
	return &((*stcs)[0]), nil
}

// FindStockTrackConfirmations finds stock.track.confirmation records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockTrackConfirmations(criteria *Criteria, options *Options) (*StockTrackConfirmations, error) {
	stcs := &StockTrackConfirmations{}
	if err := c.SearchRead(StockTrackConfirmationModel, criteria, options, stcs); err != nil {
		return nil, err
	}
	return stcs, nil
}

// FindStockTrackConfirmationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockTrackConfirmationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockTrackConfirmationModel, criteria, options)
}

// FindStockTrackConfirmationId finds record id by querying it with criteria.
func (c *Client) FindStockTrackConfirmationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockTrackConfirmationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
