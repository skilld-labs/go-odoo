package odoo

// ProductReplenish represents product.replenish model.
type ProductReplenish struct {
	AllowedRouteIds      *Relation `xmlrpc:"allowed_route_ids,omitempty"`
	CompanyId            *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate           *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid            *Many2One `xmlrpc:"create_uid,omitempty"`
	DatePlanned          *Time     `xmlrpc:"date_planned,omitempty"`
	DisplayName          *String   `xmlrpc:"display_name,omitempty"`
	ForecastUomId        *Many2One `xmlrpc:"forecast_uom_id,omitempty"`
	ForecastedQuantity   *Float    `xmlrpc:"forecasted_quantity,omitempty"`
	Id                   *Int      `xmlrpc:"id,omitempty"`
	ProductHasVariants   *Bool     `xmlrpc:"product_has_variants,omitempty"`
	ProductId            *Many2One `xmlrpc:"product_id,omitempty"`
	ProductTmplId        *Many2One `xmlrpc:"product_tmpl_id,omitempty"`
	ProductUomCategoryId *Many2One `xmlrpc:"product_uom_category_id,omitempty"`
	ProductUomId         *Many2One `xmlrpc:"product_uom_id,omitempty"`
	Quantity             *Float    `xmlrpc:"quantity,omitempty"`
	RouteId              *Many2One `xmlrpc:"route_id,omitempty"`
	ShowVendor           *Bool     `xmlrpc:"show_vendor,omitempty"`
	SupplierId           *Many2One `xmlrpc:"supplier_id,omitempty"`
	WarehouseId          *Many2One `xmlrpc:"warehouse_id,omitempty"`
	WriteDate            *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid             *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductReplenishs represents array of product.replenish model.
type ProductReplenishs []ProductReplenish

// ProductReplenishModel is the odoo model name.
const ProductReplenishModel = "product.replenish"

// Many2One convert ProductReplenish to *Many2One.
func (pr *ProductReplenish) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreateProductReplenish creates a new product.replenish model and returns its id.
func (c *Client) CreateProductReplenish(pr *ProductReplenish) (int64, error) {
	ids, err := c.CreateProductReplenishs([]*ProductReplenish{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductReplenish creates a new product.replenish model and returns its id.
func (c *Client) CreateProductReplenishs(prs []*ProductReplenish) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(ProductReplenishModel, vv, nil)
}

// UpdateProductReplenish updates an existing product.replenish record.
func (c *Client) UpdateProductReplenish(pr *ProductReplenish) error {
	return c.UpdateProductReplenishs([]int64{pr.Id.Get()}, pr)
}

// UpdateProductReplenishs updates existing product.replenish records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdateProductReplenishs(ids []int64, pr *ProductReplenish) error {
	return c.Update(ProductReplenishModel, ids, pr, nil)
}

// DeleteProductReplenish deletes an existing product.replenish record.
func (c *Client) DeleteProductReplenish(id int64) error {
	return c.DeleteProductReplenishs([]int64{id})
}

// DeleteProductReplenishs deletes existing product.replenish records.
func (c *Client) DeleteProductReplenishs(ids []int64) error {
	return c.Delete(ProductReplenishModel, ids)
}

// GetProductReplenish gets product.replenish existing record.
func (c *Client) GetProductReplenish(id int64) (*ProductReplenish, error) {
	prs, err := c.GetProductReplenishs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetProductReplenishs gets product.replenish existing records.
func (c *Client) GetProductReplenishs(ids []int64) (*ProductReplenishs, error) {
	prs := &ProductReplenishs{}
	if err := c.Read(ProductReplenishModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProductReplenish finds product.replenish record by querying it with criteria.
func (c *Client) FindProductReplenish(criteria *Criteria) (*ProductReplenish, error) {
	prs := &ProductReplenishs{}
	if err := c.SearchRead(ProductReplenishModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindProductReplenishs finds product.replenish records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductReplenishs(criteria *Criteria, options *Options) (*ProductReplenishs, error) {
	prs := &ProductReplenishs{}
	if err := c.SearchRead(ProductReplenishModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProductReplenishIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductReplenishIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductReplenishModel, criteria, options)
}

// FindProductReplenishId finds record id by querying it with criteria.
func (c *Client) FindProductReplenishId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductReplenishModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
