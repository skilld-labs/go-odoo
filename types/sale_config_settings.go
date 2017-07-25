package types

import (
	"time"
)

type SaleConfigSettings struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	AliasDomain              string    `xmlrpc:"alias_domain"`
	AliasPrefix              string    `xmlrpc:"alias_prefix"`
	AutoDoneSetting          string    `xmlrpc:"auto_done_setting"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	DefaultInvoicePolicy     string    `xmlrpc:"default_invoice_policy"`
	DefaultPickingPolicy     string    `xmlrpc:"default_picking_policy"`
	DepositProductIdSetting  Many2One  `xmlrpc:"deposit_product_id_setting"`
	DisplayName              string    `xmlrpc:"display_name"`
	GenerateSalesTeamAlias   bool      `xmlrpc:"generate_sales_team_alias"`
	GroupDiscountPerSoLine   string    `xmlrpc:"group_discount_per_so_line"`
	GroupDisplayIncoterm     string    `xmlrpc:"group_display_incoterm"`
	GroupMrpProperties       string    `xmlrpc:"group_mrp_properties"`
	GroupPricelistItem       bool      `xmlrpc:"group_pricelist_item"`
	GroupProductPricelist    bool      `xmlrpc:"group_product_pricelist"`
	GroupProductVariant      string    `xmlrpc:"group_product_variant"`
	GroupRouteSoLines        string    `xmlrpc:"group_route_so_lines"`
	GroupSaleDeliveryAddress string    `xmlrpc:"group_sale_delivery_address"`
	GroupSaleLayout          string    `xmlrpc:"group_sale_layout"`
	GroupSalePricelist       bool      `xmlrpc:"group_sale_pricelist"`
	GroupShowPriceSubtotal   bool      `xmlrpc:"group_show_price_subtotal"`
	GroupShowPriceTotal      bool      `xmlrpc:"group_show_price_total"`
	GroupUom                 string    `xmlrpc:"group_uom"`
	GroupUseLead             string    `xmlrpc:"group_use_lead"`
	GroupWarningSale         string    `xmlrpc:"group_warning_sale"`
	Id                       int64     `xmlrpc:"id"`
	ModuleCrmVoip            bool      `xmlrpc:"module_crm_voip"`
	ModuleDelivery           string    `xmlrpc:"module_delivery"`
	ModuleSaleContract       bool      `xmlrpc:"module_sale_contract"`
	ModuleSaleMargin         string    `xmlrpc:"module_sale_margin"`
	ModuleSaleOrderDates     string    `xmlrpc:"module_sale_order_dates"`
	ModuleWebsiteQuote       string    `xmlrpc:"module_website_quote"`
	ModuleWebsiteSaleDigital bool      `xmlrpc:"module_website_sale_digital"`
	ModuleWebsiteSign        bool      `xmlrpc:"module_website_sign"`
	SaleNote                 string    `xmlrpc:"sale_note"`
	SalePricelistSetting     string    `xmlrpc:"sale_pricelist_setting"`
	SaleShowTax              string    `xmlrpc:"sale_show_tax"`
	SecurityLead             float64   `xmlrpc:"security_lead"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type SaleConfigSettingsNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	AliasDomain              interface{} `xmlrpc:"alias_domain"`
	AliasPrefix              interface{} `xmlrpc:"alias_prefix"`
	AutoDoneSetting          interface{} `xmlrpc:"auto_done_setting"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	DefaultInvoicePolicy     interface{} `xmlrpc:"default_invoice_policy"`
	DefaultPickingPolicy     interface{} `xmlrpc:"default_picking_policy"`
	DepositProductIdSetting  interface{} `xmlrpc:"deposit_product_id_setting"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	GenerateSalesTeamAlias   bool        `xmlrpc:"generate_sales_team_alias"`
	GroupDiscountPerSoLine   interface{} `xmlrpc:"group_discount_per_so_line"`
	GroupDisplayIncoterm     interface{} `xmlrpc:"group_display_incoterm"`
	GroupMrpProperties       interface{} `xmlrpc:"group_mrp_properties"`
	GroupPricelistItem       bool        `xmlrpc:"group_pricelist_item"`
	GroupProductPricelist    bool        `xmlrpc:"group_product_pricelist"`
	GroupProductVariant      interface{} `xmlrpc:"group_product_variant"`
	GroupRouteSoLines        interface{} `xmlrpc:"group_route_so_lines"`
	GroupSaleDeliveryAddress interface{} `xmlrpc:"group_sale_delivery_address"`
	GroupSaleLayout          interface{} `xmlrpc:"group_sale_layout"`
	GroupSalePricelist       bool        `xmlrpc:"group_sale_pricelist"`
	GroupShowPriceSubtotal   bool        `xmlrpc:"group_show_price_subtotal"`
	GroupShowPriceTotal      bool        `xmlrpc:"group_show_price_total"`
	GroupUom                 interface{} `xmlrpc:"group_uom"`
	GroupUseLead             interface{} `xmlrpc:"group_use_lead"`
	GroupWarningSale         interface{} `xmlrpc:"group_warning_sale"`
	Id                       interface{} `xmlrpc:"id"`
	ModuleCrmVoip            bool        `xmlrpc:"module_crm_voip"`
	ModuleDelivery           interface{} `xmlrpc:"module_delivery"`
	ModuleSaleContract       bool        `xmlrpc:"module_sale_contract"`
	ModuleSaleMargin         interface{} `xmlrpc:"module_sale_margin"`
	ModuleSaleOrderDates     interface{} `xmlrpc:"module_sale_order_dates"`
	ModuleWebsiteQuote       interface{} `xmlrpc:"module_website_quote"`
	ModuleWebsiteSaleDigital bool        `xmlrpc:"module_website_sale_digital"`
	ModuleWebsiteSign        bool        `xmlrpc:"module_website_sign"`
	SaleNote                 interface{} `xmlrpc:"sale_note"`
	SalePricelistSetting     interface{} `xmlrpc:"sale_pricelist_setting"`
	SaleShowTax              interface{} `xmlrpc:"sale_show_tax"`
	SecurityLead             interface{} `xmlrpc:"security_lead"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var SaleConfigSettingsModel string = "sale.config.settings"

type SaleConfigSettingss []SaleConfigSettings

type SaleConfigSettingssNil []SaleConfigSettingsNil

func (s *SaleConfigSettings) NilableType_() interface{} {
	return &SaleConfigSettingsNil{}
}

func (ns *SaleConfigSettingsNil) Type_() interface{} {
	s := &SaleConfigSettings{}
	return load(ns, s)
}

func (s *SaleConfigSettingss) NilableType_() interface{} {
	return &SaleConfigSettingssNil{}
}

func (ns *SaleConfigSettingssNil) Type_() interface{} {
	s := &SaleConfigSettingss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*SaleConfigSettings))
	}
	return s
}
