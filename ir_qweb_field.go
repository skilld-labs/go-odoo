package odoo

// IrQwebField represents ir.qweb.field model.
type IrQwebField struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFields represents array of ir.qweb.field model.
type IrQwebFields []IrQwebField

// IrQwebFieldModel is the odoo model name.
const IrQwebFieldModel = "ir.qweb.field"

// Many2One convert IrQwebField to *Many2One.
func (iqf *IrQwebField) Many2One() *Many2One {
	return NewMany2One(iqf.Id.Get(), "")
}

// CreateIrQwebField creates a new ir.qweb.field model and returns its id.
func (c *Client) CreateIrQwebField(iqf *IrQwebField) (int64, error) {
	ids, err := c.CreateIrQwebFields([]*IrQwebField{iqf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebField creates a new ir.qweb.field model and returns its id.
func (c *Client) CreateIrQwebFields(iqfs []*IrQwebField) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfs {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldModel, vv, nil)
}

// UpdateIrQwebField updates an existing ir.qweb.field record.
func (c *Client) UpdateIrQwebField(iqf *IrQwebField) error {
	return c.UpdateIrQwebFields([]int64{iqf.Id.Get()}, iqf)
}

// UpdateIrQwebFields updates existing ir.qweb.field records.
// All records (represented by ids) will be updated by iqf values.
func (c *Client) UpdateIrQwebFields(ids []int64, iqf *IrQwebField) error {
	return c.Update(IrQwebFieldModel, ids, iqf, nil)
}

// DeleteIrQwebField deletes an existing ir.qweb.field record.
func (c *Client) DeleteIrQwebField(id int64) error {
	return c.DeleteIrQwebFields([]int64{id})
}

// DeleteIrQwebFields deletes existing ir.qweb.field records.
func (c *Client) DeleteIrQwebFields(ids []int64) error {
	return c.Delete(IrQwebFieldModel, ids)
}

// GetIrQwebField gets ir.qweb.field existing record.
func (c *Client) GetIrQwebField(id int64) (*IrQwebField, error) {
	iqfs, err := c.GetIrQwebFields([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqfs)[0]), nil
}

// GetIrQwebFields gets ir.qweb.field existing records.
func (c *Client) GetIrQwebFields(ids []int64) (*IrQwebFields, error) {
	iqfs := &IrQwebFields{}
	if err := c.Read(IrQwebFieldModel, ids, nil, iqfs); err != nil {
		return nil, err
	}
	return iqfs, nil
}

// FindIrQwebField finds ir.qweb.field record by querying it with criteria.
func (c *Client) FindIrQwebField(criteria *Criteria) (*IrQwebField, error) {
	iqfs := &IrQwebFields{}
	if err := c.SearchRead(IrQwebFieldModel, criteria, NewOptions().Limit(1), iqfs); err != nil {
		return nil, err
	}
	return &((*iqfs)[0]), nil
}

// FindIrQwebFields finds ir.qweb.field records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFields(criteria *Criteria, options *Options) (*IrQwebFields, error) {
	iqfs := &IrQwebFields{}
	if err := c.SearchRead(IrQwebFieldModel, criteria, options, iqfs); err != nil {
		return nil, err
	}
	return iqfs, nil
}

// FindIrQwebFieldIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebFieldModel, criteria, options)
}

// FindIrQwebFieldId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
