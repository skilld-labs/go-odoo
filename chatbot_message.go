package odoo

// ChatbotMessage represents chatbot.message model.
type ChatbotMessage struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DiscussChannelId   *Many2One `xmlrpc:"discuss_channel_id,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	MailMessageId      *Many2One `xmlrpc:"mail_message_id,omitempty"`
	ScriptStepId       *Many2One `xmlrpc:"script_step_id,omitempty"`
	UserRawAnswer      *String   `xmlrpc:"user_raw_answer,omitempty"`
	UserScriptAnswerId *Many2One `xmlrpc:"user_script_answer_id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ChatbotMessages represents array of chatbot.message model.
type ChatbotMessages []ChatbotMessage

// ChatbotMessageModel is the odoo model name.
const ChatbotMessageModel = "chatbot.message"

// Many2One convert ChatbotMessage to *Many2One.
func (cm *ChatbotMessage) Many2One() *Many2One {
	return NewMany2One(cm.Id.Get(), "")
}

// CreateChatbotMessage creates a new chatbot.message model and returns its id.
func (c *Client) CreateChatbotMessage(cm *ChatbotMessage) (int64, error) {
	ids, err := c.CreateChatbotMessages([]*ChatbotMessage{cm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChatbotMessage creates a new chatbot.message model and returns its id.
func (c *Client) CreateChatbotMessages(cms []*ChatbotMessage) ([]int64, error) {
	var vv []interface{}
	for _, v := range cms {
		vv = append(vv, v)
	}
	return c.Create(ChatbotMessageModel, vv, nil)
}

// UpdateChatbotMessage updates an existing chatbot.message record.
func (c *Client) UpdateChatbotMessage(cm *ChatbotMessage) error {
	return c.UpdateChatbotMessages([]int64{cm.Id.Get()}, cm)
}

// UpdateChatbotMessages updates existing chatbot.message records.
// All records (represented by ids) will be updated by cm values.
func (c *Client) UpdateChatbotMessages(ids []int64, cm *ChatbotMessage) error {
	return c.Update(ChatbotMessageModel, ids, cm, nil)
}

// DeleteChatbotMessage deletes an existing chatbot.message record.
func (c *Client) DeleteChatbotMessage(id int64) error {
	return c.DeleteChatbotMessages([]int64{id})
}

// DeleteChatbotMessages deletes existing chatbot.message records.
func (c *Client) DeleteChatbotMessages(ids []int64) error {
	return c.Delete(ChatbotMessageModel, ids)
}

// GetChatbotMessage gets chatbot.message existing record.
func (c *Client) GetChatbotMessage(id int64) (*ChatbotMessage, error) {
	cms, err := c.GetChatbotMessages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cms)[0]), nil
}

// GetChatbotMessages gets chatbot.message existing records.
func (c *Client) GetChatbotMessages(ids []int64) (*ChatbotMessages, error) {
	cms := &ChatbotMessages{}
	if err := c.Read(ChatbotMessageModel, ids, nil, cms); err != nil {
		return nil, err
	}
	return cms, nil
}

// FindChatbotMessage finds chatbot.message record by querying it with criteria.
func (c *Client) FindChatbotMessage(criteria *Criteria) (*ChatbotMessage, error) {
	cms := &ChatbotMessages{}
	if err := c.SearchRead(ChatbotMessageModel, criteria, NewOptions().Limit(1), cms); err != nil {
		return nil, err
	}
	return &((*cms)[0]), nil
}

// FindChatbotMessages finds chatbot.message records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotMessages(criteria *Criteria, options *Options) (*ChatbotMessages, error) {
	cms := &ChatbotMessages{}
	if err := c.SearchRead(ChatbotMessageModel, criteria, options, cms); err != nil {
		return nil, err
	}
	return cms, nil
}

// FindChatbotMessageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotMessageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ChatbotMessageModel, criteria, options)
}

// FindChatbotMessageId finds record id by querying it with criteria.
func (c *Client) FindChatbotMessageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChatbotMessageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
