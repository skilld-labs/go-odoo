package odoo

import (
	"fmt"
)

// ResUsers represents res.users model.
type ResUsers struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omitempty"`
	AccessesCount                 *Int       `xmlrpc:"accesses_count,omitempty"`
	ActionId                      *Many2One  `xmlrpc:"action_id,omitempty"`
	Active                        *Bool      `xmlrpc:"active,omitempty"`
	ActiveLangCount               *Int       `xmlrpc:"active_lang_count,omitempty"`
	ActivePartner                 *Bool      `xmlrpc:"active_partner,omitempty"`
	ActivityDateDeadline          *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration   *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon         *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                   *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                 *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon              *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AdditionalInfo                *String    `xmlrpc:"additional_info,omitempty"`
	ApiKeyIds                     *Relation  `xmlrpc:"api_key_ids,omitempty"`
	BankAccountCount              *Int       `xmlrpc:"bank_account_count,omitempty"`
	BankIds                       *Relation  `xmlrpc:"bank_ids,omitempty"`
	Barcode                       *String    `xmlrpc:"barcode,omitempty"`
	CategoryId                    *Relation  `xmlrpc:"category_id,omitempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omitempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omitempty"`
	City                          *String    `xmlrpc:"city,omitempty"`
	Color                         *Int       `xmlrpc:"color,omitempty"`
	Comment                       *String    `xmlrpc:"comment,omitempty"`
	CommercialCompanyName         *String    `xmlrpc:"commercial_company_name,omitempty"`
	CommercialPartnerId           *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompaniesCount                *Int       `xmlrpc:"companies_count,omitempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyIds                    *Relation  `xmlrpc:"company_ids,omitempty"`
	CompanyName                   *String    `xmlrpc:"company_name,omitempty"`
	CompanyType                   *Selection `xmlrpc:"company_type,omitempty"`
	ContactAddress                *String    `xmlrpc:"contact_address,omitempty"`
	ContractIds                   *Relation  `xmlrpc:"contract_ids,omitempty"`
	CountryId                     *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omitempty"`
	Credit                        *Float     `xmlrpc:"credit,omitempty"`
	CreditLimit                   *Float     `xmlrpc:"credit_limit,omitempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omitempty"`
	CustomerRank                  *Int       `xmlrpc:"customer_rank,omitempty"`
	Date                          *Time      `xmlrpc:"date,omitempty"`
	Debit                         *Float     `xmlrpc:"debit,omitempty"`
	DebitLimit                    *Float     `xmlrpc:"debit_limit,omitempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omitempty"`
	Email                         *String    `xmlrpc:"email,omitempty"`
	EmailFormatted                *String    `xmlrpc:"email_formatted,omitempty"`
	EmailNormalized               *String    `xmlrpc:"email_normalized,omitempty"`
	Employee                      *Bool      `xmlrpc:"employee,omitempty"`
	Function                      *String    `xmlrpc:"function,omitempty"`
	GroupsCount                   *Int       `xmlrpc:"groups_count,omitempty"`
	GroupsId                      *Relation  `xmlrpc:"groups_id,omitempty"`
	HasUnreconciledEntries        *Bool      `xmlrpc:"has_unreconciled_entries,omitempty"`
	Id                            *Int       `xmlrpc:"id,omitempty"`
	ImStatus                      *String    `xmlrpc:"im_status,omitempty"`
	Image1024                     *String    `xmlrpc:"image_1024,omitempty"`
	Image128                      *String    `xmlrpc:"image_128,omitempty"`
	Image1920                     *String    `xmlrpc:"image_1920,omitempty"`
	Image256                      *String    `xmlrpc:"image_256,omitempty"`
	Image512                      *String    `xmlrpc:"image_512,omitempty"`
	IndustryId                    *Many2One  `xmlrpc:"industry_id,omitempty"`
	InvoiceIds                    *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceWarn                   *Selection `xmlrpc:"invoice_warn,omitempty"`
	InvoiceWarnMsg                *String    `xmlrpc:"invoice_warn_msg,omitempty"`
	IsBlacklisted                 *Bool      `xmlrpc:"is_blacklisted,omitempty"`
	IsCompany                     *Bool      `xmlrpc:"is_company,omitempty"`
	IsModerator                   *Bool      `xmlrpc:"is_moderator,omitempty"`
	JournalItemCount              *Int       `xmlrpc:"journal_item_count,omitempty"`
	Lang                          *Selection `xmlrpc:"lang,omitempty"`
	LastTimeEntriesChecked        *Time      `xmlrpc:"last_time_entries_checked,omitempty"`
	LogIds                        *Relation  `xmlrpc:"log_ids,omitempty"`
	Login                         *String    `xmlrpc:"login,omitempty"`
	LoginDate                     *Time      `xmlrpc:"login_date,omitempty"`
	MessageAttachmentCount        *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageBounce                 *Int       `xmlrpc:"message_bounce,omitempty"`
	MessageChannelIds             *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError               *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter        *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError            *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId       *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                 *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter          *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Mobile                        *String    `xmlrpc:"mobile,omitempty"`
	MobileBlacklisted             *Bool      `xmlrpc:"mobile_blacklisted,omitempty"`
	ModerationChannelIds          *Relation  `xmlrpc:"moderation_channel_ids,omitempty"`
	ModerationCounter             *Int       `xmlrpc:"moderation_counter,omitempty"`
	MyActivityDateDeadline        *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty"`
	NewPassword                   *String    `xmlrpc:"new_password,omitempty"`
	NotificationType              *Selection `xmlrpc:"notification_type,omitempty"`
	OdoobotFailed                 *Bool      `xmlrpc:"odoobot_failed,omitempty"`
	OdoobotState                  *Selection `xmlrpc:"odoobot_state,omitempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omitempty"`
	ParentName                    *String    `xmlrpc:"parent_name,omitempty"`
	PartnerGid                    *Int       `xmlrpc:"partner_gid,omitempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerLatitude               *Float     `xmlrpc:"partner_latitude,omitempty"`
	PartnerLongitude              *Float     `xmlrpc:"partner_longitude,omitempty"`
	PartnerShare                  *Bool      `xmlrpc:"partner_share,omitempty"`
	Password                      *String    `xmlrpc:"password,omitempty"`
	PaymentTokenCount             *Int       `xmlrpc:"payment_token_count,omitempty"`
	PaymentTokenIds               *Relation  `xmlrpc:"payment_token_ids,omitempty"`
	Phone                         *String    `xmlrpc:"phone,omitempty"`
	PhoneBlacklisted              *Bool      `xmlrpc:"phone_blacklisted,omitempty"`
	PhoneSanitized                *String    `xmlrpc:"phone_sanitized,omitempty"`
	PhoneSanitizedBlacklisted     *Bool      `xmlrpc:"phone_sanitized_blacklisted,omitempty"`
	PropertyAccountPayableId      *Many2One  `xmlrpc:"property_account_payable_id,omitempty"`
	PropertyAccountPositionId     *Many2One  `xmlrpc:"property_account_position_id,omitempty"`
	PropertyAccountReceivableId   *Many2One  `xmlrpc:"property_account_receivable_id,omitempty"`
	PropertyPaymentMethodId       *Many2One  `xmlrpc:"property_payment_method_id,omitempty"`
	PropertyPaymentTermId         *Many2One  `xmlrpc:"property_payment_term_id,omitempty"`
	PropertyProductPricelist      *Many2One  `xmlrpc:"property_product_pricelist,omitempty"`
	PropertySupplierPaymentTermId *Many2One  `xmlrpc:"property_supplier_payment_term_id,omitempty"`
	Ref                           *String    `xmlrpc:"ref,omitempty"`
	RefCompanyIds                 *Relation  `xmlrpc:"ref_company_ids,omitempty"`
	ResourceCalendarId            *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceIds                   *Relation  `xmlrpc:"resource_ids,omitempty"`
	RulesCount                    *Int       `xmlrpc:"rules_count,omitempty"`
	SaleOrderCount                *Int       `xmlrpc:"sale_order_count,omitempty"`
	SaleOrderIds                  *Relation  `xmlrpc:"sale_order_ids,omitempty"`
	SaleTeamId                    *Many2One  `xmlrpc:"sale_team_id,omitempty"`
	SaleWarn                      *Selection `xmlrpc:"sale_warn,omitempty"`
	SaleWarnMsg                   *String    `xmlrpc:"sale_warn_msg,omitempty"`
	SameVatPartnerId              *Many2One  `xmlrpc:"same_vat_partner_id,omitempty"`
	Self                          *Many2One  `xmlrpc:"self,omitempty"`
	Share                         *Bool      `xmlrpc:"share,omitempty"`
	Signature                     *String    `xmlrpc:"signature,omitempty"`
	SignupExpiration              *Time      `xmlrpc:"signup_expiration,omitempty"`
	SignupToken                   *String    `xmlrpc:"signup_token,omitempty"`
	SignupType                    *String    `xmlrpc:"signup_type,omitempty"`
	SignupUrl                     *String    `xmlrpc:"signup_url,omitempty"`
	SignupValid                   *Bool      `xmlrpc:"signup_valid,omitempty"`
	State                         *Selection `xmlrpc:"state,omitempty"`
	StateId                       *Many2One  `xmlrpc:"state_id,omitempty"`
	Street                        *String    `xmlrpc:"street,omitempty"`
	Street2                       *String    `xmlrpc:"street2,omitempty"`
	SupplierRank                  *Int       `xmlrpc:"supplier_rank,omitempty"`
	TeamId                        *Many2One  `xmlrpc:"team_id,omitempty"`
	Title                         *Many2One  `xmlrpc:"title,omitempty"`
	TotalInvoiced                 *Float     `xmlrpc:"total_invoiced,omitempty"`
	TotpEnabled                   *Bool      `xmlrpc:"totp_enabled,omitempty"`
	TotpSecret                    *String    `xmlrpc:"totp_secret,omitempty"`
	Trust                         *Selection `xmlrpc:"trust,omitempty"`
	Type                          *Selection `xmlrpc:"type,omitempty"`
	Tz                            *Selection `xmlrpc:"tz,omitempty"`
	TzOffset                      *String    `xmlrpc:"tz_offset,omitempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omitempty"`
	UserIds                       *Relation  `xmlrpc:"user_ids,omitempty"`
	Vat                           *String    `xmlrpc:"vat,omitempty"`
	Website                       *String    `xmlrpc:"website,omitempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omitempty"`
	Zip                           *String    `xmlrpc:"zip,omitempty"`
}

// ResUserss represents array of res.users model.
type ResUserss []ResUsers

// ResUsersModel is the odoo model name.
const ResUsersModel = "res.users"

// Many2One convert ResUsers to *Many2One.
func (ru *ResUsers) Many2One() *Many2One {
	return NewMany2One(ru.Id.Get(), "")
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUsers(ru *ResUsers) (int64, error) {
	return c.Create(ResUsersModel, ru)
}

// UpdateResUsers updates an existing res.users record.
func (c *Client) UpdateResUsers(ru *ResUsers) error {
	return c.UpdateResUserss([]int64{ru.Id.Get()}, ru)
}

// UpdateResUserss updates existing res.users records.
// All records (represented by ids) will be updated by ru values.
func (c *Client) UpdateResUserss(ids []int64, ru *ResUsers) error {
	return c.Update(ResUsersModel, ids, ru)
}

// DeleteResUsers deletes an existing res.users record.
func (c *Client) DeleteResUsers(id int64) error {
	return c.DeleteResUserss([]int64{id})
}

// DeleteResUserss deletes existing res.users records.
func (c *Client) DeleteResUserss(ids []int64) error {
	return c.Delete(ResUsersModel, ids)
}

// GetResUsers gets res.users existing record.
func (c *Client) GetResUsers(id int64) (*ResUsers, error) {
	rus, err := c.GetResUserss([]int64{id})
	if err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.users not found", id)
}

// GetResUserss gets res.users existing records.
func (c *Client) GetResUserss(ids []int64) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.Read(ResUsersModel, ids, nil, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsers finds res.users record by querying it with criteria.
func (c *Client) FindResUsers(criteria *Criteria) (*ResUsers, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, NewOptions().Limit(1), rus); err != nil {
		return nil, err
	}
	if rus != nil && len(*rus) > 0 {
		return &((*rus)[0]), nil
	}
	return nil, fmt.Errorf("res.users was not found")
}

// FindResUserss finds res.users records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUserss(criteria *Criteria, options *Options) (*ResUserss, error) {
	rus := &ResUserss{}
	if err := c.SearchRead(ResUsersModel, criteria, options, rus); err != nil {
		return nil, err
	}
	return rus, nil
}

// FindResUsersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResUsersId finds record id by querying it with criteria.
func (c *Client) FindResUsersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.users was not found")
}
