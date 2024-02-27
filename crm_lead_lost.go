package odoo

// CrmLeadLost represents crm.lead.lost model.
type CrmLeadLost struct {
	LastUpdate   *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	LostReasonId *Many2One `xmlrpc:"lost_reason_id,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmLeadLosts represents array of crm.lead.lost model.
type CrmLeadLosts []CrmLeadLost

// CrmLeadLostModel is the odoo model name.
const CrmLeadLostModel = "crm.lead.lost"

// Many2One convert CrmLeadLost to *Many2One.
func (cll *CrmLeadLost) Many2One() *Many2One {
	return NewMany2One(cll.Id.Get(), "")
}

// CreateCrmLeadLost creates a new crm.lead.lost model and returns its id.
func (c *Client) CreateCrmLeadLost(cll *CrmLeadLost) (int64, error) {
	ids, err := c.CreateCrmLeadLosts([]*CrmLeadLost{cll})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmLeadLosts creates a new crm.lead.lost model and returns its id.
func (c *Client) CreateCrmLeadLosts(clls []*CrmLeadLost) ([]int64, error) {
	var vv []interface{}
	for _, v := range clls {
		vv = append(vv, v)
	}
	return c.Create(CrmLeadLostModel, vv, nil)
}

// UpdateCrmLeadLost updates an existing crm.lead.lost record.
func (c *Client) UpdateCrmLeadLost(cll *CrmLeadLost) error {
	return c.UpdateCrmLeadLosts([]int64{cll.Id.Get()}, cll)
}

// UpdateCrmLeadLosts updates existing crm.lead.lost records.
// All records (represented by ids) will be updated by cll values.
func (c *Client) UpdateCrmLeadLosts(ids []int64, cll *CrmLeadLost) error {
	return c.Update(CrmLeadLostModel, ids, cll, nil)
}

// DeleteCrmLeadLost deletes an existing crm.lead.lost record.
func (c *Client) DeleteCrmLeadLost(id int64) error {
	return c.DeleteCrmLeadLosts([]int64{id})
}

// DeleteCrmLeadLosts deletes existing crm.lead.lost records.
func (c *Client) DeleteCrmLeadLosts(ids []int64) error {
	return c.Delete(CrmLeadLostModel, ids)
}

// GetCrmLeadLost gets crm.lead.lost existing record.
func (c *Client) GetCrmLeadLost(id int64) (*CrmLeadLost, error) {
	clls, err := c.GetCrmLeadLosts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*clls)[0]), nil
}

// GetCrmLeadLosts gets crm.lead.lost existing records.
func (c *Client) GetCrmLeadLosts(ids []int64) (*CrmLeadLosts, error) {
	clls := &CrmLeadLosts{}
	if err := c.Read(CrmLeadLostModel, ids, nil, clls); err != nil {
		return nil, err
	}
	return clls, nil
}

// FindCrmLeadLost finds crm.lead.lost record by querying it with criteria.
func (c *Client) FindCrmLeadLost(criteria *Criteria) (*CrmLeadLost, error) {
	clls := &CrmLeadLosts{}
	if err := c.SearchRead(CrmLeadLostModel, criteria, NewOptions().Limit(1), clls); err != nil {
		return nil, err
	}
	return &((*clls)[0]), nil
}

// FindCrmLeadLosts finds crm.lead.lost records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadLosts(criteria *Criteria, options *Options) (*CrmLeadLosts, error) {
	clls := &CrmLeadLosts{}
	if err := c.SearchRead(CrmLeadLostModel, criteria, options, clls); err != nil {
		return nil, err
	}
	return clls, nil
}

// FindCrmLeadLostIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadLostIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmLeadLostModel, criteria, options)
}

// FindCrmLeadLostId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadLostId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadLostModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
