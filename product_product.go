package odoo

// ProductProduct represents product.product model.
type ProductProduct struct {
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
	AdditionalProductTagIds                *Relation   `xmlrpc:"additional_product_tag_ids,omitempty"`
	AlertTime                              *Int        `xmlrpc:"alert_time,omitempty"`
	AllProductTagIds                       *Relation   `xmlrpc:"all_product_tag_ids,omitempty"`
	AttributeLineIds                       *Relation   `xmlrpc:"attribute_line_ids,omitempty"`
	AvailableInPos                         *Bool       `xmlrpc:"available_in_pos,omitempty"`
	AvgCost                                *Float      `xmlrpc:"avg_cost,omitempty"`
	Barcode                                *String     `xmlrpc:"barcode,omitempty"`
	BomCount                               *Int        `xmlrpc:"bom_count,omitempty"`
	BomIds                                 *Relation   `xmlrpc:"bom_ids,omitempty"`
	BomLineIds                             *Relation   `xmlrpc:"bom_line_ids,omitempty"`
	CanImage1024BeZoomed                   *Bool       `xmlrpc:"can_image_1024_be_zoomed,omitempty"`
	CanImageVariant1024BeZoomed            *Bool       `xmlrpc:"can_image_variant_1024_be_zoomed,omitempty"`
	CategId                                *Many2One   `xmlrpc:"categ_id,omitempty"`
	Code                                   *String     `xmlrpc:"code,omitempty"`
	Color                                  *Int        `xmlrpc:"color,omitempty"`
	CombinationIndices                     *String     `xmlrpc:"combination_indices,omitempty"`
	ComboIds                               *Relation   `xmlrpc:"combo_ids,omitempty"`
	CompanyCurrencyId                      *Many2One   `xmlrpc:"company_currency_id,omitempty"`
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
	FreeQty                                *Float      `xmlrpc:"free_qty,omitempty"`
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
	ImageVariant1024                       *String     `xmlrpc:"image_variant_1024,omitempty"`
	ImageVariant128                        *String     `xmlrpc:"image_variant_128,omitempty"`
	ImageVariant1920                       *String     `xmlrpc:"image_variant_1920,omitempty"`
	ImageVariant256                        *String     `xmlrpc:"image_variant_256,omitempty"`
	ImageVariant512                        *String     `xmlrpc:"image_variant_512,omitempty"`
	IncomingQty                            *Float      `xmlrpc:"incoming_qty,omitempty"`
	InvoicePolicy                          *Selection  `xmlrpc:"invoice_policy,omitempty"`
	IsKits                                 *Bool       `xmlrpc:"is_kits,omitempty"`
	IsProductVariant                       *Bool       `xmlrpc:"is_product_variant,omitempty"`
	ListPrice                              *Float      `xmlrpc:"list_price,omitempty"`
	LocationId                             *Many2One   `xmlrpc:"location_id,omitempty"`
	LotPropertiesDefinition                interface{} `xmlrpc:"lot_properties_definition,omitempty"`
	LstPrice                               *Float      `xmlrpc:"lst_price,omitempty"`
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
	OrderpointIds                          *Relation   `xmlrpc:"orderpoint_ids,omitempty"`
	OutgoingQty                            *Float      `xmlrpc:"outgoing_qty,omitempty"`
	PackagingIds                           *Relation   `xmlrpc:"packaging_ids,omitempty"`
	PartnerRef                             *String     `xmlrpc:"partner_ref,omitempty"`
	PosCategIds                            *Relation   `xmlrpc:"pos_categ_ids,omitempty"`
	PriceExtra                             *Float      `xmlrpc:"price_extra,omitempty"`
	PricelistItemCount                     *Int        `xmlrpc:"pricelist_item_count,omitempty"`
	Priority                               *Selection  `xmlrpc:"priority,omitempty"`
	ProductCatalogProductIsInSaleOrder     *Bool       `xmlrpc:"product_catalog_product_is_in_sale_order,omitempty"`
	ProductDocumentCount                   *Int        `xmlrpc:"product_document_count,omitempty"`
	ProductDocumentIds                     *Relation   `xmlrpc:"product_document_ids,omitempty"`
	ProductProperties                      interface{} `xmlrpc:"product_properties,omitempty"`
	ProductTagIds                          *Relation   `xmlrpc:"product_tag_ids,omitempty"`
	ProductTemplateAttributeValueIds       *Relation   `xmlrpc:"product_template_attribute_value_ids,omitempty"`
	ProductTemplateVariantValueIds         *Relation   `xmlrpc:"product_template_variant_value_ids,omitempty"`
	ProductTmplId                          *Many2One   `xmlrpc:"product_tmpl_id,omitempty"`
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
	PurchaseOrderLineIds                   *Relation   `xmlrpc:"purchase_order_line_ids,omitempty"`
	PurchasedProductQty                    *Float      `xmlrpc:"purchased_product_qty,omitempty"`
	PutawayRuleIds                         *Relation   `xmlrpc:"putaway_rule_ids,omitempty"`
	QtyAvailable                           *Float      `xmlrpc:"qty_available,omitempty"`
	QualityControlPointQty                 *Int        `xmlrpc:"quality_control_point_qty,omitempty"`
	QualityFailQty                         *Int        `xmlrpc:"quality_fail_qty,omitempty"`
	QualityPassQty                         *Int        `xmlrpc:"quality_pass_qty,omitempty"`
	QuantitySvl                            *Float      `xmlrpc:"quantity_svl,omitempty"`
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
	StockMoveIds                           *Relation   `xmlrpc:"stock_move_ids,omitempty"`
	StockQuantIds                          *Relation   `xmlrpc:"stock_quant_ids,omitempty"`
	StockValuationLayerIds                 *Relation   `xmlrpc:"stock_valuation_layer_ids,omitempty"`
	StorageCategoryCapacityIds             *Relation   `xmlrpc:"storage_category_capacity_ids,omitempty"`
	SupplierTaxesId                        *Relation   `xmlrpc:"supplier_taxes_id,omitempty"`
	TaxString                              *String     `xmlrpc:"tax_string,omitempty"`
	TaxesId                                *Relation   `xmlrpc:"taxes_id,omitempty"`
	ToWeight                               *Bool       `xmlrpc:"to_weight,omitempty"`
	TotalValue                             *Float      `xmlrpc:"total_value,omitempty"`
	Tracking                               *Selection  `xmlrpc:"tracking,omitempty"`
	Type                                   *Selection  `xmlrpc:"type,omitempty"`
	UomId                                  *Many2One   `xmlrpc:"uom_id,omitempty"`
	UomName                                *String     `xmlrpc:"uom_name,omitempty"`
	UomPoId                                *Many2One   `xmlrpc:"uom_po_id,omitempty"`
	UseExpirationDate                      *Bool       `xmlrpc:"use_expiration_date,omitempty"`
	UseTime                                *Int        `xmlrpc:"use_time,omitempty"`
	UsedInBomCount                         *Int        `xmlrpc:"used_in_bom_count,omitempty"`
	ValidEan                               *Bool       `xmlrpc:"valid_ean,omitempty"`
	ValidProductTemplateAttributeLineIds   *Relation   `xmlrpc:"valid_product_template_attribute_line_ids,omitempty"`
	Valuation                              *Selection  `xmlrpc:"valuation,omitempty"`
	ValueSvl                               *Float      `xmlrpc:"value_svl,omitempty"`
	VariantBomIds                          *Relation   `xmlrpc:"variant_bom_ids,omitempty"`
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
	ids, err := c.CreateProductProducts([]*ProductProduct{pp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductProduct creates a new product.product model and returns its id.
func (c *Client) CreateProductProducts(pps []*ProductProduct) ([]int64, error) {
	var vv []interface{}
	for _, v := range pps {
		vv = append(vv, v)
	}
	return c.Create(ProductProductModel, vv, nil)
}

// UpdateProductProduct updates an existing product.product record.
func (c *Client) UpdateProductProduct(pp *ProductProduct) error {
	return c.UpdateProductProducts([]int64{pp.Id.Get()}, pp)
}

// UpdateProductProducts updates existing product.product records.
// All records (represented by ids) will be updated by pp values.
func (c *Client) UpdateProductProducts(ids []int64, pp *ProductProduct) error {
	return c.Update(ProductProductModel, ids, pp, nil)
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
	return &((*pps)[0]), nil
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
	return &((*pps)[0]), nil
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
	return c.Search(ProductProductModel, criteria, options)
}

// FindProductProductId finds record id by querying it with criteria.
func (c *Client) FindProductProductId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductProductModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
