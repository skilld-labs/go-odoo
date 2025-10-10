package odoo

// ProductLabelLayout represents product.label.layout model.
type ProductLabelLayout struct {
	Columns        *Int       `xmlrpc:"columns,omitempty"`
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	CustomQuantity *Int       `xmlrpc:"custom_quantity,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	ExtraHtml      *String    `xmlrpc:"extra_html,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	MoveIds        *Relation  `xmlrpc:"move_ids,omitempty"`
	MoveQuantity   *Selection `xmlrpc:"move_quantity,omitempty"`
	PricelistId    *Many2One  `xmlrpc:"pricelist_id,omitempty"`
	PrintFormat    *Selection `xmlrpc:"print_format,omitempty"`
	ProductIds     *Relation  `xmlrpc:"product_ids,omitempty"`
	ProductTmplIds *Relation  `xmlrpc:"product_tmpl_ids,omitempty"`
	Rows           *Int       `xmlrpc:"rows,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProductLabelLayouts represents array of product.label.layout model.
type ProductLabelLayouts []ProductLabelLayout

// ProductLabelLayoutModel is the odoo model name.
const ProductLabelLayoutModel = "product.label.layout"

// Many2One convert ProductLabelLayout to *Many2One.
func (pll *ProductLabelLayout) Many2One() *Many2One {
	return NewMany2One(pll.Id.Get(), "")
}

// CreateProductLabelLayout creates a new product.label.layout model and returns its id.
func (c *Client) CreateProductLabelLayout(pll *ProductLabelLayout) (int64, error) {
	ids, err := c.CreateProductLabelLayouts([]*ProductLabelLayout{pll})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductLabelLayout creates a new product.label.layout model and returns its id.
func (c *Client) CreateProductLabelLayouts(plls []*ProductLabelLayout) ([]int64, error) {
	var vv []interface{}
	for _, v := range plls {
		vv = append(vv, v)
	}
	return c.Create(ProductLabelLayoutModel, vv, nil)
}

// UpdateProductLabelLayout updates an existing product.label.layout record.
func (c *Client) UpdateProductLabelLayout(pll *ProductLabelLayout) error {
	return c.UpdateProductLabelLayouts([]int64{pll.Id.Get()}, pll)
}

// UpdateProductLabelLayouts updates existing product.label.layout records.
// All records (represented by ids) will be updated by pll values.
func (c *Client) UpdateProductLabelLayouts(ids []int64, pll *ProductLabelLayout) error {
	return c.Update(ProductLabelLayoutModel, ids, pll, nil)
}

// DeleteProductLabelLayout deletes an existing product.label.layout record.
func (c *Client) DeleteProductLabelLayout(id int64) error {
	return c.DeleteProductLabelLayouts([]int64{id})
}

// DeleteProductLabelLayouts deletes existing product.label.layout records.
func (c *Client) DeleteProductLabelLayouts(ids []int64) error {
	return c.Delete(ProductLabelLayoutModel, ids)
}

// GetProductLabelLayout gets product.label.layout existing record.
func (c *Client) GetProductLabelLayout(id int64) (*ProductLabelLayout, error) {
	plls, err := c.GetProductLabelLayouts([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*plls)[0]), nil
}

// GetProductLabelLayouts gets product.label.layout existing records.
func (c *Client) GetProductLabelLayouts(ids []int64) (*ProductLabelLayouts, error) {
	plls := &ProductLabelLayouts{}
	if err := c.Read(ProductLabelLayoutModel, ids, nil, plls); err != nil {
		return nil, err
	}
	return plls, nil
}

// FindProductLabelLayout finds product.label.layout record by querying it with criteria.
func (c *Client) FindProductLabelLayout(criteria *Criteria) (*ProductLabelLayout, error) {
	plls := &ProductLabelLayouts{}
	if err := c.SearchRead(ProductLabelLayoutModel, criteria, NewOptions().Limit(1), plls); err != nil {
		return nil, err
	}
	return &((*plls)[0]), nil
}

// FindProductLabelLayouts finds product.label.layout records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductLabelLayouts(criteria *Criteria, options *Options) (*ProductLabelLayouts, error) {
	plls := &ProductLabelLayouts{}
	if err := c.SearchRead(ProductLabelLayoutModel, criteria, options, plls); err != nil {
		return nil, err
	}
	return plls, nil
}

// FindProductLabelLayoutIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductLabelLayoutIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductLabelLayoutModel, criteria, options)
}

// FindProductLabelLayoutId finds record id by querying it with criteria.
func (c *Client) FindProductLabelLayoutId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductLabelLayoutModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
