package odoo

// BaseModuleInstallReview represents base.module.install.review model.
type BaseModuleInstallReview struct {
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	ModuleId           *Many2One `xmlrpc:"module_id,omitempty"`
	ModuleIds          *Relation `xmlrpc:"module_ids,omitempty"`
	ModulesDescription *String   `xmlrpc:"modules_description,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// BaseModuleInstallReviews represents array of base.module.install.review model.
type BaseModuleInstallReviews []BaseModuleInstallReview

// BaseModuleInstallReviewModel is the odoo model name.
const BaseModuleInstallReviewModel = "base.module.install.review"

// Many2One convert BaseModuleInstallReview to *Many2One.
func (bmir *BaseModuleInstallReview) Many2One() *Many2One {
	return NewMany2One(bmir.Id.Get(), "")
}

// CreateBaseModuleInstallReview creates a new base.module.install.review model and returns its id.
func (c *Client) CreateBaseModuleInstallReview(bmir *BaseModuleInstallReview) (int64, error) {
	ids, err := c.CreateBaseModuleInstallReviews([]*BaseModuleInstallReview{bmir})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseModuleInstallReview creates a new base.module.install.review model and returns its id.
func (c *Client) CreateBaseModuleInstallReviews(bmirs []*BaseModuleInstallReview) ([]int64, error) {
	var vv []interface{}
	for _, v := range bmirs {
		vv = append(vv, v)
	}
	return c.Create(BaseModuleInstallReviewModel, vv, nil)
}

// UpdateBaseModuleInstallReview updates an existing base.module.install.review record.
func (c *Client) UpdateBaseModuleInstallReview(bmir *BaseModuleInstallReview) error {
	return c.UpdateBaseModuleInstallReviews([]int64{bmir.Id.Get()}, bmir)
}

// UpdateBaseModuleInstallReviews updates existing base.module.install.review records.
// All records (represented by ids) will be updated by bmir values.
func (c *Client) UpdateBaseModuleInstallReviews(ids []int64, bmir *BaseModuleInstallReview) error {
	return c.Update(BaseModuleInstallReviewModel, ids, bmir, nil)
}

// DeleteBaseModuleInstallReview deletes an existing base.module.install.review record.
func (c *Client) DeleteBaseModuleInstallReview(id int64) error {
	return c.DeleteBaseModuleInstallReviews([]int64{id})
}

// DeleteBaseModuleInstallReviews deletes existing base.module.install.review records.
func (c *Client) DeleteBaseModuleInstallReviews(ids []int64) error {
	return c.Delete(BaseModuleInstallReviewModel, ids)
}

// GetBaseModuleInstallReview gets base.module.install.review existing record.
func (c *Client) GetBaseModuleInstallReview(id int64) (*BaseModuleInstallReview, error) {
	bmirs, err := c.GetBaseModuleInstallReviews([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bmirs)[0]), nil
}

// GetBaseModuleInstallReviews gets base.module.install.review existing records.
func (c *Client) GetBaseModuleInstallReviews(ids []int64) (*BaseModuleInstallReviews, error) {
	bmirs := &BaseModuleInstallReviews{}
	if err := c.Read(BaseModuleInstallReviewModel, ids, nil, bmirs); err != nil {
		return nil, err
	}
	return bmirs, nil
}

// FindBaseModuleInstallReview finds base.module.install.review record by querying it with criteria.
func (c *Client) FindBaseModuleInstallReview(criteria *Criteria) (*BaseModuleInstallReview, error) {
	bmirs := &BaseModuleInstallReviews{}
	if err := c.SearchRead(BaseModuleInstallReviewModel, criteria, NewOptions().Limit(1), bmirs); err != nil {
		return nil, err
	}
	return &((*bmirs)[0]), nil
}

// FindBaseModuleInstallReviews finds base.module.install.review records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleInstallReviews(criteria *Criteria, options *Options) (*BaseModuleInstallReviews, error) {
	bmirs := &BaseModuleInstallReviews{}
	if err := c.SearchRead(BaseModuleInstallReviewModel, criteria, options, bmirs); err != nil {
		return nil, err
	}
	return bmirs, nil
}

// FindBaseModuleInstallReviewIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseModuleInstallReviewIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseModuleInstallReviewModel, criteria, options)
}

// FindBaseModuleInstallReviewId finds record id by querying it with criteria.
func (c *Client) FindBaseModuleInstallReviewId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseModuleInstallReviewModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
