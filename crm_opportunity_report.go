package odoo

// CrmOpportunityReport represents crm.opportunity.report model.
type CrmOpportunityReport struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omitempty"`
	Active              *Bool      `xmlrpc:"active,omitempty"`
	CampaignId          *Many2One  `xmlrpc:"campaign_id,omitempty"`
	City                *String    `xmlrpc:"city,omitempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omitempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omitempty"`
	DateClosed          *Time      `xmlrpc:"date_closed,omitempty"`
	DateConversion      *Time      `xmlrpc:"date_conversion,omitempty"`
	DateDeadline        *Time      `xmlrpc:"date_deadline,omitempty"`
	DateLastStageUpdate *Time      `xmlrpc:"date_last_stage_update,omitempty"`
	DelayClose          *Float     `xmlrpc:"delay_close,omitempty"`
	DelayExpected       *Float     `xmlrpc:"delay_expected,omitempty"`
	DelayOpen           *Float     `xmlrpc:"delay_open,omitempty"`
	DisplayName         *String    `xmlrpc:"display_name,omitempty"`
	ExpectedRevenue     *Float     `xmlrpc:"expected_revenue,omitempty"`
	Id                  *Int       `xmlrpc:"id,omitempty"`
	LostReason          *Many2One  `xmlrpc:"lost_reason,omitempty"`
	MediumId            *Many2One  `xmlrpc:"medium_id,omitempty"`
	NbrActivities       *Int       `xmlrpc:"nbr_activities,omitempty"`
	OpeningDate         *Time      `xmlrpc:"opening_date,omitempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omitempty"`
	Priority            *Selection `xmlrpc:"priority,omitempty"`
	Probability         *Float     `xmlrpc:"probability,omitempty"`
	SourceId            *Many2One  `xmlrpc:"source_id,omitempty"`
	StageId             *Many2One  `xmlrpc:"stage_id,omitempty"`
	StageName           *String    `xmlrpc:"stage_name,omitempty"`
	TeamId              *Many2One  `xmlrpc:"team_id,omitempty"`
	TotalRevenue        *Float     `xmlrpc:"total_revenue,omitempty"`
	Type                *Selection `xmlrpc:"type,omitempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omitempty"`
}

// CrmOpportunityReports represents array of crm.opportunity.report model.
type CrmOpportunityReports []CrmOpportunityReport

// CrmOpportunityReportModel is the odoo model name.
const CrmOpportunityReportModel = "crm.opportunity.report"

// Many2One convert CrmOpportunityReport to *Many2One.
func (cor *CrmOpportunityReport) Many2One() *Many2One {
	return NewMany2One(cor.Id.Get(), "")
}

// CreateCrmOpportunityReport creates a new crm.opportunity.report model and returns its id.
func (c *Client) CreateCrmOpportunityReport(cor *CrmOpportunityReport) (int64, error) {
	ids, err := c.CreateCrmOpportunityReports([]*CrmOpportunityReport{cor})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmOpportunityReports creates a new crm.opportunity.report model and returns its id.
func (c *Client) CreateCrmOpportunityReports(cors []*CrmOpportunityReport) ([]int64, error) {
	var vv []interface{}
	for _, v := range cors {
		vv = append(vv, v)
	}
	return c.Create(CrmOpportunityReportModel, vv, nil)
}

// UpdateCrmOpportunityReport updates an existing crm.opportunity.report record.
func (c *Client) UpdateCrmOpportunityReport(cor *CrmOpportunityReport) error {
	return c.UpdateCrmOpportunityReports([]int64{cor.Id.Get()}, cor)
}

// UpdateCrmOpportunityReports updates existing crm.opportunity.report records.
// All records (represented by ids) will be updated by cor values.
func (c *Client) UpdateCrmOpportunityReports(ids []int64, cor *CrmOpportunityReport) error {
	return c.Update(CrmOpportunityReportModel, ids, cor, nil)
}

// DeleteCrmOpportunityReport deletes an existing crm.opportunity.report record.
func (c *Client) DeleteCrmOpportunityReport(id int64) error {
	return c.DeleteCrmOpportunityReports([]int64{id})
}

// DeleteCrmOpportunityReports deletes existing crm.opportunity.report records.
func (c *Client) DeleteCrmOpportunityReports(ids []int64) error {
	return c.Delete(CrmOpportunityReportModel, ids)
}

// GetCrmOpportunityReport gets crm.opportunity.report existing record.
func (c *Client) GetCrmOpportunityReport(id int64) (*CrmOpportunityReport, error) {
	cors, err := c.GetCrmOpportunityReports([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cors)[0]), nil
}

// GetCrmOpportunityReports gets crm.opportunity.report existing records.
func (c *Client) GetCrmOpportunityReports(ids []int64) (*CrmOpportunityReports, error) {
	cors := &CrmOpportunityReports{}
	if err := c.Read(CrmOpportunityReportModel, ids, nil, cors); err != nil {
		return nil, err
	}
	return cors, nil
}

// FindCrmOpportunityReport finds crm.opportunity.report record by querying it with criteria.
func (c *Client) FindCrmOpportunityReport(criteria *Criteria) (*CrmOpportunityReport, error) {
	cors := &CrmOpportunityReports{}
	if err := c.SearchRead(CrmOpportunityReportModel, criteria, NewOptions().Limit(1), cors); err != nil {
		return nil, err
	}
	return &((*cors)[0]), nil
}

// FindCrmOpportunityReports finds crm.opportunity.report records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmOpportunityReports(criteria *Criteria, options *Options) (*CrmOpportunityReports, error) {
	cors := &CrmOpportunityReports{}
	if err := c.SearchRead(CrmOpportunityReportModel, criteria, options, cors); err != nil {
		return nil, err
	}
	return cors, nil
}

// FindCrmOpportunityReportIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmOpportunityReportIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmOpportunityReportModel, criteria, options)
}

// FindCrmOpportunityReportId finds record id by querying it with criteria.
func (c *Client) FindCrmOpportunityReportId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmOpportunityReportModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
