package odoo

// BaseModuleUninstall represents base.module.uninstall model.
type BaseModuleUninstall struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ModelIds    *Relation `xmlrpc:"model_ids,omitempty"`
	ModuleId    *Many2One `xmlrpc:"module_id,omitempty"`
	ModuleIds   *Relation `xmlrpc:"module_ids,omitempty"`
	ShowAll     *Bool     `xmlrpc:"show_all,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseModuleUninstalls represents array of base.module.uninstall model.
type BaseModuleUninstalls []BaseModuleUninstall

// BaseModuleUninstallModel is the odoo model name.
const BaseModuleUninstallModel = "base.module.uninstall"

// Many2One convert BaseModuleUninstall to *Many2One.
func (bmu *BaseModuleUninstall) Many2One() *Many2One {
	return NewMany2One(bmu.Id.Get(), "")
}

// CreateBaseModuleUninstall creates a new base.module.uninstall model and returns its id.
func (c *Client) CreateBaseModuleUninstall(bmu *BaseModuleUninstall) (int64, error) {
	ids, err := c.CreateBaseModuleUninstalls([]*BaseModuleUninstall{bmu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseModuleUninstall creates a new base.module.uninstall model and returns its id.
func (c *Client) CreateBaseModuleUninstalls(bmus []*BaseModuleUninstall) ([]int64, error) {
	var vv []interface{}
	for _, v := range bmus {
		vv = append(vv, v)
	}
	return c.Create(BaseModuleUninstallModel, vv, nil)
}

// UpdateBaseModuleUninstall updates an existing base.module.uninstall record.
func (c *Client) UpdateBaseModuleUninstall(bmu *BaseModuleUninstall) error {
	return c.UpdateBaseModuleUninstalls([]int64{bmu.Id.Get()}, bmu)
}

// UpdateBaseModuleUninstalls updates existing base.module.uninstall records.
// All records (represented by ids) will be updated by bmu values.
func (c *Client) UpdateBaseModuleUninstalls(ids []int64, bmu *BaseModuleUninstall) error {
	return c.Update(BaseModuleUninstallModel, ids, bmu, nil)
}

// DeleteBaseModuleUninstall deletes an existing base.module.uninstall record.
func (c *Client) DeleteBaseModuleUninstall(id int64) error {
	return c.DeleteBaseModuleUninstalls([]int64{id})
}

// DeleteBaseModuleUninstalls deletes existing base.module.uninstall records.
func (c *Client) DeleteBaseModuleUninstalls(ids []int64) error {
	return c.Delete(BaseModuleUninstallModel, ids)
}

// GetBaseModuleUninstall gets base.module.uninstall existing record.
func (c *Client) GetBaseModuleUninstall(id int64) (*BaseModuleUninstall, error) {
	bmus, err := c.GetBaseModuleUninstalls([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bmus)[0]), nil
}

// GetBaseModuleUninstalls gets base.module.uninstall existing records.
func (c *Client) GetBaseModuleUninstalls(ids []int64) (*BaseModuleUninstalls, error) {
	bmus := &BaseModuleUninstalls{}
	if err := c.Read(BaseModuleUninstallModel, ids, nil, bmus); err != nil {
		return nil, err
	}
	return bmus, nil
}

// FindBaseModuleUninstall finds base.module.uninstall record by querying it with criteria.
func (c *Client) FindBaseModuleUninstall(criteria *Criteria) (*BaseModuleUninstall, error) {
	bmus := &BaseModuleUninstalls{}
	if err := c.SearchRead(BaseModuleUninstallModel, criteria, NewOptions().Limit(1), bmus); err != nil {
		return nil, err
	}
	return &((*bmus)[0]), nil
}

// FindBaseModuleUninstalls finds base.module.uninstall records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleUninstalls(criteria *Criteria, options *Options) (*BaseModuleUninstalls, error) {
	bmus := &BaseModuleUninstalls{}
	if err := c.SearchRead(BaseModuleUninstallModel, criteria, options, bmus); err != nil {
		return nil, err
	}
	return bmus, nil
}

// FindBaseModuleUninstallIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleUninstallIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseModuleUninstallModel, criteria, options)
}

// FindBaseModuleUninstallId finds record id by querying it with criteria.
func (c *Client) FindBaseModuleUninstallId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModuleUninstallModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
