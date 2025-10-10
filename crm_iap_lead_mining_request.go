package odoo

// CrmIapLeadMiningRequest represents crm.iap.lead.mining.request model.
type CrmIapLeadMiningRequest struct {
	AvailableStateIds   *Relation  `xmlrpc:"available_state_ids,omitempty"`
	CompanySizeMax      *Int       `xmlrpc:"company_size_max,omitempty"`
	CompanySizeMin      *Int       `xmlrpc:"company_size_min,omitempty"`
	ContactFilterType   *Selection `xmlrpc:"contact_filter_type,omitempty"`
	ContactNumber       *Int       `xmlrpc:"contact_number,omitempty"`
	CountryIds          *Relation  `xmlrpc:"country_ids,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	ErrorType           *Selection `xmlrpc:"error_type,omitempty"`
	FilterOnSize        *Bool      `xmlrpc:"filter_on_size,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	IndustryIds         *Relation  `xmlrpc:"industry_ids,omitempty"`
	LeadContactsCredits *String    `xmlrpc:"lead_contacts_credits,omitempty"`
	LeadCount           *Int       `xmlrpc:"lead_count,omitempty"`
	LeadCredits         *String    `xmlrpc:"lead_credits,omitempty"`
	LeadIds             *Relation  `xmlrpc:"lead_ids,omitempty"`
	LeadNumber          *Int       `xmlrpc:"lead_number,omitempty"`
	LeadTotalCredits    *String    `xmlrpc:"lead_total_credits,omitempty"`
	LeadType            *Selection `xmlrpc:"lead_type,omitempty"`
	Name                *String    `xmlrpc:"name,omitempty"`
	PreferredRoleId     *Many2One  `xmlrpc:"preferred_role_id,omitempty"`
	RoleIds             *Relation  `xmlrpc:"role_ids,omitempty"`
	SearchType          *Selection `xmlrpc:"search_type,omitempty"`
	SeniorityId         *Many2One  `xmlrpc:"seniority_id,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	StateIds            *Relation  `xmlrpc:"state_ids,omitempty"`
	TagIds              *Relation  `xmlrpc:"tag_ids,omitempty"`
	TeamId              *Many2One  `xmlrpc:"team_id,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CrmIapLeadMiningRequests represents array of crm.iap.lead.mining.request model.
type CrmIapLeadMiningRequests []CrmIapLeadMiningRequest

// CrmIapLeadMiningRequestModel is the odoo model name.
const CrmIapLeadMiningRequestModel = "crm.iap.lead.mining.request"

// Many2One convert CrmIapLeadMiningRequest to *Many2One.
func (cilmr *CrmIapLeadMiningRequest) Many2One() *Many2One {
	return NewMany2One(cilmr.Id.Get(), "")
}

// CreateCrmIapLeadMiningRequest creates a new crm.iap.lead.mining.request model and returns its id.
func (c *Client) CreateCrmIapLeadMiningRequest(cilmr *CrmIapLeadMiningRequest) (int64, error) {
	ids, err := c.CreateCrmIapLeadMiningRequests([]*CrmIapLeadMiningRequest{cilmr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmIapLeadMiningRequest creates a new crm.iap.lead.mining.request model and returns its id.
func (c *Client) CreateCrmIapLeadMiningRequests(cilmrs []*CrmIapLeadMiningRequest) ([]int64, error) {
	var vv []interface{}
	for _, v := range cilmrs {
		vv = append(vv, v)
	}
	return c.Create(CrmIapLeadMiningRequestModel, vv, nil)
}

// UpdateCrmIapLeadMiningRequest updates an existing crm.iap.lead.mining.request record.
func (c *Client) UpdateCrmIapLeadMiningRequest(cilmr *CrmIapLeadMiningRequest) error {
	return c.UpdateCrmIapLeadMiningRequests([]int64{cilmr.Id.Get()}, cilmr)
}

// UpdateCrmIapLeadMiningRequests updates existing crm.iap.lead.mining.request records.
// All records (represented by ids) will be updated by cilmr values.
func (c *Client) UpdateCrmIapLeadMiningRequests(ids []int64, cilmr *CrmIapLeadMiningRequest) error {
	return c.Update(CrmIapLeadMiningRequestModel, ids, cilmr, nil)
}

// DeleteCrmIapLeadMiningRequest deletes an existing crm.iap.lead.mining.request record.
func (c *Client) DeleteCrmIapLeadMiningRequest(id int64) error {
	return c.DeleteCrmIapLeadMiningRequests([]int64{id})
}

// DeleteCrmIapLeadMiningRequests deletes existing crm.iap.lead.mining.request records.
func (c *Client) DeleteCrmIapLeadMiningRequests(ids []int64) error {
	return c.Delete(CrmIapLeadMiningRequestModel, ids)
}

// GetCrmIapLeadMiningRequest gets crm.iap.lead.mining.request existing record.
func (c *Client) GetCrmIapLeadMiningRequest(id int64) (*CrmIapLeadMiningRequest, error) {
	cilmrs, err := c.GetCrmIapLeadMiningRequests([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cilmrs)[0]), nil
}

// GetCrmIapLeadMiningRequests gets crm.iap.lead.mining.request existing records.
func (c *Client) GetCrmIapLeadMiningRequests(ids []int64) (*CrmIapLeadMiningRequests, error) {
	cilmrs := &CrmIapLeadMiningRequests{}
	if err := c.Read(CrmIapLeadMiningRequestModel, ids, nil, cilmrs); err != nil {
		return nil, err
	}
	return cilmrs, nil
}

// FindCrmIapLeadMiningRequest finds crm.iap.lead.mining.request record by querying it with criteria.
func (c *Client) FindCrmIapLeadMiningRequest(criteria *Criteria) (*CrmIapLeadMiningRequest, error) {
	cilmrs := &CrmIapLeadMiningRequests{}
	if err := c.SearchRead(CrmIapLeadMiningRequestModel, criteria, NewOptions().Limit(1), cilmrs); err != nil {
		return nil, err
	}
	return &((*cilmrs)[0]), nil
}

// FindCrmIapLeadMiningRequests finds crm.iap.lead.mining.request records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadMiningRequests(criteria *Criteria, options *Options) (*CrmIapLeadMiningRequests, error) {
	cilmrs := &CrmIapLeadMiningRequests{}
	if err := c.SearchRead(CrmIapLeadMiningRequestModel, criteria, options, cilmrs); err != nil {
		return nil, err
	}
	return cilmrs, nil
}

// FindCrmIapLeadMiningRequestIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadMiningRequestIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmIapLeadMiningRequestModel, criteria, options)
}

// FindCrmIapLeadMiningRequestId finds record id by querying it with criteria.
func (c *Client) FindCrmIapLeadMiningRequestId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmIapLeadMiningRequestModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
