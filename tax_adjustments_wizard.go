package odoo

// TaxAdjustmentsWizard represents tax.adjustments.wizard model.
type TaxAdjustmentsWizard struct {
	LastUpdate        *Time      `xmlrpc:"__last_update,omitempty"`
	AdjustmentType    *Selection `xmlrpc:"adjustment_type,omitempty"`
	Amount            *Float     `xmlrpc:"amount,omitempty"`
	CompanyCurrencyId *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	CreditAccountId   *Many2One  `xmlrpc:"credit_account_id,omitempty"`
	Date              *Time      `xmlrpc:"date,omitempty"`
	DebitAccountId    *Many2One  `xmlrpc:"debit_account_id,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	JournalId         *Many2One  `xmlrpc:"journal_id,omitempty"`
	Reason            *String    `xmlrpc:"reason,omitempty"`
	TaxId             *Many2One  `xmlrpc:"tax_id,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// TaxAdjustmentsWizards represents array of tax.adjustments.wizard model.
type TaxAdjustmentsWizards []TaxAdjustmentsWizard

// TaxAdjustmentsWizardModel is the odoo model name.
const TaxAdjustmentsWizardModel = "tax.adjustments.wizard"

// Many2One convert TaxAdjustmentsWizard to *Many2One.
func (taw *TaxAdjustmentsWizard) Many2One() *Many2One {
	return NewMany2One(taw.Id.Get(), "")
}

// CreateTaxAdjustmentsWizard creates a new tax.adjustments.wizard model and returns its id.
func (c *Client) CreateTaxAdjustmentsWizard(taw *TaxAdjustmentsWizard) (int64, error) {
	ids, err := c.CreateTaxAdjustmentsWizards([]*TaxAdjustmentsWizard{taw})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateTaxAdjustmentsWizard creates a new tax.adjustments.wizard model and returns its id.
func (c *Client) CreateTaxAdjustmentsWizards(taws []*TaxAdjustmentsWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range taws {
		vv = append(vv, v)
	}
	return c.Create(TaxAdjustmentsWizardModel, vv, nil)
}

// UpdateTaxAdjustmentsWizard updates an existing tax.adjustments.wizard record.
func (c *Client) UpdateTaxAdjustmentsWizard(taw *TaxAdjustmentsWizard) error {
	return c.UpdateTaxAdjustmentsWizards([]int64{taw.Id.Get()}, taw)
}

// UpdateTaxAdjustmentsWizards updates existing tax.adjustments.wizard records.
// All records (represented by ids) will be updated by taw values.
func (c *Client) UpdateTaxAdjustmentsWizards(ids []int64, taw *TaxAdjustmentsWizard) error {
	return c.Update(TaxAdjustmentsWizardModel, ids, taw, nil)
}

// DeleteTaxAdjustmentsWizard deletes an existing tax.adjustments.wizard record.
func (c *Client) DeleteTaxAdjustmentsWizard(id int64) error {
	return c.DeleteTaxAdjustmentsWizards([]int64{id})
}

// DeleteTaxAdjustmentsWizards deletes existing tax.adjustments.wizard records.
func (c *Client) DeleteTaxAdjustmentsWizards(ids []int64) error {
	return c.Delete(TaxAdjustmentsWizardModel, ids)
}

// GetTaxAdjustmentsWizard gets tax.adjustments.wizard existing record.
func (c *Client) GetTaxAdjustmentsWizard(id int64) (*TaxAdjustmentsWizard, error) {
	taws, err := c.GetTaxAdjustmentsWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*taws)[0]), nil
}

// GetTaxAdjustmentsWizards gets tax.adjustments.wizard existing records.
func (c *Client) GetTaxAdjustmentsWizards(ids []int64) (*TaxAdjustmentsWizards, error) {
	taws := &TaxAdjustmentsWizards{}
	if err := c.Read(TaxAdjustmentsWizardModel, ids, nil, taws); err != nil {
		return nil, err
	}
	return taws, nil
}

// FindTaxAdjustmentsWizard finds tax.adjustments.wizard record by querying it with criteria.
func (c *Client) FindTaxAdjustmentsWizard(criteria *Criteria) (*TaxAdjustmentsWizard, error) {
	taws := &TaxAdjustmentsWizards{}
	if err := c.SearchRead(TaxAdjustmentsWizardModel, criteria, NewOptions().Limit(1), taws); err != nil {
		return nil, err
	}
	return &((*taws)[0]), nil
}

// FindTaxAdjustmentsWizards finds tax.adjustments.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindTaxAdjustmentsWizards(criteria *Criteria, options *Options) (*TaxAdjustmentsWizards, error) {
	taws := &TaxAdjustmentsWizards{}
	if err := c.SearchRead(TaxAdjustmentsWizardModel, criteria, options, taws); err != nil {
		return nil, err
	}
	return taws, nil
}

// FindTaxAdjustmentsWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindTaxAdjustmentsWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(TaxAdjustmentsWizardModel, criteria, options)
}

// FindTaxAdjustmentsWizardId finds record id by querying it with criteria.
func (c *Client) FindTaxAdjustmentsWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(TaxAdjustmentsWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
