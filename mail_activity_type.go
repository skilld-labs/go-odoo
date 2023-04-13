package odoo

import (
	"fmt"
)

// MailActivityType represents mail.activity.type model.
type MailActivityType struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	Category        *Selection `xmlrpc:"category,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	Days            *Int       `xmlrpc:"days,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	Icon            *String    `xmlrpc:"icon,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	NextTypeIds     *Relation  `xmlrpc:"next_type_ids,omptempty"`
	PreviousTypeIds *Relation  `xmlrpc:"previous_type_ids,omptempty"`
	ResModelId      *Many2One  `xmlrpc:"res_model_id,omptempty"`
	Sequence        *Int       `xmlrpc:"sequence,omptempty"`
	Summary         *String    `xmlrpc:"summary,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	ids, err := c.Create(MailActivityTypeModel, []interface{}{mat})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailActivityType creates a new mail.activity.type model and returns its id.
func (c *Client) CreateMailActivityTypes(mats []*MailActivityType) ([]int64, error) {
	var vv []interface{}
	for _, v := range mats {
		vv = append(vv, v)
	}
	return c.Create(MailActivityTypeModel, vv)
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
	return nil, fmt.Errorf("mail.activity.type was not found with criteria %v", criteria)
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
	return -1, fmt.Errorf("mail.activity.type was not found with criteria %v and options %v", criteria, options)
}
