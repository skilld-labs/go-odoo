package odoo

import (
	"fmt"
)

// SnailmailLetterFormatError represents snailmail.letter.format.error model.
type SnailmailLetterFormatError struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	MessageId      *Many2One `xmlrpc:"message_id,omitempty"`
	SnailmailCover *Bool     `xmlrpc:"snailmail_cover,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SnailmailLetterFormatErrors represents array of snailmail.letter.format.error model.
type SnailmailLetterFormatErrors []SnailmailLetterFormatError

// SnailmailLetterFormatErrorModel is the odoo model name.
const SnailmailLetterFormatErrorModel = "snailmail.letter.format.error"

// Many2One convert SnailmailLetterFormatError to *Many2One.
func (slfe *SnailmailLetterFormatError) Many2One() *Many2One {
	return NewMany2One(slfe.Id.Get(), "")
}

// CreateSnailmailLetterFormatError creates a new snailmail.letter.format.error model and returns its id.
func (c *Client) CreateSnailmailLetterFormatError(slfe *SnailmailLetterFormatError) (int64, error) {
	return c.Create(SnailmailLetterFormatErrorModel, slfe)
}

// UpdateSnailmailLetterFormatError updates an existing snailmail.letter.format.error record.
func (c *Client) UpdateSnailmailLetterFormatError(slfe *SnailmailLetterFormatError) error {
	return c.UpdateSnailmailLetterFormatErrors([]int64{slfe.Id.Get()}, slfe)
}

// UpdateSnailmailLetterFormatErrors updates existing snailmail.letter.format.error records.
// All records (represented by IDs) will be updated by slfe values.
func (c *Client) UpdateSnailmailLetterFormatErrors(ids []int64, slfe *SnailmailLetterFormatError) error {
	return c.Update(SnailmailLetterFormatErrorModel, ids, slfe)
}

// DeleteSnailmailLetterFormatError deletes an existing snailmail.letter.format.error record.
func (c *Client) DeleteSnailmailLetterFormatError(id int64) error {
	return c.DeleteSnailmailLetterFormatErrors([]int64{id})
}

// DeleteSnailmailLetterFormatErrors deletes existing snailmail.letter.format.error records.
func (c *Client) DeleteSnailmailLetterFormatErrors(ids []int64) error {
	return c.Delete(SnailmailLetterFormatErrorModel, ids)
}

// GetSnailmailLetterFormatError gets snailmail.letter.format.error existing record.
func (c *Client) GetSnailmailLetterFormatError(id int64) (*SnailmailLetterFormatError, error) {
	slfes, err := c.GetSnailmailLetterFormatErrors([]int64{id})
	if err != nil {
		return nil, err
	}
	if slfes != nil && len(*slfes) > 0 {
		return &((*slfes)[0]), nil
	}
	return nil, fmt.Errorf("id %v of snailmail.letter.format.error not found", id)
}

// GetSnailmailLetterFormatErrors gets snailmail.letter.format.error existing records.
func (c *Client) GetSnailmailLetterFormatErrors(ids []int64) (*SnailmailLetterFormatErrors, error) {
	slfes := &SnailmailLetterFormatErrors{}
	if err := c.Read(SnailmailLetterFormatErrorModel, ids, nil, slfes); err != nil {
		return nil, err
	}
	return slfes, nil
}

// FindSnailmailLetterFormatError finds snailmail.letter.format.error record by querying it with criteria.
func (c *Client) FindSnailmailLetterFormatError(criteria *Criteria) (*SnailmailLetterFormatError, error) {
	slfes := &SnailmailLetterFormatErrors{}
	if err := c.SearchRead(SnailmailLetterFormatErrorModel, criteria, NewOptions().Limit(1), slfes); err != nil {
		return nil, err
	}
	if slfes != nil && len(*slfes) > 0 {
		return &((*slfes)[0]), nil
	}
	return nil, fmt.Errorf("snailmail.letter.format.error was not found")
}

// FindSnailmailLetterFormatErrors finds snailmail.letter.format.error records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailLetterFormatErrors(criteria *Criteria, options *Options) (*SnailmailLetterFormatErrors, error) {
	slfes := &SnailmailLetterFormatErrors{}
	if err := c.SearchRead(SnailmailLetterFormatErrorModel, criteria, options, slfes); err != nil {
		return nil, err
	}
	return slfes, nil
}

// FindSnailmailLetterFormatErrorIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindSnailmailLetterFormatErrorIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(SnailmailLetterFormatErrorModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindSnailmailLetterFormatErrorId finds record id by querying it with criteria.
func (c *Client) FindSnailmailLetterFormatErrorId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SnailmailLetterFormatErrorModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("snailmail.letter.format.error was not found")
}
