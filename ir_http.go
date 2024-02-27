package odoo

// IrHttp represents ir.http model.
type IrHttp struct {
	LastUpdate  *Time   `xmlrpc:"__last_update,omitempty"`
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrHttps represents array of ir.http model.
type IrHttps []IrHttp

// IrHttpModel is the odoo model name.
const IrHttpModel = "ir.http"

// Many2One convert IrHttp to *Many2One.
func (ih *IrHttp) Many2One() *Many2One {
	return NewMany2One(ih.Id.Get(), "")
}

// CreateIrHttp creates a new ir.http model and returns its id.
func (c *Client) CreateIrHttp(ih *IrHttp) (int64, error) {
	ids, err := c.CreateIrHttps([]*IrHttp{ih})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrHttp creates a new ir.http model and returns its id.
func (c *Client) CreateIrHttps(ihs []*IrHttp) ([]int64, error) {
	var vv []interface{}
	for _, v := range ihs {
		vv = append(vv, v)
	}
	return c.Create(IrHttpModel, vv, nil)
}

// UpdateIrHttp updates an existing ir.http record.
func (c *Client) UpdateIrHttp(ih *IrHttp) error {
	return c.UpdateIrHttps([]int64{ih.Id.Get()}, ih)
}

// UpdateIrHttps updates existing ir.http records.
// All records (represented by ids) will be updated by ih values.
func (c *Client) UpdateIrHttps(ids []int64, ih *IrHttp) error {
	return c.Update(IrHttpModel, ids, ih, nil)
}

// DeleteIrHttp deletes an existing ir.http record.
func (c *Client) DeleteIrHttp(id int64) error {
	return c.DeleteIrHttps([]int64{id})
}

// DeleteIrHttps deletes existing ir.http records.
func (c *Client) DeleteIrHttps(ids []int64) error {
	return c.Delete(IrHttpModel, ids)
}

// GetIrHttp gets ir.http existing record.
func (c *Client) GetIrHttp(id int64) (*IrHttp, error) {
	ihs, err := c.GetIrHttps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ihs)[0]), nil
}

// GetIrHttps gets ir.http existing records.
func (c *Client) GetIrHttps(ids []int64) (*IrHttps, error) {
	ihs := &IrHttps{}
	if err := c.Read(IrHttpModel, ids, nil, ihs); err != nil {
		return nil, err
	}
	return ihs, nil
}

// FindIrHttp finds ir.http record by querying it with criteria.
func (c *Client) FindIrHttp(criteria *Criteria) (*IrHttp, error) {
	ihs := &IrHttps{}
	if err := c.SearchRead(IrHttpModel, criteria, NewOptions().Limit(1), ihs); err != nil {
		return nil, err
	}
	return &((*ihs)[0]), nil
}

// FindIrHttps finds ir.http records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrHttps(criteria *Criteria, options *Options) (*IrHttps, error) {
	ihs := &IrHttps{}
	if err := c.SearchRead(IrHttpModel, criteria, options, ihs); err != nil {
		return nil, err
	}
	return ihs, nil
}

// FindIrHttpIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrHttpIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrHttpModel, criteria, options)
}

// FindIrHttpId finds record id by querying it with criteria.
func (c *Client) FindIrHttpId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrHttpModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
