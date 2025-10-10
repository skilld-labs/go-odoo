package odoo

// ChatbotScriptAnswer represents chatbot.script.answer model.
type ChatbotScriptAnswer struct {
	ChatbotScriptId *Many2One `xmlrpc:"chatbot_script_id,omitempty"`
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Name            *String   `xmlrpc:"name,omitempty"`
	RedirectLink    *String   `xmlrpc:"redirect_link,omitempty"`
	ScriptStepId    *Many2One `xmlrpc:"script_step_id,omitempty"`
	Sequence        *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ChatbotScriptAnswers represents array of chatbot.script.answer model.
type ChatbotScriptAnswers []ChatbotScriptAnswer

// ChatbotScriptAnswerModel is the odoo model name.
const ChatbotScriptAnswerModel = "chatbot.script.answer"

// Many2One convert ChatbotScriptAnswer to *Many2One.
func (csa *ChatbotScriptAnswer) Many2One() *Many2One {
	return NewMany2One(csa.Id.Get(), "")
}

// CreateChatbotScriptAnswer creates a new chatbot.script.answer model and returns its id.
func (c *Client) CreateChatbotScriptAnswer(csa *ChatbotScriptAnswer) (int64, error) {
	ids, err := c.CreateChatbotScriptAnswers([]*ChatbotScriptAnswer{csa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChatbotScriptAnswer creates a new chatbot.script.answer model and returns its id.
func (c *Client) CreateChatbotScriptAnswers(csas []*ChatbotScriptAnswer) ([]int64, error) {
	var vv []interface{}
	for _, v := range csas {
		vv = append(vv, v)
	}
	return c.Create(ChatbotScriptAnswerModel, vv, nil)
}

// UpdateChatbotScriptAnswer updates an existing chatbot.script.answer record.
func (c *Client) UpdateChatbotScriptAnswer(csa *ChatbotScriptAnswer) error {
	return c.UpdateChatbotScriptAnswers([]int64{csa.Id.Get()}, csa)
}

// UpdateChatbotScriptAnswers updates existing chatbot.script.answer records.
// All records (represented by ids) will be updated by csa values.
func (c *Client) UpdateChatbotScriptAnswers(ids []int64, csa *ChatbotScriptAnswer) error {
	return c.Update(ChatbotScriptAnswerModel, ids, csa, nil)
}

// DeleteChatbotScriptAnswer deletes an existing chatbot.script.answer record.
func (c *Client) DeleteChatbotScriptAnswer(id int64) error {
	return c.DeleteChatbotScriptAnswers([]int64{id})
}

// DeleteChatbotScriptAnswers deletes existing chatbot.script.answer records.
func (c *Client) DeleteChatbotScriptAnswers(ids []int64) error {
	return c.Delete(ChatbotScriptAnswerModel, ids)
}

// GetChatbotScriptAnswer gets chatbot.script.answer existing record.
func (c *Client) GetChatbotScriptAnswer(id int64) (*ChatbotScriptAnswer, error) {
	csas, err := c.GetChatbotScriptAnswers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*csas)[0]), nil
}

// GetChatbotScriptAnswers gets chatbot.script.answer existing records.
func (c *Client) GetChatbotScriptAnswers(ids []int64) (*ChatbotScriptAnswers, error) {
	csas := &ChatbotScriptAnswers{}
	if err := c.Read(ChatbotScriptAnswerModel, ids, nil, csas); err != nil {
		return nil, err
	}
	return csas, nil
}

// FindChatbotScriptAnswer finds chatbot.script.answer record by querying it with criteria.
func (c *Client) FindChatbotScriptAnswer(criteria *Criteria) (*ChatbotScriptAnswer, error) {
	csas := &ChatbotScriptAnswers{}
	if err := c.SearchRead(ChatbotScriptAnswerModel, criteria, NewOptions().Limit(1), csas); err != nil {
		return nil, err
	}
	return &((*csas)[0]), nil
}

// FindChatbotScriptAnswers finds chatbot.script.answer records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotScriptAnswers(criteria *Criteria, options *Options) (*ChatbotScriptAnswers, error) {
	csas := &ChatbotScriptAnswers{}
	if err := c.SearchRead(ChatbotScriptAnswerModel, criteria, options, csas); err != nil {
		return nil, err
	}
	return csas, nil
}

// FindChatbotScriptAnswerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotScriptAnswerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ChatbotScriptAnswerModel, criteria, options)
}

// FindChatbotScriptAnswerId finds record id by querying it with criteria.
func (c *Client) FindChatbotScriptAnswerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChatbotScriptAnswerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
