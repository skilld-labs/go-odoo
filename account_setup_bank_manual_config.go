package odoo

// AccountSetupBankManualConfig represents account.setup.bank.manual.config model.
type AccountSetupBankManualConfig struct {
	AccHolderName                   *String    `xmlrpc:"acc_holder_name,omitempty"`
	AccNumber                       *String    `xmlrpc:"acc_number,omitempty"`
	AccType                         *Selection `xmlrpc:"acc_type,omitempty"`
	Active                          *Bool      `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId         *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline            *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration     *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon           *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                     *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                   *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                 *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                  *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                  *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AllowOutPayment                 *Bool      `xmlrpc:"allow_out_payment,omitempty"`
	BankBic                         *String    `xmlrpc:"bank_bic,omitempty"`
	BankId                          *Many2One  `xmlrpc:"bank_id,omitempty"`
	BankName                        *String    `xmlrpc:"bank_name,omitempty"`
	CompanyId                       *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryCode                     *String    `xmlrpc:"country_code,omitempty"`
	CreateDate                      *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                       *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                      *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName                     *String    `xmlrpc:"display_name,omitempty"`
	HasIbanWarning                  *Bool      `xmlrpc:"has_iban_warning,omitempty"`
	HasMessage                      *Bool      `xmlrpc:"has_message,omitempty"`
	HasMoneyTransferWarning         *Bool      `xmlrpc:"has_money_transfer_warning,omitempty"`
	Id                              *Int       `xmlrpc:"id,omitempty"`
	JournalId                       *Relation  `xmlrpc:"journal_id,omitempty"`
	LinkedJournalId                 *Many2One  `xmlrpc:"linked_journal_id,omitempty"`
	LockTrustFields                 *Bool      `xmlrpc:"lock_trust_fields,omitempty"`
	MessageAttachmentCount          *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds              *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                 *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter          *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError              *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                      *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower               *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction               *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter        *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds               *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MoneyTransferService            *String    `xmlrpc:"money_transfer_service,omitempty"`
	MyActivityDateDeadline          *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	NewJournalName                  *String    `xmlrpc:"new_journal_name,omitempty"`
	NumJournalsWithoutAccount       *Int       `xmlrpc:"num_journals_without_account,omitempty"`
	PartnerCountryName              *String    `xmlrpc:"partner_country_name,omitempty"`
	PartnerCustomerRank             *Int       `xmlrpc:"partner_customer_rank,omitempty"`
	PartnerId                       *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerSupplierRank             *Int       `xmlrpc:"partner_supplier_rank,omitempty"`
	RatingIds                       *Relation  `xmlrpc:"rating_ids,omitempty"`
	RelatedMoves                    *Relation  `xmlrpc:"related_moves,omitempty"`
	ResPartnerBankId                *Many2One  `xmlrpc:"res_partner_bank_id,omitempty"`
	SanitizedAccNumber              *String    `xmlrpc:"sanitized_acc_number,omitempty"`
	Sequence                        *Int       `xmlrpc:"sequence,omitempty"`
	UserHasGroupValidateBankAccount *Bool      `xmlrpc:"user_has_group_validate_bank_account,omitempty"`
	WebsiteMessageIds               *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountSetupBankManualConfigs represents array of account.setup.bank.manual.config model.
type AccountSetupBankManualConfigs []AccountSetupBankManualConfig

// AccountSetupBankManualConfigModel is the odoo model name.
const AccountSetupBankManualConfigModel = "account.setup.bank.manual.config"

// Many2One convert AccountSetupBankManualConfig to *Many2One.
func (asbmc *AccountSetupBankManualConfig) Many2One() *Many2One {
	return NewMany2One(asbmc.Id.Get(), "")
}

// CreateAccountSetupBankManualConfig creates a new account.setup.bank.manual.config model and returns its id.
func (c *Client) CreateAccountSetupBankManualConfig(asbmc *AccountSetupBankManualConfig) (int64, error) {
	ids, err := c.CreateAccountSetupBankManualConfigs([]*AccountSetupBankManualConfig{asbmc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountSetupBankManualConfig creates a new account.setup.bank.manual.config model and returns its id.
func (c *Client) CreateAccountSetupBankManualConfigs(asbmcs []*AccountSetupBankManualConfig) ([]int64, error) {
	var vv []interface{}
	for _, v := range asbmcs {
		vv = append(vv, v)
	}
	return c.Create(AccountSetupBankManualConfigModel, vv, nil)
}

// UpdateAccountSetupBankManualConfig updates an existing account.setup.bank.manual.config record.
func (c *Client) UpdateAccountSetupBankManualConfig(asbmc *AccountSetupBankManualConfig) error {
	return c.UpdateAccountSetupBankManualConfigs([]int64{asbmc.Id.Get()}, asbmc)
}

// UpdateAccountSetupBankManualConfigs updates existing account.setup.bank.manual.config records.
// All records (represented by ids) will be updated by asbmc values.
func (c *Client) UpdateAccountSetupBankManualConfigs(ids []int64, asbmc *AccountSetupBankManualConfig) error {
	return c.Update(AccountSetupBankManualConfigModel, ids, asbmc, nil)
}

// DeleteAccountSetupBankManualConfig deletes an existing account.setup.bank.manual.config record.
func (c *Client) DeleteAccountSetupBankManualConfig(id int64) error {
	return c.DeleteAccountSetupBankManualConfigs([]int64{id})
}

// DeleteAccountSetupBankManualConfigs deletes existing account.setup.bank.manual.config records.
func (c *Client) DeleteAccountSetupBankManualConfigs(ids []int64) error {
	return c.Delete(AccountSetupBankManualConfigModel, ids)
}

// GetAccountSetupBankManualConfig gets account.setup.bank.manual.config existing record.
func (c *Client) GetAccountSetupBankManualConfig(id int64) (*AccountSetupBankManualConfig, error) {
	asbmcs, err := c.GetAccountSetupBankManualConfigs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*asbmcs)[0]), nil
}

// GetAccountSetupBankManualConfigs gets account.setup.bank.manual.config existing records.
func (c *Client) GetAccountSetupBankManualConfigs(ids []int64) (*AccountSetupBankManualConfigs, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.Read(AccountSetupBankManualConfigModel, ids, nil, asbmcs); err != nil {
		return nil, err
	}
	return asbmcs, nil
}

// FindAccountSetupBankManualConfig finds account.setup.bank.manual.config record by querying it with criteria.
func (c *Client) FindAccountSetupBankManualConfig(criteria *Criteria) (*AccountSetupBankManualConfig, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.SearchRead(AccountSetupBankManualConfigModel, criteria, NewOptions().Limit(1), asbmcs); err != nil {
		return nil, err
	}
	return &((*asbmcs)[0]), nil
}

// FindAccountSetupBankManualConfigs finds account.setup.bank.manual.config records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSetupBankManualConfigs(criteria *Criteria, options *Options) (*AccountSetupBankManualConfigs, error) {
	asbmcs := &AccountSetupBankManualConfigs{}
	if err := c.SearchRead(AccountSetupBankManualConfigModel, criteria, options, asbmcs); err != nil {
		return nil, err
	}
	return asbmcs, nil
}

// FindAccountSetupBankManualConfigIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountSetupBankManualConfigIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountSetupBankManualConfigModel, criteria, options)
}

// FindAccountSetupBankManualConfigId finds record id by querying it with criteria.
func (c *Client) FindAccountSetupBankManualConfigId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountSetupBankManualConfigModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
