package odoo

// CrmActivityReport represents crm.activity.report model.
type CrmActivityReport struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omitempty"`
	Active             *Bool     `xmlrpc:"active,omitempty"`
	AuthorId           *Many2One `xmlrpc:"author_id,omitempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omitempty"`
	CountryId          *Many2One `xmlrpc:"country_id,omitempty"`
	Date               *Time     `xmlrpc:"date,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	LeadId             *Many2One `xmlrpc:"lead_id,omitempty"`
	LeadType           *String   `xmlrpc:"lead_type,omitempty"`
	MailActivityTypeId *Many2One `xmlrpc:"mail_activity_type_id,omitempty"`
	PartnerId          *Many2One `xmlrpc:"partner_id,omitempty"`
	Probability        *Float    `xmlrpc:"probability,omitempty"`
	StageId            *Many2One `xmlrpc:"stage_id,omitempty"`
	Subject            *String   `xmlrpc:"subject,omitempty"`
	SubtypeId          *Many2One `xmlrpc:"subtype_id,omitempty"`
	TeamId             *Many2One `xmlrpc:"team_id,omitempty"`
	UserId             *Many2One `xmlrpc:"user_id,omitempty"`
}

// CrmActivityReports represents array of crm.activity.report model.
type CrmActivityReports []CrmActivityReport

// CrmActivityReportModel is the odoo model name.
const CrmActivityReportModel = "crm.activity.report"

// Many2One convert CrmActivityReport to *Many2One.
func (car *CrmActivityReport) Many2One() *Many2One {
	return NewMany2One(car.Id.Get(), "")
}

// CreateCrmActivityReport creates a new crm.activity.report model and returns its id.
func (c *Client) CreateCrmActivityReport(car *CrmActivityReport) (int64, error) {
	ids, err := c.CreateCrmActivityReports([]*CrmActivityReport{car})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmActivityReports creates a new crm.activity.report model and returns its id.
func (c *Client) CreateCrmActivityReports(cars []*CrmActivityReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range cars {
		vv = append(vv, v)
	}
	return c.Create(CrmActivityReportModel, vv, nil)
}

// UpdateCrmActivityReport updates an existing crm.activity.report record.
func (c *Client) UpdateCrmActivityReport(car *CrmActivityReport) error {
	return c.UpdateCrmActivityReports([]int64{car.Id.Get()}, car)
}

// UpdateCrmActivityReports updates existing crm.activity.report records.
// All records (represented by ids) will be updated by car values.
func (c *Client) UpdateCrmActivityReports(ids []int64, car *CrmActivityReport) error {
	return c.Update(CrmActivityReportModel, ids, car, nil)
}

// DeleteCrmActivityReport deletes an existing crm.activity.report record.
func (c *Client) DeleteCrmActivityReport(id int64) error {
	return c.DeleteCrmActivityReports([]int64{id})
}

// DeleteCrmActivityReports deletes existing crm.activity.report records.
func (c *Client) DeleteCrmActivityReports(ids []int64) error {
	return c.Delete(CrmActivityReportModel, ids)
}

// GetCrmActivityReport gets crm.activity.report existing record.
func (c *Client) GetCrmActivityReport(id int64) (*CrmActivityReport, error) {
	cars, err := c.GetCrmActivityReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cars)[0]), nil
}

// GetCrmActivityReports gets crm.activity.report existing records.
func (c *Client) GetCrmActivityReports(ids []int64) (*CrmActivityReports, error) {
	cars := &CrmActivityReports{}
	if err := c.Read(CrmActivityReportModel, ids, nil, cars); err != nil {
		return nil, err
	}
	return cars, nil
}

// FindCrmActivityReport finds crm.activity.report record by querying it with criteria.
func (c *Client) FindCrmActivityReport(criteria *Criteria) (*CrmActivityReport, error) {
	cars := &CrmActivityReports{}
	if err := c.SearchRead(CrmActivityReportModel, criteria, NewOptions().Limit(1), cars); err != nil {
		return nil, err
	}
	return &((*cars)[0]), nil
}

// FindCrmActivityReports finds crm.activity.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmActivityReports(criteria *Criteria, options *Options) (*CrmActivityReports, error) {
	cars := &CrmActivityReports{}
	if err := c.SearchRead(CrmActivityReportModel, criteria, options, cars); err != nil {
		return nil, err
	}
	return cars, nil
}

// FindCrmActivityReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmActivityReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmActivityReportModel, criteria, options)
}

// FindCrmActivityReportId finds record id by querying it with criteria.
func (c *Client) FindCrmActivityReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmActivityReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
