package odoo

// ProductProduct represents product.product model.
type ProductProduct struct {
	AccessoryProductIds                    *Relation   `xmlrpc:"accessory_product_ids,omitempty"`
	AccountTagIds                          *Relation   `xmlrpc:"account_tag_ids,omitempty"`
	Active                                 *Bool       `xmlrpc:"active,omitempty"`
	ActivityCalendarEventId                *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
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
	AllProductTagIds                       *Relation   `xmlrpc:"all_product_tag_ids,omitempty"`
	AllowOutOfStockOrder                   *Bool       `xmlrpc:"allow_out_of_stock_order,omitempty"`
	AlternativeProductIds                  *Relation   `xmlrpc:"alternative_product_ids,omitempty"`
	AttributeLineIds                       *Relation   `xmlrpc:"attribute_line_ids,omitempty"`
	AvailableInPos                         *Bool       `xmlrpc:"available_in_pos,omitempty"`
	AvailableThreshold                     *Float      `xmlrpc:"available_threshold,omitempty"`
	AvgCost                                *Float      `xmlrpc:"avg_cost,omitempty"`
	Barcode                                *String     `xmlrpc:"barcode,omitempty"`
	BaseUnitCount                          *Float      `xmlrpc:"base_unit_count,omitempty"`
	BaseUnitId                             *Many2One   `xmlrpc:"base_unit_id,omitempty"`
	BaseUnitName                           *String     `xmlrpc:"base_unit_name,omitempty"`
	BaseUnitPrice                          *Float      `xmlrpc:"base_unit_price,omitempty"`
	CanImage1024BeZoomed                   *Bool       `xmlrpc:"can_image_1024_be_zoomed,omitempty"`
	CanImageVariant1024BeZoomed            *Bool       `xmlrpc:"can_image_variant_1024_be_zoomed,omitempty"`
	CanPublish                             *Bool       `xmlrpc:"can_publish,omitempty"`
	CategId                                *Many2One   `xmlrpc:"categ_id,omitempty"`
	Code                                   *String     `xmlrpc:"code,omitempty"`
	Color                                  *Int        `xmlrpc:"color,omitempty"`
	CombinationIndices                     *String     `xmlrpc:"combination_indices,omitempty"`
	ComboIds                               *Relation   `xmlrpc:"combo_ids,omitempty"`
	CompanyCurrencyId                      *Many2One   `xmlrpc:"company_currency_id,omitempty"`
	CompanyId                              *Many2One   `xmlrpc:"company_id,omitempty"`
	CompareListPrice                       *Float      `xmlrpc:"compare_list_price,omitempty"`
	CostCurrencyId                         *Many2One   `xmlrpc:"cost_currency_id,omitempty"`
	CostMethod                             *Selection  `xmlrpc:"cost_method,omitempty"`
	CountryOfOrigin                        *Many2One   `xmlrpc:"country_of_origin,omitempty"`
	CreateDate                             *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                              *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                             *Many2One   `xmlrpc:"currency_id,omitempty"`
	DefaultCode                            *String     `xmlrpc:"default_code,omitempty"`
	Description                            *String     `xmlrpc:"description,omitempty"`
	DescriptionEcommerce                   *String     `xmlrpc:"description_ecommerce,omitempty"`
	DescriptionPicking                     *String     `xmlrpc:"description_picking,omitempty"`
	DescriptionPickingin                   *String     `xmlrpc:"description_pickingin,omitempty"`
	DescriptionPickingout                  *String     `xmlrpc:"description_pickingout,omitempty"`
	DescriptionPurchase                    *String     `xmlrpc:"description_purchase,omitempty"`
	DescriptionSale                        *String     `xmlrpc:"description_sale,omitempty"`
	DisplayName                            *String     `xmlrpc:"display_name,omitempty"`
	EventTicketIds                         *Relation   `xmlrpc:"event_ticket_ids,omitempty"`
	ExpensePolicy                          *Selection  `xmlrpc:"expense_policy,omitempty"`
	FiscalCountryCodes                     *String     `xmlrpc:"fiscal_country_codes,omitempty"`
	ForceSaleMaxQty                        *Bool       `xmlrpc:"force_sale_max_qty,omitempty"`
	ForceSaleMinQty                        *Bool       `xmlrpc:"force_sale_min_qty,omitempty"`
	FreeQty                                *Float      `xmlrpc:"free_qty,omitempty"`
	HasAvailableRouteIds                   *Bool       `xmlrpc:"has_available_route_ids,omitempty"`
	HasConfigurableAttributes              *Bool       `xmlrpc:"has_configurable_attributes,omitempty"`
	HasImage                               *Bool       `xmlrpc:"has_image,omitempty"`
	HasMessage                             *Bool       `xmlrpc:"has_message,omitempty"`
	HasRessourceBookingType                *Bool       `xmlrpc:"has_ressource_booking_type,omitempty"`
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
	IsFavorite                             *Bool       `xmlrpc:"is_favorite,omitempty"`
	IsInPurchaseOrder                      *Bool       `xmlrpc:"is_in_purchase_order,omitempty"`
	IsManagementFee                        *Bool       `xmlrpc:"is_management_fee,omitempty"`
	IsProductVariant                       *Bool       `xmlrpc:"is_product_variant,omitempty"`
	IsPublished                            *Bool       `xmlrpc:"is_published,omitempty"`
	IsSeoOptimized                         *Bool       `xmlrpc:"is_seo_optimized,omitempty"`
	IsStorable                             *Bool       `xmlrpc:"is_storable,omitempty"`
	ListPrice                              *Float      `xmlrpc:"list_price,omitempty"`
	LocationId                             *Many2One   `xmlrpc:"location_id,omitempty"`
	LotPropertiesDefinition                interface{} `xmlrpc:"lot_properties_definition,omitempty"`
	LotValuated                            *Bool       `xmlrpc:"lot_valuated,omitempty"`
	LstPrice                               *Float      `xmlrpc:"lst_price,omitempty"`
	ManualForceSaleMaxQty                  *Selection  `xmlrpc:"manual_force_sale_max_qty,omitempty"`
	ManualForceSaleMinQty                  *Selection  `xmlrpc:"manual_force_sale_min_qty,omitempty"`
	ManualSaleMaxQty                       *Float      `xmlrpc:"manual_sale_max_qty,omitempty"`
	ManualSaleMinQty                       *Float      `xmlrpc:"manual_sale_min_qty,omitempty"`
	ManualSaleMultipleQty                  *Float      `xmlrpc:"manual_sale_multiple_qty,omitempty"`
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
	MyActivityDateDeadline                 *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                                   *String     `xmlrpc:"name,omitempty"`
	NbrMovesIn                             *Int        `xmlrpc:"nbr_moves_in,omitempty"`
	NbrMovesOut                            *Int        `xmlrpc:"nbr_moves_out,omitempty"`
	NbrReorderingRules                     *Int        `xmlrpc:"nbr_reordering_rules,omitempty"`
	OptionalProductIds                     *Relation   `xmlrpc:"optional_product_ids,omitempty"`
	OrderpointIds                          *Relation   `xmlrpc:"orderpoint_ids,omitempty"`
	OutOfStockMessage                      *String     `xmlrpc:"out_of_stock_message,omitempty"`
	OutgoingQty                            *Float      `xmlrpc:"outgoing_qty,omitempty"`
	PackagingIds                           *Relation   `xmlrpc:"packaging_ids,omitempty"`
	PartnerRef                             *String     `xmlrpc:"partner_ref,omitempty"`
	PosCategIds                            *Relation   `xmlrpc:"pos_categ_ids,omitempty"`
	PriceExtra                             *Float      `xmlrpc:"price_extra,omitempty"`
	PricelistItemCount                     *Int        `xmlrpc:"pricelist_item_count,omitempty"`
	ProductAddMode                         *Selection  `xmlrpc:"product_add_mode,omitempty"`
	ProductCatalogProductIsInSaleOrder     *Bool       `xmlrpc:"product_catalog_product_is_in_sale_order,omitempty"`
	ProductDocumentCount                   *Int        `xmlrpc:"product_document_count,omitempty"`
	ProductDocumentIds                     *Relation   `xmlrpc:"product_document_ids,omitempty"`
	ProductProperties                      interface{} `xmlrpc:"product_properties,omitempty"`
	ProductTagIds                          *Relation   `xmlrpc:"product_tag_ids,omitempty"`
	ProductTemplateAttributeValueIds       *Relation   `xmlrpc:"product_template_attribute_value_ids,omitempty"`
	ProductTemplateImageIds                *Relation   `xmlrpc:"product_template_image_ids,omitempty"`
	ProductTemplateVariantValueIds         *Relation   `xmlrpc:"product_template_variant_value_ids,omitempty"`
	ProductTmplId                          *Many2One   `xmlrpc:"product_tmpl_id,omitempty"`
	ProductTooltip                         *String     `xmlrpc:"product_tooltip,omitempty"`
	ProductVariantCount                    *Int        `xmlrpc:"product_variant_count,omitempty"`
	ProductVariantId                       *Many2One   `xmlrpc:"product_variant_id,omitempty"`
	ProductVariantIds                      *Relation   `xmlrpc:"product_variant_ids,omitempty"`
	ProductVariantImageIds                 *Relation   `xmlrpc:"product_variant_image_ids,omitempty"`
	ProjectId                              *Many2One   `xmlrpc:"project_id,omitempty"`
	ProjectTemplateId                      *Many2One   `xmlrpc:"project_template_id,omitempty"`
	PropertyAccountCreditorPriceDifference *Many2One   `xmlrpc:"property_account_creditor_price_difference,omitempty"`
	PropertyAccountExpenseId               *Many2One   `xmlrpc:"property_account_expense_id,omitempty"`
	PropertyAccountIncomeId                *Many2One   `xmlrpc:"property_account_income_id,omitempty"`
	PropertyStockInventory                 *Many2One   `xmlrpc:"property_stock_inventory,omitempty"`
	PropertyStockProduction                *Many2One   `xmlrpc:"property_stock_production,omitempty"`
	PublicCategIds                         *Relation   `xmlrpc:"public_categ_ids,omitempty"`
	PublicDescription                      *String     `xmlrpc:"public_description,omitempty"`
	PurchaseLineWarn                       *Selection  `xmlrpc:"purchase_line_warn,omitempty"`
	PurchaseLineWarnMsg                    *String     `xmlrpc:"purchase_line_warn_msg,omitempty"`
	PurchaseMethod                         *Selection  `xmlrpc:"purchase_method,omitempty"`
	PurchaseOk                             *Bool       `xmlrpc:"purchase_ok,omitempty"`
	PurchaseOrderLineIds                   *Relation   `xmlrpc:"purchase_order_line_ids,omitempty"`
	PurchasedProductQty                    *Float      `xmlrpc:"purchased_product_qty,omitempty"`
	PutawayRuleIds                         *Relation   `xmlrpc:"putaway_rule_ids,omitempty"`
	QtyAvailable                           *Float      `xmlrpc:"qty_available,omitempty"`
	QuantitySvl                            *Float      `xmlrpc:"quantity_svl,omitempty"`
	RatingAvg                              *Float      `xmlrpc:"rating_avg,omitempty"`
	RatingAvgText                          *Selection  `xmlrpc:"rating_avg_text,omitempty"`
	RatingCount                            *Int        `xmlrpc:"rating_count,omitempty"`
	RatingIds                              *Relation   `xmlrpc:"rating_ids,omitempty"`
	RatingLastFeedback                     *String     `xmlrpc:"rating_last_feedback,omitempty"`
	RatingLastImage                        *String     `xmlrpc:"rating_last_image,omitempty"`
	RatingLastText                         *Selection  `xmlrpc:"rating_last_text,omitempty"`
	RatingLastValue                        *Float      `xmlrpc:"rating_last_value,omitempty"`
	RatingPercentageSatisfaction           *Float      `xmlrpc:"rating_percentage_satisfaction,omitempty"`
	ReorderingMaxQty                       *Float      `xmlrpc:"reordering_max_qty,omitempty"`
	ReorderingMinQty                       *Float      `xmlrpc:"reordering_min_qty,omitempty"`
	ResourceBookingCount                   *Int        `xmlrpc:"resource_booking_count,omitempty"`
	ResourceBookingExpiration              *Time       `xmlrpc:"resource_booking_expiration,omitempty"`
	ResourceBookingIds                     *Relation   `xmlrpc:"resource_booking_ids,omitempty"`
	ResourceBookingTimeout                 *Float      `xmlrpc:"resource_booking_timeout,omitempty"`
	ResourceBookingTypeCombinationRelId    *Many2One   `xmlrpc:"resource_booking_type_combination_rel_id,omitempty"`
	ResourceBookingTypeId                  *Many2One   `xmlrpc:"resource_booking_type_id,omitempty"`
	ResponsibleId                          *Many2One   `xmlrpc:"responsible_id,omitempty"`
	RouteFromCategIds                      *Relation   `xmlrpc:"route_from_categ_ids,omitempty"`
	RouteIds                               *Relation   `xmlrpc:"route_ids,omitempty"`
	SaleDelay                              *Int        `xmlrpc:"sale_delay,omitempty"`
	SaleLineWarn                           *Selection  `xmlrpc:"sale_line_warn,omitempty"`
	SaleLineWarnMsg                        *String     `xmlrpc:"sale_line_warn_msg,omitempty"`
	SaleMaxQty                             *Float      `xmlrpc:"sale_max_qty,omitempty"`
	SaleMinQty                             *Float      `xmlrpc:"sale_min_qty,omitempty"`
	SaleMultipleQty                        *Float      `xmlrpc:"sale_multiple_qty,omitempty"`
	SaleOk                                 *Bool       `xmlrpc:"sale_ok,omitempty"`
	SalesCount                             *Float      `xmlrpc:"sales_count,omitempty"`
	SellerIds                              *Relation   `xmlrpc:"seller_ids,omitempty"`
	SeoName                                *String     `xmlrpc:"seo_name,omitempty"`
	Sequence                               *Int        `xmlrpc:"sequence,omitempty"`
	ServicePolicy                          *Selection  `xmlrpc:"service_policy,omitempty"`
	ServiceToPurchase                      *Bool       `xmlrpc:"service_to_purchase,omitempty"`
	ServiceTracking                        *Selection  `xmlrpc:"service_tracking,omitempty"`
	ServiceType                            *Selection  `xmlrpc:"service_type,omitempty"`
	ServiceUpsellThreshold                 *Float      `xmlrpc:"service_upsell_threshold,omitempty"`
	ServiceUpsellThresholdRatio            *String     `xmlrpc:"service_upsell_threshold_ratio,omitempty"`
	ShowAvailability                       *Bool       `xmlrpc:"show_availability,omitempty"`
	ShowForecastedQtyStatusButton          *Bool       `xmlrpc:"show_forecasted_qty_status_button,omitempty"`
	ShowOnHandQtyStatusButton              *Bool       `xmlrpc:"show_on_hand_qty_status_button,omitempty"`
	StandardPrice                          *Float      `xmlrpc:"standard_price,omitempty"`
	StockMoveIds                           *Relation   `xmlrpc:"stock_move_ids,omitempty"`
	StockNotificationPartnerIds            *Relation   `xmlrpc:"stock_notification_partner_ids,omitempty"`
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
	UomCategoryId                          *Many2One   `xmlrpc:"uom_category_id,omitempty"`
	UomId                                  *Many2One   `xmlrpc:"uom_id,omitempty"`
	UomName                                *String     `xmlrpc:"uom_name,omitempty"`
	UomPoId                                *Many2One   `xmlrpc:"uom_po_id,omitempty"`
	ValidEan                               *Bool       `xmlrpc:"valid_ean,omitempty"`
	ValidProductTemplateAttributeLineIds   *Relation   `xmlrpc:"valid_product_template_attribute_line_ids,omitempty"`
	Valuation                              *Selection  `xmlrpc:"valuation,omitempty"`
	ValueSvl                               *Float      `xmlrpc:"value_svl,omitempty"`
	VariantRibbonId                        *Many2One   `xmlrpc:"variant_ribbon_id,omitempty"`
	VariantSellerIds                       *Relation   `xmlrpc:"variant_seller_ids,omitempty"`
	VirtualAvailable                       *Float      `xmlrpc:"virtual_available,omitempty"`
	VisibleExpensePolicy                   *Bool       `xmlrpc:"visible_expense_policy,omitempty"`
	Volume                                 *Float      `xmlrpc:"volume,omitempty"`
	VolumeUomName                          *String     `xmlrpc:"volume_uom_name,omitempty"`
	WarehouseId                            *Many2One   `xmlrpc:"warehouse_id,omitempty"`
	WebsiteDescription                     *String     `xmlrpc:"website_description,omitempty"`
	WebsiteId                              *Many2One   `xmlrpc:"website_id,omitempty"`
	WebsiteMessageIds                      *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WebsiteMetaDescription                 *String     `xmlrpc:"website_meta_description,omitempty"`
	WebsiteMetaKeywords                    *String     `xmlrpc:"website_meta_keywords,omitempty"`
	WebsiteMetaOgImg                       *String     `xmlrpc:"website_meta_og_img,omitempty"`
	WebsiteMetaTitle                       *String     `xmlrpc:"website_meta_title,omitempty"`
	WebsitePublished                       *Bool       `xmlrpc:"website_published,omitempty"`
	WebsiteRibbonId                        *Many2One   `xmlrpc:"website_ribbon_id,omitempty"`
	WebsiteSequence                        *Int        `xmlrpc:"website_sequence,omitempty"`
	WebsiteSizeX                           *Int        `xmlrpc:"website_size_x,omitempty"`
	WebsiteSizeY                           *Int        `xmlrpc:"website_size_y,omitempty"`
	WebsiteUrl                             *String     `xmlrpc:"website_url,omitempty"`
	Weight                                 *Float      `xmlrpc:"weight,omitempty"`
	WeightUomName                          *String     `xmlrpc:"weight_uom_name,omitempty"`
	WristbandsPrinting                     *Bool       `xmlrpc:"wristbands_printing,omitempty"`
	WriteDate                              *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                               *Many2One   `xmlrpc:"write_uid,omitempty"`
	XStudioEquipement                      *Relation   `xmlrpc:"x_studio_equipement,omitempty"`
	XStudioEquipement1                     *Relation   `xmlrpc:"x_studio_equipement_1,omitempty"`
	XStudioMachines                        *Relation   `xmlrpc:"x_studio_machines,omitempty"`
	XStudioMany2ManyField1V0L8             *Relation   `xmlrpc:"x_studio_many2many_field_1v0l8,omitempty"`
	XStudioMany2ManyField8Rt1J4Kj06A3      *Relation   `xmlrpc:"x_studio_many2many_field_8rt_1j4kj06a3,omitempty"`
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
