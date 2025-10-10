package odoo

// MailActivityTodoCreate represents mail.activity.todo.create model.
type MailActivityTodoCreate struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DateDeadline *Time     `xmlrpc:"date_deadline,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	Note         *String   `xmlrpc:"note,omitempty"`
	Summary      *String   `xmlrpc:"summary,omitempty"`
	UserId       *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// MailActivityTodoCreates represents array of mail.activity.todo.create model.
type MailActivityTodoCreates []MailActivityTodoCreate

// MailActivityTodoCreateModel is the odoo model name.
const MailActivityTodoCreateModel = "mail.activity.todo.create"

// Many2One convert MailActivityTodoCreate to *Many2One.
func (matc *MailActivityTodoCreate) Many2One() *Many2One {
	return NewMany2One(matc.Id.Get(), "")
}

// CreateMailActivityTodoCreate creates a new mail.activity.todo.create model and returns its id.
func (c *Client) CreateMailActivityTodoCreate(matc *MailActivityTodoCreate) (int64, error) {
	ids, err := c.CreateMailActivityTodoCreates([]*MailActivityTodoCreate{matc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailActivityTodoCreate creates a new mail.activity.todo.create model and returns its id.
func (c *Client) CreateMailActivityTodoCreates(matcs []*MailActivityTodoCreate) ([]int64, error) {
	var vv []interface{}
	for _, v := range matcs {
		vv = append(vv, v)
	}
	return c.Create(MailActivityTodoCreateModel, vv, nil)
}

// UpdateMailActivityTodoCreate updates an existing mail.activity.todo.create record.
func (c *Client) UpdateMailActivityTodoCreate(matc *MailActivityTodoCreate) error {
	return c.UpdateMailActivityTodoCreates([]int64{matc.Id.Get()}, matc)
}

// UpdateMailActivityTodoCreates updates existing mail.activity.todo.create records.
// All records (represented by ids) will be updated by matc values.
func (c *Client) UpdateMailActivityTodoCreates(ids []int64, matc *MailActivityTodoCreate) error {
	return c.Update(MailActivityTodoCreateModel, ids, matc, nil)
}

// DeleteMailActivityTodoCreate deletes an existing mail.activity.todo.create record.
func (c *Client) DeleteMailActivityTodoCreate(id int64) error {
	return c.DeleteMailActivityTodoCreates([]int64{id})
}

// DeleteMailActivityTodoCreates deletes existing mail.activity.todo.create records.
func (c *Client) DeleteMailActivityTodoCreates(ids []int64) error {
	return c.Delete(MailActivityTodoCreateModel, ids)
}

// GetMailActivityTodoCreate gets mail.activity.todo.create existing record.
func (c *Client) GetMailActivityTodoCreate(id int64) (*MailActivityTodoCreate, error) {
	matcs, err := c.GetMailActivityTodoCreates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*matcs)[0]), nil
}

// GetMailActivityTodoCreates gets mail.activity.todo.create existing records.
func (c *Client) GetMailActivityTodoCreates(ids []int64) (*MailActivityTodoCreates, error) {
	matcs := &MailActivityTodoCreates{}
	if err := c.Read(MailActivityTodoCreateModel, ids, nil, matcs); err != nil {
		return nil, err
	}
	return matcs, nil
}

// FindMailActivityTodoCreate finds mail.activity.todo.create record by querying it with criteria.
func (c *Client) FindMailActivityTodoCreate(criteria *Criteria) (*MailActivityTodoCreate, error) {
	matcs := &MailActivityTodoCreates{}
	if err := c.SearchRead(MailActivityTodoCreateModel, criteria, NewOptions().Limit(1), matcs); err != nil {
		return nil, err
	}
	return &((*matcs)[0]), nil
}

// FindMailActivityTodoCreates finds mail.activity.todo.create records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityTodoCreates(criteria *Criteria, options *Options) (*MailActivityTodoCreates, error) {
	matcs := &MailActivityTodoCreates{}
	if err := c.SearchRead(MailActivityTodoCreateModel, criteria, options, matcs); err != nil {
		return nil, err
	}
	return matcs, nil
}

// FindMailActivityTodoCreateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityTodoCreateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailActivityTodoCreateModel, criteria, options)
}

// FindMailActivityTodoCreateId finds record id by querying it with criteria.
func (c *Client) FindMailActivityTodoCreateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityTodoCreateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
