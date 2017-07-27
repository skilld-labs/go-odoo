package types

import (
	"time"
)

type ProductTemplate struct {
	LastUpdate                             time.Time `xmlrpc:"__last_update"`
	Active                                 bool      `xmlrpc:"active"`
	AttributeLineIds                       []int64   `xmlrpc:"attribute_line_ids"`
	Barcode                                string    `xmlrpc:"barcode"`
	CategId                                Many2One  `xmlrpc:"categ_id"`
	Color                                  int64     `xmlrpc:"color"`
	CompanyId                              Many2One  `xmlrpc:"company_id"`
	CostMethod                             string    `xmlrpc:"cost_method"`
	CreateDate                             time.Time `xmlrpc:"create_date"`
	CreateUid                              Many2One  `xmlrpc:"create_uid"`
	CurrencyId                             Many2One  `xmlrpc:"currency_id"`
	DefaultCode                            string    `xmlrpc:"default_code"`
	Description                            string    `xmlrpc:"description"`
	DescriptionPicking                     string    `xmlrpc:"description_picking"`
	DescriptionPurchase                    string    `xmlrpc:"description_purchase"`
	DescriptionSale                        string    `xmlrpc:"description_sale"`
	DisplayName                            string    `xmlrpc:"display_name"`
	ExpensePolicy                          string    `xmlrpc:"expense_policy"`
	Id                                     int64     `xmlrpc:"id"`
	Image                                  string    `xmlrpc:"image"`
	ImageMedium                            string    `xmlrpc:"image_medium"`
	ImageSmall                             string    `xmlrpc:"image_small"`
	IncomingQty                            float64   `xmlrpc:"incoming_qty"`
	InvoicePolicy                          string    `xmlrpc:"invoice_policy"`
	ItemIds                                []int64   `xmlrpc:"item_ids"`
	ListPrice                              float64   `xmlrpc:"list_price"`
	LocationId                             Many2One  `xmlrpc:"location_id"`
	LstPrice                               float64   `xmlrpc:"lst_price"`
	MessageChannelIds                      []int64   `xmlrpc:"message_channel_ids"`
	MessageFollowerIds                     []int64   `xmlrpc:"message_follower_ids"`
	MessageIds                             []int64   `xmlrpc:"message_ids"`
	MessageIsFollower                      bool      `xmlrpc:"message_is_follower"`
	MessageLastPost                        time.Time `xmlrpc:"message_last_post"`
	MessageNeedaction                      bool      `xmlrpc:"message_needaction"`
	MessageNeedactionCounter               int64     `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds                      []int64   `xmlrpc:"message_partner_ids"`
	MessageUnread                          bool      `xmlrpc:"message_unread"`
	MessageUnreadCounter                   int64     `xmlrpc:"message_unread_counter"`
	Name                                   string    `xmlrpc:"name"`
	NbrReorderingRules                     int64     `xmlrpc:"nbr_reordering_rules"`
	OutgoingQty                            float64   `xmlrpc:"outgoing_qty"`
	PackagingIds                           []int64   `xmlrpc:"packaging_ids"`
	Price                                  float64   `xmlrpc:"price"`
	PricelistId                            Many2One  `xmlrpc:"pricelist_id"`
	ProductVariantCount                    int64     `xmlrpc:"product_variant_count"`
	ProductVariantId                       Many2One  `xmlrpc:"product_variant_id"`
	ProductVariantIds                      []int64   `xmlrpc:"product_variant_ids"`
	ProjectId                              Many2One  `xmlrpc:"project_id"`
	PropertyAccountCreditorPriceDifference Many2One  `xmlrpc:"property_account_creditor_price_difference"`
	PropertyAccountExpenseId               Many2One  `xmlrpc:"property_account_expense_id"`
	PropertyAccountIncomeId                Many2One  `xmlrpc:"property_account_income_id"`
	PropertyCostMethod                     string    `xmlrpc:"property_cost_method"`
	PropertyStockAccountInput              Many2One  `xmlrpc:"property_stock_account_input"`
	PropertyStockAccountOutput             Many2One  `xmlrpc:"property_stock_account_output"`
	PropertyStockInventory                 Many2One  `xmlrpc:"property_stock_inventory"`
	PropertyStockProcurement               Many2One  `xmlrpc:"property_stock_procurement"`
	PropertyStockProduction                Many2One  `xmlrpc:"property_stock_production"`
	PropertyValuation                      string    `xmlrpc:"property_valuation"`
	PurchaseCount                          int64     `xmlrpc:"purchase_count"`
	PurchaseLineWarn                       string    `xmlrpc:"purchase_line_warn"`
	PurchaseLineWarnMsg                    string    `xmlrpc:"purchase_line_warn_msg"`
	PurchaseMethod                         string    `xmlrpc:"purchase_method"`
	PurchaseOk                             bool      `xmlrpc:"purchase_ok"`
	QtyAvailable                           float64   `xmlrpc:"qty_available"`
	Rental                                 bool      `xmlrpc:"rental"`
	ReorderingMaxQty                       float64   `xmlrpc:"reordering_max_qty"`
	ReorderingMinQty                       float64   `xmlrpc:"reordering_min_qty"`
	RouteFromCategIds                      []int64   `xmlrpc:"route_from_categ_ids"`
	RouteIds                               []int64   `xmlrpc:"route_ids"`
	SaleDelay                              float64   `xmlrpc:"sale_delay"`
	SaleLineWarn                           string    `xmlrpc:"sale_line_warn"`
	SaleLineWarnMsg                        string    `xmlrpc:"sale_line_warn_msg"`
	SaleOk                                 bool      `xmlrpc:"sale_ok"`
	SalesCount                             int64     `xmlrpc:"sales_count"`
	SellerIds                              []int64   `xmlrpc:"seller_ids"`
	Sequence                               int64     `xmlrpc:"sequence"`
	StandardPrice                          float64   `xmlrpc:"standard_price"`
	SupplierTaxesId                        []int64   `xmlrpc:"supplier_taxes_id"`
	TaxesId                                []int64   `xmlrpc:"taxes_id"`
	TrackService                           string    `xmlrpc:"track_service"`
	Tracking                               string    `xmlrpc:"tracking"`
	Type                                   string    `xmlrpc:"type"`
	UomId                                  Many2One  `xmlrpc:"uom_id"`
	UomPoId                                Many2One  `xmlrpc:"uom_po_id"`
	Valuation                              string    `xmlrpc:"valuation"`
	VirtualAvailable                       float64   `xmlrpc:"virtual_available"`
	Volume                                 float64   `xmlrpc:"volume"`
	WarehouseId                            Many2One  `xmlrpc:"warehouse_id"`
	Warranty                               float64   `xmlrpc:"warranty"`
	Weight                                 float64   `xmlrpc:"weight"`
	WriteDate                              time.Time `xmlrpc:"write_date"`
	WriteUid                               Many2One  `xmlrpc:"write_uid"`
}

type ProductTemplateNil struct {
	LastUpdate                             interface{} `xmlrpc:"__last_update"`
	Active                                 bool        `xmlrpc:"active"`
	AttributeLineIds                       interface{} `xmlrpc:"attribute_line_ids"`
	Barcode                                interface{} `xmlrpc:"barcode"`
	CategId                                interface{} `xmlrpc:"categ_id"`
	Color                                  interface{} `xmlrpc:"color"`
	CompanyId                              interface{} `xmlrpc:"company_id"`
	CostMethod                             interface{} `xmlrpc:"cost_method"`
	CreateDate                             interface{} `xmlrpc:"create_date"`
	CreateUid                              interface{} `xmlrpc:"create_uid"`
	CurrencyId                             interface{} `xmlrpc:"currency_id"`
	DefaultCode                            interface{} `xmlrpc:"default_code"`
	Description                            interface{} `xmlrpc:"description"`
	DescriptionPicking                     interface{} `xmlrpc:"description_picking"`
	DescriptionPurchase                    interface{} `xmlrpc:"description_purchase"`
	DescriptionSale                        interface{} `xmlrpc:"description_sale"`
	DisplayName                            interface{} `xmlrpc:"display_name"`
	ExpensePolicy                          interface{} `xmlrpc:"expense_policy"`
	Id                                     interface{} `xmlrpc:"id"`
	Image                                  interface{} `xmlrpc:"image"`
	ImageMedium                            interface{} `xmlrpc:"image_medium"`
	ImageSmall                             interface{} `xmlrpc:"image_small"`
	IncomingQty                            interface{} `xmlrpc:"incoming_qty"`
	InvoicePolicy                          interface{} `xmlrpc:"invoice_policy"`
	ItemIds                                interface{} `xmlrpc:"item_ids"`
	ListPrice                              interface{} `xmlrpc:"list_price"`
	LocationId                             interface{} `xmlrpc:"location_id"`
	LstPrice                               interface{} `xmlrpc:"lst_price"`
	MessageChannelIds                      interface{} `xmlrpc:"message_channel_ids"`
	MessageFollowerIds                     interface{} `xmlrpc:"message_follower_ids"`
	MessageIds                             interface{} `xmlrpc:"message_ids"`
	MessageIsFollower                      bool        `xmlrpc:"message_is_follower"`
	MessageLastPost                        interface{} `xmlrpc:"message_last_post"`
	MessageNeedaction                      bool        `xmlrpc:"message_needaction"`
	MessageNeedactionCounter               interface{} `xmlrpc:"message_needaction_counter"`
	MessagePartnerIds                      interface{} `xmlrpc:"message_partner_ids"`
	MessageUnread                          bool        `xmlrpc:"message_unread"`
	MessageUnreadCounter                   interface{} `xmlrpc:"message_unread_counter"`
	Name                                   interface{} `xmlrpc:"name"`
	NbrReorderingRules                     interface{} `xmlrpc:"nbr_reordering_rules"`
	OutgoingQty                            interface{} `xmlrpc:"outgoing_qty"`
	PackagingIds                           interface{} `xmlrpc:"packaging_ids"`
	Price                                  interface{} `xmlrpc:"price"`
	PricelistId                            interface{} `xmlrpc:"pricelist_id"`
	ProductVariantCount                    interface{} `xmlrpc:"product_variant_count"`
	ProductVariantId                       interface{} `xmlrpc:"product_variant_id"`
	ProductVariantIds                      interface{} `xmlrpc:"product_variant_ids"`
	ProjectId                              interface{} `xmlrpc:"project_id"`
	PropertyAccountCreditorPriceDifference interface{} `xmlrpc:"property_account_creditor_price_difference"`
	PropertyAccountExpenseId               interface{} `xmlrpc:"property_account_expense_id"`
	PropertyAccountIncomeId                interface{} `xmlrpc:"property_account_income_id"`
	PropertyCostMethod                     interface{} `xmlrpc:"property_cost_method"`
	PropertyStockAccountInput              interface{} `xmlrpc:"property_stock_account_input"`
	PropertyStockAccountOutput             interface{} `xmlrpc:"property_stock_account_output"`
	PropertyStockInventory                 interface{} `xmlrpc:"property_stock_inventory"`
	PropertyStockProcurement               interface{} `xmlrpc:"property_stock_procurement"`
	PropertyStockProduction                interface{} `xmlrpc:"property_stock_production"`
	PropertyValuation                      interface{} `xmlrpc:"property_valuation"`
	PurchaseCount                          interface{} `xmlrpc:"purchase_count"`
	PurchaseLineWarn                       interface{} `xmlrpc:"purchase_line_warn"`
	PurchaseLineWarnMsg                    interface{} `xmlrpc:"purchase_line_warn_msg"`
	PurchaseMethod                         interface{} `xmlrpc:"purchase_method"`
	PurchaseOk                             bool        `xmlrpc:"purchase_ok"`
	QtyAvailable                           interface{} `xmlrpc:"qty_available"`
	Rental                                 bool        `xmlrpc:"rental"`
	ReorderingMaxQty                       interface{} `xmlrpc:"reordering_max_qty"`
	ReorderingMinQty                       interface{} `xmlrpc:"reordering_min_qty"`
	RouteFromCategIds                      interface{} `xmlrpc:"route_from_categ_ids"`
	RouteIds                               interface{} `xmlrpc:"route_ids"`
	SaleDelay                              interface{} `xmlrpc:"sale_delay"`
	SaleLineWarn                           interface{} `xmlrpc:"sale_line_warn"`
	SaleLineWarnMsg                        interface{} `xmlrpc:"sale_line_warn_msg"`
	SaleOk                                 bool        `xmlrpc:"sale_ok"`
	SalesCount                             interface{} `xmlrpc:"sales_count"`
	SellerIds                              interface{} `xmlrpc:"seller_ids"`
	Sequence                               interface{} `xmlrpc:"sequence"`
	StandardPrice                          interface{} `xmlrpc:"standard_price"`
	SupplierTaxesId                        interface{} `xmlrpc:"supplier_taxes_id"`
	TaxesId                                interface{} `xmlrpc:"taxes_id"`
	TrackService                           interface{} `xmlrpc:"track_service"`
	Tracking                               interface{} `xmlrpc:"tracking"`
	Type                                   interface{} `xmlrpc:"type"`
	UomId                                  interface{} `xmlrpc:"uom_id"`
	UomPoId                                interface{} `xmlrpc:"uom_po_id"`
	Valuation                              interface{} `xmlrpc:"valuation"`
	VirtualAvailable                       interface{} `xmlrpc:"virtual_available"`
	Volume                                 interface{} `xmlrpc:"volume"`
	WarehouseId                            interface{} `xmlrpc:"warehouse_id"`
	Warranty                               interface{} `xmlrpc:"warranty"`
	Weight                                 interface{} `xmlrpc:"weight"`
	WriteDate                              interface{} `xmlrpc:"write_date"`
	WriteUid                               interface{} `xmlrpc:"write_uid"`
}

var ProductTemplateModel string = "product.template"

type ProductTemplates []ProductTemplate

type ProductTemplatesNil []ProductTemplateNil

func (s *ProductTemplate) NilableType_() interface{} {
	return &ProductTemplateNil{}
}

func (ns *ProductTemplateNil) Type_() interface{} {
	s := &ProductTemplate{}
	return load(ns, s)
}

func (s *ProductTemplates) NilableType_() interface{} {
	return &ProductTemplatesNil{}
}

func (ns *ProductTemplatesNil) Type_() interface{} {
	s := &ProductTemplates{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductTemplate))
	}
	return s
}
