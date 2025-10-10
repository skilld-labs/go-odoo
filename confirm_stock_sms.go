package odoo

// ConfirmStockSms represents confirm.stock.sms model.
type ConfirmStockSms struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	PickIds     *Relation `xmlrpc:"pick_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ConfirmStockSmss represents array of confirm.stock.sms model.
type ConfirmStockSmss []ConfirmStockSms

// ConfirmStockSmsModel is the odoo model name.
const ConfirmStockSmsModel = "confirm.stock.sms"

// Many2One convert ConfirmStockSms to *Many2One.
func (css *ConfirmStockSms) Many2One() *Many2One {
	return NewMany2One(css.Id.Get(), "")
}

// CreateConfirmStockSms creates a new confirm.stock.sms model and returns its id.
func (c *Client) CreateConfirmStockSms(css *ConfirmStockSms) (int64, error) {
	ids, err := c.CreateConfirmStockSmss([]*ConfirmStockSms{css})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateConfirmStockSms creates a new confirm.stock.sms model and returns its id.
func (c *Client) CreateConfirmStockSmss(csss []*ConfirmStockSms) ([]int64, error) {
	var vv []interface{}
	for _, v := range csss {
		vv = append(vv, v)
	}
	return c.Create(ConfirmStockSmsModel, vv, nil)
}

// UpdateConfirmStockSms updates an existing confirm.stock.sms record.
func (c *Client) UpdateConfirmStockSms(css *ConfirmStockSms) error {
	return c.UpdateConfirmStockSmss([]int64{css.Id.Get()}, css)
}

// UpdateConfirmStockSmss updates existing confirm.stock.sms records.
// All records (represented by ids) will be updated by css values.
func (c *Client) UpdateConfirmStockSmss(ids []int64, css *ConfirmStockSms) error {
	return c.Update(ConfirmStockSmsModel, ids, css, nil)
}

// DeleteConfirmStockSms deletes an existing confirm.stock.sms record.
func (c *Client) DeleteConfirmStockSms(id int64) error {
	return c.DeleteConfirmStockSmss([]int64{id})
}

// DeleteConfirmStockSmss deletes existing confirm.stock.sms records.
func (c *Client) DeleteConfirmStockSmss(ids []int64) error {
	return c.Delete(ConfirmStockSmsModel, ids)
}

// GetConfirmStockSms gets confirm.stock.sms existing record.
func (c *Client) GetConfirmStockSms(id int64) (*ConfirmStockSms, error) {
	csss, err := c.GetConfirmStockSmss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*csss)[0]), nil
}

// GetConfirmStockSmss gets confirm.stock.sms existing records.
func (c *Client) GetConfirmStockSmss(ids []int64) (*ConfirmStockSmss, error) {
	csss := &ConfirmStockSmss{}
	if err := c.Read(ConfirmStockSmsModel, ids, nil, csss); err != nil {
		return nil, err
	}
	return csss, nil
}

// FindConfirmStockSms finds confirm.stock.sms record by querying it with criteria.
func (c *Client) FindConfirmStockSms(criteria *Criteria) (*ConfirmStockSms, error) {
	csss := &ConfirmStockSmss{}
	if err := c.SearchRead(ConfirmStockSmsModel, criteria, NewOptions().Limit(1), csss); err != nil {
		return nil, err
	}
	return &((*csss)[0]), nil
}

// FindConfirmStockSmss finds confirm.stock.sms records by querying it
// and filtering it with criteria and options.
func (c *Client) FindConfirmStockSmss(criteria *Criteria, options *Options) (*ConfirmStockSmss, error) {
	csss := &ConfirmStockSmss{}
	if err := c.SearchRead(ConfirmStockSmsModel, criteria, options, csss); err != nil {
		return nil, err
	}
	return csss, nil
}

// FindConfirmStockSmsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindConfirmStockSmsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ConfirmStockSmsModel, criteria, options)
}

// FindConfirmStockSmsId finds record id by querying it with criteria.
func (c *Client) FindConfirmStockSmsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ConfirmStockSmsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
