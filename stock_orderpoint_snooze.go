package odoo

// StockOrderpointSnooze represents stock.orderpoint.snooze model.
type StockOrderpointSnooze struct {
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	OrderpointIds  *Relation  `xmlrpc:"orderpoint_ids,omitempty"`
	PredefinedDate *Selection `xmlrpc:"predefined_date,omitempty"`
	SnoozedUntil   *Time      `xmlrpc:"snoozed_until,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockOrderpointSnoozes represents array of stock.orderpoint.snooze model.
type StockOrderpointSnoozes []StockOrderpointSnooze

// StockOrderpointSnoozeModel is the odoo model name.
const StockOrderpointSnoozeModel = "stock.orderpoint.snooze"

// Many2One convert StockOrderpointSnooze to *Many2One.
func (sos *StockOrderpointSnooze) Many2One() *Many2One {
	return NewMany2One(sos.Id.Get(), "")
}

// CreateStockOrderpointSnooze creates a new stock.orderpoint.snooze model and returns its id.
func (c *Client) CreateStockOrderpointSnooze(sos *StockOrderpointSnooze) (int64, error) {
	ids, err := c.CreateStockOrderpointSnoozes([]*StockOrderpointSnooze{sos})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockOrderpointSnooze creates a new stock.orderpoint.snooze model and returns its id.
func (c *Client) CreateStockOrderpointSnoozes(soss []*StockOrderpointSnooze) ([]int64, error) {
	var vv []interface{}
	for _, v := range soss {
		vv = append(vv, v)
	}
	return c.Create(StockOrderpointSnoozeModel, vv, nil)
}

// UpdateStockOrderpointSnooze updates an existing stock.orderpoint.snooze record.
func (c *Client) UpdateStockOrderpointSnooze(sos *StockOrderpointSnooze) error {
	return c.UpdateStockOrderpointSnoozes([]int64{sos.Id.Get()}, sos)
}

// UpdateStockOrderpointSnoozes updates existing stock.orderpoint.snooze records.
// All records (represented by ids) will be updated by sos values.
func (c *Client) UpdateStockOrderpointSnoozes(ids []int64, sos *StockOrderpointSnooze) error {
	return c.Update(StockOrderpointSnoozeModel, ids, sos, nil)
}

// DeleteStockOrderpointSnooze deletes an existing stock.orderpoint.snooze record.
func (c *Client) DeleteStockOrderpointSnooze(id int64) error {
	return c.DeleteStockOrderpointSnoozes([]int64{id})
}

// DeleteStockOrderpointSnoozes deletes existing stock.orderpoint.snooze records.
func (c *Client) DeleteStockOrderpointSnoozes(ids []int64) error {
	return c.Delete(StockOrderpointSnoozeModel, ids)
}

// GetStockOrderpointSnooze gets stock.orderpoint.snooze existing record.
func (c *Client) GetStockOrderpointSnooze(id int64) (*StockOrderpointSnooze, error) {
	soss, err := c.GetStockOrderpointSnoozes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*soss)[0]), nil
}

// GetStockOrderpointSnoozes gets stock.orderpoint.snooze existing records.
func (c *Client) GetStockOrderpointSnoozes(ids []int64) (*StockOrderpointSnoozes, error) {
	soss := &StockOrderpointSnoozes{}
	if err := c.Read(StockOrderpointSnoozeModel, ids, nil, soss); err != nil {
		return nil, err
	}
	return soss, nil
}

// FindStockOrderpointSnooze finds stock.orderpoint.snooze record by querying it with criteria.
func (c *Client) FindStockOrderpointSnooze(criteria *Criteria) (*StockOrderpointSnooze, error) {
	soss := &StockOrderpointSnoozes{}
	if err := c.SearchRead(StockOrderpointSnoozeModel, criteria, NewOptions().Limit(1), soss); err != nil {
		return nil, err
	}
	return &((*soss)[0]), nil
}

// FindStockOrderpointSnoozes finds stock.orderpoint.snooze records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockOrderpointSnoozes(criteria *Criteria, options *Options) (*StockOrderpointSnoozes, error) {
	soss := &StockOrderpointSnoozes{}
	if err := c.SearchRead(StockOrderpointSnoozeModel, criteria, options, soss); err != nil {
		return nil, err
	}
	return soss, nil
}

// FindStockOrderpointSnoozeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockOrderpointSnoozeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockOrderpointSnoozeModel, criteria, options)
}

// FindStockOrderpointSnoozeId finds record id by querying it with criteria.
func (c *Client) FindStockOrderpointSnoozeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockOrderpointSnoozeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
