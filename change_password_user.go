package odoo

// ChangePasswordUser represents change.password.user model.
type ChangePasswordUser struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	NewPasswd   *String   `xmlrpc:"new_passwd,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	UserLogin   *String   `xmlrpc:"user_login,omitempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ChangePasswordUsers represents array of change.password.user model.
type ChangePasswordUsers []ChangePasswordUser

// ChangePasswordUserModel is the odoo model name.
const ChangePasswordUserModel = "change.password.user"

// Many2One convert ChangePasswordUser to *Many2One.
func (cpu *ChangePasswordUser) Many2One() *Many2One {
	return NewMany2One(cpu.Id.Get(), "")
}

// CreateChangePasswordUser creates a new change.password.user model and returns its id.
func (c *Client) CreateChangePasswordUser(cpu *ChangePasswordUser) (int64, error) {
	ids, err := c.CreateChangePasswordUsers([]*ChangePasswordUser{cpu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChangePasswordUser creates a new change.password.user model and returns its id.
func (c *Client) CreateChangePasswordUsers(cpus []*ChangePasswordUser) ([]int64, error) {
	var vv []interface{}
	for _, v := range cpus {
		vv = append(vv, v)
	}
	return c.Create(ChangePasswordUserModel, vv, nil)
}

// UpdateChangePasswordUser updates an existing change.password.user record.
func (c *Client) UpdateChangePasswordUser(cpu *ChangePasswordUser) error {
	return c.UpdateChangePasswordUsers([]int64{cpu.Id.Get()}, cpu)
}

// UpdateChangePasswordUsers updates existing change.password.user records.
// All records (represented by ids) will be updated by cpu values.
func (c *Client) UpdateChangePasswordUsers(ids []int64, cpu *ChangePasswordUser) error {
	return c.Update(ChangePasswordUserModel, ids, cpu, nil)
}

// DeleteChangePasswordUser deletes an existing change.password.user record.
func (c *Client) DeleteChangePasswordUser(id int64) error {
	return c.DeleteChangePasswordUsers([]int64{id})
}

// DeleteChangePasswordUsers deletes existing change.password.user records.
func (c *Client) DeleteChangePasswordUsers(ids []int64) error {
	return c.Delete(ChangePasswordUserModel, ids)
}

// GetChangePasswordUser gets change.password.user existing record.
func (c *Client) GetChangePasswordUser(id int64) (*ChangePasswordUser, error) {
	cpus, err := c.GetChangePasswordUsers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cpus)[0]), nil
}

// GetChangePasswordUsers gets change.password.user existing records.
func (c *Client) GetChangePasswordUsers(ids []int64) (*ChangePasswordUsers, error) {
	cpus := &ChangePasswordUsers{}
	if err := c.Read(ChangePasswordUserModel, ids, nil, cpus); err != nil {
		return nil, err
	}
	return cpus, nil
}

// FindChangePasswordUser finds change.password.user record by querying it with criteria.
func (c *Client) FindChangePasswordUser(criteria *Criteria) (*ChangePasswordUser, error) {
	cpus := &ChangePasswordUsers{}
	if err := c.SearchRead(ChangePasswordUserModel, criteria, NewOptions().Limit(1), cpus); err != nil {
		return nil, err
	}
	return &((*cpus)[0]), nil
}

// FindChangePasswordUsers finds change.password.user records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChangePasswordUsers(criteria *Criteria, options *Options) (*ChangePasswordUsers, error) {
	cpus := &ChangePasswordUsers{}
	if err := c.SearchRead(ChangePasswordUserModel, criteria, options, cpus); err != nil {
		return nil, err
	}
	return cpus, nil
}

// FindChangePasswordUserIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChangePasswordUserIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ChangePasswordUserModel, criteria, options)
}

// FindChangePasswordUserId finds record id by querying it with criteria.
func (c *Client) FindChangePasswordUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChangePasswordUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
