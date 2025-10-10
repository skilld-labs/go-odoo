package odoo

// SpreadsheetDashboard represents spreadsheet.dashboard model.
type SpreadsheetDashboard struct {
	CompanyId               *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate              *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid               *Many2One `xmlrpc:"create_uid,omitempty"`
	DashboardGroupId        *Many2One `xmlrpc:"dashboard_group_id,omitempty"`
	DisplayName             *String   `xmlrpc:"display_name,omitempty"`
	GroupIds                *Relation `xmlrpc:"group_ids,omitempty"`
	Id                      *Int      `xmlrpc:"id,omitempty"`
	IsPublished             *Bool     `xmlrpc:"is_published,omitempty"`
	MainDataModelIds        *Relation `xmlrpc:"main_data_model_ids,omitempty"`
	Name                    *String   `xmlrpc:"name,omitempty"`
	SampleDashboardFilePath *String   `xmlrpc:"sample_dashboard_file_path,omitempty"`
	Sequence                *Int      `xmlrpc:"sequence,omitempty"`
	SpreadsheetBinaryData   *String   `xmlrpc:"spreadsheet_binary_data,omitempty"`
	SpreadsheetData         *String   `xmlrpc:"spreadsheet_data,omitempty"`
	SpreadsheetFileName     *String   `xmlrpc:"spreadsheet_file_name,omitempty"`
	Thumbnail               *String   `xmlrpc:"thumbnail,omitempty"`
	WriteDate               *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SpreadsheetDashboards represents array of spreadsheet.dashboard model.
type SpreadsheetDashboards []SpreadsheetDashboard

// SpreadsheetDashboardModel is the odoo model name.
const SpreadsheetDashboardModel = "spreadsheet.dashboard"

// Many2One convert SpreadsheetDashboard to *Many2One.
func (sd *SpreadsheetDashboard) Many2One() *Many2One {
	return NewMany2One(sd.Id.Get(), "")
}

// CreateSpreadsheetDashboard creates a new spreadsheet.dashboard model and returns its id.
func (c *Client) CreateSpreadsheetDashboard(sd *SpreadsheetDashboard) (int64, error) {
	ids, err := c.CreateSpreadsheetDashboards([]*SpreadsheetDashboard{sd})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetDashboard creates a new spreadsheet.dashboard model and returns its id.
func (c *Client) CreateSpreadsheetDashboards(sds []*SpreadsheetDashboard) ([]int64, error) {
	var vv []interface{}
	for _, v := range sds {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetDashboardModel, vv, nil)
}

// UpdateSpreadsheetDashboard updates an existing spreadsheet.dashboard record.
func (c *Client) UpdateSpreadsheetDashboard(sd *SpreadsheetDashboard) error {
	return c.UpdateSpreadsheetDashboards([]int64{sd.Id.Get()}, sd)
}

// UpdateSpreadsheetDashboards updates existing spreadsheet.dashboard records.
// All records (represented by ids) will be updated by sd values.
func (c *Client) UpdateSpreadsheetDashboards(ids []int64, sd *SpreadsheetDashboard) error {
	return c.Update(SpreadsheetDashboardModel, ids, sd, nil)
}

// DeleteSpreadsheetDashboard deletes an existing spreadsheet.dashboard record.
func (c *Client) DeleteSpreadsheetDashboard(id int64) error {
	return c.DeleteSpreadsheetDashboards([]int64{id})
}

// DeleteSpreadsheetDashboards deletes existing spreadsheet.dashboard records.
func (c *Client) DeleteSpreadsheetDashboards(ids []int64) error {
	return c.Delete(SpreadsheetDashboardModel, ids)
}

// GetSpreadsheetDashboard gets spreadsheet.dashboard existing record.
func (c *Client) GetSpreadsheetDashboard(id int64) (*SpreadsheetDashboard, error) {
	sds, err := c.GetSpreadsheetDashboards([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sds)[0]), nil
}

// GetSpreadsheetDashboards gets spreadsheet.dashboard existing records.
func (c *Client) GetSpreadsheetDashboards(ids []int64) (*SpreadsheetDashboards, error) {
	sds := &SpreadsheetDashboards{}
	if err := c.Read(SpreadsheetDashboardModel, ids, nil, sds); err != nil {
		return nil, err
	}
	return sds, nil
}

// FindSpreadsheetDashboard finds spreadsheet.dashboard record by querying it with criteria.
func (c *Client) FindSpreadsheetDashboard(criteria *Criteria) (*SpreadsheetDashboard, error) {
	sds := &SpreadsheetDashboards{}
	if err := c.SearchRead(SpreadsheetDashboardModel, criteria, NewOptions().Limit(1), sds); err != nil {
		return nil, err
	}
	return &((*sds)[0]), nil
}

// FindSpreadsheetDashboards finds spreadsheet.dashboard records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboards(criteria *Criteria, options *Options) (*SpreadsheetDashboards, error) {
	sds := &SpreadsheetDashboards{}
	if err := c.SearchRead(SpreadsheetDashboardModel, criteria, options, sds); err != nil {
		return nil, err
	}
	return sds, nil
}

// FindSpreadsheetDashboardIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboardIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SpreadsheetDashboardModel, criteria, options)
}

// FindSpreadsheetDashboardId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetDashboardId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetDashboardModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
