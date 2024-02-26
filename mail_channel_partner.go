package odoo

// MailChannelPartner represents mail.channel.partner model.
type MailChannelPartner struct {
	LastUpdate    *Time      `xmlrpc:"__last_update,omptempty"`
	ChannelId     *Many2One  `xmlrpc:"channel_id,omptempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName   *String    `xmlrpc:"display_name,omptempty"`
	FoldState     *Selection `xmlrpc:"fold_state,omptempty"`
	Id            *Int       `xmlrpc:"id,omptempty"`
	IsMinimized   *Bool      `xmlrpc:"is_minimized,omptempty"`
	IsPinned      *Bool      `xmlrpc:"is_pinned,omptempty"`
	PartnerEmail  *String    `xmlrpc:"partner_email,omptempty"`
	PartnerId     *Many2One  `xmlrpc:"partner_id,omptempty"`
	SeenMessageId *Many2One  `xmlrpc:"seen_message_id,omptempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// MailChannelPartners represents array of mail.channel.partner model.
type MailChannelPartners []MailChannelPartner

// MailChannelPartnerModel is the odoo model name.
const MailChannelPartnerModel = "mail.channel.partner"

// Many2One convert MailChannelPartner to *Many2One.
func (mcp *MailChannelPartner) Many2One() *Many2One {
	return NewMany2One(mcp.Id.Get(), "")
}

// CreateMailChannelPartner creates a new mail.channel.partner model and returns its id.
func (c *Client) CreateMailChannelPartner(mcp *MailChannelPartner) (int64, error) {
	ids, err := c.CreateMailChannelPartners([]*MailChannelPartner{mcp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailChannelPartner creates a new mail.channel.partner model and returns its id.
func (c *Client) CreateMailChannelPartners(mcps []*MailChannelPartner) ([]int64, error) {
	var vv []interface{}
	for _, v := range mcps {
		vv = append(vv, v)
	}
	return c.Create(MailChannelPartnerModel, vv, nil)
}

// UpdateMailChannelPartner updates an existing mail.channel.partner record.
func (c *Client) UpdateMailChannelPartner(mcp *MailChannelPartner) error {
	return c.UpdateMailChannelPartners([]int64{mcp.Id.Get()}, mcp)
}

// UpdateMailChannelPartners updates existing mail.channel.partner records.
// All records (represented by ids) will be updated by mcp values.
func (c *Client) UpdateMailChannelPartners(ids []int64, mcp *MailChannelPartner) error {
	return c.Update(MailChannelPartnerModel, ids, mcp, nil)
}

// DeleteMailChannelPartner deletes an existing mail.channel.partner record.
func (c *Client) DeleteMailChannelPartner(id int64) error {
	return c.DeleteMailChannelPartners([]int64{id})
}

// DeleteMailChannelPartners deletes existing mail.channel.partner records.
func (c *Client) DeleteMailChannelPartners(ids []int64) error {
	return c.Delete(MailChannelPartnerModel, ids)
}

// GetMailChannelPartner gets mail.channel.partner existing record.
func (c *Client) GetMailChannelPartner(id int64) (*MailChannelPartner, error) {
	mcps, err := c.GetMailChannelPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mcps)[0]), nil
}

// GetMailChannelPartners gets mail.channel.partner existing records.
func (c *Client) GetMailChannelPartners(ids []int64) (*MailChannelPartners, error) {
	mcps := &MailChannelPartners{}
	if err := c.Read(MailChannelPartnerModel, ids, nil, mcps); err != nil {
		return nil, err
	}
	return mcps, nil
}

// FindMailChannelPartner finds mail.channel.partner record by querying it with criteria.
func (c *Client) FindMailChannelPartner(criteria *Criteria) (*MailChannelPartner, error) {
	mcps := &MailChannelPartners{}
	if err := c.SearchRead(MailChannelPartnerModel, criteria, NewOptions().Limit(1), mcps); err != nil {
		return nil, err
	}
	return &((*mcps)[0]), nil
}

// FindMailChannelPartners finds mail.channel.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelPartners(criteria *Criteria, options *Options) (*MailChannelPartners, error) {
	mcps := &MailChannelPartners{}
	if err := c.SearchRead(MailChannelPartnerModel, criteria, options, mcps); err != nil {
		return nil, err
	}
	return mcps, nil
}

// FindMailChannelPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailChannelPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailChannelPartnerModel, criteria, options)
}

// FindMailChannelPartnerId finds record id by querying it with criteria.
func (c *Client) FindMailChannelPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailChannelPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
