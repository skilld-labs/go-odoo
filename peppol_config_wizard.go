package odoo

// PeppolConfigWizard represents peppol.config.wizard model.
type PeppolConfigWizard struct {
	AccountPeppolContactEmail           *String     `xmlrpc:"account_peppol_contact_email,omitempty"`
	AccountPeppolEdiIdentification      *String     `xmlrpc:"account_peppol_edi_identification,omitempty"`
	AccountPeppolEdiUser                *Many2One   `xmlrpc:"account_peppol_edi_user,omitempty"`
	AccountPeppolMigrationKey           *String     `xmlrpc:"account_peppol_migration_key,omitempty"`
	AccountPeppolProxyState             *Selection  `xmlrpc:"account_peppol_proxy_state,omitempty"`
	CompanyId                           *Many2One   `xmlrpc:"company_id,omitempty"`
	CreateDate                          *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                           *Many2One   `xmlrpc:"create_uid,omitempty"`
	DisplayName                         *String     `xmlrpc:"display_name,omitempty"`
	Id                                  *Int        `xmlrpc:"id,omitempty"`
	PeppolActivateSelfBilling           *Bool       `xmlrpc:"peppol_activate_self_billing,omitempty"`
	PeppolSelfBillingReceptionJournalId *Many2One   `xmlrpc:"peppol_self_billing_reception_journal_id,omitempty"`
	ServiceIds                          *Relation   `xmlrpc:"service_ids,omitempty"`
	ServiceInfo                         *String     `xmlrpc:"service_info,omitempty"`
	ServiceJson                         interface{} `xmlrpc:"service_json,omitempty"`
	WriteDate                           *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                            *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// PeppolConfigWizards represents array of peppol.config.wizard model.
type PeppolConfigWizards []PeppolConfigWizard

// PeppolConfigWizardModel is the odoo model name.
const PeppolConfigWizardModel = "peppol.config.wizard"

// Many2One convert PeppolConfigWizard to *Many2One.
func (pcw *PeppolConfigWizard) Many2One() *Many2One {
	return NewMany2One(pcw.Id.Get(), "")
}

// CreatePeppolConfigWizard creates a new peppol.config.wizard model and returns its id.
func (c *Client) CreatePeppolConfigWizard(pcw *PeppolConfigWizard) (int64, error) {
	ids, err := c.CreatePeppolConfigWizards([]*PeppolConfigWizard{pcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePeppolConfigWizard creates a new peppol.config.wizard model and returns its id.
func (c *Client) CreatePeppolConfigWizards(pcws []*PeppolConfigWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcws {
		vv = append(vv, v)
	}
	return c.Create(PeppolConfigWizardModel, vv, nil)
}

// UpdatePeppolConfigWizard updates an existing peppol.config.wizard record.
func (c *Client) UpdatePeppolConfigWizard(pcw *PeppolConfigWizard) error {
	return c.UpdatePeppolConfigWizards([]int64{pcw.Id.Get()}, pcw)
}

// UpdatePeppolConfigWizards updates existing peppol.config.wizard records.
// All records (represented by ids) will be updated by pcw values.
func (c *Client) UpdatePeppolConfigWizards(ids []int64, pcw *PeppolConfigWizard) error {
	return c.Update(PeppolConfigWizardModel, ids, pcw, nil)
}

// DeletePeppolConfigWizard deletes an existing peppol.config.wizard record.
func (c *Client) DeletePeppolConfigWizard(id int64) error {
	return c.DeletePeppolConfigWizards([]int64{id})
}

// DeletePeppolConfigWizards deletes existing peppol.config.wizard records.
func (c *Client) DeletePeppolConfigWizards(ids []int64) error {
	return c.Delete(PeppolConfigWizardModel, ids)
}

// GetPeppolConfigWizard gets peppol.config.wizard existing record.
func (c *Client) GetPeppolConfigWizard(id int64) (*PeppolConfigWizard, error) {
	pcws, err := c.GetPeppolConfigWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcws)[0]), nil
}

// GetPeppolConfigWizards gets peppol.config.wizard existing records.
func (c *Client) GetPeppolConfigWizards(ids []int64) (*PeppolConfigWizards, error) {
	pcws := &PeppolConfigWizards{}
	if err := c.Read(PeppolConfigWizardModel, ids, nil, pcws); err != nil {
		return nil, err
	}
	return pcws, nil
}

// FindPeppolConfigWizard finds peppol.config.wizard record by querying it with criteria.
func (c *Client) FindPeppolConfigWizard(criteria *Criteria) (*PeppolConfigWizard, error) {
	pcws := &PeppolConfigWizards{}
	if err := c.SearchRead(PeppolConfigWizardModel, criteria, NewOptions().Limit(1), pcws); err != nil {
		return nil, err
	}
	return &((*pcws)[0]), nil
}

// FindPeppolConfigWizards finds peppol.config.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPeppolConfigWizards(criteria *Criteria, options *Options) (*PeppolConfigWizards, error) {
	pcws := &PeppolConfigWizards{}
	if err := c.SearchRead(PeppolConfigWizardModel, criteria, options, pcws); err != nil {
		return nil, err
	}
	return pcws, nil
}

// FindPeppolConfigWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPeppolConfigWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PeppolConfigWizardModel, criteria, options)
}

// FindPeppolConfigWizardId finds record id by querying it with criteria.
func (c *Client) FindPeppolConfigWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PeppolConfigWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
