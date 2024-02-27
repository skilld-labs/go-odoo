package odoo

// ResConfig represents res.config model.
type ResConfig struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResConfigs represents array of res.config model.
type ResConfigs []ResConfig

// ResConfigModel is the odoo model name.
const ResConfigModel = "res.config"

// Many2One convert ResConfig to *Many2One.
func (rc *ResConfig) Many2One() *Many2One {
	return NewMany2One(rc.Id.Get(), "")
}

// CreateResConfig creates a new res.config model and returns its id.
func (c *Client) CreateResConfig(rc *ResConfig) (int64, error) {
	ids, err := c.CreateResConfigs([]*ResConfig{rc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfig creates a new res.config model and returns its id.
func (c *Client) CreateResConfigs(rcs []*ResConfig) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcs {
		vv = append(vv, v)
	}
	return c.Create(ResConfigModel, vv, nil)
}

// UpdateResConfig updates an existing res.config record.
func (c *Client) UpdateResConfig(rc *ResConfig) error {
	return c.UpdateResConfigs([]int64{rc.Id.Get()}, rc)
}

// UpdateResConfigs updates existing res.config records.
// All records (represented by ids) will be updated by rc values.
func (c *Client) UpdateResConfigs(ids []int64, rc *ResConfig) error {
	return c.Update(ResConfigModel, ids, rc, nil)
}

// DeleteResConfig deletes an existing res.config record.
func (c *Client) DeleteResConfig(id int64) error {
	return c.DeleteResConfigs([]int64{id})
}

// DeleteResConfigs deletes existing res.config records.
func (c *Client) DeleteResConfigs(ids []int64) error {
	return c.Delete(ResConfigModel, ids)
}

// GetResConfig gets res.config existing record.
func (c *Client) GetResConfig(id int64) (*ResConfig, error) {
	rcs, err := c.GetResConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// GetResConfigs gets res.config existing records.
func (c *Client) GetResConfigs(ids []int64) (*ResConfigs, error) {
	rcs := &ResConfigs{}
	if err := c.Read(ResConfigModel, ids, nil, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResConfig finds res.config record by querying it with criteria.
func (c *Client) FindResConfig(criteria *Criteria) (*ResConfig, error) {
	rcs := &ResConfigs{}
	if err := c.SearchRead(ResConfigModel, criteria, NewOptions().Limit(1), rcs); err != nil {
		return nil, err
	}
	return &((*rcs)[0]), nil
}

// FindResConfigs finds res.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigs(criteria *Criteria, options *Options) (*ResConfigs, error) {
	rcs := &ResConfigs{}
	if err := c.SearchRead(ResConfigModel, criteria, options, rcs); err != nil {
		return nil, err
	}
	return rcs, nil
}

// FindResConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResConfigModel, criteria, options)
}

// FindResConfigId finds record id by querying it with criteria.
func (c *Client) FindResConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
