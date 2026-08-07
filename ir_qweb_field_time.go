package odoo

// IrQwebFieldTime represents ir.qweb.field.time model.
type IrQwebFieldTime struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldTimes represents array of ir.qweb.field.time model.
type IrQwebFieldTimes []IrQwebFieldTime

// IrQwebFieldTimeModel is the odoo model name.
const IrQwebFieldTimeModel = "ir.qweb.field.time"

// Many2One convert IrQwebFieldTime to *Many2One.
func (iqft *IrQwebFieldTime) Many2One() *Many2One {
	return NewMany2One(iqft.Id.Get(), "")
}

// CreateIrQwebFieldTime creates a new ir.qweb.field.time model and returns its id.
func (c *Client) CreateIrQwebFieldTime(iqft *IrQwebFieldTime) (int64, error) {
	ids, err := c.CreateIrQwebFieldTimes([]*IrQwebFieldTime{iqft})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldTime creates a new ir.qweb.field.time model and returns its id.
func (c *Client) CreateIrQwebFieldTimes(iqfts []*IrQwebFieldTime) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfts {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldTimeModel, vv, nil)
}

// UpdateIrQwebFieldTime updates an existing ir.qweb.field.time record.
func (c *Client) UpdateIrQwebFieldTime(iqft *IrQwebFieldTime) error {
	return c.UpdateIrQwebFieldTimes([]int64{iqft.Id.Get()}, iqft)
}

// UpdateIrQwebFieldTimes updates existing ir.qweb.field.time records.
// All records (represented by ids) will be updated by iqft values.
func (c *Client) UpdateIrQwebFieldTimes(ids []int64, iqft *IrQwebFieldTime) error {
	return c.Update(IrQwebFieldTimeModel, ids, iqft, nil)
}

// DeleteIrQwebFieldTime deletes an existing ir.qweb.field.time record.
func (c *Client) DeleteIrQwebFieldTime(id int64) error {
	return c.DeleteIrQwebFieldTimes([]int64{id})
}

// DeleteIrQwebFieldTimes deletes existing ir.qweb.field.time records.
func (c *Client) DeleteIrQwebFieldTimes(ids []int64) error {
	return c.Delete(IrQwebFieldTimeModel, ids)
}

// GetIrQwebFieldTime gets ir.qweb.field.time existing record.
func (c *Client) GetIrQwebFieldTime(id int64) (*IrQwebFieldTime, error) {
	iqfts, err := c.GetIrQwebFieldTimes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqfts)[0]), nil
}

// GetIrQwebFieldTimes gets ir.qweb.field.time existing records.
func (c *Client) GetIrQwebFieldTimes(ids []int64) (*IrQwebFieldTimes, error) {
	iqfts := &IrQwebFieldTimes{}
	if err := c.Read(IrQwebFieldTimeModel, ids, nil, iqfts); err != nil {
		return nil, err
	}
	return iqfts, nil
}

// FindIrQwebFieldTime finds ir.qweb.field.time record by querying it with criteria.
func (c *Client) FindIrQwebFieldTime(criteria *Criteria) (*IrQwebFieldTime, error) {
	iqfts := &IrQwebFieldTimes{}
	if err := c.SearchRead(IrQwebFieldTimeModel, criteria, NewOptions().Limit(1), iqfts); err != nil {
		return nil, err
	}
	return &((*iqfts)[0]), nil
}

// FindIrQwebFieldTimes finds ir.qweb.field.time records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldTimes(criteria *Criteria, options *Options) (*IrQwebFieldTimes, error) {
	iqfts := &IrQwebFieldTimes{}
	if err := c.SearchRead(IrQwebFieldTimeModel, criteria, options, iqfts); err != nil {
		return nil, err
	}
	return iqfts, nil
}

// FindIrQwebFieldTimeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldTimeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebFieldTimeModel, criteria, options)
}

// FindIrQwebFieldTimeId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldTimeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldTimeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
