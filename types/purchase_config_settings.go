package types

import (
	"time"
)

type PurchaseConfigSettings struct {
	LastUpdate                time.Time `xmlrpc:"__last_update"`
	CompanyCurrencyId         Many2One  `xmlrpc:"company_currency_id"`
	CompanyId                 Many2One  `xmlrpc:"company_id"`
	CreateDate                time.Time `xmlrpc:"create_date"`
	CreateUid                 Many2One  `xmlrpc:"create_uid"`
	DisplayName               string    `xmlrpc:"display_name"`
	GroupCostingMethod        string    `xmlrpc:"group_costing_method"`
	GroupManageVendorPrice    string    `xmlrpc:"group_manage_vendor_price"`
	GroupProductVariant       string    `xmlrpc:"group_product_variant"`
	GroupUom                  string    `xmlrpc:"group_uom"`
	GroupWarningPurchase      string    `xmlrpc:"group_warning_purchase"`
	Id                        int64     `xmlrpc:"id"`
	ModulePurchaseRequisition string    `xmlrpc:"module_purchase_requisition"`
	ModuleStockDropshipping   string    `xmlrpc:"module_stock_dropshipping"`
	PoDoubleValidation        string    `xmlrpc:"po_double_validation"`
	PoDoubleValidationAmount  float64   `xmlrpc:"po_double_validation_amount"`
	PoLead                    float64   `xmlrpc:"po_lead"`
	PoLock                    string    `xmlrpc:"po_lock"`
	WriteDate                 time.Time `xmlrpc:"write_date"`
	WriteUid                  Many2One  `xmlrpc:"write_uid"`
}

type PurchaseConfigSettingsNil struct {
	LastUpdate                interface{} `xmlrpc:"__last_update"`
	CompanyCurrencyId         interface{} `xmlrpc:"company_currency_id"`
	CompanyId                 interface{} `xmlrpc:"company_id"`
	CreateDate                interface{} `xmlrpc:"create_date"`
	CreateUid                 interface{} `xmlrpc:"create_uid"`
	DisplayName               interface{} `xmlrpc:"display_name"`
	GroupCostingMethod        interface{} `xmlrpc:"group_costing_method"`
	GroupManageVendorPrice    interface{} `xmlrpc:"group_manage_vendor_price"`
	GroupProductVariant       interface{} `xmlrpc:"group_product_variant"`
	GroupUom                  interface{} `xmlrpc:"group_uom"`
	GroupWarningPurchase      interface{} `xmlrpc:"group_warning_purchase"`
	Id                        interface{} `xmlrpc:"id"`
	ModulePurchaseRequisition interface{} `xmlrpc:"module_purchase_requisition"`
	ModuleStockDropshipping   interface{} `xmlrpc:"module_stock_dropshipping"`
	PoDoubleValidation        interface{} `xmlrpc:"po_double_validation"`
	PoDoubleValidationAmount  interface{} `xmlrpc:"po_double_validation_amount"`
	PoLead                    interface{} `xmlrpc:"po_lead"`
	PoLock                    interface{} `xmlrpc:"po_lock"`
	WriteDate                 interface{} `xmlrpc:"write_date"`
	WriteUid                  interface{} `xmlrpc:"write_uid"`
}

var PurchaseConfigSettingsModel string = "purchase.config.settings"

type PurchaseConfigSettingss []PurchaseConfigSettings

type PurchaseConfigSettingssNil []PurchaseConfigSettingsNil

func (s *PurchaseConfigSettings) NilableType_() interface{} {
	return &PurchaseConfigSettingsNil{}
}

func (ns *PurchaseConfigSettingsNil) Type_() interface{} {
	s := &PurchaseConfigSettings{}
	return load(ns, s)
}

func (s *PurchaseConfigSettingss) NilableType_() interface{} {
	return &PurchaseConfigSettingssNil{}
}

func (ns *PurchaseConfigSettingssNil) Type_() interface{} {
	s := &PurchaseConfigSettingss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PurchaseConfigSettings))
	}
	return s
}
