package odoo

// SpreadsheetDashboardGroup represents spreadsheet.dashboard.group model.
type SpreadsheetDashboardGroup struct {
	CreateDate            *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid             *Many2One `xmlrpc:"create_uid,omitempty"`
	DashboardIds          *Relation `xmlrpc:"dashboard_ids,omitempty"`
	DisplayName           *String   `xmlrpc:"display_name,omitempty"`
	Id                    *Int      `xmlrpc:"id,omitempty"`
	Name                  *String   `xmlrpc:"name,omitempty"`
	PublishedDashboardIds *Relation `xmlrpc:"published_dashboard_ids,omitempty"`
	Sequence              *Int      `xmlrpc:"sequence,omitempty"`
	WriteDate             *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid              *Many2One `xmlrpc:"write_uid,omitempty"`
}

// SpreadsheetDashboardGroups represents array of spreadsheet.dashboard.group model.
type SpreadsheetDashboardGroups []SpreadsheetDashboardGroup

// SpreadsheetDashboardGroupModel is the odoo model name.
const SpreadsheetDashboardGroupModel = "spreadsheet.dashboard.group"

// Many2One convert SpreadsheetDashboardGroup to *Many2One.
func (sdg *SpreadsheetDashboardGroup) Many2One() *Many2One {
	return NewMany2One(sdg.Id.Get(), "")
}

// CreateSpreadsheetDashboardGroup creates a new spreadsheet.dashboard.group model and returns its id.
func (c *Client) CreateSpreadsheetDashboardGroup(sdg *SpreadsheetDashboardGroup) (int64, error) {
	ids, err := c.CreateSpreadsheetDashboardGroups([]*SpreadsheetDashboardGroup{sdg})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateSpreadsheetDashboardGroup creates a new spreadsheet.dashboard.group model and returns its id.
func (c *Client) CreateSpreadsheetDashboardGroups(sdgs []*SpreadsheetDashboardGroup) ([]int64, error) {
	var vv []interface{}
	for _, v := range sdgs {
		vv = append(vv, v)
	}
	return c.Create(SpreadsheetDashboardGroupModel, vv, nil)
}

// UpdateSpreadsheetDashboardGroup updates an existing spreadsheet.dashboard.group record.
func (c *Client) UpdateSpreadsheetDashboardGroup(sdg *SpreadsheetDashboardGroup) error {
	return c.UpdateSpreadsheetDashboardGroups([]int64{sdg.Id.Get()}, sdg)
}

// UpdateSpreadsheetDashboardGroups updates existing spreadsheet.dashboard.group records.
// All records (represented by ids) will be updated by sdg values.
func (c *Client) UpdateSpreadsheetDashboardGroups(ids []int64, sdg *SpreadsheetDashboardGroup) error {
	return c.Update(SpreadsheetDashboardGroupModel, ids, sdg, nil)
}

// DeleteSpreadsheetDashboardGroup deletes an existing spreadsheet.dashboard.group record.
func (c *Client) DeleteSpreadsheetDashboardGroup(id int64) error {
	return c.DeleteSpreadsheetDashboardGroups([]int64{id})
}

// DeleteSpreadsheetDashboardGroups deletes existing spreadsheet.dashboard.group records.
func (c *Client) DeleteSpreadsheetDashboardGroups(ids []int64) error {
	return c.Delete(SpreadsheetDashboardGroupModel, ids)
}

// GetSpreadsheetDashboardGroup gets spreadsheet.dashboard.group existing record.
func (c *Client) GetSpreadsheetDashboardGroup(id int64) (*SpreadsheetDashboardGroup, error) {
	sdgs, err := c.GetSpreadsheetDashboardGroups([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*sdgs)[0]), nil
}

// GetSpreadsheetDashboardGroups gets spreadsheet.dashboard.group existing records.
func (c *Client) GetSpreadsheetDashboardGroups(ids []int64) (*SpreadsheetDashboardGroups, error) {
	sdgs := &SpreadsheetDashboardGroups{}
	if err := c.Read(SpreadsheetDashboardGroupModel, ids, nil, sdgs); err != nil {
		return nil, err
	}
	return sdgs, nil
}

// FindSpreadsheetDashboardGroup finds spreadsheet.dashboard.group record by querying it with criteria.
func (c *Client) FindSpreadsheetDashboardGroup(criteria *Criteria) (*SpreadsheetDashboardGroup, error) {
	sdgs := &SpreadsheetDashboardGroups{}
	if err := c.SearchRead(SpreadsheetDashboardGroupModel, criteria, NewOptions().Limit(1), sdgs); err != nil {
		return nil, err
	}
	return &((*sdgs)[0]), nil
}

// FindSpreadsheetDashboardGroups finds spreadsheet.dashboard.group records by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboardGroups(criteria *Criteria, options *Options) (*SpreadsheetDashboardGroups, error) {
	sdgs := &SpreadsheetDashboardGroups{}
	if err := c.SearchRead(SpreadsheetDashboardGroupModel, criteria, options, sdgs); err != nil {
		return nil, err
	}
	return sdgs, nil
}

// FindSpreadsheetDashboardGroupIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindSpreadsheetDashboardGroupIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(SpreadsheetDashboardGroupModel, criteria, options)
}

// FindSpreadsheetDashboardGroupId finds record id by querying it with criteria.
func (c *Client) FindSpreadsheetDashboardGroupId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(SpreadsheetDashboardGroupModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
