package odoo

// ResUsersSettingsEmbeddedAction represents res.users.settings.embedded.action model.
type ResUsersSettingsEmbeddedAction struct {
	ActionId                  *Many2One `xmlrpc:"action_id,omitempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	EmbeddedActionsOrder      *String   `xmlrpc:"embedded_actions_order,omitempty"`
	EmbeddedActionsVisibility *String   `xmlrpc:"embedded_actions_visibility,omitempty"`
	EmbeddedVisibility        *Bool     `xmlrpc:"embedded_visibility,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	ResId                     *Int      `xmlrpc:"res_id,omitempty"`
	ResModel                  *String   `xmlrpc:"res_model,omitempty"`
	UserSettingId             *Many2One `xmlrpc:"user_setting_id,omitempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResUsersSettingsEmbeddedActions represents array of res.users.settings.embedded.action model.
type ResUsersSettingsEmbeddedActions []ResUsersSettingsEmbeddedAction

// ResUsersSettingsEmbeddedActionModel is the odoo model name.
const ResUsersSettingsEmbeddedActionModel = "res.users.settings.embedded.action"

// Many2One convert ResUsersSettingsEmbeddedAction to *Many2One.
func (rusea *ResUsersSettingsEmbeddedAction) Many2One() *Many2One {
	return NewMany2One(rusea.Id.Get(), "")
}

// CreateResUsersSettingsEmbeddedAction creates a new res.users.settings.embedded.action model and returns its id.
func (c *Client) CreateResUsersSettingsEmbeddedAction(rusea *ResUsersSettingsEmbeddedAction) (int64, error) {
	ids, err := c.CreateResUsersSettingsEmbeddedActions([]*ResUsersSettingsEmbeddedAction{rusea})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersSettingsEmbeddedAction creates a new res.users.settings.embedded.action model and returns its id.
func (c *Client) CreateResUsersSettingsEmbeddedActions(ruseas []*ResUsersSettingsEmbeddedAction) ([]int64, error) {
	var vv []interface{}
	for _, v := range ruseas {
		vv = append(vv, v)
	}
	return c.Create(ResUsersSettingsEmbeddedActionModel, vv, nil)
}

// UpdateResUsersSettingsEmbeddedAction updates an existing res.users.settings.embedded.action record.
func (c *Client) UpdateResUsersSettingsEmbeddedAction(rusea *ResUsersSettingsEmbeddedAction) error {
	return c.UpdateResUsersSettingsEmbeddedActions([]int64{rusea.Id.Get()}, rusea)
}

// UpdateResUsersSettingsEmbeddedActions updates existing res.users.settings.embedded.action records.
// All records (represented by ids) will be updated by rusea values.
func (c *Client) UpdateResUsersSettingsEmbeddedActions(ids []int64, rusea *ResUsersSettingsEmbeddedAction) error {
	return c.Update(ResUsersSettingsEmbeddedActionModel, ids, rusea, nil)
}

// DeleteResUsersSettingsEmbeddedAction deletes an existing res.users.settings.embedded.action record.
func (c *Client) DeleteResUsersSettingsEmbeddedAction(id int64) error {
	return c.DeleteResUsersSettingsEmbeddedActions([]int64{id})
}

// DeleteResUsersSettingsEmbeddedActions deletes existing res.users.settings.embedded.action records.
func (c *Client) DeleteResUsersSettingsEmbeddedActions(ids []int64) error {
	return c.Delete(ResUsersSettingsEmbeddedActionModel, ids)
}

// GetResUsersSettingsEmbeddedAction gets res.users.settings.embedded.action existing record.
func (c *Client) GetResUsersSettingsEmbeddedAction(id int64) (*ResUsersSettingsEmbeddedAction, error) {
	ruseas, err := c.GetResUsersSettingsEmbeddedActions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ruseas)[0]), nil
}

// GetResUsersSettingsEmbeddedActions gets res.users.settings.embedded.action existing records.
func (c *Client) GetResUsersSettingsEmbeddedActions(ids []int64) (*ResUsersSettingsEmbeddedActions, error) {
	ruseas := &ResUsersSettingsEmbeddedActions{}
	if err := c.Read(ResUsersSettingsEmbeddedActionModel, ids, nil, ruseas); err != nil {
		return nil, err
	}
	return ruseas, nil
}

// FindResUsersSettingsEmbeddedAction finds res.users.settings.embedded.action record by querying it with criteria.
func (c *Client) FindResUsersSettingsEmbeddedAction(criteria *Criteria) (*ResUsersSettingsEmbeddedAction, error) {
	ruseas := &ResUsersSettingsEmbeddedActions{}
	if err := c.SearchRead(ResUsersSettingsEmbeddedActionModel, criteria, NewOptions().Limit(1), ruseas); err != nil {
		return nil, err
	}
	return &((*ruseas)[0]), nil
}

// FindResUsersSettingsEmbeddedActions finds res.users.settings.embedded.action records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingsEmbeddedActions(criteria *Criteria, options *Options) (*ResUsersSettingsEmbeddedActions, error) {
	ruseas := &ResUsersSettingsEmbeddedActions{}
	if err := c.SearchRead(ResUsersSettingsEmbeddedActionModel, criteria, options, ruseas); err != nil {
		return nil, err
	}
	return ruseas, nil
}

// FindResUsersSettingsEmbeddedActionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersSettingsEmbeddedActionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResUsersSettingsEmbeddedActionModel, criteria, options)
}

// FindResUsersSettingsEmbeddedActionId finds record id by querying it with criteria.
func (c *Client) FindResUsersSettingsEmbeddedActionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersSettingsEmbeddedActionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
