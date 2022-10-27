package odoo

import (
	"fmt"
)

// ProductProduct represents product.product model.
type ProductProduct struct {
	LastUpdate                           *Time      `xmlrpc:"__last_update,omitempty"`
	Active                               *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline                 *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration          *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                          *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                        *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                      *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                     *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                       *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                       *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AttributeLineIds                     *Relation  `xmlrpc:"attribute_line_ids,omitempty"`
	Barcode                              *String    `xmlrpc:"barcode,omitempty"`
	CanImage1024BeZoomed                 *Bool      `xmlrpc:"can_image_1024_be_zoomed,omitempty"`
	CanImageVariant1024BeZoomed          *Bool      `xmlrpc:"can_image_variant_1024_be_zoomed,omitempty"`
	CategId                              *Many2One  `xmlrpc:"categ_id,omitempty"`
	Code                                 *String    `xmlrpc:"code,omitempty"`
	Color                                *Int       `xmlrpc:"color,omitempty"`
	CombinationIndices                   *String    `xmlrpc:"combination_indices,omitempty"`
	CompanyId                            *Many2One  `xmlrpc:"company_id,omitempty"`
	CostCurrencyId                       *Many2One  `xmlrpc:"cost_currency_id,omitempty"`
	CreateDate                           *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                            *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                           *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultCode                          *String    `xmlrpc:"default_code,omitempty"`
	Description                          *String    `xmlrpc:"description,omitempty"`
	DescriptionPurchase                  *String    `xmlrpc:"description_purchase,omitempty"`
	DescriptionSale                      *String    `xmlrpc:"description_sale,omitempty"`
	DisplayName                          *String    `xmlrpc:"display_name,omitempty"`
	ExpensePolicy                        *Selection `xmlrpc:"expense_policy,omitempty"`
	HasConfigurableAttributes            *Bool      `xmlrpc:"has_configurable_attributes,omitempty"`
	Id                                   *Int       `xmlrpc:"id,omitempty"`
	Image1024                            *String    `xmlrpc:"image_1024,omitempty"`
	Image128                             *String    `xmlrpc:"image_128,omitempty"`
	Image1920                            *String    `xmlrpc:"image_1920,omitempty"`
	Image256                             *String    `xmlrpc:"image_256,omitempty"`
	Image512                             *String    `xmlrpc:"image_512,omitempty"`
	ImageVariant1024                     *String    `xmlrpc:"image_variant_1024,omitempty"`
	ImageVariant128                      *String    `xmlrpc:"image_variant_128,omitempty"`
	ImageVariant1920                     *String    `xmlrpc:"image_variant_1920,omitempty"`
	ImageVariant256                      *String    `xmlrpc:"image_variant_256,omitempty"`
	ImageVariant512                      *String    `xmlrpc:"image_variant_512,omitempty"`
	InvoicePolicy                        *Selection `xmlrpc:"invoice_policy,omitempty"`
	IsProductVariant                     *Bool      `xmlrpc:"is_product_variant,omitempty"`
	ListPrice                            *Float     `xmlrpc:"list_price,omitempty"`
	LstPrice                             *Float     `xmlrpc:"lst_price,omitempty"`
	MessageAttachmentCount               *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageChannelIds                    *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds                   *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                      *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter               *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                   *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                           *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                    *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId              *Many2One  `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction                    *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter             *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                    *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                        *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter                 *Int       `xmlrpc:"message_unread_counter,omitempty"`
	MyActivityDateDeadline               *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                 *String    `xmlrpc:"name,omitempty"`
	PackagingIds                         *Relation  `xmlrpc:"packaging_ids,omitempty"`
	PartnerRef                           *String    `xmlrpc:"partner_ref,omitempty"`
	Price                                *Float     `xmlrpc:"price,omitempty"`
	PriceExtra                           *Float     `xmlrpc:"price_extra,omitempty"`
	PricelistId                          *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	PricelistItemCount                   *Int       `xmlrpc:"pricelist_item_count,omitempty"`
	ProductTemplateAttributeValueIds     *Relation  `xmlrpc:"product_template_attribute_value_ids,omitempty"`
	ProductTmplId                        *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
	ProductVariantCount                  *Int       `xmlrpc:"product_variant_count,omitempty"`
	ProductVariantId                     *Many2One  `xmlrpc:"product_variant_id,omitempty"`
	ProductVariantIds                    *Relation  `xmlrpc:"product_variant_ids,omitempty"`
	PropertyAccountExpenseId             *Many2One  `xmlrpc:"property_account_expense_id,omitempty"`
	PropertyAccountIncomeId              *Many2One  `xmlrpc:"property_account_income_id,omitempty"`
	PurchaseOk                           *Bool      `xmlrpc:"purchase_ok,omitempty"`
	SaleLineWarn                         *Selection `xmlrpc:"sale_line_warn,omitempty"`
	SaleLineWarnMsg                      *String    `xmlrpc:"sale_line_warn_msg,omitempty"`
	SaleOk                               *Bool      `xmlrpc:"sale_ok,omitempty"`
	SalesCount                           *Float     `xmlrpc:"sales_count,omitempty"`
	SellerIds                            *Relation  `xmlrpc:"seller_ids,omitempty"`
	Sequence                             *Int       `xmlrpc:"sequence,omitempty"`
	ServiceType                          *Selection `xmlrpc:"service_type,omitempty"`
	StandardPrice                        *Float     `xmlrpc:"standard_price,omitempty"`
	SupplierTaxesId                      *Relation  `xmlrpc:"supplier_taxes_id,omitempty"`
	TaxesId                              *Relation  `xmlrpc:"taxes_id,omitempty"`
	Type                                 *Selection `xmlrpc:"type,omitempty"`
	UomId                                *Many2One  `xmlrpc:"uom_id,omitempty"`
	UomName                              *String    `xmlrpc:"uom_name,omitempty"`
	UomPoId                              *Many2One  `xmlrpc:"uom_po_id,omitempty"`
	ValidProductTemplateAttributeLineIds *Relation  `xmlrpc:"valid_product_template_attribute_line_ids,omitempty"`
	VariantSellerIds                     *Relation  `xmlrpc:"variant_seller_ids,omitempty"`
	VisibleExpensePolicy                 *Bool      `xmlrpc:"visible_expense_policy,omitempty"`
	VisibleQtyConfigurator               *Bool      `xmlrpc:"visible_qty_configurator,omitempty"`
	Volume                               *Float     `xmlrpc:"volume,omitempty"`
	VolumeUomName                        *String    `xmlrpc:"volume_uom_name,omitempty"`
	WebsiteMessageIds                    *Relation  `xmlrpc:"website_message_ids,omitempty"`
	Weight                               *Float     `xmlrpc:"weight,omitempty"`
	WeightUomName                        *String    `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate                            *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                             *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductProducts represents array of product.product model.
type ProductProducts []ProductProduct

// ProductProductModel is the odoo model name.
const ProductProductModel = "product.product"

// Many2One convert ProductProduct to *Many2One.
func (pp *ProductProduct) Many2One() *Many2One {
	return NewMany2One(pp.Id.Get(), "")
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProduct(pp *ProductProduct) (int64, error) {
	return c.Create(ProductProductModel, pp)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp)
}

// DeleteProductProduct deletes an existing product.product record.
func (c *Client) DeleteProductProduct(id int64) error {
	return c.DeleteProductProducts([]int64{id})
}

// DeleteProductProducts deletes existing product.product records.
func (c *Client) DeleteProductProducts(ids []int64) error {
	return c.Delete(ProductProductModel, ids)
}

// GetProductProduct gets product.product existing record.
func (c *Client) GetProductProduct(id int64) (*ProductProduct, error) {
	pps, err := c.GetProductProducts([]int64{id})
	if err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.product not found", id)
}

// GetProductProducts gets product.product existing records.
func (c *Client) GetProductProducts(ids []int64) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.Read(ProductProductModel, ids, nil, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProduct finds product.product record by querying it with criteria.
func (c *Client) FindProductProduct(criteria *Criteria) (*ProductProduct, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, NewOptions().Limit(1), pps); err != nil {
		return nil, err
	}
	if pps != nil && len(*pps) > 0 {
		return &((*pps)[0]), nil
	}
	return nil, fmt.Errorf("product.product was not found")
}

// FindProductProducts finds product.product records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProducts(criteria *Criteria, options *Options) (*ProductProducts, error) {
	pps := &ProductProducts{}
	if err := c.SearchRead(ProductProductModel, criteria, options, pps); err != nil {
		return nil, err
	}
	return pps, nil
}

// FindProductProductIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductProductIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.product was not found")
}
