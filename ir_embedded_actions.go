package odoo

// IrEmbeddedActions represents ir.embedded.actions model.
type IrEmbeddedActions struct {
	ActionId        *Many2One `xmlrpc:"action_id,omitempty"`
	Context         *String   `xmlrpc:"context,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DefaultViewMode *String   `xmlrpc:"default_view_mode,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Domain          *String   `xmlrpc:"domain,omitempty"`
	FilterIds       *Relation `xmlrpc:"filter_ids,omitempty"`
	GroupsIds       *Relation `xmlrpc:"groups_ids,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	IsDeletable     *Bool     `xmlrpc:"is_deletable,omitempty"`
	IsVisible       *Bool     `xmlrpc:"is_visible,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	ParentActionId  *Many2One `xmlrpc:"parent_action_id,omitempty"`
	ParentResId     *Int      `xmlrpc:"parent_res_id,omitempty"`
	ParentResModel  *String   `xmlrpc:"parent_res_model,omitempty"`
	PythonMethod    *String   `xmlrpc:"python_method,omitempty"`
	Sequence        *Int      `xmlrpc:"sequence,omitempty"`
	UserId          *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrEmbeddedActionss represents array of ir.embedded.actions model.
type IrEmbeddedActionss []IrEmbeddedActions

// IrEmbeddedActionsModel is the odoo model name.
const IrEmbeddedActionsModel = "ir.embedded.actions"

// Many2One convert IrEmbeddedActions to *Many2One.
func (iea *IrEmbeddedActions) Many2One() *Many2One {
	return NewMany2One(iea.Id.Get(), "")
}

// CreateIrEmbeddedActions creates a new ir.embedded.actions model and returns its id.
func (c *Client) CreateIrEmbeddedActions(iea *IrEmbeddedActions) (int64, error) {
	ids, err := c.CreateIrEmbeddedActionss([]*IrEmbeddedActions{iea})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrEmbeddedActions creates a new ir.embedded.actions model and returns its id.
func (c *Client) CreateIrEmbeddedActionss(ieas []*IrEmbeddedActions) ([]int64, error) {
	var vv []interface{}
	for _, v := range ieas {
		vv = append(vv, v)
	}
	return c.Create(IrEmbeddedActionsModel, vv, nil)
}

// UpdateIrEmbeddedActions updates an existing ir.embedded.actions record.
func (c *Client) UpdateIrEmbeddedActions(iea *IrEmbeddedActions) error {
	return c.UpdateIrEmbeddedActionss([]int64{iea.Id.Get()}, iea)
}

// UpdateIrEmbeddedActionss updates existing ir.embedded.actions records.
// All records (represented by ids) will be updated by iea values.
func (c *Client) UpdateIrEmbeddedActionss(ids []int64, iea *IrEmbeddedActions) error {
	return c.Update(IrEmbeddedActionsModel, ids, iea, nil)
}

// DeleteIrEmbeddedActions deletes an existing ir.embedded.actions record.
func (c *Client) DeleteIrEmbeddedActions(id int64) error {
	return c.DeleteIrEmbeddedActionss([]int64{id})
}

// DeleteIrEmbeddedActionss deletes existing ir.embedded.actions records.
func (c *Client) DeleteIrEmbeddedActionss(ids []int64) error {
	return c.Delete(IrEmbeddedActionsModel, ids)
}

// GetIrEmbeddedActions gets ir.embedded.actions existing record.
func (c *Client) GetIrEmbeddedActions(id int64) (*IrEmbeddedActions, error) {
	ieas, err := c.GetIrEmbeddedActionss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ieas)[0]), nil
}

// GetIrEmbeddedActionss gets ir.embedded.actions existing records.
func (c *Client) GetIrEmbeddedActionss(ids []int64) (*IrEmbeddedActionss, error) {
	ieas := &IrEmbeddedActionss{}
	if err := c.Read(IrEmbeddedActionsModel, ids, nil, ieas); err != nil {
		return nil, err
	}
	return ieas, nil
}

// FindIrEmbeddedActions finds ir.embedded.actions record by querying it with criteria.
func (c *Client) FindIrEmbeddedActions(criteria *Criteria) (*IrEmbeddedActions, error) {
	ieas := &IrEmbeddedActionss{}
	if err := c.SearchRead(IrEmbeddedActionsModel, criteria, NewOptions().Limit(1), ieas); err != nil {
		return nil, err
	}
	return &((*ieas)[0]), nil
}

// FindIrEmbeddedActionss finds ir.embedded.actions records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrEmbeddedActionss(criteria *Criteria, options *Options) (*IrEmbeddedActionss, error) {
	ieas := &IrEmbeddedActionss{}
	if err := c.SearchRead(IrEmbeddedActionsModel, criteria, options, ieas); err != nil {
		return nil, err
	}
	return ieas, nil
}

// FindIrEmbeddedActionsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrEmbeddedActionsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrEmbeddedActionsModel, criteria, options)
}

// FindIrEmbeddedActionsId finds record id by querying it with criteria.
func (c *Client) FindIrEmbeddedActionsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrEmbeddedActionsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
