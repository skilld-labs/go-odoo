package odoo

// IrActionsServer represents ir.actions.server model.
type IrActionsServer struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	BindingModelId *Many2One  `xmlrpc:"binding_model_id,omitempty"`
	BindingType    *Selection `xmlrpc:"binding_type,omitempty"`
	ChannelIds     *Relation  `xmlrpc:"channel_ids,omitempty"`
	ChildIds       *Relation  `xmlrpc:"child_ids,omitempty"`
	Code           *String    `xmlrpc:"code,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	CrudModelId    *Many2One  `xmlrpc:"crud_model_id,omitempty"`
	CrudModelName  *String    `xmlrpc:"crud_model_name,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	FieldsLines    *Relation  `xmlrpc:"fields_lines,omitempty"`
	Help           *String    `xmlrpc:"help,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	LinkFieldId    *Many2One  `xmlrpc:"link_field_id,omitempty"`
	ModelId        *Many2One  `xmlrpc:"model_id,omitempty"`
	ModelName      *String    `xmlrpc:"model_name,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	PartnerIds     *Relation  `xmlrpc:"partner_ids,omitempty"`
	Sequence       *Int       `xmlrpc:"sequence,omitempty"`
	State          *Selection `xmlrpc:"state,omitempty"`
	TemplateId     *Many2One  `xmlrpc:"template_id,omitempty"`
	Type           *String    `xmlrpc:"type,omitempty"`
	Usage          *Selection `xmlrpc:"usage,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId          *String    `xmlrpc:"xml_id,omitempty"`
}

// IrActionsServers represents array of ir.actions.server model.
type IrActionsServers []IrActionsServer

// IrActionsServerModel is the odoo model name.
const IrActionsServerModel = "ir.actions.server"

// Many2One convert IrActionsServer to *Many2One.
func (ias *IrActionsServer) Many2One() *Many2One {
	return NewMany2One(ias.Id.Get(), "")
}

// CreateIrActionsServer creates a new ir.actions.server model and returns its id.
func (c *Client) CreateIrActionsServer(ias *IrActionsServer) (int64, error) {
	ids, err := c.CreateIrActionsServers([]*IrActionsServer{ias})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrActionsServer creates a new ir.actions.server model and returns its id.
func (c *Client) CreateIrActionsServers(iass []*IrActionsServer) ([]int64, error) {
	var vv []interface{}
	for _, v := range iass {
		vv = append(vv, v)
	}
	return c.Create(IrActionsServerModel, vv, nil)
}

// UpdateIrActionsServer updates an existing ir.actions.server record.
func (c *Client) UpdateIrActionsServer(ias *IrActionsServer) error {
	return c.UpdateIrActionsServers([]int64{ias.Id.Get()}, ias)
}

// UpdateIrActionsServers updates existing ir.actions.server records.
// All records (represented by ids) will be updated by ias values.
func (c *Client) UpdateIrActionsServers(ids []int64, ias *IrActionsServer) error {
	return c.Update(IrActionsServerModel, ids, ias, nil)
}

// DeleteIrActionsServer deletes an existing ir.actions.server record.
func (c *Client) DeleteIrActionsServer(id int64) error {
	return c.DeleteIrActionsServers([]int64{id})
}

// DeleteIrActionsServers deletes existing ir.actions.server records.
func (c *Client) DeleteIrActionsServers(ids []int64) error {
	return c.Delete(IrActionsServerModel, ids)
}

// GetIrActionsServer gets ir.actions.server existing record.
func (c *Client) GetIrActionsServer(id int64) (*IrActionsServer, error) {
	iass, err := c.GetIrActionsServers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iass)[0]), nil
}

// GetIrActionsServers gets ir.actions.server existing records.
func (c *Client) GetIrActionsServers(ids []int64) (*IrActionsServers, error) {
	iass := &IrActionsServers{}
	if err := c.Read(IrActionsServerModel, ids, nil, iass); err != nil {
		return nil, err
	}
	return iass, nil
}

// FindIrActionsServer finds ir.actions.server record by querying it with criteria.
func (c *Client) FindIrActionsServer(criteria *Criteria) (*IrActionsServer, error) {
	iass := &IrActionsServers{}
	if err := c.SearchRead(IrActionsServerModel, criteria, NewOptions().Limit(1), iass); err != nil {
		return nil, err
	}
	return &((*iass)[0]), nil
}

// FindIrActionsServers finds ir.actions.server records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsServers(criteria *Criteria, options *Options) (*IrActionsServers, error) {
	iass := &IrActionsServers{}
	if err := c.SearchRead(IrActionsServerModel, criteria, options, iass); err != nil {
		return nil, err
	}
	return iass, nil
}

// FindIrActionsServerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsServerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrActionsServerModel, criteria, options)
}

// FindIrActionsServerId finds record id by querying it with criteria.
func (c *Client) FindIrActionsServerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsServerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
