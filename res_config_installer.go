package odoo

// ResConfigInstaller represents res.config.installer model.
type ResConfigInstaller struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResConfigInstallers represents array of res.config.installer model.
type ResConfigInstallers []ResConfigInstaller

// ResConfigInstallerModel is the odoo model name.
const ResConfigInstallerModel = "res.config.installer"

// Many2One convert ResConfigInstaller to *Many2One.
func (rci *ResConfigInstaller) Many2One() *Many2One {
	return NewMany2One(rci.Id.Get(), "")
}

// CreateResConfigInstaller creates a new res.config.installer model and returns its id.
func (c *Client) CreateResConfigInstaller(rci *ResConfigInstaller) (int64, error) {
	ids, err := c.CreateResConfigInstallers([]*ResConfigInstaller{rci})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResConfigInstallers creates a new res.config.installer model and returns its id.
func (c *Client) CreateResConfigInstallers(rcis []*ResConfigInstaller) ([]int64, error) {
	var vv []interface{}
	for _, v := range rcis {
		vv = append(vv, v)
	}
	return c.Create(ResConfigInstallerModel, vv, nil)
}

// UpdateResConfigInstaller updates an existing res.config.installer record.
func (c *Client) UpdateResConfigInstaller(rci *ResConfigInstaller) error {
	return c.UpdateResConfigInstallers([]int64{rci.Id.Get()}, rci)
}

// UpdateResConfigInstallers updates existing res.config.installer records.
// All records (represented by ids) will be updated by rci values.
func (c *Client) UpdateResConfigInstallers(ids []int64, rci *ResConfigInstaller) error {
	return c.Update(ResConfigInstallerModel, ids, rci, nil)
}

// DeleteResConfigInstaller deletes an existing res.config.installer record.
func (c *Client) DeleteResConfigInstaller(id int64) error {
	return c.DeleteResConfigInstallers([]int64{id})
}

// DeleteResConfigInstallers deletes existing res.config.installer records.
func (c *Client) DeleteResConfigInstallers(ids []int64) error {
	return c.Delete(ResConfigInstallerModel, ids)
}

// GetResConfigInstaller gets res.config.installer existing record.
func (c *Client) GetResConfigInstaller(id int64) (*ResConfigInstaller, error) {
	rcis, err := c.GetResConfigInstallers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rcis)[0]), nil
}

// GetResConfigInstallers gets res.config.installer existing records.
func (c *Client) GetResConfigInstallers(ids []int64) (*ResConfigInstallers, error) {
	rcis := &ResConfigInstallers{}
	if err := c.Read(ResConfigInstallerModel, ids, nil, rcis); err != nil {
		return nil, err
	}
	return rcis, nil
}

// FindResConfigInstaller finds res.config.installer record by querying it with criteria.
func (c *Client) FindResConfigInstaller(criteria *Criteria) (*ResConfigInstaller, error) {
	rcis := &ResConfigInstallers{}
	if err := c.SearchRead(ResConfigInstallerModel, criteria, NewOptions().Limit(1), rcis); err != nil {
		return nil, err
	}
	return &((*rcis)[0]), nil
}

// FindResConfigInstallers finds res.config.installer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigInstallers(criteria *Criteria, options *Options) (*ResConfigInstallers, error) {
	rcis := &ResConfigInstallers{}
	if err := c.SearchRead(ResConfigInstallerModel, criteria, options, rcis); err != nil {
		return nil, err
	}
	return rcis, nil
}

// FindResConfigInstallerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResConfigInstallerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResConfigInstallerModel, criteria, options)
}

// FindResConfigInstallerId finds record id by querying it with criteria.
func (c *Client) FindResConfigInstallerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResConfigInstallerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
