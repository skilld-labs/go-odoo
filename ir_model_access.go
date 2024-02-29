package odoo

// IrModelAccess represents ir.model.access model.
type IrModelAccess struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	GroupId     *Many2One `xmlrpc:"group_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ModelId     *Many2One `xmlrpc:"model_id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	PermCreate  *Bool     `xmlrpc:"perm_create,omitempty"`
	PermRead    *Bool     `xmlrpc:"perm_read,omitempty"`
	PermUnlink  *Bool     `xmlrpc:"perm_unlink,omitempty"`
	PermWrite   *Bool     `xmlrpc:"perm_write,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrModelAccesss represents array of ir.model.access model.
type IrModelAccesss []IrModelAccess

// IrModelAccessModel is the odoo model name.
const IrModelAccessModel = "ir.model.access"

// Many2One convert IrModelAccess to *Many2One.
func (ima *IrModelAccess) Many2One() *Many2One {
	return NewMany2One(ima.Id.Get(), "")
}

// CreateIrModelAccess creates a new ir.model.access model and returns its id.
func (c *Client) CreateIrModelAccess(ima *IrModelAccess) (int64, error) {
	ids, err := c.CreateIrModelAccesss([]*IrModelAccess{ima})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrModelAccesss creates a new ir.model.access model and returns its id.
func (c *Client) CreateIrModelAccesss(imas []*IrModelAccess) ([]int64, error) {
	var vv []interface{}
	for _, v := range imas {
		vv = append(vv, v)
	}
	return c.Create(IrModelAccessModel, vv, nil)
}

// UpdateIrModelAccess updates an existing ir.model.access record.
func (c *Client) UpdateIrModelAccess(ima *IrModelAccess) error {
	return c.UpdateIrModelAccesss([]int64{ima.Id.Get()}, ima)
}

// UpdateIrModelAccesss updates existing ir.model.access records.
// All records (represented by ids) will be updated by ima values.
func (c *Client) UpdateIrModelAccesss(ids []int64, ima *IrModelAccess) error {
	return c.Update(IrModelAccessModel, ids, ima, nil)
}

// DeleteIrModelAccess deletes an existing ir.model.access record.
func (c *Client) DeleteIrModelAccess(id int64) error {
	return c.DeleteIrModelAccesss([]int64{id})
}

// DeleteIrModelAccesss deletes existing ir.model.access records.
func (c *Client) DeleteIrModelAccesss(ids []int64) error {
	return c.Delete(IrModelAccessModel, ids)
}

// GetIrModelAccess gets ir.model.access existing record.
func (c *Client) GetIrModelAccess(id int64) (*IrModelAccess, error) {
	imas, err := c.GetIrModelAccesss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*imas)[0]), nil
}

// GetIrModelAccesss gets ir.model.access existing records.
func (c *Client) GetIrModelAccesss(ids []int64) (*IrModelAccesss, error) {
	imas := &IrModelAccesss{}
	if err := c.Read(IrModelAccessModel, ids, nil, imas); err != nil {
		return nil, err
	}
	return imas, nil
}

// FindIrModelAccess finds ir.model.access record by querying it with criteria.
func (c *Client) FindIrModelAccess(criteria *Criteria) (*IrModelAccess, error) {
	imas := &IrModelAccesss{}
	if err := c.SearchRead(IrModelAccessModel, criteria, NewOptions().Limit(1), imas); err != nil {
		return nil, err
	}
	return &((*imas)[0]), nil
}

// FindIrModelAccesss finds ir.model.access records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelAccesss(criteria *Criteria, options *Options) (*IrModelAccesss, error) {
	imas := &IrModelAccesss{}
	if err := c.SearchRead(IrModelAccessModel, criteria, options, imas); err != nil {
		return nil, err
	}
	return imas, nil
}

// FindIrModelAccessIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrModelAccessIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrModelAccessModel, criteria, options)
}

// FindIrModelAccessId finds record id by querying it with criteria.
func (c *Client) FindIrModelAccessId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrModelAccessModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
