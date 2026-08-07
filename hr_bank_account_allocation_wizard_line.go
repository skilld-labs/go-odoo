package odoo

// HrBankAccountAllocationWizardLine represents hr.bank.account.allocation.wizard.line model.
type HrBankAccountAllocationWizardLine struct {
	AccNumber     *String    `xmlrpc:"acc_number,omitempty"`
	Amount        *Float     `xmlrpc:"amount,omitempty"`
	AmountType    *Selection `xmlrpc:"amount_type,omitempty"`
	BankAccountId *Many2One  `xmlrpc:"bank_account_id,omitempty"`
	CreateDate    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName   *String    `xmlrpc:"display_name,omitempty"`
	Id            *Int       `xmlrpc:"id,omitempty"`
	Sequence      *Int       `xmlrpc:"sequence,omitempty"`
	Symbol        *String    `xmlrpc:"symbol,omitempty"`
	Trusted       *Bool      `xmlrpc:"trusted,omitempty"`
	WizardId      *Many2One  `xmlrpc:"wizard_id,omitempty"`
	WriteDate     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrBankAccountAllocationWizardLines represents array of hr.bank.account.allocation.wizard.line model.
type HrBankAccountAllocationWizardLines []HrBankAccountAllocationWizardLine

// HrBankAccountAllocationWizardLineModel is the odoo model name.
const HrBankAccountAllocationWizardLineModel = "hr.bank.account.allocation.wizard.line"

// Many2One convert HrBankAccountAllocationWizardLine to *Many2One.
func (hbaawl *HrBankAccountAllocationWizardLine) Many2One() *Many2One {
	return NewMany2One(hbaawl.Id.Get(), "")
}

// CreateHrBankAccountAllocationWizardLine creates a new hr.bank.account.allocation.wizard.line model and returns its id.
func (c *Client) CreateHrBankAccountAllocationWizardLine(hbaawl *HrBankAccountAllocationWizardLine) (int64, error) {
	ids, err := c.CreateHrBankAccountAllocationWizardLines([]*HrBankAccountAllocationWizardLine{hbaawl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrBankAccountAllocationWizardLine creates a new hr.bank.account.allocation.wizard.line model and returns its id.
func (c *Client) CreateHrBankAccountAllocationWizardLines(hbaawls []*HrBankAccountAllocationWizardLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range hbaawls {
		vv = append(vv, v)
	}
	return c.Create(HrBankAccountAllocationWizardLineModel, vv, nil)
}

// UpdateHrBankAccountAllocationWizardLine updates an existing hr.bank.account.allocation.wizard.line record.
func (c *Client) UpdateHrBankAccountAllocationWizardLine(hbaawl *HrBankAccountAllocationWizardLine) error {
	return c.UpdateHrBankAccountAllocationWizardLines([]int64{hbaawl.Id.Get()}, hbaawl)
}

// UpdateHrBankAccountAllocationWizardLines updates existing hr.bank.account.allocation.wizard.line records.
// All records (represented by ids) will be updated by hbaawl values.
func (c *Client) UpdateHrBankAccountAllocationWizardLines(ids []int64, hbaawl *HrBankAccountAllocationWizardLine) error {
	return c.Update(HrBankAccountAllocationWizardLineModel, ids, hbaawl, nil)
}

// DeleteHrBankAccountAllocationWizardLine deletes an existing hr.bank.account.allocation.wizard.line record.
func (c *Client) DeleteHrBankAccountAllocationWizardLine(id int64) error {
	return c.DeleteHrBankAccountAllocationWizardLines([]int64{id})
}

// DeleteHrBankAccountAllocationWizardLines deletes existing hr.bank.account.allocation.wizard.line records.
func (c *Client) DeleteHrBankAccountAllocationWizardLines(ids []int64) error {
	return c.Delete(HrBankAccountAllocationWizardLineModel, ids)
}

// GetHrBankAccountAllocationWizardLine gets hr.bank.account.allocation.wizard.line existing record.
func (c *Client) GetHrBankAccountAllocationWizardLine(id int64) (*HrBankAccountAllocationWizardLine, error) {
	hbaawls, err := c.GetHrBankAccountAllocationWizardLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hbaawls)[0]), nil
}

// GetHrBankAccountAllocationWizardLines gets hr.bank.account.allocation.wizard.line existing records.
func (c *Client) GetHrBankAccountAllocationWizardLines(ids []int64) (*HrBankAccountAllocationWizardLines, error) {
	hbaawls := &HrBankAccountAllocationWizardLines{}
	if err := c.Read(HrBankAccountAllocationWizardLineModel, ids, nil, hbaawls); err != nil {
		return nil, err
	}
	return hbaawls, nil
}

// FindHrBankAccountAllocationWizardLine finds hr.bank.account.allocation.wizard.line record by querying it with criteria.
func (c *Client) FindHrBankAccountAllocationWizardLine(criteria *Criteria) (*HrBankAccountAllocationWizardLine, error) {
	hbaawls := &HrBankAccountAllocationWizardLines{}
	if err := c.SearchRead(HrBankAccountAllocationWizardLineModel, criteria, NewOptions().Limit(1), hbaawls); err != nil {
		return nil, err
	}
	return &((*hbaawls)[0]), nil
}

// FindHrBankAccountAllocationWizardLines finds hr.bank.account.allocation.wizard.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrBankAccountAllocationWizardLines(criteria *Criteria, options *Options) (*HrBankAccountAllocationWizardLines, error) {
	hbaawls := &HrBankAccountAllocationWizardLines{}
	if err := c.SearchRead(HrBankAccountAllocationWizardLineModel, criteria, options, hbaawls); err != nil {
		return nil, err
	}
	return hbaawls, nil
}

// FindHrBankAccountAllocationWizardLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrBankAccountAllocationWizardLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrBankAccountAllocationWizardLineModel, criteria, options)
}

// FindHrBankAccountAllocationWizardLineId finds record id by querying it with criteria.
func (c *Client) FindHrBankAccountAllocationWizardLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrBankAccountAllocationWizardLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
