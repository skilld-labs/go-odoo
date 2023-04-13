package odoo

import (
	"fmt"
)

// MailMassMailingStage represents mail.mass_mailing.stage model.
type MailMassMailingStage struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omptempty"`
	DisplayName *String   `xmlrpc:"display_name,omptempty"`
	Id          *Int      `xmlrpc:"id,omptempty"`
	Name        *String   `xmlrpc:"name,omptempty"`
	Sequence    *Int      `xmlrpc:"sequence,omptempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omptempty"`
}

// MailMassMailingStages represents array of mail.mass_mailing.stage model.
type MailMassMailingStages []MailMassMailingStage

// MailMassMailingStageModel is the odoo model name.
const MailMassMailingStageModel = "mail.mass_mailing.stage"

// Many2One convert MailMassMailingStage to *Many2One.
func (mms *MailMassMailingStage) Many2One() *Many2One {
	return NewMany2One(mms.Id.Get(), "")
}

// CreateMailMassMailingStage creates a new mail.mass_mailing.stage model and returns its id.
func (c *Client) CreateMailMassMailingStage(mms *MailMassMailingStage) (int64, error) {
	ids, err := c.CreateMailMassMailingStages([]*MailMassMailingStage{mms})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMassMailingStage creates a new mail.mass_mailing.stage model and returns its id.
func (c *Client) CreateMailMassMailingStages(mmss []*MailMassMailingStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmss {
		vv = append(vv, v)
	}
	return c.Create(MailMassMailingStageModel, vv)
}

// UpdateMailMassMailingStage updates an existing mail.mass_mailing.stage record.
func (c *Client) UpdateMailMassMailingStage(mms *MailMassMailingStage) error {
	return c.UpdateMailMassMailingStages([]int64{mms.Id.Get()}, mms)
}

// UpdateMailMassMailingStages updates existing mail.mass_mailing.stage records.
// All records (represented by ids) will be updated by mms values.
func (c *Client) UpdateMailMassMailingStages(ids []int64, mms *MailMassMailingStage) error {
	return c.Update(MailMassMailingStageModel, ids, mms)
}

// DeleteMailMassMailingStage deletes an existing mail.mass_mailing.stage record.
func (c *Client) DeleteMailMassMailingStage(id int64) error {
	return c.DeleteMailMassMailingStages([]int64{id})
}

// DeleteMailMassMailingStages deletes existing mail.mass_mailing.stage records.
func (c *Client) DeleteMailMassMailingStages(ids []int64) error {
	return c.Delete(MailMassMailingStageModel, ids)
}

// GetMailMassMailingStage gets mail.mass_mailing.stage existing record.
func (c *Client) GetMailMassMailingStage(id int64) (*MailMassMailingStage, error) {
	mmss, err := c.GetMailMassMailingStages([]int64{id})
	if err != nil {
		return nil, err
	}
	if mmss != nil && len(*mmss) > 0 {
		return &((*mmss)[0]), nil
	}
	return nil, fmt.Errorf("id %v of mail.mass_mailing.stage not found", id)
}

// GetMailMassMailingStages gets mail.mass_mailing.stage existing records.
func (c *Client) GetMailMassMailingStages(ids []int64) (*MailMassMailingStages, error) {
	mmss := &MailMassMailingStages{}
	if err := c.Read(MailMassMailingStageModel, ids, nil, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMassMailingStage finds mail.mass_mailing.stage record by querying it with criteria.
func (c *Client) FindMailMassMailingStage(criteria *Criteria) (*MailMassMailingStage, error) {
	mmss := &MailMassMailingStages{}
	if err := c.SearchRead(MailMassMailingStageModel, criteria, NewOptions().Limit(1), mmss); err != nil {
		return nil, err
	}
	if mmss != nil && len(*mmss) > 0 {
		return &((*mmss)[0]), nil
	}
	return nil, fmt.Errorf("mail.mass_mailing.stage was not found with criteria %v", criteria)
}

// FindMailMassMailingStages finds mail.mass_mailing.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingStages(criteria *Criteria, options *Options) (*MailMassMailingStages, error) {
	mmss := &MailMassMailingStages{}
	if err := c.SearchRead(MailMassMailingStageModel, criteria, options, mmss); err != nil {
		return nil, err
	}
	return mmss, nil
}

// FindMailMassMailingStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMassMailingStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(MailMassMailingStageModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindMailMassMailingStageId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("mail.mass_mailing.stage was not found with criteria %v and options %v", criteria, options)
}
