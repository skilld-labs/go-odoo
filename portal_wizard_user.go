package odoo

// PortalWizardUser represents portal.wizard.user model.
type PortalWizardUser struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Email       *String   `xmlrpc:"email,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	InPortal    *Bool     `xmlrpc:"in_portal,omitempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PortalWizardUsers represents array of portal.wizard.user model.
type PortalWizardUsers []PortalWizardUser

// PortalWizardUserModel is the odoo model name.
const PortalWizardUserModel = "portal.wizard.user"

// Many2One convert PortalWizardUser to *Many2One.
func (pwu *PortalWizardUser) Many2One() *Many2One {
	return NewMany2One(pwu.Id.Get(), "")
}

// CreatePortalWizardUser creates a new portal.wizard.user model and returns its id.
func (c *Client) CreatePortalWizardUser(pwu *PortalWizardUser) (int64, error) {
	ids, err := c.CreatePortalWizardUsers([]*PortalWizardUser{pwu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePortalWizardUsers creates a new portal.wizard.user model and returns its id.
func (c *Client) CreatePortalWizardUsers(pwus []*PortalWizardUser) ([]int64, error) {
	var vv []interface{}
	for _, v := range pwus {
		vv = append(vv, v)
	}
	return c.Create(PortalWizardUserModel, vv, nil)
}

// UpdatePortalWizardUser updates an existing portal.wizard.user record.
func (c *Client) UpdatePortalWizardUser(pwu *PortalWizardUser) error {
	return c.UpdatePortalWizardUsers([]int64{pwu.Id.Get()}, pwu)
}

// UpdatePortalWizardUsers updates existing portal.wizard.user records.
// All records (represented by ids) will be updated by pwu values.
func (c *Client) UpdatePortalWizardUsers(ids []int64, pwu *PortalWizardUser) error {
	return c.Update(PortalWizardUserModel, ids, pwu, nil)
}

// DeletePortalWizardUser deletes an existing portal.wizard.user record.
func (c *Client) DeletePortalWizardUser(id int64) error {
	return c.DeletePortalWizardUsers([]int64{id})
}

// DeletePortalWizardUsers deletes existing portal.wizard.user records.
func (c *Client) DeletePortalWizardUsers(ids []int64) error {
	return c.Delete(PortalWizardUserModel, ids)
}

// GetPortalWizardUser gets portal.wizard.user existing record.
func (c *Client) GetPortalWizardUser(id int64) (*PortalWizardUser, error) {
	pwus, err := c.GetPortalWizardUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pwus)[0]), nil
}

// GetPortalWizardUsers gets portal.wizard.user existing records.
func (c *Client) GetPortalWizardUsers(ids []int64) (*PortalWizardUsers, error) {
	pwus := &PortalWizardUsers{}
	if err := c.Read(PortalWizardUserModel, ids, nil, pwus); err != nil {
		return nil, err
	}
	return pwus, nil
}

// FindPortalWizardUser finds portal.wizard.user record by querying it with criteria.
func (c *Client) FindPortalWizardUser(criteria *Criteria) (*PortalWizardUser, error) {
	pwus := &PortalWizardUsers{}
	if err := c.SearchRead(PortalWizardUserModel, criteria, NewOptions().Limit(1), pwus); err != nil {
		return nil, err
	}
	return &((*pwus)[0]), nil
}

// FindPortalWizardUsers finds portal.wizard.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPortalWizardUsers(criteria *Criteria, options *Options) (*PortalWizardUsers, error) {
	pwus := &PortalWizardUsers{}
	if err := c.SearchRead(PortalWizardUserModel, criteria, options, pwus); err != nil {
		return nil, err
	}
	return pwus, nil
}

// FindPortalWizardUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPortalWizardUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PortalWizardUserModel, criteria, options)
}

// FindPortalWizardUserId finds record id by querying it with criteria.
func (c *Client) FindPortalWizardUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PortalWizardUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
