package odoo

// IrActionsActWindow represents ir.actions.act_window model.
type IrActionsActWindow struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	AutoSearch     *Bool      `xmlrpc:"auto_search,omitempty"`
	BindingModelId *Many2One  `xmlrpc:"binding_model_id,omitempty"`
	BindingType    *Selection `xmlrpc:"binding_type,omitempty"`
	Context        *String    `xmlrpc:"context,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Domain         *String    `xmlrpc:"domain,omitempty"`
	Filter         *Bool      `xmlrpc:"filter,omitempty"`
	GroupsId       *Relation  `xmlrpc:"groups_id,omitempty"`
	Help           *String    `xmlrpc:"help,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	Limit          *Int       `xmlrpc:"limit,omitempty"`
	Multi          *Bool      `xmlrpc:"multi,omitempty"`
	Name           *String    `xmlrpc:"name,omitempty"`
	ResId          *Int       `xmlrpc:"res_id,omitempty"`
	ResModel       *String    `xmlrpc:"res_model,omitempty"`
	SearchView     *String    `xmlrpc:"search_view,omitempty"`
	SearchViewId   *Many2One  `xmlrpc:"search_view_id,omitempty"`
	SrcModel       *String    `xmlrpc:"src_model,omitempty"`
	Target         *Selection `xmlrpc:"target,omitempty"`
	Type           *String    `xmlrpc:"type,omitempty"`
	Usage          *String    `xmlrpc:"usage,omitempty"`
	ViewId         *Many2One  `xmlrpc:"view_id,omitempty"`
	ViewIds        *Relation  `xmlrpc:"view_ids,omitempty"`
	ViewMode       *String    `xmlrpc:"view_mode,omitempty"`
	ViewType       *Selection `xmlrpc:"view_type,omitempty"`
	Views          *String    `xmlrpc:"views,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
	XmlId          *String    `xmlrpc:"xml_id,omitempty"`
}

// IrActionsActWindows represents array of ir.actions.act_window model.
type IrActionsActWindows []IrActionsActWindow

// IrActionsActWindowModel is the odoo model name.
const IrActionsActWindowModel = "ir.actions.act_window"

// Many2One convert IrActionsActWindow to *Many2One.
func (iaa *IrActionsActWindow) Many2One() *Many2One {
	return NewMany2One(iaa.Id.Get(), "")
}

// CreateIrActionsActWindow creates a new ir.actions.act_window model and returns its id.
func (c *Client) CreateIrActionsActWindow(iaa *IrActionsActWindow) (int64, error) {
	ids, err := c.CreateIrActionsActWindows([]*IrActionsActWindow{iaa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrActionsActWindows creates a new ir.actions.act_window model and returns its id.
func (c *Client) CreateIrActionsActWindows(iaas []*IrActionsActWindow) ([]int64, error) {
	var vv []interface{}
	for _, v := range iaas {
		vv = append(vv, v)
	}
	return c.Create(IrActionsActWindowModel, vv, nil)
}

// UpdateIrActionsActWindow updates an existing ir.actions.act_window record.
func (c *Client) UpdateIrActionsActWindow(iaa *IrActionsActWindow) error {
	return c.UpdateIrActionsActWindows([]int64{iaa.Id.Get()}, iaa)
}

// UpdateIrActionsActWindows updates existing ir.actions.act_window records.
// All records (represented by ids) will be updated by iaa values.
func (c *Client) UpdateIrActionsActWindows(ids []int64, iaa *IrActionsActWindow) error {
	return c.Update(IrActionsActWindowModel, ids, iaa, nil)
}

// DeleteIrActionsActWindow deletes an existing ir.actions.act_window record.
func (c *Client) DeleteIrActionsActWindow(id int64) error {
	return c.DeleteIrActionsActWindows([]int64{id})
}

// DeleteIrActionsActWindows deletes existing ir.actions.act_window records.
func (c *Client) DeleteIrActionsActWindows(ids []int64) error {
	return c.Delete(IrActionsActWindowModel, ids)
}

// GetIrActionsActWindow gets ir.actions.act_window existing record.
func (c *Client) GetIrActionsActWindow(id int64) (*IrActionsActWindow, error) {
	iaas, err := c.GetIrActionsActWindows([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iaas)[0]), nil
}

// GetIrActionsActWindows gets ir.actions.act_window existing records.
func (c *Client) GetIrActionsActWindows(ids []int64) (*IrActionsActWindows, error) {
	iaas := &IrActionsActWindows{}
	if err := c.Read(IrActionsActWindowModel, ids, nil, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActWindow finds ir.actions.act_window record by querying it with criteria.
func (c *Client) FindIrActionsActWindow(criteria *Criteria) (*IrActionsActWindow, error) {
	iaas := &IrActionsActWindows{}
	if err := c.SearchRead(IrActionsActWindowModel, criteria, NewOptions().Limit(1), iaas); err != nil {
		return nil, err
	}
	return &((*iaas)[0]), nil
}

// FindIrActionsActWindows finds ir.actions.act_window records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActWindows(criteria *Criteria, options *Options) (*IrActionsActWindows, error) {
	iaas := &IrActionsActWindows{}
	if err := c.SearchRead(IrActionsActWindowModel, criteria, options, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActWindowIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActWindowIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrActionsActWindowModel, criteria, options)
}

// FindIrActionsActWindowId finds record id by querying it with criteria.
func (c *Client) FindIrActionsActWindowId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsActWindowModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
