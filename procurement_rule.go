package odoo

import (
	"fmt"
)

// ProcurementRule represents procurement.rule model.
type ProcurementRule struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	Action                 *Selection `xmlrpc:"action,omptempty"`
	Active                 *Bool      `xmlrpc:"active,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	Delay                  *Int       `xmlrpc:"delay,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	GroupId                *Many2One  `xmlrpc:"group_id,omptempty"`
	GroupPropagationOption *Selection `xmlrpc:"group_propagation_option,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	LocationId             *Many2One  `xmlrpc:"location_id,omptempty"`
	LocationSrcId          *Many2One  `xmlrpc:"location_src_id,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	PartnerAddressId       *Many2One  `xmlrpc:"partner_address_id,omptempty"`
	PickingTypeId          *Many2One  `xmlrpc:"picking_type_id,omptempty"`
	ProcureMethod          *Selection `xmlrpc:"procure_method,omptempty"`
	Propagate              *Bool      `xmlrpc:"propagate,omptempty"`
	PropagateWarehouseId   *Many2One  `xmlrpc:"propagate_warehouse_id,omptempty"`
	RouteId                *Many2One  `xmlrpc:"route_id,omptempty"`
	RouteSequence          *Int       `xmlrpc:"route_sequence,omptempty"`
	Sequence               *Int       `xmlrpc:"sequence,omptempty"`
	WarehouseId            *Many2One  `xmlrpc:"warehouse_id,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
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
	return c.Create(ProcurementRuleModel, pr)
}

// UpdateProcurementRule updates an existing procurement.rule record.
func (c *Client) UpdateProcurementRule(pr *ProcurementRule) error {
	return c.UpdateProcurementRules([]int64{pr.Id.Get()}, pr)
}

// UpdateProcurementRules updates existing procurement.rule records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdateProcurementRules(ids []int64, pr *ProcurementRule) error {
	return c.Update(ProcurementRuleModel, ids, pr)
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
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of procurement.rule not found", id)
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
	if prs != nil && len(*prs) > 0 {
		return &((*prs)[0]), nil
	}
	return nil, fmt.Errorf("no procurement.rule was found with criteria %v", criteria)
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
	ids, err := c.Search(ProcurementRuleModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProcurementRuleId finds record id by querying it with criteria.
func (c *Client) FindProcurementRuleId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProcurementRuleModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("no procurement.rule was found with criteria %v and options %v", criteria, options)
}
