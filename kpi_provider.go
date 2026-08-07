package odoo

// KpiProvider represents kpi.provider model.
type KpiProvider struct {
	DisplayName *String `xmlrpc:"display_name,omitempty"`
	Id          *Int    `xmlrpc:"id,omitempty"`
}

// KpiProviders represents array of kpi.provider model.
type KpiProviders []KpiProvider

// KpiProviderModel is the odoo model name.
const KpiProviderModel = "kpi.provider"

// Many2One convert KpiProvider to *Many2One.
func (kp *KpiProvider) Many2One() *Many2One {
	return NewMany2One(kp.Id.Get(), "")
}

// CreateKpiProvider creates a new kpi.provider model and returns its id.
func (c *Client) CreateKpiProvider(kp *KpiProvider) (int64, error) {
	ids, err := c.CreateKpiProviders([]*KpiProvider{kp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateKpiProvider creates a new kpi.provider model and returns its id.
func (c *Client) CreateKpiProviders(kps []*KpiProvider) ([]int64, error) {
	var vv []interface{}
	for _, v := range kps {
		vv = append(vv, v)
	}
	return c.Create(KpiProviderModel, vv, nil)
}

// UpdateKpiProvider updates an existing kpi.provider record.
func (c *Client) UpdateKpiProvider(kp *KpiProvider) error {
	return c.UpdateKpiProviders([]int64{kp.Id.Get()}, kp)
}

// UpdateKpiProviders updates existing kpi.provider records.
// All records (represented by ids) will be updated by kp values.
func (c *Client) UpdateKpiProviders(ids []int64, kp *KpiProvider) error {
	return c.Update(KpiProviderModel, ids, kp, nil)
}

// DeleteKpiProvider deletes an existing kpi.provider record.
func (c *Client) DeleteKpiProvider(id int64) error {
	return c.DeleteKpiProviders([]int64{id})
}

// DeleteKpiProviders deletes existing kpi.provider records.
func (c *Client) DeleteKpiProviders(ids []int64) error {
	return c.Delete(KpiProviderModel, ids)
}

// GetKpiProvider gets kpi.provider existing record.
func (c *Client) GetKpiProvider(id int64) (*KpiProvider, error) {
	kps, err := c.GetKpiProviders([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*kps)[0]), nil
}

// GetKpiProviders gets kpi.provider existing records.
func (c *Client) GetKpiProviders(ids []int64) (*KpiProviders, error) {
	kps := &KpiProviders{}
	if err := c.Read(KpiProviderModel, ids, nil, kps); err != nil {
		return nil, err
	}
	return kps, nil
}

// FindKpiProvider finds kpi.provider record by querying it with criteria.
func (c *Client) FindKpiProvider(criteria *Criteria) (*KpiProvider, error) {
	kps := &KpiProviders{}
	if err := c.SearchRead(KpiProviderModel, criteria, NewOptions().Limit(1), kps); err != nil {
		return nil, err
	}
	return &((*kps)[0]), nil
}

// FindKpiProviders finds kpi.provider records by querying it
// and filtering it with criteria and options.
func (c *Client) FindKpiProviders(criteria *Criteria, options *Options) (*KpiProviders, error) {
	kps := &KpiProviders{}
	if err := c.SearchRead(KpiProviderModel, criteria, options, kps); err != nil {
		return nil, err
	}
	return kps, nil
}

// FindKpiProviderIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindKpiProviderIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(KpiProviderModel, criteria, options)
}

// FindKpiProviderId finds record id by querying it with criteria.
func (c *Client) FindKpiProviderId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(KpiProviderModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
