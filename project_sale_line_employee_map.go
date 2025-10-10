package odoo

// ProjectSaleLineEmployeeMap represents project.sale.line.employee.map model.
type ProjectSaleLineEmployeeMap struct {
	CompanyId           *Many2One `xmlrpc:"company_id,omitempty"`
	Cost                *Float    `xmlrpc:"cost,omitempty"`
	CostCurrencyId      *Many2One `xmlrpc:"cost_currency_id,omitempty"`
	CreateDate          *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid           *Many2One `xmlrpc:"create_uid,omitempty"`
	CurrencyId          *Many2One `xmlrpc:"currency_id,omitempty"`
	DisplayCost         *Float    `xmlrpc:"display_cost,omitempty"`
	DisplayName         *String   `xmlrpc:"display_name,omitempty"`
	EmployeeId          *Many2One `xmlrpc:"employee_id,omitempty"`
	ExistingEmployeeIds *Relation `xmlrpc:"existing_employee_ids,omitempty"`
	Id                  *Int      `xmlrpc:"id,omitempty"`
	IsCostChanged       *Bool     `xmlrpc:"is_cost_changed,omitempty"`
	PartnerId           *Many2One `xmlrpc:"partner_id,omitempty"`
	PriceUnit           *Float    `xmlrpc:"price_unit,omitempty"`
	ProjectId           *Many2One `xmlrpc:"project_id,omitempty"`
	SaleLineId          *Many2One `xmlrpc:"sale_line_id,omitempty"`
	SaleOrderId         *Many2One `xmlrpc:"sale_order_id,omitempty"`
	WriteDate           *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid            *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ProjectSaleLineEmployeeMaps represents array of project.sale.line.employee.map model.
type ProjectSaleLineEmployeeMaps []ProjectSaleLineEmployeeMap

// ProjectSaleLineEmployeeMapModel is the odoo model name.
const ProjectSaleLineEmployeeMapModel = "project.sale.line.employee.map"

// Many2One convert ProjectSaleLineEmployeeMap to *Many2One.
func (pslem *ProjectSaleLineEmployeeMap) Many2One() *Many2One {
	return NewMany2One(pslem.Id.Get(), "")
}

// CreateProjectSaleLineEmployeeMap creates a new project.sale.line.employee.map model and returns its id.
func (c *Client) CreateProjectSaleLineEmployeeMap(pslem *ProjectSaleLineEmployeeMap) (int64, error) {
	ids, err := c.CreateProjectSaleLineEmployeeMaps([]*ProjectSaleLineEmployeeMap{pslem})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectSaleLineEmployeeMap creates a new project.sale.line.employee.map model and returns its id.
func (c *Client) CreateProjectSaleLineEmployeeMaps(pslems []*ProjectSaleLineEmployeeMap) ([]int64, error) {
	var vv []interface{}
	for _, v := range pslems {
		vv = append(vv, v)
	}
	return c.Create(ProjectSaleLineEmployeeMapModel, vv, nil)
}

// UpdateProjectSaleLineEmployeeMap updates an existing project.sale.line.employee.map record.
func (c *Client) UpdateProjectSaleLineEmployeeMap(pslem *ProjectSaleLineEmployeeMap) error {
	return c.UpdateProjectSaleLineEmployeeMaps([]int64{pslem.Id.Get()}, pslem)
}

// UpdateProjectSaleLineEmployeeMaps updates existing project.sale.line.employee.map records.
// All records (represented by ids) will be updated by pslem values.
func (c *Client) UpdateProjectSaleLineEmployeeMaps(ids []int64, pslem *ProjectSaleLineEmployeeMap) error {
	return c.Update(ProjectSaleLineEmployeeMapModel, ids, pslem, nil)
}

// DeleteProjectSaleLineEmployeeMap deletes an existing project.sale.line.employee.map record.
func (c *Client) DeleteProjectSaleLineEmployeeMap(id int64) error {
	return c.DeleteProjectSaleLineEmployeeMaps([]int64{id})
}

// DeleteProjectSaleLineEmployeeMaps deletes existing project.sale.line.employee.map records.
func (c *Client) DeleteProjectSaleLineEmployeeMaps(ids []int64) error {
	return c.Delete(ProjectSaleLineEmployeeMapModel, ids)
}

// GetProjectSaleLineEmployeeMap gets project.sale.line.employee.map existing record.
func (c *Client) GetProjectSaleLineEmployeeMap(id int64) (*ProjectSaleLineEmployeeMap, error) {
	pslems, err := c.GetProjectSaleLineEmployeeMaps([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pslems)[0]), nil
}

// GetProjectSaleLineEmployeeMaps gets project.sale.line.employee.map existing records.
func (c *Client) GetProjectSaleLineEmployeeMaps(ids []int64) (*ProjectSaleLineEmployeeMaps, error) {
	pslems := &ProjectSaleLineEmployeeMaps{}
	if err := c.Read(ProjectSaleLineEmployeeMapModel, ids, nil, pslems); err != nil {
		return nil, err
	}
	return pslems, nil
}

// FindProjectSaleLineEmployeeMap finds project.sale.line.employee.map record by querying it with criteria.
func (c *Client) FindProjectSaleLineEmployeeMap(criteria *Criteria) (*ProjectSaleLineEmployeeMap, error) {
	pslems := &ProjectSaleLineEmployeeMaps{}
	if err := c.SearchRead(ProjectSaleLineEmployeeMapModel, criteria, NewOptions().Limit(1), pslems); err != nil {
		return nil, err
	}
	return &((*pslems)[0]), nil
}

// FindProjectSaleLineEmployeeMaps finds project.sale.line.employee.map records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectSaleLineEmployeeMaps(criteria *Criteria, options *Options) (*ProjectSaleLineEmployeeMaps, error) {
	pslems := &ProjectSaleLineEmployeeMaps{}
	if err := c.SearchRead(ProjectSaleLineEmployeeMapModel, criteria, options, pslems); err != nil {
		return nil, err
	}
	return pslems, nil
}

// FindProjectSaleLineEmployeeMapIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectSaleLineEmployeeMapIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectSaleLineEmployeeMapModel, criteria, options)
}

// FindProjectSaleLineEmployeeMapId finds record id by querying it with criteria.
func (c *Client) FindProjectSaleLineEmployeeMapId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectSaleLineEmployeeMapModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
