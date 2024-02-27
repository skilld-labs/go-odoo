package odoo

// IrConfigParameter represents ir.config_parameter model.
type IrConfigParameter struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Key         *String   `xmlrpc:"key,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrConfigParameters represents array of ir.config_parameter model.
type IrConfigParameters []IrConfigParameter

// IrConfigParameterModel is the odoo model name.
const IrConfigParameterModel = "ir.config_parameter"

// Many2One convert IrConfigParameter to *Many2One.
func (ic *IrConfigParameter) Many2One() *Many2One {
	return NewMany2One(ic.Id.Get(), "")
}

// CreateIrConfigParameter creates a new ir.config_parameter model and returns its id.
func (c *Client) CreateIrConfigParameter(ic *IrConfigParameter) (int64, error) {
	ids, err := c.CreateIrConfigParameters([]*IrConfigParameter{ic})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrConfigParameter creates a new ir.config_parameter model and returns its id.
func (c *Client) CreateIrConfigParameters(ics []*IrConfigParameter) ([]int64, error) {
	var vv []interface{}
	for _, v := range ics {
		vv = append(vv, v)
	}
	return c.Create(IrConfigParameterModel, vv, nil)
}

// UpdateIrConfigParameter updates an existing ir.config_parameter record.
func (c *Client) UpdateIrConfigParameter(ic *IrConfigParameter) error {
	return c.UpdateIrConfigParameters([]int64{ic.Id.Get()}, ic)
}

// UpdateIrConfigParameters updates existing ir.config_parameter records.
// All records (represented by ids) will be updated by ic values.
func (c *Client) UpdateIrConfigParameters(ids []int64, ic *IrConfigParameter) error {
	return c.Update(IrConfigParameterModel, ids, ic, nil)
}

// DeleteIrConfigParameter deletes an existing ir.config_parameter record.
func (c *Client) DeleteIrConfigParameter(id int64) error {
	return c.DeleteIrConfigParameters([]int64{id})
}

// DeleteIrConfigParameters deletes existing ir.config_parameter records.
func (c *Client) DeleteIrConfigParameters(ids []int64) error {
	return c.Delete(IrConfigParameterModel, ids)
}

// GetIrConfigParameter gets ir.config_parameter existing record.
func (c *Client) GetIrConfigParameter(id int64) (*IrConfigParameter, error) {
	ics, err := c.GetIrConfigParameters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ics)[0]), nil
}

// GetIrConfigParameters gets ir.config_parameter existing records.
func (c *Client) GetIrConfigParameters(ids []int64) (*IrConfigParameters, error) {
	ics := &IrConfigParameters{}
	if err := c.Read(IrConfigParameterModel, ids, nil, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindIrConfigParameter finds ir.config_parameter record by querying it with criteria.
func (c *Client) FindIrConfigParameter(criteria *Criteria) (*IrConfigParameter, error) {
	ics := &IrConfigParameters{}
	if err := c.SearchRead(IrConfigParameterModel, criteria, NewOptions().Limit(1), ics); err != nil {
		return nil, err
	}
	return &((*ics)[0]), nil
}

// FindIrConfigParameters finds ir.config_parameter records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrConfigParameters(criteria *Criteria, options *Options) (*IrConfigParameters, error) {
	ics := &IrConfigParameters{}
	if err := c.SearchRead(IrConfigParameterModel, criteria, options, ics); err != nil {
		return nil, err
	}
	return ics, nil
}

// FindIrConfigParameterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrConfigParameterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrConfigParameterModel, criteria, options)
}

// FindIrConfigParameterId finds record id by querying it with criteria.
func (c *Client) FindIrConfigParameterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrConfigParameterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
