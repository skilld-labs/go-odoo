package odoo

// ProcurementRule represents procurement.rule model.
type ProcurementRule struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omitempty"`
	Action                 *Selection `xmlrpc:"action,omitempty"`
	Active                 *Bool      `xmlrpc:"active,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	Delay                  *Int       `xmlrpc:"delay,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	GroupId                *Many2One  `xmlrpc:"group_id,omitempty"`
	GroupPropagationOption *Selection `xmlrpc:"group_propagation_option,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	LocationId             *Many2One  `xmlrpc:"location_id,omitempty"`
	LocationSrcId          *Many2One  `xmlrpc:"location_src_id,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	PartnerAddressId       *Many2One  `xmlrpc:"partner_address_id,omitempty"`
	PickingTypeId          *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	ProcureMethod          *Selection `xmlrpc:"procure_method,omitempty"`
	Propagate              *Bool      `xmlrpc:"propagate,omitempty"`
	PropagateWarehouseId   *Many2One  `xmlrpc:"propagate_warehouse_id,omitempty"`
	RouteId                *Many2One  `xmlrpc:"route_id,omitempty"`
	RouteSequence          *Int       `xmlrpc:"route_sequence,omitempty"`
	Sequence               *Int       `xmlrpc:"sequence,omitempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProcurementRules represents array of procurement.rule model.
type ProcurementRules []ProcurementRule

// ProcurementRuleModel is the odoo model name.
const ProcurementRuleModel = "procurement.rule"

// Many2One convert ProcurementRule to *Many2One.
func (pr *ProcurementRule) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreateProcurementRule creates a new procurement.rule model and returns its id.
func (c *Client) CreateProcurementRule(pr *ProcurementRule) (int64, error) {
	ids, err := c.CreateProcurementRules([]*ProcurementRule{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProcurementRules creates a new procurement.rule model and returns its id.
func (c *Client) CreateProcurementRules(prs []*ProcurementRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(ProcurementRuleModel, vv, nil)
}

// UpdateProcurementRule updates an existing procurement.rule record.
func (c *Client) UpdateProcurementRule(pr *ProcurementRule) error {
	return c.UpdateProcurementRules([]int64{pr.Id.Get()}, pr)
}

// UpdateProcurementRules updates existing procurement.rule records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdateProcurementRules(ids []int64, pr *ProcurementRule) error {
	return c.Update(ProcurementRuleModel, ids, pr, nil)
}

// DeleteProcurementRule deletes an existing procurement.rule record.
func (c *Client) DeleteProcurementRule(id int64) error {
	return c.DeleteProcurementRules([]int64{id})
}

// DeleteProcurementRules deletes existing procurement.rule records.
func (c *Client) DeleteProcurementRules(ids []int64) error {
	return c.Delete(ProcurementRuleModel, ids)
}

// GetProcurementRule gets procurement.rule existing record.
func (c *Client) GetProcurementRule(id int64) (*ProcurementRule, error) {
	prs, err := c.GetProcurementRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetProcurementRules gets procurement.rule existing records.
func (c *Client) GetProcurementRules(ids []int64) (*ProcurementRules, error) {
	prs := &ProcurementRules{}
	if err := c.Read(ProcurementRuleModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProcurementRule finds procurement.rule record by querying it with criteria.
func (c *Client) FindProcurementRule(criteria *Criteria) (*ProcurementRule, error) {
	prs := &ProcurementRules{}
	if err := c.SearchRead(ProcurementRuleModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindProcurementRules finds procurement.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProcurementRules(criteria *Criteria, options *Options) (*ProcurementRules, error) {
	prs := &ProcurementRules{}
	if err := c.SearchRead(ProcurementRuleModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProcurementRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProcurementRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProcurementRuleModel, criteria, options)
}

// FindProcurementRuleId finds record id by querying it with criteria.
func (c *Client) FindProcurementRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProcurementRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
