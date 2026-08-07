package odoo

// ServerActionHistoryWizard represents server.action.history.wizard model.
type ServerActionHistoryWizard struct {
	ActionId    *Many2One `xmlrpc:"action_id,omitempty"`
	CodeDiff    *String   `xmlrpc:"code_diff,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrentCode *String   `xmlrpc:"current_code,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Revision    *Many2One `xmlrpc:"revision,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ServerActionHistoryWizards represents array of server.action.history.wizard model.
type ServerActionHistoryWizards []ServerActionHistoryWizard

// ServerActionHistoryWizardModel is the odoo model name.
const ServerActionHistoryWizardModel = "server.action.history.wizard"

// Many2One convert ServerActionHistoryWizard to *Many2One.
func (sahw *ServerActionHistoryWizard) Many2One() *Many2One {
	return NewMany2One(sahw.Id.Get(), "")
}

// CreateServerActionHistoryWizard creates a new server.action.history.wizard model and returns its id.
func (c *Client) CreateServerActionHistoryWizard(sahw *ServerActionHistoryWizard) (int64, error) {
	ids, err := c.CreateServerActionHistoryWizards([]*ServerActionHistoryWizard{sahw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateServerActionHistoryWizard creates a new server.action.history.wizard model and returns its id.
func (c *Client) CreateServerActionHistoryWizards(sahws []*ServerActionHistoryWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range sahws {
		vv = append(vv, v)
	}
	return c.Create(ServerActionHistoryWizardModel, vv, nil)
}

// UpdateServerActionHistoryWizard updates an existing server.action.history.wizard record.
func (c *Client) UpdateServerActionHistoryWizard(sahw *ServerActionHistoryWizard) error {
	return c.UpdateServerActionHistoryWizards([]int64{sahw.Id.Get()}, sahw)
}

// UpdateServerActionHistoryWizards updates existing server.action.history.wizard records.
// All records (represented by ids) will be updated by sahw values.
func (c *Client) UpdateServerActionHistoryWizards(ids []int64, sahw *ServerActionHistoryWizard) error {
	return c.Update(ServerActionHistoryWizardModel, ids, sahw, nil)
}

// DeleteServerActionHistoryWizard deletes an existing server.action.history.wizard record.
func (c *Client) DeleteServerActionHistoryWizard(id int64) error {
	return c.DeleteServerActionHistoryWizards([]int64{id})
}

// DeleteServerActionHistoryWizards deletes existing server.action.history.wizard records.
func (c *Client) DeleteServerActionHistoryWizards(ids []int64) error {
	return c.Delete(ServerActionHistoryWizardModel, ids)
}

// GetServerActionHistoryWizard gets server.action.history.wizard existing record.
func (c *Client) GetServerActionHistoryWizard(id int64) (*ServerActionHistoryWizard, error) {
	sahws, err := c.GetServerActionHistoryWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sahws)[0]), nil
}

// GetServerActionHistoryWizards gets server.action.history.wizard existing records.
func (c *Client) GetServerActionHistoryWizards(ids []int64) (*ServerActionHistoryWizards, error) {
	sahws := &ServerActionHistoryWizards{}
	if err := c.Read(ServerActionHistoryWizardModel, ids, nil, sahws); err != nil {
		return nil, err
	}
	return sahws, nil
}

// FindServerActionHistoryWizard finds server.action.history.wizard record by querying it with criteria.
func (c *Client) FindServerActionHistoryWizard(criteria *Criteria) (*ServerActionHistoryWizard, error) {
	sahws := &ServerActionHistoryWizards{}
	if err := c.SearchRead(ServerActionHistoryWizardModel, criteria, NewOptions().Limit(1), sahws); err != nil {
		return nil, err
	}
	return &((*sahws)[0]), nil
}

// FindServerActionHistoryWizards finds server.action.history.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindServerActionHistoryWizards(criteria *Criteria, options *Options) (*ServerActionHistoryWizards, error) {
	sahws := &ServerActionHistoryWizards{}
	if err := c.SearchRead(ServerActionHistoryWizardModel, criteria, options, sahws); err != nil {
		return nil, err
	}
	return sahws, nil
}

// FindServerActionHistoryWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindServerActionHistoryWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ServerActionHistoryWizardModel, criteria, options)
}

// FindServerActionHistoryWizardId finds record id by querying it with criteria.
func (c *Client) FindServerActionHistoryWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ServerActionHistoryWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
