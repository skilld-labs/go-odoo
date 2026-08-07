package odoo

// ProductValue represents product.value model.
type ProductValue struct {
	CompanyId                *Many2One `xmlrpc:"company_id,omitempty"`
	ComputedValueDescription *String   `xmlrpc:"computed_value_description,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId               *Many2One `xmlrpc:"currency_id,omitempty"`
	CurrentValue             *Float    `xmlrpc:"current_value,omitempty"`
	CurrentValueDescription  *String   `xmlrpc:"current_value_description,omitempty"`
	CurrentValueDetails      *String   `xmlrpc:"current_value_details,omitempty"`
	Date                     *Time     `xmlrpc:"date,omitempty"`
	Description              *String   `xmlrpc:"description,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	LotId                    *Many2One `xmlrpc:"lot_id,omitempty"`
	MoveId                   *Many2One `xmlrpc:"move_id,omitempty"`
	ProductId                *Many2One `xmlrpc:"product_id,omitempty"`
	UserId                   *Many2One `xmlrpc:"user_id,omitempty"`
	Value                    *Float    `xmlrpc:"value,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductValues represents array of product.value model.
type ProductValues []ProductValue

// ProductValueModel is the odoo model name.
const ProductValueModel = "product.value"

// Many2One convert ProductValue to *Many2One.
func (pv *ProductValue) Many2One() *Many2One {
	return NewMany2One(pv.Id.Get(), "")
}

// CreateProductValue creates a new product.value model and returns its id.
func (c *Client) CreateProductValue(pv *ProductValue) (int64, error) {
	ids, err := c.CreateProductValues([]*ProductValue{pv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductValue creates a new product.value model and returns its id.
func (c *Client) CreateProductValues(pvs []*ProductValue) ([]int64, error) {
	var vv []interface{}
	for _, v := range pvs {
		vv = append(vv, v)
	}
	return c.Create(ProductValueModel, vv, nil)
}

// UpdateProductValue updates an existing product.value record.
func (c *Client) UpdateProductValue(pv *ProductValue) error {
	return c.UpdateProductValues([]int64{pv.Id.Get()}, pv)
}

// UpdateProductValues updates existing product.value records.
// All records (represented by ids) will be updated by pv values.
func (c *Client) UpdateProductValues(ids []int64, pv *ProductValue) error {
	return c.Update(ProductValueModel, ids, pv, nil)
}

// DeleteProductValue deletes an existing product.value record.
func (c *Client) DeleteProductValue(id int64) error {
	return c.DeleteProductValues([]int64{id})
}

// DeleteProductValues deletes existing product.value records.
func (c *Client) DeleteProductValues(ids []int64) error {
	return c.Delete(ProductValueModel, ids)
}

// GetProductValue gets product.value existing record.
func (c *Client) GetProductValue(id int64) (*ProductValue, error) {
	pvs, err := c.GetProductValues([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pvs)[0]), nil
}

// GetProductValues gets product.value existing records.
func (c *Client) GetProductValues(ids []int64) (*ProductValues, error) {
	pvs := &ProductValues{}
	if err := c.Read(ProductValueModel, ids, nil, pvs); err != nil {
		return nil, err
	}
	return pvs, nil
}

// FindProductValue finds product.value record by querying it with criteria.
func (c *Client) FindProductValue(criteria *Criteria) (*ProductValue, error) {
	pvs := &ProductValues{}
	if err := c.SearchRead(ProductValueModel, criteria, NewOptions().Limit(1), pvs); err != nil {
		return nil, err
	}
	return &((*pvs)[0]), nil
}

// FindProductValues finds product.value records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductValues(criteria *Criteria, options *Options) (*ProductValues, error) {
	pvs := &ProductValues{}
	if err := c.SearchRead(ProductValueModel, criteria, options, pvs); err != nil {
		return nil, err
	}
	return pvs, nil
}

// FindProductValueIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductValueIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductValueModel, criteria, options)
}

// FindProductValueId finds record id by querying it with criteria.
func (c *Client) FindProductValueId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductValueModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
