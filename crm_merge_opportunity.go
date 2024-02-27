package odoo

// CrmMergeOpportunity represents crm.merge.opportunity model.
type CrmMergeOpportunity struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	OpportunityIds *Relation `xmlrpc:"opportunity_ids,omitempty"`
	TeamId         *Many2One `xmlrpc:"team_id,omitempty"`
	UserId         *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmMergeOpportunitys represents array of crm.merge.opportunity model.
type CrmMergeOpportunitys []CrmMergeOpportunity

// CrmMergeOpportunityModel is the odoo model name.
const CrmMergeOpportunityModel = "crm.merge.opportunity"

// Many2One convert CrmMergeOpportunity to *Many2One.
func (cmo *CrmMergeOpportunity) Many2One() *Many2One {
	return NewMany2One(cmo.Id.Get(), "")
}

// CreateCrmMergeOpportunity creates a new crm.merge.opportunity model and returns its id.
func (c *Client) CreateCrmMergeOpportunity(cmo *CrmMergeOpportunity) (int64, error) {
	ids, err := c.CreateCrmMergeOpportunitys([]*CrmMergeOpportunity{cmo})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmMergeOpportunity creates a new crm.merge.opportunity model and returns its id.
func (c *Client) CreateCrmMergeOpportunitys(cmos []*CrmMergeOpportunity) ([]int64, error) {
	var vv []interface{}
	for _, v := range cmos {
		vv = append(vv, v)
	}
	return c.Create(CrmMergeOpportunityModel, vv, nil)
}

// UpdateCrmMergeOpportunity updates an existing crm.merge.opportunity record.
func (c *Client) UpdateCrmMergeOpportunity(cmo *CrmMergeOpportunity) error {
	return c.UpdateCrmMergeOpportunitys([]int64{cmo.Id.Get()}, cmo)
}

// UpdateCrmMergeOpportunitys updates existing crm.merge.opportunity records.
// All records (represented by ids) will be updated by cmo values.
func (c *Client) UpdateCrmMergeOpportunitys(ids []int64, cmo *CrmMergeOpportunity) error {
	return c.Update(CrmMergeOpportunityModel, ids, cmo, nil)
}

// DeleteCrmMergeOpportunity deletes an existing crm.merge.opportunity record.
func (c *Client) DeleteCrmMergeOpportunity(id int64) error {
	return c.DeleteCrmMergeOpportunitys([]int64{id})
}

// DeleteCrmMergeOpportunitys deletes existing crm.merge.opportunity records.
func (c *Client) DeleteCrmMergeOpportunitys(ids []int64) error {
	return c.Delete(CrmMergeOpportunityModel, ids)
}

// GetCrmMergeOpportunity gets crm.merge.opportunity existing record.
func (c *Client) GetCrmMergeOpportunity(id int64) (*CrmMergeOpportunity, error) {
	cmos, err := c.GetCrmMergeOpportunitys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cmos)[0]), nil
}

// GetCrmMergeOpportunitys gets crm.merge.opportunity existing records.
func (c *Client) GetCrmMergeOpportunitys(ids []int64) (*CrmMergeOpportunitys, error) {
	cmos := &CrmMergeOpportunitys{}
	if err := c.Read(CrmMergeOpportunityModel, ids, nil, cmos); err != nil {
		return nil, err
	}
	return cmos, nil
}

// FindCrmMergeOpportunity finds crm.merge.opportunity record by querying it with criteria.
func (c *Client) FindCrmMergeOpportunity(criteria *Criteria) (*CrmMergeOpportunity, error) {
	cmos := &CrmMergeOpportunitys{}
	if err := c.SearchRead(CrmMergeOpportunityModel, criteria, NewOptions().Limit(1), cmos); err != nil {
		return nil, err
	}
	return &((*cmos)[0]), nil
}

// FindCrmMergeOpportunitys finds crm.merge.opportunity records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmMergeOpportunitys(criteria *Criteria, options *Options) (*CrmMergeOpportunitys, error) {
	cmos := &CrmMergeOpportunitys{}
	if err := c.SearchRead(CrmMergeOpportunityModel, criteria, options, cmos); err != nil {
		return nil, err
	}
	return cmos, nil
}

// FindCrmMergeOpportunityIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmMergeOpportunityIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmMergeOpportunityModel, criteria, options)
}

// FindCrmMergeOpportunityId finds record id by querying it with criteria.
func (c *Client) FindCrmMergeOpportunityId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmMergeOpportunityModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
