package odoo

// L10NFrPdpReportsSendWizard represents l10n.fr.pdp.reports.send.wizard model.
type L10NFrPdpReportsSendWizard struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	FlowId      *Many2One `xmlrpc:"flow_id,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Warning     *String   `xmlrpc:"warning,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// L10NFrPdpReportsSendWizards represents array of l10n.fr.pdp.reports.send.wizard model.
type L10NFrPdpReportsSendWizards []L10NFrPdpReportsSendWizard

// L10NFrPdpReportsSendWizardModel is the odoo model name.
const L10NFrPdpReportsSendWizardModel = "l10n.fr.pdp.reports.send.wizard"

// Many2One convert L10NFrPdpReportsSendWizard to *Many2One.
func (lfprsw *L10NFrPdpReportsSendWizard) Many2One() *Many2One {
	return NewMany2One(lfprsw.Id.Get(), "")
}

// CreateL10NFrPdpReportsSendWizard creates a new l10n.fr.pdp.reports.send.wizard model and returns its id.
func (c *Client) CreateL10NFrPdpReportsSendWizard(lfprsw *L10NFrPdpReportsSendWizard) (int64, error) {
	ids, err := c.CreateL10NFrPdpReportsSendWizards([]*L10NFrPdpReportsSendWizard{lfprsw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NFrPdpReportsSendWizard creates a new l10n.fr.pdp.reports.send.wizard model and returns its id.
func (c *Client) CreateL10NFrPdpReportsSendWizards(lfprsws []*L10NFrPdpReportsSendWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range lfprsws {
		vv = append(vv, v)
	}
	return c.Create(L10NFrPdpReportsSendWizardModel, vv, nil)
}

// UpdateL10NFrPdpReportsSendWizard updates an existing l10n.fr.pdp.reports.send.wizard record.
func (c *Client) UpdateL10NFrPdpReportsSendWizard(lfprsw *L10NFrPdpReportsSendWizard) error {
	return c.UpdateL10NFrPdpReportsSendWizards([]int64{lfprsw.Id.Get()}, lfprsw)
}

// UpdateL10NFrPdpReportsSendWizards updates existing l10n.fr.pdp.reports.send.wizard records.
// All records (represented by ids) will be updated by lfprsw values.
func (c *Client) UpdateL10NFrPdpReportsSendWizards(ids []int64, lfprsw *L10NFrPdpReportsSendWizard) error {
	return c.Update(L10NFrPdpReportsSendWizardModel, ids, lfprsw, nil)
}

// DeleteL10NFrPdpReportsSendWizard deletes an existing l10n.fr.pdp.reports.send.wizard record.
func (c *Client) DeleteL10NFrPdpReportsSendWizard(id int64) error {
	return c.DeleteL10NFrPdpReportsSendWizards([]int64{id})
}

// DeleteL10NFrPdpReportsSendWizards deletes existing l10n.fr.pdp.reports.send.wizard records.
func (c *Client) DeleteL10NFrPdpReportsSendWizards(ids []int64) error {
	return c.Delete(L10NFrPdpReportsSendWizardModel, ids)
}

// GetL10NFrPdpReportsSendWizard gets l10n.fr.pdp.reports.send.wizard existing record.
func (c *Client) GetL10NFrPdpReportsSendWizard(id int64) (*L10NFrPdpReportsSendWizard, error) {
	lfprsws, err := c.GetL10NFrPdpReportsSendWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lfprsws)[0]), nil
}

// GetL10NFrPdpReportsSendWizards gets l10n.fr.pdp.reports.send.wizard existing records.
func (c *Client) GetL10NFrPdpReportsSendWizards(ids []int64) (*L10NFrPdpReportsSendWizards, error) {
	lfprsws := &L10NFrPdpReportsSendWizards{}
	if err := c.Read(L10NFrPdpReportsSendWizardModel, ids, nil, lfprsws); err != nil {
		return nil, err
	}
	return lfprsws, nil
}

// FindL10NFrPdpReportsSendWizard finds l10n.fr.pdp.reports.send.wizard record by querying it with criteria.
func (c *Client) FindL10NFrPdpReportsSendWizard(criteria *Criteria) (*L10NFrPdpReportsSendWizard, error) {
	lfprsws := &L10NFrPdpReportsSendWizards{}
	if err := c.SearchRead(L10NFrPdpReportsSendWizardModel, criteria, NewOptions().Limit(1), lfprsws); err != nil {
		return nil, err
	}
	return &((*lfprsws)[0]), nil
}

// FindL10NFrPdpReportsSendWizards finds l10n.fr.pdp.reports.send.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrPdpReportsSendWizards(criteria *Criteria, options *Options) (*L10NFrPdpReportsSendWizards, error) {
	lfprsws := &L10NFrPdpReportsSendWizards{}
	if err := c.SearchRead(L10NFrPdpReportsSendWizardModel, criteria, options, lfprsws); err != nil {
		return nil, err
	}
	return lfprsws, nil
}

// FindL10NFrPdpReportsSendWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrPdpReportsSendWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NFrPdpReportsSendWizardModel, criteria, options)
}

// FindL10NFrPdpReportsSendWizardId finds record id by querying it with criteria.
func (c *Client) FindL10NFrPdpReportsSendWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NFrPdpReportsSendWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
