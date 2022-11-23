package odoo

import "encoding/json"

// SaleOrderModelID is the model identifier we use to fetch a SaleOrder from Odoo.
const SaleOrderModelID = "sale.order"

// SaleOrder represents a sales order within Odoo.
type SaleOrder struct {
	ID                       int           `json:"id"`
	AuthorizedTransactionIds []interface{} `json:"authorized_transaction_ids"`
	State                    string        `json:"state"`
	PartnerCreditWarning     string        `json:"partner_credit_warning"`
	InvoiceCount             int           `json:"invoice_count"`
	Name                     string        `json:"name"`
	PartnerId                []interface{} `json:"partner_id"`
	SaleOrderTemplateId      bool          `json:"sale_order_template_id"`
	ValidityDate             bool          `json:"validity_date"`
	DateOrder                string        `json:"date_order"`
	ShowUpdatePricelist      bool          `json:"show_update_pricelist"`
	PricelistId              []interface{} `json:"pricelist_id"`
	CompanyId                []interface{} `json:"company_id"`
	CurrencyId               []interface{} `json:"currency_id"`
	TaxCountryId             []interface{} `json:"tax_country_id"`
	PaymentTermId            bool          `json:"payment_term_id"`
	OrderLine                []int         `json:"order_line"`
	Note                     bool          `json:"note"`
	TaxTotals                struct {
		AmountUntaxed          float64 `json:"amount_untaxed"`
		AmountTotal            float64 `json:"amount_total"`
		FormattedAmountTotal   string  `json:"formatted_amount_total"`
		FormattedAmountUntaxed string  `json:"formatted_amount_untaxed"`
		GroupsBySubtotal       struct {
		} `json:"groups_by_subtotal"`
		Subtotals      []interface{} `json:"subtotals"`
		SubtotalsOrder []interface{} `json:"subtotals_order"`
		DisplayTaxBase bool          `json:"display_tax_base"`
	} `json:"tax_totals"`
	SaleOrderOptionIds []interface{} `json:"sale_order_option_ids"`
	UserId             []interface{} `json:"user_id"`
	TeamId             []interface{} `json:"team_id"`
	RequireSignature   bool          `json:"require_signature"`
	RequirePayment     bool          `json:"require_payment"`
	Reference          bool          `json:"reference"`
	ClientOrderRef     bool          `json:"client_order_ref"`
	TagIds             []interface{} `json:"tag_ids"`
	ShowUpdateFpos     bool          `json:"show_update_fpos"`
	FiscalPositionId   bool          `json:"fiscal_position_id"`
	PartnerInvoiceId   []interface{} `json:"partner_invoice_id"`
	InvoiceStatus      string        `json:"invoice_status"`
	CommitmentDate     bool          `json:"commitment_date"`
	ExpectedDate       string        `json:"expected_date"`
	Origin             bool          `json:"origin"`
	CampaignId         []interface{} `json:"campaign_id"`
	MediumId           []interface{} `json:"medium_id"`
	SourceId           []interface{} `json:"source_id"`
	DisplayName        string        `json:"display_name"`
}

// GetSalesOrder fetches a sales order from Odoo.
func (c *Client) GetSalesOrder(id int) ([]byte, error) {
	saleOrder, err := c.Read(SaleOrderModelID, []int{id}, nil)
	if err != nil {
		return nil, err
	}

	return json.Marshal(saleOrder)
}
