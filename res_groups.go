package odoo

// ResGroups represents res.groups model.
type ResGroups struct {
	LastUpdate      *Time     `xmlrpc:"__last_update,omitempty"`
	CategoryId      *Many2One `xmlrpc:"category_id,omitempty"`
	Color           *Int      `xmlrpc:"color,omitempty"`
	Comment         *String   `xmlrpc:"comment,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	FullName        *String   `xmlrpc:"full_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	ImpliedIds      *Relation `xmlrpc:"implied_ids,omitempty"`
	IsPortal        *Bool     `xmlrpc:"is_portal,omitempty"`
	MenuAccess      *Relation `xmlrpc:"menu_access,omitempty"`
	ModelAccess     *Relation `xmlrpc:"model_access,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	RuleGroups      *Relation `xmlrpc:"rule_groups,omitempty"`
	Share           *Bool     `xmlrpc:"share,omitempty"`
	TransImpliedIds *Relation `xmlrpc:"trans_implied_ids,omitempty"`
	Users           *Relation `xmlrpc:"users,omitempty"`
	ViewAccess      *Relation `xmlrpc:"view_access,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResGroupss represents array of res.groups model.
type ResGroupss []ResGroups

// ResGroupsModel is the odoo model name.
const ResGroupsModel = "res.groups"

// Many2One convert ResGroups to *Many2One.
func (rg *ResGroups) Many2One() *Many2One {
	return NewMany2One(rg.Id.Get(), "")
}

// CreateResGroups creates a new res.groups model and returns its id.
func (c *Client) CreateResGroups(rg *ResGroups) (int64, error) {
	ids, err := c.CreateResGroupss([]*ResGroups{rg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResGroupss creates a new res.groups model and returns its id.
func (c *Client) CreateResGroupss(rgs []*ResGroups) ([]int64, error) {
	var vv []interface{}
	for _, v := range rgs {
		vv = append(vv, v)
	}
	return c.Create(ResGroupsModel, vv, nil)
}

// UpdateResGroups updates an existing res.groups record.
func (c *Client) UpdateResGroups(rg *ResGroups) error {
	return c.UpdateResGroupss([]int64{rg.Id.Get()}, rg)
}

// UpdateResGroupss updates existing res.groups records.
// All records (represented by ids) will be updated by rg values.
func (c *Client) UpdateResGroupss(ids []int64, rg *ResGroups) error {
	return c.Update(ResGroupsModel, ids, rg, nil)
}

// DeleteResGroups deletes an existing res.groups record.
func (c *Client) DeleteResGroups(id int64) error {
	return c.DeleteResGroupss([]int64{id})
}

// DeleteResGroupss deletes existing res.groups records.
func (c *Client) DeleteResGroupss(ids []int64) error {
	return c.Delete(ResGroupsModel, ids)
}

// GetResGroups gets res.groups existing record.
func (c *Client) GetResGroups(id int64) (*ResGroups, error) {
	rgs, err := c.GetResGroupss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rgs)[0]), nil
}

// GetResGroupss gets res.groups existing records.
func (c *Client) GetResGroupss(ids []int64) (*ResGroupss, error) {
	rgs := &ResGroupss{}
	if err := c.Read(ResGroupsModel, ids, nil, rgs); err != nil {
		return nil, err
	}
	return rgs, nil
}

// FindResGroups finds res.groups record by querying it with criteria.
func (c *Client) FindResGroups(criteria *Criteria) (*ResGroups, error) {
	rgs := &ResGroupss{}
	if err := c.SearchRead(ResGroupsModel, criteria, NewOptions().Limit(1), rgs); err != nil {
		return nil, err
	}
	return &((*rgs)[0]), nil
}

// FindResGroupss finds res.groups records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResGroupss(criteria *Criteria, options *Options) (*ResGroupss, error) {
	rgs := &ResGroupss{}
	if err := c.SearchRead(ResGroupsModel, criteria, options, rgs); err != nil {
		return nil, err
	}
	return rgs, nil
}

// FindResGroupsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResGroupsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResGroupsModel, criteria, options)
}

// FindResGroupsId finds record id by querying it with criteria.
func (c *Client) FindResGroupsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResGroupsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
