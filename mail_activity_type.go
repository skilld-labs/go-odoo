package odoo

import (
	"fmt"
)

// MailActivityType represents mail.activity.type model.
type MailActivityType struct {
	LastUpdate         *Time      `xmlrpc:"__last_update,omitempty"`
	Active             *Bool      `xmlrpc:"active,omitempty"`
	Category           *Selection `xmlrpc:"category,omitempty"`
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DecorationType     *Selection `xmlrpc:"decoration_type,omitempty"`
	DefaultDescription *String    `xmlrpc:"default_description,omitempty"`
	DefaultNextTypeId  *Many2One  `xmlrpc:"default_next_type_id,omitempty"`
	DefaultUserId      *Many2One  `xmlrpc:"default_user_id,omitempty"`
	DelayCount         *Int       `xmlrpc:"delay_count,omitempty"`
	DelayFrom          *Selection `xmlrpc:"delay_from,omitempty"`
	DelayLabel         *String    `xmlrpc:"delay_label,omitempty"`
	DelayUnit          *Selection `xmlrpc:"delay_unit,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	ForceNext          *Bool      `xmlrpc:"force_next,omitempty"`
	Icon               *String    `xmlrpc:"icon,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	InitialResModelId  *Many2One  `xmlrpc:"initial_res_model_id,omitempty"`
	MailTemplateIds    *Relation  `xmlrpc:"mail_template_ids,omitempty"`
	Name               *String    `xmlrpc:"name,omitempty"`
	NextTypeIds        *Relation  `xmlrpc:"next_type_ids,omitempty"`
	PreviousTypeIds    *Relation  `xmlrpc:"previous_type_ids,omitempty"`
	ResModelChange     *Bool      `xmlrpc:"res_model_change,omitempty"`
	ResModelId         *Many2One  `xmlrpc:"res_model_id,omitempty"`
	Sequence           *Int       `xmlrpc:"sequence,omitempty"`
	Summary            *String    `xmlrpc:"summary,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailActivityTypes represents array of mail.activity.type model.
type MailActivityTypes []MailActivityType

// MailActivityTypeModel is the odoo model name.
const MailActivityTypeModel = "mail.activity.type"

// Many2One convert MailActivityType to *Many2One.
func (mat *MailActivityType) Many2One() *Many2One {
	return NewMany2One(mat.Id.Get(), "")
}

// CreateMailActivityType creates a new mail.activity.type model and returns its id.
func (c *Client) CreateMailActivityType(mat *MailActivityType) (int64, error) {
	return c.Create(MailActivityTypeModel, mat)
}

// UpdateMailActivityType updates an existing mail.activity.type record.
func (c *Client) UpdateMailActivityType(mat *MailActivityType) error {
	return c.UpdateMailActivityTypes([]int64{mat.Id.Get()}, mat)
}

// UpdateMailActivityTypes updates existing mail.activity.type records.
// All records (represented by ids) will be updated by mat values.
func (c *Client) UpdateMailActivityTypes(ids []int64, mat *MailActivityType) error {
	return c.Update(MailActivityTypeModel, ids, mat)
}

// DeleteMailActivityType deletes an existing mail.activity.type record.
func (c *Client) DeleteMailActivityType(id int64) error {
	return c.DeleteMailActivityTypes([]int64{id})
}

// DeleteMailActivityTypes deletes existing mail.activity.type records.
func (c *Client) DeleteMailActivityTypes(ids []int64) error {
	return c.Delete(MailActivityTypeModel, ids)
}

// GetMailActivityType gets mail.activity.type existing record.
func (c *Client) GetMailActivityType(id int64) (*MailActivityType, error) {
	mats, err := c.GetMailActivityTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if mats != nil && len(*mats) > 0 {
		return &((*mats)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.activity.type not found", id)
}

// GetMailActivityTypes gets mail.activity.type existing records.
func (c *Client) GetMailActivityTypes(ids []int64) (*MailActivityTypes, error) {
	mats := &MailActivityTypes{}
	if err := c.Read(MailActivityTypeModel, ids, nil, mats); err != nil {
		return nil, err
	}
	return mats, nil
}

// FindMailActivityType finds mail.activity.type record by querying it with criteria.
func (c *Client) FindMailActivityType(criteria *Criteria) (*MailActivityType, error) {
	mats := &MailActivityTypes{}
	if err := c.SearchRead(MailActivityTypeModel, criteria, NewOptions().Limit(1), mats); err != nil {
		return nil, err
	}
	if mats != nil && len(*mats) > 0 {
		return &((*mats)[0]), nil
	}
	return nil, fmt.Errorf("mail.activity.type was not found")
}

// FindMailActivityTypes finds mail.activity.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityTypes(criteria *Criteria, options *Options) (*MailActivityTypes, error) {
	mats := &MailActivityTypes{}
	if err := c.SearchRead(MailActivityTypeModel, criteria, options, mats); err != nil {
		return nil, err
	}
	return mats, nil
}

// FindMailActivityTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailActivityTypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailActivityTypeId finds record id by querying it with criteria.
func (c *Client) FindMailActivityTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.activity.type was not found")
}
