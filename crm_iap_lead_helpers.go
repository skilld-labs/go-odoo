package odoo

// CrmIapLeadHelpers represents crm.iap.lead.helpers model.
type CrmIapLeadHelpers struct {
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmIapLeadHelperss represents array of crm.iap.lead.helpers model.
type CrmIapLeadHelperss []CrmIapLeadHelpers

// CrmIapLeadHelpersModel is the odoo model name.
const CrmIapLeadHelpersModel = "crm.iap.lead.helpers"

// Many2One convert CrmIapLeadHelpers to *Many2One.
func (cilh *CrmIapLeadHelpers) Many2One() *Many2One {
	return NewMany2One(cilh.Id.Get(), "")
}

// CreateCrmIapLeadHelpers creates a new crm.iap.lead.helpers model and returns its id.
func (c *Client) CreateCrmIapLeadHelpers(cilh *CrmIapLeadHelpers) (int64, error) {
	ids, err := c.CreateCrmIapLeadHelperss([]*CrmIapLeadHelpers{cilh})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmIapLeadHelpers creates a new crm.iap.lead.helpers model and returns its id.
func (c *Client) CreateCrmIapLeadHelperss(cilhs []*CrmIapLeadHelpers) ([]int64, error) {
	var vv []interface{}
	for _, v := range cilhs {
		vv = append(vv, v)
	}
	return c.Create(CrmIapLeadHelpersModel, vv, nil)
}

// UpdateCrmIapLeadHelpers updates an existing crm.iap.lead.helpers record.
func (c *Client) UpdateCrmIapLeadHelpers(cilh *CrmIapLeadHelpers) error {
	return c.UpdateCrmIapLeadHelperss([]int64{cilh.Id.Get()}, cilh)
}

// UpdateCrmIapLeadHelperss updates existing crm.iap.lead.helpers records.
// All records (represented by ids) will be updated by cilh values.
func (c *Client) UpdateCrmIapLeadHelperss(ids []int64, cilh *CrmIapLeadHelpers) error {
	return c.Update(CrmIapLeadHelpersModel, ids, cilh, nil)
}

// DeleteCrmIapLeadHelpers deletes an existing crm.iap.lead.helpers record.
func (c *Client) DeleteCrmIapLeadHelpers(id int64) error {
	return c.DeleteCrmIapLeadHelperss([]int64{id})
}

// DeleteCrmIapLeadHelperss deletes existing crm.iap.lead.helpers records.
func (c *Client) DeleteCrmIapLeadHelperss(ids []int64) error {
	return c.Delete(CrmIapLeadHelpersModel, ids)
}

// GetCrmIapLeadHelpers gets crm.iap.lead.helpers existing record.
func (c *Client) GetCrmIapLeadHelpers(id int64) (*CrmIapLeadHelpers, error) {
	cilhs, err := c.GetCrmIapLeadHelperss([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cilhs)[0]), nil
}

// GetCrmIapLeadHelperss gets crm.iap.lead.helpers existing records.
func (c *Client) GetCrmIapLeadHelperss(ids []int64) (*CrmIapLeadHelperss, error) {
	cilhs := &CrmIapLeadHelperss{}
	if err := c.Read(CrmIapLeadHelpersModel, ids, nil, cilhs); err != nil {
		return nil, err
	}
	return cilhs, nil
}

// FindCrmIapLeadHelpers finds crm.iap.lead.helpers record by querying it with criteria.
func (c *Client) FindCrmIapLeadHelpers(criteria *Criteria) (*CrmIapLeadHelpers, error) {
	cilhs := &CrmIapLeadHelperss{}
	if err := c.SearchRead(CrmIapLeadHelpersModel, criteria, NewOptions().Limit(1), cilhs); err != nil {
		return nil, err
	}
	return &((*cilhs)[0]), nil
}

// FindCrmIapLeadHelperss finds crm.iap.lead.helpers records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadHelperss(criteria *Criteria, options *Options) (*CrmIapLeadHelperss, error) {
	cilhs := &CrmIapLeadHelperss{}
	if err := c.SearchRead(CrmIapLeadHelpersModel, criteria, options, cilhs); err != nil {
		return nil, err
	}
	return cilhs, nil
}

// FindCrmIapLeadHelpersIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadHelpersIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmIapLeadHelpersModel, criteria, options)
}

// FindCrmIapLeadHelpersId finds record id by querying it with criteria.
func (c *Client) FindCrmIapLeadHelpersId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmIapLeadHelpersModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
