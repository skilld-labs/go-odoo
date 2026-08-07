package odoo

// ProductCatalogMixin represents product.catalog.mixin model.
type ProductCatalogMixin struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// ProductCatalogMixins represents array of product.catalog.mixin model.
type ProductCatalogMixins []ProductCatalogMixin

// ProductCatalogMixinModel is the odoo model name.
const ProductCatalogMixinModel = "product.catalog.mixin"

// Many2One convert ProductCatalogMixin to *Many2One.
func (pcm *ProductCatalogMixin) Many2One() *Many2One {
	return NewMany2One(pcm.Id.Get(), "")
}

// CreateProductCatalogMixin creates a new product.catalog.mixin model and returns its id.
func (c *Client) CreateProductCatalogMixin(pcm *ProductCatalogMixin) (int64, error) {
	ids, err := c.CreateProductCatalogMixins([]*ProductCatalogMixin{pcm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductCatalogMixin creates a new product.catalog.mixin model and returns its id.
func (c *Client) CreateProductCatalogMixins(pcms []*ProductCatalogMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range pcms {
		vv = append(vv, v)
	}
	return c.Create(ProductCatalogMixinModel, vv, nil)
}

// UpdateProductCatalogMixin updates an existing product.catalog.mixin record.
func (c *Client) UpdateProductCatalogMixin(pcm *ProductCatalogMixin) error {
	return c.UpdateProductCatalogMixins([]int64{pcm.Id.Get()}, pcm)
}

// UpdateProductCatalogMixins updates existing product.catalog.mixin records.
// All records (represented by ids) will be updated by pcm values.
func (c *Client) UpdateProductCatalogMixins(ids []int64, pcm *ProductCatalogMixin) error {
	return c.Update(ProductCatalogMixinModel, ids, pcm, nil)
}

// DeleteProductCatalogMixin deletes an existing product.catalog.mixin record.
func (c *Client) DeleteProductCatalogMixin(id int64) error {
	return c.DeleteProductCatalogMixins([]int64{id})
}

// DeleteProductCatalogMixins deletes existing product.catalog.mixin records.
func (c *Client) DeleteProductCatalogMixins(ids []int64) error {
	return c.Delete(ProductCatalogMixinModel, ids)
}

// GetProductCatalogMixin gets product.catalog.mixin existing record.
func (c *Client) GetProductCatalogMixin(id int64) (*ProductCatalogMixin, error) {
	pcms, err := c.GetProductCatalogMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pcms)[0]), nil
}

// GetProductCatalogMixins gets product.catalog.mixin existing records.
func (c *Client) GetProductCatalogMixins(ids []int64) (*ProductCatalogMixins, error) {
	pcms := &ProductCatalogMixins{}
	if err := c.Read(ProductCatalogMixinModel, ids, nil, pcms); err != nil {
		return nil, err
	}
	return pcms, nil
}

// FindProductCatalogMixin finds product.catalog.mixin record by querying it with criteria.
func (c *Client) FindProductCatalogMixin(criteria *Criteria) (*ProductCatalogMixin, error) {
	pcms := &ProductCatalogMixins{}
	if err := c.SearchRead(ProductCatalogMixinModel, criteria, NewOptions().Limit(1), pcms); err != nil {
		return nil, err
	}
	return &((*pcms)[0]), nil
}

// FindProductCatalogMixins finds product.catalog.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductCatalogMixins(criteria *Criteria, options *Options) (*ProductCatalogMixins, error) {
	pcms := &ProductCatalogMixins{}
	if err := c.SearchRead(ProductCatalogMixinModel, criteria, options, pcms); err != nil {
		return nil, err
	}
	return pcms, nil
}

// FindProductCatalogMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductCatalogMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductCatalogMixinModel, criteria, options)
}

// FindProductCatalogMixinId finds record id by querying it with criteria.
func (c *Client) FindProductCatalogMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductCatalogMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
