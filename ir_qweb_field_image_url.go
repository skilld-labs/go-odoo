package odoo

// IrQwebFieldImageUrl represents ir.qweb.field.image_url model.
type IrQwebFieldImageUrl struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IrQwebFieldImageUrls represents array of ir.qweb.field.image_url model.
type IrQwebFieldImageUrls []IrQwebFieldImageUrl

// IrQwebFieldImageUrlModel is the odoo model name.
const IrQwebFieldImageUrlModel = "ir.qweb.field.image_url"

// Many2One convert IrQwebFieldImageUrl to *Many2One.
func (iqfi *IrQwebFieldImageUrl) Many2One() *Many2One {
	return NewMany2One(iqfi.Id.Get(), "")
}

// CreateIrQwebFieldImageUrl creates a new ir.qweb.field.image_url model and returns its id.
func (c *Client) CreateIrQwebFieldImageUrl(iqfi *IrQwebFieldImageUrl) (int64, error) {
	ids, err := c.CreateIrQwebFieldImageUrls([]*IrQwebFieldImageUrl{iqfi})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIrQwebFieldImageUrl creates a new ir.qweb.field.image_url model and returns its id.
func (c *Client) CreateIrQwebFieldImageUrls(iqfis []*IrQwebFieldImageUrl) ([]int64, error) {
	var vv []interface{}
	for _, v := range iqfis {
		vv = append(vv, v)
	}
	return c.Create(IrQwebFieldImageUrlModel, vv, nil)
}

// UpdateIrQwebFieldImageUrl updates an existing ir.qweb.field.image_url record.
func (c *Client) UpdateIrQwebFieldImageUrl(iqfi *IrQwebFieldImageUrl) error {
	return c.UpdateIrQwebFieldImageUrls([]int64{iqfi.Id.Get()}, iqfi)
}

// UpdateIrQwebFieldImageUrls updates existing ir.qweb.field.image_url records.
// All records (represented by ids) will be updated by iqfi values.
func (c *Client) UpdateIrQwebFieldImageUrls(ids []int64, iqfi *IrQwebFieldImageUrl) error {
	return c.Update(IrQwebFieldImageUrlModel, ids, iqfi, nil)
}

// DeleteIrQwebFieldImageUrl deletes an existing ir.qweb.field.image_url record.
func (c *Client) DeleteIrQwebFieldImageUrl(id int64) error {
	return c.DeleteIrQwebFieldImageUrls([]int64{id})
}

// DeleteIrQwebFieldImageUrls deletes existing ir.qweb.field.image_url records.
func (c *Client) DeleteIrQwebFieldImageUrls(ids []int64) error {
	return c.Delete(IrQwebFieldImageUrlModel, ids)
}

// GetIrQwebFieldImageUrl gets ir.qweb.field.image_url existing record.
func (c *Client) GetIrQwebFieldImageUrl(id int64) (*IrQwebFieldImageUrl, error) {
	iqfis, err := c.GetIrQwebFieldImageUrls([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iqfis)[0]), nil
}

// GetIrQwebFieldImageUrls gets ir.qweb.field.image_url existing records.
func (c *Client) GetIrQwebFieldImageUrls(ids []int64) (*IrQwebFieldImageUrls, error) {
	iqfis := &IrQwebFieldImageUrls{}
	if err := c.Read(IrQwebFieldImageUrlModel, ids, nil, iqfis); err != nil {
		return nil, err
	}
	return iqfis, nil
}

// FindIrQwebFieldImageUrl finds ir.qweb.field.image_url record by querying it with criteria.
func (c *Client) FindIrQwebFieldImageUrl(criteria *Criteria) (*IrQwebFieldImageUrl, error) {
	iqfis := &IrQwebFieldImageUrls{}
	if err := c.SearchRead(IrQwebFieldImageUrlModel, criteria, NewOptions().Limit(1), iqfis); err != nil {
		return nil, err
	}
	return &((*iqfis)[0]), nil
}

// FindIrQwebFieldImageUrls finds ir.qweb.field.image_url records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldImageUrls(criteria *Criteria, options *Options) (*IrQwebFieldImageUrls, error) {
	iqfis := &IrQwebFieldImageUrls{}
	if err := c.SearchRead(IrQwebFieldImageUrlModel, criteria, options, iqfis); err != nil {
		return nil, err
	}
	return iqfis, nil
}

// FindIrQwebFieldImageUrlIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIrQwebFieldImageUrlIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IrQwebFieldImageUrlModel, criteria, options)
}

// FindIrQwebFieldImageUrlId finds record id by querying it with criteria.
func (c *Client) FindIrQwebFieldImageUrlId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IrQwebFieldImageUrlModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
