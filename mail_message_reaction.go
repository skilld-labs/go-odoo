package odoo

// MailMessageReaction represents mail.message.reaction model.
type MailMessageReaction struct {
	Content     *String   `xmlrpc:"content,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	GuestId     *Many2One `xmlrpc:"guest_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	MessageId   *Many2One `xmlrpc:"message_id,omitempty"`
	PartnerId   *Many2One `xmlrpc:"partner_id,omitempty"`
}

// MailMessageReactions represents array of mail.message.reaction model.
type MailMessageReactions []MailMessageReaction

// MailMessageReactionModel is the odoo model name.
const MailMessageReactionModel = "mail.message.reaction"

// Many2One convert MailMessageReaction to *Many2One.
func (mmr *MailMessageReaction) Many2One() *Many2One {
	return NewMany2One(mmr.Id.Get(), "")
}

// CreateMailMessageReaction creates a new mail.message.reaction model and returns its id.
func (c *Client) CreateMailMessageReaction(mmr *MailMessageReaction) (int64, error) {
	ids, err := c.CreateMailMessageReactions([]*MailMessageReaction{mmr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailMessageReaction creates a new mail.message.reaction model and returns its id.
func (c *Client) CreateMailMessageReactions(mmrs []*MailMessageReaction) ([]int64, error) {
	var vv []interface{}
	for _, v := range mmrs {
		vv = append(vv, v)
	}
	return c.Create(MailMessageReactionModel, vv, nil)
}

// UpdateMailMessageReaction updates an existing mail.message.reaction record.
func (c *Client) UpdateMailMessageReaction(mmr *MailMessageReaction) error {
	return c.UpdateMailMessageReactions([]int64{mmr.Id.Get()}, mmr)
}

// UpdateMailMessageReactions updates existing mail.message.reaction records.
// All records (represented by ids) will be updated by mmr values.
func (c *Client) UpdateMailMessageReactions(ids []int64, mmr *MailMessageReaction) error {
	return c.Update(MailMessageReactionModel, ids, mmr, nil)
}

// DeleteMailMessageReaction deletes an existing mail.message.reaction record.
func (c *Client) DeleteMailMessageReaction(id int64) error {
	return c.DeleteMailMessageReactions([]int64{id})
}

// DeleteMailMessageReactions deletes existing mail.message.reaction records.
func (c *Client) DeleteMailMessageReactions(ids []int64) error {
	return c.Delete(MailMessageReactionModel, ids)
}

// GetMailMessageReaction gets mail.message.reaction existing record.
func (c *Client) GetMailMessageReaction(id int64) (*MailMessageReaction, error) {
	mmrs, err := c.GetMailMessageReactions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mmrs)[0]), nil
}

// GetMailMessageReactions gets mail.message.reaction existing records.
func (c *Client) GetMailMessageReactions(ids []int64) (*MailMessageReactions, error) {
	mmrs := &MailMessageReactions{}
	if err := c.Read(MailMessageReactionModel, ids, nil, mmrs); err != nil {
		return nil, err
	}
	return mmrs, nil
}

// FindMailMessageReaction finds mail.message.reaction record by querying it with criteria.
func (c *Client) FindMailMessageReaction(criteria *Criteria) (*MailMessageReaction, error) {
	mmrs := &MailMessageReactions{}
	if err := c.SearchRead(MailMessageReactionModel, criteria, NewOptions().Limit(1), mmrs); err != nil {
		return nil, err
	}
	return &((*mmrs)[0]), nil
}

// FindMailMessageReactions finds mail.message.reaction records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageReactions(criteria *Criteria, options *Options) (*MailMessageReactions, error) {
	mmrs := &MailMessageReactions{}
	if err := c.SearchRead(MailMessageReactionModel, criteria, options, mmrs); err != nil {
		return nil, err
	}
	return mmrs, nil
}

// FindMailMessageReactionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailMessageReactionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailMessageReactionModel, criteria, options)
}

// FindMailMessageReactionId finds record id by querying it with criteria.
func (c *Client) FindMailMessageReactionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailMessageReactionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
