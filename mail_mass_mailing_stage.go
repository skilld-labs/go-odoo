package odoo

// MailMassMailingStage represents mail.mass_mailing.stage model.
type MailMassMailingStage struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
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
	return c.Create(MailMassMailingStageModel, vv, nil)
}

// UpdateMailMassMailingStage updates an existing mail.mass_mailing.stage record.
func (c *Client) UpdateMailMassMailingStage(mms *MailMassMailingStage) error {
	return c.UpdateMailMassMailingStages([]int64{mms.Id.Get()}, mms)
}

// UpdateMailMassMailingStages updates existing mail.mass_mailing.stage records.
// All records (represented by ids) will be updated by mms values.
func (c *Client) UpdateMailMassMailingStages(ids []int64, mms *MailMassMailingStage) error {
	return c.Update(MailMassMailingStageModel, ids, mms, nil)
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
	return &((*mmss)[0]), nil
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
	return &((*mmss)[0]), nil
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
	return c.Search(MailMassMailingStageModel, criteria, options)
}

// FindMailMassMailingStageId finds record id by querying it with criteria.
func (c *Client) FindMailMassMailingStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMassMailingStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
