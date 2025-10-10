package odoo

// IrDemoFailure represents ir.demo_failure model.
type IrDemoFailure struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Error       *String   `xmlrpc:"error,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	ModuleId    *Many2One `xmlrpc:"module_id,omitempty"`
	WizardId    *Many2One `xmlrpc:"wizard_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// IrDemoFailures represents array of ir.demo_failure model.
type IrDemoFailures []IrDemoFailure

// IrDemoFailureModel is the odoo model name.
const IrDemoFailureModel = "ir.demo_failure"

// Many2One convert IrDemoFailure to *Many2One.
func (ID *IrDemoFailure) Many2One() *Many2One {
	return NewMany2One(ID.Id.Get(), "")
}

// CreateIrDemoFailure creates a new ir.demo_failure model and returns its id.
func (c *Client) CreateIrDemoFailure(ID *IrDemoFailure) (int64, error) {
	ids, err := c.CreateIrDemoFailures([]*IrDemoFailure{ID})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrDemoFailure creates a new ir.demo_failure model and returns its id.
func (c *Client) CreateIrDemoFailures(IDs []*IrDemoFailure) ([]int64, error) {
	var vv []interface{}
	for _, v := range IDs {
		vv = append(vv, v)
	}
	return c.Create(IrDemoFailureModel, vv, nil)
}

// UpdateIrDemoFailure updates an existing ir.demo_failure record.
func (c *Client) UpdateIrDemoFailure(ID *IrDemoFailure) error {
	return c.UpdateIrDemoFailures([]int64{ID.Id.Get()}, ID)
}

// UpdateIrDemoFailures updates existing ir.demo_failure records.
// All records (represented by ids) will be updated by ID values.
func (c *Client) UpdateIrDemoFailures(ids []int64, ID *IrDemoFailure) error {
	return c.Update(IrDemoFailureModel, ids, ID, nil)
}

// DeleteIrDemoFailure deletes an existing ir.demo_failure record.
func (c *Client) DeleteIrDemoFailure(id int64) error {
	return c.DeleteIrDemoFailures([]int64{id})
}

// DeleteIrDemoFailures deletes existing ir.demo_failure records.
func (c *Client) DeleteIrDemoFailures(ids []int64) error {
	return c.Delete(IrDemoFailureModel, ids)
}

// GetIrDemoFailure gets ir.demo_failure existing record.
func (c *Client) GetIrDemoFailure(id int64) (*IrDemoFailure, error) {
	IDs, err := c.GetIrDemoFailures([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*IDs)[0]), nil
}

// GetIrDemoFailures gets ir.demo_failure existing records.
func (c *Client) GetIrDemoFailures(ids []int64) (*IrDemoFailures, error) {
	IDs := &IrDemoFailures{}
	if err := c.Read(IrDemoFailureModel, ids, nil, IDs); err != nil {
		return nil, err
	}
	return IDs, nil
}

// FindIrDemoFailure finds ir.demo_failure record by querying it with criteria.
func (c *Client) FindIrDemoFailure(criteria *Criteria) (*IrDemoFailure, error) {
	IDs := &IrDemoFailures{}
	if err := c.SearchRead(IrDemoFailureModel, criteria, NewOptions().Limit(1), IDs); err != nil {
		return nil, err
	}
	return &((*IDs)[0]), nil
}

// FindIrDemoFailures finds ir.demo_failure records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDemoFailures(criteria *Criteria, options *Options) (*IrDemoFailures, error) {
	IDs := &IrDemoFailures{}
	if err := c.SearchRead(IrDemoFailureModel, criteria, options, IDs); err != nil {
		return nil, err
	}
	return IDs, nil
}

// FindIrDemoFailureIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrDemoFailureIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrDemoFailureModel, criteria, options)
}

// FindIrDemoFailureId finds record id by querying it with criteria.
func (c *Client) FindIrDemoFailureId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrDemoFailureModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
