package odoo

import (
	"fmt"
)

// MailMessageSubtype represents mail.message.subtype model.
type MailMessageSubtype struct {
	LastUpdate    *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omptempty"`
	Default       *Bool     `xmlrpc:"default,omptempty"`
	Description   *String   `xmlrpc:"description,omptempty"`
	DisplayName   *String   `xmlrpc:"display_name,omptempty"`
	Hidden        *Bool     `xmlrpc:"hidden,omptempty"`
	Id            *Int      `xmlrpc:"id,omptempty"`
	Internal      *Bool     `xmlrpc:"internal,omptempty"`
	Name          *String   `xmlrpc:"name,omptempty"`
	ParentId      *Many2One `xmlrpc:"parent_id,omptempty"`
	RelationField *String   `xmlrpc:"relation_field,omptempty"`
	ResModel      *String   `xmlrpc:"res_model,omptempty"`
	Sequence      *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailMessageSubtypes represents array of mail.message.subtype model.
type MailMessageSubtypes []MailMessageSubtype

// MailMessageSubtypeModel is the odoo model name.
const MailMessageSubtypeModel = "mail.message.subtype"

// Many2One convert MailMessageSubtype to *Many2One.
func (mms *MailMessageSubtype) Many2One() *Many2One {
	return NewMany2One(mms.Id.Get(), "")
}

// CreateMailMessageSubtype creates a new mail.message.subtype model and returns its id.
func (c *Client) CreateMailMessageSubtype(mms *MailMessageSubtype) (int64, error) {
	return c.Create(MailMessageSubtypeModel, mms)
}

// UpdateMailMessageSubtype updates an existing mail.message.subtype record.
func (c *Client) UpdateMailMessageSubtype(mms *MailMessageSubtype) error {
	return c.UpdateMailMessageSubtypes([]int64{mms.Id.Get()}, mms)
}

// UpdateMailMessageSubtypes updates existing mail.message.subtype records.
// All records (represented by ids) will be updated by mms values.
func (c *Client) UpdateMailMessageSubtypes(ids []int64, mms *MailMessageSubtype) error {
	return c.Update(MailMessageSubtypeModel, ids, mms)
}

// DeleteMailMessageSubtype deletes an existing mail.message.subtype record.
func (c *Client) DeleteMailMessageSubtype(id int64) error {
	return c.DeleteMailMessageSubtypes([]int64{id})
}

// DeleteMailMessageSubtypes deletes existing mail.message.subtype records.
func (c *Client) DeleteMailMessageSubtypes(ids []int64) error {
	return c.Delete(MailMessageSubtypeModel, ids)
}

// GetMailMessageSubtype gets mail.message.subtype existing record.
func (c *Client) GetMailMessageSubtype(id int64) (*MailMessageSubtype, error) {
	mmss, err := c.GetMailMessageSubtypes([]int64{id})
	if err != nil {
		return nil, err
	}
	if mmss != nil && len(*mmss) > 0 {
		return &((*mmss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.message.subtype not found", id)
}

// GetMailMessageSubtypes gets mail.message.subtype existing records.
func (c *Client) GetMailMessageSubtypes(ids []int64) (*MailMessageSubtypes, error) {
	mmss := &MailMessageSubtypes{}
	if err := c.Read(MailMessageSubtypeModel, ids, nil, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMessageSubtype finds mail.message.subtype record by querying it with criteria.
func (c *Client) FindMailMessageSubtype(criteria *Criteria) (*MailMessageSubtype, error) {
	mmss := &MailMessageSubtypes{}
	if err := c.SearchRead(MailMessageSubtypeModel, criteria, NewOptions().Limit(1), mmss); err != nil {
		return nil, err
	}
	if mmss != nil && len(*mmss) > 0 {
		return &((*mmss)[0]), nil
	}
	return nil, fmt.Errorf("mail.message.subtype was not found with criteria %v", criteria)
}

// FindMailMessageSubtypes finds mail.message.subtype records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageSubtypes(criteria *Criteria, options *Options) (*MailMessageSubtypes, error) {
	mmss := &MailMessageSubtypes{}
	if err := c.SearchRead(MailMessageSubtypeModel, criteria, options, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMessageSubtypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageSubtypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMessageSubtypeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMessageSubtypeId finds record id by querying it with criteria.
func (c *Client) FindMailMessageSubtypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageSubtypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.message.subtype was not found with criteria %v and options %v", criteria, options)
}
