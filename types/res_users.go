package types

import (
	"time"
)

type ResUsers struct {
	LastUpdate                    time.Time `xmlrpc:"__last_update"`
	ActionId                      Many2One  `xmlrpc:"action_id"`
	Active                        bool      `xmlrpc:"active"`
	ActivitiesCount               int64     `xmlrpc:"activities_count"`
	AliasContact                  string    `xmlrpc:"alias_contact"`
	AliasId                       Many2One  `xmlrpc:"alias_id"`
	BankAccountCount              int64     `xmlrpc:"bank_account_count"`
	BankIds                       []int64   `xmlrpc:"bank_ids"`
	Barcode                       string    `xmlrpc:"barcode"`
	CalendarLastNotifAck          time.Time `xmlrpc:"calendar_last_notif_ack"`
	CategoryId                    []int64   `xmlrpc:"category_id"`
	ChannelIds                    []int64   `xmlrpc:"channel_ids"`
	ChildIds                      []int64   `xmlrpc:"child_ids"`
	City                          string    `xmlrpc:"city"`
	Color                         int64     `xmlrpc:"color"`
	Comment                       string    `xmlrpc:"comment"`
	CommercialCompanyName         string    `xmlrpc:"commercial_company_name"`
	CommercialPartnerId           Many2One  `xmlrpc:"commercial_partner_id"`
	CompaniesCount                int64     `xmlrpc:"companies_count"`
	CompanyId                     Many2One  `xmlrpc:"company_id"`
	CompanyIds                    []int64   `xmlrpc:"company_ids"`
	CompanyName                   string    `xmlrpc:"company_name"`
	CompanyType                   string    `xmlrpc:"company_type"`
	ContactAddress                string    `xmlrpc:"contact_address"`
	ContractIds                   []int64   `xmlrpc:"contract_ids"`
	ContractsCount                int64     `xmlrpc:"contracts_count"`
	CountryId                     Many2One  `xmlrpc:"country_id"`
	CreateDate                    time.Time `xmlrpc:"create_date"`
	CreateUid                     Many2One  `xmlrpc:"create_uid"`
	Credit                        float64   `xmlrpc:"credit"`
	CreditLimit                   float64   `xmlrpc:"credit_limit"`
	CurrencyId                    Many2One  `xmlrpc:"currency_id"`
	Customer                      bool      `xmlrpc:"customer"`
	Date                          time.Time `xmlrpc:"date"`
	Debit                         float64   `xmlrpc:"debit"`
	DebitLimit                    float64   `xmlrpc:"debit_limit"`
	DisplayName                   string    `xmlrpc:"display_name"`
	Email                         string    `xmlrpc:"email"`
	EmailFormatted                string    `xmlrpc:"email_formatted"`
	Employee                      bool      `xmlrpc:"employee"`
	EmployeeIds                   []int64   `xmlrpc:"employee_ids"`
	Fax                           string    `xmlrpc:"fax"`
	Function                      string    `xmlrpc:"function"`
	GroupsId                      []int64   `xmlrpc:"groups_id"`
	HasUnreconciledEntries        bool      `xmlrpc:"has_unreconciled_entries"`
	Id                            int64     `xmlrpc:"id"`
	ImStatus                      string    `xmlrpc:"im_status"`
	Image                         string    `xmlrpc:"image"`
	ImageMedium                   string    `xmlrpc:"image_medium"`
	ImageSmall                    string    `xmlrpc:"image_small"`
	InvoiceIds                    []int64   `xmlrpc:"invoice_ids"`
	InvoiceWarn                   string    `xmlrpc:"invoice_warn"`
	InvoiceWarnMsg                string    `xmlrpc:"invoice_warn_msg"`
	IsCompany                     bool      `xmlrpc:"is_company"`
	IssuedTotal                   float64   `xmlrpc:"issued_total"`
	JournalItemCount              int64     `xmlrpc:"journal_item_count"`
	Lang                          string    `xmlrpc:"lang"`
	LastTimeEntriesChecked        time.Time `xmlrpc:"last_time_entries_checked"`
	LogIds                        []int64   `xmlrpc:"log_ids"`
	Login                         string    `xmlrpc:"login"`
	LoginDate                     time.Time `xmlrpc:"login_date"`
	MachineCompanyName            string    `xmlrpc:"machine_company_name"`
	MachineUserEmail              string    `xmlrpc:"machine_user_email"`
	MachineUserLogin              string    `xmlrpc:"machine_user_login"`
	MeetingCount                  int64     `xmlrpc:"meeting_count"`
	MeetingIds                    []int64   `xmlrpc:"meeting_ids"`
	MessageBounce                 int64     `xmlrpc:"message_bounce"`
	MessageChannelIds             []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds            []int64   `xmlrpc:"message_follower_ids"`
	MessageIds                    []int64   `xmlrpc:"message_ids"`
	MessageIsFollower             bool      `xmlrpc:"message_is_follower"`
	MessageLastPost               time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction             bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter      int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds             []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread                 bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter          int64     `xmlrpc:"message_unread_counter"`
	Mobile                        string    `xmlrpc:"mobile"`
	Name                          string    `xmlrpc:"name"`
	NewPassword                   string    `xmlrpc:"new_password"`
	NotifyEmail                   string    `xmlrpc:"notify_email"`
	OpportunityCount              int64     `xmlrpc:"opportunity_count"`
	OpportunityIds                []int64   `xmlrpc:"opportunity_ids"`
	OptOut                        bool      `xmlrpc:"opt_out"`
	ParentId                      Many2One  `xmlrpc:"parent_id"`
	ParentName                    string    `xmlrpc:"parent_name"`
	PartnerId                     Many2One  `xmlrpc:"partner_id"`
	PartnerShare                  bool      `xmlrpc:"partner_share"`
	Password                      string    `xmlrpc:"password"`
	PasswordCrypt                 string    `xmlrpc:"password_crypt"`
	PaymentTokenCount             int64     `xmlrpc:"payment_token_count"`
	PaymentTokenIds               []int64   `xmlrpc:"payment_token_ids"`
	Phone                         string    `xmlrpc:"phone"`
	PickingWarn                   string    `xmlrpc:"picking_warn"`
	PickingWarnMsg                string    `xmlrpc:"picking_warn_msg"`
	PropertyAccountPayableId      Many2One  `xmlrpc:"property_account_payable_id"`
	PropertyAccountPositionId     Many2One  `xmlrpc:"property_account_position_id"`
	PropertyAccountReceivableId   Many2One  `xmlrpc:"property_account_receivable_id"`
	PropertyAutosalesConfig       Many2One  `xmlrpc:"property_autosales_config"`
	PropertyPaymentTermId         Many2One  `xmlrpc:"property_payment_term_id"`
	PropertyProductPricelist      Many2One  `xmlrpc:"property_product_pricelist"`
	PropertyPurchaseCurrencyId    Many2One  `xmlrpc:"property_purchase_currency_id"`
	PropertyStockCustomer         Many2One  `xmlrpc:"property_stock_customer"`
	PropertyStockSupplier         Many2One  `xmlrpc:"property_stock_supplier"`
	PropertySupplierPaymentTermId Many2One  `xmlrpc:"property_supplier_payment_term_id"`
	PurchaseOrderCount            int64     `xmlrpc:"purchase_order_count"`
	PurchaseWarn                  string    `xmlrpc:"purchase_warn"`
	PurchaseWarnMsg               string    `xmlrpc:"purchase_warn_msg"`
	Ref                           string    `xmlrpc:"ref"`
	RefCompanyIds                 []int64   `xmlrpc:"ref_company_ids"`
	SaleOrderCount                int64     `xmlrpc:"sale_order_count"`
	SaleOrderIds                  []int64   `xmlrpc:"sale_order_ids"`
	SaleTeamId                    Many2One  `xmlrpc:"sale_team_id"`
	SaleWarn                      string    `xmlrpc:"sale_warn"`
	SaleWarnMsg                   string    `xmlrpc:"sale_warn_msg"`
	Self                          Many2One  `xmlrpc:"self"`
	Share                         bool      `xmlrpc:"share"`
	Signature                     string    `xmlrpc:"signature"`
	SignupExpiration              time.Time `xmlrpc:"signup_expiration"`
	SignupToken                   string    `xmlrpc:"signup_token"`
	SignupType                    string    `xmlrpc:"signup_type"`
	SignupUrl                     string    `xmlrpc:"signup_url"`
	SignupValid                   bool      `xmlrpc:"signup_valid"`
	State                         string    `xmlrpc:"state"`
	StateId                       Many2One  `xmlrpc:"state_id"`
	Street                        string    `xmlrpc:"street"`
	Street2                       string    `xmlrpc:"street2"`
	Supplier                      bool      `xmlrpc:"supplier"`
	SupplierInvoiceCount          int64     `xmlrpc:"supplier_invoice_count"`
	TargetSalesDone               int64     `xmlrpc:"target_sales_done"`
	TargetSalesInvoiced           int64     `xmlrpc:"target_sales_invoiced"`
	TargetSalesWon                int64     `xmlrpc:"target_sales_won"`
	TaskCount                     int64     `xmlrpc:"task_count"`
	TaskIds                       []int64   `xmlrpc:"task_ids"`
	TeamId                        Many2One  `xmlrpc:"team_id"`
	Title                         Many2One  `xmlrpc:"title"`
	TotalInvoiced                 float64   `xmlrpc:"total_invoiced"`
	Trust                         string    `xmlrpc:"trust"`
	Type                          string    `xmlrpc:"type"`
	Tz                            string    `xmlrpc:"tz"`
	TzOffset                      string    `xmlrpc:"tz_offset"`
	UserId                        Many2One  `xmlrpc:"user_id"`
	UserIds                       []int64   `xmlrpc:"user_ids"`
	Vat                           string    `xmlrpc:"vat"`
	Website                       string    `xmlrpc:"website"`
	WriteDate                     time.Time `xmlrpc:"write_date"`
	WriteUid                      Many2One  `xmlrpc:"write_uid"`
	Zip                           string    `xmlrpc:"zip"`
}

