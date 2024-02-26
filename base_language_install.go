package odoo

// BaseLanguageInstall represents base.language.install model.
type BaseLanguageInstall struct {
	LastUpdate  *Time      `xmlrpc:"__last_update,omptempty"`
	CreateDate  *Time      `xmlrpc:"create_date,omptempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omptempty"`
	DisplayName *String    `xmlrpc:"display_name,omptempty"`
	Id          *Int       `xmlrpc:"id,omptempty"`
	Lang        *Selection `xmlrpc:"lang,omptempty"`
	Overwrite   *Bool      `xmlrpc:"overwrite,omptempty"`
	State       *Selection `xmlrpc:"state,omptempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omptempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omptempty"`
}

// BaseLanguageInstalls represents array of base.language.install model.
type BaseLanguageInstalls []BaseLanguageInstall

// BaseLanguageInstallModel is the odoo model name.
const BaseLanguageInstallModel = "base.language.install"

// Many2One convert BaseLanguageInstall to *Many2One.
func (bli *BaseLanguageInstall) Many2One() *Many2One {
	return NewMany2One(bli.Id.Get(), "")
}

// CreateBaseLanguageInstall creates a new base.language.install model and returns its id.
func (c *Client) CreateBaseLanguageInstall(bli *BaseLanguageInstall) (int64, error) {
	ids, err := c.CreateBaseLanguageInstalls([]*BaseLanguageInstall{bli})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBaseLanguageInstall creates a new base.language.install model and returns its id.
func (c *Client) CreateBaseLanguageInstalls(blis []*BaseLanguageInstall) ([]int64, error) {
	var vv []interface{}
	for _, v := range blis {
		vv = append(vv, v)
	}
	return c.Create(BaseLanguageInstallModel, vv, nil)
}

// UpdateBaseLanguageInstall updates an existing base.language.install record.
func (c *Client) UpdateBaseLanguageInstall(bli *BaseLanguageInstall) error {
	return c.UpdateBaseLanguageInstalls([]int64{bli.Id.Get()}, bli)
}

// UpdateBaseLanguageInstalls updates existing base.language.install records.
// All records (represented by ids) will be updated by bli values.
func (c *Client) UpdateBaseLanguageInstalls(ids []int64, bli *BaseLanguageInstall) error {
	return c.Update(BaseLanguageInstallModel, ids, bli, nil)
}

// DeleteBaseLanguageInstall deletes an existing base.language.install record.
func (c *Client) DeleteBaseLanguageInstall(id int64) error {
	return c.DeleteBaseLanguageInstalls([]int64{id})
}

// DeleteBaseLanguageInstalls deletes existing base.language.install records.
func (c *Client) DeleteBaseLanguageInstalls(ids []int64) error {
	return c.Delete(BaseLanguageInstallModel, ids)
}

// GetBaseLanguageInstall gets base.language.install existing record.
func (c *Client) GetBaseLanguageInstall(id int64) (*BaseLanguageInstall, error) {
	blis, err := c.GetBaseLanguageInstalls([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*blis)[0]), nil
}

// GetBaseLanguageInstalls gets base.language.install existing records.
func (c *Client) GetBaseLanguageInstalls(ids []int64) (*BaseLanguageInstalls, error) {
	blis := &BaseLanguageInstalls{}
	if err := c.Read(BaseLanguageInstallModel, ids, nil, blis); err != nil {
		return nil, err
	}
	return blis, nil
}

// FindBaseLanguageInstall finds base.language.install record by querying it with criteria.
func (c *Client) FindBaseLanguageInstall(criteria *Criteria) (*BaseLanguageInstall, error) {
	blis := &BaseLanguageInstalls{}
	if err := c.SearchRead(BaseLanguageInstallModel, criteria, NewOptions().Limit(1), blis); err != nil {
		return nil, err
	}
	return &((*blis)[0]), nil
}

// FindBaseLanguageInstalls finds base.language.install records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseLanguageInstalls(criteria *Criteria, options *Options) (*BaseLanguageInstalls, error) {
	blis := &BaseLanguageInstalls{}
	if err := c.SearchRead(BaseLanguageInstallModel, criteria, options, blis); err != nil {
		return nil, err
	}
	return blis, nil
}

// FindBaseLanguageInstallIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBaseLanguageInstallIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BaseLanguageInstallModel, criteria, options)
}

// FindBaseLanguageInstallId finds record id by querying it with criteria.
func (c *Client) FindBaseLanguageInstallId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BaseLanguageInstallModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
