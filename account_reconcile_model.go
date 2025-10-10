package odoo

// AccountReconcileModel represents account.reconcile.model model.
type AccountReconcileModel struct {
	Active                     *Bool      `xmlrpc:"active,omitempty"`
	AllowPaymentTolerance      *Bool      `xmlrpc:"allow_payment_tolerance,omitempty"`
	AutoReconcile              *Bool      `xmlrpc:"auto_reconcile,omitempty"`
	CompanyId                  *Many2One  `xmlrpc:"company_id,omitempty"`
	CounterpartType            *Selection `xmlrpc:"counterpart_type,omitempty"`
	CreateDate                 *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                  *Many2One  `xmlrpc:"create_uid,omitempty"`
	DecimalSeparator           *String    `xmlrpc:"decimal_separator,omitempty"`
	DisplayName                *String    `xmlrpc:"display_name,omitempty"`
	HasMessage                 *Bool      `xmlrpc:"has_message,omitempty"`
	Id                         *Int       `xmlrpc:"id,omitempty"`
	LineIds                    *Relation  `xmlrpc:"line_ids,omitempty"`
	MatchAmount                *Selection `xmlrpc:"match_amount,omitempty"`
	MatchAmountMax             *Float     `xmlrpc:"match_amount_max,omitempty"`
	MatchAmountMin             *Float     `xmlrpc:"match_amount_min,omitempty"`
	MatchJournalIds            *Relation  `xmlrpc:"match_journal_ids,omitempty"`
	MatchLabel                 *Selection `xmlrpc:"match_label,omitempty"`
	MatchLabelParam            *String    `xmlrpc:"match_label_param,omitempty"`
	MatchNature                *Selection `xmlrpc:"match_nature,omitempty"`
	MatchNote                  *Selection `xmlrpc:"match_note,omitempty"`
	MatchNoteParam             *String    `xmlrpc:"match_note_param,omitempty"`
	MatchPartner               *Bool      `xmlrpc:"match_partner,omitempty"`
	MatchPartnerCategoryIds    *Relation  `xmlrpc:"match_partner_category_ids,omitempty"`
	MatchPartnerIds            *Relation  `xmlrpc:"match_partner_ids,omitempty"`
	MatchSameCurrency          *Bool      `xmlrpc:"match_same_currency,omitempty"`
	MatchTextLocationLabel     *Bool      `xmlrpc:"match_text_location_label,omitempty"`
	MatchTextLocationNote      *Bool      `xmlrpc:"match_text_location_note,omitempty"`
	MatchTextLocationReference *Bool      `xmlrpc:"match_text_location_reference,omitempty"`
	MatchTransactionType       *Selection `xmlrpc:"match_transaction_type,omitempty"`
	MatchTransactionTypeParam  *String    `xmlrpc:"match_transaction_type_param,omitempty"`
	MatchingOrder              *Selection `xmlrpc:"matching_order,omitempty"`
	MessageAttachmentCount     *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds         *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError            *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter     *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError         *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                 *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower          *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction          *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter   *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds          *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Name                       *String    `xmlrpc:"name,omitempty"`
	NumberEntries              *Int       `xmlrpc:"number_entries,omitempty"`
	PartnerMappingLineIds      *Relation  `xmlrpc:"partner_mapping_line_ids,omitempty"`
	PastMonthsLimit            *Int       `xmlrpc:"past_months_limit,omitempty"`
	PaymentToleranceParam      *Float     `xmlrpc:"payment_tolerance_param,omitempty"`
	PaymentToleranceType       *Selection `xmlrpc:"payment_tolerance_type,omitempty"`
	RatingIds                  *Relation  `xmlrpc:"rating_ids,omitempty"`
	RuleType                   *Selection `xmlrpc:"rule_type,omitempty"`
	Sequence                   *Int       `xmlrpc:"sequence,omitempty"`
	ShowDecimalSeparator       *Bool      `xmlrpc:"show_decimal_separator,omitempty"`
	ToCheck                    *Bool      `xmlrpc:"to_check,omitempty"`
	WebsiteMessageIds          *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                  *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                   *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AccountReconcileModels represents array of account.reconcile.model model.
type AccountReconcileModels []AccountReconcileModel

// AccountReconcileModelModel is the odoo model name.
const AccountReconcileModelModel = "account.reconcile.model"

// Many2One convert AccountReconcileModel to *Many2One.
func (arm *AccountReconcileModel) Many2One() *Many2One {
	return NewMany2One(arm.Id.Get(), "")
}

// CreateAccountReconcileModel creates a new account.reconcile.model model and returns its id.
func (c *Client) CreateAccountReconcileModel(arm *AccountReconcileModel) (int64, error) {
	ids, err := c.CreateAccountReconcileModels([]*AccountReconcileModel{arm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountReconcileModel creates a new account.reconcile.model model and returns its id.
func (c *Client) CreateAccountReconcileModels(arms []*AccountReconcileModel) ([]int64, error) {
	var vv []interface{}
	for _, v := range arms {
		vv = append(vv, v)
	}
	return c.Create(AccountReconcileModelModel, vv, nil)
}

// UpdateAccountReconcileModel updates an existing account.reconcile.model record.
func (c *Client) UpdateAccountReconcileModel(arm *AccountReconcileModel) error {
	return c.UpdateAccountReconcileModels([]int64{arm.Id.Get()}, arm)
}

// UpdateAccountReconcileModels updates existing account.reconcile.model records.
// All records (represented by ids) will be updated by arm values.
func (c *Client) UpdateAccountReconcileModels(ids []int64, arm *AccountReconcileModel) error {
	return c.Update(AccountReconcileModelModel, ids, arm, nil)
}

// DeleteAccountReconcileModel deletes an existing account.reconcile.model record.
func (c *Client) DeleteAccountReconcileModel(id int64) error {
	return c.DeleteAccountReconcileModels([]int64{id})
}

// DeleteAccountReconcileModels deletes existing account.reconcile.model records.
func (c *Client) DeleteAccountReconcileModels(ids []int64) error {
	return c.Delete(AccountReconcileModelModel, ids)
}

// GetAccountReconcileModel gets account.reconcile.model existing record.
func (c *Client) GetAccountReconcileModel(id int64) (*AccountReconcileModel, error) {
	arms, err := c.GetAccountReconcileModels([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*arms)[0]), nil
}

// GetAccountReconcileModels gets account.reconcile.model existing records.
func (c *Client) GetAccountReconcileModels(ids []int64) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.Read(AccountReconcileModelModel, ids, nil, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModel finds account.reconcile.model record by querying it with criteria.
func (c *Client) FindAccountReconcileModel(criteria *Criteria) (*AccountReconcileModel, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, NewOptions().Limit(1), arms); err != nil {
		return nil, err
	}
	return &((*arms)[0]), nil
}

// FindAccountReconcileModels finds account.reconcile.model records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModels(criteria *Criteria, options *Options) (*AccountReconcileModels, error) {
	arms := &AccountReconcileModels{}
	if err := c.SearchRead(AccountReconcileModelModel, criteria, options, arms); err != nil {
		return nil, err
	}
	return arms, nil
}

// FindAccountReconcileModelIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountReconcileModelIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountReconcileModelModel, criteria, options)
}

// FindAccountReconcileModelId finds record id by querying it with criteria.
func (c *Client) FindAccountReconcileModelId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountReconcileModelModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
