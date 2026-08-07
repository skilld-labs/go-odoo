package odoo

// L10NFrFecExportWizard represents l10n_fr.fec.export.wizard model.
type L10NFrFecExportWizard struct {
	CreateDate         *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateFrom           *Time      `xmlrpc:"date_from,omitempty"`
	DateTo             *Time      `xmlrpc:"date_to,omitempty"`
	DisplayName        *String    `xmlrpc:"display_name,omitempty"`
	ExcludedJournalIds *Relation  `xmlrpc:"excluded_journal_ids,omitempty"`
	ExportType         *Selection `xmlrpc:"export_type,omitempty"`
	Filename           *String    `xmlrpc:"filename,omitempty"`
	Id                 *Int       `xmlrpc:"id,omitempty"`
	TestFile           *Bool      `xmlrpc:"test_file,omitempty"`
	WriteDate          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// L10NFrFecExportWizards represents array of l10n_fr.fec.export.wizard model.
type L10NFrFecExportWizards []L10NFrFecExportWizard

// L10NFrFecExportWizardModel is the odoo model name.
const L10NFrFecExportWizardModel = "l10n_fr.fec.export.wizard"

// Many2One convert L10NFrFecExportWizard to *Many2One.
func (lfew *L10NFrFecExportWizard) Many2One() *Many2One {
	return NewMany2One(lfew.Id.Get(), "")
}

// CreateL10NFrFecExportWizard creates a new l10n_fr.fec.export.wizard model and returns its id.
func (c *Client) CreateL10NFrFecExportWizard(lfew *L10NFrFecExportWizard) (int64, error) {
	ids, err := c.CreateL10NFrFecExportWizards([]*L10NFrFecExportWizard{lfew})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NFrFecExportWizard creates a new l10n_fr.fec.export.wizard model and returns its id.
func (c *Client) CreateL10NFrFecExportWizards(lfews []*L10NFrFecExportWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range lfews {
		vv = append(vv, v)
	}
	return c.Create(L10NFrFecExportWizardModel, vv, nil)
}

// UpdateL10NFrFecExportWizard updates an existing l10n_fr.fec.export.wizard record.
func (c *Client) UpdateL10NFrFecExportWizard(lfew *L10NFrFecExportWizard) error {
	return c.UpdateL10NFrFecExportWizards([]int64{lfew.Id.Get()}, lfew)
}

// UpdateL10NFrFecExportWizards updates existing l10n_fr.fec.export.wizard records.
// All records (represented by ids) will be updated by lfew values.
func (c *Client) UpdateL10NFrFecExportWizards(ids []int64, lfew *L10NFrFecExportWizard) error {
	return c.Update(L10NFrFecExportWizardModel, ids, lfew, nil)
}

// DeleteL10NFrFecExportWizard deletes an existing l10n_fr.fec.export.wizard record.
func (c *Client) DeleteL10NFrFecExportWizard(id int64) error {
	return c.DeleteL10NFrFecExportWizards([]int64{id})
}

// DeleteL10NFrFecExportWizards deletes existing l10n_fr.fec.export.wizard records.
func (c *Client) DeleteL10NFrFecExportWizards(ids []int64) error {
	return c.Delete(L10NFrFecExportWizardModel, ids)
}

// GetL10NFrFecExportWizard gets l10n_fr.fec.export.wizard existing record.
func (c *Client) GetL10NFrFecExportWizard(id int64) (*L10NFrFecExportWizard, error) {
	lfews, err := c.GetL10NFrFecExportWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lfews)[0]), nil
}

// GetL10NFrFecExportWizards gets l10n_fr.fec.export.wizard existing records.
func (c *Client) GetL10NFrFecExportWizards(ids []int64) (*L10NFrFecExportWizards, error) {
	lfews := &L10NFrFecExportWizards{}
	if err := c.Read(L10NFrFecExportWizardModel, ids, nil, lfews); err != nil {
		return nil, err
	}
	return lfews, nil
}

// FindL10NFrFecExportWizard finds l10n_fr.fec.export.wizard record by querying it with criteria.
func (c *Client) FindL10NFrFecExportWizard(criteria *Criteria) (*L10NFrFecExportWizard, error) {
	lfews := &L10NFrFecExportWizards{}
	if err := c.SearchRead(L10NFrFecExportWizardModel, criteria, NewOptions().Limit(1), lfews); err != nil {
		return nil, err
	}
	return &((*lfews)[0]), nil
}

// FindL10NFrFecExportWizards finds l10n_fr.fec.export.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrFecExportWizards(criteria *Criteria, options *Options) (*L10NFrFecExportWizards, error) {
	lfews := &L10NFrFecExportWizards{}
	if err := c.SearchRead(L10NFrFecExportWizardModel, criteria, options, lfews); err != nil {
		return nil, err
	}
	return lfews, nil
}

// FindL10NFrFecExportWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrFecExportWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NFrFecExportWizardModel, criteria, options)
}

// FindL10NFrFecExportWizardId finds record id by querying it with criteria.
func (c *Client) FindL10NFrFecExportWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NFrFecExportWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
