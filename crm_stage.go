package odoo

// CrmStage represents crm.stage model.
type CrmStage struct {
	LastUpdate     *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate     *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String   `xmlrpc:"display_name,omitempty"`
	Fold           *Bool     `xmlrpc:"fold,omitempty"`
	Id             *Int      `xmlrpc:"id,omitempty"`
	LegendPriority *String   `xmlrpc:"legend_priority,omitempty"`
	Name           *String   `xmlrpc:"name,omitempty"`
	OnChange       *Bool     `xmlrpc:"on_change,omitempty"`
	Probability    *Float    `xmlrpc:"probability,omitempty"`
	Requirements   *String   `xmlrpc:"requirements,omitempty"`
	Sequence       *Int      `xmlrpc:"sequence,omitempty"`
	TeamId         *Many2One `xmlrpc:"team_id,omitempty"`
	WriteDate      *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmStages represents array of crm.stage model.
type CrmStages []CrmStage

// CrmStageModel is the odoo model name.
const CrmStageModel = "crm.stage"

// Many2One convert CrmStage to *Many2One.
func (cs *CrmStage) Many2One() *Many2One {
	return NewMany2One(cs.Id.Get(), "")
}

// CreateCrmStage creates a new crm.stage model and returns its id.
func (c *Client) CreateCrmStage(cs *CrmStage) (int64, error) {
	ids, err := c.CreateCrmStages([]*CrmStage{cs})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmStages creates a new crm.stage model and returns its id.
func (c *Client) CreateCrmStages(css []*CrmStage) ([]int64, error) {
	var vv []interface{}
	for _, v := range css {
		vv = append(vv, v)
	}
	return c.Create(CrmStageModel, vv, nil)
}

// UpdateCrmStage updates an existing crm.stage record.
func (c *Client) UpdateCrmStage(cs *CrmStage) error {
	return c.UpdateCrmStages([]int64{cs.Id.Get()}, cs)
}

// UpdateCrmStages updates existing crm.stage records.
// All records (represented by ids) will be updated by cs values.
func (c *Client) UpdateCrmStages(ids []int64, cs *CrmStage) error {
	return c.Update(CrmStageModel, ids, cs, nil)
}

// DeleteCrmStage deletes an existing crm.stage record.
func (c *Client) DeleteCrmStage(id int64) error {
	return c.DeleteCrmStages([]int64{id})
}

// DeleteCrmStages deletes existing crm.stage records.
func (c *Client) DeleteCrmStages(ids []int64) error {
	return c.Delete(CrmStageModel, ids)
}

// GetCrmStage gets crm.stage existing record.
func (c *Client) GetCrmStage(id int64) (*CrmStage, error) {
	css, err := c.GetCrmStages([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*css)[0]), nil
}

// GetCrmStages gets crm.stage existing records.
func (c *Client) GetCrmStages(ids []int64) (*CrmStages, error) {
	css := &CrmStages{}
	if err := c.Read(CrmStageModel, ids, nil, css); err != nil {
		return nil, err
	}
	return css, nil
}

// FindCrmStage finds crm.stage record by querying it with criteria.
func (c *Client) FindCrmStage(criteria *Criteria) (*CrmStage, error) {
	css := &CrmStages{}
	if err := c.SearchRead(CrmStageModel, criteria, NewOptions().Limit(1), css); err != nil {
		return nil, err
	}
	return &((*css)[0]), nil
}

// FindCrmStages finds crm.stage records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmStages(criteria *Criteria, options *Options) (*CrmStages, error) {
	css := &CrmStages{}
	if err := c.SearchRead(CrmStageModel, criteria, options, css); err != nil {
		return nil, err
	}
	return css, nil
}

// FindCrmStageIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmStageIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmStageModel, criteria, options)
}

// FindCrmStageId finds record id by querying it with criteria.
func (c *Client) FindCrmStageId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmStageModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
