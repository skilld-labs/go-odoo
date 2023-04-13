package odoo

import (
	"fmt"
)

// MailMassMailingTag represents mail.mass_mailing.tag model.
type MailMassMailingTag struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	Color       *Int      `xmlrpc:"color,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailMassMailingTags represents array of mail.mass_mailing.tag model.
type MailMassMailingTags []MailMassMailingTag

// MailMassMailingTagModel is the odoo model name.
const MailMassMailingTagModel = "mail.mass_mailing.tag"

// Many2One convert MailMassMailingTag to *Many2One.
func (mmt *MailMassMailingTag) Many2One() *Many2One {
	return NewMany2One(mmt.Id.Get(), "")
}

// CreateMailMassMailingTag creates a new mail.mass_mailing.tag model and returns its id.
func (c *Client) CreateMailMassMailingTag(mmt *MailMassMailingTag) (int64, error) {
	ids, err := c.CreateMailMassMailingTags([]*MailMassMailingTag{mmt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMassMailingTag creates a new mail.mass_mailing.tag model and returns its id.
func (c *Client) CreateMailMassMailingTags(mmts []*MailMassMailingTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmts {
		vv = append(vv, v)
	}
	return c.Create(MailMassMailingTagModel, vv)
}

// UpdateMailMassMailingTag updates an existing mail.mass_mailing.tag record.
func (c *Client) UpdateMailMassMailingTag(mmt *MailMassMailingTag) error {
	return c.UpdateMailMassMailingTags([]int64{mmt.Id.Get()}, mmt)
}

// UpdateMailMassMailingTags updates existing mail.mass_mailing.tag records.
// All records (represented by ids) will be updated by mmt values.
func (c *Client) UpdateMailMassMailingTags(ids []int64, mmt *MailMassMailingTag) error {
	return c.Update(MailMassMailingTagModel, ids, mmt)
}

// DeleteMailMassMailingTag deletes an existing mail.mass_mailing.tag record.
func (c *Client) DeleteMailMassMailingTag(id int64) error {
	return c.DeleteMailMassMailingTags([]int64{id})
}

// DeleteMailMassMailingTags deletes existing mail.mass_mailing.tag records.
func (c *Client) DeleteMailMassMailingTags(ids []int64) error {
	return c.Delete(MailMassMailingTagModel, ids)
}

// GetMailMassMailingTag gets mail.mass_mailing.tag existing record.
func (c *Client) GetMailMassMailingTag(id int64) (*MailMassMailingTag, error) {
	mmts, err := c.GetMailMassMailingTags([]int64{id})
	if err != nil {
		return nil, err
	}
	if mmts != nil && len(*mmts) > 0 {
		return &((*mmts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.mass_mailing.tag not found", id)
}

// GetMailMassMailingTags gets mail.mass_mailing.tag existing records.
func (c *Client) GetMailMassMailingTags(ids []int64) (*MailMassMailingTags, error) {
	mmts := &MailMassMailingTags{}
	if err := c.Read(MailMassMailingTagModel, ids, nil, mmts); err != nil {
		return nil, err
	}
	return mmts, nil
}

// FindMailMassMailingTag finds mail.mass_mailing.tag record by querying it with criteria.
func (c *Client) FindMailMassMailingTag(criteria *Criteria) (*MailMassMailingTag, error) {
	mmts := &MailMassMailingTags{}
	if err := c.SearchRead(MailMassMailingTagModel, criteria, NewOptions().Limit(1), mmts); err != nil {
		return nil, err
	}
	if mmts != nil && len(*mmts) > 0 {
		return &((*mmts)[0]), nil
	}
	return nil, fmt.Errorf("mail.mass_mailing.tag was not found with criteria %v", criteria)
}

// FindMailMassMailingTags finds mail.mass_mailing.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingTags(criteria *Criteria, options *Options) (*MailMassMailingTags, error) {
	mmts := &MailMassMailingTags{}
	if err := c.SearchRead(MailMassMailingTagModel, criteria, options, mmts); err != nil {
		return nil, err
	}
	return mmts, nil
}

// FindMailMassMailingTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMassMailingTagModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMassMailingTagId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.mass_mailing.tag was not found with criteria %v and options %v", criteria, options)
}
