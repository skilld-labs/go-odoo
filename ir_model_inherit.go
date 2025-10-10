package odoo

// IrModelInherit represents ir.model.inherit model.
type IrModelInherit struct {
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	ModelId       *Many2One `xmlrpc:"model_id,omitempty"`
	ParentFieldId *Many2One `xmlrpc:"parent_field_id,omitempty"`
	ParentId      *Many2One `xmlrpc:"parent_id,omitempty"`
}

// IrModelInherits represents array of ir.model.inherit model.
type IrModelInherits []IrModelInherit

// IrModelInheritModel is the odoo model name.
const IrModelInheritModel = "ir.model.inherit"

// Many2One convert IrModelInherit to *Many2One.
func (imi *IrModelInherit) Many2One() *Many2One {
	return NewMany2One(imi.Id.Get(), "")
}

// CreateIrModelInherit creates a new ir.model.inherit model and returns its id.
func (c *Client) CreateIrModelInherit(imi *IrModelInherit) (int64, error) {
	ids, err := c.CreateIrModelInherits([]*IrModelInherit{imi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModelInherit creates a new ir.model.inherit model and returns its id.
func (c *Client) CreateIrModelInherits(imis []*IrModelInherit) ([]int64, error) {
	var vv []interface{}
	for _, v := range imis {
		vv = append(vv, v)
	}
	return c.Create(IrModelInheritModel, vv, nil)
}

// UpdateIrModelInherit updates an existing ir.model.inherit record.
func (c *Client) UpdateIrModelInherit(imi *IrModelInherit) error {
	return c.UpdateIrModelInherits([]int64{imi.Id.Get()}, imi)
}

// UpdateIrModelInherits updates existing ir.model.inherit records.
// All records (represented by ids) will be updated by imi values.
func (c *Client) UpdateIrModelInherits(ids []int64, imi *IrModelInherit) error {
	return c.Update(IrModelInheritModel, ids, imi, nil)
}

// DeleteIrModelInherit deletes an existing ir.model.inherit record.
func (c *Client) DeleteIrModelInherit(id int64) error {
	return c.DeleteIrModelInherits([]int64{id})
}

// DeleteIrModelInherits deletes existing ir.model.inherit records.
func (c *Client) DeleteIrModelInherits(ids []int64) error {
	return c.Delete(IrModelInheritModel, ids)
}

// GetIrModelInherit gets ir.model.inherit existing record.
func (c *Client) GetIrModelInherit(id int64) (*IrModelInherit, error) {
	imis, err := c.GetIrModelInherits([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imis)[0]), nil
}

// GetIrModelInherits gets ir.model.inherit existing records.
func (c *Client) GetIrModelInherits(ids []int64) (*IrModelInherits, error) {
	imis := &IrModelInherits{}
	if err := c.Read(IrModelInheritModel, ids, nil, imis); err != nil {
		return nil, err
	}
	return imis, nil
}

// FindIrModelInherit finds ir.model.inherit record by querying it with criteria.
func (c *Client) FindIrModelInherit(criteria *Criteria) (*IrModelInherit, error) {
	imis := &IrModelInherits{}
	if err := c.SearchRead(IrModelInheritModel, criteria, NewOptions().Limit(1), imis); err != nil {
		return nil, err
	}
	return &((*imis)[0]), nil
}

// FindIrModelInherits finds ir.model.inherit records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelInherits(criteria *Criteria, options *Options) (*IrModelInherits, error) {
	imis := &IrModelInherits{}
	if err := c.SearchRead(IrModelInheritModel, criteria, options, imis); err != nil {
		return nil, err
	}
	return imis, nil
}

// FindIrModelInheritIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelInheritIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrModelInheritModel, criteria, options)
}

// FindIrModelInheritId finds record id by querying it with criteria.
func (c *Client) FindIrModelInheritId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelInheritModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
