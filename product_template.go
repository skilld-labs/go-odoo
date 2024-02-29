package odoo

// ProductTemplate represents product.template model.
type ProductTemplate struct {
	LastUpdate                             *Time      `xmlrpc:"__last_update,omitempty"`
	Active                                 *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline                   *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds                            *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                          *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                        *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId                         *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                         *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AttributeLineIds                       *Relation  `xmlrpc:"attribute_line_ids,omitempty"`
	Barcode                                *String    `xmlrpc:"barcode,omitempty"`
	CategId                                *Many2One  `xmlrpc:"categ_id,omitempty"`
	Color                                  *Int       `xmlrpc:"color,omitempty"`
	CompanyId                              *Many2One  `xmlrpc:"company_id,omitempty"`
	CostCurrencyId                         *Many2One  `xmlrpc:"cost_currency_id,omitempty"`
	CostMethod                             *String    `xmlrpc:"cost_method,omitempty"`
	CreateDate                             *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                              *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                             *Many2One  `xmlrpc:"currency_id,omitempty"`
	DefaultCode                            *String    `xmlrpc:"default_code,omitempty"`
	Description                            *String    `xmlrpc:"description,omitempty"`
	DescriptionPicking                     *String    `xmlrpc:"description_picking,omitempty"`
	DescriptionPickingin                   *String    `xmlrpc:"description_pickingin,omitempty"`
	DescriptionPickingout                  *String    `xmlrpc:"description_pickingout,omitempty"`
	DescriptionPurchase                    *String    `xmlrpc:"description_purchase,omitempty"`
	DescriptionSale                        *String    `xmlrpc:"description_sale,omitempty"`
	DisplayName                            *String    `xmlrpc:"display_name,omitempty"`
	ExpensePolicy                          *Selection `xmlrpc:"expense_policy,omitempty"`
	Id                                     *Int       `xmlrpc:"id,omitempty"`
	Image                                  *String    `xmlrpc:"image,omitempty"`
	ImageMedium                            *String    `xmlrpc:"image_medium,omitempty"`
	ImageSmall                             *String    `xmlrpc:"image_small,omitempty"`
	IncomingQty                            *Float     `xmlrpc:"incoming_qty,omitempty"`
	InvoicePolicy                          *Selection `xmlrpc:"invoice_policy,omitempty"`
	IsProductVariant                       *Bool      `xmlrpc:"is_product_variant,omitempty"`
	ItemIds                                *Relation  `xmlrpc:"item_ids,omitempty"`
	ListPrice                              *Float     `xmlrpc:"list_price,omitempty"`
	LocationId                             *Many2One  `xmlrpc:"location_id,omitempty"`
	LstPrice                               *Float     `xmlrpc:"lst_price,omitempty"`
	MessageChannelIds                      *Relation  `xmlrpc:"message_channel_ids,omitempty"`
	MessageFollowerIds                     *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageIds                             *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower                      *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageLastPost                        *Time      `xmlrpc:"message_last_post,omitempty"`
	MessageNeedaction                      *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter               *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds                      *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MessageUnread                          *Bool      `xmlrpc:"message_unread,omitempty"`
	MessageUnreadCounter                   *Int       `xmlrpc:"message_unread_counter,omitempty"`
	Name                                   *String    `xmlrpc:"name,omitempty"`
	NbrReorderingRules                     *Int       `xmlrpc:"nbr_reordering_rules,omitempty"`
	OutgoingQty                            *Float     `xmlrpc:"outgoing_qty,omitempty"`
	PackagingIds                           *Relation  `xmlrpc:"packaging_ids,omitempty"`
	Price                                  *Float     `xmlrpc:"price,omitempty"`
	PricelistId                            *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	ProductVariantCount                    *Int       `xmlrpc:"product_variant_count,omitempty"`
	ProductVariantId                       *Many2One  `xmlrpc:"product_variant_id,omitempty"`
	ProductVariantIds                      *Relation  `xmlrpc:"product_variant_ids,omitempty"`
	ProjectId                              *Many2One  `xmlrpc:"project_id,omitempty"`
	PropertyAccountCreditorPriceDifference *Many2One  `xmlrpc:"property_account_creditor_price_difference,omitempty"`
	PropertyAccountExpenseId               *Many2One  `xmlrpc:"property_account_expense_id,omitempty"`
	PropertyAccountIncomeId                *Many2One  `xmlrpc:"property_account_income_id,omitempty"`
	PropertyCostMethod                     *Selection `xmlrpc:"property_cost_method,omitempty"`
	PropertyStockAccountInput              *Many2One  `xmlrpc:"property_stock_account_input,omitempty"`
	PropertyStockAccountOutput             *Many2One  `xmlrpc:"property_stock_account_output,omitempty"`
	PropertyStockInventory                 *Many2One  `xmlrpc:"property_stock_inventory,omitempty"`
	PropertyStockProduction                *Many2One  `xmlrpc:"property_stock_production,omitempty"`
	PropertyValuation                      *Selection `xmlrpc:"property_valuation,omitempty"`
	PurchaseCount                          *Int       `xmlrpc:"purchase_count,omitempty"`
	PurchaseLineWarn                       *Selection `xmlrpc:"purchase_line_warn,omitempty"`
	PurchaseLineWarnMsg                    *String    `xmlrpc:"purchase_line_warn_msg,omitempty"`
	PurchaseMethod                         *Selection `xmlrpc:"purchase_method,omitempty"`
	PurchaseOk                             *Bool      `xmlrpc:"purchase_ok,omitempty"`
	QtyAvailable                           *Float     `xmlrpc:"qty_available,omitempty"`
	Rental                                 *Bool      `xmlrpc:"rental,omitempty"`
	ReorderingMaxQty                       *Float     `xmlrpc:"reordering_max_qty,omitempty"`
	ReorderingMinQty                       *Float     `xmlrpc:"reordering_min_qty,omitempty"`
	ResponsibleId                          *Many2One  `xmlrpc:"responsible_id,omitempty"`
	RouteFromCategIds                      *Relation  `xmlrpc:"route_from_categ_ids,omitempty"`
	RouteIds                               *Relation  `xmlrpc:"route_ids,omitempty"`
	SaleDelay                              *Float     `xmlrpc:"sale_delay,omitempty"`
	SaleLineWarn                           *Selection `xmlrpc:"sale_line_warn,omitempty"`
	SaleLineWarnMsg                        *String    `xmlrpc:"sale_line_warn_msg,omitempty"`
	SaleOk                                 *Bool      `xmlrpc:"sale_ok,omitempty"`
	SalesCount                             *Int       `xmlrpc:"sales_count,omitempty"`
	SellerIds                              *Relation  `xmlrpc:"seller_ids,omitempty"`
	Sequence                               *Int       `xmlrpc:"sequence,omitempty"`
	ServicePolicy                          *Selection `xmlrpc:"service_policy,omitempty"`
	ServiceTracking                        *Selection `xmlrpc:"service_tracking,omitempty"`
	ServiceType                            *Selection `xmlrpc:"service_type,omitempty"`
	StandardPrice                          *Float     `xmlrpc:"standard_price,omitempty"`
	SupplierTaxesId                        *Relation  `xmlrpc:"supplier_taxes_id,omitempty"`
	TaxesId                                *Relation  `xmlrpc:"taxes_id,omitempty"`
	Tracking                               *Selection `xmlrpc:"tracking,omitempty"`
	Type                                   *Selection `xmlrpc:"type,omitempty"`
	UomId                                  *Many2One  `xmlrpc:"uom_id,omitempty"`
	UomPoId                                *Many2One  `xmlrpc:"uom_po_id,omitempty"`
	Valuation                              *String    `xmlrpc:"valuation,omitempty"`
	VariantSellerIds                       *Relation  `xmlrpc:"variant_seller_ids,omitempty"`
	VirtualAvailable                       *Float     `xmlrpc:"virtual_available,omitempty"`
	Volume                                 *Float     `xmlrpc:"volume,omitempty"`
	WarehouseId                            *Many2One  `xmlrpc:"warehouse_id,omitempty"`
	WebsiteMessageIds                      *Relation  `xmlrpc:"website_message_ids,omitempty"`
	Weight                                 *Float     `xmlrpc:"weight,omitempty"`
	WriteDate                              *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                               *Many2One  `xmlrpc:"write_uid,omitempty"`
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

// CreateProductTemplates creates a new product.template model and returns its id.
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
