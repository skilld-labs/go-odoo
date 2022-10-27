package odoo

import (
	"fmt"
)

// ProductTemplate represents product.template model.
type ProductTemplate struct {
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
	CategId                              *Many2One  `xmlrpc:"categ_id,omitempty"`
	Color                                *Int       `xmlrpc:"color,omitempty"`
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
	Price                                *Float     `xmlrpc:"price,omitempty"`
	PricelistId                          *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	PricelistItemCount                   *Int       `xmlrpc:"pricelist_item_count,omitempty"`
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

// ProductTemplates represents array of product.template model.
type ProductTemplates []ProductTemplate

// ProductTemplateModel is the odoo model name.
const ProductTemplateModel = "product.template"

// Many2One convert ProductTemplate to *Many2One.
func (pt *ProductTemplate) Many2One() *Many2One {
	return NewMany2One(pt.Id.Get(), "")
}

// CreateProductTemplate creates a new product.template model and returns its id.
func (c *Client) CreateProductTemplate(pt *ProductTemplate) (int64, error) {
	return c.Create(ProductTemplateModel, pt)
}

// UpdateProductTemplate updates an existing product.template record.
func (c *Client) UpdateProductTemplate(pt *ProductTemplate) error {
	return c.UpdateProductTemplates([]int64{pt.Id.Get()}, pt)
}

// UpdateProductTemplates updates existing product.template records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProductTemplates(ids []int64, pt *ProductTemplate) error {
	return c.Update(ProductTemplateModel, ids, pt)
}

// DeleteProductTemplate deletes an existing product.template record.
func (c *Client) DeleteProductTemplate(id int64) error {
	return c.DeleteProductTemplates([]int64{id})
}

// DeleteProductTemplates deletes existing product.template records.
func (c *Client) DeleteProductTemplates(ids []int64) error {
	return c.Delete(ProductTemplateModel, ids)
}

// GetProductTemplate gets product.template existing record.
func (c *Client) GetProductTemplate(id int64) (*ProductTemplate, error) {
	pts, err := c.GetProductTemplates([]int64{id})
	if err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.template not found", id)
}

// GetProductTemplates gets product.template existing records.
func (c *Client) GetProductTemplates(ids []int64) (*ProductTemplates, error) {
	pts := &ProductTemplates{}
	if err := c.Read(ProductTemplateModel, ids, nil, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTemplate finds product.template record by querying it with criteria.
func (c *Client) FindProductTemplate(criteria *Criteria) (*ProductTemplate, error) {
	pts := &ProductTemplates{}
	if err := c.SearchRead(ProductTemplateModel, criteria, NewOptions().Limit(1), pts); err != nil {
		return nil, err
	}
	if pts != nil && len(*pts) > 0 {
		return &((*pts)[0]), nil
	}
	return nil, fmt.Errorf("product.template was not found")
}

// FindProductTemplates finds product.template records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplates(criteria *Criteria, options *Options) (*ProductTemplates, error) {
	pts := &ProductTemplates{}
	if err := c.SearchRead(ProductTemplateModel, criteria, options, pts); err != nil {
		return nil, err
	}
	return pts, nil
}

// FindProductTemplateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductTemplateIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductTemplateModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductTemplateId finds record id by querying it with criteria.
func (c *Client) FindProductTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.template was not found")
}
