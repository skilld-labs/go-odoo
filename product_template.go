package odoo

// ProductTemplate represents product.template model.
type ProductTemplate struct {
	AccountTagIds                          *Relation   `xmlrpc:"account_tag_ids,omitempty"`
	Active                                 *Bool       `xmlrpc:"active,omitempty"`
	ActivityDateDeadline                   *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration            *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon                  *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                            *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState                          *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                        *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon                       *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId                         *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                         *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AlertTime                              *Int        `xmlrpc:"alert_time,omitempty"`
	AttributeLineIds                       *Relation   `xmlrpc:"attribute_line_ids,omitempty"`
	AvailableInPos                         *Bool       `xmlrpc:"available_in_pos,omitempty"`
	Barcode                                *String     `xmlrpc:"barcode,omitempty"`
	BomCount                               *Int        `xmlrpc:"bom_count,omitempty"`
	BomIds                                 *Relation   `xmlrpc:"bom_ids,omitempty"`
	BomLineIds                             *Relation   `xmlrpc:"bom_line_ids,omitempty"`
	CanImage1024BeZoomed                   *Bool       `xmlrpc:"can_image_1024_be_zoomed,omitempty"`
	CategId                                *Many2One   `xmlrpc:"categ_id,omitempty"`
	Color                                  *Int        `xmlrpc:"color,omitempty"`
	ComboIds                               *Relation   `xmlrpc:"combo_ids,omitempty"`
	CompanyId                              *Many2One   `xmlrpc:"company_id,omitempty"`
	CostCurrencyId                         *Many2One   `xmlrpc:"cost_currency_id,omitempty"`
	CostMethod                             *Selection  `xmlrpc:"cost_method,omitempty"`
	CountryOfOrigin                        *Many2One   `xmlrpc:"country_of_origin,omitempty"`
	CreateDate                             *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                              *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                             *Many2One   `xmlrpc:"currency_id,omitempty"`
	DefaultCode                            *String     `xmlrpc:"default_code,omitempty"`
	Description                            *String     `xmlrpc:"description,omitempty"`
	DescriptionPicking                     *String     `xmlrpc:"description_picking,omitempty"`
	DescriptionPickingin                   *String     `xmlrpc:"description_pickingin,omitempty"`
	DescriptionPickingout                  *String     `xmlrpc:"description_pickingout,omitempty"`
	DescriptionPurchase                    *String     `xmlrpc:"description_purchase,omitempty"`
	DescriptionSale                        *String     `xmlrpc:"description_sale,omitempty"`
	DescriptionSelfOrder                   *String     `xmlrpc:"description_self_order,omitempty"`
	DetailedType                           *Selection  `xmlrpc:"detailed_type,omitempty"`
	DisplayName                            *String     `xmlrpc:"display_name,omitempty"`
	ExpensePolicy                          *Selection  `xmlrpc:"expense_policy,omitempty"`
	ExpirationTime                         *Int        `xmlrpc:"expiration_time,omitempty"`
	FiscalCountryCodes                     *String     `xmlrpc:"fiscal_country_codes,omitempty"`
	HasAvailableRouteIds                   *Bool       `xmlrpc:"has_available_route_ids,omitempty"`
	HasConfigurableAttributes              *Bool       `xmlrpc:"has_configurable_attributes,omitempty"`
	HasMessage                             *Bool       `xmlrpc:"has_message,omitempty"`
	HsCode                                 *String     `xmlrpc:"hs_code,omitempty"`
	Id                                     *Int        `xmlrpc:"id,omitempty"`
	Image1024                              *String     `xmlrpc:"image_1024,omitempty"`
	Image128                               *String     `xmlrpc:"image_128,omitempty"`
	Image1920                              *String     `xmlrpc:"image_1920,omitempty"`
	Image256                               *String     `xmlrpc:"image_256,omitempty"`
	Image512                               *String     `xmlrpc:"image_512,omitempty"`
	IncomingQty                            *Float      `xmlrpc:"incoming_qty,omitempty"`
	InvoicePolicy                          *Selection  `xmlrpc:"invoice_policy,omitempty"`
	IsKits                                 *Bool       `xmlrpc:"is_kits,omitempty"`
	IsProductVariant                       *Bool       `xmlrpc:"is_product_variant,omitempty"`
	ListPrice                              *Float      `xmlrpc:"list_price,omitempty"`
	LocationId                             *Many2One   `xmlrpc:"location_id,omitempty"`
	MessageAttachmentCount                 *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds                     *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError                        *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter                 *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError                     *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                             *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                      *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction                      *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter               *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                      *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MrpProductQty                          *Float      `xmlrpc:"mrp_product_qty,omitempty"`
	MyActivityDateDeadline                 *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                   *String     `xmlrpc:"name,omitempty"`
	NbrMovesIn                             *Int        `xmlrpc:"nbr_moves_in,omitempty"`
	NbrMovesOut                            *Int        `xmlrpc:"nbr_moves_out,omitempty"`
	NbrReorderingRules                     *Int        `xmlrpc:"nbr_reordering_rules,omitempty"`
	OptionalProductIds                     *Relation   `xmlrpc:"optional_product_ids,omitempty"`
	OutgoingQty                            *Float      `xmlrpc:"outgoing_qty,omitempty"`
	PackagingIds                           *Relation   `xmlrpc:"packaging_ids,omitempty"`
	PosCategIds                            *Relation   `xmlrpc:"pos_categ_ids,omitempty"`
	PricelistItemCount                     *Int        `xmlrpc:"pricelist_item_count,omitempty"`
	Priority                               *Selection  `xmlrpc:"priority,omitempty"`
	ProductDocumentCount                   *Int        `xmlrpc:"product_document_count,omitempty"`
	ProductDocumentIds                     *Relation   `xmlrpc:"product_document_ids,omitempty"`
	ProductProperties                      interface{} `xmlrpc:"product_properties,omitempty"`
	ProductTagIds                          *Relation   `xmlrpc:"product_tag_ids,omitempty"`
	ProductTooltip                         *String     `xmlrpc:"product_tooltip,omitempty"`
	ProductVariantCount                    *Int        `xmlrpc:"product_variant_count,omitempty"`
	ProductVariantId                       *Many2One   `xmlrpc:"product_variant_id,omitempty"`
	ProductVariantIds                      *Relation   `xmlrpc:"product_variant_ids,omitempty"`
	PropertyAccountCreditorPriceDifference *Many2One   `xmlrpc:"property_account_creditor_price_difference,omitempty"`
	PropertyAccountExpenseId               *Many2One   `xmlrpc:"property_account_expense_id,omitempty"`
	PropertyAccountIncomeId                *Many2One   `xmlrpc:"property_account_income_id,omitempty"`
	PropertyStockInventory                 *Many2One   `xmlrpc:"property_stock_inventory,omitempty"`
	PropertyStockProduction                *Many2One   `xmlrpc:"property_stock_production,omitempty"`
	PurchaseLineWarn                       *Selection  `xmlrpc:"purchase_line_warn,omitempty"`
	PurchaseLineWarnMsg                    *String     `xmlrpc:"purchase_line_warn_msg,omitempty"`
	PurchaseMethod                         *Selection  `xmlrpc:"purchase_method,omitempty"`
	PurchaseOk                             *Bool       `xmlrpc:"purchase_ok,omitempty"`
	PurchasedProductQty                    *Float      `xmlrpc:"purchased_product_qty,omitempty"`
	QtyAvailable                           *Float      `xmlrpc:"qty_available,omitempty"`
	QualityControlPointQty                 *Int        `xmlrpc:"quality_control_point_qty,omitempty"`
	QualityFailQty                         *Int        `xmlrpc:"quality_fail_qty,omitempty"`
	QualityPassQty                         *Int        `xmlrpc:"quality_pass_qty,omitempty"`
	RatingIds                              *Relation   `xmlrpc:"rating_ids,omitempty"`
	RemovalTime                            *Int        `xmlrpc:"removal_time,omitempty"`
	ReorderingMaxQty                       *Float      `xmlrpc:"reordering_max_qty,omitempty"`
	ReorderingMinQty                       *Float      `xmlrpc:"reordering_min_qty,omitempty"`
	ResponsibleId                          *Many2One   `xmlrpc:"responsible_id,omitempty"`
	RouteFromCategIds                      *Relation   `xmlrpc:"route_from_categ_ids,omitempty"`
	RouteIds                               *Relation   `xmlrpc:"route_ids,omitempty"`
	SaleDelay                              *Int        `xmlrpc:"sale_delay,omitempty"`
	SaleLineWarn                           *Selection  `xmlrpc:"sale_line_warn,omitempty"`
	SaleLineWarnMsg                        *String     `xmlrpc:"sale_line_warn_msg,omitempty"`
	SaleOk                                 *Bool       `xmlrpc:"sale_ok,omitempty"`
	SalesCount                             *Float      `xmlrpc:"sales_count,omitempty"`
	ScheduleCount                          *Int        `xmlrpc:"schedule_count,omitempty"`
	SelfOrderAvailable                     *Bool       `xmlrpc:"self_order_available,omitempty"`
	SellerIds                              *Relation   `xmlrpc:"seller_ids,omitempty"`
	Sequence                               *Int        `xmlrpc:"sequence,omitempty"`
	ServiceToPurchase                      *Bool       `xmlrpc:"service_to_purchase,omitempty"`
	ServiceType                            *Selection  `xmlrpc:"service_type,omitempty"`
	ShowForecastedQtyStatusButton          *Bool       `xmlrpc:"show_forecasted_qty_status_button,omitempty"`
	ShowOnHandQtyStatusButton              *Bool       `xmlrpc:"show_on_hand_qty_status_button,omitempty"`
	StandardPrice                          *Float      `xmlrpc:"standard_price,omitempty"`
	SupplierTaxesId                        *Relation   `xmlrpc:"supplier_taxes_id,omitempty"`
	TaxString                              *String     `xmlrpc:"tax_string,omitempty"`
	TaxesId                                *Relation   `xmlrpc:"taxes_id,omitempty"`
	ToWeight                               *Bool       `xmlrpc:"to_weight,omitempty"`
	Tracking                               *Selection  `xmlrpc:"tracking,omitempty"`
	Type                                   *Selection  `xmlrpc:"type,omitempty"`
	UomId                                  *Many2One   `xmlrpc:"uom_id,omitempty"`
	UomName                                *String     `xmlrpc:"uom_name,omitempty"`
	UomPoId                                *Many2One   `xmlrpc:"uom_po_id,omitempty"`
	UseExpirationDate                      *Bool       `xmlrpc:"use_expiration_date,omitempty"`
	UseTime                                *Int        `xmlrpc:"use_time,omitempty"`
	UsedInBomCount                         *Int        `xmlrpc:"used_in_bom_count,omitempty"`
	ValidProductTemplateAttributeLineIds   *Relation   `xmlrpc:"valid_product_template_attribute_line_ids,omitempty"`
	Valuation                              *Selection  `xmlrpc:"valuation,omitempty"`
	VariantSellerIds                       *Relation   `xmlrpc:"variant_seller_ids,omitempty"`
	VirtualAvailable                       *Float      `xmlrpc:"virtual_available,omitempty"`
	VisibleExpensePolicy                   *Bool       `xmlrpc:"visible_expense_policy,omitempty"`
	Volume                                 *Float      `xmlrpc:"volume,omitempty"`
	VolumeUomName                          *String     `xmlrpc:"volume_uom_name,omitempty"`
	WarehouseId                            *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WebsiteMessageIds                      *Relation   `xmlrpc:"website_message_ids,omitempty"`
	Weight                                 *Float      `xmlrpc:"weight,omitempty"`
	WeightUomName                          *String     `xmlrpc:"weight_uom_name,omitempty"`
	WriteDate                              *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                               *Many2One   `xmlrpc:"write_uid,omitempty"`
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
	ids, err := c.CreateProductTemplates([]*ProductTemplate{pt})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductTemplate creates a new product.template model and returns its id.
func (c *Client) CreateProductTemplates(pts []*ProductTemplate) ([]int64, error) {
	var vv []interface{}
	for _, v := range pts {
		vv = append(vv, v)
	}
	return c.Create(ProductTemplateModel, vv, nil)
}

// UpdateProductTemplate updates an existing product.template record.
func (c *Client) UpdateProductTemplate(pt *ProductTemplate) error {
	return c.UpdateProductTemplates([]int64{pt.Id.Get()}, pt)
}

// UpdateProductTemplates updates existing product.template records.
// All records (represented by ids) will be updated by pt values.
func (c *Client) UpdateProductTemplates(ids []int64, pt *ProductTemplate) error {
	return c.Update(ProductTemplateModel, ids, pt, nil)
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
	return &((*pts)[0]), nil
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
	return &((*pts)[0]), nil
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
	return c.Search(ProductTemplateModel, criteria, options)
}

// FindProductTemplateId finds record id by querying it with criteria.
func (c *Client) FindProductTemplateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductTemplateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
