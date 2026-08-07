package odoo

// PeppolRegistration represents peppol.registration model.
type PeppolRegistration struct {
	AccountPeppolProxyState             *Selection  `xmlrpc:"account_peppol_proxy_state,omitempty"`
	ActiveParentCompany                 *Many2One   `xmlrpc:"active_parent_company,omitempty"`
	ActiveParentCompanyName             *String     `xmlrpc:"active_parent_company_name,omitempty"`
	CanUseParentConnection              *Bool       `xmlrpc:"can_use_parent_connection,omitempty"`
	CompanyId                           *Many2One   `xmlrpc:"company_id,omitempty"`
	ContactEmail                        *String     `xmlrpc:"contact_email,omitempty"`
	CreateDate                          *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                           *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayItsmeLogin                   *Bool       `xmlrpc:"display_itsme_login,omitempty"`
	DisplayName                         *String     `xmlrpc:"display_name,omitempty"`
	DisplayNoAuthButtons                *Bool       `xmlrpc:"display_no_auth_buttons,omitempty"`
	DisplayUseParentConnectionSelection *Bool       `xmlrpc:"display_use_parent_connection_selection,omitempty"`
	EdiMode                             *Selection  `xmlrpc:"edi_mode,omitempty"`
	EdiUserId                           *Many2One   `xmlrpc:"edi_user_id,omitempty"`
	Id                                  *Int        `xmlrpc:"id,omitempty"`
	IsBranchCompany                     *Bool       `xmlrpc:"is_branch_company,omitempty"`
	ParentCompanyId                     *Many2One   `xmlrpc:"parent_company_id,omitempty"`
	ParentCompanyName                   *String     `xmlrpc:"parent_company_name,omitempty"`
	PeppolCanConnectData                interface{} `xmlrpc:"peppol_can_connect_data,omitempty"`
	PeppolEas                           *Selection  `xmlrpc:"peppol_eas,omitempty"`
	PeppolEndpoint                      *String     `xmlrpc:"peppol_endpoint,omitempty"`
	PeppolExternalProvider              *String     `xmlrpc:"peppol_external_provider,omitempty"`
	PeppolWarnings                      interface{} `xmlrpc:"peppol_warnings,omitempty"`
	PhoneNumber                         *String     `xmlrpc:"phone_number,omitempty"`
	SelectedCompanyId                   *Many2One   `xmlrpc:"selected_company_id,omitempty"`
	SmpRegistration                     *Bool       `xmlrpc:"smp_registration,omitempty"`
	UseParentConnection                 *Bool       `xmlrpc:"use_parent_connection,omitempty"`
	UseParentConnectionSelection        *Selection  `xmlrpc:"use_parent_connection_selection,omitempty"`
	WriteDate                           *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                            *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// PeppolRegistrations represents array of peppol.registration model.
type PeppolRegistrations []PeppolRegistration

// PeppolRegistrationModel is the odoo model name.
const PeppolRegistrationModel = "peppol.registration"

// Many2One convert PeppolRegistration to *Many2One.
func (pr *PeppolRegistration) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreatePeppolRegistration creates a new peppol.registration model and returns its id.
func (c *Client) CreatePeppolRegistration(pr *PeppolRegistration) (int64, error) {
	ids, err := c.CreatePeppolRegistrations([]*PeppolRegistration{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePeppolRegistration creates a new peppol.registration model and returns its id.
func (c *Client) CreatePeppolRegistrations(prs []*PeppolRegistration) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(PeppolRegistrationModel, vv, nil)
}

// UpdatePeppolRegistration updates an existing peppol.registration record.
func (c *Client) UpdatePeppolRegistration(pr *PeppolRegistration) error {
	return c.UpdatePeppolRegistrations([]int64{pr.Id.Get()}, pr)
}

// UpdatePeppolRegistrations updates existing peppol.registration records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdatePeppolRegistrations(ids []int64, pr *PeppolRegistration) error {
	return c.Update(PeppolRegistrationModel, ids, pr, nil)
}

// DeletePeppolRegistration deletes an existing peppol.registration record.
func (c *Client) DeletePeppolRegistration(id int64) error {
	return c.DeletePeppolRegistrations([]int64{id})
}

// DeletePeppolRegistrations deletes existing peppol.registration records.
func (c *Client) DeletePeppolRegistrations(ids []int64) error {
	return c.Delete(PeppolRegistrationModel, ids)
}

// GetPeppolRegistration gets peppol.registration existing record.
func (c *Client) GetPeppolRegistration(id int64) (*PeppolRegistration, error) {
	prs, err := c.GetPeppolRegistrations([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetPeppolRegistrations gets peppol.registration existing records.
func (c *Client) GetPeppolRegistrations(ids []int64) (*PeppolRegistrations, error) {
	prs := &PeppolRegistrations{}
	if err := c.Read(PeppolRegistrationModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPeppolRegistration finds peppol.registration record by querying it with criteria.
func (c *Client) FindPeppolRegistration(criteria *Criteria) (*PeppolRegistration, error) {
	prs := &PeppolRegistrations{}
	if err := c.SearchRead(PeppolRegistrationModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindPeppolRegistrations finds peppol.registration records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPeppolRegistrations(criteria *Criteria, options *Options) (*PeppolRegistrations, error) {
	prs := &PeppolRegistrations{}
	if err := c.SearchRead(PeppolRegistrationModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindPeppolRegistrationIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPeppolRegistrationIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PeppolRegistrationModel, criteria, options)
}

// FindPeppolRegistrationId finds record id by querying it with criteria.
func (c *Client) FindPeppolRegistrationId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PeppolRegistrationModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
