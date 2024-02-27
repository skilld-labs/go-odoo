package odoo

// HrEmployeeCategory represents hr.employee.category model.
type HrEmployeeCategory struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	Color       *Int      `xmlrpc:"color,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	EmployeeIds *Relation `xmlrpc:"employee_ids,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	Name        *String   `xmlrpc:"name,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// HrEmployeeCategorys represents array of hr.employee.category model.
type HrEmployeeCategorys []HrEmployeeCategory

// HrEmployeeCategoryModel is the odoo model name.
const HrEmployeeCategoryModel = "hr.employee.category"

// Many2One convert HrEmployeeCategory to *Many2One.
func (hec *HrEmployeeCategory) Many2One() *Many2One {
	return NewMany2One(hec.Id.Get(), "")
}

// CreateHrEmployeeCategory creates a new hr.employee.category model and returns its id.
func (c *Client) CreateHrEmployeeCategory(hec *HrEmployeeCategory) (int64, error) {
	ids, err := c.CreateHrEmployeeCategorys([]*HrEmployeeCategory{hec})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployeeCategorys creates a new hr.employee.category model and returns its id.
func (c *Client) CreateHrEmployeeCategorys(hecs []*HrEmployeeCategory) ([]int64, error) {
	var vv []interface{}
	for _, v := range hecs {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeCategoryModel, vv, nil)
}

// UpdateHrEmployeeCategory updates an existing hr.employee.category record.
func (c *Client) UpdateHrEmployeeCategory(hec *HrEmployeeCategory) error {
	return c.UpdateHrEmployeeCategorys([]int64{hec.Id.Get()}, hec)
}

// UpdateHrEmployeeCategorys updates existing hr.employee.category records.
// All records (represented by ids) will be updated by hec values.
func (c *Client) UpdateHrEmployeeCategorys(ids []int64, hec *HrEmployeeCategory) error {
	return c.Update(HrEmployeeCategoryModel, ids, hec, nil)
}

// DeleteHrEmployeeCategory deletes an existing hr.employee.category record.
func (c *Client) DeleteHrEmployeeCategory(id int64) error {
	return c.DeleteHrEmployeeCategorys([]int64{id})
}

// DeleteHrEmployeeCategorys deletes existing hr.employee.category records.
func (c *Client) DeleteHrEmployeeCategorys(ids []int64) error {
	return c.Delete(HrEmployeeCategoryModel, ids)
}

// GetHrEmployeeCategory gets hr.employee.category existing record.
func (c *Client) GetHrEmployeeCategory(id int64) (*HrEmployeeCategory, error) {
	hecs, err := c.GetHrEmployeeCategorys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hecs)[0]), nil
}

// GetHrEmployeeCategorys gets hr.employee.category existing records.
func (c *Client) GetHrEmployeeCategorys(ids []int64) (*HrEmployeeCategorys, error) {
	hecs := &HrEmployeeCategorys{}
	if err := c.Read(HrEmployeeCategoryModel, ids, nil, hecs); err != nil {
		return nil, err
	}
	return hecs, nil
}

// FindHrEmployeeCategory finds hr.employee.category record by querying it with criteria.
func (c *Client) FindHrEmployeeCategory(criteria *Criteria) (*HrEmployeeCategory, error) {
	hecs := &HrEmployeeCategorys{}
	if err := c.SearchRead(HrEmployeeCategoryModel, criteria, NewOptions().Limit(1), hecs); err != nil {
		return nil, err
	}
	return &((*hecs)[0]), nil
}

// FindHrEmployeeCategorys finds hr.employee.category records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeCategorys(criteria *Criteria, options *Options) (*HrEmployeeCategorys, error) {
	hecs := &HrEmployeeCategorys{}
	if err := c.SearchRead(HrEmployeeCategoryModel, criteria, options, hecs); err != nil {
		return nil, err
	}
	return hecs, nil
}

// FindHrEmployeeCategoryIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeCategoryIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeCategoryModel, criteria, options)
}

// FindHrEmployeeCategoryId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeCategoryId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeCategoryModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
