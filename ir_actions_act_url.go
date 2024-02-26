package odoo

// IrActionsActUrl represents ir.actions.act_url model.
type IrActionsActUrl struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omptempty"`
	BindingModelId *Many2One  `xmlrpc:"binding_model_id,omptempty"`
	BindingType    *Selection `xmlrpc:"binding_type,omptempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName    *String    `xmlrpc:"display_name,omptempty"`
	Help           *String    `xmlrpc:"help,omptempty"`
	Id             *Int       `xmlrpc:"id,omptempty"`
	Name           *String    `xmlrpc:"name,omptempty"`
	Target         *Selection `xmlrpc:"target,omptempty"`
	Type           *String    `xmlrpc:"type,omptempty"`
	Url            *String    `xmlrpc:"url,omptempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omptempty"`
	XmlId          *String    `xmlrpc:"xml_id,omptempty"`
}

// IrActionsActUrls represents array of ir.actions.act_url model.
type IrActionsActUrls []IrActionsActUrl

// IrActionsActUrlModel is the odoo model name.
const IrActionsActUrlModel = "ir.actions.act_url"

// Many2One convert IrActionsActUrl to *Many2One.
func (iaa *IrActionsActUrl) Many2One() *Many2One {
	return NewMany2One(iaa.Id.Get(), "")
}

// CreateIrActionsActUrl creates a new ir.actions.act_url model and returns its id.
func (c *Client) CreateIrActionsActUrl(iaa *IrActionsActUrl) (int64, error) {
	ids, err := c.CreateIrActionsActUrls([]*IrActionsActUrl{iaa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrActionsActUrl creates a new ir.actions.act_url model and returns its id.
func (c *Client) CreateIrActionsActUrls(iaas []*IrActionsActUrl) ([]int64, error) {
	var vv []interface{}
	for _, v := range iaas {
		vv = append(vv, v)
	}
	return c.Create(IrActionsActUrlModel, vv, nil)
}

// UpdateIrActionsActUrl updates an existing ir.actions.act_url record.
func (c *Client) UpdateIrActionsActUrl(iaa *IrActionsActUrl) error {
	return c.UpdateIrActionsActUrls([]int64{iaa.Id.Get()}, iaa)
}

// UpdateIrActionsActUrls updates existing ir.actions.act_url records.
// All records (represented by ids) will be updated by iaa values.
func (c *Client) UpdateIrActionsActUrls(ids []int64, iaa *IrActionsActUrl) error {
	return c.Update(IrActionsActUrlModel, ids, iaa, nil)
}

// DeleteIrActionsActUrl deletes an existing ir.actions.act_url record.
func (c *Client) DeleteIrActionsActUrl(id int64) error {
	return c.DeleteIrActionsActUrls([]int64{id})
}

// DeleteIrActionsActUrls deletes existing ir.actions.act_url records.
func (c *Client) DeleteIrActionsActUrls(ids []int64) error {
	return c.Delete(IrActionsActUrlModel, ids)
}

// GetIrActionsActUrl gets ir.actions.act_url existing record.
func (c *Client) GetIrActionsActUrl(id int64) (*IrActionsActUrl, error) {
	iaas, err := c.GetIrActionsActUrls([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iaas)[0]), nil
}

// GetIrActionsActUrls gets ir.actions.act_url existing records.
func (c *Client) GetIrActionsActUrls(ids []int64) (*IrActionsActUrls, error) {
	iaas := &IrActionsActUrls{}
	if err := c.Read(IrActionsActUrlModel, ids, nil, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActUrl finds ir.actions.act_url record by querying it with criteria.
func (c *Client) FindIrActionsActUrl(criteria *Criteria) (*IrActionsActUrl, error) {
	iaas := &IrActionsActUrls{}
	if err := c.SearchRead(IrActionsActUrlModel, criteria, NewOptions().Limit(1), iaas); err != nil {
		return nil, err
	}
	return &((*iaas)[0]), nil
}

// FindIrActionsActUrls finds ir.actions.act_url records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActUrls(criteria *Criteria, options *Options) (*IrActionsActUrls, error) {
	iaas := &IrActionsActUrls{}
	if err := c.SearchRead(IrActionsActUrlModel, criteria, options, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIrActionsActUrlIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrActionsActUrlIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrActionsActUrlModel, criteria, options)
}

// FindIrActionsActUrlId finds record id by querying it with criteria.
func (c *Client) FindIrActionsActUrlId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrActionsActUrlModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
