package odoo

// MailCannedResponse represents mail.canned.response model.
type MailCannedResponse struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	Description  *String   `xmlrpc:"description,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	GroupIds     *Relation `xmlrpc:"group_ids,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	IsEditable   *Bool     `xmlrpc:"is_editable,omitempty"`
	IsShared     *Bool     `xmlrpc:"is_shared,omitempty"`
	LastUsed     *Time     `xmlrpc:"last_used,omitempty"`
	Source       *String   `xmlrpc:"source,omitempty"`
	Substitution *String   `xmlrpc:"substitution,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailCannedResponses represents array of mail.canned.response model.
type MailCannedResponses []MailCannedResponse

// MailCannedResponseModel is the odoo model name.
const MailCannedResponseModel = "mail.canned.response"

// Many2One convert MailCannedResponse to *Many2One.
func (mcr *MailCannedResponse) Many2One() *Many2One {
	return NewMany2One(mcr.Id.Get(), "")
}

// CreateMailCannedResponse creates a new mail.canned.response model and returns its id.
func (c *Client) CreateMailCannedResponse(mcr *MailCannedResponse) (int64, error) {
	ids, err := c.CreateMailCannedResponses([]*MailCannedResponse{mcr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailCannedResponse creates a new mail.canned.response model and returns its id.
func (c *Client) CreateMailCannedResponses(mcrs []*MailCannedResponse) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcrs {
		vv = append(vv, v)
	}
	return c.Create(MailCannedResponseModel, vv, nil)
}

// UpdateMailCannedResponse updates an existing mail.canned.response record.
func (c *Client) UpdateMailCannedResponse(mcr *MailCannedResponse) error {
	return c.UpdateMailCannedResponses([]int64{mcr.Id.Get()}, mcr)
}

// UpdateMailCannedResponses updates existing mail.canned.response records.
// All records (represented by ids) will be updated by mcr values.
func (c *Client) UpdateMailCannedResponses(ids []int64, mcr *MailCannedResponse) error {
	return c.Update(MailCannedResponseModel, ids, mcr, nil)
}

// DeleteMailCannedResponse deletes an existing mail.canned.response record.
func (c *Client) DeleteMailCannedResponse(id int64) error {
	return c.DeleteMailCannedResponses([]int64{id})
}

// DeleteMailCannedResponses deletes existing mail.canned.response records.
func (c *Client) DeleteMailCannedResponses(ids []int64) error {
	return c.Delete(MailCannedResponseModel, ids)
}

// GetMailCannedResponse gets mail.canned.response existing record.
func (c *Client) GetMailCannedResponse(id int64) (*MailCannedResponse, error) {
	mcrs, err := c.GetMailCannedResponses([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcrs)[0]), nil
}

// GetMailCannedResponses gets mail.canned.response existing records.
func (c *Client) GetMailCannedResponses(ids []int64) (*MailCannedResponses, error) {
	mcrs := &MailCannedResponses{}
	if err := c.Read(MailCannedResponseModel, ids, nil, mcrs); err != nil {
		return nil, err
	}
	return mcrs, nil
}

// FindMailCannedResponse finds mail.canned.response record by querying it with criteria.
func (c *Client) FindMailCannedResponse(criteria *Criteria) (*MailCannedResponse, error) {
	mcrs := &MailCannedResponses{}
	if err := c.SearchRead(MailCannedResponseModel, criteria, NewOptions().Limit(1), mcrs); err != nil {
		return nil, err
	}
	return &((*mcrs)[0]), nil
}

// FindMailCannedResponses finds mail.canned.response records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailCannedResponses(criteria *Criteria, options *Options) (*MailCannedResponses, error) {
	mcrs := &MailCannedResponses{}
	if err := c.SearchRead(MailCannedResponseModel, criteria, options, mcrs); err != nil {
		return nil, err
	}
	return mcrs, nil
}

// FindMailCannedResponseIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailCannedResponseIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailCannedResponseModel, criteria, options)
}

// FindMailCannedResponseId finds record id by querying it with criteria.
func (c *Client) FindMailCannedResponseId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailCannedResponseModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
