package odoo

// CrmIapLeadRole represents crm.iap.lead.role model.
type CrmIapLeadRole struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	RevealId    *String   `xmlrpc:"reveal_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmIapLeadRoles represents array of crm.iap.lead.role model.
type CrmIapLeadRoles []CrmIapLeadRole

// CrmIapLeadRoleModel is the odoo model name.
const CrmIapLeadRoleModel = "crm.iap.lead.role"

// Many2One convert CrmIapLeadRole to *Many2One.
func (cilr *CrmIapLeadRole) Many2One() *Many2One {
	return NewMany2One(cilr.Id.Get(), "")
}

// CreateCrmIapLeadRole creates a new crm.iap.lead.role model and returns its id.
func (c *Client) CreateCrmIapLeadRole(cilr *CrmIapLeadRole) (int64, error) {
	ids, err := c.CreateCrmIapLeadRoles([]*CrmIapLeadRole{cilr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmIapLeadRole creates a new crm.iap.lead.role model and returns its id.
func (c *Client) CreateCrmIapLeadRoles(cilrs []*CrmIapLeadRole) ([]int64, error) {
	var vv []interface{}
	for _, v := range cilrs {
		vv = append(vv, v)
	}
	return c.Create(CrmIapLeadRoleModel, vv, nil)
}

// UpdateCrmIapLeadRole updates an existing crm.iap.lead.role record.
func (c *Client) UpdateCrmIapLeadRole(cilr *CrmIapLeadRole) error {
	return c.UpdateCrmIapLeadRoles([]int64{cilr.Id.Get()}, cilr)
}

// UpdateCrmIapLeadRoles updates existing crm.iap.lead.role records.
// All records (represented by ids) will be updated by cilr values.
func (c *Client) UpdateCrmIapLeadRoles(ids []int64, cilr *CrmIapLeadRole) error {
	return c.Update(CrmIapLeadRoleModel, ids, cilr, nil)
}

// DeleteCrmIapLeadRole deletes an existing crm.iap.lead.role record.
func (c *Client) DeleteCrmIapLeadRole(id int64) error {
	return c.DeleteCrmIapLeadRoles([]int64{id})
}

// DeleteCrmIapLeadRoles deletes existing crm.iap.lead.role records.
func (c *Client) DeleteCrmIapLeadRoles(ids []int64) error {
	return c.Delete(CrmIapLeadRoleModel, ids)
}

// GetCrmIapLeadRole gets crm.iap.lead.role existing record.
func (c *Client) GetCrmIapLeadRole(id int64) (*CrmIapLeadRole, error) {
	cilrs, err := c.GetCrmIapLeadRoles([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cilrs)[0]), nil
}

// GetCrmIapLeadRoles gets crm.iap.lead.role existing records.
func (c *Client) GetCrmIapLeadRoles(ids []int64) (*CrmIapLeadRoles, error) {
	cilrs := &CrmIapLeadRoles{}
	if err := c.Read(CrmIapLeadRoleModel, ids, nil, cilrs); err != nil {
		return nil, err
	}
	return cilrs, nil
}

// FindCrmIapLeadRole finds crm.iap.lead.role record by querying it with criteria.
func (c *Client) FindCrmIapLeadRole(criteria *Criteria) (*CrmIapLeadRole, error) {
	cilrs := &CrmIapLeadRoles{}
	if err := c.SearchRead(CrmIapLeadRoleModel, criteria, NewOptions().Limit(1), cilrs); err != nil {
		return nil, err
	}
	return &((*cilrs)[0]), nil
}

// FindCrmIapLeadRoles finds crm.iap.lead.role records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadRoles(criteria *Criteria, options *Options) (*CrmIapLeadRoles, error) {
	cilrs := &CrmIapLeadRoles{}
	if err := c.SearchRead(CrmIapLeadRoleModel, criteria, options, cilrs); err != nil {
		return nil, err
	}
	return cilrs, nil
}

// FindCrmIapLeadRoleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadRoleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmIapLeadRoleModel, criteria, options)
}

// FindCrmIapLeadRoleId finds record id by querying it with criteria.
func (c *Client) FindCrmIapLeadRoleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmIapLeadRoleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
