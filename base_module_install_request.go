package odoo

// BaseModuleInstallRequest represents base.module.install.request model.
type BaseModuleInstallRequest struct {
	BodyHtml    *String   `xmlrpc:"body_html,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ModuleId    *Many2One `xmlrpc:"module_id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	UserIds     *Relation `xmlrpc:"user_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseModuleInstallRequests represents array of base.module.install.request model.
type BaseModuleInstallRequests []BaseModuleInstallRequest

// BaseModuleInstallRequestModel is the odoo model name.
const BaseModuleInstallRequestModel = "base.module.install.request"

// Many2One convert BaseModuleInstallRequest to *Many2One.
func (bmir *BaseModuleInstallRequest) Many2One() *Many2One {
	return NewMany2One(bmir.Id.Get(), "")
}

// CreateBaseModuleInstallRequest creates a new base.module.install.request model and returns its id.
func (c *Client) CreateBaseModuleInstallRequest(bmir *BaseModuleInstallRequest) (int64, error) {
	ids, err := c.CreateBaseModuleInstallRequests([]*BaseModuleInstallRequest{bmir})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseModuleInstallRequest creates a new base.module.install.request model and returns its id.
func (c *Client) CreateBaseModuleInstallRequests(bmirs []*BaseModuleInstallRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range bmirs {
		vv = append(vv, v)
	}
	return c.Create(BaseModuleInstallRequestModel, vv, nil)
}

// UpdateBaseModuleInstallRequest updates an existing base.module.install.request record.
func (c *Client) UpdateBaseModuleInstallRequest(bmir *BaseModuleInstallRequest) error {
	return c.UpdateBaseModuleInstallRequests([]int64{bmir.Id.Get()}, bmir)
}

// UpdateBaseModuleInstallRequests updates existing base.module.install.request records.
// All records (represented by ids) will be updated by bmir values.
func (c *Client) UpdateBaseModuleInstallRequests(ids []int64, bmir *BaseModuleInstallRequest) error {
	return c.Update(BaseModuleInstallRequestModel, ids, bmir, nil)
}

// DeleteBaseModuleInstallRequest deletes an existing base.module.install.request record.
func (c *Client) DeleteBaseModuleInstallRequest(id int64) error {
	return c.DeleteBaseModuleInstallRequests([]int64{id})
}

// DeleteBaseModuleInstallRequests deletes existing base.module.install.request records.
func (c *Client) DeleteBaseModuleInstallRequests(ids []int64) error {
	return c.Delete(BaseModuleInstallRequestModel, ids)
}

// GetBaseModuleInstallRequest gets base.module.install.request existing record.
func (c *Client) GetBaseModuleInstallRequest(id int64) (*BaseModuleInstallRequest, error) {
	bmirs, err := c.GetBaseModuleInstallRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bmirs)[0]), nil
}

// GetBaseModuleInstallRequests gets base.module.install.request existing records.
func (c *Client) GetBaseModuleInstallRequests(ids []int64) (*BaseModuleInstallRequests, error) {
	bmirs := &BaseModuleInstallRequests{}
	if err := c.Read(BaseModuleInstallRequestModel, ids, nil, bmirs); err != nil {
		return nil, err
	}
	return bmirs, nil
}

// FindBaseModuleInstallRequest finds base.module.install.request record by querying it with criteria.
func (c *Client) FindBaseModuleInstallRequest(criteria *Criteria) (*BaseModuleInstallRequest, error) {
	bmirs := &BaseModuleInstallRequests{}
	if err := c.SearchRead(BaseModuleInstallRequestModel, criteria, NewOptions().Limit(1), bmirs); err != nil {
		return nil, err
	}
	return &((*bmirs)[0]), nil
}

// FindBaseModuleInstallRequests finds base.module.install.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleInstallRequests(criteria *Criteria, options *Options) (*BaseModuleInstallRequests, error) {
	bmirs := &BaseModuleInstallRequests{}
	if err := c.SearchRead(BaseModuleInstallRequestModel, criteria, options, bmirs); err != nil {
		return nil, err
	}
	return bmirs, nil
}

// FindBaseModuleInstallRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleInstallRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseModuleInstallRequestModel, criteria, options)
}

// FindBaseModuleInstallRequestId finds record id by querying it with criteria.
func (c *Client) FindBaseModuleInstallRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModuleInstallRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
