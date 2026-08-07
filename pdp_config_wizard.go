package odoo

// PdpConfigWizard represents pdp.config.wizard model.
type PdpConfigWizard struct {
	AccountPeppolContactEmail      *String    `xmlrpc:"account_peppol_contact_email,omitempty"`
	AccountPeppolEdiIdentification *String    `xmlrpc:"account_peppol_edi_identification,omitempty"`
	AccountPeppolEdiUser           *Many2One  `xmlrpc:"account_peppol_edi_user,omitempty"`
	AccountPeppolProxyState        *Selection `xmlrpc:"account_peppol_proxy_state,omitempty"`
	CompanyId                      *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                    *String    `xmlrpc:"display_name,omitempty"`
	Id                             *Int       `xmlrpc:"id,omitempty"`
	WriteDate                      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PdpConfigWizards represents array of pdp.config.wizard model.
type PdpConfigWizards []PdpConfigWizard

// PdpConfigWizardModel is the odoo model name.
const PdpConfigWizardModel = "pdp.config.wizard"

// Many2One convert PdpConfigWizard to *Many2One.
func (pcw *PdpConfigWizard) Many2One() *Many2One {
	return NewMany2One(pcw.Id.Get(), "")
}

// CreatePdpConfigWizard creates a new pdp.config.wizard model and returns its id.
func (c *Client) CreatePdpConfigWizard(pcw *PdpConfigWizard) (int64, error) {
	ids, err := c.CreatePdpConfigWizards([]*PdpConfigWizard{pcw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePdpConfigWizard creates a new pdp.config.wizard model and returns its id.
func (c *Client) CreatePdpConfigWizards(pcws []*PdpConfigWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcws {
		vv = append(vv, v)
	}
	return c.Create(PdpConfigWizardModel, vv, nil)
}

// UpdatePdpConfigWizard updates an existing pdp.config.wizard record.
func (c *Client) UpdatePdpConfigWizard(pcw *PdpConfigWizard) error {
	return c.UpdatePdpConfigWizards([]int64{pcw.Id.Get()}, pcw)
}

// UpdatePdpConfigWizards updates existing pdp.config.wizard records.
// All records (represented by ids) will be updated by pcw values.
func (c *Client) UpdatePdpConfigWizards(ids []int64, pcw *PdpConfigWizard) error {
	return c.Update(PdpConfigWizardModel, ids, pcw, nil)
}

// DeletePdpConfigWizard deletes an existing pdp.config.wizard record.
func (c *Client) DeletePdpConfigWizard(id int64) error {
	return c.DeletePdpConfigWizards([]int64{id})
}

// DeletePdpConfigWizards deletes existing pdp.config.wizard records.
func (c *Client) DeletePdpConfigWizards(ids []int64) error {
	return c.Delete(PdpConfigWizardModel, ids)
}

// GetPdpConfigWizard gets pdp.config.wizard existing record.
func (c *Client) GetPdpConfigWizard(id int64) (*PdpConfigWizard, error) {
	pcws, err := c.GetPdpConfigWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcws)[0]), nil
}

// GetPdpConfigWizards gets pdp.config.wizard existing records.
func (c *Client) GetPdpConfigWizards(ids []int64) (*PdpConfigWizards, error) {
	pcws := &PdpConfigWizards{}
	if err := c.Read(PdpConfigWizardModel, ids, nil, pcws); err != nil {
		return nil, err
	}
	return pcws, nil
}

// FindPdpConfigWizard finds pdp.config.wizard record by querying it with criteria.
func (c *Client) FindPdpConfigWizard(criteria *Criteria) (*PdpConfigWizard, error) {
	pcws := &PdpConfigWizards{}
	if err := c.SearchRead(PdpConfigWizardModel, criteria, NewOptions().Limit(1), pcws); err != nil {
		return nil, err
	}
	return &((*pcws)[0]), nil
}

// FindPdpConfigWizards finds pdp.config.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPdpConfigWizards(criteria *Criteria, options *Options) (*PdpConfigWizards, error) {
	pcws := &PdpConfigWizards{}
	if err := c.SearchRead(PdpConfigWizardModel, criteria, options, pcws); err != nil {
		return nil, err
	}
	return pcws, nil
}

// FindPdpConfigWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPdpConfigWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PdpConfigWizardModel, criteria, options)
}

// FindPdpConfigWizardId finds record id by querying it with criteria.
func (c *Client) FindPdpConfigWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PdpConfigWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
