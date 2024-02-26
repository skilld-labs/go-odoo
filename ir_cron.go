package odoo

// IrCron represents ir.cron model.
type IrCron struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omptempty"`
	Active            *Bool      `xmlrpc:"active,omptempty"`
	BindingModelId    *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType       *Selection `xmlrpc:"binding_type,omptempty"`
	ChannelIds        *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds          *Relation  `xmlrpc:"child_ids,omptempty"`
	Code              *String    `xmlrpc:"code,omptempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omptempty"`
	CronName          *String    `xmlrpc:"cron_name,omptempty"`
	CrudModelId       *Many2One  `xmlrpc:"crud_model_id,omptempty"`
	CrudModelName     *String    `xmlrpc:"crud_model_name,omptempty"`
	DisplayName       *String    `xmlrpc:"display_name,omptempty"`
	Doall             *Bool      `xmlrpc:"doall,omptempty"`
	FieldsLines       *Relation  `xmlrpc:"fields_lines,omptempty"`
	Help              *String    `xmlrpc:"help,omptempty"`
	Id                *Int       `xmlrpc:"id,omptempty"`
	IntervalNumber    *Int       `xmlrpc:"interval_number,omptempty"`
	IntervalType      *Selection `xmlrpc:"interval_type,omptempty"`
	IrActionsServerId *Many2One  `xmlrpc:"ir_actions_server_id,omptempty"`
	LinkFieldId       *Many2One  `xmlrpc:"link_field_id,omptempty"`
	ModelId           *Many2One  `xmlrpc:"model_id,omptempty"`
	ModelName         *String    `xmlrpc:"model_name,omptempty"`
	Name              *String    `xmlrpc:"name,omptempty"`
	Nextcall          *Time      `xmlrpc:"nextcall,omptempty"`
	Numbercall        *Int       `xmlrpc:"numbercall,omptempty"`
	PartnerIds        *Relation  `xmlrpc:"partner_ids,omptempty"`
	Priority          *Int       `xmlrpc:"priority,omptempty"`
	Sequence          *Int       `xmlrpc:"sequence,omptempty"`
	State             *Selection `xmlrpc:"state,omptempty"`
	TemplateId        *Many2One  `xmlrpc:"template_id,omptempty"`
	Type              *String    `xmlrpc:"type,omptempty"`
	Usage             *Selection `xmlrpc:"usage,omptempty"`
	UserId            *Many2One  `xmlrpc:"user_id,omptempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId             *String    `xmlrpc:"xml_id,omptempty"`
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

// CreateIrCron creates a new ir.cron model and returns its id.
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
