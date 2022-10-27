package odoo

import (
	"fmt"
)

// MailShortcode represents mail.shortcode model.
type MailShortcode struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	Description  *String   `xmlrpc:"description,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	MessageIds   *Many2One `xmlrpc:"message_ids,omitempty"`
	Source       *String   `xmlrpc:"source,omitempty"`
	Substitution *String   `xmlrpc:"substitution,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailShortcodes represents array of mail.shortcode model.
type MailShortcodes []MailShortcode

// MailShortcodeModel is the odoo model name.
const MailShortcodeModel = "mail.shortcode"

// Many2One convert MailShortcode to *Many2One.
func (ms *MailShortcode) Many2One() *Many2One {
	return NewMany2One(ms.Id.Get(), "")
}

// CreateMailShortcode creates a new mail.shortcode model and returns its id.
func (c *Client) CreateMailShortcode(ms *MailShortcode) (int64, error) {
	return c.Create(MailShortcodeModel, ms)
}

// UpdateMailShortcode updates an existing mail.shortcode record.
func (c *Client) UpdateMailShortcode(ms *MailShortcode) error {
	return c.UpdateMailShortcodes([]int64{ms.Id.Get()}, ms)
}

// UpdateMailShortcodes updates existing mail.shortcode records.
// All records (represented by ids) will be updated by ms values.
func (c *Client) UpdateMailShortcodes(ids []int64, ms *MailShortcode) error {
	return c.Update(MailShortcodeModel, ids, ms)
}

// DeleteMailShortcode deletes an existing mail.shortcode record.
func (c *Client) DeleteMailShortcode(id int64) error {
	return c.DeleteMailShortcodes([]int64{id})
}

// DeleteMailShortcodes deletes existing mail.shortcode records.
func (c *Client) DeleteMailShortcodes(ids []int64) error {
	return c.Delete(MailShortcodeModel, ids)
}

// GetMailShortcode gets mail.shortcode existing record.
func (c *Client) GetMailShortcode(id int64) (*MailShortcode, error) {
	mss, err := c.GetMailShortcodes([]int64{id})
	if err != nil {
		return nil, err
	}
	if mss != nil && len(*mss) > 0 {
		return &((*mss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.shortcode not found", id)
}

// GetMailShortcodes gets mail.shortcode existing records.
func (c *Client) GetMailShortcodes(ids []int64) (*MailShortcodes, error) {
	mss := &MailShortcodes{}
	if err := c.Read(MailShortcodeModel, ids, nil, mss); err != nil {
		return nil, err
	}
	return mss, nil
}

// FindMailShortcode finds mail.shortcode record by querying it with criteria.
func (c *Client) FindMailShortcode(criteria *Criteria) (*MailShortcode, error) {
	mss := &MailShortcodes{}
	if err := c.SearchRead(MailShortcodeModel, criteria, NewOptions().Limit(1), mss); err != nil {
		return nil, err
	}
	if mss != nil && len(*mss) > 0 {
		return &((*mss)[0]), nil
	}
	return nil, fmt.Errorf("mail.shortcode was not found")
}

// FindMailShortcodes finds mail.shortcode records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailShortcodes(criteria *Criteria, options *Options) (*MailShortcodes, error) {
	mss := &MailShortcodes{}
	if err := c.SearchRead(MailShortcodeModel, criteria, options, mss); err != nil {
		return nil, err
	}
	return mss, nil
}

// FindMailShortcodeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailShortcodeIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailShortcodeModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailShortcodeId finds record id by querying it with criteria.
func (c *Client) FindMailShortcodeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailShortcodeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.shortcode was not found")
}
