package odoo

// ResUsersIdentitycheck represents res.users.identitycheck model.
type ResUsersIdentitycheck struct {
	AuthMethod  *Selection `xmlrpc:"auth_method,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Password    *String    `xmlrpc:"password,omitempty"`
	Request     *String    `xmlrpc:"request,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResUsersIdentitychecks represents array of res.users.identitycheck model.
type ResUsersIdentitychecks []ResUsersIdentitycheck

// ResUsersIdentitycheckModel is the odoo model name.
const ResUsersIdentitycheckModel = "res.users.identitycheck"

// Many2One convert ResUsersIdentitycheck to *Many2One.
func (rui *ResUsersIdentitycheck) Many2One() *Many2One {
	return NewMany2One(rui.Id.Get(), "")
}

// CreateResUsersIdentitycheck creates a new res.users.identitycheck model and returns its id.
func (c *Client) CreateResUsersIdentitycheck(rui *ResUsersIdentitycheck) (int64, error) {
	ids, err := c.CreateResUsersIdentitychecks([]*ResUsersIdentitycheck{rui})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersIdentitycheck creates a new res.users.identitycheck model and returns its id.
func (c *Client) CreateResUsersIdentitychecks(ruis []*ResUsersIdentitycheck) ([]int64, error) {
	var vv []interface{}
	for _, v := range ruis {
		vv = append(vv, v)
	}
	return c.Create(ResUsersIdentitycheckModel, vv, nil)
}

// UpdateResUsersIdentitycheck updates an existing res.users.identitycheck record.
func (c *Client) UpdateResUsersIdentitycheck(rui *ResUsersIdentitycheck) error {
	return c.UpdateResUsersIdentitychecks([]int64{rui.Id.Get()}, rui)
}

// UpdateResUsersIdentitychecks updates existing res.users.identitycheck records.
// All records (represented by ids) will be updated by rui values.
func (c *Client) UpdateResUsersIdentitychecks(ids []int64, rui *ResUsersIdentitycheck) error {
	return c.Update(ResUsersIdentitycheckModel, ids, rui, nil)
}

// DeleteResUsersIdentitycheck deletes an existing res.users.identitycheck record.
func (c *Client) DeleteResUsersIdentitycheck(id int64) error {
	return c.DeleteResUsersIdentitychecks([]int64{id})
}

// DeleteResUsersIdentitychecks deletes existing res.users.identitycheck records.
func (c *Client) DeleteResUsersIdentitychecks(ids []int64) error {
	return c.Delete(ResUsersIdentitycheckModel, ids)
}

// GetResUsersIdentitycheck gets res.users.identitycheck existing record.
func (c *Client) GetResUsersIdentitycheck(id int64) (*ResUsersIdentitycheck, error) {
	ruis, err := c.GetResUsersIdentitychecks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ruis)[0]), nil
}

// GetResUsersIdentitychecks gets res.users.identitycheck existing records.
func (c *Client) GetResUsersIdentitychecks(ids []int64) (*ResUsersIdentitychecks, error) {
	ruis := &ResUsersIdentitychecks{}
	if err := c.Read(ResUsersIdentitycheckModel, ids, nil, ruis); err != nil {
		return nil, err
	}
	return ruis, nil
}

// FindResUsersIdentitycheck finds res.users.identitycheck record by querying it with criteria.
func (c *Client) FindResUsersIdentitycheck(criteria *Criteria) (*ResUsersIdentitycheck, error) {
	ruis := &ResUsersIdentitychecks{}
	if err := c.SearchRead(ResUsersIdentitycheckModel, criteria, NewOptions().Limit(1), ruis); err != nil {
		return nil, err
	}
	return &((*ruis)[0]), nil
}

// FindResUsersIdentitychecks finds res.users.identitycheck records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersIdentitychecks(criteria *Criteria, options *Options) (*ResUsersIdentitychecks, error) {
	ruis := &ResUsersIdentitychecks{}
	if err := c.SearchRead(ResUsersIdentitycheckModel, criteria, options, ruis); err != nil {
		return nil, err
	}
	return ruis, nil
}

// FindResUsersIdentitycheckIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersIdentitycheckIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResUsersIdentitycheckModel, criteria, options)
}

// FindResUsersIdentitycheckId finds record id by querying it with criteria.
func (c *Client) FindResUsersIdentitycheckId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersIdentitycheckModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
