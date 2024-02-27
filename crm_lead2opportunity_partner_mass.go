package odoo

// CrmLead2OpportunityPartnerMass represents crm.lead2opportunity.partner.mass model.
type CrmLead2OpportunityPartnerMass struct {
	LastUpdate       *Time      `xmlrpc:"__last_update,omitempty"`
	Action           *Selection `xmlrpc:"action,omitempty"`
	CreateDate       *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid        *Many2One  `xmlrpc:"create_uid,omitempty"`
	Deduplicate      *Bool      `xmlrpc:"deduplicate,omitempty"`
	DisplayName      *String    `xmlrpc:"display_name,omitempty"`
	ForceAssignation *Bool      `xmlrpc:"force_assignation,omitempty"`
	Id               *Int       `xmlrpc:"id,omitempty"`
	Name             *Selection `xmlrpc:"name,omitempty"`
	OpportunityIds   *Relation  `xmlrpc:"opportunity_ids,omitempty"`
	PartnerId        *Many2One  `xmlrpc:"partner_id,omitempty"`
	TeamId           *Many2One  `xmlrpc:"team_id,omitempty"`
	UserId           *Many2One  `xmlrpc:"user_id,omitempty"`
	UserIds          *Relation  `xmlrpc:"user_ids,omitempty"`
	WriteDate        *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid         *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CrmLead2OpportunityPartnerMasss represents array of crm.lead2opportunity.partner.mass model.
type CrmLead2OpportunityPartnerMasss []CrmLead2OpportunityPartnerMass

// CrmLead2OpportunityPartnerMassModel is the odoo model name.
const CrmLead2OpportunityPartnerMassModel = "crm.lead2opportunity.partner.mass"

// Many2One convert CrmLead2OpportunityPartnerMass to *Many2One.
func (clpm *CrmLead2OpportunityPartnerMass) Many2One() *Many2One {
	return NewMany2One(clpm.Id.Get(), "")
}

// CreateCrmLead2OpportunityPartnerMass creates a new crm.lead2opportunity.partner.mass model and returns its id.
func (c *Client) CreateCrmLead2OpportunityPartnerMass(clpm *CrmLead2OpportunityPartnerMass) (int64, error) {
	ids, err := c.CreateCrmLead2OpportunityPartnerMasss([]*CrmLead2OpportunityPartnerMass{clpm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmLead2OpportunityPartnerMasss creates a new crm.lead2opportunity.partner.mass model and returns its id.
func (c *Client) CreateCrmLead2OpportunityPartnerMasss(clpms []*CrmLead2OpportunityPartnerMass) ([]int64, error) {
	var vv []interface{}
	for _, v := range clpms {
		vv = append(vv, v)
	}
	return c.Create(CrmLead2OpportunityPartnerMassModel, vv, nil)
}

// UpdateCrmLead2OpportunityPartnerMass updates an existing crm.lead2opportunity.partner.mass record.
func (c *Client) UpdateCrmLead2OpportunityPartnerMass(clpm *CrmLead2OpportunityPartnerMass) error {
	return c.UpdateCrmLead2OpportunityPartnerMasss([]int64{clpm.Id.Get()}, clpm)
}

// UpdateCrmLead2OpportunityPartnerMasss updates existing crm.lead2opportunity.partner.mass records.
// All records (represented by ids) will be updated by clpm values.
func (c *Client) UpdateCrmLead2OpportunityPartnerMasss(ids []int64, clpm *CrmLead2OpportunityPartnerMass) error {
	return c.Update(CrmLead2OpportunityPartnerMassModel, ids, clpm, nil)
}

// DeleteCrmLead2OpportunityPartnerMass deletes an existing crm.lead2opportunity.partner.mass record.
func (c *Client) DeleteCrmLead2OpportunityPartnerMass(id int64) error {
	return c.DeleteCrmLead2OpportunityPartnerMasss([]int64{id})
}

// DeleteCrmLead2OpportunityPartnerMasss deletes existing crm.lead2opportunity.partner.mass records.
func (c *Client) DeleteCrmLead2OpportunityPartnerMasss(ids []int64) error {
	return c.Delete(CrmLead2OpportunityPartnerMassModel, ids)
}

// GetCrmLead2OpportunityPartnerMass gets crm.lead2opportunity.partner.mass existing record.
func (c *Client) GetCrmLead2OpportunityPartnerMass(id int64) (*CrmLead2OpportunityPartnerMass, error) {
	clpms, err := c.GetCrmLead2OpportunityPartnerMasss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*clpms)[0]), nil
}

// GetCrmLead2OpportunityPartnerMasss gets crm.lead2opportunity.partner.mass existing records.
func (c *Client) GetCrmLead2OpportunityPartnerMasss(ids []int64) (*CrmLead2OpportunityPartnerMasss, error) {
	clpms := &CrmLead2OpportunityPartnerMasss{}
	if err := c.Read(CrmLead2OpportunityPartnerMassModel, ids, nil, clpms); err != nil {
		return nil, err
	}
	return clpms, nil
}

// FindCrmLead2OpportunityPartnerMass finds crm.lead2opportunity.partner.mass record by querying it with criteria.
func (c *Client) FindCrmLead2OpportunityPartnerMass(criteria *Criteria) (*CrmLead2OpportunityPartnerMass, error) {
	clpms := &CrmLead2OpportunityPartnerMasss{}
	if err := c.SearchRead(CrmLead2OpportunityPartnerMassModel, criteria, NewOptions().Limit(1), clpms); err != nil {
		return nil, err
	}
	return &((*clpms)[0]), nil
}

// FindCrmLead2OpportunityPartnerMasss finds crm.lead2opportunity.partner.mass records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLead2OpportunityPartnerMasss(criteria *Criteria, options *Options) (*CrmLead2OpportunityPartnerMasss, error) {
	clpms := &CrmLead2OpportunityPartnerMasss{}
	if err := c.SearchRead(CrmLead2OpportunityPartnerMassModel, criteria, options, clpms); err != nil {
		return nil, err
	}
	return clpms, nil
}

// FindCrmLead2OpportunityPartnerMassIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmLead2OpportunityPartnerMassIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmLead2OpportunityPartnerMassModel, criteria, options)
}

// FindCrmLead2OpportunityPartnerMassId finds record id by querying it with criteria.
func (c *Client) FindCrmLead2OpportunityPartnerMassId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmLead2OpportunityPartnerMassModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
