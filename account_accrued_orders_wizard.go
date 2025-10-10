package odoo

// AccountAccruedOrdersWizard represents account.accrued.orders.wizard model.
type AccountAccruedOrdersWizard struct {
	AccountId     *Many2One `xmlrpc:"account_id,omitempty"`
	Amount        *Float    `xmlrpc:"amount,omitempty"`
	CompanyId     *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate    *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid     *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId    *Many2One `xmlrpc:"currency_id,omitempty"`
	Date          *Time     `xmlrpc:"date,omitempty"`
	DisplayAmount *Bool     `xmlrpc:"display_amount,omitempty"`
	DisplayName   *String   `xmlrpc:"display_name,omitempty"`
	Id            *Int      `xmlrpc:"id,omitempty"`
	JournalId     *Many2One `xmlrpc:"journal_id,omitempty"`
	PreviewData   *String   `xmlrpc:"preview_data,omitempty"`
	ReversalDate  *Time     `xmlrpc:"reversal_date,omitempty"`
	WriteDate     *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid      *Many2One `xmlrpc:"write_uid,omitempty"`
}

// AccountAccruedOrdersWizards represents array of account.accrued.orders.wizard model.
type AccountAccruedOrdersWizards []AccountAccruedOrdersWizard

// AccountAccruedOrdersWizardModel is the odoo model name.
const AccountAccruedOrdersWizardModel = "account.accrued.orders.wizard"

// Many2One convert AccountAccruedOrdersWizard to *Many2One.
func (aaow *AccountAccruedOrdersWizard) Many2One() *Many2One {
	return NewMany2One(aaow.Id.Get(), "")
}

// CreateAccountAccruedOrdersWizard creates a new account.accrued.orders.wizard model and returns its id.
func (c *Client) CreateAccountAccruedOrdersWizard(aaow *AccountAccruedOrdersWizard) (int64, error) {
	ids, err := c.CreateAccountAccruedOrdersWizards([]*AccountAccruedOrdersWizard{aaow})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAccruedOrdersWizard creates a new account.accrued.orders.wizard model and returns its id.
func (c *Client) CreateAccountAccruedOrdersWizards(aaows []*AccountAccruedOrdersWizard) ([]int64, error) {
	var vv []interface{}
	for _, v := range aaows {
		vv = append(vv, v)
	}
	return c.Create(AccountAccruedOrdersWizardModel, vv, nil)
}

// UpdateAccountAccruedOrdersWizard updates an existing account.accrued.orders.wizard record.
func (c *Client) UpdateAccountAccruedOrdersWizard(aaow *AccountAccruedOrdersWizard) error {
	return c.UpdateAccountAccruedOrdersWizards([]int64{aaow.Id.Get()}, aaow)
}

// UpdateAccountAccruedOrdersWizards updates existing account.accrued.orders.wizard records.
// All records (represented by ids) will be updated by aaow values.
func (c *Client) UpdateAccountAccruedOrdersWizards(ids []int64, aaow *AccountAccruedOrdersWizard) error {
	return c.Update(AccountAccruedOrdersWizardModel, ids, aaow, nil)
}

// DeleteAccountAccruedOrdersWizard deletes an existing account.accrued.orders.wizard record.
func (c *Client) DeleteAccountAccruedOrdersWizard(id int64) error {
	return c.DeleteAccountAccruedOrdersWizards([]int64{id})
}

// DeleteAccountAccruedOrdersWizards deletes existing account.accrued.orders.wizard records.
func (c *Client) DeleteAccountAccruedOrdersWizards(ids []int64) error {
	return c.Delete(AccountAccruedOrdersWizardModel, ids)
}

// GetAccountAccruedOrdersWizard gets account.accrued.orders.wizard existing record.
func (c *Client) GetAccountAccruedOrdersWizard(id int64) (*AccountAccruedOrdersWizard, error) {
	aaows, err := c.GetAccountAccruedOrdersWizards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aaows)[0]), nil
}

// GetAccountAccruedOrdersWizards gets account.accrued.orders.wizard existing records.
func (c *Client) GetAccountAccruedOrdersWizards(ids []int64) (*AccountAccruedOrdersWizards, error) {
	aaows := &AccountAccruedOrdersWizards{}
	if err := c.Read(AccountAccruedOrdersWizardModel, ids, nil, aaows); err != nil {
		return nil, err
	}
	return aaows, nil
}

// FindAccountAccruedOrdersWizard finds account.accrued.orders.wizard record by querying it with criteria.
func (c *Client) FindAccountAccruedOrdersWizard(criteria *Criteria) (*AccountAccruedOrdersWizard, error) {
	aaows := &AccountAccruedOrdersWizards{}
	if err := c.SearchRead(AccountAccruedOrdersWizardModel, criteria, NewOptions().Limit(1), aaows); err != nil {
		return nil, err
	}
	return &((*aaows)[0]), nil
}

// FindAccountAccruedOrdersWizards finds account.accrued.orders.wizard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccruedOrdersWizards(criteria *Criteria, options *Options) (*AccountAccruedOrdersWizards, error) {
	aaows := &AccountAccruedOrdersWizards{}
	if err := c.SearchRead(AccountAccruedOrdersWizardModel, criteria, options, aaows); err != nil {
		return nil, err
	}
	return aaows, nil
}

// FindAccountAccruedOrdersWizardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccruedOrdersWizardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAccruedOrdersWizardModel, criteria, options)
}

// FindAccountAccruedOrdersWizardId finds record id by querying it with criteria.
func (c *Client) FindAccountAccruedOrdersWizardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccruedOrdersWizardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
