package odoo

// IapAutocompleteApi represents iap.autocomplete.api model.
type IapAutocompleteApi struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// IapAutocompleteApis represents array of iap.autocomplete.api model.
type IapAutocompleteApis []IapAutocompleteApi

// IapAutocompleteApiModel is the odoo model name.
const IapAutocompleteApiModel = "iap.autocomplete.api"

// Many2One convert IapAutocompleteApi to *Many2One.
func (iaa *IapAutocompleteApi) Many2One() *Many2One {
	return NewMany2One(iaa.Id.Get(), "")
}

// CreateIapAutocompleteApi creates a new iap.autocomplete.api model and returns its id.
func (c *Client) CreateIapAutocompleteApi(iaa *IapAutocompleteApi) (int64, error) {
	ids, err := c.CreateIapAutocompleteApis([]*IapAutocompleteApi{iaa})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateIapAutocompleteApi creates a new iap.autocomplete.api model and returns its id.
func (c *Client) CreateIapAutocompleteApis(iaas []*IapAutocompleteApi) ([]int64, error) {
	var vv []interface{}
	for _, v := range iaas {
		vv = append(vv, v)
	}
	return c.Create(IapAutocompleteApiModel, vv, nil)
}

// UpdateIapAutocompleteApi updates an existing iap.autocomplete.api record.
func (c *Client) UpdateIapAutocompleteApi(iaa *IapAutocompleteApi) error {
	return c.UpdateIapAutocompleteApis([]int64{iaa.Id.Get()}, iaa)
}

// UpdateIapAutocompleteApis updates existing iap.autocomplete.api records.
// All records (represented by ids) will be updated by iaa values.
func (c *Client) UpdateIapAutocompleteApis(ids []int64, iaa *IapAutocompleteApi) error {
	return c.Update(IapAutocompleteApiModel, ids, iaa, nil)
}

// DeleteIapAutocompleteApi deletes an existing iap.autocomplete.api record.
func (c *Client) DeleteIapAutocompleteApi(id int64) error {
	return c.DeleteIapAutocompleteApis([]int64{id})
}

// DeleteIapAutocompleteApis deletes existing iap.autocomplete.api records.
func (c *Client) DeleteIapAutocompleteApis(ids []int64) error {
	return c.Delete(IapAutocompleteApiModel, ids)
}

// GetIapAutocompleteApi gets iap.autocomplete.api existing record.
func (c *Client) GetIapAutocompleteApi(id int64) (*IapAutocompleteApi, error) {
	iaas, err := c.GetIapAutocompleteApis([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*iaas)[0]), nil
}

// GetIapAutocompleteApis gets iap.autocomplete.api existing records.
func (c *Client) GetIapAutocompleteApis(ids []int64) (*IapAutocompleteApis, error) {
	iaas := &IapAutocompleteApis{}
	if err := c.Read(IapAutocompleteApiModel, ids, nil, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIapAutocompleteApi finds iap.autocomplete.api record by querying it with criteria.
func (c *Client) FindIapAutocompleteApi(criteria *Criteria) (*IapAutocompleteApi, error) {
	iaas := &IapAutocompleteApis{}
	if err := c.SearchRead(IapAutocompleteApiModel, criteria, NewOptions().Limit(1), iaas); err != nil {
		return nil, err
	}
	return &((*iaas)[0]), nil
}

// FindIapAutocompleteApis finds iap.autocomplete.api records by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapAutocompleteApis(criteria *Criteria, options *Options) (*IapAutocompleteApis, error) {
	iaas := &IapAutocompleteApis{}
	if err := c.SearchRead(IapAutocompleteApiModel, criteria, options, iaas); err != nil {
		return nil, err
	}
	return iaas, nil
}

// FindIapAutocompleteApiIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindIapAutocompleteApiIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(IapAutocompleteApiModel, criteria, options)
}

// FindIapAutocompleteApiId finds record id by querying it with criteria.
func (c *Client) FindIapAutocompleteApiId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(IapAutocompleteApiModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
