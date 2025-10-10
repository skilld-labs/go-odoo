package odoo

// AccountPaymentRegister represents account.payment.register model.
type AccountPaymentRegister struct {
	ActionableErrors              interface{} `xmlrpc:"actionable_errors,omitempty"`
	Amount                        *Float      `xmlrpc:"amount,omitempty"`
	AvailableJournalIds           *Relation   `xmlrpc:"available_journal_ids,omitempty"`
	AvailablePartnerBankIds       *Relation   `xmlrpc:"available_partner_bank_ids,omitempty"`
	AvailablePaymentMethodLineIds *Relation   `xmlrpc:"available_payment_method_line_ids,omitempty"`
	Batches                       *String     `xmlrpc:"batches,omitempty"`
	CanEditWizard                 *Bool       `xmlrpc:"can_edit_wizard,omitempty"`
	CanGroupPayments              *Bool       `xmlrpc:"can_group_payments,omitempty"`
	Communication                 *String     `xmlrpc:"communication,omitempty"`
	CompanyCurrencyId             *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                     *Many2One   `xmlrpc:"company_id,omitempty"`
	CountryCode                   *String     `xmlrpc:"country_code,omitempty"`
	CreateDate                    *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                    *Many2One   `xmlrpc:"currency_id,omitempty"`
	CustomUserAmount              *Float      `xmlrpc:"custom_user_amount,omitempty"`
	CustomUserCurrencyId          *Many2One   `xmlrpc:"custom_user_currency_id,omitempty"`
	DisplayName                   *String     `xmlrpc:"display_name,omitempty"`
	DuplicatePaymentIds           *Relation   `xmlrpc:"duplicate_payment_ids,omitempty"`
	EarlyPaymentDiscountMode      *Bool       `xmlrpc:"early_payment_discount_mode,omitempty"`
	GroupPayment                  *Bool       `xmlrpc:"group_payment,omitempty"`
	HideWriteoffSection           *Bool       `xmlrpc:"hide_writeoff_section,omitempty"`
	Id                            *Int        `xmlrpc:"id,omitempty"`
	InstallmentsMode              *Selection  `xmlrpc:"installments_mode,omitempty"`
	InstallmentsSwitchAmount      *Float      `xmlrpc:"installments_switch_amount,omitempty"`
	InstallmentsSwitchHtml        *String     `xmlrpc:"installments_switch_html,omitempty"`
	IsRegisterPaymentOnDraft      *Bool       `xmlrpc:"is_register_payment_on_draft,omitempty"`
	JournalId                     *Many2One   `xmlrpc:"journal_id,omitempty"`
	LineIds                       *Relation   `xmlrpc:"line_ids,omitempty"`
	MissingAccountPartners        *Relation   `xmlrpc:"missing_account_partners,omitempty"`
	PartnerBankId                 *Many2One   `xmlrpc:"partner_bank_id,omitempty"`
	PartnerId                     *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerType                   *Selection  `xmlrpc:"partner_type,omitempty"`
	PaymentDate                   *Time       `xmlrpc:"payment_date,omitempty"`
	PaymentDifference             *Float      `xmlrpc:"payment_difference,omitempty"`
	PaymentDifferenceHandling     *Selection  `xmlrpc:"payment_difference_handling,omitempty"`
	PaymentMethodCode             *String     `xmlrpc:"payment_method_code,omitempty"`
	PaymentMethodLineId           *Many2One   `xmlrpc:"payment_method_line_id,omitempty"`
	PaymentTokenId                *Many2One   `xmlrpc:"payment_token_id,omitempty"`
	PaymentType                   *Selection  `xmlrpc:"payment_type,omitempty"`
	QrCode                        *String     `xmlrpc:"qr_code,omitempty"`
	RequirePartnerBankAccount     *Bool       `xmlrpc:"require_partner_bank_account,omitempty"`
	ShowPartnerBankAccount        *Bool       `xmlrpc:"show_partner_bank_account,omitempty"`
	ShowPaymentDifference         *Bool       `xmlrpc:"show_payment_difference,omitempty"`
	SourceAmount                  *Float      `xmlrpc:"source_amount,omitempty"`
	SourceAmountCurrency          *Float      `xmlrpc:"source_amount_currency,omitempty"`
	SourceCurrencyId              *Many2One   `xmlrpc:"source_currency_id,omitempty"`
	SuitablePaymentTokenIds       *Relation   `xmlrpc:"suitable_payment_token_ids,omitempty"`
	TotalPaymentsAmount           *Int        `xmlrpc:"total_payments_amount,omitempty"`
	UntrustedBankIds              *Relation   `xmlrpc:"untrusted_bank_ids,omitempty"`
	UntrustedPaymentsCount        *Int        `xmlrpc:"untrusted_payments_count,omitempty"`
	UseElectronicPaymentMethod    *Bool       `xmlrpc:"use_electronic_payment_method,omitempty"`
	WriteDate                     *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One   `xmlrpc:"write_uid,omitempty"`
	WriteoffAccountId             *Many2One   `xmlrpc:"writeoff_account_id,omitempty"`
	WriteoffIsExchangeAccount     *Bool       `xmlrpc:"writeoff_is_exchange_account,omitempty"`
	WriteoffLabel                 *String     `xmlrpc:"writeoff_label,omitempty"`
}

