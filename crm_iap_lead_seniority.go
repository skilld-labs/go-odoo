package odoo

// CrmIapLeadSeniority represents crm.iap.lead.seniority model.
type CrmIapLeadSeniority struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	RevealId    *String   `xmlrpc:"reveal_id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmIapLeadSenioritys represents array of crm.iap.lead.seniority model.
type CrmIapLeadSenioritys []CrmIapLeadSeniority

// CrmIapLeadSeniorityModel is the odoo model name.
const CrmIapLeadSeniorityModel = "crm.iap.lead.seniority"

// Many2One convert CrmIapLeadSeniority to *Many2One.
func (cils *CrmIapLeadSeniority) Many2One() *Many2One {
	return NewMany2One(cils.Id.Get(), "")
}

// CreateCrmIapLeadSeniority creates a new crm.iap.lead.seniority model and returns its id.
func (c *Client) CreateCrmIapLeadSeniority(cils *CrmIapLeadSeniority) (int64, error) {
	ids, err := c.CreateCrmIapLeadSenioritys([]*CrmIapLeadSeniority{cils})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmIapLeadSeniority creates a new crm.iap.lead.seniority model and returns its id.
func (c *Client) CreateCrmIapLeadSenioritys(cilss []*CrmIapLeadSeniority) ([]int64, error) {
	var vv []interface{}
	for _, v := range cilss {
		vv = append(vv, v)
	}
	return c.Create(CrmIapLeadSeniorityModel, vv, nil)
}

// UpdateCrmIapLeadSeniority updates an existing crm.iap.lead.seniority record.
func (c *Client) UpdateCrmIapLeadSeniority(cils *CrmIapLeadSeniority) error {
	return c.UpdateCrmIapLeadSenioritys([]int64{cils.Id.Get()}, cils)
}

// UpdateCrmIapLeadSenioritys updates existing crm.iap.lead.seniority records.
// All records (represented by ids) will be updated by cils values.
func (c *Client) UpdateCrmIapLeadSenioritys(ids []int64, cils *CrmIapLeadSeniority) error {
	return c.Update(CrmIapLeadSeniorityModel, ids, cils, nil)
}

// DeleteCrmIapLeadSeniority deletes an existing crm.iap.lead.seniority record.
func (c *Client) DeleteCrmIapLeadSeniority(id int64) error {
	return c.DeleteCrmIapLeadSenioritys([]int64{id})
}

// DeleteCrmIapLeadSenioritys deletes existing crm.iap.lead.seniority records.
func (c *Client) DeleteCrmIapLeadSenioritys(ids []int64) error {
	return c.Delete(CrmIapLeadSeniorityModel, ids)
}

// GetCrmIapLeadSeniority gets crm.iap.lead.seniority existing record.
func (c *Client) GetCrmIapLeadSeniority(id int64) (*CrmIapLeadSeniority, error) {
	cilss, err := c.GetCrmIapLeadSenioritys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cilss)[0]), nil
}

// GetCrmIapLeadSenioritys gets crm.iap.lead.seniority existing records.
func (c *Client) GetCrmIapLeadSenioritys(ids []int64) (*CrmIapLeadSenioritys, error) {
	cilss := &CrmIapLeadSenioritys{}
	if err := c.Read(CrmIapLeadSeniorityModel, ids, nil, cilss); err != nil {
		return nil, err
	}
	return cilss, nil
}

// FindCrmIapLeadSeniority finds crm.iap.lead.seniority record by querying it with criteria.
func (c *Client) FindCrmIapLeadSeniority(criteria *Criteria) (*CrmIapLeadSeniority, error) {
	cilss := &CrmIapLeadSenioritys{}
	if err := c.SearchRead(CrmIapLeadSeniorityModel, criteria, NewOptions().Limit(1), cilss); err != nil {
		return nil, err
	}
	return &((*cilss)[0]), nil
}

// FindCrmIapLeadSenioritys finds crm.iap.lead.seniority records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadSenioritys(criteria *Criteria, options *Options) (*CrmIapLeadSenioritys, error) {
	cilss := &CrmIapLeadSenioritys{}
	if err := c.SearchRead(CrmIapLeadSeniorityModel, criteria, options, cilss); err != nil {
		return nil, err
	}
	return cilss, nil
}

// FindCrmIapLeadSeniorityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadSeniorityIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmIapLeadSeniorityModel, criteria, options)
}

// FindCrmIapLeadSeniorityId finds record id by querying it with criteria.
func (c *Client) FindCrmIapLeadSeniorityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmIapLeadSeniorityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
