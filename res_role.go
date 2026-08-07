package odoo

// ResRole represents res.role model.
type ResRole struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	UserIds     *Relation `xmlrpc:"user_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResRoles represents array of res.role model.
type ResRoles []ResRole

// ResRoleModel is the odoo model name.
const ResRoleModel = "res.role"

// Many2One convert ResRole to *Many2One.
func (rr *ResRole) Many2One() *Many2One {
	return NewMany2One(rr.Id.Get(), "")
}

// CreateResRole creates a new res.role model and returns its id.
func (c *Client) CreateResRole(rr *ResRole) (int64, error) {
	ids, err := c.CreateResRoles([]*ResRole{rr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResRole creates a new res.role model and returns its id.
func (c *Client) CreateResRoles(rrs []*ResRole) ([]int64, error) {
	var vv []interface{}
	for _, v := range rrs {
		vv = append(vv, v)
	}
	return c.Create(ResRoleModel, vv, nil)
}

// UpdateResRole updates an existing res.role record.
func (c *Client) UpdateResRole(rr *ResRole) error {
	return c.UpdateResRoles([]int64{rr.Id.Get()}, rr)
}

// UpdateResRoles updates existing res.role records.
// All records (represented by ids) will be updated by rr values.
func (c *Client) UpdateResRoles(ids []int64, rr *ResRole) error {
	return c.Update(ResRoleModel, ids, rr, nil)
}

// DeleteResRole deletes an existing res.role record.
func (c *Client) DeleteResRole(id int64) error {
	return c.DeleteResRoles([]int64{id})
}

// DeleteResRoles deletes existing res.role records.
func (c *Client) DeleteResRoles(ids []int64) error {
	return c.Delete(ResRoleModel, ids)
}

// GetResRole gets res.role existing record.
func (c *Client) GetResRole(id int64) (*ResRole, error) {
	rrs, err := c.GetResRoles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rrs)[0]), nil
}

// GetResRoles gets res.role existing records.
func (c *Client) GetResRoles(ids []int64) (*ResRoles, error) {
	rrs := &ResRoles{}
	if err := c.Read(ResRoleModel, ids, nil, rrs); err != nil {
		return nil, err
	}
	return rrs, nil
}

// FindResRole finds res.role record by querying it with criteria.
func (c *Client) FindResRole(criteria *Criteria) (*ResRole, error) {
	rrs := &ResRoles{}
	if err := c.SearchRead(ResRoleModel, criteria, NewOptions().Limit(1), rrs); err != nil {
		return nil, err
	}
	return &((*rrs)[0]), nil
}

// FindResRoles finds res.role records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResRoles(criteria *Criteria, options *Options) (*ResRoles, error) {
	rrs := &ResRoles{}
	if err := c.SearchRead(ResRoleModel, criteria, options, rrs); err != nil {
		return nil, err
	}
	return rrs, nil
}

// FindResRoleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResRoleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResRoleModel, criteria, options)
}

// FindResRoleId finds record id by querying it with criteria.
func (c *Client) FindResRoleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResRoleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
