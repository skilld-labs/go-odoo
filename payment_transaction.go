package odoo

// PaymentTransaction represents payment.transaction model.
type PaymentTransaction struct {
	Amount              *Float     `xmlrpc:"amount,omitempty"`
	ChildTransactionIds *Relation  `xmlrpc:"child_transaction_ids,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId          *Many2One  `xmlrpc:"currency_id,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	InvoiceIds          *Relation  `xmlrpc:"invoice_ids,omitempty"`
	InvoicesCount       *Int       `xmlrpc:"invoices_count,omitempty"`
	IsPostProcessed     *Bool      `xmlrpc:"is_post_processed,omitempty"`
	LandingRoute        *String    `xmlrpc:"landing_route,omitempty"`
	LastStateChange     *Time      `xmlrpc:"last_state_change,omitempty"`
	Operation           *Selection `xmlrpc:"operation,omitempty"`
	PartnerAddress      *String    `xmlrpc:"partner_address,omitempty"`
	PartnerCity         *String    `xmlrpc:"partner_city,omitempty"`
	PartnerCountryId    *Many2One  `xmlrpc:"partner_country_id,omitempty"`
	PartnerEmail        *String    `xmlrpc:"partner_email,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	PartnerLang         *Selection `xmlrpc:"partner_lang,omitempty"`
	PartnerName         *String    `xmlrpc:"partner_name,omitempty"`
	PartnerPhone        *String    `xmlrpc:"partner_phone,omitempty"`
	PartnerStateId      *Many2One  `xmlrpc:"partner_state_id,omitempty"`
	PartnerZip          *String    `xmlrpc:"partner_zip,omitempty"`
	PaymentId           *Many2One  `xmlrpc:"payment_id,omitempty"`
	PaymentMethodCode   *String    `xmlrpc:"payment_method_code,omitempty"`
	PaymentMethodId     *Many2One  `xmlrpc:"payment_method_id,omitempty"`
	ProviderCode        *Selection `xmlrpc:"provider_code,omitempty"`
	ProviderId          *Many2One  `xmlrpc:"provider_id,omitempty"`
	ProviderReference   *String    `xmlrpc:"provider_reference,omitempty"`
	Reference           *String    `xmlrpc:"reference,omitempty"`
	RefundsCount        *Int       `xmlrpc:"refunds_count,omitempty"`
	SaleOrderIds        *Relation  `xmlrpc:"sale_order_ids,omitempty"`
	SaleOrderIdsNbr     *Int       `xmlrpc:"sale_order_ids_nbr,omitempty"`
	SourceTransactionId *Many2One  `xmlrpc:"source_transaction_id,omitempty"`
	State               *Selection `xmlrpc:"state,omitempty"`
	StateMessage        *String    `xmlrpc:"state_message,omitempty"`
	TokenId             *Many2One  `xmlrpc:"token_id,omitempty"`
	Tokenize            *Bool      `xmlrpc:"tokenize,omitempty"`
	WriteDate           *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// PaymentTransactions represents array of payment.transaction model.
type PaymentTransactions []PaymentTransaction

// PaymentTransactionModel is the odoo model name.
const PaymentTransactionModel = "payment.transaction"

// Many2One convert PaymentTransaction to *Many2One.
func (pt *PaymentTransaction) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreatePaymentTransaction creates a new payment.transaction model and returns its id.
func (c *Client) CreatePaymentTransaction(pt *PaymentTransaction) (int64, error) {
	ids, err := c.CreatePaymentTransactions([]*PaymentTransaction{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePaymentTransaction creates a new payment.transaction model and returns its id.
func (c *Client) CreatePaymentTransactions(pts []*PaymentTransaction) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(PaymentTransactionModel, vv, nil)
}

// UpdatePaymentTransaction updates an existing payment.transaction record.
func (c *Client) UpdatePaymentTransaction(pt *PaymentTransaction) error {
	return c.UpdatePaymentTransactions([]int64{pt.Id.Get()}, pt)
}

// UpdatePaymentTransactions updates existing payment.transaction records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdatePaymentTransactions(ids []int64, pt *PaymentTransaction) error {
	return c.Update(PaymentTransactionModel, ids, pt, nil)
}

// DeletePaymentTransaction deletes an existing payment.transaction record.
func (c *Client) DeletePaymentTransaction(id int64) error {
	return c.DeletePaymentTransactions([]int64{id})
}

// DeletePaymentTransactions deletes existing payment.transaction records.
func (c *Client) DeletePaymentTransactions(ids []int64) error {
	return c.Delete(PaymentTransactionModel, ids)
}

// GetPaymentTransaction gets payment.transaction existing record.
func (c *Client) GetPaymentTransaction(id int64) (*PaymentTransaction, error) {
	pts, err := c.GetPaymentTransactions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pts)[0]), nil
}

// GetPaymentTransactions gets payment.transaction existing records.
func (c *Client) GetPaymentTransactions(ids []int64) (*PaymentTransactions, error) {
	pts := &PaymentTransactions{}
	if err := c.Read(PaymentTransactionModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindPaymentTransaction finds payment.transaction record by querying it with criteria.
func (c *Client) FindPaymentTransaction(criteria *Criteria) (*PaymentTransaction, error) {
	pts := &PaymentTransactions{}
	if err := c.SearchRead(PaymentTransactionModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	return &((*pts)[0]), nil
}

// FindPaymentTransactions finds payment.transaction records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentTransactions(criteria *Criteria, options *Options) (*PaymentTransactions, error) {
	pts := &PaymentTransactions{}
	if err := c.SearchRead(PaymentTransactionModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindPaymentTransactionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPaymentTransactionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PaymentTransactionModel, criteria, options)
}

// FindPaymentTransactionId finds record id by querying it with criteria.
func (c *Client) FindPaymentTransactionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PaymentTransactionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
