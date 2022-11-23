package odoo

import (
	"fmt"
)

// CrmActivityReport represents crm.activity.report model.
type CrmActivityReport struct {
	LastUpdate         *Time     `xmlrpc:"__last_update,omptempty"`
	Active             *Bool     `xmlrpc:"active,omptempty"`
	AuthorId           *Many2One `xmlrpc:"author_id,omptempty"`
	CompanyId          *Many2One `xmlrpc:"company_id,omptempty"`
	CountryId          *Many2One `xmlrpc:"country_id,omptempty"`
	Date               *Time     `xmlrpc:"date,omptempty"`
	DisplayName        *String   `xmlrpc:"display_name,omptempty"`
	Id                 *Int      `xmlrpc:"id,omptempty"`
	LeadId             *Many2One `xmlrpc:"lead_id,omptempty"`
	LeadType           *String   `xmlrpc:"lead_type,omptempty"`
	MailActivityTypeId *Many2One `xmlrpc:"mail_activity_type_id,omptempty"`
	PartnerId          *Many2One `xmlrpc:"partner_id,omptempty"`
	Probability        *Float    `xmlrpc:"probability,omptempty"`
	StageId            *Many2One `xmlrpc:"stage_id,omptempty"`
	Subject            *String   `xmlrpc:"subject,omptempty"`
	SubtypeId          *Many2One `xmlrpc:"subtype_id,omptempty"`
	TeamId             *Many2One `xmlrpc:"team_id,omptempty"`
	UserId             *Many2One `xmlrpc:"user_id,omptempty"`
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
	return c.Create(CrmActivityReportModel, car)
}

// UpdateCrmActivityReport updates an existing crm.activity.report record.
func (c *Client) UpdateCrmActivityReport(car *CrmActivityReport) error {
	return c.UpdateCrmActivityReports([]int64{car.Id.Get()}, car)
}

// UpdateCrmActivityReports updates existing crm.activity.report records.
// All records (represented by ids) will be updated by car values.
func (c *Client) UpdateCrmActivityReports(ids []int64, car *CrmActivityReport) error {
	return c.Update(CrmActivityReportModel, ids, car)
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
	if cars != nil && len(*cars) > 0 {
		return &((*cars)[0]), nil
	}
	return nil, fmt.Errorf("id %v of crm.activity.report not found", id)
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
	if cars != nil && len(*cars) > 0 {
		return &((*cars)[0]), nil
	}
	return nil, fmt.Errorf("crm.activity.report was not found with criteria %v", criteria)
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
	ids, err := c.Search(CrmActivityReportModel, criteria, options)
	if err != nil {
		return []int64{}, err
	}
	return ids, nil
}

// FindCrmActivityReportId finds record id by querying it with criteria.
func (c *Client) FindCrmActivityReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmActivityReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	if len(ids) > 0 {
		return ids[0], nil
	}
	return -1, fmt.Errorf("crm.activity.report was not found with criteria %v and options %v", criteria, options)
}
