package odoo

// IrQwebFieldOne2Many represents ir.qweb.field.one2many model.
type IrQwebFieldOne2Many struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldOne2Manys represents array of ir.qweb.field.one2many model.
type IrQwebFieldOne2Manys []IrQwebFieldOne2Many

// IrQwebFieldOne2ManyModel is the odoo model name.
const IrQwebFieldOne2ManyModel = "ir.qweb.field.one2many"

// Many2One convert IrQwebFieldOne2Many to *Many2One.
func (iqfo *IrQwebFieldOne2Many) Many2One() *Many2One {
	return NewMany2One(iqfo.Id.Get(), "")
}

// CreateIrQwebFieldOne2Many creates a new ir.qweb.field.one2many model and returns its id.
func (c *Client) CreateIrQwebFieldOne2Many(iqfo *IrQwebFieldOne2Many) (int64, error) {
	ids, err := c.CreateIrQwebFieldOne2Manys([]*IrQwebFieldOne2Many{iqfo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldOne2Many creates a new ir.qweb.field.one2many model and returns its id.
func (c *Client) CreateIrQwebFieldOne2Manys(iqfos []*IrQwebFieldOne2Many) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfos {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldOne2ManyModel, vv, nil)
}

// UpdateIrQwebFieldOne2Many updates an existing ir.qweb.field.one2many record.
func (c *Client) UpdateIrQwebFieldOne2Many(iqfo *IrQwebFieldOne2Many) error {
	return c.UpdateIrQwebFieldOne2Manys([]int64{iqfo.Id.Get()}, iqfo)
}

// UpdateIrQwebFieldOne2Manys updates existing ir.qweb.field.one2many records.
// All records (represented by ids) will be updated by iqfo values.
func (c *Client) UpdateIrQwebFieldOne2Manys(ids []int64, iqfo *IrQwebFieldOne2Many) error {
	return c.Update(IrQwebFieldOne2ManyModel, ids, iqfo, nil)
}

// DeleteIrQwebFieldOne2Many deletes an existing ir.qweb.field.one2many record.
func (c *Client) DeleteIrQwebFieldOne2Many(id int64) error {
	return c.DeleteIrQwebFieldOne2Manys([]int64{id})
}

// DeleteIrQwebFieldOne2Manys deletes existing ir.qweb.field.one2many records.
func (c *Client) DeleteIrQwebFieldOne2Manys(ids []int64) error {
	return c.Delete(IrQwebFieldOne2ManyModel, ids)
}

// GetIrQwebFieldOne2Many gets ir.qweb.field.one2many existing record.
func (c *Client) GetIrQwebFieldOne2Many(id int64) (*IrQwebFieldOne2Many, error) {
	iqfos, err := c.GetIrQwebFieldOne2Manys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqfos)[0]), nil
}

// GetIrQwebFieldOne2Manys gets ir.qweb.field.one2many existing records.
func (c *Client) GetIrQwebFieldOne2Manys(ids []int64) (*IrQwebFieldOne2Manys, error) {
	iqfos := &IrQwebFieldOne2Manys{}
	if err := c.Read(IrQwebFieldOne2ManyModel, ids, nil, iqfos); err != nil {
		return nil, err
	}
	return iqfos, nil
}

// FindIrQwebFieldOne2Many finds ir.qweb.field.one2many record by querying it with criteria.
func (c *Client) FindIrQwebFieldOne2Many(criteria *Criteria) (*IrQwebFieldOne2Many, error) {
	iqfos := &IrQwebFieldOne2Manys{}
	if err := c.SearchRead(IrQwebFieldOne2ManyModel, criteria, NewOptions().Limit(1), iqfos); err != nil {
		return nil, err
	}
	return &((*iqfos)[0]), nil
}

// FindIrQwebFieldOne2Manys finds ir.qweb.field.one2many records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldOne2Manys(criteria *Criteria, options *Options) (*IrQwebFieldOne2Manys, error) {
	iqfos := &IrQwebFieldOne2Manys{}
	if err := c.SearchRead(IrQwebFieldOne2ManyModel, criteria, options, iqfos); err != nil {
		return nil, err
	}
	return iqfos, nil
}

// FindIrQwebFieldOne2ManyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldOne2ManyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebFieldOne2ManyModel, criteria, options)
}

// FindIrQwebFieldOne2ManyId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldOne2ManyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldOne2ManyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
