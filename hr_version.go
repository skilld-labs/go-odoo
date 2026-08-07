package odoo

// HrVersion represents hr.version model.
type HrVersion struct {
	Active                      *Bool      `xmlrpc:"active,omitempty"`
	ActiveEmployee              *Bool      `xmlrpc:"active_employee,omitempty"`
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
	AdditionalNote              *String    `xmlrpc:"additional_note,omitempty"`
	AddressId                   *Many2One  `xmlrpc:"address_id,omitempty"`
	AllowedCountryStateIds      *Relation  `xmlrpc:"allowed_country_state_ids,omitempty"`
	Children                    *Int       `xmlrpc:"children,omitempty"`
	CompanyCountryId            *Many2One  `xmlrpc:"company_country_id,omitempty"`
	CompanyId                   *Many2One  `xmlrpc:"company_id,omitempty"`
	ContractDateEnd             *Time      `xmlrpc:"contract_date_end,omitempty"`
	ContractDateStart           *Time      `xmlrpc:"contract_date_start,omitempty"`
	ContractTemplateId          *Many2One  `xmlrpc:"contract_template_id,omitempty"`
	ContractTypeId              *Many2One  `xmlrpc:"contract_type_id,omitempty"`
	ContractWage                *Float     `xmlrpc:"contract_wage,omitempty"`
	CountryCode                 *String    `xmlrpc:"country_code,omitempty"`
	CountryId                   *Many2One  `xmlrpc:"country_id,omitempty"`
	CreateDate                  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One  `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One  `xmlrpc:"currency_id,omitempty"`
	DateEnd                     *Time      `xmlrpc:"date_end,omitempty"`
	DateStart                   *Time      `xmlrpc:"date_start,omitempty"`
	DateVersion                 *Time      `xmlrpc:"date_version,omitempty"`
	DepartmentId                *Many2One  `xmlrpc:"department_id,omitempty"`
	DepartureDate               *Time      `xmlrpc:"departure_date,omitempty"`
	DepartureDescription        *String    `xmlrpc:"departure_description,omitempty"`
	DepartureReasonId           *Many2One  `xmlrpc:"departure_reason_id,omitempty"`
	DisplayName                 *String    `xmlrpc:"display_name,omitempty"`
	DistanceHomeWork            *Int       `xmlrpc:"distance_home_work,omitempty"`
	DistanceHomeWorkUnit        *Selection `xmlrpc:"distance_home_work_unit,omitempty"`
	EmployeeId                  *Many2One  `xmlrpc:"employee_id,omitempty"`
	EmployeeType                *Selection `xmlrpc:"employee_type,omitempty"`
	HasMessage                  *Bool      `xmlrpc:"has_message,omitempty"`
	HrResponsibleId             *Many2One  `xmlrpc:"hr_responsible_id,omitempty"`
	Id                          *Int       `xmlrpc:"id,omitempty"`
	IdentificationId            *String    `xmlrpc:"identification_id,omitempty"`
	IsCurrent                   *Bool      `xmlrpc:"is_current,omitempty"`
	IsCustomJobTitle            *Bool      `xmlrpc:"is_custom_job_title,omitempty"`
	IsFlexible                  *Bool      `xmlrpc:"is_flexible,omitempty"`
	IsFullyFlexible             *Bool      `xmlrpc:"is_fully_flexible,omitempty"`
	IsFuture                    *Bool      `xmlrpc:"is_future,omitempty"`
	IsInContract                *Bool      `xmlrpc:"is_in_contract,omitempty"`
	IsPast                      *Bool      `xmlrpc:"is_past,omitempty"`
	JobId                       *Many2One  `xmlrpc:"job_id,omitempty"`
	JobTitle                    *String    `xmlrpc:"job_title,omitempty"`
	KmHomeWork                  *Int       `xmlrpc:"km_home_work,omitempty"`
	LastModifiedDate            *Time      `xmlrpc:"last_modified_date,omitempty"`
	LastModifiedUid             *Many2One  `xmlrpc:"last_modified_uid,omitempty"`
	Marital                     *Selection `xmlrpc:"marital,omitempty"`
	MemberOfDepartment          *Bool      `xmlrpc:"member_of_department,omitempty"`
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
	MyActivityDateDeadline      *Time      `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String    `xmlrpc:"name,omitempty"`
	PassportExpirationDate      *Time      `xmlrpc:"passport_expiration_date,omitempty"`
	PassportId                  *String    `xmlrpc:"passport_id,omitempty"`
	PrivateCity                 *String    `xmlrpc:"private_city,omitempty"`
	PrivateCountryId            *Many2One  `xmlrpc:"private_country_id,omitempty"`
	PrivateStateId              *Many2One  `xmlrpc:"private_state_id,omitempty"`
	PrivateStreet               *String    `xmlrpc:"private_street,omitempty"`
	PrivateStreet2              *String    `xmlrpc:"private_street2,omitempty"`
	PrivateZip                  *String    `xmlrpc:"private_zip,omitempty"`
	RatingIds                   *Relation  `xmlrpc:"rating_ids,omitempty"`
	ResourceCalendarId          *Many2One  `xmlrpc:"resource_calendar_id,omitempty"`
	Sex                         *Selection `xmlrpc:"sex,omitempty"`
	SpouseBirthdate             *Time      `xmlrpc:"spouse_birthdate,omitempty"`
	SpouseCompleteName          *String    `xmlrpc:"spouse_complete_name,omitempty"`
	Ssnid                       *String    `xmlrpc:"ssnid,omitempty"`
	StructureTypeId             *Many2One  `xmlrpc:"structure_type_id,omitempty"`
	TrialDateEnd                *Time      `xmlrpc:"trial_date_end,omitempty"`
	Tz                          *Selection `xmlrpc:"tz,omitempty"`
	Wage                        *Float     `xmlrpc:"wage,omitempty"`
	WebsiteMessageIds           *Relation  `xmlrpc:"website_message_ids,omitempty"`
	WorkLocationId              *Many2One  `xmlrpc:"work_location_id,omitempty"`
	WriteDate                   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// HrVersions represents array of hr.version model.
type HrVersions []HrVersion

// HrVersionModel is the odoo model name.
const HrVersionModel = "hr.version"

// Many2One convert HrVersion to *Many2One.
func (hv *HrVersion) Many2One() *Many2One {
	return NewMany2One(hv.Id.Get(), "")
}

// CreateHrVersion creates a new hr.version model and returns its id.
func (c *Client) CreateHrVersion(hv *HrVersion) (int64, error) {
	ids, err := c.CreateHrVersions([]*HrVersion{hv})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrVersion creates a new hr.version model and returns its id.
func (c *Client) CreateHrVersions(hvs []*HrVersion) ([]int64, error) {
	var vv []interface{}
	for _, v := range hvs {
		vv = append(vv, v)
	}
	return c.Create(HrVersionModel, vv, nil)
}

// UpdateHrVersion updates an existing hr.version record.
func (c *Client) UpdateHrVersion(hv *HrVersion) error {
	return c.UpdateHrVersions([]int64{hv.Id.Get()}, hv)
}

// UpdateHrVersions updates existing hr.version records.
// All records (represented by ids) will be updated by hv values.
func (c *Client) UpdateHrVersions(ids []int64, hv *HrVersion) error {
	return c.Update(HrVersionModel, ids, hv, nil)
}

// DeleteHrVersion deletes an existing hr.version record.
func (c *Client) DeleteHrVersion(id int64) error {
	return c.DeleteHrVersions([]int64{id})
}

// DeleteHrVersions deletes existing hr.version records.
func (c *Client) DeleteHrVersions(ids []int64) error {
	return c.Delete(HrVersionModel, ids)
}

// GetHrVersion gets hr.version existing record.
func (c *Client) GetHrVersion(id int64) (*HrVersion, error) {
	hvs, err := c.GetHrVersions([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hvs)[0]), nil
}

// GetHrVersions gets hr.version existing records.
func (c *Client) GetHrVersions(ids []int64) (*HrVersions, error) {
	hvs := &HrVersions{}
	if err := c.Read(HrVersionModel, ids, nil, hvs); err != nil {
		return nil, err
	}
	return hvs, nil
}

// FindHrVersion finds hr.version record by querying it with criteria.
func (c *Client) FindHrVersion(criteria *Criteria) (*HrVersion, error) {
	hvs := &HrVersions{}
	if err := c.SearchRead(HrVersionModel, criteria, NewOptions().Limit(1), hvs); err != nil {
		return nil, err
	}
	return &((*hvs)[0]), nil
}

// FindHrVersions finds hr.version records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrVersions(criteria *Criteria, options *Options) (*HrVersions, error) {
	hvs := &HrVersions{}
	if err := c.SearchRead(HrVersionModel, criteria, options, hvs); err != nil {
		return nil, err
	}
	return hvs, nil
}

// FindHrVersionIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrVersionIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrVersionModel, criteria, options)
}

// FindHrVersionId finds record id by querying it with criteria.
func (c *Client) FindHrVersionId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrVersionModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
