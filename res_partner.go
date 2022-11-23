package odoo

import (
	"fmt"
)

// ResPartner represents res.partner model.
type ResPartner struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omitempty"`
	Active                        *Bool      `xmlrpc:"active,omitempty"`
	ActiveLangCount               *Int       `xmlrpc:"active_lang_count,omitempty"`
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
	CompanyId                     *Many2One  `xmlrpc:"company_id,omitempty"`
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
	JournalItemCount              *Int       `xmlrpc:"journal_item_count,omitempty"`
	Lang                          *Selection `xmlrpc:"lang,omitempty"`
	LastTimeEntriesChecked        *Time      `xmlrpc:"last_time_entries_checked,omitempty"`
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
	MyActivityDateDeadline        *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                          *String    `xmlrpc:"name,omitempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omitempty"`
	ParentName                    *String    `xmlrpc:"parent_name,omitempty"`
	PartnerGid                    *Int       `xmlrpc:"partner_gid,omitempty"`
	PartnerLatitude               *Float     `xmlrpc:"partner_latitude,omitempty"`
	PartnerLongitude              *Float     `xmlrpc:"partner_longitude,omitempty"`
	PartnerShare                  *Bool      `xmlrpc:"partner_share,omitempty"`
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
	SaleOrderCount                *Int       `xmlrpc:"sale_order_count,omitempty"`
	SaleOrderIds                  *Relation  `xmlrpc:"sale_order_ids,omitempty"`
	SaleWarn                      *Selection `xmlrpc:"sale_warn,omitempty"`
	SaleWarnMsg                   *String    `xmlrpc:"sale_warn_msg,omitempty"`
	SameVatPartnerId              *Many2One  `xmlrpc:"same_vat_partner_id,omitempty"`
	Self                          *Many2One  `xmlrpc:"self,omitempty"`
	SignupExpiration              *Time      `xmlrpc:"signup_expiration,omitempty"`
	SignupToken                   *String    `xmlrpc:"signup_token,omitempty"`
	SignupType                    *String    `xmlrpc:"signup_type,omitempty"`
	SignupUrl                     *String    `xmlrpc:"signup_url,omitempty"`
	SignupValid                   *Bool      `xmlrpc:"signup_valid,omitempty"`
	StateId                       *Many2One  `xmlrpc:"state_id,omitempty"`
	Street                        *String    `xmlrpc:"street,omitempty"`
	Street2                       *String    `xmlrpc:"street2,omitempty"`
	SupplierRank                  *Int       `xmlrpc:"supplier_rank,omitempty"`
	TeamId                        *Many2One  `xmlrpc:"team_id,omitempty"`
	Title                         *Many2One  `xmlrpc:"title,omitempty"`
	TotalInvoiced                 *Float     `xmlrpc:"total_invoiced,omitempty"`
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

// ResPartners represents array of res.partner model.
type ResPartners []ResPartner

// ResPartnerModel is the odoo model name.
const ResPartnerModel = "res.partner"

// Many2One convert ResPartner to *Many2One.
func (rp *ResPartner) Many2One() *Many2One {
	return NewMany2One(rp.Id.Get(), "")
}

// CreateResPartner creates a new res.partner model and returns its id.
func (c *Client) CreateResPartner(rp *ResPartner) (int64, error) {
	return c.Create(ResPartnerModel, rp)
}

// UpdateResPartner updates an existing res.partner record.
func (c *Client) UpdateResPartner(rp *ResPartner) error {
	return c.UpdateResPartners([]int64{rp.Id.Get()}, rp)
}

// UpdateResPartners updates existing res.partner records.
// All records (represented by IDs) will be updated by rp values.
func (c *Client) UpdateResPartners(ids []int64, rp *ResPartner) error {
	return c.Update(ResPartnerModel, ids, rp)
}

// DeleteResPartner deletes an existing res.partner record.
func (c *Client) DeleteResPartner(id int64) error {
	return c.DeleteResPartners([]int64{id})
}

// DeleteResPartners deletes existing res.partner records.
func (c *Client) DeleteResPartners(ids []int64) error {
	return c.Delete(ResPartnerModel, ids)
}

// GetResPartner gets res.partner existing record.
func (c *Client) GetResPartner(id int64) (*ResPartner, error) {
	rps, err := c.GetResPartners([]int64{id})
	if err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of res.partner not found", id)
}

// GetResPartners gets res.partner existing records.
func (c *Client) GetResPartners(ids []int64) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.Read(ResPartnerModel, ids, nil, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartner finds res.partner record by querying it with criteria.
func (c *Client) FindResPartner(criteria *Criteria) (*ResPartner, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, NewOptions().Limit(1), rps); err != nil {
		return nil, err
	}
	if rps != nil && len(*rps) > 0 {
		return &((*rps)[0]), nil
	}
	return nil, fmt.Errorf("res.partner was not found")
}

// FindResPartners finds res.partner records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartners(criteria *Criteria, options *Options) (*ResPartners, error) {
	rps := &ResPartners{}
	if err := c.SearchRead(ResPartnerModel, criteria, options, rps); err != nil {
		return nil, err
	}
	return rps, nil
}

// FindResPartnerIds finds records IDs by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindResPartnerId finds record id by querying it with criteria.
func (c *Client) FindResPartnerId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("res.partner was not found")
}
