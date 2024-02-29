package odoo

// IrQwebFieldDatetime represents ir.qweb.field.datetime model.
type IrQwebFieldDatetime struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldDatetimes represents array of ir.qweb.field.datetime model.
type IrQwebFieldDatetimes []IrQwebFieldDatetime

// IrQwebFieldDatetimeModel is the odoo model name.
const IrQwebFieldDatetimeModel = "ir.qweb.field.datetime"

// Many2One convert IrQwebFieldDatetime to *Many2One.
func (iqfd *IrQwebFieldDatetime) Many2One() *Many2One {
	return NewMany2One(iqfd.Id.Get(), "")
}

// CreateIrQwebFieldDatetime creates a new ir.qweb.field.datetime model and returns its id.
func (c *Client) CreateIrQwebFieldDatetime(iqfd *IrQwebFieldDatetime) (int64, error) {
	ids, err := c.CreateIrQwebFieldDatetimes([]*IrQwebFieldDatetime{iqfd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldDatetimes creates a new ir.qweb.field.datetime model and returns its id.
func (c *Client) CreateIrQwebFieldDatetimes(iqfds []*IrQwebFieldDatetime) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfds {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldDatetimeModel, vv, nil)
}

// UpdateIrQwebFieldDatetime updates an existing ir.qweb.field.datetime record.
func (c *Client) UpdateIrQwebFieldDatetime(iqfd *IrQwebFieldDatetime) error {
	return c.UpdateIrQwebFieldDatetimes([]int64{iqfd.Id.Get()}, iqfd)
}

// UpdateIrQwebFieldDatetimes updates existing ir.qweb.field.datetime records.
// All records (represented by ids) will be updated by iqfd values.
func (c *Client) UpdateIrQwebFieldDatetimes(ids []int64, iqfd *IrQwebFieldDatetime) error {
	return c.Update(IrQwebFieldDatetimeModel, ids, iqfd, nil)
}

// DeleteIrQwebFieldDatetime deletes an existing ir.qweb.field.datetime record.
func (c *Client) DeleteIrQwebFieldDatetime(id int64) error {
	return c.DeleteIrQwebFieldDatetimes([]int64{id})
}

// DeleteIrQwebFieldDatetimes deletes existing ir.qweb.field.datetime records.
func (c *Client) DeleteIrQwebFieldDatetimes(ids []int64) error {
	return c.Delete(IrQwebFieldDatetimeModel, ids)
}

// GetIrQwebFieldDatetime gets ir.qweb.field.datetime existing record.
func (c *Client) GetIrQwebFieldDatetime(id int64) (*IrQwebFieldDatetime, error) {
	iqfds, err := c.GetIrQwebFieldDatetimes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqfds)[0]), nil
}

// GetIrQwebFieldDatetimes gets ir.qweb.field.datetime existing records.
func (c *Client) GetIrQwebFieldDatetimes(ids []int64) (*IrQwebFieldDatetimes, error) {
	iqfds := &IrQwebFieldDatetimes{}
	if err := c.Read(IrQwebFieldDatetimeModel, ids, nil, iqfds); err != nil {
		return nil, err
	}
	return iqfds, nil
}

// FindIrQwebFieldDatetime finds ir.qweb.field.datetime record by querying it with criteria.
func (c *Client) FindIrQwebFieldDatetime(criteria *Criteria) (*IrQwebFieldDatetime, error) {
	iqfds := &IrQwebFieldDatetimes{}
	if err := c.SearchRead(IrQwebFieldDatetimeModel, criteria, NewOptions().Limit(1), iqfds); err != nil {
		return nil, err
	}
	return &((*iqfds)[0]), nil
}

// FindIrQwebFieldDatetimes finds ir.qweb.field.datetime records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldDatetimes(criteria *Criteria, options *Options) (*IrQwebFieldDatetimes, error) {
	iqfds := &IrQwebFieldDatetimes{}
	if err := c.SearchRead(IrQwebFieldDatetimeModel, criteria, options, iqfds); err != nil {
		return nil, err
	}
	return iqfds, nil
}

// FindIrQwebFieldDatetimeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldDatetimeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebFieldDatetimeModel, criteria, options)
}

// FindIrQwebFieldDatetimeId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldDatetimeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldDatetimeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