// AccountPaymentRegisters represents array of account.payment.register model.
type AccountPaymentRegisters []AccountPaymentRegister

// AccountPaymentRegisterModel is the odoo model name.
const AccountPaymentRegisterModel = "account.payment.register"

// Many2One convert AccountPaymentRegister to *Many2One.
func (apr *AccountPaymentRegister) Many2One() *Many2One {
	return NewMany2One(apr.Id.Get(), "")
}

// CreateAccountPaymentRegister creates a new account.payment.register model and returns its id.
func (c *Client) CreateAccountPaymentRegister(apr *AccountPaymentRegister) (int64, error) {
	ids, err := c.CreateAccountPaymentRegisters([]*AccountPaymentRegister{apr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountPaymentRegister creates a new account.payment.register model and returns its id.
func (c *Client) CreateAccountPaymentRegisters(aprs []*AccountPaymentRegister) ([]int64, error) {
	var vv []interface{}
	for _, v := range aprs {
		vv = append(vv, v)
	}
	return c.Create(AccountPaymentRegisterModel, vv, nil)
}

// UpdateAccountPaymentRegister updates an existing account.payment.register record.
func (c *Client) UpdateAccountPaymentRegister(apr *AccountPaymentRegister) error {
	return c.UpdateAccountPaymentRegisters([]int64{apr.Id.Get()}, apr)
}

// UpdateAccountPaymentRegisters updates existing account.payment.register records.
// All records (represented by ids) will be updated by apr values.
func (c *Client) UpdateAccountPaymentRegisters(ids []int64, apr *AccountPaymentRegister) error {
	return c.Update(AccountPaymentRegisterModel, ids, apr, nil)
}

// DeleteAccountPaymentRegister deletes an existing account.payment.register record.
func (c *Client) DeleteAccountPaymentRegister(id int64) error {
	return c.DeleteAccountPaymentRegisters([]int64{id})
}

// DeleteAccountPaymentRegisters deletes existing account.payment.register records.
func (c *Client) DeleteAccountPaymentRegisters(ids []int64) error {
	return c.Delete(AccountPaymentRegisterModel, ids)
}

// GetAccountPaymentRegister gets account.payment.register existing record.
func (c *Client) GetAccountPaymentRegister(id int64) (*AccountPaymentRegister, error) {
	aprs, err := c.GetAccountPaymentRegisters([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aprs)[0]), nil
}

// GetAccountPaymentRegisters gets account.payment.register existing records.
func (c *Client) GetAccountPaymentRegisters(ids []int64) (*AccountPaymentRegisters, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.Read(AccountPaymentRegisterModel, ids, nil, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPaymentRegister finds account.payment.register record by querying it with criteria.
func (c *Client) FindAccountPaymentRegister(criteria *Criteria) (*AccountPaymentRegister, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.SearchRead(AccountPaymentRegisterModel, criteria, NewOptions().Limit(1), aprs); err != nil {
		return nil, err
	}
	return &((*aprs)[0]), nil
}

// FindAccountPaymentRegisters finds account.payment.register records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentRegisters(criteria *Criteria, options *Options) (*AccountPaymentRegisters, error) {
	aprs := &AccountPaymentRegisters{}
	if err := c.SearchRead(AccountPaymentRegisterModel, criteria, options, aprs); err != nil {
		return nil, err
	}
	return aprs, nil
}

// FindAccountPaymentRegisterIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountPaymentRegisterIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountPaymentRegisterModel, criteria, options)
}

// FindAccountPaymentRegisterId finds record id by querying it with criteria.
func (c *Client) FindAccountPaymentRegisterId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountPaymentRegisterModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
