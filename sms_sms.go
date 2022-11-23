package odoo

import (
	"fmt"
)

// SmsSms represents sms.sms model.
type SmsSms struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omitempty"`
	Body          *String    `xmlrpc:"body,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	ErrorCode     *Selection `xmlrpc:"error_code,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	MailMessageId *Many2One  `xmlrpc:"mail_message_id,omitempty"`
	Number        *String    `xmlrpc:"number,omitempty"`
	PartnerId     *Many2One  `xmlrpc:"partner_id,omitempty"`
	State         *Selection `xmlrpc:"state,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// SmsSmss represents array of sms.sms model.
type SmsSmss []SmsSms

// SmsSmsModel is the odoo model name.
const SmsSmsModel = "sms.sms"

// Many2One convert SmsSms to *Many2One.
func (ss *SmsSms) Many2One() *Many2One {
	return NewMany2One(ss.Id.Get(), "")
}

// CreateSmsSms creates a new sms.sms model and returns its id.
func (c *Client) CreateSmsSms(ss *SmsSms) (int64, error) {
	return c.Create(SmsSmsModel, ss)
}

// UpdateSmsSms updates an existing sms.sms record.
func (c *Client) UpdateSmsSms(ss *SmsSms) error {
	return c.UpdateSmsSmss([]int64{ss.Id.Get()}, ss)
}

// UpdateSmsSmss updates existing sms.sms records.
// All records (represented by IDs) will be updated by ss values.
func (c *Client) UpdateSmsSmss(ids []int64, ss *SmsSms) error {
	return c.Update(SmsSmsModel, ids, ss)
}

// DeleteSmsSms deletes an existing sms.sms record.
func (c *Client) DeleteSmsSms(id int64) error {
	return c.DeleteSmsSmss([]int64{id})
}

// DeleteSmsSmss deletes existing sms.sms records.
func (c *Client) DeleteSmsSmss(ids []int64) error {
	return c.Delete(SmsSmsModel, ids)
}

// GetSmsSms gets sms.sms existing record.
func (c *Client) GetSmsSms(id int64) (*SmsSms, error) {
	sss, err := c.GetSmsSmss([]int64{id})
	if err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.sms not found", id)
}

// GetSmsSmss gets sms.sms existing records.
func (c *Client) GetSmsSmss(ids []int64) (*SmsSmss, error) {
	sss := &SmsSmss{}
	if err := c.Read(SmsSmsModel, ids, nil, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSmsSms finds sms.sms record by querying it with criteria.
func (c *Client) FindSmsSms(criteria *Criteria) (*SmsSms, error) {
	sss := &SmsSmss{}
	if err := c.SearchRead(SmsSmsModel, criteria, NewOptions().Limit(1), sss); err != nil {
		return nil, err
	}
	if sss != nil && len(*sss) > 0 {
		return &((*sss)[0]), nil
	}
	return nil, fmt.Errorf("sms.sms was not found")
}

// FindSmsSmss finds sms.sms records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsSmss(criteria *Criteria, options *Options) (*SmsSmss, error) {
	sss := &SmsSmss{}
	if err := c.SearchRead(SmsSmsModel, criteria, options, sss); err != nil {
		return nil, err
	}
	return sss, nil
}

// FindSmsSmsIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsSmsIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsSmsModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsSmsId finds record id by querying it with criteria.
func (c *Client) FindSmsSmsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsSmsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.sms was not found")
}
