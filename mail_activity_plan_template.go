package odoo

// MailActivityPlanTemplate represents mail.activity.plan.template model.
type MailActivityPlanTemplate struct {
	ActivityTypeId  *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omitempty"`
	DelayCount      *Int       `xmlrpc:"delay_count,omitempty"`
	DelayFrom       *Selection `xmlrpc:"delay_from,omitempty"`
	DelayUnit       *Selection `xmlrpc:"delay_unit,omitempty"`
	DisplayName     *String    `xmlrpc:"display_name,omitempty"`
	Icon            *String    `xmlrpc:"icon,omitempty"`
	Id              *Int       `xmlrpc:"id,omitempty"`
	Note            *String    `xmlrpc:"note,omitempty"`
	PlanId          *Many2One  `xmlrpc:"plan_id,omitempty"`
	ResModel        *Selection `xmlrpc:"res_model,omitempty"`
	ResponsibleId   *Many2One  `xmlrpc:"responsible_id,omitempty"`
	ResponsibleType *Selection `xmlrpc:"responsible_type,omitempty"`
	Sequence        *Int       `xmlrpc:"sequence,omitempty"`
	Summary         *String    `xmlrpc:"summary,omitempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// MailActivityPlanTemplates represents array of mail.activity.plan.template model.
type MailActivityPlanTemplates []MailActivityPlanTemplate

// MailActivityPlanTemplateModel is the odoo model name.
const MailActivityPlanTemplateModel = "mail.activity.plan.template"

// Many2One convert MailActivityPlanTemplate to *Many2One.
func (mapt *MailActivityPlanTemplate) Many2One() *Many2One {
	return NewMany2One(mapt.Id.Get(), "")
}

// CreateMailActivityPlanTemplate creates a new mail.activity.plan.template model and returns its id.
func (c *Client) CreateMailActivityPlanTemplate(mapt *MailActivityPlanTemplate) (int64, error) {
	ids, err := c.CreateMailActivityPlanTemplates([]*MailActivityPlanTemplate{mapt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateMailActivityPlanTemplate creates a new mail.activity.plan.template model and returns its id.
func (c *Client) CreateMailActivityPlanTemplates(mapts []*MailActivityPlanTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range mapts {
		vv = append(vv, v)
	}
	return c.Create(MailActivityPlanTemplateModel, vv, nil)
}

// UpdateMailActivityPlanTemplate updates an existing mail.activity.plan.template record.
func (c *Client) UpdateMailActivityPlanTemplate(mapt *MailActivityPlanTemplate) error {
	return c.UpdateMailActivityPlanTemplates([]int64{mapt.Id.Get()}, mapt)
}

// UpdateMailActivityPlanTemplates updates existing mail.activity.plan.template records.
// All records (represented by ids) will be updated by mapt values.
func (c *Client) UpdateMailActivityPlanTemplates(ids []int64, mapt *MailActivityPlanTemplate) error {
	return c.Update(MailActivityPlanTemplateModel, ids, mapt, nil)
}

// DeleteMailActivityPlanTemplate deletes an existing mail.activity.plan.template record.
func (c *Client) DeleteMailActivityPlanTemplate(id int64) error {
	return c.DeleteMailActivityPlanTemplates([]int64{id})
}

// DeleteMailActivityPlanTemplates deletes existing mail.activity.plan.template records.
func (c *Client) DeleteMailActivityPlanTemplates(ids []int64) error {
	return c.Delete(MailActivityPlanTemplateModel, ids)
}

// GetMailActivityPlanTemplate gets mail.activity.plan.template existing record.
func (c *Client) GetMailActivityPlanTemplate(id int64) (*MailActivityPlanTemplate, error) {
	mapts, err := c.GetMailActivityPlanTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*mapts)[0]), nil
}

// GetMailActivityPlanTemplates gets mail.activity.plan.template existing records.
func (c *Client) GetMailActivityPlanTemplates(ids []int64) (*MailActivityPlanTemplates, error) {
	mapts := &MailActivityPlanTemplates{}
	if err := c.Read(MailActivityPlanTemplateModel, ids, nil, mapts); err != nil {
		return nil, err
	}
	return mapts, nil
}

// FindMailActivityPlanTemplate finds mail.activity.plan.template record by querying it with criteria.
func (c *Client) FindMailActivityPlanTemplate(criteria *Criteria) (*MailActivityPlanTemplate, error) {
	mapts := &MailActivityPlanTemplates{}
	if err := c.SearchRead(MailActivityPlanTemplateModel, criteria, NewOptions().Limit(1), mapts); err != nil {
		return nil, err
	}
	return &((*mapts)[0]), nil
}

// FindMailActivityPlanTemplates finds mail.activity.plan.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityPlanTemplates(criteria *Criteria, options *Options) (*MailActivityPlanTemplates, error) {
	mapts := &MailActivityPlanTemplates{}
	if err := c.SearchRead(MailActivityPlanTemplateModel, criteria, options, mapts); err != nil {
		return nil, err
	}
	return mapts, nil
}

// FindMailActivityPlanTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindMailActivityPlanTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(MailActivityPlanTemplateModel, criteria, options)
}

// FindMailActivityPlanTemplateId finds record id by querying it with criteria.
func (c *Client) FindMailActivityPlanTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(MailActivityPlanTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
