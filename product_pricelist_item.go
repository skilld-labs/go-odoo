package odoo

// ProductPricelistItem represents product.pricelist.item model.
type ProductPricelistItem struct {
	LastUpdate      *Time      `xmlrpc:"__last_update,omptempty"`
	AppliedOn       *Selection `xmlrpc:"applied_on,omptempty"`
	Base            *Selection `xmlrpc:"base,omptempty"`
	BasePricelistId *Many2One  `xmlrpc:"base_pricelist_id,omptempty"`
	CategId         *Many2One  `xmlrpc:"categ_id,omptempty"`
	CompanyId       *Many2One  `xmlrpc:"company_id,omptempty"`
	ComputePrice    *Selection `xmlrpc:"compute_price,omptempty"`
	CreateDate      *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid       *Many2One  `xmlrpc:"create_uid,omptempty"`
	CurrencyId      *Many2One  `xmlrpc:"currency_id,omptempty"`
	DateEnd         *Time      `xmlrpc:"date_end,omptempty"`
	DateStart       *Time      `xmlrpc:"date_start,omptempty"`
	DisplayName     *String    `xmlrpc:"display_name,omptempty"`
	FixedPrice      *Float     `xmlrpc:"fixed_price,omptempty"`
	Id              *Int       `xmlrpc:"id,omptempty"`
	MinQuantity     *Int       `xmlrpc:"min_quantity,omptempty"`
	Name            *String    `xmlrpc:"name,omptempty"`
	PercentPrice    *Float     `xmlrpc:"percent_price,omptempty"`
	Price           *String    `xmlrpc:"price,omptempty"`
	PriceDiscount   *Float     `xmlrpc:"price_discount,omptempty"`
	PriceMaxMargin  *Float     `xmlrpc:"price_max_margin,omptempty"`
	PriceMinMargin  *Float     `xmlrpc:"price_min_margin,omptempty"`
	PriceRound      *Float     `xmlrpc:"price_round,omptempty"`
	PriceSurcharge  *Float     `xmlrpc:"price_surcharge,omptempty"`
	PricelistId     *Many2One  `xmlrpc:"pricelist_id,omptempty"`
	ProductId       *Many2One  `xmlrpc:"product_id,omptempty"`
	ProductTmplId   *Many2One  `xmlrpc:"product_tmpl_id,omptempty"`
	WriteDate       *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid        *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// ProductPricelistItems represents array of product.pricelist.item model.
type ProductPricelistItems []ProductPricelistItem

// ProductPricelistItemModel is the odoo model name.
const ProductPricelistItemModel = "product.pricelist.item"

// Many2One convert ProductPricelistItem to *Many2One.
func (ppi *ProductPricelistItem) Many2One() *Many2One {
	return NewMany2One(ppi.Id.Get(), "")
}

// CreateProductPricelistItem creates a new product.pricelist.item model and returns its id.
func (c *Client) CreateProductPricelistItem(ppi *ProductPricelistItem) (int64, error) {
	ids, err := c.CreateProductPricelistItems([]*ProductPricelistItem{ppi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductPricelistItem creates a new product.pricelist.item model and returns its id.
func (c *Client) CreateProductPricelistItems(ppis []*ProductPricelistItem) ([]int64, error) {
	var vv []interface{}
	for _, v := range ppis {
		vv = append(vv, v)
	}
	return c.Create(ProductPricelistItemModel, vv, nil)
}

// UpdateProductPricelistItem updates an existing product.pricelist.item record.
func (c *Client) UpdateProductPricelistItem(ppi *ProductPricelistItem) error {
	return c.UpdateProductPricelistItems([]int64{ppi.Id.Get()}, ppi)
}

// UpdateProductPricelistItems updates existing product.pricelist.item records.
// All records (represented by ids) will be updated by ppi values.
func (c *Client) UpdateProductPricelistItems(ids []int64, ppi *ProductPricelistItem) error {
	return c.Update(ProductPricelistItemModel, ids, ppi, nil)
}

// DeleteProductPricelistItem deletes an existing product.pricelist.item record.
func (c *Client) DeleteProductPricelistItem(id int64) error {
	return c.DeleteProductPricelistItems([]int64{id})
}

// DeleteProductPricelistItems deletes existing product.pricelist.item records.
func (c *Client) DeleteProductPricelistItems(ids []int64) error {
	return c.Delete(ProductPricelistItemModel, ids)
}

// GetProductPricelistItem gets product.pricelist.item existing record.
func (c *Client) GetProductPricelistItem(id int64) (*ProductPricelistItem, error) {
	ppis, err := c.GetProductPricelistItems([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ppis)[0]), nil
}

// GetProductPricelistItems gets product.pricelist.item existing records.
func (c *Client) GetProductPricelistItems(ids []int64) (*ProductPricelistItems, error) {
	ppis := &ProductPricelistItems{}
	if err := c.Read(ProductPricelistItemModel, ids, nil, ppis); err != nil {
		return nil, err
	}
	return ppis, nil
}

// FindProductPricelistItem finds product.pricelist.item record by querying it with criteria.
func (c *Client) FindProductPricelistItem(criteria *Criteria) (*ProductPricelistItem, error) {
	ppis := &ProductPricelistItems{}
	if err := c.SearchRead(ProductPricelistItemModel, criteria, NewOptions().Limit(1), ppis); err != nil {
		return nil, err
	}
	return &((*ppis)[0]), nil
}

// FindProductPricelistItems finds product.pricelist.item records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelistItems(criteria *Criteria, options *Options) (*ProductPricelistItems, error) {
	ppis := &ProductPricelistItems{}
	if err := c.SearchRead(ProductPricelistItemModel, criteria, options, ppis); err != nil {
		return nil, err
	}
	return ppis, nil
}

// FindProductPricelistItemIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductPricelistItemIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductPricelistItemModel, criteria, options)
}

// FindProductPricelistItemId finds record id by querying it with criteria.
func (c *Client) FindProductPricelistItemId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductPricelistItemModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
