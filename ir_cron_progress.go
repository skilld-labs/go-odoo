package odoo

// IrCronProgress represents ir.cron.progress model.
type IrCronProgress struct {
	CreateDate      *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid       *Many2One `xmlrpc:"create_uid,omitempty"`
	CronId          *Many2One `xmlrpc:"cron_id,omitempty"`
	Deactivate      *Bool     `xmlrpc:"deactivate,omitempty"`
	DisplayName     *String   `xmlrpc:"display_name,omitempty"`
	Done            *Int      `xmlrpc:"done,omitempty"`
	Id              *Int      `xmlrpc:"id,omitempty"`
	Remaining       *Int      `xmlrpc:"remaining,omitempty"`
	TimedOutCounter *Int      `xmlrpc:"timed_out_counter,omitempty"`
	WriteDate       *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid        *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrCronProgresss represents array of ir.cron.progress model.
type IrCronProgresss []IrCronProgress

// IrCronProgressModel is the odoo model name.
const IrCronProgressModel = "ir.cron.progress"

// Many2One convert IrCronProgress to *Many2One.
func (icp *IrCronProgress) Many2One() *Many2One {
	return NewMany2One(icp.Id.Get(), "")
}

// CreateIrCronProgress creates a new ir.cron.progress model and returns its id.
func (c *Client) CreateIrCronProgress(icp *IrCronProgress) (int64, error) {
	ids, err := c.CreateIrCronProgresss([]*IrCronProgress{icp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrCronProgress creates a new ir.cron.progress model and returns its id.
func (c *Client) CreateIrCronProgresss(icps []*IrCronProgress) ([]int64, error) {
	var vv []interface{}
	for _, v := range icps {
		vv = append(vv, v)
	}
	return c.Create(IrCronProgressModel, vv, nil)
}

// UpdateIrCronProgress updates an existing ir.cron.progress record.
func (c *Client) UpdateIrCronProgress(icp *IrCronProgress) error {
	return c.UpdateIrCronProgresss([]int64{icp.Id.Get()}, icp)
}

// UpdateIrCronProgresss updates existing ir.cron.progress records.
// All records (represented by ids) will be updated by icp values.
func (c *Client) UpdateIrCronProgresss(ids []int64, icp *IrCronProgress) error {
	return c.Update(IrCronProgressModel, ids, icp, nil)
}

// DeleteIrCronProgress deletes an existing ir.cron.progress record.
func (c *Client) DeleteIrCronProgress(id int64) error {
	return c.DeleteIrCronProgresss([]int64{id})
}

// DeleteIrCronProgresss deletes existing ir.cron.progress records.
func (c *Client) DeleteIrCronProgresss(ids []int64) error {
	return c.Delete(IrCronProgressModel, ids)
}

// GetIrCronProgress gets ir.cron.progress existing record.
func (c *Client) GetIrCronProgress(id int64) (*IrCronProgress, error) {
	icps, err := c.GetIrCronProgresss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*icps)[0]), nil
}

// GetIrCronProgresss gets ir.cron.progress existing records.
func (c *Client) GetIrCronProgresss(ids []int64) (*IrCronProgresss, error) {
	icps := &IrCronProgresss{}
	if err := c.Read(IrCronProgressModel, ids, nil, icps); err != nil {
		return nil, err
	}
	return icps, nil
}

// FindIrCronProgress finds ir.cron.progress record by querying it with criteria.
func (c *Client) FindIrCronProgress(criteria *Criteria) (*IrCronProgress, error) {
	icps := &IrCronProgresss{}
	if err := c.SearchRead(IrCronProgressModel, criteria, NewOptions().Limit(1), icps); err != nil {
		return nil, err
	}
	return &((*icps)[0]), nil
}

// FindIrCronProgresss finds ir.cron.progress records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrCronProgresss(criteria *Criteria, options *Options) (*IrCronProgresss, error) {
	icps := &IrCronProgresss{}
	if err := c.SearchRead(IrCronProgressModel, criteria, options, icps); err != nil {
		return nil, err
	}
	return icps, nil
}

// FindIrCronProgressIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrCronProgressIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrCronProgressModel, criteria, options)
}

// FindIrCronProgressId finds record id by querying it with criteria.
func (c *Client) FindIrCronProgressId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrCronProgressModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
