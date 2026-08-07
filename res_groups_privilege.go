package odoo

// ResGroupsPrivilege represents res.groups.privilege model.
type ResGroupsPrivilege struct {
	CategoryId  *Many2One `xmlrpc:"category_id,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	Description *String   `xmlrpc:"description,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	GroupIds    *Relation `xmlrpc:"group_ids,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Placeholder *String   `xmlrpc:"placeholder,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResGroupsPrivileges represents array of res.groups.privilege model.
type ResGroupsPrivileges []ResGroupsPrivilege

// ResGroupsPrivilegeModel is the odoo model name.
const ResGroupsPrivilegeModel = "res.groups.privilege"

// Many2One convert ResGroupsPrivilege to *Many2One.
func (rgp *ResGroupsPrivilege) Many2One() *Many2One {
	return NewMany2One(rgp.Id.Get(), "")
}

// CreateResGroupsPrivilege creates a new res.groups.privilege model and returns its id.
func (c *Client) CreateResGroupsPrivilege(rgp *ResGroupsPrivilege) (int64, error) {
	ids, err := c.CreateResGroupsPrivileges([]*ResGroupsPrivilege{rgp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResGroupsPrivilege creates a new res.groups.privilege model and returns its id.
func (c *Client) CreateResGroupsPrivileges(rgps []*ResGroupsPrivilege) ([]int64, error) {
	var vv []interface{}
	for _, v := range rgps {
		vv = append(vv, v)
	}
	return c.Create(ResGroupsPrivilegeModel, vv, nil)
}

// UpdateResGroupsPrivilege updates an existing res.groups.privilege record.
func (c *Client) UpdateResGroupsPrivilege(rgp *ResGroupsPrivilege) error {
	return c.UpdateResGroupsPrivileges([]int64{rgp.Id.Get()}, rgp)
}

// UpdateResGroupsPrivileges updates existing res.groups.privilege records.
// All records (represented by ids) will be updated by rgp values.
func (c *Client) UpdateResGroupsPrivileges(ids []int64, rgp *ResGroupsPrivilege) error {
	return c.Update(ResGroupsPrivilegeModel, ids, rgp, nil)
}

// DeleteResGroupsPrivilege deletes an existing res.groups.privilege record.
func (c *Client) DeleteResGroupsPrivilege(id int64) error {
	return c.DeleteResGroupsPrivileges([]int64{id})
}

// DeleteResGroupsPrivileges deletes existing res.groups.privilege records.
func (c *Client) DeleteResGroupsPrivileges(ids []int64) error {
	return c.Delete(ResGroupsPrivilegeModel, ids)
}

// GetResGroupsPrivilege gets res.groups.privilege existing record.
func (c *Client) GetResGroupsPrivilege(id int64) (*ResGroupsPrivilege, error) {
	rgps, err := c.GetResGroupsPrivileges([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rgps)[0]), nil
}

// GetResGroupsPrivileges gets res.groups.privilege existing records.
func (c *Client) GetResGroupsPrivileges(ids []int64) (*ResGroupsPrivileges, error) {
	rgps := &ResGroupsPrivileges{}
	if err := c.Read(ResGroupsPrivilegeModel, ids, nil, rgps); err != nil {
		return nil, err
	}
	return rgps, nil
}

// FindResGroupsPrivilege finds res.groups.privilege record by querying it with criteria.
func (c *Client) FindResGroupsPrivilege(criteria *Criteria) (*ResGroupsPrivilege, error) {
	rgps := &ResGroupsPrivileges{}
	if err := c.SearchRead(ResGroupsPrivilegeModel, criteria, NewOptions().Limit(1), rgps); err != nil {
		return nil, err
	}
	return &((*rgps)[0]), nil
}

// FindResGroupsPrivileges finds res.groups.privilege records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResGroupsPrivileges(criteria *Criteria, options *Options) (*ResGroupsPrivileges, error) {
	rgps := &ResGroupsPrivileges{}
	if err := c.SearchRead(ResGroupsPrivilegeModel, criteria, options, rgps); err != nil {
		return nil, err
	}
	return rgps, nil
}

// FindResGroupsPrivilegeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResGroupsPrivilegeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResGroupsPrivilegeModel, criteria, options)
}

// FindResGroupsPrivilegeId finds record id by querying it with criteria.
func (c *Client) FindResGroupsPrivilegeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResGroupsPrivilegeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
