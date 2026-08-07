package odoo

// IapEnrichApi represents iap.enrich.api model.
type IapEnrichApi struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IapEnrichApis represents array of iap.enrich.api model.
type IapEnrichApis []IapEnrichApi

// IapEnrichApiModel is the odoo model name.
const IapEnrichApiModel = "iap.enrich.api"

// Many2One convert IapEnrichApi to *Many2One.
func (iea *IapEnrichApi) Many2One() *Many2One {
	return NewMany2One(iea.Id.Get(), "")
}

// CreateIapEnrichApi creates a new iap.enrich.api model and returns its id.
func (c *Client) CreateIapEnrichApi(iea *IapEnrichApi) (int64, error) {
	ids, err := c.CreateIapEnrichApis([]*IapEnrichApi{iea})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIapEnrichApi creates a new iap.enrich.api model and returns its id.
func (c *Client) CreateIapEnrichApis(ieas []*IapEnrichApi) ([]int64, error) {
	var vv []interface{}
	for _, v := range ieas {
		vv = append(vv, v)
	}
	return c.Create(IapEnrichApiModel, vv, nil)
}

// UpdateIapEnrichApi updates an existing iap.enrich.api record.
func (c *Client) UpdateIapEnrichApi(iea *IapEnrichApi) error {
	return c.UpdateIapEnrichApis([]int64{iea.Id.Get()}, iea)
}

// UpdateIapEnrichApis updates existing iap.enrich.api records.
// All records (represented by ids) will be updated by iea values.
func (c *Client) UpdateIapEnrichApis(ids []int64, iea *IapEnrichApi) error {
	return c.Update(IapEnrichApiModel, ids, iea, nil)
}

// DeleteIapEnrichApi deletes an existing iap.enrich.api record.
func (c *Client) DeleteIapEnrichApi(id int64) error {
	return c.DeleteIapEnrichApis([]int64{id})
}

// DeleteIapEnrichApis deletes existing iap.enrich.api records.
func (c *Client) DeleteIapEnrichApis(ids []int64) error {
	return c.Delete(IapEnrichApiModel, ids)
}

// GetIapEnrichApi gets iap.enrich.api existing record.
func (c *Client) GetIapEnrichApi(id int64) (*IapEnrichApi, error) {
	ieas, err := c.GetIapEnrichApis([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ieas)[0]), nil
}

// GetIapEnrichApis gets iap.enrich.api existing records.
func (c *Client) GetIapEnrichApis(ids []int64) (*IapEnrichApis, error) {
	ieas := &IapEnrichApis{}
	if err := c.Read(IapEnrichApiModel, ids, nil, ieas); err != nil {
		return nil, err
	}
	return ieas, nil
}

// FindIapEnrichApi finds iap.enrich.api record by querying it with criteria.
func (c *Client) FindIapEnrichApi(criteria *Criteria) (*IapEnrichApi, error) {
	ieas := &IapEnrichApis{}
	if err := c.SearchRead(IapEnrichApiModel, criteria, NewOptions().Limit(1), ieas); err != nil {
		return nil, err
	}
	return &((*ieas)[0]), nil
}

// FindIapEnrichApis finds iap.enrich.api records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapEnrichApis(criteria *Criteria, options *Options) (*IapEnrichApis, error) {
	ieas := &IapEnrichApis{}
	if err := c.SearchRead(IapEnrichApiModel, criteria, options, ieas); err != nil {
		return nil, err
	}
	return ieas, nil
}

// FindIapEnrichApiIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapEnrichApiIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IapEnrichApiModel, criteria, options)
}

// FindIapEnrichApiId finds record id by querying it with criteria.
func (c *Client) FindIapEnrichApiId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IapEnrichApiModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