type ResUsersNil struct {
	LastUpdate                    interface{} `xmlrpc:"__last_update"`
	ActionId                      interface{} `xmlrpc:"action_id"`
	Active                        bool        `xmlrpc:"active"`
	ActivitiesCount               interface{} `xmlrpc:"activities_count"`
	AliasContact                  interface{} `xmlrpc:"alias_contact"`
	AliasId                       interface{} `xmlrpc:"alias_id"`
	BankAccountCount              interface{} `xmlrpc:"bank_account_count"`
	BankIds                       interface{} `xmlrpc:"bank_ids"`
	Barcode                       interface{} `xmlrpc:"barcode"`
	CalendarLastNotifAck          interface{} `xmlrpc:"calendar_last_notif_ack"`
	CategoryId                    interface{} `xmlrpc:"category_id"`
	ChannelIds                    interface{} `xmlrpc:"channel_ids"`
	ChildIds                      interface{} `xmlrpc:"child_ids"`
	City                          interface{} `xmlrpc:"city"`
	Color                         interface{} `xmlrpc:"color"`
	Comment                       interface{} `xmlrpc:"comment"`
	CommercialCompanyName         interface{} `xmlrpc:"commercial_company_name"`
	CommercialPartnerId           interface{} `xmlrpc:"commercial_partner_id"`
	CompaniesCount                interface{} `xmlrpc:"companies_count"`
	CompanyId                     interface{} `xmlrpc:"company_id"`
	CompanyIds                    interface{} `xmlrpc:"company_ids"`
	CompanyName                   interface{} `xmlrpc:"company_name"`
	CompanyType                   interface{} `xmlrpc:"company_type"`
	ContactAddress                interface{} `xmlrpc:"contact_address"`
	ContractIds                   interface{} `xmlrpc:"contract_ids"`
	ContractsCount                interface{} `xmlrpc:"contracts_count"`
	CountryId                     interface{} `xmlrpc:"country_id"`
	CreateDate                    interface{} `xmlrpc:"create_date"`
	CreateUid                     interface{} `xmlrpc:"create_uid"`
	Credit                        interface{} `xmlrpc:"credit"`
	CreditLimit                   interface{} `xmlrpc:"credit_limit"`
	CurrencyId                    interface{} `xmlrpc:"currency_id"`
	Customer                      bool        `xmlrpc:"customer"`
	Date                          interface{} `xmlrpc:"date"`
	Debit                         interface{} `xmlrpc:"debit"`
	DebitLimit                    interface{} `xmlrpc:"debit_limit"`
	DisplayName                   interface{} `xmlrpc:"display_name"`
	Email                         interface{} `xmlrpc:"email"`
	EmailFormatted                interface{} `xmlrpc:"email_formatted"`
	Employee                      bool        `xmlrpc:"employee"`
	EmployeeIds                   interface{} `xmlrpc:"employee_ids"`
	Fax                           interface{} `xmlrpc:"fax"`
	Function                      interface{} `xmlrpc:"function"`
	GroupsId                      interface{} `xmlrpc:"groups_id"`
	HasUnreconciledEntries        bool        `xmlrpc:"has_unreconciled_entries"`
	Id                            interface{} `xmlrpc:"id"`
	ImStatus                      interface{} `xmlrpc:"im_status"`
	Image                         interface{} `xmlrpc:"image"`
	ImageMedium                   interface{} `xmlrpc:"image_medium"`
	ImageSmall                    interface{} `xmlrpc:"image_small"`
	InvoiceIds                    interface{} `xmlrpc:"invoice_ids"`
	InvoiceWarn                   interface{} `xmlrpc:"invoice_warn"`
	InvoiceWarnMsg                interface{} `xmlrpc:"invoice_warn_msg"`
	IsCompany                     bool        `xmlrpc:"is_company"`
	IssuedTotal                   interface{} `xmlrpc:"issued_total"`
	JournalItemCount              interface{} `xmlrpc:"journal_item_count"`
	Lang                          interface{} `xmlrpc:"lang"`
	LastTimeEntriesChecked        interface{} `xmlrpc:"last_time_entries_checked"`
	LogIds                        interface{} `xmlrpc:"log_ids"`
	Login                         interface{} `xmlrpc:"login"`
	LoginDate                     interface{} `xmlrpc:"login_date"`
	MachineCompanyName            interface{} `xmlrpc:"machine_company_name"`
	MachineUserEmail              interface{} `xmlrpc:"machine_user_email"`
	MachineUserLogin              interface{} `xmlrpc:"machine_user_login"`
	MeetingCount                  interface{} `xmlrpc:"meeting_count"`
	MeetingIds                    interface{} `xmlrpc:"meeting_ids"`
	MessageBounce                 interface{} `xmlrpc:"message_bounce"`
	MessageChannelIds             interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds            interface{} `xmlrpc:"message_follower_ids"`
	MessageIds                    interface{} `xmlrpc:"message_ids"`
	MessageIsFollower             bool        `xmlrpc:"message_is_follower"`
	MessageLastPost               interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction             bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter      interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds             interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread                 bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter          interface{} `xmlrpc:"message_unread_counter"`
	Mobile                        interface{} `xmlrpc:"mobile"`
	Name                          interface{} `xmlrpc:"name"`
	NewPassword                   interface{} `xmlrpc:"new_password"`
	NotifyEmail                   interface{} `xmlrpc:"notify_email"`
	OpportunityCount              interface{} `xmlrpc:"opportunity_count"`
	OpportunityIds                interface{} `xmlrpc:"opportunity_ids"`
	OptOut                        bool        `xmlrpc:"opt_out"`
	ParentId                      interface{} `xmlrpc:"parent_id"`
	ParentName                    interface{} `xmlrpc:"parent_name"`
	PartnerId                     interface{} `xmlrpc:"partner_id"`
	PartnerShare                  bool        `xmlrpc:"partner_share"`
	Password                      interface{} `xmlrpc:"password"`
	PasswordCrypt                 interface{} `xmlrpc:"password_crypt"`
	PaymentTokenCount             interface{} `xmlrpc:"payment_token_count"`
	PaymentTokenIds               interface{} `xmlrpc:"payment_token_ids"`
	Phone                         interface{} `xmlrpc:"phone"`
	PickingWarn                   interface{} `xmlrpc:"picking_warn"`
	PickingWarnMsg                interface{} `xmlrpc:"picking_warn_msg"`
	PropertyAccountPayableId      interface{} `xmlrpc:"property_account_payable_id"`
	PropertyAccountPositionId     interface{} `xmlrpc:"property_account_position_id"`
	PropertyAccountReceivableId   interface{} `xmlrpc:"property_account_receivable_id"`
	PropertyAutosalesConfig       interface{} `xmlrpc:"property_autosales_config"`
	PropertyPaymentTermId         interface{} `xmlrpc:"property_payment_term_id"`
	PropertyProductPricelist      interface{} `xmlrpc:"property_product_pricelist"`
	PropertyPurchaseCurrencyId    interface{} `xmlrpc:"property_purchase_currency_id"`
	PropertyStockCustomer         interface{} `xmlrpc:"property_stock_customer"`
	PropertyStockSupplier         interface{} `xmlrpc:"property_stock_supplier"`
	PropertySupplierPaymentTermId interface{} `xmlrpc:"property_supplier_payment_term_id"`
	PurchaseOrderCount            interface{} `xmlrpc:"purchase_order_count"`
	PurchaseWarn                  interface{} `xmlrpc:"purchase_warn"`
	PurchaseWarnMsg               interface{} `xmlrpc:"purchase_warn_msg"`
	Ref                           interface{} `xmlrpc:"ref"`
	RefCompanyIds                 interface{} `xmlrpc:"ref_company_ids"`
	SaleOrderCount                interface{} `xmlrpc:"sale_order_count"`
	SaleOrderIds                  interface{} `xmlrpc:"sale_order_ids"`
	SaleTeamId                    interface{} `xmlrpc:"sale_team_id"`
	SaleWarn                      interface{} `xmlrpc:"sale_warn"`
	SaleWarnMsg                   interface{} `xmlrpc:"sale_warn_msg"`
	Self                          interface{} `xmlrpc:"self"`
	Share                         bool        `xmlrpc:"share"`
	Signature                     interface{} `xmlrpc:"signature"`
	SignupExpiration              interface{} `xmlrpc:"signup_expiration"`
	SignupToken                   interface{} `xmlrpc:"signup_token"`
	SignupType                    interface{} `xmlrpc:"signup_type"`
	SignupUrl                     interface{} `xmlrpc:"signup_url"`
	SignupValid                   bool        `xmlrpc:"signup_valid"`
	State                         interface{} `xmlrpc:"state"`
	StateId                       interface{} `xmlrpc:"state_id"`
	Street                        interface{} `xmlrpc:"street"`
	Street2                       interface{} `xmlrpc:"street2"`
	Supplier                      bool        `xmlrpc:"supplier"`
	SupplierInvoiceCount          interface{} `xmlrpc:"supplier_invoice_count"`
	TargetSalesDone               interface{} `xmlrpc:"target_sales_done"`
	TargetSalesInvoiced           interface{} `xmlrpc:"target_sales_invoiced"`
	TargetSalesWon                interface{} `xmlrpc:"target_sales_won"`
	TaskCount                     interface{} `xmlrpc:"task_count"`
	TaskIds                       interface{} `xmlrpc:"task_ids"`
	TeamId                        interface{} `xmlrpc:"team_id"`
	Title                         interface{} `xmlrpc:"title"`
	TotalInvoiced                 interface{} `xmlrpc:"total_invoiced"`
	Trust                         interface{} `xmlrpc:"trust"`
	Type                          interface{} `xmlrpc:"type"`
	Tz                            interface{} `xmlrpc:"tz"`
	TzOffset                      interface{} `xmlrpc:"tz_offset"`
	UserId                        interface{} `xmlrpc:"user_id"`
	UserIds                       interface{} `xmlrpc:"user_ids"`
	Vat                           interface{} `xmlrpc:"vat"`
	Website                       interface{} `xmlrpc:"website"`
	WriteDate                     interface{} `xmlrpc:"write_date"`
	WriteUid                      interface{} `xmlrpc:"write_uid"`
	Zip                           interface{} `xmlrpc:"zip"`
}

var ResUsersModel string = "res.users"

type ResUserss []ResUsers

type ResUserssNil []ResUsersNil

func (s *ResUsers) NilableType_() interface{} {
	return &ResUsersNil{}
}

func (ns *ResUsersNil) Type_() interface{} {
	s := &ResUsers{}
	return load(ns, s)
}

func (s *ResUserss) NilableType_() interface{} {
	return &ResUserssNil{}
}

func (ns *ResUserssNil) Type_() interface{} {
	s := &ResUserss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResUsers))
	}
	return s
}
