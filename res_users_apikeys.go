package odoo

// ResUsersApikeys represents res.users.apikeys model.
type ResUsersApikeys struct {
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	ExpirationDate *Time     `xmlrpc:"expiration_date,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	Scope          *String   `xmlrpc:"scope,omitempty"`
	UserId         *Many2One `xmlrpc:"user_id,omitempty"`
}

// ResUsersApikeyss represents array of res.users.apikeys model.
type ResUsersApikeyss []ResUsersApikeys

// ResUsersApikeysModel is the odoo model name.
const ResUsersApikeysModel = "res.users.apikeys"

// Many2One convert ResUsersApikeys to *Many2One.
func (rua *ResUsersApikeys) Many2One() *Many2One {
	return NewMany2One(rua.Id.Get(), "")
}

// CreateResUsersApikeys creates a new res.users.apikeys model and returns its id.
func (c *Client) CreateResUsersApikeys(rua *ResUsersApikeys) (int64, error) {
	ids, err := c.CreateResUsersApikeyss([]*ResUsersApikeys{rua})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersApikeys creates a new res.users.apikeys model and returns its id.
func (c *Client) CreateResUsersApikeyss(ruas []*ResUsersApikeys) ([]int64, error) {
	var vv []interface{}
	for _, v := range ruas {
		vv = append(vv, v)
	}
	return c.Create(ResUsersApikeysModel, vv, nil)
}

// UpdateResUsersApikeys updates an existing res.users.apikeys record.
func (c *Client) UpdateResUsersApikeys(rua *ResUsersApikeys) error {
	return c.UpdateResUsersApikeyss([]int64{rua.Id.Get()}, rua)
}

// UpdateResUsersApikeyss updates existing res.users.apikeys records.
// All records (represented by ids) will be updated by rua values.
func (c *Client) UpdateResUsersApikeyss(ids []int64, rua *ResUsersApikeys) error {
	return c.Update(ResUsersApikeysModel, ids, rua, nil)
}

// DeleteResUsersApikeys deletes an existing res.users.apikeys record.
func (c *Client) DeleteResUsersApikeys(id int64) error {
	return c.DeleteResUsersApikeyss([]int64{id})
}

// DeleteResUsersApikeyss deletes existing res.users.apikeys records.
func (c *Client) DeleteResUsersApikeyss(ids []int64) error {
	return c.Delete(ResUsersApikeysModel, ids)
}

// GetResUsersApikeys gets res.users.apikeys existing record.
func (c *Client) GetResUsersApikeys(id int64) (*ResUsersApikeys, error) {
	ruas, err := c.GetResUsersApikeyss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ruas)[0]), nil
}

// GetResUsersApikeyss gets res.users.apikeys existing records.
func (c *Client) GetResUsersApikeyss(ids []int64) (*ResUsersApikeyss, error) {
	ruas := &ResUsersApikeyss{}
	if err := c.Read(ResUsersApikeysModel, ids, nil, ruas); err != nil {
		return nil, err
	}
	return ruas, nil
}

// FindResUsersApikeys finds res.users.apikeys record by querying it with criteria.
func (c *Client) FindResUsersApikeys(criteria *Criteria) (*ResUsersApikeys, error) {
	ruas := &ResUsersApikeyss{}
	if err := c.SearchRead(ResUsersApikeysModel, criteria, NewOptions().Limit(1), ruas); err != nil {
		return nil, err
	}
	return &((*ruas)[0]), nil
}

// FindResUsersApikeyss finds res.users.apikeys records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersApikeyss(criteria *Criteria, options *Options) (*ResUsersApikeyss, error) {
	ruas := &ResUsersApikeyss{}
	if err := c.SearchRead(ResUsersApikeysModel, criteria, options, ruas); err != nil {
		return nil, err
	}
	return ruas, nil
}

// FindResUsersApikeysIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersApikeysIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResUsersApikeysModel, criteria, options)
}

// FindResUsersApikeysId finds record id by querying it with criteria.
func (c *Client) FindResUsersApikeysId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersApikeysModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
