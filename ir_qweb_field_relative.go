package odoo

// IrQwebFieldRelative represents ir.qweb.field.relative model.
type IrQwebFieldRelative struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldRelatives represents array of ir.qweb.field.relative model.
type IrQwebFieldRelatives []IrQwebFieldRelative

// IrQwebFieldRelativeModel is the odoo model name.
const IrQwebFieldRelativeModel = "ir.qweb.field.relative"

// Many2One convert IrQwebFieldRelative to *Many2One.
func (iqfr *IrQwebFieldRelative) Many2One() *Many2One {
	return NewMany2One(iqfr.Id.Get(), "")
}

// CreateIrQwebFieldRelative creates a new ir.qweb.field.relative model and returns its id.
func (c *Client) CreateIrQwebFieldRelative(iqfr *IrQwebFieldRelative) (int64, error) {
	ids, err := c.CreateIrQwebFieldRelatives([]*IrQwebFieldRelative{iqfr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldRelative creates a new ir.qweb.field.relative model and returns its id.
func (c *Client) CreateIrQwebFieldRelatives(iqfrs []*IrQwebFieldRelative) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfrs {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldRelativeModel, vv, nil)
}

// UpdateIrQwebFieldRelative updates an existing ir.qweb.field.relative record.
func (c *Client) UpdateIrQwebFieldRelative(iqfr *IrQwebFieldRelative) error {
	return c.UpdateIrQwebFieldRelatives([]int64{iqfr.Id.Get()}, iqfr)
}

// UpdateIrQwebFieldRelatives updates existing ir.qweb.field.relative records.
// All records (represented by ids) will be updated by iqfr values.
func (c *Client) UpdateIrQwebFieldRelatives(ids []int64, iqfr *IrQwebFieldRelative) error {
	return c.Update(IrQwebFieldRelativeModel, ids, iqfr, nil)
}

// DeleteIrQwebFieldRelative deletes an existing ir.qweb.field.relative record.
func (c *Client) DeleteIrQwebFieldRelative(id int64) error {
	return c.DeleteIrQwebFieldRelatives([]int64{id})
}

// DeleteIrQwebFieldRelatives deletes existing ir.qweb.field.relative records.
func (c *Client) DeleteIrQwebFieldRelatives(ids []int64) error {
	return c.Delete(IrQwebFieldRelativeModel, ids)
}

// GetIrQwebFieldRelative gets ir.qweb.field.relative existing record.
func (c *Client) GetIrQwebFieldRelative(id int64) (*IrQwebFieldRelative, error) {
	iqfrs, err := c.GetIrQwebFieldRelatives([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqfrs)[0]), nil
}

// GetIrQwebFieldRelatives gets ir.qweb.field.relative existing records.
func (c *Client) GetIrQwebFieldRelatives(ids []int64) (*IrQwebFieldRelatives, error) {
	iqfrs := &IrQwebFieldRelatives{}
	if err := c.Read(IrQwebFieldRelativeModel, ids, nil, iqfrs); err != nil {
		return nil, err
	}
	return iqfrs, nil
}

// FindIrQwebFieldRelative finds ir.qweb.field.relative record by querying it with criteria.
func (c *Client) FindIrQwebFieldRelative(criteria *Criteria) (*IrQwebFieldRelative, error) {
	iqfrs := &IrQwebFieldRelatives{}
	if err := c.SearchRead(IrQwebFieldRelativeModel, criteria, NewOptions().Limit(1), iqfrs); err != nil {
		return nil, err
	}
	return &((*iqfrs)[0]), nil
}

// FindIrQwebFieldRelatives finds ir.qweb.field.relative records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldRelatives(criteria *Criteria, options *Options) (*IrQwebFieldRelatives, error) {
	iqfrs := &IrQwebFieldRelatives{}
	if err := c.SearchRead(IrQwebFieldRelativeModel, criteria, options, iqfrs); err != nil {
		return nil, err
	}
	return iqfrs, nil
}

// FindIrQwebFieldRelativeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldRelativeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebFieldRelativeModel, criteria, options)
}

// FindIrQwebFieldRelativeId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldRelativeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldRelativeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
