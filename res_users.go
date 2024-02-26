package odoo

// ResUsers represents res.users model.
type ResUsers struct {
	LastUpdate                    *Time      `xmlrpc:"__last_update,omptempty"`
	ActionId                      *Many2One  `xmlrpc:"action_id,omptempty"`
	Active                        *Bool      `xmlrpc:"active,omptempty"`
	ActivityDateDeadline          *Time      `xmlrpc:"activity_date_deadline,omptempty"`
	ActivityIds                   *Relation  `xmlrpc:"activity_ids,omptempty"`
	ActivityState                 *Selection `xmlrpc:"activity_state,omptempty"`
	ActivitySummary               *String    `xmlrpc:"activity_summary,omptempty"`
	ActivityTypeId                *Many2One  `xmlrpc:"activity_type_id,omptempty"`
	ActivityUserId                *Many2One  `xmlrpc:"activity_user_id,omptempty"`
	AliasContact                  *Selection `xmlrpc:"alias_contact,omptempty"`
	AliasId                       *Many2One  `xmlrpc:"alias_id,omptempty"`
	BankAccountCount              *Int       `xmlrpc:"bank_account_count,omptempty"`
	BankIds                       *Relation  `xmlrpc:"bank_ids,omptempty"`
	Barcode                       *String    `xmlrpc:"barcode,omptempty"`
	CalendarLastNotifAck          *Time      `xmlrpc:"calendar_last_notif_ack,omptempty"`
	CategoryId                    *Relation  `xmlrpc:"category_id,omptempty"`
	ChannelIds                    *Relation  `xmlrpc:"channel_ids,omptempty"`
	ChildIds                      *Relation  `xmlrpc:"child_ids,omptempty"`
	City                          *String    `xmlrpc:"city,omptempty"`
	Color                         *Int       `xmlrpc:"color,omptempty"`
	Comment                       *String    `xmlrpc:"comment,omptempty"`
	CommercialCompanyName         *String    `xmlrpc:"commercial_company_name,omptempty"`
	CommercialPartnerCountryId    *Many2One  `xmlrpc:"commercial_partner_country_id,omptempty"`
	CommercialPartnerId           *Many2One  `xmlrpc:"commercial_partner_id,omptempty"`
	CompaniesCount                *Int       `xmlrpc:"companies_count,omptempty"`
	CompanyId                     *Many2One  `xmlrpc:"company_id,omptempty"`
	CompanyIds                    *Relation  `xmlrpc:"company_ids,omptempty"`
	CompanyName                   *String    `xmlrpc:"company_name,omptempty"`
	CompanyType                   *Selection `xmlrpc:"company_type,omptempty"`
	ContactAddress                *String    `xmlrpc:"contact_address,omptempty"`
	ContractIds                   *Relation  `xmlrpc:"contract_ids,omptempty"`
	ContractsCount                *Int       `xmlrpc:"contracts_count,omptempty"`
	CountryId                     *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate                    *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                     *Many2One  `xmlrpc:"create_uid,omptempty"`
	Credit                        *Float     `xmlrpc:"credit,omptempty"`
	CreditLimit                   *Float     `xmlrpc:"credit_limit,omptempty"`
	CurrencyId                    *Many2One  `xmlrpc:"currency_id,omptempty"`
	Customer                      *Bool      `xmlrpc:"customer,omptempty"`
	Date                          *Time      `xmlrpc:"date,omptempty"`
	Debit                         *Float     `xmlrpc:"debit,omptempty"`
	DebitLimit                    *Float     `xmlrpc:"debit_limit,omptempty"`
	DisplayName                   *String    `xmlrpc:"display_name,omptempty"`
	Email                         *String    `xmlrpc:"email,omptempty"`
	EmailFormatted                *String    `xmlrpc:"email_formatted,omptempty"`
	Employee                      *Bool      `xmlrpc:"employee,omptempty"`
	EmployeeIds                   *Relation  `xmlrpc:"employee_ids,omptempty"`
	Function                      *String    `xmlrpc:"function,omptempty"`
	GroupsId                      *Relation  `xmlrpc:"groups_id,omptempty"`
	HasUnreconciledEntries        *Bool      `xmlrpc:"has_unreconciled_entries,omptempty"`
	Id                            *Int       `xmlrpc:"id,omptempty"`
	ImStatus                      *String    `xmlrpc:"im_status,omptempty"`
	Image                         *String    `xmlrpc:"image,omptempty"`
	ImageMedium                   *String    `xmlrpc:"image_medium,omptempty"`
	ImageSmall                    *String    `xmlrpc:"image_small,omptempty"`
	IndustryId                    *Many2One  `xmlrpc:"industry_id,omptempty"`
	InvoiceIds                    *Relation  `xmlrpc:"invoice_ids,omptempty"`
	InvoiceWarn                   *Selection `xmlrpc:"invoice_warn,omptempty"`
	InvoiceWarnMsg                *String    `xmlrpc:"invoice_warn_msg,omptempty"`
	IsCompany                     *Bool      `xmlrpc:"is_company,omptempty"`
	JournalItemCount              *Int       `xmlrpc:"journal_item_count,omptempty"`
	Lang                          *Selection `xmlrpc:"lang,omptempty"`
	LastTimeEntriesChecked        *Time      `xmlrpc:"last_time_entries_checked,omptempty"`
	LogIds                        *Relation  `xmlrpc:"log_ids,omptempty"`
	Login                         *String    `xmlrpc:"login,omptempty"`
	LoginDate                     *Time      `xmlrpc:"login_date,omptempty"`
	MachineCompanyName            *String    `xmlrpc:"machine_company_name,omptempty"`
	MachineOrganizationName       *String    `xmlrpc:"machine_organization_name,omptempty"`
	MachineUserEmail              *String    `xmlrpc:"machine_user_email,omptempty"`
	MachineUserLogin              *String    `xmlrpc:"machine_user_login,omptempty"`
	MeetingCount                  *Int       `xmlrpc:"meeting_count,omptempty"`
	MeetingIds                    *Relation  `xmlrpc:"meeting_ids,omptempty"`
	MessageBounce                 *Int       `xmlrpc:"message_bounce,omptempty"`
	MessageChannelIds             *Relation  `xmlrpc:"message_channel_ids,omptempty"`
	MessageFollowerIds            *Relation  `xmlrpc:"message_follower_ids,omptempty"`
	MessageIds                    *Relation  `xmlrpc:"message_ids,omptempty"`
	MessageIsFollower             *Bool      `xmlrpc:"message_is_follower,omptempty"`
	MessageLastPost               *Time      `xmlrpc:"message_last_post,omptempty"`
	MessageNeedaction             *Bool      `xmlrpc:"message_needaction,omptempty"`
	MessageNeedactionCounter      *Int       `xmlrpc:"message_needaction_counter,omptempty"`
	MessagePartnerIds             *Relation  `xmlrpc:"message_partner_ids,omptempty"`
	MessageUnread                 *Bool      `xmlrpc:"message_unread,omptempty"`
	MessageUnreadCounter          *Int       `xmlrpc:"message_unread_counter,omptempty"`
	Mobile                        *String    `xmlrpc:"mobile,omptempty"`
	Name                          *String    `xmlrpc:"name,omptempty"`
	NewPassword                   *String    `xmlrpc:"new_password,omptempty"`
	NotificationType              *Selection `xmlrpc:"notification_type,omptempty"`
	OpportunityCount              *Int       `xmlrpc:"opportunity_count,omptempty"`
	OpportunityIds                *Relation  `xmlrpc:"opportunity_ids,omptempty"`
	OptOut                        *Bool      `xmlrpc:"opt_out,omptempty"`
	ParentId                      *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentName                    *String    `xmlrpc:"parent_name,omptempty"`
	PartnerId                     *Many2One  `xmlrpc:"partner_id,omptempty"`
	PartnerShare                  *Bool      `xmlrpc:"partner_share,omptempty"`
	Password                      *String    `xmlrpc:"password,omptempty"`
	PasswordCrypt                 *String    `xmlrpc:"password_crypt,omptempty"`
	PaymentTokenCount             *Int       `xmlrpc:"payment_token_count,omptempty"`
	PaymentTokenIds               *Relation  `xmlrpc:"payment_token_ids,omptempty"`
	Phone                         *String    `xmlrpc:"phone,omptempty"`
	PickingWarn                   *Selection `xmlrpc:"picking_warn,omptempty"`
	PickingWarnMsg                *String    `xmlrpc:"picking_warn_msg,omptempty"`
	PropertyAccountPayableId      *Many2One  `xmlrpc:"property_account_payable_id,omptempty"`
	PropertyAccountPositionId     *Many2One  `xmlrpc:"property_account_position_id,omptempty"`
	PropertyAccountReceivableId   *Many2One  `xmlrpc:"property_account_receivable_id,omptempty"`
	PropertyAutosalesConfig       *Many2One  `xmlrpc:"property_autosales_config,omptempty"`
	PropertyPaymentTermId         *Many2One  `xmlrpc:"property_payment_term_id,omptempty"`
	PropertyProductPricelist      *Many2One  `xmlrpc:"property_product_pricelist,omptempty"`
	PropertyPurchaseCurrencyId    *Many2One  `xmlrpc:"property_purchase_currency_id,omptempty"`
	PropertyStockCustomer         *Many2One  `xmlrpc:"property_stock_customer,omptempty"`
	PropertyStockSupplier         *Many2One  `xmlrpc:"property_stock_supplier,omptempty"`
	PropertySupplierPaymentTermId *Many2One  `xmlrpc:"property_supplier_payment_term_id,omptempty"`
	PurchaseOrderCount            *Int       `xmlrpc:"purchase_order_count,omptempty"`
	PurchaseWarn                  *Selection `xmlrpc:"purchase_warn,omptempty"`
	PurchaseWarnMsg               *String    `xmlrpc:"purchase_warn_msg,omptempty"`
	Ref                           *String    `xmlrpc:"ref,omptempty"`
	RefCompanyIds                 *Relation  `xmlrpc:"ref_company_ids,omptempty"`
	ResourceCalendarId            *Many2One  `xmlrpc:"resource_calendar_id,omptempty"`
	ResourceIds                   *Relation  `xmlrpc:"resource_ids,omptempty"`
	SaleOrderCount                *Int       `xmlrpc:"sale_order_count,omptempty"`
	SaleOrderIds                  *Relation  `xmlrpc:"sale_order_ids,omptempty"`
	SaleTeamId                    *Many2One  `xmlrpc:"sale_team_id,omptempty"`
	SaleWarn                      *Selection `xmlrpc:"sale_warn,omptempty"`
	SaleWarnMsg                   *String    `xmlrpc:"sale_warn_msg,omptempty"`
	Self                          *Many2One  `xmlrpc:"self,omptempty"`
	Share                         *Bool      `xmlrpc:"share,omptempty"`
	Signature                     *String    `xmlrpc:"signature,omptempty"`
	Siret                         *String    `xmlrpc:"siret,omptempty"`
	StateId                       *Many2One  `xmlrpc:"state_id,omptempty"`
	Street                        *String    `xmlrpc:"street,omptempty"`
	Street2                       *String    `xmlrpc:"street2,omptempty"`
	Supplier                      *Bool      `xmlrpc:"supplier,omptempty"`
	SupplierInvoiceCount          *Int       `xmlrpc:"supplier_invoice_count,omptempty"`
	TargetSalesDone               *Int       `xmlrpc:"target_sales_done,omptempty"`
	TargetSalesInvoiced           *Int       `xmlrpc:"target_sales_invoiced,omptempty"`
	TargetSalesWon                *Int       `xmlrpc:"target_sales_won,omptempty"`
	TaskCount                     *Int       `xmlrpc:"task_count,omptempty"`
	TaskIds                       *Relation  `xmlrpc:"task_ids,omptempty"`
	TeamId                        *Many2One  `xmlrpc:"team_id,omptempty"`
	Title                         *Many2One  `xmlrpc:"title,omptempty"`
	TotalInvoiced                 *Float     `xmlrpc:"total_invoiced,omptempty"`
	Trust                         *Selection `xmlrpc:"trust,omptempty"`
	Type                          *Selection `xmlrpc:"type,omptempty"`
	Tz                            *Selection `xmlrpc:"tz,omptempty"`
	TzOffset                      *String    `xmlrpc:"tz_offset,omptempty"`
	UserId                        *Many2One  `xmlrpc:"user_id,omptempty"`
	UserIds                       *Relation  `xmlrpc:"user_ids,omptempty"`
	Vat                           *String    `xmlrpc:"vat,omptempty"`
	Website                       *String    `xmlrpc:"website,omptempty"`
	WebsiteMessageIds             *Relation  `xmlrpc:"website_message_ids,omptempty"`
	WriteDate                     *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                      *Many2One  `xmlrpc:"write_uid,omptempty"`
	Zip                           *String    `xmlrpc:"zip,omptempty"`
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
