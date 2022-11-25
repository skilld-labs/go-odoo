package odoo

import (
	"encoding/json"
	"log"
	"time"
)

// ProductModelID is the model identifier for a product within Odoo.
const ProductModelID = "product.product"

// Product represents a product within Odoo.
type Product struct {
	LastUpdate                           string        `json:"__last_update"`
	AccountTagIds                        []interface{} `json:"account_tag_ids"`
	Active                               bool          `json:"active"`
	ActivityDateDeadline                 bool          `json:"activity_date_deadline"`
	ActivityExceptionDecoration          bool          `json:"activity_exception_decoration"`
	ActivityExceptionIcon                bool          `json:"activity_exception_icon"`
	ActivityIds                          []interface{} `json:"activity_ids"`
	ActivityState                        bool          `json:"activity_state"`
	ActivitySummary                      bool          `json:"activity_summary"`
	ActivityTypeIcon                     bool          `json:"activity_type_icon"`
	ActivityTypeId                       bool          `json:"activity_type_id"`
	ActivityUserId                       bool          `json:"activity_user_id"`
	AdditionalProductTagIds              []interface{} `json:"additional_product_tag_ids"`
	AllProductTagIds                     []interface{} `json:"all_product_tag_ids"`
	AttributeLineIds                     []interface{} `json:"attribute_line_ids"`
	Barcode                              bool          `json:"barcode"`
	CanImage1024BeZoomed                 bool          `json:"can_image_1024_be_zoomed"`
	CanImageVariant1024BeZoomed          bool          `json:"can_image_variant_1024_be_zoomed"`
	CategId                              []interface{} `json:"categ_id"`
	Code                                 string        `json:"code"`
	Color                                int           `json:"color"`
	CombinationIndices                   string        `json:"combination_indices"`
	CompanyId                            bool          `json:"company_id"`
	CostCurrencyId                       []interface{} `json:"cost_currency_id"`
	CreateDate                           string        `json:"create_date"`
	CreateUid                            []interface{} `json:"create_uid"`
	CurrencyId                           []interface{} `json:"currency_id"`
	DefaultCode                          string        `json:"default_code"`
	Description                          bool          `json:"description"`
	DescriptionPurchase                  bool          `json:"description_purchase"`
	DescriptionSale                      string        `json:"description_sale"`
	DetailedType                         string        `json:"detailed_type"`
	DisplayName                          string        `json:"display_name"`
	ExpensePolicy                        string        `json:"expense_policy"`
	HasConfigurableAttributes            bool          `json:"has_configurable_attributes"`
	HasMessage                           bool          `json:"has_message"`
	ID                                   int           `json:"id"`
	Image1024                            string        `json:"image_1024"`
	Image128                             string        `json:"image_128"`
	Image1920                            string        `json:"image_1920"`
	Image256                             string        `json:"image_256"`
	Image512                             string        `json:"image_512"`
	ImageVariant1024                     bool          `json:"image_variant_1024"`
	ImageVariant128                      bool          `json:"image_variant_128"`
	ImageVariant1920                     bool          `json:"image_variant_1920"`
	ImageVariant256                      bool          `json:"image_variant_256"`
	ImageVariant512                      bool          `json:"image_variant_512"`
	InvoicePolicy                        string        `json:"invoice_policy"`
	IsProductVariant                     bool          `json:"is_product_variant"`
	ListPrice                            int           `json:"list_price"`
	LstPrice                             int           `json:"lst_price"`
	MessageAttachmentCount               int           `json:"message_attachment_count"`
	MessageFollowerIds                   []int         `json:"message_follower_ids"`
	MessageHasError                      bool          `json:"message_has_error"`
	MessageHasErrorCounter               int           `json:"message_has_error_counter"`
	MessageHasSmsError                   bool          `json:"message_has_sms_error"`
	MessageIds                           []int         `json:"message_ids"`
	MessageIsFollower                    bool          `json:"message_is_follower"`
	MessageMainAttachmentId              bool          `json:"message_main_attachment_id"`
	MessageNeedaction                    bool          `json:"message_needaction"`
	MessageNeedactionCounter             int           `json:"message_needaction_counter"`
	MessagePartnerIds                    []interface{} `json:"message_partner_ids"`
	MyActivityDateDeadline               bool          `json:"my_activity_date_deadline"`
	Name                                 string        `json:"name"`
	OptionalProductIds                   []interface{} `json:"optional_product_ids"`
	PackagingIds                         []interface{} `json:"packaging_ids"`
	PartnerRef                           string        `json:"partner_ref"`
	PriceExtra                           int           `json:"price_extra"`
	PricelistItemCount                   int           `json:"pricelist_item_count"`
	Priority                             string        `json:"priority"`
	ProductTagIds                        []interface{} `json:"product_tag_ids"`
	ProductTemplateAttributeValueIds     []interface{} `json:"product_template_attribute_value_ids"`
	ProductTemplateVariantValueIds       []interface{} `json:"product_template_variant_value_ids"`
	ProductTmplId                        []interface{} `json:"product_tmpl_id"`
	ProductTooltip                       string        `json:"product_tooltip"`
	ProductVariantCount                  int           `json:"product_variant_count"`
	ProductVariantId                     []interface{} `json:"product_variant_id"`
	ProductVariantIds                    []int         `json:"product_variant_ids"`
	PropertyAccountExpenseId             bool          `json:"property_account_expense_id"`
	PropertyAccountIncomeId              bool          `json:"property_account_income_id"`
	PurchaseOk                           bool          `json:"purchase_ok"`
	SaleLineWarn                         string        `json:"sale_line_warn"`
	SaleLineWarnMsg                      bool          `json:"sale_line_warn_msg"`
	SaleOk                               bool          `json:"sale_ok"`
	SalesCount                           int           `json:"sales_count"`
	SellerIds                            []interface{} `json:"seller_ids"`
	Sequence                             int           `json:"sequence"`
	ServiceType                          string        `json:"service_type"`
	StandardPrice                        int           `json:"standard_price"`
	SupplierTaxesId                      []interface{} `json:"supplier_taxes_id"`
	TaxString                            string        `json:"tax_string"`
	TaxesId                              []interface{} `json:"taxes_id"`
	Type                                 string        `json:"type"`
	UomId                                []interface{} `json:"uom_id"`
	UomName                              string        `json:"uom_name"`
	UomPoId                              []interface{} `json:"uom_po_id"`
	ValidProductTemplateAttributeLineIds []interface{} `json:"valid_product_template_attribute_line_ids"`
	VariantSellerIds                     []interface{} `json:"variant_seller_ids"`
	VisibleExpensePolicy                 bool          `json:"visible_expense_policy"`
	VisibleQtyConfigurator               bool          `json:"visible_qty_configurator"`
	Volume                               int           `json:"volume"`
	VolumeUomName                        string        `json:"volume_uom_name"`
	WebsiteMessageIds                    []interface{} `json:"website_message_ids"`
	Weight                               float64       `json:"weight"`
	WeightUomName                        string        `json:"weight_uom_name"`
	WriteDate                            string        `json:"write_date"`
	WriteUid                             []interface{} `json:"write_uid"`
}

// GetProductsCreatedAfter gets all product.product records that have been created after the time passed in.
// This allows us to run a cron job every x minutes and return new products. We return encoded JSON
// within this method which makes it easy to use within XFuze.
func (c *Client) GetProductsCreatedAfter(t time.Time) ([]byte, error) {
	dtLayout := "2006-01-02 15:04:05"
	log.Printf("Looking for products created after %s", t.Format(dtLayout))

	product, err := c.SearchRead(ProductModelID, List{List{"create_date", ">=", t.Format(dtLayout)}}, nil)
	if err != nil {
		return nil, err
	}

	return json.Marshal(product)
}

// GetProduct returns a JSON encoded byte slice of an Odoo product.
func (c *Client) GetProduct(id int) ([]byte, error) {
	product, err := c.Read(ProductModelID, []int{id}, nil)
	if err != nil {
		return nil, err
	}

	return json.Marshal(product)
}
