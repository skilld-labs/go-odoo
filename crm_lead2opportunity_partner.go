package odoo

// CrmLead2OpportunityPartner represents crm.lead2opportunity.partner model.
type CrmLead2OpportunityPartner struct {
	LastUpdate     *Time      `xmlrpc:"__last_update,omitempty"`
	Action         *Selection `xmlrpc:"action,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	Name           *Selection `xmlrpc:"name,omitempty"`
	OpportunityIds *Relation  `xmlrpc:"opportunity_ids,omitempty"`
	PartnerId      *Many2One  `xmlrpc:"partner_id,omitempty"`
	TeamId         *Many2One  `xmlrpc:"team_id,omitempty"`
	UserId         *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CrmLead2OpportunityPartners represents array of crm.lead2opportunity.partner model.
type CrmLead2OpportunityPartners []CrmLead2OpportunityPartner

// CrmLead2OpportunityPartnerModel is the odoo model name.
const CrmLead2OpportunityPartnerModel = "crm.lead2opportunity.partner"

// Many2One convert CrmLead2OpportunityPartner to *Many2One.
func (clp *CrmLead2OpportunityPartner) Many2One() *Many2One {
	return NewMany2One(clp.Id.Get(), "")
}

// CreateCrmLead2OpportunityPartner creates a new crm.lead2opportunity.partner model and returns its id.
func (c *Client) CreateCrmLead2OpportunityPartner(clp *CrmLead2OpportunityPartner) (int64, error) {
	ids, err := c.CreateCrmLead2OpportunityPartners([]*CrmLead2OpportunityPartner{clp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmLead2OpportunityPartner creates a new crm.lead2opportunity.partner model and returns its id.
func (c *Client) CreateCrmLead2OpportunityPartners(clps []*CrmLead2OpportunityPartner) ([]int64, error) {
	var vv []interface{}
	for _, v := range clps {
		vv = append(vv, v)
	}
	return c.Create(CrmLead2OpportunityPartnerModel, vv, nil)
}

// UpdateCrmLead2OpportunityPartner updates an existing crm.lead2opportunity.partner record.
func (c *Client) UpdateCrmLead2OpportunityPartner(clp *CrmLead2OpportunityPartner) error {
	return c.UpdateCrmLead2OpportunityPartners([]int64{clp.Id.Get()}, clp)
}

// UpdateCrmLead2OpportunityPartners updates existing crm.lead2opportunity.partner records.
// All records (represented by ids) will be updated by clp values.
func (c *Client) UpdateCrmLead2OpportunityPartners(ids []int64, clp *CrmLead2OpportunityPartner) error {
	return c.Update(CrmLead2OpportunityPartnerModel, ids, clp, nil)
}

// DeleteCrmLead2OpportunityPartner deletes an existing crm.lead2opportunity.partner record.
func (c *Client) DeleteCrmLead2OpportunityPartner(id int64) error {
	return c.DeleteCrmLead2OpportunityPartners([]int64{id})
}

// DeleteCrmLead2OpportunityPartners deletes existing crm.lead2opportunity.partner records.
func (c *Client) DeleteCrmLead2OpportunityPartners(ids []int64) error {
	return c.Delete(CrmLead2OpportunityPartnerModel, ids)
}

// GetCrmLead2OpportunityPartner gets crm.lead2opportunity.partner existing record.
func (c *Client) GetCrmLead2OpportunityPartner(id int64) (*CrmLead2OpportunityPartner, error) {
	clps, err := c.GetCrmLead2OpportunityPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*clps)[0]), nil
}

// GetCrmLead2OpportunityPartners gets crm.lead2opportunity.partner existing records.
func (c *Client) GetCrmLead2OpportunityPartners(ids []int64) (*CrmLead2OpportunityPartners, error) {
	clps := &CrmLead2OpportunityPartners{}
	if err := c.Read(CrmLead2OpportunityPartnerModel, ids, nil, clps); err != nil {
		return nil, err
	}
	return clps, nil
}

// FindCrmLead2OpportunityPartner finds crm.lead2opportunity.partner record by querying it with criteria.
func (c *Client) FindCrmLead2OpportunityPartner(criteria *Criteria) (*CrmLead2OpportunityPartner, error) {
	clps := &CrmLead2OpportunityPartners{}
	if err := c.SearchRead(CrmLead2OpportunityPartnerModel, criteria, NewOptions().Limit(1), clps); err != nil {
		return nil, err
	}
	return &((*clps)[0]), nil
}

// FindCrmLead2OpportunityPartners finds crm.lead2opportunity.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLead2OpportunityPartners(criteria *Criteria, options *Options) (*CrmLead2OpportunityPartners, error) {
	clps := &CrmLead2OpportunityPartners{}
	if err := c.SearchRead(CrmLead2OpportunityPartnerModel, criteria, options, clps); err != nil {
		return nil, err
	}
	return clps, nil
}

// FindCrmLead2OpportunityPartnerIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLead2OpportunityPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmLead2OpportunityPartnerModel, criteria, options)
}

// FindCrmLead2OpportunityPartnerId finds record id by querying it with criteria.
func (c *Client) FindCrmLead2OpportunityPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLead2OpportunityPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
