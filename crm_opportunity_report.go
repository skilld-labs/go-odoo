package odoo

// CrmOpportunityReport represents crm.opportunity.report model.
type CrmOpportunityReport struct {
	LastUpdate          *Time      `xmlrpc:"__last_update,omptempty"`
	Active              *Bool      `xmlrpc:"active,omptempty"`
	CampaignId          *Many2One  `xmlrpc:"campaign_id,omptempty"`
	City                *String    `xmlrpc:"city,omptempty"`
	CompanyId           *Many2One  `xmlrpc:"company_id,omptempty"`
	CountryId           *Many2One  `xmlrpc:"country_id,omptempty"`
	CreateDate          *Time      `xmlrpc:"create_date,omptempty"`
	DateClosed          *Time      `xmlrpc:"date_closed,omptempty"`
	DateConversion      *Time      `xmlrpc:"date_conversion,omptempty"`
	DateDeadline        *Time      `xmlrpc:"date_deadline,omptempty"`
	DateLastStageUpdate *Time      `xmlrpc:"date_last_stage_update,omptempty"`
	DelayClose          *Float     `xmlrpc:"delay_close,omptempty"`
	DelayExpected       *Float     `xmlrpc:"delay_expected,omptempty"`
	DelayOpen           *Float     `xmlrpc:"delay_open,omptempty"`
	DisplayName         *String    `xmlrpc:"display_name,omptempty"`
	ExpectedRevenue     *Float     `xmlrpc:"expected_revenue,omptempty"`
	Id                  *Int       `xmlrpc:"id,omptempty"`
	LostReason          *Many2One  `xmlrpc:"lost_reason,omptempty"`
	MediumId            *Many2One  `xmlrpc:"medium_id,omptempty"`
	NbrActivities       *Int       `xmlrpc:"nbr_activities,omptempty"`
	OpeningDate         *Time      `xmlrpc:"opening_date,omptempty"`
	PartnerId           *Many2One  `xmlrpc:"partner_id,omptempty"`
	Priority            *Selection `xmlrpc:"priority,omptempty"`
	Probability         *Float     `xmlrpc:"probability,omptempty"`
	SourceId            *Many2One  `xmlrpc:"source_id,omptempty"`
	StageId             *Many2One  `xmlrpc:"stage_id,omptempty"`
	StageName           *String    `xmlrpc:"stage_name,omptempty"`
	TeamId              *Many2One  `xmlrpc:"team_id,omptempty"`
	TotalRevenue        *Float     `xmlrpc:"total_revenue,omptempty"`
	Type                *Selection `xmlrpc:"type,omptempty"`
	UserId              *Many2One  `xmlrpc:"user_id,omptempty"`
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

// CreateCrmOpportunityReport creates a new crm.opportunity.report model and returns its id.
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
