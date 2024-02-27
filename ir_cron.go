package odoo

// IrCron represents ir.cron model.
type IrCron struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omitempty"`
	Active            *Bool      `xmlrpc:"active,omitempty"`
	BindingModelId    *Many2One  `xmlrpc:"binding_model_id,omitempty"`
	BindingType       *Selection `xmlrpc:"binding_type,omitempty"`
	ChannelIds        *Relation  `xmlrpc:"channel_ids,omitempty"`
	ChildIds          *Relation  `xmlrpc:"child_ids,omitempty"`
	Code              *String    `xmlrpc:"code,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	CronName          *String    `xmlrpc:"cron_name,omitempty"`
	CrudModelId       *Many2One  `xmlrpc:"crud_model_id,omitempty"`
	CrudModelName     *String    `xmlrpc:"crud_model_name,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Doall             *Bool      `xmlrpc:"doall,omitempty"`
	FieldsLines       *Relation  `xmlrpc:"fields_lines,omitempty"`
	Help              *String    `xmlrpc:"help,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	IntervalNumber    *Int       `xmlrpc:"interval_number,omitempty"`
	IntervalType      *Selection `xmlrpc:"interval_type,omitempty"`
	IrActionsServerId *Many2One  `xmlrpc:"ir_actions_server_id,omitempty"`
	LinkFieldId       *Many2One  `xmlrpc:"link_field_id,omitempty"`
	ModelId           *Many2One  `xmlrpc:"model_id,omitempty"`
	ModelName         *String    `xmlrpc:"model_name,omitempty"`
	Name              *String    `xmlrpc:"name,omitempty"`
	Nextcall          *Time      `xmlrpc:"nextcall,omitempty"`
	Numbercall        *Int       `xmlrpc:"numbercall,omitempty"`
	PartnerIds        *Relation  `xmlrpc:"partner_ids,omitempty"`
	Priority          *Int       `xmlrpc:"priority,omitempty"`
	Sequence          *Int       `xmlrpc:"sequence,omitempty"`
	State             *Selection `xmlrpc:"state,omitempty"`
	TemplateId        *Many2One  `xmlrpc:"template_id,omitempty"`
	Type              *String    `xmlrpc:"type,omitempty"`
	Usage             *Selection `xmlrpc:"usage,omitempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId             *String    `xmlrpc:"xml_id,omitempty"`
}

// IrCrons represents array of ir.cron model.
type IrCrons []IrCron

// IrCronModel is the odoo model name.
const IrCronModel = "ir.cron"

// Many2One convert IrCron to *Many2One.
func (ic *IrCron) Many2One() *Many2One {
	return NewMany2One(ic.Id.Get(), "")
}

// CreateIrCron creates a new ir.cron model and returns its id.
func (c *Client) CreateIrCron(ic *IrCron) (int64, error) {
	ids, err := c.CreateIrCrons([]*IrCron{ic})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrCrons creates a new ir.cron model and returns its id.
func (c *Client) CreateIrCrons(ics []*IrCron) ([]int64, error) {
	var vv []interface{}
	for _, v := range ics {
		vv = append(vv, v)
	}
	return c.Create(IrCronModel, vv, nil)
}

// UpdateIrCron updates an existing ir.cron record.
func (c *Client) UpdateIrCron(ic *IrCron) error {
	return c.UpdateIrCrons([]int64{ic.Id.Get()}, ic)
}

// UpdateIrCrons updates existing ir.cron records.
// All records (represented by ids) will be updated by ic values.
func (c *Client) UpdateIrCrons(ids []int64, ic *IrCron) error {
	return c.Update(IrCronModel, ids, ic, nil)
}

// DeleteIrCron deletes an existing ir.cron record.
func (c *Client) DeleteIrCron(id int64) error {
	return c.DeleteIrCrons([]int64{id})
}

// DeleteIrCrons deletes existing ir.cron records.
func (c *Client) DeleteIrCrons(ids []int64) error {
	return c.Delete(IrCronModel, ids)
}

// GetIrCron gets ir.cron existing record.
func (c *Client) GetIrCron(id int64) (*IrCron, error) {
	ics, err := c.GetIrCrons([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ics)[0]), nil
}

// GetIrCrons gets ir.cron existing records.
func (c *Client) GetIrCrons(ids []int64) (*IrCrons, error) {
	ics := &IrCrons{}
	if err := c.Read(IrCronModel, ids, nil, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindIrCron finds ir.cron record by querying it with criteria.
func (c *Client) FindIrCron(criteria *Criteria) (*IrCron, error) {
	ics := &IrCrons{}
	if err := c.SearchRead(IrCronModel, criteria, NewOptions().Limit(1), ics); err != nil {
		return nil, err
	}
	return &((*ics)[0]), nil
}

// FindIrCrons finds ir.cron records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrCrons(criteria *Criteria, options *Options) (*IrCrons, error) {
	ics := &IrCrons{}
	if err := c.SearchRead(IrCronModel, criteria, options, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindIrCronIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrCronIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrCronModel, criteria, options)
}

// FindIrCronId finds record id by querying it with criteria.
func (c *Client) FindIrCronId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrCronModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
