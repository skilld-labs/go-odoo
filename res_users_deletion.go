package odoo

// ResUsersDeletion represents res.users.deletion model.
type ResUsersDeletion struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	State       *Selection `xmlrpc:"state,omitempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omitempty"`
	UserIdInt   *Int       `xmlrpc:"user_id_int,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResUsersDeletions represents array of res.users.deletion model.
type ResUsersDeletions []ResUsersDeletion

// ResUsersDeletionModel is the odoo model name.
const ResUsersDeletionModel = "res.users.deletion"

// Many2One convert ResUsersDeletion to *Many2One.
func (rud *ResUsersDeletion) Many2One() *Many2One {
	return NewMany2One(rud.Id.Get(), "")
}

// CreateResUsersDeletion creates a new res.users.deletion model and returns its id.
func (c *Client) CreateResUsersDeletion(rud *ResUsersDeletion) (int64, error) {
	ids, err := c.CreateResUsersDeletions([]*ResUsersDeletion{rud})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersDeletion creates a new res.users.deletion model and returns its id.
func (c *Client) CreateResUsersDeletions(ruds []*ResUsersDeletion) ([]int64, error) {
	var vv []interface{}
	for _, v := range ruds {
		vv = append(vv, v)
	}
	return c.Create(ResUsersDeletionModel, vv, nil)
}

// UpdateResUsersDeletion updates an existing res.users.deletion record.
func (c *Client) UpdateResUsersDeletion(rud *ResUsersDeletion) error {
	return c.UpdateResUsersDeletions([]int64{rud.Id.Get()}, rud)
}

// UpdateResUsersDeletions updates existing res.users.deletion records.
// All records (represented by ids) will be updated by rud values.
func (c *Client) UpdateResUsersDeletions(ids []int64, rud *ResUsersDeletion) error {
	return c.Update(ResUsersDeletionModel, ids, rud, nil)
}

// DeleteResUsersDeletion deletes an existing res.users.deletion record.
func (c *Client) DeleteResUsersDeletion(id int64) error {
	return c.DeleteResUsersDeletions([]int64{id})
}

// DeleteResUsersDeletions deletes existing res.users.deletion records.
func (c *Client) DeleteResUsersDeletions(ids []int64) error {
	return c.Delete(ResUsersDeletionModel, ids)
}

// GetResUsersDeletion gets res.users.deletion existing record.
func (c *Client) GetResUsersDeletion(id int64) (*ResUsersDeletion, error) {
	ruds, err := c.GetResUsersDeletions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ruds)[0]), nil
}

// GetResUsersDeletions gets res.users.deletion existing records.
func (c *Client) GetResUsersDeletions(ids []int64) (*ResUsersDeletions, error) {
	ruds := &ResUsersDeletions{}
	if err := c.Read(ResUsersDeletionModel, ids, nil, ruds); err != nil {
		return nil, err
	}
	return ruds, nil
}

// FindResUsersDeletion finds res.users.deletion record by querying it with criteria.
func (c *Client) FindResUsersDeletion(criteria *Criteria) (*ResUsersDeletion, error) {
	ruds := &ResUsersDeletions{}
	if err := c.SearchRead(ResUsersDeletionModel, criteria, NewOptions().Limit(1), ruds); err != nil {
		return nil, err
	}
	return &((*ruds)[0]), nil
}

// FindResUsersDeletions finds res.users.deletion records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersDeletions(criteria *Criteria, options *Options) (*ResUsersDeletions, error) {
	ruds := &ResUsersDeletions{}
	if err := c.SearchRead(ResUsersDeletionModel, criteria, options, ruds); err != nil {
		return nil, err
	}
	return ruds, nil
}

// FindResUsersDeletionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersDeletionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResUsersDeletionModel, criteria, options)
}

// FindResUsersDeletionId finds record id by querying it with criteria.
func (c *Client) FindResUsersDeletionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersDeletionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
