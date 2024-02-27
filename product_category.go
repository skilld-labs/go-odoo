package odoo

// ProductCategory represents product.category model.
type ProductCategory struct {
	LastUpdate                                  *Time      `xmlrpc:"__last_update,omitempty"`
	ChildId                                     *Relation  `xmlrpc:"child_id,omitempty"`
	CompleteName                                *String    `xmlrpc:"complete_name,omitempty"`
	CreateDate                                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                                 *String    `xmlrpc:"display_name,omitempty"`
	Id                                          *Int       `xmlrpc:"id,omitempty"`
	Name                                        *String    `xmlrpc:"name,omitempty"`
	ParentId                                    *Many2One  `xmlrpc:"parent_id,omitempty"`
	ParentLeft                                  *Int       `xmlrpc:"parent_left,omitempty"`
	ParentRight                                 *Int       `xmlrpc:"parent_right,omitempty"`
	ProductCount                                *Int       `xmlrpc:"product_count,omitempty"`
	PropertyAccountCreditorPriceDifferenceCateg *Many2One  `xmlrpc:"property_account_creditor_price_difference_categ,omitempty"`
	PropertyAccountExpenseCategId               *Many2One  `xmlrpc:"property_account_expense_categ_id,omitempty"`
	PropertyAccountIncomeCategId                *Many2One  `xmlrpc:"property_account_income_categ_id,omitempty"`
	PropertyCostMethod                          *Selection `xmlrpc:"property_cost_method,omitempty"`
	PropertyStockAccountInputCategId            *Many2One  `xmlrpc:"property_stock_account_input_categ_id,omitempty"`
	PropertyStockAccountOutputCategId           *Many2One  `xmlrpc:"property_stock_account_output_categ_id,omitempty"`
	PropertyStockJournal                        *Many2One  `xmlrpc:"property_stock_journal,omitempty"`
	PropertyStockValuationAccountId             *Many2One  `xmlrpc:"property_stock_valuation_account_id,omitempty"`
	PropertyValuation                           *Selection `xmlrpc:"property_valuation,omitempty"`
	RemovalStrategyId                           *Many2One  `xmlrpc:"removal_strategy_id,omitempty"`
	RouteIds                                    *Relation  `xmlrpc:"route_ids,omitempty"`
	TotalRouteIds                               *Relation  `xmlrpc:"total_route_ids,omitempty"`
	WriteDate                                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                                    *Many2One  `xmlrpc:"write_uid,omitempty"`
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

// CreateProductCategorys creates a new product.category model and returns its id.
func (c *Client) CreateProductCategorys(pcs []*ProductCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcs {
		vv = append(vv, v)
	}
	return c.Create(ProductCategoryModel, vv, nil)
}

// UpdateProductCategory updates an existing product.category record.
func (c *Client) UpdateProductCategory(pc *ProductCategory) error {
	return c.UpdateProductCategorys([]int64{pc.Id.Get()}, pc)
}

// UpdateProductCategorys updates existing product.category records.
// All records (represented by ids) will be updated by pc values.
func (c *Client) UpdateProductCategorys(ids []int64, pc *ProductCategory) error {
	return c.Update(ProductCategoryModel, ids, pc, nil)
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
	return &((*pcs)[0]), nil
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
	return &((*pcs)[0]), nil
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
	return c.Search(ProductCategoryModel, criteria, options)
}

// FindProductCategoryId finds record id by querying it with criteria.
func (c *Client) FindProductCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
