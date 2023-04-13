package odoo

import (
	"fmt"
)

// ChangePasswordUser represents change.password.user model.
type ChangePasswordUser struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	NewPasswd   *String   `xmlrpc:"new_passwd,omptempty"`
	UserId      *Many2One `xmlrpc:"user_id,omptempty"`
	UserLogin   *String   `xmlrpc:"user_login,omptempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.Create(ChangePasswordUserModel, []interface{}{cpu})
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
	return c.Create(ChangePasswordUserModel, vv)
}

// UpdateChangePasswordUser updates an existing change.password.user record.
func (c *Client) UpdateChangePasswordUser(cpu *ChangePasswordUser) error {
	return c.UpdateChangePasswordUsers([]int64{cpu.Id.Get()}, cpu)
}

// UpdateChangePasswordUsers updates existing change.password.user records.
// All records (represented by ids) will be updated by cpu values.
func (c *Client) UpdateChangePasswordUsers(ids []int64, cpu *ChangePasswordUser) error {
	return c.Update(ChangePasswordUserModel, ids, cpu)
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
	if cpus != nil && len(*cpus) > 0 {
		return &((*cpus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of change.password.user not found", id)
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
	if cpus != nil && len(*cpus) > 0 {
		return &((*cpus)[0]), nil
	}
	return nil, fmt.Errorf("change.password.user was not found with criteria %v", criteria)
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
	ids, err := c.Search(ChangePasswordUserModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindChangePasswordUserId finds record id by querying it with criteria.
func (c *Client) FindChangePasswordUserId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChangePasswordUserModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("change.password.user was not found with criteria %v and options %v", criteria, options)
}
