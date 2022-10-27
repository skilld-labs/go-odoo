package odoo

import (
	"fmt"
)

// MailResendCancel represents mail.resend.cancel model.
type MailResendCancel struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	HelpMessage *String   `xmlrpc:"help_message,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Model       *String   `xmlrpc:"model,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailResendCancels represents array of mail.resend.cancel model.
type MailResendCancels []MailResendCancel

// MailResendCancelModel is the odoo model name.
const MailResendCancelModel = "mail.resend.cancel"

// Many2One convert MailResendCancel to *Many2One.
func (mrc *MailResendCancel) Many2One() *Many2One {
	return NewMany2One(mrc.Id.Get(), "")
}

// CreateMailResendCancel creates a new mail.resend.cancel model and returns its id.
func (c *Client) CreateMailResendCancel(mrc *MailResendCancel) (int64, error) {
	return c.Create(MailResendCancelModel, mrc)
}

// UpdateMailResendCancel updates an existing mail.resend.cancel record.
func (c *Client) UpdateMailResendCancel(mrc *MailResendCancel) error {
	return c.UpdateMailResendCancels([]int64{mrc.Id.Get()}, mrc)
}

// UpdateMailResendCancels updates existing mail.resend.cancel records.
// All records (represented by ids) will be updated by mrc values.
func (c *Client) UpdateMailResendCancels(ids []int64, mrc *MailResendCancel) error {
	return c.Update(MailResendCancelModel, ids, mrc)
}

// DeleteMailResendCancel deletes an existing mail.resend.cancel record.
func (c *Client) DeleteMailResendCancel(id int64) error {
	return c.DeleteMailResendCancels([]int64{id})
}

// DeleteMailResendCancels deletes existing mail.resend.cancel records.
func (c *Client) DeleteMailResendCancels(ids []int64) error {
	return c.Delete(MailResendCancelModel, ids)
}

// GetMailResendCancel gets mail.resend.cancel existing record.
func (c *Client) GetMailResendCancel(id int64) (*MailResendCancel, error) {
	mrcs, err := c.GetMailResendCancels([]int64{id})
	if err != nil {
		return nil, err
	}
	if mrcs != nil && len(*mrcs) > 0 {
		return &((*mrcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.resend.cancel not found", id)
}

// GetMailResendCancels gets mail.resend.cancel existing records.
func (c *Client) GetMailResendCancels(ids []int64) (*MailResendCancels, error) {
	mrcs := &MailResendCancels{}
	if err := c.Read(MailResendCancelModel, ids, nil, mrcs); err != nil {
		return nil, err
	}
	return mrcs, nil
}

// FindMailResendCancel finds mail.resend.cancel record by querying it with criteria.
func (c *Client) FindMailResendCancel(criteria *Criteria) (*MailResendCancel, error) {
	mrcs := &MailResendCancels{}
	if err := c.SearchRead(MailResendCancelModel, criteria, NewOptions().Limit(1), mrcs); err != nil {
		return nil, err
	}
	if mrcs != nil && len(*mrcs) > 0 {
		return &((*mrcs)[0]), nil
	}
	return nil, fmt.Errorf("mail.resend.cancel was not found")
}

// FindMailResendCancels finds mail.resend.cancel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailResendCancels(criteria *Criteria, options *Options) (*MailResendCancels, error) {
	mrcs := &MailResendCancels{}
	if err := c.SearchRead(MailResendCancelModel, criteria, options, mrcs); err != nil {
		return nil, err
	}
	return mrcs, nil
}

// FindMailResendCancelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailResendCancelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailResendCancelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailResendCancelId finds record id by querying it with criteria.
func (c *Client) FindMailResendCancelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailResendCancelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.resend.cancel was not found")
}
