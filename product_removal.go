package odoo

// ProductRemoval represents product.removal model.
type ProductRemoval struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Method      *String   `xmlrpc:"method,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProductRemovals represents array of product.removal model.
type ProductRemovals []ProductRemoval

// ProductRemovalModel is the odoo model name.
const ProductRemovalModel = "product.removal"

// Many2One convert ProductRemoval to *Many2One.
func (pr *ProductRemoval) Many2One() *Many2One {
	return NewMany2One(pr.Id.Get(), "")
}

// CreateProductRemoval creates a new product.removal model and returns its id.
func (c *Client) CreateProductRemoval(pr *ProductRemoval) (int64, error) {
	ids, err := c.CreateProductRemovals([]*ProductRemoval{pr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProductRemoval creates a new product.removal model and returns its id.
func (c *Client) CreateProductRemovals(prs []*ProductRemoval) ([]int64, error) {
	var vv []interface{}
	for _, v := range prs {
		vv = append(vv, v)
	}
	return c.Create(ProductRemovalModel, vv, nil)
}

// UpdateProductRemoval updates an existing product.removal record.
func (c *Client) UpdateProductRemoval(pr *ProductRemoval) error {
	return c.UpdateProductRemovals([]int64{pr.Id.Get()}, pr)
}

// UpdateProductRemovals updates existing product.removal records.
// All records (represented by ids) will be updated by pr values.
func (c *Client) UpdateProductRemovals(ids []int64, pr *ProductRemoval) error {
	return c.Update(ProductRemovalModel, ids, pr, nil)
}

// DeleteProductRemoval deletes an existing product.removal record.
func (c *Client) DeleteProductRemoval(id int64) error {
	return c.DeleteProductRemovals([]int64{id})
}

// DeleteProductRemovals deletes existing product.removal records.
func (c *Client) DeleteProductRemovals(ids []int64) error {
	return c.Delete(ProductRemovalModel, ids)
}

// GetProductRemoval gets product.removal existing record.
func (c *Client) GetProductRemoval(id int64) (*ProductRemoval, error) {
	prs, err := c.GetProductRemovals([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// GetProductRemovals gets product.removal existing records.
func (c *Client) GetProductRemovals(ids []int64) (*ProductRemovals, error) {
	prs := &ProductRemovals{}
	if err := c.Read(ProductRemovalModel, ids, nil, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProductRemoval finds product.removal record by querying it with criteria.
func (c *Client) FindProductRemoval(criteria *Criteria) (*ProductRemoval, error) {
	prs := &ProductRemovals{}
	if err := c.SearchRead(ProductRemovalModel, criteria, NewOptions().Limit(1), prs); err != nil {
		return nil, err
	}
	return &((*prs)[0]), nil
}

// FindProductRemovals finds product.removal records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductRemovals(criteria *Criteria, options *Options) (*ProductRemovals, error) {
	prs := &ProductRemovals{}
	if err := c.SearchRead(ProductRemovalModel, criteria, options, prs); err != nil {
		return nil, err
	}
	return prs, nil
}

// FindProductRemovalIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProductRemovalIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProductRemovalModel, criteria, options)
}

// FindProductRemovalId finds record id by querying it with criteria.
func (c *Client) FindProductRemovalId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProductRemovalModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
