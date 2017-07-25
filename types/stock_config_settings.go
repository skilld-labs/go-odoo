package types

import (
	"time"
)

type StockConfigSettings struct {
	LastUpdate                     time.Time `xmlrpc:"__last_update"`
	CompanyId                      Many2One  `xmlrpc:"company_id"`
	CreateDate                     time.Time `xmlrpc:"create_date"`
	CreateUid                      Many2One  `xmlrpc:"create_uid"`
	DecimalPrecision               int64     `xmlrpc:"decimal_precision"`
	DisplayName                    string    `xmlrpc:"display_name"`
	GroupProductVariant            string    `xmlrpc:"group_product_variant"`
	GroupStockAdvLocation          string    `xmlrpc:"group_stock_adv_location"`
	GroupStockInventoryValuation   string    `xmlrpc:"group_stock_inventory_valuation"`
	GroupStockMultiLocations       bool      `xmlrpc:"group_stock_multi_locations"`
	GroupStockMultiWarehouses      bool      `xmlrpc:"group_stock_multi_warehouses"`
	GroupStockPackaging            string    `xmlrpc:"group_stock_packaging"`
	GroupStockProductionLot        string    `xmlrpc:"group_stock_production_lot"`
	GroupStockTrackingLot          string    `xmlrpc:"group_stock_tracking_lot"`
	GroupStockTrackingOwner        string    `xmlrpc:"group_stock_tracking_owner"`
	GroupUom                       string    `xmlrpc:"group_uom"`
	GroupWarningStock              string    `xmlrpc:"group_warning_stock"`
	Id                             int64     `xmlrpc:"id"`
	ModuleDeliveryDhl              bool      `xmlrpc:"module_delivery_dhl"`
	ModuleDeliveryFedex            bool      `xmlrpc:"module_delivery_fedex"`
	ModuleDeliveryTemando          bool      `xmlrpc:"module_delivery_temando"`
	ModuleDeliveryUps              bool      `xmlrpc:"module_delivery_ups"`
	ModuleDeliveryUsps             bool      `xmlrpc:"module_delivery_usps"`
	ModuleProcurementJit           string    `xmlrpc:"module_procurement_jit"`
	ModuleProductExpiry            string    `xmlrpc:"module_product_expiry"`
	ModuleQuality                  bool      `xmlrpc:"module_quality"`
	ModuleStockBarcode             bool      `xmlrpc:"module_stock_barcode"`
	ModuleStockCalendar            string    `xmlrpc:"module_stock_calendar"`
	ModuleStockDropshipping        string    `xmlrpc:"module_stock_dropshipping"`
	ModuleStockLandedCosts         string    `xmlrpc:"module_stock_landed_costs"`
	ModuleStockPickingWave         string    `xmlrpc:"module_stock_picking_wave"`
	PropagationMinimumDelta        int64     `xmlrpc:"propagation_minimum_delta"`
	WarehouseAndLocationUsageLevel string    `xmlrpc:"warehouse_and_location_usage_level"`
	WriteDate                      time.Time `xmlrpc:"write_date"`
	WriteUid                       Many2One  `xmlrpc:"write_uid"`
}

type StockConfigSettingsNil struct {
	LastUpdate                     interface{} `xmlrpc:"__last_update"`
	CompanyId                      interface{} `xmlrpc:"company_id"`
	CreateDate                     interface{} `xmlrpc:"create_date"`
	CreateUid                      interface{} `xmlrpc:"create_uid"`
	DecimalPrecision               interface{} `xmlrpc:"decimal_precision"`
	DisplayName                    interface{} `xmlrpc:"display_name"`
	GroupProductVariant            interface{} `xmlrpc:"group_product_variant"`
	GroupStockAdvLocation          interface{} `xmlrpc:"group_stock_adv_location"`
	GroupStockInventoryValuation   interface{} `xmlrpc:"group_stock_inventory_valuation"`
	GroupStockMultiLocations       bool        `xmlrpc:"group_stock_multi_locations"`
	GroupStockMultiWarehouses      bool        `xmlrpc:"group_stock_multi_warehouses"`
	GroupStockPackaging            interface{} `xmlrpc:"group_stock_packaging"`
	GroupStockProductionLot        interface{} `xmlrpc:"group_stock_production_lot"`
	GroupStockTrackingLot          interface{} `xmlrpc:"group_stock_tracking_lot"`
	GroupStockTrackingOwner        interface{} `xmlrpc:"group_stock_tracking_owner"`
	GroupUom                       interface{} `xmlrpc:"group_uom"`
	GroupWarningStock              interface{} `xmlrpc:"group_warning_stock"`
	Id                             interface{} `xmlrpc:"id"`
	ModuleDeliveryDhl              bool        `xmlrpc:"module_delivery_dhl"`
	ModuleDeliveryFedex            bool        `xmlrpc:"module_delivery_fedex"`
	ModuleDeliveryTemando          bool        `xmlrpc:"module_delivery_temando"`
	ModuleDeliveryUps              bool        `xmlrpc:"module_delivery_ups"`
	ModuleDeliveryUsps             bool        `xmlrpc:"module_delivery_usps"`
	ModuleProcurementJit           interface{} `xmlrpc:"module_procurement_jit"`
	ModuleProductExpiry            interface{} `xmlrpc:"module_product_expiry"`
	ModuleQuality                  bool        `xmlrpc:"module_quality"`
	ModuleStockBarcode             bool        `xmlrpc:"module_stock_barcode"`
	ModuleStockCalendar            interface{} `xmlrpc:"module_stock_calendar"`
	ModuleStockDropshipping        interface{} `xmlrpc:"module_stock_dropshipping"`
	ModuleStockLandedCosts         interface{} `xmlrpc:"module_stock_landed_costs"`
	ModuleStockPickingWave         interface{} `xmlrpc:"module_stock_picking_wave"`
	PropagationMinimumDelta        interface{} `xmlrpc:"propagation_minimum_delta"`
	WarehouseAndLocationUsageLevel interface{} `xmlrpc:"warehouse_and_location_usage_level"`
	WriteDate                      interface{} `xmlrpc:"write_date"`
	WriteUid                       interface{} `xmlrpc:"write_uid"`
}

var StockConfigSettingsModel string = "stock.config.settings"

type StockConfigSettingss []StockConfigSettings

type StockConfigSettingssNil []StockConfigSettingsNil

func (s *StockConfigSettings) NilableType_() interface{} {
	return &StockConfigSettingsNil{}
}

func (ns *StockConfigSettingsNil) Type_() interface{} {
	s := &StockConfigSettings{}
	return load(ns, s)
}

func (s *StockConfigSettingss) NilableType_() interface{} {
	return &StockConfigSettingssNil{}
}

func (ns *StockConfigSettingssNil) Type_() interface{} {
	s := &StockConfigSettingss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*StockConfigSettings))
	}
	return s
}
