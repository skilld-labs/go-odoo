package odoo

// HrPayrollStructureType represents hr.payroll.structure.type model.
type HrPayrollStructureType struct {
	CountryCode               *String   `xmlrpc:"country_code,omitempty"`
	CountryId                 *Many2One `xmlrpc:"country_id,omitempty"`
	CreateDate                *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                 *Many2One `xmlrpc:"create_uid,omitempty"`
	DefaultResourceCalendarId *Many2One `xmlrpc:"default_resource_calendar_id,omitempty"`
	DisplayName               *String   `xmlrpc:"display_name,omitempty"`
	Id                        *Int      `xmlrpc:"id,omitempty"`
	Name                      *String   `xmlrpc:"name,omitempty"`
	WriteDate                 *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                  *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrPayrollStructureTypes represents array of hr.payroll.structure.type model.
type HrPayrollStructureTypes []HrPayrollStructureType

// HrPayrollStructureTypeModel is the odoo model name.
const HrPayrollStructureTypeModel = "hr.payroll.structure.type"

// Many2One convert HrPayrollStructureType to *Many2One.
func (hpst *HrPayrollStructureType) Many2One() *Many2One {
	return NewMany2One(hpst.Id.Get(), "")
}

// CreateHrPayrollStructureType creates a new hr.payroll.structure.type model and returns its id.
func (c *Client) CreateHrPayrollStructureType(hpst *HrPayrollStructureType) (int64, error) {
	ids, err := c.CreateHrPayrollStructureTypes([]*HrPayrollStructureType{hpst})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrPayrollStructureType creates a new hr.payroll.structure.type model and returns its id.
func (c *Client) CreateHrPayrollStructureTypes(hpsts []*HrPayrollStructureType) ([]int64, error) {
	var vv []interface{}
	for _, v := range hpsts {
		vv = append(vv, v)
	}
	return c.Create(HrPayrollStructureTypeModel, vv, nil)
}

// UpdateHrPayrollStructureType updates an existing hr.payroll.structure.type record.
func (c *Client) UpdateHrPayrollStructureType(hpst *HrPayrollStructureType) error {
	return c.UpdateHrPayrollStructureTypes([]int64{hpst.Id.Get()}, hpst)
}

// UpdateHrPayrollStructureTypes updates existing hr.payroll.structure.type records.
// All records (represented by ids) will be updated by hpst values.
func (c *Client) UpdateHrPayrollStructureTypes(ids []int64, hpst *HrPayrollStructureType) error {
	return c.Update(HrPayrollStructureTypeModel, ids, hpst, nil)
}

// DeleteHrPayrollStructureType deletes an existing hr.payroll.structure.type record.
func (c *Client) DeleteHrPayrollStructureType(id int64) error {
	return c.DeleteHrPayrollStructureTypes([]int64{id})
}

// DeleteHrPayrollStructureTypes deletes existing hr.payroll.structure.type records.
func (c *Client) DeleteHrPayrollStructureTypes(ids []int64) error {
	return c.Delete(HrPayrollStructureTypeModel, ids)
}

// GetHrPayrollStructureType gets hr.payroll.structure.type existing record.
func (c *Client) GetHrPayrollStructureType(id int64) (*HrPayrollStructureType, error) {
	hpsts, err := c.GetHrPayrollStructureTypes([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hpsts)[0]), nil
}

// GetHrPayrollStructureTypes gets hr.payroll.structure.type existing records.
func (c *Client) GetHrPayrollStructureTypes(ids []int64) (*HrPayrollStructureTypes, error) {
	hpsts := &HrPayrollStructureTypes{}
	if err := c.Read(HrPayrollStructureTypeModel, ids, nil, hpsts); err != nil {
		return nil, err
	}
	return hpsts, nil
}

// FindHrPayrollStructureType finds hr.payroll.structure.type record by querying it with criteria.
func (c *Client) FindHrPayrollStructureType(criteria *Criteria) (*HrPayrollStructureType, error) {
	hpsts := &HrPayrollStructureTypes{}
	if err := c.SearchRead(HrPayrollStructureTypeModel, criteria, NewOptions().Limit(1), hpsts); err != nil {
		return nil, err
	}
	return &((*hpsts)[0]), nil
}

// FindHrPayrollStructureTypes finds hr.payroll.structure.type records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollStructureTypes(criteria *Criteria, options *Options) (*HrPayrollStructureTypes, error) {
	hpsts := &HrPayrollStructureTypes{}
	if err := c.SearchRead(HrPayrollStructureTypeModel, criteria, options, hpsts); err != nil {
		return nil, err
	}
	return hpsts, nil
}

// FindHrPayrollStructureTypeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrPayrollStructureTypeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrPayrollStructureTypeModel, criteria, options)
}

// FindHrPayrollStructureTypeId finds record id by querying it with criteria.
func (c *Client) FindHrPayrollStructureTypeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrPayrollStructureTypeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
