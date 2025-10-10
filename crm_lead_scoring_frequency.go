package odoo

// CrmLeadScoringFrequency represents crm.lead.scoring.frequency model.
type CrmLeadScoringFrequency struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	LostCount   *Float    `xmlrpc:"lost_count,omitempty"`
	TeamId      *Many2One `xmlrpc:"team_id,omitempty"`
	Value       *String   `xmlrpc:"value,omitempty"`
	Variable    *String   `xmlrpc:"variable,omitempty"`
	WonCount    *Float    `xmlrpc:"won_count,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmLeadScoringFrequencys represents array of crm.lead.scoring.frequency model.
type CrmLeadScoringFrequencys []CrmLeadScoringFrequency

// CrmLeadScoringFrequencyModel is the odoo model name.
const CrmLeadScoringFrequencyModel = "crm.lead.scoring.frequency"

// Many2One convert CrmLeadScoringFrequency to *Many2One.
func (clsf *CrmLeadScoringFrequency) Many2One() *Many2One {
	return NewMany2One(clsf.Id.Get(), "")
}

// CreateCrmLeadScoringFrequency creates a new crm.lead.scoring.frequency model and returns its id.
func (c *Client) CreateCrmLeadScoringFrequency(clsf *CrmLeadScoringFrequency) (int64, error) {
	ids, err := c.CreateCrmLeadScoringFrequencys([]*CrmLeadScoringFrequency{clsf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmLeadScoringFrequency creates a new crm.lead.scoring.frequency model and returns its id.
func (c *Client) CreateCrmLeadScoringFrequencys(clsfs []*CrmLeadScoringFrequency) ([]int64, error) {
	var vv []interface{}
	for _, v := range clsfs {
		vv = append(vv, v)
	}
	return c.Create(CrmLeadScoringFrequencyModel, vv, nil)
}

// UpdateCrmLeadScoringFrequency updates an existing crm.lead.scoring.frequency record.
func (c *Client) UpdateCrmLeadScoringFrequency(clsf *CrmLeadScoringFrequency) error {
	return c.UpdateCrmLeadScoringFrequencys([]int64{clsf.Id.Get()}, clsf)
}

// UpdateCrmLeadScoringFrequencys updates existing crm.lead.scoring.frequency records.
// All records (represented by ids) will be updated by clsf values.
func (c *Client) UpdateCrmLeadScoringFrequencys(ids []int64, clsf *CrmLeadScoringFrequency) error {
	return c.Update(CrmLeadScoringFrequencyModel, ids, clsf, nil)
}

// DeleteCrmLeadScoringFrequency deletes an existing crm.lead.scoring.frequency record.
func (c *Client) DeleteCrmLeadScoringFrequency(id int64) error {
	return c.DeleteCrmLeadScoringFrequencys([]int64{id})
}

// DeleteCrmLeadScoringFrequencys deletes existing crm.lead.scoring.frequency records.
func (c *Client) DeleteCrmLeadScoringFrequencys(ids []int64) error {
	return c.Delete(CrmLeadScoringFrequencyModel, ids)
}

// GetCrmLeadScoringFrequency gets crm.lead.scoring.frequency existing record.
func (c *Client) GetCrmLeadScoringFrequency(id int64) (*CrmLeadScoringFrequency, error) {
	clsfs, err := c.GetCrmLeadScoringFrequencys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*clsfs)[0]), nil
}

// GetCrmLeadScoringFrequencys gets crm.lead.scoring.frequency existing records.
func (c *Client) GetCrmLeadScoringFrequencys(ids []int64) (*CrmLeadScoringFrequencys, error) {
	clsfs := &CrmLeadScoringFrequencys{}
	if err := c.Read(CrmLeadScoringFrequencyModel, ids, nil, clsfs); err != nil {
		return nil, err
	}
	return clsfs, nil
}

// FindCrmLeadScoringFrequency finds crm.lead.scoring.frequency record by querying it with criteria.
func (c *Client) FindCrmLeadScoringFrequency(criteria *Criteria) (*CrmLeadScoringFrequency, error) {
	clsfs := &CrmLeadScoringFrequencys{}
	if err := c.SearchRead(CrmLeadScoringFrequencyModel, criteria, NewOptions().Limit(1), clsfs); err != nil {
		return nil, err
	}
	return &((*clsfs)[0]), nil
}

// FindCrmLeadScoringFrequencys finds crm.lead.scoring.frequency records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadScoringFrequencys(criteria *Criteria, options *Options) (*CrmLeadScoringFrequencys, error) {
	clsfs := &CrmLeadScoringFrequencys{}
	if err := c.SearchRead(CrmLeadScoringFrequencyModel, criteria, options, clsfs); err != nil {
		return nil, err
	}
	return clsfs, nil
}

// FindCrmLeadScoringFrequencyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLeadScoringFrequencyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmLeadScoringFrequencyModel, criteria, options)
}

// FindCrmLeadScoringFrequencyId finds record id by querying it with criteria.
func (c *Client) FindCrmLeadScoringFrequencyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLeadScoringFrequencyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
