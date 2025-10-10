package odoo

// StockRule represents stock.rule model.
type StockRule struct {
	Action                 *Selection `xmlrpc:"action,omitempty"`
	Active                 *Bool      `xmlrpc:"active,omitempty"`
	Auto                   *Selection `xmlrpc:"auto,omitempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omitempty"`
	Delay                  *Int       `xmlrpc:"delay,omitempty"`
	DisplayName            *String    `xmlrpc:"display_name,omitempty"`
	GroupId                *Many2One  `xmlrpc:"group_id,omitempty"`
	GroupPropagationOption *Selection `xmlrpc:"group_propagation_option,omitempty"`
	Id                     *Int       `xmlrpc:"id,omitempty"`
	LocationDestFromRule   *Bool      `xmlrpc:"location_dest_from_rule,omitempty"`
	LocationDestId         *Many2One  `xmlrpc:"location_dest_id,omitempty"`
	LocationSrcId          *Many2One  `xmlrpc:"location_src_id,omitempty"`
	Name                   *String    `xmlrpc:"name,omitempty"`
	PartnerAddressId       *Many2One  `xmlrpc:"partner_address_id,omitempty"`
	PickingTypeCodeDomain  *String    `xmlrpc:"picking_type_code_domain,omitempty"`
	PickingTypeId          *Many2One  `xmlrpc:"picking_type_id,omitempty"`
	ProcureMethod          *Selection `xmlrpc:"procure_method,omitempty"`
	PropagateCancel        *Bool      `xmlrpc:"propagate_cancel,omitempty"`
	PropagateCarrier       *Bool      `xmlrpc:"propagate_carrier,omitempty"`
	PropagateWarehouseId   *Many2One  `xmlrpc:"propagate_warehouse_id,omitempty"`
	PushDomain             *String    `xmlrpc:"push_domain,omitempty"`
	RouteCompanyId         *Many2One  `xmlrpc:"route_company_id,omitempty"`
	RouteId                *Many2One  `xmlrpc:"route_id,omitempty"`
	RouteSequence          *Int       `xmlrpc:"route_sequence,omitempty"`
	RuleMessage            *String    `xmlrpc:"rule_message,omitempty"`
	Sequence               *Int       `xmlrpc:"sequence,omitempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// StockRules represents array of stock.rule model.
type StockRules []StockRule

// StockRuleModel is the odoo model name.
const StockRuleModel = "stock.rule"

// Many2One convert StockRule to *Many2One.
func (sr *StockRule) Many2One() *Many2One {
	return NewMany2One(sr.Id.Get(), "")
}

// CreateStockRule creates a new stock.rule model and returns its id.
func (c *Client) CreateStockRule(sr *StockRule) (int64, error) {
	ids, err := c.CreateStockRules([]*StockRule{sr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateStockRule creates a new stock.rule model and returns its id.
func (c *Client) CreateStockRules(srs []*StockRule) ([]int64, error) {
	var vv []interface{}
	for _, v := range srs {
		vv = append(vv, v)
	}
	return c.Create(StockRuleModel, vv, nil)
}

// UpdateStockRule updates an existing stock.rule record.
func (c *Client) UpdateStockRule(sr *StockRule) error {
	return c.UpdateStockRules([]int64{sr.Id.Get()}, sr)
}

// UpdateStockRules updates existing stock.rule records.
// All records (represented by ids) will be updated by sr values.
func (c *Client) UpdateStockRules(ids []int64, sr *StockRule) error {
	return c.Update(StockRuleModel, ids, sr, nil)
}

// DeleteStockRule deletes an existing stock.rule record.
func (c *Client) DeleteStockRule(id int64) error {
	return c.DeleteStockRules([]int64{id})
}

// DeleteStockRules deletes existing stock.rule records.
func (c *Client) DeleteStockRules(ids []int64) error {
	return c.Delete(StockRuleModel, ids)
}

// GetStockRule gets stock.rule existing record.
func (c *Client) GetStockRule(id int64) (*StockRule, error) {
	srs, err := c.GetStockRules([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// GetStockRules gets stock.rule existing records.
func (c *Client) GetStockRules(ids []int64) (*StockRules, error) {
	srs := &StockRules{}
	if err := c.Read(StockRuleModel, ids, nil, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockRule finds stock.rule record by querying it with criteria.
func (c *Client) FindStockRule(criteria *Criteria) (*StockRule, error) {
	srs := &StockRules{}
	if err := c.SearchRead(StockRuleModel, criteria, NewOptions().Limit(1), srs); err != nil {
		return nil, err
	}
	return &((*srs)[0]), nil
}

// FindStockRules finds stock.rule records by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRules(criteria *Criteria, options *Options) (*StockRules, error) {
	srs := &StockRules{}
	if err := c.SearchRead(StockRuleModel, criteria, options, srs); err != nil {
		return nil, err
	}
	return srs, nil
}

// FindStockRuleIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindStockRuleIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(StockRuleModel, criteria, options)
}

// FindStockRuleId finds record id by querying it with criteria.
func (c *Client) FindStockRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(StockRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
