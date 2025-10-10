package odoo

// AccountInvoiceReport represents account.invoice.report model.
type AccountInvoiceReport struct {
	AccountId             *Many2One  `xmlrpc:"account_id,omitempty"`
	CommercialPartnerId   *Many2One  `xmlrpc:"commercial_partner_id,omitempty"`
	CompanyCurrencyId     *Many2One  `xmlrpc:"company_currency_id,omitempty"`
	CompanyId             *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId             *Many2One  `xmlrpc:"country_id,omitempty"`
	CurrencyId            *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName           *String    `xmlrpc:"display_name,omitempty"`
	FiscalPositionId      *Many2One  `xmlrpc:"fiscal_position_id,omitempty"`
	Id                    *Int       `xmlrpc:"id,omitempty"`
	InventoryValue        *Float     `xmlrpc:"inventory_value,omitempty"`
	InvoiceDate           *Time      `xmlrpc:"invoice_date,omitempty"`
	InvoiceDateDue        *Time      `xmlrpc:"invoice_date_due,omitempty"`
	InvoiceUserId         *Many2One  `xmlrpc:"invoice_user_id,omitempty"`
	JournalId             *Many2One  `xmlrpc:"journal_id,omitempty"`
	MoveId                *Many2One  `xmlrpc:"move_id,omitempty"`
	MoveType              *Selection `xmlrpc:"move_type,omitempty"`
	PartnerId             *Many2One  `xmlrpc:"partner_id,omitempty"`
	PaymentState          *Selection `xmlrpc:"payment_state,omitempty"`
	PriceAverage          *Float     `xmlrpc:"price_average,omitempty"`
	PriceMargin           *Float     `xmlrpc:"price_margin,omitempty"`
	PriceSubtotal         *Float     `xmlrpc:"price_subtotal,omitempty"`
	PriceSubtotalCurrency *Float     `xmlrpc:"price_subtotal_currency,omitempty"`
	PriceTotal            *Float     `xmlrpc:"price_total,omitempty"`
	ProductCategId        *Many2One  `xmlrpc:"product_categ_id,omitempty"`
	ProductId             *Many2One  `xmlrpc:"product_id,omitempty"`
	ProductUomId          *Many2One  `xmlrpc:"product_uom_id,omitempty"`
	Quantity              *Float     `xmlrpc:"quantity,omitempty"`
	State                 *Selection `xmlrpc:"state,omitempty"`
	TeamId                *Many2One  `xmlrpc:"team_id,omitempty"`
}

// AccountInvoiceReports represents array of account.invoice.report model.
type AccountInvoiceReports []AccountInvoiceReport

// AccountInvoiceReportModel is the odoo model name.
const AccountInvoiceReportModel = "account.invoice.report"

// Many2One convert AccountInvoiceReport to *Many2One.
func (air *AccountInvoiceReport) Many2One() *Many2One {
	return NewMany2One(air.Id.Get(), "")
}

// CreateAccountInvoiceReport creates a new account.invoice.report model and returns its id.
func (c *Client) CreateAccountInvoiceReport(air *AccountInvoiceReport) (int64, error) {
	ids, err := c.CreateAccountInvoiceReports([]*AccountInvoiceReport{air})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountInvoiceReport creates a new account.invoice.report model and returns its id.
func (c *Client) CreateAccountInvoiceReports(airs []*AccountInvoiceReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range airs {
		vv = append(vv, v)
	}
	return c.Create(AccountInvoiceReportModel, vv, nil)
}

// UpdateAccountInvoiceReport updates an existing account.invoice.report record.
func (c *Client) UpdateAccountInvoiceReport(air *AccountInvoiceReport) error {
	return c.UpdateAccountInvoiceReports([]int64{air.Id.Get()}, air)
}

// UpdateAccountInvoiceReports updates existing account.invoice.report records.
// All records (represented by ids) will be updated by air values.
func (c *Client) UpdateAccountInvoiceReports(ids []int64, air *AccountInvoiceReport) error {
	return c.Update(AccountInvoiceReportModel, ids, air, nil)
}

// DeleteAccountInvoiceReport deletes an existing account.invoice.report record.
func (c *Client) DeleteAccountInvoiceReport(id int64) error {
	return c.DeleteAccountInvoiceReports([]int64{id})
}

// DeleteAccountInvoiceReports deletes existing account.invoice.report records.
func (c *Client) DeleteAccountInvoiceReports(ids []int64) error {
	return c.Delete(AccountInvoiceReportModel, ids)
}

// GetAccountInvoiceReport gets account.invoice.report existing record.
func (c *Client) GetAccountInvoiceReport(id int64) (*AccountInvoiceReport, error) {
	airs, err := c.GetAccountInvoiceReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*airs)[0]), nil
}

// GetAccountInvoiceReports gets account.invoice.report existing records.
func (c *Client) GetAccountInvoiceReports(ids []int64) (*AccountInvoiceReports, error) {
	airs := &AccountInvoiceReports{}
	if err := c.Read(AccountInvoiceReportModel, ids, nil, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountInvoiceReport finds account.invoice.report record by querying it with criteria.
func (c *Client) FindAccountInvoiceReport(criteria *Criteria) (*AccountInvoiceReport, error) {
	airs := &AccountInvoiceReports{}
	if err := c.SearchRead(AccountInvoiceReportModel, criteria, NewOptions().Limit(1), airs); err != nil {
		return nil, err
	}
	return &((*airs)[0]), nil
}

// FindAccountInvoiceReports finds account.invoice.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceReports(criteria *Criteria, options *Options) (*AccountInvoiceReports, error) {
	airs := &AccountInvoiceReports{}
	if err := c.SearchRead(AccountInvoiceReportModel, criteria, options, airs); err != nil {
		return nil, err
	}
	return airs, nil
}

// FindAccountInvoiceReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountInvoiceReportModel, criteria, options)
}

// FindAccountInvoiceReportId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
