package odoo

import (
	"fmt"
)

// ProductCategory represents product.category model.
type ProductCategory struct {
	LastUpdate                                  *Time      `xmlrpc:"__last_update,omptempty"`
	ChildId                                     *Relation  `xmlrpc:"child_id,omptempty"`
	CompleteName                                *String    `xmlrpc:"complete_name,omptempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omptempty"`
	Id                                          *Int       `xmlrpc:"id,omptempty"`
	Name                                        *String    `xmlrpc:"name,omptempty"`
	ParentId                                    *Many2One  `xmlrpc:"parent_id,omptempty"`
	ParentLeft                                  *Int       `xmlrpc:"parent_left,omptempty"`
	ParentRight                                 *Int       `xmlrpc:"parent_right,omptempty"`
	ProductCount                                *Int       `xmlrpc:"product_count,omptempty"`
	PropertyAccountCreditorPriceDifferenceCateg *Many2One  `xmlrpc:"property_account_creditor_price_difference_categ,omptempty"`
	PropertyAccountExpenseCategId               *Many2One  `xmlrpc:"property_account_expense_categ_id,omptempty"`
	PropertyAccountIncomeCategId                *Many2One  `xmlrpc:"property_account_income_categ_id,omptempty"`
	PropertyCostMethod                          *Selection `xmlrpc:"property_cost_method,omptempty"`
	PropertyStockAccountInputCategId            *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omptempty"`
	PropertyStockAccountOutputCategId           *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omptempty"`
	PropertyStockJournal                        *Many2One  `xmlrpc:"property_stock_journal,omptempty"`
	PropertyStockValuationAccountId             *Many2One  `xmlrpc:"property_stock_valuation_account_id,omptempty"`
	PropertyValuation                           *Selection `xmlrpc:"property_valuation,omptempty"`
	RemovalStrategyId                           *Many2One  `xmlrpc:"removal_strategy_id,omptempty"`
	RouteIds                                    *Relation  `xmlrpc:"route_ids,omptempty"`
	TotalRouteIds                               *Relation  `xmlrpc:"total_route_ids,omptempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductCategorys represents array of product.category model.
type ProductCategorys []ProductCategory

// ProductCategoryModel is the odoo model name.
const ProductCategoryModel = "product.category"

// Many2One convert ProductCategory to *Many2One.
func (pc *ProductCategory) Many2One() *Many2One {
	return NewMany2One(pc.Id.Get(), "")
}

// CreateProductCategory creates a new product.category model and returns its id.
func (c *Client) CreateProductCategory(pc *ProductCategory) (int64, error) {
	ids, err := c.CreateProductCategorys([]*ProductCategory{pc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductCategory creates a new product.category model and returns its id.
func (c *Client) CreateProductCategorys(pcs []*ProductCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcs {
		vv = append(vv, v)
	}
	return c.Create(ProductCategoryModel, vv)
}

// UpdateProductCategory updates an existing product.category record.
func (c *Client) UpdateProductCategory(pc *ProductCategory) error {
	return c.UpdateProductCategorys([]int64{pc.Id.Get()}, pc)
}

// UpdateProductCategorys updates existing product.category records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdateProductCategorys(ids []int64, pc *ProductCategory) error {
	return c.Update(ProductCategoryModel, ids, pc)
}

// DeleteProductCategory deletes an existing product.category record.
func (c *Client) DeleteProductCategory(id int64) error {
	return c.DeleteProductCategorys([]int64{id})
}

// DeleteProductCategorys deletes existing product.category records.
func (c *Client) DeleteProductCategorys(ids []int64) error {
	return c.Delete(ProductCategoryModel, ids)
}

// GetProductCategory gets product.category existing record.
func (c *Client) GetProductCategory(id int64) (*ProductCategory, error) {
	pcs, err := c.GetProductCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("id %v of product.category not found", id)
}

// GetProductCategorys gets product.category existing records.
func (c *Client) GetProductCategorys(ids []int64) (*ProductCategorys, error) {
	pcs := &ProductCategorys{}
	if err := c.Read(ProductCategoryModel, ids, nil, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindProductCategory finds product.category record by querying it with criteria.
func (c *Client) FindProductCategory(criteria *Criteria) (*ProductCategory, error) {
	pcs := &ProductCategorys{}
	if err := c.SearchRead(ProductCategoryModel, criteria, NewOptions().Limit(1), pcs); err != nil {
		return nil, err
	}
	if pcs != nil && len(*pcs) > 0 {
		return &((*pcs)[0]), nil
	}
	return nil, fmt.Errorf("product.category was not found with criteria %v", criteria)
}

// FindProductCategorys finds product.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductCategorys(criteria *Criteria, options *Options) (*ProductCategorys, error) {
	pcs := &ProductCategorys{}
	if err := c.SearchRead(ProductCategoryModel, criteria, options, pcs); err != nil {
		return nil, err
	}
	return pcs, nil
}

// FindProductCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	ids, err := c.Search(ProductCategoryModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindProductCategoryId finds record id by querying it with criteria.
func (c *Client) FindProductCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("product.category was not found with criteria %v and options %v", criteria, options)
}
