package odoo

// ResUsers represents res.users model.
type ResUsers struct {
	AccessesCount                        *Int        `xmlrpc:"accesses_count,omitempty"`
	AccountMoveCount                     *Int        `xmlrpc:"account_move_count,omitempty"`
	ActionId                             *Many2One   `xmlrpc:"action_id,omitempty"`
	Active                               *Bool       `xmlrpc:"active,omitempty"`
	ActiveLangCount                      *Int        `xmlrpc:"active_lang_count,omitempty"`
	ActivePartner                        *Bool       `xmlrpc:"active_partner,omitempty"`
	ActivityCalendarEventId              *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline                 *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration          *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                          *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                        *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                      *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                     *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                       *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                       *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AdditionalNote                       *String     `xmlrpc:"additional_note,omitempty"`
	AllGroupIds                          *Relation   `xmlrpc:"all_group_ids,omitempty"`
	ApiKeyIds                            *Relation   `xmlrpc:"api_key_ids,omitempty"`
	ApplicationStatistics                interface{} `xmlrpc:"application_statistics,omitempty"`
	AuthPasskeyKeyIds                    *Relation   `xmlrpc:"auth_passkey_key_ids,omitempty"`
	AutopostBills                        *Selection  `xmlrpc:"autopost_bills,omitempty"`
	AvailableInvoiceTemplatePdfReportIds *Relation   `xmlrpc:"available_invoice_template_pdf_report_ids,omitempty"`
	AvailablePeppolEas                   interface{} `xmlrpc:"available_peppol_eas,omitempty"`
	AvailablePeppolEdiFormats            interface{} `xmlrpc:"available_peppol_edi_formats,omitempty"`
	AvailablePeppolSendingMethods        interface{} `xmlrpc:"available_peppol_sending_methods,omitempty"`
	Avatar1024                           *String     `xmlrpc:"avatar_1024,omitempty"`
	Avatar128                            *String     `xmlrpc:"avatar_128,omitempty"`
	Avatar1920                           *String     `xmlrpc:"avatar_1920,omitempty"`
	Avatar256                            *String     `xmlrpc:"avatar_256,omitempty"`
	Avatar512                            *String     `xmlrpc:"avatar_512,omitempty"`
	BankAccountCount                     *Int        `xmlrpc:"bank_account_count,omitempty"`
	BankAccountIds                       *Relation   `xmlrpc:"bank_account_ids,omitempty"`
	BankIds                              *Relation   `xmlrpc:"bank_ids,omitempty"`
	Barcode                              *String     `xmlrpc:"barcode,omitempty"`
	BuyerId                              *Many2One   `xmlrpc:"buyer_id,omitempty"`
	CalendarDefaultPrivacy               *Selection  `xmlrpc:"calendar_default_privacy,omitempty"`
	CalendarLastNotifAck                 *Time       `xmlrpc:"calendar_last_notif_ack,omitempty"`
	CanEditRole                          *Bool       `xmlrpc:"can_edit_role,omitempty"`
	CategoryId                           *Relation   `xmlrpc:"category_id,omitempty"`
	CategoryIds                          *Relation   `xmlrpc:"category_ids,omitempty"`
	ChannelIds                           *Relation   `xmlrpc:"channel_ids,omitempty"`
	ChannelMemberIds                     *Relation   `xmlrpc:"channel_member_ids,omitempty"`
	ChatbotScriptIds                     *Relation   `xmlrpc:"chatbot_script_ids,omitempty"`
	ChildIds                             *Relation   `xmlrpc:"child_ids,omitempty"`
	City                                 *String     `xmlrpc:"city,omitempty"`
	Color                                *Int        `xmlrpc:"color,omitempty"`
	Comment                              *String     `xmlrpc:"comment,omitempty"`
	CommercialCompanyName                *String     `xmlrpc:"commercial_company_name,omitempty"`
	CommercialPartnerId                  *Many2One   `xmlrpc:"commercial_partner_id,omitempty"`
	CompaniesCount                       *Int        `xmlrpc:"companies_count,omitempty"`
	CompanyId                            *Many2One   `xmlrpc:"company_id,omitempty"`
	CompanyIds                           *Relation   `xmlrpc:"company_ids,omitempty"`
	CompanyName                          *String     `xmlrpc:"company_name,omitempty"`
	CompanyRegistry                      *String     `xmlrpc:"company_registry,omitempty"`
	CompanyRegistryLabel                 *String     `xmlrpc:"company_registry_label,omitempty"`
	CompanyRegistryPlaceholder           *String     `xmlrpc:"company_registry_placeholder,omitempty"`
	CompanyType                          *Selection  `xmlrpc:"company_type,omitempty"`
	CompleteName                         *String     `xmlrpc:"complete_name,omitempty"`
	ContactAddress                       *String     `xmlrpc:"contact_address,omitempty"`
	ContactAddressInline                 *String     `xmlrpc:"contact_address_inline,omitempty"`
	ContractIds                          *Relation   `xmlrpc:"contract_ids,omitempty"`
	CountryCode                          *String     `xmlrpc:"country_code,omitempty"`
	CountryId                            *Many2One   `xmlrpc:"country_id,omitempty"`
	CreateDate                           *Time       `xmlrpc:"create_date,omitempty"`
	CreateEmployee                       *Bool       `xmlrpc:"create_employee,omitempty"`
	CreateEmployeeId                     *Many2One   `xmlrpc:"create_employee_id,omitempty"`
	CreateUid                            *Many2One   `xmlrpc:"create_uid,omitempty"`
	Credit                               *Float      `xmlrpc:"credit,omitempty"`
	CreditLimit                          *Float      `xmlrpc:"credit_limit,omitempty"`
	CreditToInvoice                      *Float      `xmlrpc:"credit_to_invoice,omitempty"`
	CrmTeamIds                           *Relation   `xmlrpc:"crm_team_ids,omitempty"`
	CrmTeamMemberIds                     *Relation   `xmlrpc:"crm_team_member_ids,omitempty"`
	CurrencyId                           *Many2One   `xmlrpc:"currency_id,omitempty"`
	CustomerRank                         *Int        `xmlrpc:"customer_rank,omitempty"`
	DaysSalesOutstanding                 *Float      `xmlrpc:"days_sales_outstanding,omitempty"`
	Debit                                *Float      `xmlrpc:"debit,omitempty"`
	DeviceIds                            *Relation   `xmlrpc:"device_ids,omitempty"`
	DisplayInvoiceEdiFormat              *Bool       `xmlrpc:"display_invoice_edi_format,omitempty"`
	DisplayInvoiceTemplatePdfReportId    *Bool       `xmlrpc:"display_invoice_template_pdf_report_id,omitempty"`
	DisplayName                          *String     `xmlrpc:"display_name,omitempty"`
	DuplicateBankPartnerIds              *Relation   `xmlrpc:"duplicate_bank_partner_ids,omitempty"`
	Email                                *String     `xmlrpc:"email,omitempty"`
	EmailDomainPlaceholder               *String     `xmlrpc:"email_domain_placeholder,omitempty"`
	EmailFormatted                       *String     `xmlrpc:"email_formatted,omitempty"`
	EmailNormalized                      *String     `xmlrpc:"email_normalized,omitempty"`
	EmergencyContact                     *String     `xmlrpc:"emergency_contact,omitempty"`
	EmergencyPhone                       *String     `xmlrpc:"emergency_phone,omitempty"`
	Employee                             *Bool       `xmlrpc:"employee,omitempty"`
	EmployeeBankAccountIds               *Relation   `xmlrpc:"employee_bank_account_ids,omitempty"`
	EmployeeCount                        *Int        `xmlrpc:"employee_count,omitempty"`
	EmployeeId                           *Many2One   `xmlrpc:"employee_id,omitempty"`
	EmployeeIds                          *Relation   `xmlrpc:"employee_ids,omitempty"`
	EmployeeResourceCalendarId           *Many2One   `xmlrpc:"employee_resource_calendar_id,omitempty"`
	EmployeeSkillIds                     *Relation   `xmlrpc:"employee_skill_ids,omitempty"`
	EmployeesCount                       *Int        `xmlrpc:"employees_count,omitempty"`
	FavoriteProjectIds                   *Relation   `xmlrpc:"favorite_project_ids,omitempty"`
	FiscalCountryCodes                   *String     `xmlrpc:"fiscal_country_codes,omitempty"`
	FiscalCountryGroupCodes              interface{} `xmlrpc:"fiscal_country_group_codes,omitempty"`
	FridayLocationId                     *Many2One   `xmlrpc:"friday_location_id,omitempty"`
	Function                             *String     `xmlrpc:"function,omitempty"`
	GlobalLocationNumber                 *String     `xmlrpc:"global_location_number,omitempty"`
	GroupIds                             *Relation   `xmlrpc:"group_ids,omitempty"`
	GroupOn                              *Selection  `xmlrpc:"group_on,omitempty"`
	GroupRfq                             *Selection  `xmlrpc:"group_rfq,omitempty"`
	GroupsCount                          *Int        `xmlrpc:"groups_count,omitempty"`
	HasAccessLivechat                    *Bool       `xmlrpc:"has_access_livechat,omitempty"`
	HasExternalMailServer                *Bool       `xmlrpc:"has_external_mail_server,omitempty"`
	HasMessage                           *Bool       `xmlrpc:"has_message,omitempty"`
	Id                                   *Int        `xmlrpc:"id,omitempty"`
	IgnoreAbnormalInvoiceAmount          *Bool       `xmlrpc:"ignore_abnormal_invoice_amount,omitempty"`
	IgnoreAbnormalInvoiceDate            *Bool       `xmlrpc:"ignore_abnormal_invoice_date,omitempty"`
	ImStatus                             *String     `xmlrpc:"im_status,omitempty"`
	Image1024                            *String     `xmlrpc:"image_1024,omitempty"`
	Image128                             *String     `xmlrpc:"image_128,omitempty"`
	Image1920                            *String     `xmlrpc:"image_1920,omitempty"`
	Image256                             *String     `xmlrpc:"image_256,omitempty"`
	Image512                             *String     `xmlrpc:"image_512,omitempty"`
	IndustryId                           *Many2One   `xmlrpc:"industry_id,omitempty"`
	InvoiceEdiFormat                     *Selection  `xmlrpc:"invoice_edi_format,omitempty"`
	InvoiceEdiFormatStore                *String     `xmlrpc:"invoice_edi_format_store,omitempty"`
	InvoiceIds                           *Relation   `xmlrpc:"invoice_ids,omitempty"`
	InvoiceSendingMethod                 *Selection  `xmlrpc:"invoice_sending_method,omitempty"`
	InvoiceTemplatePdfReportId           *Many2One   `xmlrpc:"invoice_template_pdf_report_id,omitempty"`
	IsBlacklisted                        *Bool       `xmlrpc:"is_blacklisted,omitempty"`
	IsCompany                            *Bool       `xmlrpc:"is_company,omitempty"`
	IsHrUser                             *Bool       `xmlrpc:"is_hr_user,omitempty"`
	IsInCall                             *Bool       `xmlrpc:"is_in_call,omitempty"`
	IsOutOfOffice                        *Bool       `xmlrpc:"is_out_of_office,omitempty"`
	IsPeppolEdiFormat                    *Bool       `xmlrpc:"is_peppol_edi_format,omitempty"`
	IsPublic                             *Bool       `xmlrpc:"is_public,omitempty"`
	IsSystem                             *Bool       `xmlrpc:"is_system,omitempty"`
	IsUblFormat                          *Bool       `xmlrpc:"is_ubl_format,omitempty"`
	JobTitle                             *String     `xmlrpc:"job_title,omitempty"`
	KmHomeWork                           *Int        `xmlrpc:"km_home_work,omitempty"`
	L10NFrIsFrench                       *Bool       `xmlrpc:"l10n_fr_is_french,omitempty"`
	Lang                                 *Selection  `xmlrpc:"lang,omitempty"`
	LeaveDateTo                          *Time       `xmlrpc:"leave_date_to,omitempty"`
	LivechatChannelCount                 *Int        `xmlrpc:"livechat_channel_count,omitempty"`
	LivechatChannelIds                   *Relation   `xmlrpc:"livechat_channel_ids,omitempty"`
	LivechatExpertiseIds                 *Relation   `xmlrpc:"livechat_expertise_ids,omitempty"`
	LivechatIsInCall                     *Bool       `xmlrpc:"livechat_is_in_call,omitempty"`
	LivechatLangIds                      *Relation   `xmlrpc:"livechat_lang_ids,omitempty"`
	LivechatOngoingSessionCount          *Int        `xmlrpc:"livechat_ongoing_session_count,omitempty"`
	LivechatUsername                     *String     `xmlrpc:"livechat_username,omitempty"`
	LogIds                               *Relation   `xmlrpc:"log_ids,omitempty"`
	Login                                *String     `xmlrpc:"login,omitempty"`
	LoginDate                            *Time       `xmlrpc:"login_date,omitempty"`
	MachineOrganizationName              *String     `xmlrpc:"machine_organization_name,omitempty"`
	MachineUserEmail                     *String     `xmlrpc:"machine_user_email,omitempty"`
	MachineUserLogin                     *String     `xmlrpc:"machine_user_login,omitempty"`
	MainUserId                           *Many2One   `xmlrpc:"main_user_id,omitempty"`
	ManualImStatus                       *Selection  `xmlrpc:"manual_im_status,omitempty"`
	MeetingCount                         *Int        `xmlrpc:"meeting_count,omitempty"`
	MeetingIds                           *Relation   `xmlrpc:"meeting_ids,omitempty"`
	MessageAttachmentCount               *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageBounce                        *Int        `xmlrpc:"message_bounce,omitempty"`
	MessageFollowerIds                   *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                      *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter               *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                   *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                           *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                    *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                    *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter             *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                    *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MobilePhone                          *String     `xmlrpc:"mobile_phone,omitempty"`
	MondayLocationId                     *Many2One   `xmlrpc:"monday_location_id,omitempty"`
	MyActivityDateDeadline               *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                 *String     `xmlrpc:"name,omitempty"`
	NewPassword                          *String     `xmlrpc:"new_password,omitempty"`
	NotificationType                     *Selection  `xmlrpc:"notification_type,omitempty"`
	OdoobotFailed                        *Bool       `xmlrpc:"odoobot_failed,omitempty"`
	OdoobotState                         *Selection  `xmlrpc:"odoobot_state,omitempty"`
	OfflineSince                         *Time       `xmlrpc:"offline_since,omitempty"`
	OnTimeRate                           *Float      `xmlrpc:"on_time_rate,omitempty"`
	OpportunityCount                     *Int        `xmlrpc:"opportunity_count,omitempty"`
	OpportunityIds                       *Relation   `xmlrpc:"opportunity_ids,omitempty"`
	OutOfOfficeFrom                      *Time       `xmlrpc:"out_of_office_from,omitempty"`
	OutOfOfficeMessage                   *String     `xmlrpc:"out_of_office_message,omitempty"`
	OutOfOfficeTo                        *Time       `xmlrpc:"out_of_office_to,omitempty"`
	OutgoingMailServerId                 *Many2One   `xmlrpc:"outgoing_mail_server_id,omitempty"`
	OutgoingMailServerType               *Selection  `xmlrpc:"outgoing_mail_server_type,omitempty"`
	ParentId                             *Many2One   `xmlrpc:"parent_id,omitempty"`
	ParentName                           *String     `xmlrpc:"parent_name,omitempty"`
	PartnerCompanyRegistryPlaceholder    *String     `xmlrpc:"partner_company_registry_placeholder,omitempty"`
	PartnerId                            *Many2One   `xmlrpc:"partner_id,omitempty"`
	PartnerLatitude                      *Float      `xmlrpc:"partner_latitude,omitempty"`
	PartnerLongitude                     *Float      `xmlrpc:"partner_longitude,omitempty"`
	PartnerShare                         *Bool       `xmlrpc:"partner_share,omitempty"`
	PartnerVatPlaceholder                *String     `xmlrpc:"partner_vat_placeholder,omitempty"`
	Password                             *String     `xmlrpc:"password,omitempty"`
	PaymentTokenCount                    *Int        `xmlrpc:"payment_token_count,omitempty"`
	PaymentTokenIds                      *Relation   `xmlrpc:"payment_token_ids,omitempty"`
	PdpVerificationDisplayState          *Selection  `xmlrpc:"pdp_verification_display_state,omitempty"`
	PeppolEas                            *Selection  `xmlrpc:"peppol_eas,omitempty"`
	PeppolEndpoint                       *String     `xmlrpc:"peppol_endpoint,omitempty"`
	PeppolResponseSupport                *Bool       `xmlrpc:"peppol_response_support,omitempty"`
	PeppolSupportedDocuments             interface{} `xmlrpc:"peppol_supported_documents,omitempty"`
	PeppolVerificationState              *Selection  `xmlrpc:"peppol_verification_state,omitempty"`
	PerformViesValidation                *Bool       `xmlrpc:"perform_vies_validation,omitempty"`
	Phone                                *String     `xmlrpc:"phone,omitempty"`
	PhoneBlacklisted                     *Bool       `xmlrpc:"phone_blacklisted,omitempty"`
	PhoneMobileSearch                    *String     `xmlrpc:"phone_mobile_search,omitempty"`
	PhoneSanitized                       *String     `xmlrpc:"phone_sanitized,omitempty"`
	PhoneSanitizedBlacklisted            *Bool       `xmlrpc:"phone_sanitized_blacklisted,omitempty"`
	PickingWarnMsg                       *String     `xmlrpc:"picking_warn_msg,omitempty"`
	Pin                                  *String     `xmlrpc:"pin,omitempty"`
	PresenceIds                          *Relation   `xmlrpc:"presence_ids,omitempty"`
	PrivateCity                          *String     `xmlrpc:"private_city,omitempty"`
	PrivateCountryId                     *Many2One   `xmlrpc:"private_country_id,omitempty"`
	PrivateEmail                         *String     `xmlrpc:"private_email,omitempty"`
	PrivatePhone                         *String     `xmlrpc:"private_phone,omitempty"`
	PrivateStateId                       *Many2One   `xmlrpc:"private_state_id,omitempty"`
	PrivateStreet                        *String     `xmlrpc:"private_street,omitempty"`
	PrivateStreet2                       *String     `xmlrpc:"private_street2,omitempty"`
	PrivateZip                           *String     `xmlrpc:"private_zip,omitempty"`
	ProjectIds                           *Relation   `xmlrpc:"project_ids,omitempty"`
	Properties                           interface{} `xmlrpc:"properties,omitempty"`
	PropertiesBaseDefinitionId           *Many2One   `xmlrpc:"properties_base_definition_id,omitempty"`
	PropertyAccountPayableId             *Many2One   `xmlrpc:"property_account_payable_id,omitempty"`
	PropertyAccountPositionId            *Many2One   `xmlrpc:"property_account_position_id,omitempty"`
	PropertyAccountReceivableId          *Many2One   `xmlrpc:"property_account_receivable_id,omitempty"`
	PropertyAutosalesConfig              *Many2One   `xmlrpc:"property_autosales_config,omitempty"`
	PropertyInboundPaymentMethodLineId   *Many2One   `xmlrpc:"property_inbound_payment_method_line_id,omitempty"`
	PropertyOutboundPaymentMethodLineId  *Many2One   `xmlrpc:"property_outbound_payment_method_line_id,omitempty"`
	PropertyPaymentTermId                *Many2One   `xmlrpc:"property_payment_term_id,omitempty"`
	PropertyProductPricelist             *Many2One   `xmlrpc:"property_product_pricelist,omitempty"`
	PropertyPurchaseCurrencyId           *Many2One   `xmlrpc:"property_purchase_currency_id,omitempty"`
	PropertyStockCustomer                *Many2One   `xmlrpc:"property_stock_customer,omitempty"`
	PropertyStockSupplier                *Many2One   `xmlrpc:"property_stock_supplier,omitempty"`
	PropertySupplierPaymentTermId        *Many2One   `xmlrpc:"property_supplier_payment_term_id,omitempty"`
	PropertyWarehouseId                  *Many2One   `xmlrpc:"property_warehouse_id,omitempty"`
	PurchaseLineIds                      *Relation   `xmlrpc:"purchase_line_ids,omitempty"`
	PurchaseOrderCount                   *Int        `xmlrpc:"purchase_order_count,omitempty"`
	PurchaseWarnMsg                      *String     `xmlrpc:"purchase_warn_msg,omitempty"`
	RatingIds                            *Relation   `xmlrpc:"rating_ids,omitempty"`
	ReceiptReminderEmail                 *Bool       `xmlrpc:"receipt_reminder_email,omitempty"`
	Ref                                  *String     `xmlrpc:"ref,omitempty"`
	RefCompanyIds                        *Relation   `xmlrpc:"ref_company_ids,omitempty"`
	ReminderDateBeforeReceipt            *Int        `xmlrpc:"reminder_date_before_receipt,omitempty"`
	ResUsersSettingsId                   *Many2One   `xmlrpc:"res_users_settings_id,omitempty"`
	ResUsersSettingsIds                  *Relation   `xmlrpc:"res_users_settings_ids,omitempty"`
	ResourceCalendarId                   *Many2One   `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceIds                          *Relation   `xmlrpc:"resource_ids,omitempty"`
	Role                                 *Selection  `xmlrpc:"role,omitempty"`
	RoleIds                              *Relation   `xmlrpc:"role_ids,omitempty"`
	RtcSessionIds                        *Relation   `xmlrpc:"rtc_session_ids,omitempty"`
	RulesCount                           *Int        `xmlrpc:"rules_count,omitempty"`
	SaleOrderCount                       *Int        `xmlrpc:"sale_order_count,omitempty"`
	SaleOrderIds                         *Relation   `xmlrpc:"sale_order_ids,omitempty"`
	SaleTeamId                           *Many2One   `xmlrpc:"sale_team_id,omitempty"`
	SaleWarnMsg                          *String     `xmlrpc:"sale_warn_msg,omitempty"`
	SameCompanyRegistryPartnerId         *Many2One   `xmlrpc:"same_company_registry_partner_id,omitempty"`
	SameVatPartnerId                     *Many2One   `xmlrpc:"same_vat_partner_id,omitempty"`
	SaturdayLocationId                   *Many2One   `xmlrpc:"saturday_location_id,omitempty"`
	Self                                 *Many2One   `xmlrpc:"self,omitempty"`
	Share                                *Bool       `xmlrpc:"share,omitempty"`
	ShowCreditLimit                      *Bool       `xmlrpc:"show_credit_limit,omitempty"`
	Signature                            *String     `xmlrpc:"signature,omitempty"`
	SignupType                           *String     `xmlrpc:"signup_type,omitempty"`
	SpecificPropertyAutosalesConfig      *Many2One   `xmlrpc:"specific_property_autosales_config,omitempty"`
	SpecificPropertyProductPricelist     *Many2One   `xmlrpc:"specific_property_product_pricelist,omitempty"`
	State                                *Selection  `xmlrpc:"state,omitempty"`
	StateId                              *Many2One   `xmlrpc:"state_id,omitempty"`
	Street                               *String     `xmlrpc:"street,omitempty"`
	Street2                              *String     `xmlrpc:"street2,omitempty"`
	SuggestBasedOn                       *String     `xmlrpc:"suggest_based_on,omitempty"`
	SuggestDays                          *Int        `xmlrpc:"suggest_days,omitempty"`
	SuggestPercent                       *Int        `xmlrpc:"suggest_percent,omitempty"`
	SundayLocationId                     *Many2One   `xmlrpc:"sunday_location_id,omitempty"`
	SupplierInvoiceCount                 *Int        `xmlrpc:"supplier_invoice_count,omitempty"`
	SupplierRank                         *Int        `xmlrpc:"supplier_rank,omitempty"`
	TaskCount                            *Int        `xmlrpc:"task_count,omitempty"`
	TaskIds                              *Relation   `xmlrpc:"task_ids,omitempty"`
	ThursdayLocationId                   *Many2One   `xmlrpc:"thursday_location_id,omitempty"`
	TotalInvoiced                        *Float      `xmlrpc:"total_invoiced,omitempty"`
	TotpEnabled                          *Bool       `xmlrpc:"totp_enabled,omitempty"`
	TotpLastCounter                      *Int        `xmlrpc:"totp_last_counter,omitempty"`
	TotpSecret                           *String     `xmlrpc:"totp_secret,omitempty"`
	TotpTrustedDeviceIds                 *Relation   `xmlrpc:"totp_trusted_device_ids,omitempty"`
	TourEnabled                          *Bool       `xmlrpc:"tour_enabled,omitempty"`
	Trust                                *Selection  `xmlrpc:"trust,omitempty"`
	TuesdayLocationId                    *Many2One   `xmlrpc:"tuesday_location_id,omitempty"`
	Type                                 *Selection  `xmlrpc:"type,omitempty"`
	TypeAddressLabel                     *String     `xmlrpc:"type_address_label,omitempty"`
	Tz                                   *Selection  `xmlrpc:"tz,omitempty"`
	TzOffset                             *String     `xmlrpc:"tz_offset,omitempty"`
	UsePartnerCreditLimit                *Bool       `xmlrpc:"use_partner_credit_limit,omitempty"`
	UserId                               *Many2One   `xmlrpc:"user_id,omitempty"`
	UserIds                              *Relation   `xmlrpc:"user_ids,omitempty"`
	UserLivechatUsername                 *String     `xmlrpc:"user_livechat_username,omitempty"`
	Vat                                  *String     `xmlrpc:"vat,omitempty"`
	VatLabel                             *String     `xmlrpc:"vat_label,omitempty"`
	ViesValid                            *Bool       `xmlrpc:"vies_valid,omitempty"`
	ViewGroupHierarchy                   interface{} `xmlrpc:"view_group_hierarchy,omitempty"`
	VisaExpire                           *Time       `xmlrpc:"visa_expire,omitempty"`
	Website                              *String     `xmlrpc:"website,omitempty"`
	WebsiteMessageIds                    *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WednesdayLocationId                  *Many2One   `xmlrpc:"wednesday_location_id,omitempty"`
	WorkContactId                        *Many2One   `xmlrpc:"work_contact_id,omitempty"`
	WorkEmail                            *String     `xmlrpc:"work_email,omitempty"`
	WorkLocationId                       *Many2One   `xmlrpc:"work_location_id,omitempty"`
	WorkLocationName                     *String     `xmlrpc:"work_location_name,omitempty"`
	WorkLocationType                     *Selection  `xmlrpc:"work_location_type,omitempty"`
	WorkPhone                            *String     `xmlrpc:"work_phone,omitempty"`
	WriteDate                            *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                             *Many2One   `xmlrpc:"write_uid,omitempty"`
	Zip                                  *String     `xmlrpc:"zip,omitempty"`
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
