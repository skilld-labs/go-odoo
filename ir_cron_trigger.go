package odoo

// IrCronTrigger represents ir.cron.trigger model.
type IrCronTrigger struct {
	CallAt      *Time     `xmlrpc:"call_at,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	CronId      *Many2One `xmlrpc:"cron_id,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrCronTriggers represents array of ir.cron.trigger model.
type IrCronTriggers []IrCronTrigger

// IrCronTriggerModel is the odoo model name.
const IrCronTriggerModel = "ir.cron.trigger"

// Many2One convert IrCronTrigger to *Many2One.
func (ict *IrCronTrigger) Many2One() *Many2One {
	return NewMany2One(ict.Id.Get(), "")
}

// CreateIrCronTrigger creates a new ir.cron.trigger model and returns its id.
func (c *Client) CreateIrCronTrigger(ict *IrCronTrigger) (int64, error) {
	ids, err := c.CreateIrCronTriggers([]*IrCronTrigger{ict})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrCronTrigger creates a new ir.cron.trigger model and returns its id.
func (c *Client) CreateIrCronTriggers(icts []*IrCronTrigger) ([]int64, error) {
	var vv []interface{}
	for _, v := range icts {
		vv = append(vv, v)
	}
	return c.Create(IrCronTriggerModel, vv, nil)
}

// UpdateIrCronTrigger updates an existing ir.cron.trigger record.
func (c *Client) UpdateIrCronTrigger(ict *IrCronTrigger) error {
	return c.UpdateIrCronTriggers([]int64{ict.Id.Get()}, ict)
}

// UpdateIrCronTriggers updates existing ir.cron.trigger records.
// All records (represented by ids) will be updated by ict values.
func (c *Client) UpdateIrCronTriggers(ids []int64, ict *IrCronTrigger) error {
	return c.Update(IrCronTriggerModel, ids, ict, nil)
}

// DeleteIrCronTrigger deletes an existing ir.cron.trigger record.
func (c *Client) DeleteIrCronTrigger(id int64) error {
	return c.DeleteIrCronTriggers([]int64{id})
}

// DeleteIrCronTriggers deletes existing ir.cron.trigger records.
func (c *Client) DeleteIrCronTriggers(ids []int64) error {
	return c.Delete(IrCronTriggerModel, ids)
}

// GetIrCronTrigger gets ir.cron.trigger existing record.
func (c *Client) GetIrCronTrigger(id int64) (*IrCronTrigger, error) {
	icts, err := c.GetIrCronTriggers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*icts)[0]), nil
}

// GetIrCronTriggers gets ir.cron.trigger existing records.
func (c *Client) GetIrCronTriggers(ids []int64) (*IrCronTriggers, error) {
	icts := &IrCronTriggers{}
	if err := c.Read(IrCronTriggerModel, ids, nil, icts); err != nil {
		return nil, err
	}
	return icts, nil
}

// FindIrCronTrigger finds ir.cron.trigger record by querying it with criteria.
func (c *Client) FindIrCronTrigger(criteria *Criteria) (*IrCronTrigger, error) {
	icts := &IrCronTriggers{}
	if err := c.SearchRead(IrCronTriggerModel, criteria, NewOptions().Limit(1), icts); err != nil {
		return nil, err
	}
	return &((*icts)[0]), nil
}

// FindIrCronTriggers finds ir.cron.trigger records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrCronTriggers(criteria *Criteria, options *Options) (*IrCronTriggers, error) {
	icts := &IrCronTriggers{}
	if err := c.SearchRead(IrCronTriggerModel, criteria, options, icts); err != nil {
		return nil, err
	}
	return icts, nil
}

// FindIrCronTriggerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrCronTriggerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrCronTriggerModel, criteria, options)
}

// FindIrCronTriggerId finds record id by querying it with criteria.
func (c *Client) FindIrCronTriggerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrCronTriggerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
