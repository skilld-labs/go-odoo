package odoo

// AccountInvoiceLine represents account.invoice.line model.
type AccountInvoiceLine struct {
	LastUpdate             *Time      `xmlrpc:"__last_update,omptempty"`
	AccountAnalyticId      *Many2One  `xmlrpc:"account_analytic_id,omptempty"`
	AccountId              *Many2One  `xmlrpc:"account_id,omptempty"`
	AnalyticTagIds         *Relation  `xmlrpc:"analytic_tag_ids,omptempty"`
	CompanyCurrencyId      *Many2One  `xmlrpc:"company_currency_id,omptempty"`
	CompanyId              *Many2One  `xmlrpc:"company_id,omptempty"`
	CreateDate             *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid              *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId             *Many2One  `xmlrpc:"currency_id,omptempty"`
	Discount               *Float     `xmlrpc:"discount,omptempty"`
	DisplayName            *String    `xmlrpc:"display_name,omptempty"`
	Id                     *Int       `xmlrpc:"id,omptempty"`
	InvoiceId              *Many2One  `xmlrpc:"invoice_id,omptempty"`
	InvoiceLineTaxIds      *Relation  `xmlrpc:"invoice_line_tax_ids,omptempty"`
	InvoiceType            *Selection `xmlrpc:"invoice_type,omptempty"`
	IsRoundingLine         *Bool      `xmlrpc:"is_rounding_line,omptempty"`
	LayoutCategoryId       *Many2One  `xmlrpc:"layout_category_id,omptempty"`
	LayoutCategorySequence *Int       `xmlrpc:"layout_category_sequence,omptempty"`
	Name                   *String    `xmlrpc:"name,omptempty"`
	Origin                 *String    `xmlrpc:"origin,omptempty"`
	PartnerId              *Many2One  `xmlrpc:"partner_id,omptempty"`
	PriceSubtotal          *Float     `xmlrpc:"price_subtotal,omptempty"`
	PriceSubtotalSigned    *Float     `xmlrpc:"price_subtotal_signed,omptempty"`
	PriceTotal             *Float     `xmlrpc:"price_total,omptempty"`
	PriceUnit              *Float     `xmlrpc:"price_unit,omptempty"`
	ProductId              *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductImage           *String    `xmlrpc:"product_image,omptempty"`
	PurchaseId             *Many2One  `xmlrpc:"purchase_id,omptempty"`
	PurchaseLineId         *Many2One  `xmlrpc:"purchase_line_id,omptempty"`
	Quantity               *Float     `xmlrpc:"quantity,omptempty"`
	SaleLineIds            *Relation  `xmlrpc:"sale_line_ids,omptempty"`
	Sequence               *Int       `xmlrpc:"sequence,omptempty"`
	UomId                  *Many2One  `xmlrpc:"uom_id,omptempty"`
	WriteDate              *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid               *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// AccountInvoiceLines represents array of account.invoice.line model.
type AccountInvoiceLines []AccountInvoiceLine

// AccountInvoiceLineModel is the odoo model name.
const AccountInvoiceLineModel = "account.invoice.line"

// Many2One convert AccountInvoiceLine to *Many2One.
func (ail *AccountInvoiceLine) Many2One() *Many2One {
	return NewMany2One(ail.Id.Get(), "")
}

// CreateAccountInvoiceLine creates a new account.invoice.line model and returns its id.
func (c *Client) CreateAccountInvoiceLine(ail *AccountInvoiceLine) (int64, error) {
	ids, err := c.CreateAccountInvoiceLines([]*AccountInvoiceLine{ail})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAccountInvoiceLine creates a new account.invoice.line model and returns its id.
func (c *Client) CreateAccountInvoiceLines(ails []*AccountInvoiceLine) ([]int64, error) {
	var vv []interface{}
	for _, v := range ails {
		vv = append(vv, v)
	}
	return c.Create(AccountInvoiceLineModel, vv, nil)
}

// UpdateAccountInvoiceLine updates an existing account.invoice.line record.
func (c *Client) UpdateAccountInvoiceLine(ail *AccountInvoiceLine) error {
	return c.UpdateAccountInvoiceLines([]int64{ail.Id.Get()}, ail)
}

// UpdateAccountInvoiceLines updates existing account.invoice.line records.
// All records (represented by ids) will be updated by ail values.
func (c *Client) UpdateAccountInvoiceLines(ids []int64, ail *AccountInvoiceLine) error {
	return c.Update(AccountInvoiceLineModel, ids, ail, nil)
}

// DeleteAccountInvoiceLine deletes an existing account.invoice.line record.
func (c *Client) DeleteAccountInvoiceLine(id int64) error {
	return c.DeleteAccountInvoiceLines([]int64{id})
}

// DeleteAccountInvoiceLines deletes existing account.invoice.line records.
func (c *Client) DeleteAccountInvoiceLines(ids []int64) error {
	return c.Delete(AccountInvoiceLineModel, ids)
}

// GetAccountInvoiceLine gets account.invoice.line existing record.
func (c *Client) GetAccountInvoiceLine(id int64) (*AccountInvoiceLine, error) {
	ails, err := c.GetAccountInvoiceLines([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ails)[0]), nil
}

// GetAccountInvoiceLines gets account.invoice.line existing records.
func (c *Client) GetAccountInvoiceLines(ids []int64) (*AccountInvoiceLines, error) {
	ails := &AccountInvoiceLines{}
	if err := c.Read(AccountInvoiceLineModel, ids, nil, ails); err != nil {
		return nil, err
	}
	return ails, nil
}

// FindAccountInvoiceLine finds account.invoice.line record by querying it with criteria.
func (c *Client) FindAccountInvoiceLine(criteria *Criteria) (*AccountInvoiceLine, error) {
	ails := &AccountInvoiceLines{}
	if err := c.SearchRead(AccountInvoiceLineModel, criteria, NewOptions().Limit(1), ails); err != nil {
		return nil, err
	}
	return &((*ails)[0]), nil
}

// FindAccountInvoiceLines finds account.invoice.line records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceLines(criteria *Criteria, options *Options) (*AccountInvoiceLines, error) {
	ails := &AccountInvoiceLines{}
	if err := c.SearchRead(AccountInvoiceLineModel, criteria, options, ails); err != nil {
		return nil, err
	}
	return ails, nil
}

// FindAccountInvoiceLineIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAccountInvoiceLineIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AccountInvoiceLineModel, criteria, options)
}

// FindAccountInvoiceLineId finds record id by querying it with criteria.
func (c *Client) FindAccountInvoiceLineId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AccountInvoiceLineModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
