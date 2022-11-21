package odoo

import (
	"fmt"
)

// MailAliasMixin represents mail.alias.mixin model.
type MailAliasMixin struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	AliasContact        *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasDefaults       *String    `xmlrpc:"alias_defaults,omptempty"`
	AliasDomain         *String    `xmlrpc:"alias_domain,omptempty"`
	AliasForceThreadId  *Int       `xmlrpc:"alias_force_thread_id,omptempty"`
	AliasId             *Many2One  `xmlrpc:"alias_id,omptempty"`
	AliasModelId        *Many2One  `xmlrpc:"alias_model_id,omptempty"`
	AliasName           *String    `xmlrpc:"alias_name,omptempty"`
	AliasParentModelId  *Many2One  `xmlrpc:"alias_parent_model_id,omptempty"`
	AliasParentThreadId *Int       `xmlrpc:"alias_parent_thread_id,omptempty"`
	AliasUserId         *Many2One  `xmlrpc:"alias_user_id,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailAliasMixins represents array of mail.alias.mixin model.
type MailAliasMixins []MailAliasMixin

// MailAliasMixinModel is the odoo model name.
const MailAliasMixinModel = "mail.alias.mixin"

// Many2One convert MailAliasMixin to *Many2One.
func (mam *MailAliasMixin) Many2One() *Many2One {
	return NewMany2One(mam.Id.Get(), "")
}

// CreateMailAliasMixin creates a new mail.alias.mixin model and returns its id.
func (c *Client) CreateMailAliasMixin(mam *MailAliasMixin) (int64, error) {
	return c.Create(MailAliasMixinModel, mam)
}

// UpdateMailAliasMixin updates an existing mail.alias.mixin record.
func (c *Client) UpdateMailAliasMixin(mam *MailAliasMixin) error {
	return c.UpdateMailAliasMixins([]int64{mam.Id.Get()}, mam)
}

// UpdateMailAliasMixins updates existing mail.alias.mixin records.
// All records (represented by ids) will be updated by mam values.
func (c *Client) UpdateMailAliasMixins(ids []int64, mam *MailAliasMixin) error {
	return c.Update(MailAliasMixinModel, ids, mam)
}

// DeleteMailAliasMixin deletes an existing mail.alias.mixin record.
func (c *Client) DeleteMailAliasMixin(id int64) error {
	return c.DeleteMailAliasMixins([]int64{id})
}

// DeleteMailAliasMixins deletes existing mail.alias.mixin records.
func (c *Client) DeleteMailAliasMixins(ids []int64) error {
	return c.Delete(MailAliasMixinModel, ids)
}

// GetMailAliasMixin gets mail.alias.mixin existing record.
func (c *Client) GetMailAliasMixin(id int64) (*MailAliasMixin, error) {
	mams, err := c.GetMailAliasMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	if mams != nil && len(*mams) > 0 {
		return &((*mams)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.alias.mixin not found", id)
}

// GetMailAliasMixins gets mail.alias.mixin existing records.
func (c *Client) GetMailAliasMixins(ids []int64) (*MailAliasMixins, error) {
	mams := &MailAliasMixins{}
	if err := c.Read(MailAliasMixinModel, ids, nil, mams); err != nil {
		return nil, err
	}
	return mams, nil
}

// FindMailAliasMixin finds mail.alias.mixin record by querying it with criteria.
func (c *Client) FindMailAliasMixin(criteria *Criteria) (*MailAliasMixin, error) {
	mams := &MailAliasMixins{}
	if err := c.SearchRead(MailAliasMixinModel, criteria, NewOptions().Limit(1), mams); err != nil {
		return nil, err
	}
	if mams != nil && len(*mams) > 0 {
		return &((*mams)[0]), nil
	}
	return nil, fmt.Errorf("no mail.alias.mixin was found with criteria %v", criteria)
}

// FindMailAliasMixins finds mail.alias.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAliasMixins(criteria *Criteria, options *Options) (*MailAliasMixins, error) {
	mams := &MailAliasMixins{}
	if err := c.SearchRead(MailAliasMixinModel, criteria, options, mams); err != nil {
		return nil, err
	}
	return mams, nil
}

// FindMailAliasMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailAliasMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailAliasMixinModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailAliasMixinId finds record id by querying it with criteria.
func (c *Client) FindMailAliasMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailAliasMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no mail.alias.mixin was found with criteria %v and options %v", criteria, options)
}
