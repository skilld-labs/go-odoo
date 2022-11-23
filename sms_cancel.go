package odoo

import (
	"fmt"
)

// SmsCancel represents sms.cancel model.
type SmsCancel struct {
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

// SmsCancels represents array of sms.cancel model.
type SmsCancels []SmsCancel

// SmsCancelModel is the odoo model name.
const SmsCancelModel = "sms.cancel"

// Many2One convert SmsCancel to *Many2One.
func (sc *SmsCancel) Many2One() *Many2One {
	return NewMany2One(sc.Id.Get(), "")
}

// CreateSmsCancel creates a new sms.cancel model and returns its id.
func (c *Client) CreateSmsCancel(sc *SmsCancel) (int64, error) {
	return c.Create(SmsCancelModel, sc)
}

// UpdateSmsCancel updates an existing sms.cancel record.
func (c *Client) UpdateSmsCancel(sc *SmsCancel) error {
	return c.UpdateSmsCancels([]int64{sc.Id.Get()}, sc)
}

// UpdateSmsCancels updates existing sms.cancel records.
// All records (represented by IDs) will be updated by sc values.
func (c *Client) UpdateSmsCancels(ids []int64, sc *SmsCancel) error {
	return c.Update(SmsCancelModel, ids, sc)
}

// DeleteSmsCancel deletes an existing sms.cancel record.
func (c *Client) DeleteSmsCancel(id int64) error {
	return c.DeleteSmsCancels([]int64{id})
}

// DeleteSmsCancels deletes existing sms.cancel records.
func (c *Client) DeleteSmsCancels(ids []int64) error {
	return c.Delete(SmsCancelModel, ids)
}

// GetSmsCancel gets sms.cancel existing record.
func (c *Client) GetSmsCancel(id int64) (*SmsCancel, error) {
	scs, err := c.GetSmsCancels([]int64{id})
	if err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of sms.cancel not found", id)
}

// GetSmsCancels gets sms.cancel existing records.
func (c *Client) GetSmsCancels(ids []int64) (*SmsCancels, error) {
	scs := &SmsCancels{}
	if err := c.Read(SmsCancelModel, ids, nil, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSmsCancel finds sms.cancel record by querying it with criteria.
func (c *Client) FindSmsCancel(criteria *Criteria) (*SmsCancel, error) {
	scs := &SmsCancels{}
	if err := c.SearchRead(SmsCancelModel, criteria, NewOptions().Limit(1), scs); err != nil {
		return nil, err
	}
	if scs != nil && len(*scs) > 0 {
		return &((*scs)[0]), nil
	}
	return nil, fmt.Errorf("sms.cancel was not found")
}

// FindSmsCancels finds sms.cancel records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsCancels(criteria *Criteria, options *Options) (*SmsCancels, error) {
	scs := &SmsCancels{}
	if err := c.SearchRead(SmsCancelModel, criteria, options, scs); err != nil {
		return nil, err
	}
	return scs, nil
}

// FindSmsCancelIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSmsCancelIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SmsCancelModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSmsCancelId finds record id by querying it with criteria.
func (c *Client) FindSmsCancelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SmsCancelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("sms.cancel was not found")
}
