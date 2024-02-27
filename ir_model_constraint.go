package odoo

// IrModelConstraint represents ir.model.constraint model.
type IrModelConstraint struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DateInit    *Time     `xmlrpc:"date_init,omitempty"`
	DateUpdate  *Time     `xmlrpc:"date_update,omitempty"`
	Definition  *String   `xmlrpc:"definition,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Model       *Many2One `xmlrpc:"model,omitempty"`
	Module      *Many2One `xmlrpc:"module,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Type        *String   `xmlrpc:"type,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrModelConstraints represents array of ir.model.constraint model.
type IrModelConstraints []IrModelConstraint

// IrModelConstraintModel is the odoo model name.
const IrModelConstraintModel = "ir.model.constraint"

// Many2One convert IrModelConstraint to *Many2One.
func (imc *IrModelConstraint) Many2One() *Many2One {
	return NewMany2One(imc.Id.Get(), "")
}

// CreateIrModelConstraint creates a new ir.model.constraint model and returns its id.
func (c *Client) CreateIrModelConstraint(imc *IrModelConstraint) (int64, error) {
	ids, err := c.CreateIrModelConstraints([]*IrModelConstraint{imc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModelConstraints creates a new ir.model.constraint model and returns its id.
func (c *Client) CreateIrModelConstraints(imcs []*IrModelConstraint) ([]int64, error) {
	var vv []interface{}
	for _, v := range imcs {
		vv = append(vv, v)
	}
	return c.Create(IrModelConstraintModel, vv, nil)
}

// UpdateIrModelConstraint updates an existing ir.model.constraint record.
func (c *Client) UpdateIrModelConstraint(imc *IrModelConstraint) error {
	return c.UpdateIrModelConstraints([]int64{imc.Id.Get()}, imc)
}

// UpdateIrModelConstraints updates existing ir.model.constraint records.
// All records (represented by ids) will be updated by imc values.
func (c *Client) UpdateIrModelConstraints(ids []int64, imc *IrModelConstraint) error {
	return c.Update(IrModelConstraintModel, ids, imc, nil)
}

// DeleteIrModelConstraint deletes an existing ir.model.constraint record.
func (c *Client) DeleteIrModelConstraint(id int64) error {
	return c.DeleteIrModelConstraints([]int64{id})
}

// DeleteIrModelConstraints deletes existing ir.model.constraint records.
func (c *Client) DeleteIrModelConstraints(ids []int64) error {
	return c.Delete(IrModelConstraintModel, ids)
}

// GetIrModelConstraint gets ir.model.constraint existing record.
func (c *Client) GetIrModelConstraint(id int64) (*IrModelConstraint, error) {
	imcs, err := c.GetIrModelConstraints([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imcs)[0]), nil
}

// GetIrModelConstraints gets ir.model.constraint existing records.
func (c *Client) GetIrModelConstraints(ids []int64) (*IrModelConstraints, error) {
	imcs := &IrModelConstraints{}
	if err := c.Read(IrModelConstraintModel, ids, nil, imcs); err != nil {
		return nil, err
	}
	return imcs, nil
}

// FindIrModelConstraint finds ir.model.constraint record by querying it with criteria.
func (c *Client) FindIrModelConstraint(criteria *Criteria) (*IrModelConstraint, error) {
	imcs := &IrModelConstraints{}
	if err := c.SearchRead(IrModelConstraintModel, criteria, NewOptions().Limit(1), imcs); err != nil {
		return nil, err
	}
	return &((*imcs)[0]), nil
}

// FindIrModelConstraints finds ir.model.constraint records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelConstraints(criteria *Criteria, options *Options) (*IrModelConstraints, error) {
	imcs := &IrModelConstraints{}
	if err := c.SearchRead(IrModelConstraintModel, criteria, options, imcs); err != nil {
		return nil, err
	}
	return imcs, nil
}

// FindIrModelConstraintIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelConstraintIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrModelConstraintModel, criteria, options)
}

// FindIrModelConstraintId finds record id by querying it with criteria.
func (c *Client) FindIrModelConstraintId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelConstraintModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
