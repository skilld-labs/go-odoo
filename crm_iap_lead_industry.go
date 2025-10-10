package odoo

// CrmIapLeadIndustry represents crm.iap.lead.industry model.
type CrmIapLeadIndustry struct {
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	RevealIds   *String   `xmlrpc:"reveal_ids,omitempty"`
	Sequence    *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmIapLeadIndustrys represents array of crm.iap.lead.industry model.
type CrmIapLeadIndustrys []CrmIapLeadIndustry

// CrmIapLeadIndustryModel is the odoo model name.
const CrmIapLeadIndustryModel = "crm.iap.lead.industry"

// Many2One convert CrmIapLeadIndustry to *Many2One.
func (cili *CrmIapLeadIndustry) Many2One() *Many2One {
	return NewMany2One(cili.Id.Get(), "")
}

// CreateCrmIapLeadIndustry creates a new crm.iap.lead.industry model and returns its id.
func (c *Client) CreateCrmIapLeadIndustry(cili *CrmIapLeadIndustry) (int64, error) {
	ids, err := c.CreateCrmIapLeadIndustrys([]*CrmIapLeadIndustry{cili})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmIapLeadIndustry creates a new crm.iap.lead.industry model and returns its id.
func (c *Client) CreateCrmIapLeadIndustrys(cilis []*CrmIapLeadIndustry) ([]int64, error) {
	var vv []interface{}
	for _, v := range cilis {
		vv = append(vv, v)
	}
	return c.Create(CrmIapLeadIndustryModel, vv, nil)
}

// UpdateCrmIapLeadIndustry updates an existing crm.iap.lead.industry record.
func (c *Client) UpdateCrmIapLeadIndustry(cili *CrmIapLeadIndustry) error {
	return c.UpdateCrmIapLeadIndustrys([]int64{cili.Id.Get()}, cili)
}

// UpdateCrmIapLeadIndustrys updates existing crm.iap.lead.industry records.
// All records (represented by ids) will be updated by cili values.
func (c *Client) UpdateCrmIapLeadIndustrys(ids []int64, cili *CrmIapLeadIndustry) error {
	return c.Update(CrmIapLeadIndustryModel, ids, cili, nil)
}

// DeleteCrmIapLeadIndustry deletes an existing crm.iap.lead.industry record.
func (c *Client) DeleteCrmIapLeadIndustry(id int64) error {
	return c.DeleteCrmIapLeadIndustrys([]int64{id})
}

// DeleteCrmIapLeadIndustrys deletes existing crm.iap.lead.industry records.
func (c *Client) DeleteCrmIapLeadIndustrys(ids []int64) error {
	return c.Delete(CrmIapLeadIndustryModel, ids)
}

// GetCrmIapLeadIndustry gets crm.iap.lead.industry existing record.
func (c *Client) GetCrmIapLeadIndustry(id int64) (*CrmIapLeadIndustry, error) {
	cilis, err := c.GetCrmIapLeadIndustrys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cilis)[0]), nil
}

// GetCrmIapLeadIndustrys gets crm.iap.lead.industry existing records.
func (c *Client) GetCrmIapLeadIndustrys(ids []int64) (*CrmIapLeadIndustrys, error) {
	cilis := &CrmIapLeadIndustrys{}
	if err := c.Read(CrmIapLeadIndustryModel, ids, nil, cilis); err != nil {
		return nil, err
	}
	return cilis, nil
}

// FindCrmIapLeadIndustry finds crm.iap.lead.industry record by querying it with criteria.
func (c *Client) FindCrmIapLeadIndustry(criteria *Criteria) (*CrmIapLeadIndustry, error) {
	cilis := &CrmIapLeadIndustrys{}
	if err := c.SearchRead(CrmIapLeadIndustryModel, criteria, NewOptions().Limit(1), cilis); err != nil {
		return nil, err
	}
	return &((*cilis)[0]), nil
}

// FindCrmIapLeadIndustrys finds crm.iap.lead.industry records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadIndustrys(criteria *Criteria, options *Options) (*CrmIapLeadIndustrys, error) {
	cilis := &CrmIapLeadIndustrys{}
	if err := c.SearchRead(CrmIapLeadIndustryModel, criteria, options, cilis); err != nil {
		return nil, err
	}
	return cilis, nil
}

// FindCrmIapLeadIndustryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmIapLeadIndustryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmIapLeadIndustryModel, criteria, options)
}

// FindCrmIapLeadIndustryId finds record id by querying it with criteria.
func (c *Client) FindCrmIapLeadIndustryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmIapLeadIndustryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
