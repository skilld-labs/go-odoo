package odoo

// IrLogging represents ir.logging model.
type IrLogging struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Int       `xmlrpc:"create_uid,omitempty"`
	Dbname      *String    `xmlrpc:"dbname,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Func        *String    `xmlrpc:"func,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Level       *String    `xmlrpc:"level,omitempty"`
	Line        *String    `xmlrpc:"line,omitempty"`
	Message     *String    `xmlrpc:"message,omitempty"`
	Name        *String    `xmlrpc:"name,omitempty"`
	Path        *String    `xmlrpc:"path,omitempty"`
	Type        *Selection `xmlrpc:"type,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// IrLoggings represents array of ir.logging model.
type IrLoggings []IrLogging

// IrLoggingModel is the odoo model name.
const IrLoggingModel = "ir.logging"

// Many2One convert IrLogging to *Many2One.
func (il *IrLogging) Many2One() *Many2One {
	return NewMany2One(il.Id.Get(), "")
}

// CreateIrLogging creates a new ir.logging model and returns its id.
func (c *Client) CreateIrLogging(il *IrLogging) (int64, error) {
	ids, err := c.CreateIrLoggings([]*IrLogging{il})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrLogging creates a new ir.logging model and returns its id.
func (c *Client) CreateIrLoggings(ils []*IrLogging) ([]int64, error) {
	var vv []interface{}
	for _, v := range ils {
		vv = append(vv, v)
	}
	return c.Create(IrLoggingModel, vv, nil)
}

// UpdateIrLogging updates an existing ir.logging record.
func (c *Client) UpdateIrLogging(il *IrLogging) error {
	return c.UpdateIrLoggings([]int64{il.Id.Get()}, il)
}

// UpdateIrLoggings updates existing ir.logging records.
// All records (represented by ids) will be updated by il values.
func (c *Client) UpdateIrLoggings(ids []int64, il *IrLogging) error {
	return c.Update(IrLoggingModel, ids, il, nil)
}

// DeleteIrLogging deletes an existing ir.logging record.
func (c *Client) DeleteIrLogging(id int64) error {
	return c.DeleteIrLoggings([]int64{id})
}

// DeleteIrLoggings deletes existing ir.logging records.
func (c *Client) DeleteIrLoggings(ids []int64) error {
	return c.Delete(IrLoggingModel, ids)
}

// GetIrLogging gets ir.logging existing record.
func (c *Client) GetIrLogging(id int64) (*IrLogging, error) {
	ils, err := c.GetIrLoggings([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ils)[0]), nil
}

// GetIrLoggings gets ir.logging existing records.
func (c *Client) GetIrLoggings(ids []int64) (*IrLoggings, error) {
	ils := &IrLoggings{}
	if err := c.Read(IrLoggingModel, ids, nil, ils); err != nil {
		return nil, err
	}
	return ils, nil
}

// FindIrLogging finds ir.logging record by querying it with criteria.
func (c *Client) FindIrLogging(criteria *Criteria) (*IrLogging, error) {
	ils := &IrLoggings{}
	if err := c.SearchRead(IrLoggingModel, criteria, NewOptions().Limit(1), ils); err != nil {
		return nil, err
	}
	return &((*ils)[0]), nil
}

// FindIrLoggings finds ir.logging records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrLoggings(criteria *Criteria, options *Options) (*IrLoggings, error) {
	ils := &IrLoggings{}
	if err := c.SearchRead(IrLoggingModel, criteria, options, ils); err != nil {
		return nil, err
	}
	return ils, nil
}

// FindIrLoggingIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrLoggingIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrLoggingModel, criteria, options)
}

// FindIrLoggingId finds record id by querying it with criteria.
func (c *Client) FindIrLoggingId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrLoggingModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
