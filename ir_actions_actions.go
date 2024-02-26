package odoo

// IrActionsActions represents ir.actions.actions model.
type IrActionsActions struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	BindingModelId *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType    *Selection `xmlrpc:"binding_type,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Help           *String    `xmlrpc:"help,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	Type           *String    `xmlrpc:"type,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId          *String    `xmlrpc:"xml_id,omptempty"`
}

// IrActionsActionss represents array of ir.actions.actions model.
type IrActionsActionss []IrActionsActions

// IrActionsActionsModel is the odoo model name.
const IrActionsActionsModel = "ir.actions.actions"

// Many2One convert IrActionsActions to *Many2One.
func (iaa *IrActionsActions) Many2One() *Many2One {
	return NewMany2One(iaa.Id.Get(), "")
}

// CreateIrActionsActions creates a new ir.actions.actions model and returns its id.
func (c *Client) CreateIrActionsActions(iaa *IrActionsActions) (int64, error) {
	ids, err := c.CreateIrActionsActionss([]*IrActionsActions{iaa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrActionsActions creates a new ir.actions.actions model and returns its id.
func (c *Client) CreateIrActionsActionss(iaas []*IrActionsActions) ([]int64, error) {
	var vv []interface{}
	for _, v := range iaas {
		vv = append(vv, v)
	}
	return c.Create(IrActionsActionsModel, vv, nil)
}

// UpdateIrActionsActions updates an existing ir.actions.actions record.
func (c *Client) UpdateIrActionsActions(iaa *IrActionsActions) error {
	return c.UpdateIrActionsActionss([]int64{iaa.Id.Get()}, iaa)
}

// UpdateIrActionsActionss updates existing ir.actions.actions records.
// All records (represented by ids) will be updated by iaa values.
func (c *Client) UpdateIrActionsActionss(ids []int64, iaa *IrActionsActions) error {
	return c.Update(IrActionsActionsModel, ids, iaa, nil)
}

// DeleteIrActionsActions deletes an existing ir.actions.actions record.
func (c *Client) DeleteIrActionsActions(id int64) error {
	return c.DeleteIrActionsActionss([]int64{id})
}

// DeleteIrActionsActionss deletes existing ir.actions.actions records.
func (c *Client) DeleteIrActionsActionss(ids []int64) error {
	return c.Delete(IrActionsActionsModel, ids)
}

// GetIrActionsActions gets ir.actions.actions existing record.
func (c *Client) GetIrActionsActions(id int64) (*IrActionsActions, error) {
	iaas, err := c.GetIrActionsActionss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iaas)[0]), nil
}

// GetIrActionsActionss gets ir.actions.actions existing records.
func (c *Client) GetIrActionsActionss(ids []int64) (*IrActionsActionss, error) {
	iaas := &IrActionsActionss{}
	if err := c.Read(IrActionsActionsModel, ids, nil, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActions finds ir.actions.actions record by querying it with criteria.
func (c *Client) FindIrActionsActions(criteria *Criteria) (*IrActionsActions, error) {
	iaas := &IrActionsActionss{}
	if err := c.SearchRead(IrActionsActionsModel, criteria, NewOptions().Limit(1), iaas); err != nil {
		return nil, err
	}
	return &((*iaas)[0]), nil
}

// FindIrActionsActionss finds ir.actions.actions records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActionss(criteria *Criteria, options *Options) (*IrActionsActionss, error) {
	iaas := &IrActionsActionss{}
	if err := c.SearchRead(IrActionsActionsModel, criteria, options, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActionsIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActionsIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrActionsActionsModel, criteria, options)
}

// FindIrActionsActionsId finds record id by querying it with criteria.
func (c *Client) FindIrActionsActionsId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsActionsModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
