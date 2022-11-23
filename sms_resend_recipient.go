package odoo

import (
	"fmt"
)

// SmsResendRecipient represents sms.resend.recipient model.
type SmsResendRecipient struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	FailureType    *Selection `xmlrpc:"failure_type,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	NotificationId *Many2One  `xmlrpc:"notification_id,omitempty"`
	PartnerId      *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerName    *String    `xmlrpc:"partner_name,omitempty"`
	Resend         *Bool      `xmlrpc:"resend,omitempty"`
	SmsNumber      *String    `xmlrpc:"sms_number,omitempty"`
	SmsResendId    *Many2One  `xmlrpc:"sms_resend_id,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SmsResendRecipients represents array of sms.resend.recipient model.
type SmsResendRecipients []SmsResendRecipient

// SmsResendRecipientModel is the odoo model name.
const SmsResendRecipientModel = "sms.resend.recipient"

// Many2One convert SmsResendRecipient to *Many2One.
func (srr *SmsResendRecipient) Many2One() *Many2One {
	return NewMany2One(srr.Id.Get(), "")
}

// CreateSmsResendRecipient creates a new sms.resend.recipient model and returns its id.
func (c *Client) CreateSmsResendRecipient(srr *SmsResendRecipient) (int64, error) {
	return c.Create(SmsResendRecipientModel, srr)
}

// UpdateSmsResendRecipient updates an existing sms.resend.recipient record.
func (c *Client) UpdateSmsResendRecipient(srr *SmsResendRecipient) error {
	return c.UpdateSmsResendRecipients([]int64{srr.Id.Get()}, srr)
}

// UpdateSmsResendRecipients updates existing sms.resend.recipient records.
// All records (represented by IDs) will be updated by srr values.
func (c *Client) UpdateSmsResendRecipients(ids []int64, srr *SmsResendRecipient) error {
	return c.Update(SmsResendRecipientModel, ids, srr)
}

// DeleteSmsResendRecipient deletes an existing sms.resend.recipient record.
func (c *Client) DeleteSmsResendRecipient(id int64) error {
	return c.DeleteSmsResendRecipients([]int64{id})
}

// DeleteSmsResendRecipients deletes existing sms.resend.recipient records.
func (c *Client) DeleteSmsResendRecipients(ids []int64) error {
	return c.Delete(SmsResendRecipientModel, ids)
}

// GetSmsResendRecipient gets sms.resend.recipient existing record.
func (c *Client) GetSmsResendRecipient(id int64) (*SmsResendRecipient, error) {
	srrs, err := c.GetSmsResendRecipients([]int64{id})
	if err != nil {
		return nil, err
	}
	if srrs != nil && len(*srrs) > 0 {
		return &((*srrs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.resend.recipient not found", id)
}

// GetSmsResendRecipients gets sms.resend.recipient existing records.
func (c *Client) GetSmsResendRecipients(ids []int64) (*SmsResendRecipients, error) {
	srrs := &SmsResendRecipients{}
	if err := c.Read(SmsResendRecipientModel, ids, nil, srrs); err != nil {
		return nil, err
	}
	return srrs, nil
}

// FindSmsResendRecipient finds sms.resend.recipient record by querying it with criteria.
func (c *Client) FindSmsResendRecipient(criteria *Criteria) (*SmsResendRecipient, error) {
	srrs := &SmsResendRecipients{}
	if err := c.SearchRead(SmsResendRecipientModel, criteria, NewOptions().Limit(1), srrs); err != nil {
		return nil, err
	}
	if srrs != nil && len(*srrs) > 0 {
		return &((*srrs)[0]), nil
	}
	return nil, fmt.Errorf("sms.resend.recipient was not found")
}

// FindSmsResendRecipients finds sms.resend.recipient records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsResendRecipients(criteria *Criteria, options *Options) (*SmsResendRecipients, error) {
	srrs := &SmsResendRecipients{}
	if err := c.SearchRead(SmsResendRecipientModel, criteria, options, srrs); err != nil {
		return nil, err
	}
	return srrs, nil
}

// FindSmsResendRecipientIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsResendRecipientIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsResendRecipientModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsResendRecipientId finds record id by querying it with criteria.
func (c *Client) FindSmsResendRecipientId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsResendRecipientModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.resend.recipient was not found")
}
