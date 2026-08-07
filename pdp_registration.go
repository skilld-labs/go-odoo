package odoo

// PdpRegistration represents pdp.registration model.
type PdpRegistration struct {
	AccountPeppolProxyState *Selection  `xmlrpc:"account_peppol_proxy_state,omitempty"`
	AuthUrlHash             *String     `xmlrpc:"auth_url_hash,omitempty"`
	CompanyId               *Many2One   `xmlrpc:"company_id,omitempty"`
	ContactEmail            *String     `xmlrpc:"contact_email,omitempty"`
	CreateDate              *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName             *String     `xmlrpc:"display_name,omitempty"`
	EdiMode                 *Selection  `xmlrpc:"edi_mode,omitempty"`
	EdiUserId               *Many2One   `xmlrpc:"edi_user_id,omitempty"`
	Id                      *Int        `xmlrpc:"id,omitempty"`
	PdpAuthenticationUuid   *String     `xmlrpc:"pdp_authentication_uuid,omitempty"`
	PdpIdentifier           *String     `xmlrpc:"pdp_identifier,omitempty"`
	PdpKycStatus            *Selection  `xmlrpc:"pdp_kyc_status,omitempty"`
	PdpPilotPhase           *Bool       `xmlrpc:"pdp_pilot_phase,omitempty"`
	SirenNumber             *String     `xmlrpc:"siren_number,omitempty"`
	Warnings                interface{} `xmlrpc:"warnings,omitempty"`
	WriteDate               *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// PdpRegistrations represents array of pdp.registration model.
type PdpRegistrations []PdpRegistration

// PdpRegistrationModel is the odoo model name.
const PdpRegistrationModel = "pdp.registration"

// Many2One convert PdpRegistration to *Many2One.
func (pr *PdpRegistration) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreatePdpRegistration creates a new pdp.registration model and returns its id.
func (c *Client) CreatePdpRegistration(pr *PdpRegistration) (int64, error) {
	ids, err := c.CreatePdpRegistrations([]*PdpRegistration{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePdpRegistration creates a new pdp.registration model and returns its id.
func (c *Client) CreatePdpRegistrations(prs []*PdpRegistration) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(PdpRegistrationModel, vv, nil)
}

// UpdatePdpRegistration updates an existing pdp.registration record.
func (c *Client) UpdatePdpRegistration(pr *PdpRegistration) error {
	return c.UpdatePdpRegistrations([]int64{pr.Id.Get()}, pr)
}

// UpdatePdpRegistrations updates existing pdp.registration records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdatePdpRegistrations(ids []int64, pr *PdpRegistration) error {
	return c.Update(PdpRegistrationModel, ids, pr, nil)
}

// DeletePdpRegistration deletes an existing pdp.registration record.
func (c *Client) DeletePdpRegistration(id int64) error {
	return c.DeletePdpRegistrations([]int64{id})
}

// DeletePdpRegistrations deletes existing pdp.registration records.
func (c *Client) DeletePdpRegistrations(ids []int64) error {
	return c.Delete(PdpRegistrationModel, ids)
}

// GetPdpRegistration gets pdp.registration existing record.
func (c *Client) GetPdpRegistration(id int64) (*PdpRegistration, error) {
	prs, err := c.GetPdpRegistrations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetPdpRegistrations gets pdp.registration existing records.
func (c *Client) GetPdpRegistrations(ids []int64) (*PdpRegistrations, error) {
	prs := &PdpRegistrations{}
	if err := c.Read(PdpRegistrationModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPdpRegistration finds pdp.registration record by querying it with criteria.
func (c *Client) FindPdpRegistration(criteria *Criteria) (*PdpRegistration, error) {
	prs := &PdpRegistrations{}
	if err := c.SearchRead(PdpRegistrationModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindPdpRegistrations finds pdp.registration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPdpRegistrations(criteria *Criteria, options *Options) (*PdpRegistrations, error) {
	prs := &PdpRegistrations{}
	if err := c.SearchRead(PdpRegistrationModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPdpRegistrationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPdpRegistrationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PdpRegistrationModel, criteria, options)
}

// FindPdpRegistrationId finds record id by querying it with criteria.
func (c *Client) FindPdpRegistrationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PdpRegistrationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
