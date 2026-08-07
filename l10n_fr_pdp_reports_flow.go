package odoo

// L10NFrPdpReportsFlow represents l10n.fr.pdp.reports.flow model.
type L10NFrPdpReportsFlow struct {
	ActivityCalendarEventId     *Many2One  `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time      `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String    `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation  `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String    `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String    `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One  `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One  `xmlrpc:"activity_user_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	DuePeriodEnd                *Time      `xmlrpc:"due_period_end,omitempty"`
	DuePeriodStart              *Time      `xmlrpc:"due_period_start,omitempty"`
	ErrorMoveMessage            *String    `xmlrpc:"error_move_message,omitempty"`
	ErrorMovesCount             *Int       `xmlrpc:"error_moves_count,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	InitialFlowId               *Many2One  `xmlrpc:"initial_flow_id,omitempty"`
	MessageAttachmentCount      *Int       `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation  `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool      `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int       `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool      `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation  `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool      `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction           *Bool      `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int       `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation  `xmlrpc:"message_partner_ids,omitempty"`
	MoveIds                     *Relation  `xmlrpc:"move_ids,omitempty"`
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	OperationType               *Selection `xmlrpc:"operation_type,omitempty"`
	PayloadId                   *Many2One  `xmlrpc:"payload_id,omitempty"`
	PdpFlowId                   *String    `xmlrpc:"pdp_flow_id,omitempty"`
	PeriodEnd                   *Time      `xmlrpc:"period_end,omitempty"`
	PeriodStart                 *Time      `xmlrpc:"period_start,omitempty"`
	PeriodStatus                *Selection `xmlrpc:"period_status,omitempty"`
	PeriodicityCode             *String    `xmlrpc:"periodicity_code,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	RectificativeFlowIds        *Relation  `xmlrpc:"rectificative_flow_ids,omitempty"`
	ReportType                  *Selection `xmlrpc:"report_type,omitempty"`
	SentMoveIds                 *Relation  `xmlrpc:"sent_move_ids,omitempty"`
	State                       *Selection `xmlrpc:"state,omitempty"`
	TransmissionType            *Selection `xmlrpc:"transmission_type,omitempty"`
	TransportMessage            *String    `xmlrpc:"transport_message,omitempty"`
	TransportStatus             *String    `xmlrpc:"transport_status,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// L10NFrPdpReportsFlows represents array of l10n.fr.pdp.reports.flow model.
type L10NFrPdpReportsFlows []L10NFrPdpReportsFlow

// L10NFrPdpReportsFlowModel is the odoo model name.
const L10NFrPdpReportsFlowModel = "l10n.fr.pdp.reports.flow"

// Many2One convert L10NFrPdpReportsFlow to *Many2One.
func (lfprf *L10NFrPdpReportsFlow) Many2One() *Many2One {
	return NewMany2One(lfprf.Id.Get(), "")
}

// CreateL10NFrPdpReportsFlow creates a new l10n.fr.pdp.reports.flow model and returns its id.
func (c *Client) CreateL10NFrPdpReportsFlow(lfprf *L10NFrPdpReportsFlow) (int64, error) {
	ids, err := c.CreateL10NFrPdpReportsFlows([]*L10NFrPdpReportsFlow{lfprf})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateL10NFrPdpReportsFlow creates a new l10n.fr.pdp.reports.flow model and returns its id.
func (c *Client) CreateL10NFrPdpReportsFlows(lfprfs []*L10NFrPdpReportsFlow) ([]int64, error) {
	var vv []interface{}
	for _, v := range lfprfs {
		vv = append(vv, v)
	}
	return c.Create(L10NFrPdpReportsFlowModel, vv, nil)
}

// UpdateL10NFrPdpReportsFlow updates an existing l10n.fr.pdp.reports.flow record.
func (c *Client) UpdateL10NFrPdpReportsFlow(lfprf *L10NFrPdpReportsFlow) error {
	return c.UpdateL10NFrPdpReportsFlows([]int64{lfprf.Id.Get()}, lfprf)
}

// UpdateL10NFrPdpReportsFlows updates existing l10n.fr.pdp.reports.flow records.
// All records (represented by ids) will be updated by lfprf values.
func (c *Client) UpdateL10NFrPdpReportsFlows(ids []int64, lfprf *L10NFrPdpReportsFlow) error {
	return c.Update(L10NFrPdpReportsFlowModel, ids, lfprf, nil)
}

// DeleteL10NFrPdpReportsFlow deletes an existing l10n.fr.pdp.reports.flow record.
func (c *Client) DeleteL10NFrPdpReportsFlow(id int64) error {
	return c.DeleteL10NFrPdpReportsFlows([]int64{id})
}

// DeleteL10NFrPdpReportsFlows deletes existing l10n.fr.pdp.reports.flow records.
func (c *Client) DeleteL10NFrPdpReportsFlows(ids []int64) error {
	return c.Delete(L10NFrPdpReportsFlowModel, ids)
}

// GetL10NFrPdpReportsFlow gets l10n.fr.pdp.reports.flow existing record.
func (c *Client) GetL10NFrPdpReportsFlow(id int64) (*L10NFrPdpReportsFlow, error) {
	lfprfs, err := c.GetL10NFrPdpReportsFlows([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*lfprfs)[0]), nil
}

// GetL10NFrPdpReportsFlows gets l10n.fr.pdp.reports.flow existing records.
func (c *Client) GetL10NFrPdpReportsFlows(ids []int64) (*L10NFrPdpReportsFlows, error) {
	lfprfs := &L10NFrPdpReportsFlows{}
	if err := c.Read(L10NFrPdpReportsFlowModel, ids, nil, lfprfs); err != nil {
		return nil, err
	}
	return lfprfs, nil
}

// FindL10NFrPdpReportsFlow finds l10n.fr.pdp.reports.flow record by querying it with criteria.
func (c *Client) FindL10NFrPdpReportsFlow(criteria *Criteria) (*L10NFrPdpReportsFlow, error) {
	lfprfs := &L10NFrPdpReportsFlows{}
	if err := c.SearchRead(L10NFrPdpReportsFlowModel, criteria, NewOptions().Limit(1), lfprfs); err != nil {
		return nil, err
	}
	return &((*lfprfs)[0]), nil
}

// FindL10NFrPdpReportsFlows finds l10n.fr.pdp.reports.flow records by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrPdpReportsFlows(criteria *Criteria, options *Options) (*L10NFrPdpReportsFlows, error) {
	lfprfs := &L10NFrPdpReportsFlows{}
	if err := c.SearchRead(L10NFrPdpReportsFlowModel, criteria, options, lfprfs); err != nil {
		return nil, err
	}
	return lfprfs, nil
}

// FindL10NFrPdpReportsFlowIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindL10NFrPdpReportsFlowIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(L10NFrPdpReportsFlowModel, criteria, options)
}

// FindL10NFrPdpReportsFlowId finds record id by querying it with criteria.
func (c *Client) FindL10NFrPdpReportsFlowId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(L10NFrPdpReportsFlowModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
