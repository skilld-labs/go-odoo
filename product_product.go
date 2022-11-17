package odoo

import (
	"fmt"
)

// ProductProduct represents product.product model.
type ProductProduct struct {
	LastUpdate                             *Time      `xmlrpc:"__last_update,omitempty"`
	Active                                 *Bool      `xmlrpc:"active,omitempty"`
	ActivityDateDeadline                   *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityIds                            *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState                          *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary                        *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeId                         *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId                         *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	AttributeLineIds                       *Relation  `xmlrpc:"attribute_line_ids,omitempty"`
	AttributeValueIds                      *Relation  `xmlrpc:"attribute_value_ids,omitempty"`
	Barcode                                *String    `xmlrpc:"barcode,omitempty"`
	CategId                                *Many2One  `xmlrpc:"categ_id,omitempty"`
	Code                                   *String    `xmlrpc:"code,omitempty"`
	Color                                  *Int       `xmlrpc:"color,omitempty"`
	CompanyId                              *Many2One  `xmlrpc:"company_id,omitempty"`
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
	ImageVariant                           *String    `xmlrpc:"image_variant,omitempty"`
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
	OrderpointIds                          *Relation  `xmlrpc:"orderpoint_ids,omitempty"`
	OutgoingQty                            *Float     `xmlrpc:"outgoing_qty,omitempty"`
	PackagingIds                           *Relation  `xmlrpc:"packaging_ids,omitempty"`
	PartnerRef                             *String    `xmlrpc:"partner_ref,omitempty"`
	Price                                  *Float     `xmlrpc:"price,omitempty"`
	PriceExtra                             *Float     `xmlrpc:"price_extra,omitempty"`
	PricelistId                            *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	PricelistItemIds                       *Relation  `xmlrpc:"pricelist_item_ids,omitempty"`
	ProductTmplId                          *Many2One  `xmlrpc:"product_tmpl_id,omitempty"`
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
	QtyAtDate                              *Float     `xmlrpc:"qty_at_date,omitempty"`
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
	StockFifoManualMoveIds                 *Relation  `xmlrpc:"stock_fifo_manual_move_ids,omitempty"`
	StockFifoRealTimeAmlIds                *Relation  `xmlrpc:"stock_fifo_real_time_aml_ids,omitempty"`
	StockMoveIds                           *Relation  `xmlrpc:"stock_move_ids,omitempty"`
	StockQuantIds                          *Relation  `xmlrpc:"stock_quant_ids,omitempty"`
	StockValue                             *Float     `xmlrpc:"stock_value,omitempty"`
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
	return nil, fmt.Errorf("no product.product was found with criteria %v", criteria)
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
	return -1, fmt.Errorf("no product.product was found with criteria %v and options %v", criteria, options)
}
