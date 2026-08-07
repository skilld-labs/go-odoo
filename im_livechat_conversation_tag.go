package odoo

// ImLivechatConversationTag represents im_livechat.conversation.tag model.
type ImLivechatConversationTag struct {
	Color           *Int      `xmlrpc:"color,omitempty"`
	ConversationIds *Relation `xmlrpc:"conversation_ids,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ImLivechatConversationTags represents array of im_livechat.conversation.tag model.
type ImLivechatConversationTags []ImLivechatConversationTag

// ImLivechatConversationTagModel is the odoo model name.
const ImLivechatConversationTagModel = "im_livechat.conversation.tag"

// Many2One convert ImLivechatConversationTag to *Many2One.
func (ict *ImLivechatConversationTag) Many2One() *Many2One {
	return NewMany2One(ict.Id.Get(), "")
}

// CreateImLivechatConversationTag creates a new im_livechat.conversation.tag model and returns its id.
func (c *Client) CreateImLivechatConversationTag(ict *ImLivechatConversationTag) (int64, error) {
	ids, err := c.CreateImLivechatConversationTags([]*ImLivechatConversationTag{ict})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateImLivechatConversationTag creates a new im_livechat.conversation.tag model and returns its id.
func (c *Client) CreateImLivechatConversationTags(icts []*ImLivechatConversationTag) ([]int64, error) {
	var vv []interface{}
	for _, v := range icts {
		vv = append(vv, v)
	}
	return c.Create(ImLivechatConversationTagModel, vv, nil)
}

// UpdateImLivechatConversationTag updates an existing im_livechat.conversation.tag record.
func (c *Client) UpdateImLivechatConversationTag(ict *ImLivechatConversationTag) error {
	return c.UpdateImLivechatConversationTags([]int64{ict.Id.Get()}, ict)
}

// UpdateImLivechatConversationTags updates existing im_livechat.conversation.tag records.
// All records (represented by ids) will be updated by ict values.
func (c *Client) UpdateImLivechatConversationTags(ids []int64, ict *ImLivechatConversationTag) error {
	return c.Update(ImLivechatConversationTagModel, ids, ict, nil)
}

// DeleteImLivechatConversationTag deletes an existing im_livechat.conversation.tag record.
func (c *Client) DeleteImLivechatConversationTag(id int64) error {
	return c.DeleteImLivechatConversationTags([]int64{id})
}

// DeleteImLivechatConversationTags deletes existing im_livechat.conversation.tag records.
func (c *Client) DeleteImLivechatConversationTags(ids []int64) error {
	return c.Delete(ImLivechatConversationTagModel, ids)
}

// GetImLivechatConversationTag gets im_livechat.conversation.tag existing record.
func (c *Client) GetImLivechatConversationTag(id int64) (*ImLivechatConversationTag, error) {
	icts, err := c.GetImLivechatConversationTags([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*icts)[0]), nil
}

// GetImLivechatConversationTags gets im_livechat.conversation.tag existing records.
func (c *Client) GetImLivechatConversationTags(ids []int64) (*ImLivechatConversationTags, error) {
	icts := &ImLivechatConversationTags{}
	if err := c.Read(ImLivechatConversationTagModel, ids, nil, icts); err != nil {
		return nil, err
	}
	return icts, nil
}

// FindImLivechatConversationTag finds im_livechat.conversation.tag record by querying it with criteria.
func (c *Client) FindImLivechatConversationTag(criteria *Criteria) (*ImLivechatConversationTag, error) {
	icts := &ImLivechatConversationTags{}
	if err := c.SearchRead(ImLivechatConversationTagModel, criteria, NewOptions().Limit(1), icts); err != nil {
		return nil, err
	}
	return &((*icts)[0]), nil
}

// FindImLivechatConversationTags finds im_livechat.conversation.tag records by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatConversationTags(criteria *Criteria, options *Options) (*ImLivechatConversationTags, error) {
	icts := &ImLivechatConversationTags{}
	if err := c.SearchRead(ImLivechatConversationTagModel, criteria, options, icts); err != nil {
		return nil, err
	}
	return icts, nil
}

// FindImLivechatConversationTagIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindImLivechatConversationTagIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ImLivechatConversationTagModel, criteria, options)
}

// FindImLivechatConversationTagId finds record id by querying it with criteria.
func (c *Client) FindImLivechatConversationTagId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ImLivechatConversationTagModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
