package types

import (
	"time"
)

type ResPartner struct {
	LastUpdate                    time.Time `xmlrpc:"__last_update"`
	Active                        bool      `xmlrpc:"active"`
	ActivitiesCount               int64     `xmlrpc:"activities_count"`
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
	CompanyId                     Many2One  `xmlrpc:"company_id"`
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
	Fax                           string    `xmlrpc:"fax"`
	Function                      string    `xmlrpc:"function"`
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
	MachineCompanyName            string    `xmlrpc:"machine_company_name"`
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
	NotifyEmail                   string    `xmlrpc:"notify_email"`
	OpportunityCount              int64     `xmlrpc:"opportunity_count"`
	OpportunityIds                []int64   `xmlrpc:"opportunity_ids"`
	OptOut                        bool      `xmlrpc:"opt_out"`
	ParentId                      Many2One  `xmlrpc:"parent_id"`
	ParentName                    string    `xmlrpc:"parent_name"`
	PartnerShare                  bool      `xmlrpc:"partner_share"`
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
	SaleWarn                      string    `xmlrpc:"sale_warn"`
	SaleWarnMsg                   string    `xmlrpc:"sale_warn_msg"`
	Self                          Many2One  `xmlrpc:"self"`
	SignupExpiration              time.Time `xmlrpc:"signup_expiration"`
	SignupToken                   string    `xmlrpc:"signup_token"`
	SignupType                    string    `xmlrpc:"signup_type"`
	SignupUrl                     string    `xmlrpc:"signup_url"`
	SignupValid                   bool      `xmlrpc:"signup_valid"`
	StateId                       Many2One  `xmlrpc:"state_id"`
	Street                        string    `xmlrpc:"street"`
	Street2                       string    `xmlrpc:"street2"`
	Supplier                      bool      `xmlrpc:"supplier"`
	SupplierInvoiceCount          int64     `xmlrpc:"supplier_invoice_count"`
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

type ResPartnerNil struct {
	LastUpdate                    interface{} `xmlrpc:"__last_update"`
	Active                        bool        `xmlrpc:"active"`
	ActivitiesCount               interface{} `xmlrpc:"activities_count"`
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
	CompanyId                     interface{} `xmlrpc:"company_id"`
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
	Fax                           interface{} `xmlrpc:"fax"`
	Function                      interface{} `xmlrpc:"function"`
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
	MachineCompanyName            interface{} `xmlrpc:"machine_company_name"`
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
	NotifyEmail                   interface{} `xmlrpc:"notify_email"`
	OpportunityCount              interface{} `xmlrpc:"opportunity_count"`
	OpportunityIds                interface{} `xmlrpc:"opportunity_ids"`
	OptOut                        bool        `xmlrpc:"opt_out"`
	ParentId                      interface{} `xmlrpc:"parent_id"`
	ParentName                    interface{} `xmlrpc:"parent_name"`
	PartnerShare                  bool        `xmlrpc:"partner_share"`
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
	SaleWarn                      interface{} `xmlrpc:"sale_warn"`
	SaleWarnMsg                   interface{} `xmlrpc:"sale_warn_msg"`
	Self                          interface{} `xmlrpc:"self"`
	SignupExpiration              interface{} `xmlrpc:"signup_expiration"`
	SignupToken                   interface{} `xmlrpc:"signup_token"`
	SignupType                    interface{} `xmlrpc:"signup_type"`
	SignupUrl                     interface{} `xmlrpc:"signup_url"`
	SignupValid                   bool        `xmlrpc:"signup_valid"`
	StateId                       interface{} `xmlrpc:"state_id"`
	Street                        interface{} `xmlrpc:"street"`
	Street2                       interface{} `xmlrpc:"street2"`
	Supplier                      bool        `xmlrpc:"supplier"`
	SupplierInvoiceCount          interface{} `xmlrpc:"supplier_invoice_count"`
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

var ResPartnerModel string = "res.partner"

type ResPartners []ResPartner

type ResPartnersNil []ResPartnerNil

func (s *ResPartner) NilableType_() interface{} {
	return &ResPartnerNil{}
}

func (ns *ResPartnerNil) Type_() interface{} {
	s := &ResPartner{}
	return load(ns, s)
}

func (s *ResPartners) NilableType_() interface{} {
	return &ResPartnersNil{}
}

func (ns *ResPartnersNil) Type_() interface{} {
	s := &ResPartners{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ResPartner))
	}
	return s
}
