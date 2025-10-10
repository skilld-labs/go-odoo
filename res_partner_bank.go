package odoo

// ResPartnerBank represents res.partner.bank model.
type ResPartnerBank struct {
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
	PartnerCountryName              *String    `xmlrpc:"partner_country_name,omitempty"`
	PartnerCustomerRank             *Int       `xmlrpc:"partner_customer_rank,omitempty"`
	PartnerId                       *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerSupplierRank             *Int       `xmlrpc:"partner_supplier_rank,omitempty"`
	RatingIds                       *Relation  `xmlrpc:"rating_ids,omitempty"`
	RelatedMoves                    *Relation  `xmlrpc:"related_moves,omitempty"`
	SanitizedAccNumber              *String    `xmlrpc:"sanitized_acc_number,omitempty"`
	Sequence                        *Int       `xmlrpc:"sequence,omitempty"`
	UserHasGroupValidateBankAccount *Bool      `xmlrpc:"user_has_group_validate_bank_account,omitempty"`
	WebsiteMessageIds               *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                       *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                        *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ResPartnerBanks represents array of res.partner.bank model.
type ResPartnerBanks []ResPartnerBank

// ResPartnerBankModel is the odoo model name.
const ResPartnerBankModel = "res.partner.bank"

// Many2One convert ResPartnerBank to *Many2One.
func (rpb *ResPartnerBank) Many2One() *Many2One {
	return NewMany2One(rpb.Id.Get(), "")
}

// CreateResPartnerBank creates a new res.partner.bank model and returns its id.
func (c *Client) CreateResPartnerBank(rpb *ResPartnerBank) (int64, error) {
	ids, err := c.CreateResPartnerBanks([]*ResPartnerBank{rpb})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResPartnerBank creates a new res.partner.bank model and returns its id.
func (c *Client) CreateResPartnerBanks(rpbs []*ResPartnerBank) ([]int64, error) {
	var vv []interface{}
	for _, v := range rpbs {
		vv = append(vv, v)
	}
	return c.Create(ResPartnerBankModel, vv, nil)
}

// UpdateResPartnerBank updates an existing res.partner.bank record.
func (c *Client) UpdateResPartnerBank(rpb *ResPartnerBank) error {
	return c.UpdateResPartnerBanks([]int64{rpb.Id.Get()}, rpb)
}

// UpdateResPartnerBanks updates existing res.partner.bank records.
// All records (represented by ids) will be updated by rpb values.
func (c *Client) UpdateResPartnerBanks(ids []int64, rpb *ResPartnerBank) error {
	return c.Update(ResPartnerBankModel, ids, rpb, nil)
}

// DeleteResPartnerBank deletes an existing res.partner.bank record.
func (c *Client) DeleteResPartnerBank(id int64) error {
	return c.DeleteResPartnerBanks([]int64{id})
}

// DeleteResPartnerBanks deletes existing res.partner.bank records.
func (c *Client) DeleteResPartnerBanks(ids []int64) error {
	return c.Delete(ResPartnerBankModel, ids)
}

// GetResPartnerBank gets res.partner.bank existing record.
func (c *Client) GetResPartnerBank(id int64) (*ResPartnerBank, error) {
	rpbs, err := c.GetResPartnerBanks([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rpbs)[0]), nil
}

// GetResPartnerBanks gets res.partner.bank existing records.
func (c *Client) GetResPartnerBanks(ids []int64) (*ResPartnerBanks, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.Read(ResPartnerBankModel, ids, nil, rpbs); err != nil {
		return nil, err
	}
	return rpbs, nil
}

// FindResPartnerBank finds res.partner.bank record by querying it with criteria.
func (c *Client) FindResPartnerBank(criteria *Criteria) (*ResPartnerBank, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.SearchRead(ResPartnerBankModel, criteria, NewOptions().Limit(1), rpbs); err != nil {
		return nil, err
	}
	return &((*rpbs)[0]), nil
}

// FindResPartnerBanks finds res.partner.bank records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerBanks(criteria *Criteria, options *Options) (*ResPartnerBanks, error) {
	rpbs := &ResPartnerBanks{}
	if err := c.SearchRead(ResPartnerBankModel, criteria, options, rpbs); err != nil {
		return nil, err
	}
	return rpbs, nil
}

// FindResPartnerBankIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerBankIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResPartnerBankModel, criteria, options)
}

// FindResPartnerBankId finds record id by querying it with criteria.
func (c *Client) FindResPartnerBankId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerBankModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
