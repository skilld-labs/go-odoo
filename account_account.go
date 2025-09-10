package odoo

// AccountAccount represents account.account model.
type AccountAccount struct {
	AccountType                 *Selection `xmlrpc:"account_type,omitempty"`
	AllowedJournalIds           *Relation  `xmlrpc:"allowed_journal_ids,omitempty"`
	AssetModelIds               *Relation  `xmlrpc:"asset_model_ids,omitempty"`
	BudgetItemIds               *Relation  `xmlrpc:"budget_item_ids,omitempty"`
	CanCreateAsset              *Bool      `xmlrpc:"can_create_asset,omitempty"`
	Code                        *String    `xmlrpc:"code,omitempty"`
	CodeMappingIds              *Relation  `xmlrpc:"code_mapping_ids,omitempty"`
	CodeStore                   *String    `xmlrpc:"code_store,omitempty"`
	CompanyCurrencyId           *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyFiscalCountryCode    *String    `xmlrpc:"company_fiscal_country_code,omitempty"`
	CompanyIds                  *Relation  `xmlrpc:"company_ids,omitempty"`
	CreateAsset                 *Selection `xmlrpc:"create_asset,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	CurrentBalance              *Float     `xmlrpc:"current_balance,omitempty"`
	Deprecated                  *Bool      `xmlrpc:"deprecated,omitempty"`
	DisplayMappingTab           *Bool      `xmlrpc:"display_mapping_tab,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	ExcludeProvisionCurrencyIds *Relation  `xmlrpc:"exclude_provision_currency_ids,omitempty"`
	FormViewRef                 *String    `xmlrpc:"form_view_ref,omitempty"`
	GroupId                     *Many2One  `xmlrpc:"group_id,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IncludeInitialBalance       *Bool      `xmlrpc:"include_initial_balance,omitempty"`
	InternalGroup               *Selection `xmlrpc:"internal_group,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MultipleAssetsPerLine       *Bool      `xmlrpc:"multiple_assets_per_line,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	NonTrade                    *Bool      `xmlrpc:"non_trade,omitempty"`
	Note                        *String    `xmlrpc:"note,omitempty"`
	OpeningBalance              *Float     `xmlrpc:"opening_balance,omitempty"`
	OpeningCredit               *Float     `xmlrpc:"opening_credit,omitempty"`
	OpeningDebit                *Float     `xmlrpc:"opening_debit,omitempty"`
	PlaceholderCode             *String    `xmlrpc:"placeholder_code,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	Reconcile                   *Bool      `xmlrpc:"reconcile,omitempty"`
	RelatedTaxesAmount          *Int       `xmlrpc:"related_taxes_amount,omitempty"`
	RootId                      *Many2One  `xmlrpc:"root_id,omitempty"`
	TagIds                      *Relation  `xmlrpc:"tag_ids,omitempty"`
	TaxIds                      *Relation  `xmlrpc:"tax_ids,omitempty"`
	Used                        *Bool      `xmlrpc:"used,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountAccounts represents array of account.account model.
type AccountAccounts []AccountAccount

// AccountAccountModel is the odoo model name.
const AccountAccountModel = "account.account"

// Many2One convert AccountAccount to *Many2One.
func (aa *AccountAccount) Many2One() *Many2One {
	return NewMany2One(aa.Id.Get(), "")
}

// CreateAccountAccount creates a new account.account model and returns its id.
func (c *Client) CreateAccountAccount(aa *AccountAccount) (int64, error) {
	ids, err := c.CreateAccountAccounts([]*AccountAccount{aa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountAccount creates a new account.account model and returns its id.
func (c *Client) CreateAccountAccounts(aas []*AccountAccount) ([]int64, error) {
	var vv []interface{}
	for _, v := range aas {
		vv = append(vv, v)
	}
	return c.Create(AccountAccountModel, vv, nil)
}

// UpdateAccountAccount updates an existing account.account record.
func (c *Client) UpdateAccountAccount(aa *AccountAccount) error {
	return c.UpdateAccountAccounts([]int64{aa.Id.Get()}, aa)
}

// UpdateAccountAccounts updates existing account.account records.
// All records (represented by ids) will be updated by aa values.
func (c *Client) UpdateAccountAccounts(ids []int64, aa *AccountAccount) error {
	return c.Update(AccountAccountModel, ids, aa, nil)
}

// DeleteAccountAccount deletes an existing account.account record.
func (c *Client) DeleteAccountAccount(id int64) error {
	return c.DeleteAccountAccounts([]int64{id})
}

// DeleteAccountAccounts deletes existing account.account records.
func (c *Client) DeleteAccountAccounts(ids []int64) error {
	return c.Delete(AccountAccountModel, ids)
}

// GetAccountAccount gets account.account existing record.
func (c *Client) GetAccountAccount(id int64) (*AccountAccount, error) {
	aas, err := c.GetAccountAccounts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*aas)[0]), nil
}

// GetAccountAccounts gets account.account existing records.
func (c *Client) GetAccountAccounts(ids []int64) (*AccountAccounts, error) {
	aas := &AccountAccounts{}
	if err := c.Read(AccountAccountModel, ids, nil, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAccount finds account.account record by querying it with criteria.
func (c *Client) FindAccountAccount(criteria *Criteria) (*AccountAccount, error) {
	aas := &AccountAccounts{}
	if err := c.SearchRead(AccountAccountModel, criteria, NewOptions().Limit(1), aas); err != nil {
		return nil, err
	}
	return &((*aas)[0]), nil
}

// FindAccountAccounts finds account.account records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccounts(criteria *Criteria, options *Options) (*AccountAccounts, error) {
	aas := &AccountAccounts{}
	if err := c.SearchRead(AccountAccountModel, criteria, options, aas); err != nil {
		return nil, err
	}
	return aas, nil
}

// FindAccountAccountIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountAccountIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountAccountModel, criteria, options)
}

// FindAccountAccountId finds record id by querying it with criteria.
func (c *Client) FindAccountAccountId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountAccountModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
