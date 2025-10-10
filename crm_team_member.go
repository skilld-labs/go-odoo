package odoo

// CrmTeamMember represents crm.team.member model.
type CrmTeamMember struct {
	Active                   *Bool     `xmlrpc:"active,omitempty"`
	AssignmentDomain         *String   `xmlrpc:"assignment_domain,omitempty"`
	AssignmentEnabled        *Bool     `xmlrpc:"assignment_enabled,omitempty"`
	AssignmentMax            *Int      `xmlrpc:"assignment_max,omitempty"`
	AssignmentOptout         *Bool     `xmlrpc:"assignment_optout,omitempty"`
	CompanyId                *Many2One `xmlrpc:"company_id,omitempty"`
	CreateDate               *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid                *Many2One `xmlrpc:"create_uid,omitempty"`
	CrmTeamId                *Many2One `xmlrpc:"crm_team_id,omitempty"`
	DisplayName              *String   `xmlrpc:"display_name,omitempty"`
	Email                    *String   `xmlrpc:"email,omitempty"`
	HasMessage               *Bool     `xmlrpc:"has_message,omitempty"`
	Id                       *Int      `xmlrpc:"id,omitempty"`
	Image128                 *String   `xmlrpc:"image_128,omitempty"`
	Image1920                *String   `xmlrpc:"image_1920,omitempty"`
	IsMembershipMulti        *Bool     `xmlrpc:"is_membership_multi,omitempty"`
	LeadDayCount             *Int      `xmlrpc:"lead_day_count,omitempty"`
	LeadMonthCount           *Int      `xmlrpc:"lead_month_count,omitempty"`
	MemberWarning            *String   `xmlrpc:"member_warning,omitempty"`
	MessageAttachmentCount   *Int      `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds       *Relation `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError          *Bool     `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter   *Int      `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError       *Bool     `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds               *Relation `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower        *Bool     `xmlrpc:"message_is_follower,omitempty"`
	MessageNeedaction        *Bool     `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter *Int      `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds        *Relation `xmlrpc:"message_partner_ids,omitempty"`
	Mobile                   *String   `xmlrpc:"mobile,omitempty"`
	Name                     *String   `xmlrpc:"name,omitempty"`
	Phone                    *String   `xmlrpc:"phone,omitempty"`
	RatingIds                *Relation `xmlrpc:"rating_ids,omitempty"`
	UserCompanyIds           *Relation `xmlrpc:"user_company_ids,omitempty"`
	UserId                   *Many2One `xmlrpc:"user_id,omitempty"`
	UserInTeamsIds           *Relation `xmlrpc:"user_in_teams_ids,omitempty"`
	WebsiteMessageIds        *Relation `xmlrpc:"website_message_ids,omitempty"`
	WriteDate                *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid                 *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CrmTeamMembers represents array of crm.team.member model.
type CrmTeamMembers []CrmTeamMember

// CrmTeamMemberModel is the odoo model name.
const CrmTeamMemberModel = "crm.team.member"

// Many2One convert CrmTeamMember to *Many2One.
func (ctm *CrmTeamMember) Many2One() *Many2One {
	return NewMany2One(ctm.Id.Get(), "")
}

// CreateCrmTeamMember creates a new crm.team.member model and returns its id.
func (c *Client) CreateCrmTeamMember(ctm *CrmTeamMember) (int64, error) {
	ids, err := c.CreateCrmTeamMembers([]*CrmTeamMember{ctm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCrmTeamMember creates a new crm.team.member model and returns its id.
func (c *Client) CreateCrmTeamMembers(ctms []*CrmTeamMember) ([]int64, error) {
	var vv []interface{}
	for _, v := range ctms {
		vv = append(vv, v)
	}
	return c.Create(CrmTeamMemberModel, vv, nil)
}

// UpdateCrmTeamMember updates an existing crm.team.member record.
func (c *Client) UpdateCrmTeamMember(ctm *CrmTeamMember) error {
	return c.UpdateCrmTeamMembers([]int64{ctm.Id.Get()}, ctm)
}

// UpdateCrmTeamMembers updates existing crm.team.member records.
// All records (represented by ids) will be updated by ctm values.
func (c *Client) UpdateCrmTeamMembers(ids []int64, ctm *CrmTeamMember) error {
	return c.Update(CrmTeamMemberModel, ids, ctm, nil)
}

// DeleteCrmTeamMember deletes an existing crm.team.member record.
func (c *Client) DeleteCrmTeamMember(id int64) error {
	return c.DeleteCrmTeamMembers([]int64{id})
}

// DeleteCrmTeamMembers deletes existing crm.team.member records.
func (c *Client) DeleteCrmTeamMembers(ids []int64) error {
	return c.Delete(CrmTeamMemberModel, ids)
}

// GetCrmTeamMember gets crm.team.member existing record.
func (c *Client) GetCrmTeamMember(id int64) (*CrmTeamMember, error) {
	ctms, err := c.GetCrmTeamMembers([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ctms)[0]), nil
}

// GetCrmTeamMembers gets crm.team.member existing records.
func (c *Client) GetCrmTeamMembers(ids []int64) (*CrmTeamMembers, error) {
	ctms := &CrmTeamMembers{}
	if err := c.Read(CrmTeamMemberModel, ids, nil, ctms); err != nil {
		return nil, err
	}
	return ctms, nil
}

// FindCrmTeamMember finds crm.team.member record by querying it with criteria.
func (c *Client) FindCrmTeamMember(criteria *Criteria) (*CrmTeamMember, error) {
	ctms := &CrmTeamMembers{}
	if err := c.SearchRead(CrmTeamMemberModel, criteria, NewOptions().Limit(1), ctms); err != nil {
		return nil, err
	}
	return &((*ctms)[0]), nil
}

// FindCrmTeamMembers finds crm.team.member records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmTeamMembers(criteria *Criteria, options *Options) (*CrmTeamMembers, error) {
	ctms := &CrmTeamMembers{}
	if err := c.SearchRead(CrmTeamMemberModel, criteria, options, ctms); err != nil {
		return nil, err
	}
	return ctms, nil
}

// FindCrmTeamMemberIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCrmTeamMemberIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CrmTeamMemberModel, criteria, options)
}

// FindCrmTeamMemberId finds record id by querying it with criteria.
func (c *Client) FindCrmTeamMemberId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CrmTeamMemberModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
