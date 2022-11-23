package odoo

import (
	"fmt"
)

// SnailmailLetter represents snailmail.letter model.
type SnailmailLetter struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omitempty"`
	AttachmentDatas *String    `xmlrpc:"attachment_datas,omitempty"`
	AttachmentFname *String    `xmlrpc:"attachment_fname,omitempty"`
	AttachmentId    *Many2One  `xmlrpc:"attachment_id,omitempty"`
	City            *String    `xmlrpc:"city,omitempty"`
	Color           *Bool      `xmlrpc:"color,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId       *Many2One  `xmlrpc:"country_id,omitempty"`
	Cover           *Bool      `xmlrpc:"cover,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Duplex          *Bool      `xmlrpc:"duplex,omitempty"`
	ErrorCode       *Selection `xmlrpc:"error_code,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	InfoMsg         *String    `xmlrpc:"info_msg,omitempty"`
	MessageId       *Many2One  `xmlrpc:"message_id,omitempty"`
	Model           *String    `xmlrpc:"model,omitempty"`
	NotificationIds *Relation  `xmlrpc:"notification_ids,omitempty"`
	PartnerId       *Many2One  `xmlrpc:"partner_id,omitempty"`
	Reference       *String    `xmlrpc:"reference,omitempty"`
	ReportTemplate  *Many2One  `xmlrpc:"report_template,omitempty"`
	ResId           *Int       `xmlrpc:"res_id,omitempty"`
	State           *Selection `xmlrpc:"state,omitempty"`
	StateId         *Many2One  `xmlrpc:"state_id,omitempty"`
	Street          *String    `xmlrpc:"street,omitempty"`
	Street2         *String    `xmlrpc:"street2,omitempty"`
	UserId          *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
	Zip             *String    `xmlrpc:"zip,omitempty"`
}

// SnailmailLetters represents array of snailmail.letter model.
type SnailmailLetters []SnailmailLetter

// SnailmailLetterModel is the odoo model name.
const SnailmailLetterModel = "snailmail.letter"

// Many2One convert SnailmailLetter to *Many2One.
func (sl *SnailmailLetter) Many2One() *Many2One {
	return NewMany2One(sl.Id.Get(), "")
}

// CreateSnailmailLetter creates a new snailmail.letter model and returns its id.
func (c *Client) CreateSnailmailLetter(sl *SnailmailLetter) (int64, error) {
	return c.Create(SnailmailLetterModel, sl)
}

// UpdateSnailmailLetter updates an existing snailmail.letter record.
func (c *Client) UpdateSnailmailLetter(sl *SnailmailLetter) error {
	return c.UpdateSnailmailLetters([]int64{sl.Id.Get()}, sl)
}

// UpdateSnailmailLetters updates existing snailmail.letter records.
// All records (represented by IDs) will be updated by sl values.
func (c *Client) UpdateSnailmailLetters(ids []int64, sl *SnailmailLetter) error {
	return c.Update(SnailmailLetterModel, ids, sl)
}

// DeleteSnailmailLetter deletes an existing snailmail.letter record.
func (c *Client) DeleteSnailmailLetter(id int64) error {
	return c.DeleteSnailmailLetters([]int64{id})
}

// DeleteSnailmailLetters deletes existing snailmail.letter records.
func (c *Client) DeleteSnailmailLetters(ids []int64) error {
	return c.Delete(SnailmailLetterModel, ids)
}

// GetSnailmailLetter gets snailmail.letter existing record.
func (c *Client) GetSnailmailLetter(id int64) (*SnailmailLetter, error) {
	sls, err := c.GetSnailmailLetters([]int64{id})
	if err != nil {
		return nil, err
	}
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("id %v of snailmail.letter not found", id)
}

// GetSnailmailLetters gets snailmail.letter existing records.
func (c *Client) GetSnailmailLetters(ids []int64) (*SnailmailLetters, error) {
	sls := &SnailmailLetters{}
	if err := c.Read(SnailmailLetterModel, ids, nil, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindSnailmailLetter finds snailmail.letter record by querying it with criteria.
func (c *Client) FindSnailmailLetter(criteria *Criteria) (*SnailmailLetter, error) {
	sls := &SnailmailLetters{}
	if err := c.SearchRead(SnailmailLetterModel, criteria, NewOptions().Limit(1), sls); err != nil {
		return nil, err
	}
	if sls != nil && len(*sls) > 0 {
		return &((*sls)[0]), nil
	}
	return nil, fmt.Errorf("snailmail.letter was not found")
}

// FindSnailmailLetters finds snailmail.letter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailLetters(criteria *Criteria, options *Options) (*SnailmailLetters, error) {
	sls := &SnailmailLetters{}
	if err := c.SearchRead(SnailmailLetterModel, criteria, options, sls); err != nil {
		return nil, err
	}
	return sls, nil
}

// FindSnailmailLetterIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailLetterIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SnailmailLetterModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSnailmailLetterId finds record id by querying it with criteria.
func (c *Client) FindSnailmailLetterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SnailmailLetterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("snailmail.letter was not found")
}
