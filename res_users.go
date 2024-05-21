package odoo

// ResUsers represents res.users model.
type ResUsers struct {
	AccessesCount                      *Int       `xmlrpc:"accesses_count,omitempty"`
	ActionId                           *Many2One  `xmlrpc:"action_id,omitempty"`
	Active                             *Bool      `xmlrpc:"active,omitempty"`
	ActiveLangCount                    *Int       `xmlrpc:"active_lang_count,omitempty"`
	ActivePartner                      *Bool      `xmlrpc:"active_partner,omitempty"`
	ActivityDateDeadline               *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration        *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon              *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                        *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                      *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                    *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                   *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                     *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                     *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AdditionalInfo                     *String    `xmlrpc:"additional_info,omitempty"`
	AdditionalNote                     *String    `xmlrpc:"additional_note,omitempty"`
	AddressId                          *Many2One  `xmlrpc:"address_id,omitempty"`
	ApiKeyIds                          *Relation  `xmlrpc:"api_key_ids,omitempty"`
	Avatar1024                         *String    `xmlrpc:"avatar_1024,omitempty"`
	Avatar128                          *String    `xmlrpc:"avatar_128,omitempty"`
	Avatar1920                         *String    `xmlrpc:"avatar_1920,omitempty"`
	Avatar256                          *String    `xmlrpc:"avatar_256,omitempty"`
	Avatar512                          *String    `xmlrpc:"avatar_512,omitempty"`
	BankAccountCount                   *Int       `xmlrpc:"bank_account_count,omitempty"`
	BankIds                            *Relation  `xmlrpc:"bank_ids,omitempty"`
	Barcode                            *String    `xmlrpc:"barcode,omitempty"`
	Birthday                           *Time      `xmlrpc:"birthday,omitempty"`
	BuyerId                            *Many2One  `xmlrpc:"buyer_id,omitempty"`
	CanEdit                            *Bool      `xmlrpc:"can_edit,omitempty"`
	CategoryId                         *Relation  `xmlrpc:"category_id,omitempty"`
	CategoryIds                        *Relation  `xmlrpc:"category_ids,omitempty"`
	Certificate                        *Selection `xmlrpc:"certificate,omitempty"`
	ChannelIds                         *Relation  `xmlrpc:"channel_ids,omitempty"`
	ChildIds                           *Relation  `xmlrpc:"child_ids,omitempty"`
	Children                           *Int       `xmlrpc:"children,omitempty"`
	City                               *String    `xmlrpc:"city,omitempty"`
	CoachId                            *Many2One  `xmlrpc:"coach_id,omitempty"`
	Color                              *Int       `xmlrpc:"color,omitempty"`
	Comment                            *String    `xmlrpc:"comment,omitempty"`
	CommercialCompanyName              *String    `xmlrpc:"commercial_company_name,omitempty"`
	CommercialPartnerId                *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompaniesCount                     *Int       `xmlrpc:"companies_count,omitempty"`
	CompanyId                          *Many2One  `xmlrpc:"company_id,omitempty"`
	CompanyIds                         *Relation  `xmlrpc:"company_ids,omitempty"`
	CompanyName                        *String    `xmlrpc:"company_name,omitempty"`
	CompanyRegistry                    *String    `xmlrpc:"company_registry,omitempty"`
	CompanyType                        *Selection `xmlrpc:"company_type,omitempty"`
	CompleteName                       *String    `xmlrpc:"complete_name,omitempty"`
	ContactAddress                     *String    `xmlrpc:"contact_address,omitempty"`
	ContactAddressComplete             *String    `xmlrpc:"contact_address_complete,omitempty"`
	ContactAddressInline               *String    `xmlrpc:"contact_address_inline,omitempty"`
	ContractIds                        *Relation  `xmlrpc:"contract_ids,omitempty"`
	CountryCode                        *String    `xmlrpc:"country_code,omitempty"`
	CountryId                          *Many2One  `xmlrpc:"country_id,omitempty"`
	CountryOfBirth                     *Many2One  `xmlrpc:"country_of_birth,omitempty"`
	CreateDate                         *Time      `xmlrpc:"create_date,omitempty"`
	CreateEmployee                     *Bool      `xmlrpc:"create_employee,omitempty"`
	CreateEmployeeId                   *Many2One  `xmlrpc:"create_employee_id,omitempty"`
	CreateUid                          *Many2One  `xmlrpc:"create_uid,omitempty"`
	Credit                             *Float     `xmlrpc:"credit,omitempty"`
	CreditLimit                        *Float     `xmlrpc:"credit_limit,omitempty"`
	CreditToInvoice                    *Float     `xmlrpc:"credit_to_invoice,omitempty"`
	CrmTeamIds                         *Relation  `xmlrpc:"crm_team_ids,omitempty"`
	CrmTeamMemberIds                   *Relation  `xmlrpc:"crm_team_member_ids,omitempty"`
	CurrencyId                         *Many2One  `xmlrpc:"currency_id,omitempty"`
	CustomerRank                       *Int       `xmlrpc:"customer_rank,omitempty"`
	Date                               *Time      `xmlrpc:"date,omitempty"`
	DaysSalesOutstanding               *Float     `xmlrpc:"days_sales_outstanding,omitempty"`
	Debit                              *Float     `xmlrpc:"debit,omitempty"`
	DebitLimit                         *Float     `xmlrpc:"debit_limit,omitempty"`
	DepartmentId                       *Many2One  `xmlrpc:"department_id,omitempty"`
	DisplayName                        *String    `xmlrpc:"display_name,omitempty"`
	DuplicatedBankAccountPartnersCount *Int       `xmlrpc:"duplicated_bank_account_partners_count,omitempty"`
	Email                              *String    `xmlrpc:"email,omitempty"`
	EmailFormatted                     *String    `xmlrpc:"email_formatted,omitempty"`
	EmailNormalized                    *String    `xmlrpc:"email_normalized,omitempty"`
	EmergencyContact                   *String    `xmlrpc:"emergency_contact,omitempty"`
	EmergencyPhone                     *String    `xmlrpc:"emergency_phone,omitempty"`
	Employee                           *Bool      `xmlrpc:"employee,omitempty"`
	EmployeeBankAccountId              *Many2One  `xmlrpc:"employee_bank_account_id,omitempty"`
	EmployeeCount                      *Int       `xmlrpc:"employee_count,omitempty"`
	EmployeeCountryId                  *Many2One  `xmlrpc:"employee_country_id,omitempty"`
	EmployeeId                         *Many2One  `xmlrpc:"employee_id,omitempty"`
	EmployeeIds                        *Relation  `xmlrpc:"employee_ids,omitempty"`
	EmployeeParentId                   *Many2One  `xmlrpc:"employee_parent_id,omitempty"`
	EmployeeResourceCalendarId         *Many2One  `xmlrpc:"employee_resource_calendar_id,omitempty"`
	EmployeeSkillIds                   *Relation  `xmlrpc:"employee_skill_ids,omitempty"`
	EmployeeType                       *Selection `xmlrpc:"employee_type,omitempty"`
	EmployeesCount                     *Int       `xmlrpc:"employees_count,omitempty"`
	FiscalCountryCodes                 *String    `xmlrpc:"fiscal_country_codes,omitempty"`
	Function                           *String    `xmlrpc:"function,omitempty"`
	Gender                             *Selection `xmlrpc:"gender,omitempty"`
	GroupsCount                        *Int       `xmlrpc:"groups_count,omitempty"`
	GroupsId                           *Relation  `xmlrpc:"groups_id,omitempty"`
	HasMessage                         *Bool      `xmlrpc:"has_message,omitempty"`
	HasUnreconciledEntries             *Bool      `xmlrpc:"has_unreconciled_entries,omitempty"`
	HrPresenceState                    *Selection `xmlrpc:"hr_presence_state,omitempty"`
	Id                                 *Int       `xmlrpc:"id,omitempty"`
	IdentificationId                   *String    `xmlrpc:"identification_id,omitempty"`
	ImStatus                           *String    `xmlrpc:"im_status,omitempty"`
	Image1024                          *String    `xmlrpc:"image_1024,omitempty"`
	Image128                           *String    `xmlrpc:"image_128,omitempty"`
	Image1920                          *String    `xmlrpc:"image_1920,omitempty"`
	Image256                           *String    `xmlrpc:"image_256,omitempty"`
	Image512                           *String    `xmlrpc:"image_512,omitempty"`
	ImageMedium                        *String    `xmlrpc:"image_medium,omitempty"`
	IndustryId                         *Many2One  `xmlrpc:"industry_id,omitempty"`
	InvoiceIds                         *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoiceWarn                        *Selection `xmlrpc:"invoice_warn,omitempty"`
	InvoiceWarnMsg                     *String    `xmlrpc:"invoice_warn_msg,omitempty"`
	IsBlacklisted                      *Bool      `xmlrpc:"is_blacklisted,omitempty"`
	IsCompany                          *Bool      `xmlrpc:"is_company,omitempty"`
	IsPublic                           *Bool      `xmlrpc:"is_public,omitempty"`
	IsSystem                           *Bool      `xmlrpc:"is_system,omitempty"`
	JobTitle                           *String    `xmlrpc:"job_title,omitempty"`
	JournalItemCount                   *Int       `xmlrpc:"journal_item_count,omitempty"`
	KmHomeWork                         *Int       `xmlrpc:"km_home_work,omitempty"`
	Lang                               *Selection `xmlrpc:"lang,omitempty"`
	LastActivity                       *Time      `xmlrpc:"last_activity,omitempty"`
	LastActivityTime                   *String    `xmlrpc:"last_activity_time,omitempty"`
	LastTimeEntriesChecked             *Time      `xmlrpc:"last_time_entries_checked,omitempty"`
	LogIds                             *Relation  `xmlrpc:"log_ids,omitempty"`
	Login                              *String    `xmlrpc:"login,omitempty"`
	LoginDate                          *Time      `xmlrpc:"login_date,omitempty"`
	Marital                            *Selection `xmlrpc:"marital,omitempty"`
	MessageAttachmentCount             *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageBounce                      *Int       `xmlrpc:"message_bounce,omitempty"`
	MessageFollowerIds                 *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                    *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter             *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                 *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                         *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                  *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                  *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter           *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                  *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	Mobile                             *String    `xmlrpc:"mobile,omitempty"`
	MobileBlacklisted                  *Bool      `xmlrpc:"mobile_blacklisted,omitempty"`
	MobilePhone                        *String    `xmlrpc:"mobile_phone,omitempty"`
	MyActivityDateDeadline             *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                               *String    `xmlrpc:"name,omitempty"`
	NewPassword                        *String    `xmlrpc:"new_password,omitempty"`
	NotificationType                   *Selection `xmlrpc:"notification_type,omitempty"`
	OauthAccessToken                   *String    `xmlrpc:"oauth_access_token,omitempty"`
	OauthProviderId                    *Many2One  `xmlrpc:"oauth_provider_id,omitempty"`
	OauthUid                           *String    `xmlrpc:"oauth_uid,omitempty"`
	OcnToken                           *String    `xmlrpc:"ocn_token,omitempty"`
	OdoobotFailed                      *Bool      `xmlrpc:"odoobot_failed,omitempty"`
	OdoobotState                       *Selection `xmlrpc:"odoobot_state,omitempty"`
	OnTimeRate                         *Float     `xmlrpc:"on_time_rate,omitempty"`
	ParentId                           *Many2One  `xmlrpc:"parent_id,omitempty"`
	ParentName                         *String    `xmlrpc:"parent_name,omitempty"`
	PartnerGid                         *Int       `xmlrpc:"partner_gid,omitempty"`
	PartnerId                          *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerLatitude                    *Float     `xmlrpc:"partner_latitude,omitempty"`
	PartnerLongitude                   *Float     `xmlrpc:"partner_longitude,omitempty"`
	PartnerShare                       *Bool      `xmlrpc:"partner_share,omitempty"`
	PassportId                         *String    `xmlrpc:"passport_id,omitempty"`
	Password                           *String    `xmlrpc:"password,omitempty"`
	PaymentTokenCount                  *Int       `xmlrpc:"payment_token_count,omitempty"`
	PaymentTokenIds                    *Relation  `xmlrpc:"payment_token_ids,omitempty"`
	PeppolEas                          *Selection `xmlrpc:"peppol_eas,omitempty"`
	PeppolEndpoint                     *String    `xmlrpc:"peppol_endpoint,omitempty"`
	PermitNo                           *String    `xmlrpc:"permit_no,omitempty"`
	Phone                              *String    `xmlrpc:"phone,omitempty"`
	PhoneBlacklisted                   *Bool      `xmlrpc:"phone_blacklisted,omitempty"`
	PhoneMobileSearch                  *String    `xmlrpc:"phone_mobile_search,omitempty"`
	PhoneSanitized                     *String    `xmlrpc:"phone_sanitized,omitempty"`
	PhoneSanitizedBlacklisted          *Bool      `xmlrpc:"phone_sanitized_blacklisted,omitempty"`
	PickingWarn                        *Selection `xmlrpc:"picking_warn,omitempty"`
	PickingWarnMsg                     *String    `xmlrpc:"picking_warn_msg,omitempty"`
	Pin                                *String    `xmlrpc:"pin,omitempty"`
	PlaceOfBirth                       *String    `xmlrpc:"place_of_birth,omitempty"`
	PosOrderCount                      *Int       `xmlrpc:"pos_order_count,omitempty"`
	PosOrderIds                        *Relation  `xmlrpc:"pos_order_ids,omitempty"`
	PrivateCity                        *String    `xmlrpc:"private_city,omitempty"`
	PrivateCountryId                   *Many2One  `xmlrpc:"private_country_id,omitempty"`
	PrivateEmail                       *String    `xmlrpc:"private_email,omitempty"`
	PrivateLang                        *Selection `xmlrpc:"private_lang,omitempty"`
	PrivatePhone                       *String    `xmlrpc:"private_phone,omitempty"`
	PrivateStateId                     *Many2One  `xmlrpc:"private_state_id,omitempty"`
	PrivateStreet                      *String    `xmlrpc:"private_street,omitempty"`
	PrivateStreet2                     *String    `xmlrpc:"private_street2,omitempty"`
	PrivateZip                         *String    `xmlrpc:"private_zip,omitempty"`
	PropertyAccountPayableId           *Many2One  `xmlrpc:"property_account_payable_id,omitempty"`
	PropertyAccountPositionId          *Many2One  `xmlrpc:"property_account_position_id,omitempty"`
	PropertyAccountReceivableId        *Many2One  `xmlrpc:"property_account_receivable_id,omitempty"`
	PropertyDeliveryCarrierId          *Many2One  `xmlrpc:"property_delivery_carrier_id,omitempty"`
	PropertyPaymentTermId              *Many2One  `xmlrpc:"property_payment_term_id,omitempty"`
	PropertyProductPricelist           *Many2One  `xmlrpc:"property_product_pricelist,omitempty"`
	PropertyPurchaseCurrencyId         *Many2One  `xmlrpc:"property_purchase_currency_id,omitempty"`
	PropertyStockCustomer              *Many2One  `xmlrpc:"property_stock_customer,omitempty"`
	PropertyStockSupplier              *Many2One  `xmlrpc:"property_stock_supplier,omitempty"`
	PropertySupplierPaymentTermId      *Many2One  `xmlrpc:"property_supplier_payment_term_id,omitempty"`
	PropertyWarehouseId                *Many2One  `xmlrpc:"property_warehouse_id,omitempty"`
	PurchaseLineIds                    *Relation  `xmlrpc:"purchase_line_ids,omitempty"`
	PurchaseOrderCount                 *Int       `xmlrpc:"purchase_order_count,omitempty"`
	PurchaseWarn                       *Selection `xmlrpc:"purchase_warn,omitempty"`
	PurchaseWarnMsg                    *String    `xmlrpc:"purchase_warn_msg,omitempty"`
	RatingIds                          *Relation  `xmlrpc:"rating_ids,omitempty"`
	ReceiptReminderEmail               *Bool      `xmlrpc:"receipt_reminder_email,omitempty"`
	Ref                                *String    `xmlrpc:"ref,omitempty"`
	RefCompanyIds                      *Relation  `xmlrpc:"ref_company_ids,omitempty"`
	ReminderDateBeforeReceipt          *Int       `xmlrpc:"reminder_date_before_receipt,omitempty"`
	ResUsersSettingsId                 *Many2One  `xmlrpc:"res_users_settings_id,omitempty"`
	ResUsersSettingsIds                *Relation  `xmlrpc:"res_users_settings_ids,omitempty"`
	ResourceCalendarId                 *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceIds                        *Relation  `xmlrpc:"resource_ids,omitempty"`
	ResumeLineIds                      *Relation  `xmlrpc:"resume_line_ids,omitempty"`
	RulesCount                         *Int       `xmlrpc:"rules_count,omitempty"`
	SaleOrderCount                     *Int       `xmlrpc:"sale_order_count,omitempty"`
	SaleOrderIds                       *Relation  `xmlrpc:"sale_order_ids,omitempty"`
	SaleTeamId                         *Many2One  `xmlrpc:"sale_team_id,omitempty"`
	SaleWarn                           *Selection `xmlrpc:"sale_warn,omitempty"`
	SaleWarnMsg                        *String    `xmlrpc:"sale_warn_msg,omitempty"`
	SameCompanyRegistryPartnerId       *Many2One  `xmlrpc:"same_company_registry_partner_id,omitempty"`
	SameVatPartnerId                   *Many2One  `xmlrpc:"same_vat_partner_id,omitempty"`
	Self                               *Many2One  `xmlrpc:"self,omitempty"`
	Share                              *Bool      `xmlrpc:"share,omitempty"`
	ShowCreditLimit                    *Bool      `xmlrpc:"show_credit_limit,omitempty"`
	Signature                          *String    `xmlrpc:"signature,omitempty"`
	SignupExpiration                   *Time      `xmlrpc:"signup_expiration,omitempty"`
	SignupToken                        *String    `xmlrpc:"signup_token,omitempty"`
	SignupType                         *String    `xmlrpc:"signup_type,omitempty"`
	SignupUrl                          *String    `xmlrpc:"signup_url,omitempty"`
	SignupValid                        *Bool      `xmlrpc:"signup_valid,omitempty"`
	SpouseBirthdate                    *Time      `xmlrpc:"spouse_birthdate,omitempty"`
	SpouseCompleteName                 *String    `xmlrpc:"spouse_complete_name,omitempty"`
	Ssnid                              *String    `xmlrpc:"ssnid,omitempty"`
	StarredMessageIds                  *Relation  `xmlrpc:"starred_message_ids,omitempty"`
	State                              *Selection `xmlrpc:"state,omitempty"`
	StateId                            *Many2One  `xmlrpc:"state_id,omitempty"`
	Street                             *String    `xmlrpc:"street,omitempty"`
	Street2                            *String    `xmlrpc:"street2,omitempty"`
	StudyField                         *String    `xmlrpc:"study_field,omitempty"`
	StudySchool                        *String    `xmlrpc:"study_school,omitempty"`
	SupplierInvoiceCount               *Int       `xmlrpc:"supplier_invoice_count,omitempty"`
	SupplierRank                       *Int       `xmlrpc:"supplier_rank,omitempty"`
	TeamId                             *Many2One  `xmlrpc:"team_id,omitempty"`
	Title                              *Many2One  `xmlrpc:"title,omitempty"`
	TotalInvoiced                      *Float     `xmlrpc:"total_invoiced,omitempty"`
	TotpEnabled                        *Bool      `xmlrpc:"totp_enabled,omitempty"`
	TotpSecret                         *String    `xmlrpc:"totp_secret,omitempty"`
	TotpTrustedDeviceIds               *Relation  `xmlrpc:"totp_trusted_device_ids,omitempty"`
	Trust                              *Selection `xmlrpc:"trust,omitempty"`
	Type                               *Selection `xmlrpc:"type,omitempty"`
	Tz                                 *Selection `xmlrpc:"tz,omitempty"`
	TzOffset                           *String    `xmlrpc:"tz_offset,omitempty"`
	UblCiiFormat                       *Selection `xmlrpc:"ubl_cii_format,omitempty"`
	UsePartnerCreditLimit              *Bool      `xmlrpc:"use_partner_credit_limit,omitempty"`
	UserGroupWarning                   *String    `xmlrpc:"user_group_warning,omitempty"`
	UserId                             *Many2One  `xmlrpc:"user_id,omitempty"`
	UserIds                            *Relation  `xmlrpc:"user_ids,omitempty"`
	Vat                                *String    `xmlrpc:"vat,omitempty"`
	VisaExpire                         *Time      `xmlrpc:"visa_expire,omitempty"`
	VisaNo                             *String    `xmlrpc:"visa_no,omitempty"`
	Website                            *String    `xmlrpc:"website,omitempty"`
	WebsiteMessageIds                  *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorkContactId                      *Many2One  `xmlrpc:"work_contact_id,omitempty"`
	WorkEmail                          *String    `xmlrpc:"work_email,omitempty"`
	WorkLocationId                     *Many2One  `xmlrpc:"work_location_id,omitempty"`
	WorkPhone                          *String    `xmlrpc:"work_phone,omitempty"`
	WriteDate                          *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                           *Many2One  `xmlrpc:"write_uid,omitempty"`
	XStudioDefaultWarehouse            *Many2One  `xmlrpc:"x_studio_default_warehouse,omitempty"`
	XStudioJobPosition                 *String    `xmlrpc:"x_studio_job_position,omitempty"`
	XStudioJobPosition1                *String    `xmlrpc:"x_studio_job_position_1,omitempty"`
	XStudioMany2ManyField5F11Htukc431  *Relation  `xmlrpc:"x_studio_many2many_field_5f1_1htukc431,omitempty"`
	XStudioMany2OneField6121Hsufid8C   *Many2One  `xmlrpc:"x_studio_many2one_field_612_1hsufid8c,omitempty"`
	XStudioMany2OneField7481Htui26F2   *Many2One  `xmlrpc:"x_studio_many2one_field_748_1htui26f2,omitempty"`
	XStudioOne2ManyField10B1Htuk9614   *Relation  `xmlrpc:"x_studio_one2many_field_10b_1htuk9614,omitempty"`
	XStudioPlaceOfPerson               *Selection `xmlrpc:"x_studio_place_of_person,omitempty"`
	XStudioPrimaryWarehouse            *Many2One  `xmlrpc:"x_studio_primary_warehouse,omitempty"`
	XStudioQualityTeam                 *Relation  `xmlrpc:"x_studio_quality_team,omitempty"`
	XStudioRelatedField4Hi1Htugmkh4    *String    `xmlrpc:"x_studio_related_field_4hi_1htugmkh4,omitempty"`
	XStudioRelatedField8Q11Huaplps4    *Many2One  `xmlrpc:"x_studio_related_field_8q1_1huaplps4,omitempty"`
	XStudioRelatedField9Hd1Htugr876    *String    `xmlrpc:"x_studio_related_field_9hd_1htugr876,omitempty"`
	XStudioRelatedFieldEp1Htug3M6L     *String    `xmlrpc:"x_studio_related_field_ep_1htug3m6l,omitempty"`
	XStudioSelectionField70C1Huaq2Mse  *Selection `xmlrpc:"x_studio_selection_field_70c_1huaq2mse,omitempty"`
	Zip                                *String    `xmlrpc:"zip,omitempty"`
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
	ids, err := c.CreateResUserss([]*ResUsers{ru})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsers creates a new res.users model and returns its id.
func (c *Client) CreateResUserss(rus []*ResUsers) ([]int64, error) {
	var vv []interface{}
	for _, v := range rus {
		vv = append(vv, v)
	}
	return c.Create(ResUsersModel, vv, nil)
}

// UpdateResUsers updates an existing res.users record.
func (c *Client) UpdateResUsers(ru *ResUsers) error {
	return c.UpdateResUserss([]int64{ru.Id.Get()}, ru)
}

// UpdateResUserss updates existing res.users records.
// All records (represented by ids) will be updated by ru values.
func (c *Client) UpdateResUserss(ids []int64, ru *ResUsers) error {
	return c.Update(ResUsersModel, ids, ru, nil)
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
	return &((*rus)[0]), nil
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
	return &((*rus)[0]), nil
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
	return c.Search(ResUsersModel, criteria, options)
}

// FindResUsersId finds record id by querying it with criteria.
func (c *Client) FindResUsersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
