package odoo

// CrmLeadPlsUpdate represents crm.lead.pls.update model.
type CrmLeadPlsUpdate struct {
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	PlsFields    *Relation `xmlrpc:"pls_fields,omitempty"`
	PlsStartDate *Time     `xmlrpc:"pls_start_date,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmLeadPlsUpdates represents array of crm.lead.pls.update model.
type CrmLeadPlsUpdates []CrmLeadPlsUpdate

// CrmLeadPlsUpdateModel is the odoo model name.
const CrmLeadPlsUpdateModel = "crm.lead.pls.update"

// Many2One convert CrmLeadPlsUpdate to *Many2One.
func (clpu *CrmLeadPlsUpdate) Many2One() *Many2One {
	return NewMany2One(clpu.Id.Get(), "")
}

// CreateCrmLeadPlsUpdate creates a new crm.lead.pls.update model and returns its id.
func (c *Client) CreateCrmLeadPlsUpdate(clpu *CrmLeadPlsUpdate) (int64, error) {
	ids, err := c.CreateCrmLeadPlsUpdates([]*CrmLeadPlsUpdate{clpu})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmLeadPlsUpdate creates a new crm.lead.pls.update model and returns its id.
func (c *Client) CreateCrmLeadPlsUpdates(clpus []*CrmLeadPlsUpdate) ([]int64, error) {
	var vv []interface{}
	for _, v := range clpus {
		vv = append(vv, v)
	}
	return c.Create(CrmLeadPlsUpdateModel, vv, nil)
}

// UpdateCrmLeadPlsUpdate updates an existing crm.lead.pls.update record.
func (c *Client) UpdateCrmLeadPlsUpdate(clpu *CrmLeadPlsUpdate) error {
	return c.UpdateCrmLeadPlsUpdates([]int64{clpu.Id.Get()}, clpu)
}

// UpdateCrmLeadPlsUpdates updates existing crm.lead.pls.update records.
// All records (represented by ids) will be updated by clpu values.
func (c *Client) UpdateCrmLeadPlsUpdates(ids []int64, clpu *CrmLeadPlsUpdate) error {
	return c.Update(CrmLeadPlsUpdateModel, ids, clpu, nil)
}

// DeleteCrmLeadPlsUpdate deletes an existing crm.lead.pls.update record.
func (c *Client) DeleteCrmLeadPlsUpdate(id int64) error {
	return c.DeleteCrmLeadPlsUpdates([]int64{id})
}

// DeleteCrmLeadPlsUpdates deletes existing crm.lead.pls.update records.
func (c *Client) DeleteCrmLeadPlsUpdates(ids []int64) error {
	return c.Delete(CrmLeadPlsUpdateModel, ids)
}

// GetCrmLeadPlsUpdate gets crm.lead.pls.update existing record.
func (c *Client) GetCrmLeadPlsUpdate(id int64) (*CrmLeadPlsUpdate, error) {
	clpus, err := c.GetCrmLeadPlsUpdates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*clpus)[0]), nil
}

// GetCrmLeadPlsUpdates gets crm.lead.pls.update existing records.
func (c *Client) GetCrmLeadPlsUpdates(ids []int64) (*CrmLeadPlsUpdates, error) {
	clpus := &CrmLeadPlsUpdates{}
	if err := c.Read(CrmLeadPlsUpdateModel, ids, nil, clpus); err != nil {
		return nil, err
	}
	return clpus, nil
}

// FindCrmLeadPlsUpdate finds crm.lead.pls.update record by querying it with criteria.
func (c *Client) FindCrmLeadPlsUpdate(criteria *Criteria) (*CrmLeadPlsUpdate, error) {
	clpus := &CrmLeadPlsUpdates{}
	if err := c.SearchRead(CrmLeadPlsUpdateModel, criteria, NewOptions().Limit(1), clpus); err != nil {
		return nil, err
	}
	return &((*clpus)[0]), nil
}

// FindCrmLeadPlsUpdates finds crm.lead.pls.update records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadPlsUpdates(criteria *Criteria, options *Options) (*CrmLeadPlsUpdates, error) {
	clpus := &CrmLeadPlsUpdates{}
	if err := c.SearchRead(CrmLeadPlsUpdateModel, criteria, options, clpus); err != nil {
		return nil, err
	}
	return clpus, nil
}

// FindCrmLeadPlsUpdateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadPlsUpdateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmLeadPlsUpdateModel, criteria, options)
}

// FindCrmLeadPlsUpdateId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadPlsUpdateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadPlsUpdateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
