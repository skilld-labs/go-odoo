package odoo

import (
	"fmt"
)

// AuthTotpWizard represents auth_totp.wizard model.
type AuthTotpWizard struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Code        *String   `xmlrpc:"code,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Qrcode      *String   `xmlrpc:"qrcode,omitempty"`
	Secret      *String   `xmlrpc:"secret,omitempty"`
	Url         *String   `xmlrpc:"url,omitempty"`
	UserId      *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AuthTotpWizards represents array of auth_totp.wizard model.
type AuthTotpWizards []AuthTotpWizard

// AuthTotpWizardModel is the odoo model name.
const AuthTotpWizardModel = "auth_totp.wizard"

// Many2One convert AuthTotpWizard to *Many2One.
func (aw *AuthTotpWizard) Many2One() *Many2One {
	return NewMany2One(aw.Id.Get(), "")
}

// CreateAuthTotpWizard creates a new auth_totp.wizard model and returns its id.
func (c *Client) CreateAuthTotpWizard(aw *AuthTotpWizard) (int64, error) {
	return c.Create(AuthTotpWizardModel, aw)
}

// UpdateAuthTotpWizard updates an existing auth_totp.wizard record.
func (c *Client) UpdateAuthTotpWizard(aw *AuthTotpWizard) error {
	return c.UpdateAuthTotpWizards([]int64{aw.Id.Get()}, aw)
}

// UpdateAuthTotpWizards updates existing auth_totp.wizard records.
// All records (represented by IDs) will be updated by aw values.
func (c *Client) UpdateAuthTotpWizards(ids []int64, aw *AuthTotpWizard) error {
	return c.Update(AuthTotpWizardModel, ids, aw)
}

// DeleteAuthTotpWizard deletes an existing auth_totp.wizard record.
func (c *Client) DeleteAuthTotpWizard(id int64) error {
	return c.DeleteAuthTotpWizards([]int64{id})
}

// DeleteAuthTotpWizards deletes existing auth_totp.wizard records.
func (c *Client) DeleteAuthTotpWizards(ids []int64) error {
	return c.Delete(AuthTotpWizardModel, ids)
}

// GetAuthTotpWizard gets auth_totp.wizard existing record.
func (c *Client) GetAuthTotpWizard(id int64) (*AuthTotpWizard, error) {
	aws, err := c.GetAuthTotpWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	if aws != nil && len(*aws) > 0 {
		return &((*aws)[0]), nil
	}
	return nil, fmt.Errorf("id %v of auth_totp.wizard not found", id)
}

// GetAuthTotpWizards gets auth_totp.wizard existing records.
func (c *Client) GetAuthTotpWizards(ids []int64) (*AuthTotpWizards, error) {
	aws := &AuthTotpWizards{}
	if err := c.Read(AuthTotpWizardModel, ids, nil, aws); err != nil {
		return nil, err
	}
	return aws, nil
}

// FindAuthTotpWizard finds auth_totp.wizard record by querying it with criteria.
func (c *Client) FindAuthTotpWizard(criteria *Criteria) (*AuthTotpWizard, error) {
	aws := &AuthTotpWizards{}
	if err := c.SearchRead(AuthTotpWizardModel, criteria, NewOptions().Limit(1), aws); err != nil {
		return nil, err
	}
	if aws != nil && len(*aws) > 0 {
		return &((*aws)[0]), nil
	}
	return nil, fmt.Errorf("auth_totp.wizard was not found")
}

// FindAuthTotpWizards finds auth_totp.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthTotpWizards(criteria *Criteria, options *Options) (*AuthTotpWizards, error) {
	aws := &AuthTotpWizards{}
	if err := c.SearchRead(AuthTotpWizardModel, criteria, options, aws); err != nil {
		return nil, err
	}
	return aws, nil
}

// FindAuthTotpWizardIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthTotpWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(AuthTotpWizardModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindAuthTotpWizardId finds record id by querying it with criteria.
func (c *Client) FindAuthTotpWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuthTotpWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("auth_totp.wizard was not found")
}
