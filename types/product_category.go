package types

import (
	"time"
)

type ProductCategory struct {
	LastUpdate                                  time.Time `xmlrpc:"__last_update"`
	ChildId                                     []int64   `xmlrpc:"child_id"`
	CreateDate                                  time.Time `xmlrpc:"create_date"`
	CreateUid                                   Many2One  `xmlrpc:"create_uid"`
	DisplayName                                 string    `xmlrpc:"display_name"`
	Id                                          int64     `xmlrpc:"id"`
	Name                                        string    `xmlrpc:"name"`
	ParentId                                    Many2One  `xmlrpc:"parent_id"`
	ParentLeft                                  int64     `xmlrpc:"parent_left"`
	ParentRight                                 int64     `xmlrpc:"parent_right"`
	ProductCount                                int64     `xmlrpc:"product_count"`
	PropertyAccountCreditorPriceDifferenceCateg Many2One  `xmlrpc:"property_account_creditor_price_difference_categ"`
	PropertyAccountExpenseCategId               Many2One  `xmlrpc:"property_account_expense_categ_id"`
	PropertyAccountIncomeCategId                Many2One  `xmlrpc:"property_account_income_categ_id"`
	PropertyCostMethod                          string    `xmlrpc:"property_cost_method"`
	PropertyStockAccountInputCategId            Many2One  `xmlrpc:"property_stock_account_input_categ_id"`
	PropertyStockAccountOutputCategId           Many2One  `xmlrpc:"property_stock_account_output_categ_id"`
	PropertyStockJournal                        Many2One  `xmlrpc:"property_stock_journal"`
	PropertyStockValuationAccountId             Many2One  `xmlrpc:"property_stock_valuation_account_id"`
	PropertyValuation                           string    `xmlrpc:"property_valuation"`
	RemovalStrategyId                           Many2One  `xmlrpc:"removal_strategy_id"`
	RouteIds                                    []int64   `xmlrpc:"route_ids"`
	TotalRouteIds                               []int64   `xmlrpc:"total_route_ids"`
	Type                                        string    `xmlrpc:"type"`
	WriteDate                                   time.Time `xmlrpc:"write_date"`
	WriteUid                                    Many2One  `xmlrpc:"write_uid"`
}

type ProductCategoryNil struct {
	LastUpdate                                  interface{} `xmlrpc:"__last_update"`
	ChildId                                     interface{} `xmlrpc:"child_id"`
	CreateDate                                  interface{} `xmlrpc:"create_date"`
	CreateUid                                   interface{} `xmlrpc:"create_uid"`
	DisplayName                                 interface{} `xmlrpc:"display_name"`
	Id                                          interface{} `xmlrpc:"id"`
	Name                                        interface{} `xmlrpc:"name"`
	ParentId                                    interface{} `xmlrpc:"parent_id"`
	ParentLeft                                  interface{} `xmlrpc:"parent_left"`
	ParentRight                                 interface{} `xmlrpc:"parent_right"`
	ProductCount                                interface{} `xmlrpc:"product_count"`
	PropertyAccountCreditorPriceDifferenceCateg interface{} `xmlrpc:"property_account_creditor_price_difference_categ"`
	PropertyAccountExpenseCategId               interface{} `xmlrpc:"property_account_expense_categ_id"`
	PropertyAccountIncomeCategId                interface{} `xmlrpc:"property_account_income_categ_id"`
	PropertyCostMethod                          interface{} `xmlrpc:"property_cost_method"`
	PropertyStockAccountInputCategId            interface{} `xmlrpc:"property_stock_account_input_categ_id"`
	PropertyStockAccountOutputCategId           interface{} `xmlrpc:"property_stock_account_output_categ_id"`
	PropertyStockJournal                        interface{} `xmlrpc:"property_stock_journal"`
	PropertyStockValuationAccountId             interface{} `xmlrpc:"property_stock_valuation_account_id"`
	PropertyValuation                           interface{} `xmlrpc:"property_valuation"`
	RemovalStrategyId                           interface{} `xmlrpc:"removal_strategy_id"`
	RouteIds                                    interface{} `xmlrpc:"route_ids"`
	TotalRouteIds                               interface{} `xmlrpc:"total_route_ids"`
	Type                                        interface{} `xmlrpc:"type"`
	WriteDate                                   interface{} `xmlrpc:"write_date"`
	WriteUid                                    interface{} `xmlrpc:"write_uid"`
}

var ProductCategoryModel string = "product.category"

type ProductCategorys []ProductCategory

type ProductCategorysNil []ProductCategoryNil

func (s *ProductCategory) NilableType_() interface{} {
	return &ProductCategoryNil{}
}

func (ns *ProductCategoryNil) Type_() interface{} {
	s := &ProductCategory{}
	return load(ns, s)
}

func (s *ProductCategorys) NilableType_() interface{} {
	return &ProductCategorysNil{}
}

func (ns *ProductCategorysNil) Type_() interface{} {
	s := &ProductCategorys{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*ProductCategory))
	}
	return s
}
