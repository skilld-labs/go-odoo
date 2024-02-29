package odoo

// ResPartnerCategory represents res.partner.category model.
type ResPartnerCategory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Active      *Bool     `xmlrpc:"active,omitempty"`
	ChildIds    *Relation `xmlrpc:"child_ids,omitempty"`
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	ParentId    *Many2One `xmlrpc:"parent_id,omitempty"`
	ParentLeft  *Int      `xmlrpc:"parent_left,omitempty"`
	ParentRight *Int      `xmlrpc:"parent_right,omitempty"`
	PartnerIds  *Relation `xmlrpc:"partner_ids,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResPartnerCategorys represents array of res.partner.category model.
type ResPartnerCategorys []ResPartnerCategory

// ResPartnerCategoryModel is the odoo model name.
const ResPartnerCategoryModel = "res.partner.category"

// Many2One convert ResPartnerCategory to *Many2One.
func (rpc *ResPartnerCategory) Many2One() *Many2One {
	return NewMany2One(rpc.Id.Get(), "")
}

// CreateResPartnerCategory creates a new res.partner.category model and returns its id.
func (c *Client) CreateResPartnerCategory(rpc *ResPartnerCategory) (int64, error) {
	ids, err := c.CreateResPartnerCategorys([]*ResPartnerCategory{rpc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResPartnerCategorys creates a new res.partner.category model and returns its id.
func (c *Client) CreateResPartnerCategorys(rpcs []*ResPartnerCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range rpcs {
		vv = append(vv, v)
	}
	return c.Create(ResPartnerCategoryModel, vv, nil)
}

// UpdateResPartnerCategory updates an existing res.partner.category record.
func (c *Client) UpdateResPartnerCategory(rpc *ResPartnerCategory) error {
	return c.UpdateResPartnerCategorys([]int64{rpc.Id.Get()}, rpc)
}

// UpdateResPartnerCategorys updates existing res.partner.category records.
// All records (represented by ids) will be updated by rpc values.
func (c *Client) UpdateResPartnerCategorys(ids []int64, rpc *ResPartnerCategory) error {
	return c.Update(ResPartnerCategoryModel, ids, rpc, nil)
}

// DeleteResPartnerCategory deletes an existing res.partner.category record.
func (c *Client) DeleteResPartnerCategory(id int64) error {
	return c.DeleteResPartnerCategorys([]int64{id})
}

// DeleteResPartnerCategorys deletes existing res.partner.category records.
func (c *Client) DeleteResPartnerCategorys(ids []int64) error {
	return c.Delete(ResPartnerCategoryModel, ids)
}

// GetResPartnerCategory gets res.partner.category existing record.
func (c *Client) GetResPartnerCategory(id int64) (*ResPartnerCategory, error) {
	rpcs, err := c.GetResPartnerCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*rpcs)[0]), nil
}

// GetResPartnerCategorys gets res.partner.category existing records.
func (c *Client) GetResPartnerCategorys(ids []int64) (*ResPartnerCategorys, error) {
	rpcs := &ResPartnerCategorys{}
	if err := c.Read(ResPartnerCategoryModel, ids, nil, rpcs); err != nil {
		return nil, err
	}
	return rpcs, nil
}

// FindResPartnerCategory finds res.partner.category record by querying it with criteria.
func (c *Client) FindResPartnerCategory(criteria *Criteria) (*ResPartnerCategory, error) {
	rpcs := &ResPartnerCategorys{}
	if err := c.SearchRead(ResPartnerCategoryModel, criteria, NewOptions().Limit(1), rpcs); err != nil {
		return nil, err
	}
	return &((*rpcs)[0]), nil
}

// FindResPartnerCategorys finds res.partner.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerCategorys(criteria *Criteria, options *Options) (*ResPartnerCategorys, error) {
	rpcs := &ResPartnerCategorys{}
	if err := c.SearchRead(ResPartnerCategoryModel, criteria, options, rpcs); err != nil {
		return nil, err
	}
	return rpcs, nil
}

// FindResPartnerCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResPartnerCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResPartnerCategoryModel, criteria, options)
}

// FindResPartnerCategoryId finds record id by querying it with criteria.
func (c *Client) FindResPartnerCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResPartnerCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
