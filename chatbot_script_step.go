package odoo

// ChatbotScriptStep represents chatbot.script.step model.
type ChatbotScriptStep struct {
	AnswerIds              *Relation  `xmlrpc:"answer_ids,omitempty"`
	ChatbotScriptId        *Many2One  `xmlrpc:"chatbot_script_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrmTeamId              *Many2One  `xmlrpc:"crm_team_id,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	IsForwardOperatorChild *Bool      `xmlrpc:"is_forward_operator_child,omitempty"`
	Message                *String    `xmlrpc:"message,omitempty"`
	Sequence               *Int       `xmlrpc:"sequence,omitempty"`
	StepType               *Selection `xmlrpc:"step_type,omitempty"`
	TriggeringAnswerIds    *Relation  `xmlrpc:"triggering_answer_ids,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ChatbotScriptSteps represents array of chatbot.script.step model.
type ChatbotScriptSteps []ChatbotScriptStep

// ChatbotScriptStepModel is the odoo model name.
const ChatbotScriptStepModel = "chatbot.script.step"

// Many2One convert ChatbotScriptStep to *Many2One.
func (css *ChatbotScriptStep) Many2One() *Many2One {
	return NewMany2One(css.Id.Get(), "")
}

// CreateChatbotScriptStep creates a new chatbot.script.step model and returns its id.
func (c *Client) CreateChatbotScriptStep(css *ChatbotScriptStep) (int64, error) {
	ids, err := c.CreateChatbotScriptSteps([]*ChatbotScriptStep{css})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateChatbotScriptStep creates a new chatbot.script.step model and returns its id.
func (c *Client) CreateChatbotScriptSteps(csss []*ChatbotScriptStep) ([]int64, error) {
	var vv []interface{}
	for _, v := range csss {
		vv = append(vv, v)
	}
	return c.Create(ChatbotScriptStepModel, vv, nil)
}

// UpdateChatbotScriptStep updates an existing chatbot.script.step record.
func (c *Client) UpdateChatbotScriptStep(css *ChatbotScriptStep) error {
	return c.UpdateChatbotScriptSteps([]int64{css.Id.Get()}, css)
}

// UpdateChatbotScriptSteps updates existing chatbot.script.step records.
// All records (represented by ids) will be updated by css values.
func (c *Client) UpdateChatbotScriptSteps(ids []int64, css *ChatbotScriptStep) error {
	return c.Update(ChatbotScriptStepModel, ids, css, nil)
}

// DeleteChatbotScriptStep deletes an existing chatbot.script.step record.
func (c *Client) DeleteChatbotScriptStep(id int64) error {
	return c.DeleteChatbotScriptSteps([]int64{id})
}

// DeleteChatbotScriptSteps deletes existing chatbot.script.step records.
func (c *Client) DeleteChatbotScriptSteps(ids []int64) error {
	return c.Delete(ChatbotScriptStepModel, ids)
}

// GetChatbotScriptStep gets chatbot.script.step existing record.
func (c *Client) GetChatbotScriptStep(id int64) (*ChatbotScriptStep, error) {
	csss, err := c.GetChatbotScriptSteps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*csss)[0]), nil
}

// GetChatbotScriptSteps gets chatbot.script.step existing records.
func (c *Client) GetChatbotScriptSteps(ids []int64) (*ChatbotScriptSteps, error) {
	csss := &ChatbotScriptSteps{}
	if err := c.Read(ChatbotScriptStepModel, ids, nil, csss); err != nil {
		return nil, err
	}
	return csss, nil
}

// FindChatbotScriptStep finds chatbot.script.step record by querying it with criteria.
func (c *Client) FindChatbotScriptStep(criteria *Criteria) (*ChatbotScriptStep, error) {
	csss := &ChatbotScriptSteps{}
	if err := c.SearchRead(ChatbotScriptStepModel, criteria, NewOptions().Limit(1), csss); err != nil {
		return nil, err
	}
	return &((*csss)[0]), nil
}

// FindChatbotScriptSteps finds chatbot.script.step records by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotScriptSteps(criteria *Criteria, options *Options) (*ChatbotScriptSteps, error) {
	csss := &ChatbotScriptSteps{}
	if err := c.SearchRead(ChatbotScriptStepModel, criteria, options, csss); err != nil {
		return nil, err
	}
	return csss, nil
}

// FindChatbotScriptStepIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindChatbotScriptStepIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ChatbotScriptStepModel, criteria, options)
}

// FindChatbotScriptStepId finds record id by querying it with criteria.
func (c *Client) FindChatbotScriptStepId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ChatbotScriptStepModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
