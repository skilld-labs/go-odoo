package odoo

// HrEmployee represents hr.employee model.
type HrEmployee struct {
	Active                      *Bool       `xmlrpc:"active,omitempty"`
	ActiveEmployee              *Bool       `xmlrpc:"active_employee,omitempty"`
	ActivityCalendarEventId     *Many2One   `xmlrpc:"activity_calendar_event_id,omitempty"`
	ActivityDateDeadline        *Time       `xmlrpc:"activity_date_deadline,omitempty"`
	ActivityExceptionDecoration *Selection  `xmlrpc:"activity_exception_decoration,omitempty"`
	ActivityExceptionIcon       *String     `xmlrpc:"activity_exception_icon,omitempty"`
	ActivityIds                 *Relation   `xmlrpc:"activity_ids,omitempty"`
	ActivityState               *Selection  `xmlrpc:"activity_state,omitempty"`
	ActivitySummary             *String     `xmlrpc:"activity_summary,omitempty"`
	ActivityTypeIcon            *String     `xmlrpc:"activity_type_icon,omitempty"`
	ActivityTypeId              *Many2One   `xmlrpc:"activity_type_id,omitempty"`
	ActivityUserId              *Many2One   `xmlrpc:"activity_user_id,omitempty"`
	AdditionalNote              *String     `xmlrpc:"additional_note,omitempty"`
	AddressId                   *Many2One   `xmlrpc:"address_id,omitempty"`
	AllocationCount             *Float      `xmlrpc:"allocation_count,omitempty"`
	AllocationDisplay           *String     `xmlrpc:"allocation_display,omitempty"`
	AllocationRemainingDisplay  *String     `xmlrpc:"allocation_remaining_display,omitempty"`
	AllocationsCount            *Int        `xmlrpc:"allocations_count,omitempty"`
	AllowedCountryStateIds      *Relation   `xmlrpc:"allowed_country_state_ids,omitempty"`
	Avatar1024                  *String     `xmlrpc:"avatar_1024,omitempty"`
	Avatar128                   *String     `xmlrpc:"avatar_128,omitempty"`
	Avatar1920                  *String     `xmlrpc:"avatar_1920,omitempty"`
	Avatar256                   *String     `xmlrpc:"avatar_256,omitempty"`
	Avatar512                   *String     `xmlrpc:"avatar_512,omitempty"`
	BankAccountIds              *Relation   `xmlrpc:"bank_account_ids,omitempty"`
	Barcode                     *String     `xmlrpc:"barcode,omitempty"`
	Birthday                    *Time       `xmlrpc:"birthday,omitempty"`
	BirthdayPublicDisplay       *Bool       `xmlrpc:"birthday_public_display,omitempty"`
	BirthdayPublicDisplayString *String     `xmlrpc:"birthday_public_display_string,omitempty"`
	CategoryIds                 *Relation   `xmlrpc:"category_ids,omitempty"`
	Certificate                 *Selection  `xmlrpc:"certificate,omitempty"`
	CertificationIds            *Relation   `xmlrpc:"certification_ids,omitempty"`
	ChildAllCount               *Int        `xmlrpc:"child_all_count,omitempty"`
	ChildCount                  *Int        `xmlrpc:"child_count,omitempty"`
	ChildIds                    *Relation   `xmlrpc:"child_ids,omitempty"`
	Children                    *Int        `xmlrpc:"children,omitempty"`
	CoachId                     *Many2One   `xmlrpc:"coach_id,omitempty"`
	Color                       *Int        `xmlrpc:"color,omitempty"`
	CompanyCountryCode          *String     `xmlrpc:"company_country_code,omitempty"`
	CompanyCountryId            *Many2One   `xmlrpc:"company_country_id,omitempty"`
	CompanyId                   *Many2One   `xmlrpc:"company_id,omitempty"`
	ContractDateEnd             *Time       `xmlrpc:"contract_date_end,omitempty"`
	ContractDateStart           *Time       `xmlrpc:"contract_date_start,omitempty"`
	ContractTemplateId          *Many2One   `xmlrpc:"contract_template_id,omitempty"`
	ContractTypeId              *Many2One   `xmlrpc:"contract_type_id,omitempty"`
	ContractWage                *Float      `xmlrpc:"contract_wage,omitempty"`
	CountryCode                 *String     `xmlrpc:"country_code,omitempty"`
	CountryId                   *Many2One   `xmlrpc:"country_id,omitempty"`
	CountryOfBirth              *Many2One   `xmlrpc:"country_of_birth,omitempty"`
	CreateDate                  *Time       `xmlrpc:"create_date,omitempty"`
	CreateUid                   *Many2One   `xmlrpc:"create_uid,omitempty"`
	CurrencyId                  *Many2One   `xmlrpc:"currency_id,omitempty"`
	CurrentDateVersion          *Time       `xmlrpc:"current_date_version,omitempty"`
	CurrentEmployeeSkillIds     *Relation   `xmlrpc:"current_employee_skill_ids,omitempty"`
	CurrentLeaveId              *Many2One   `xmlrpc:"current_leave_id,omitempty"`
	CurrentLeaveState           *Selection  `xmlrpc:"current_leave_state,omitempty"`
	CurrentVersionId            *Many2One   `xmlrpc:"current_version_id,omitempty"`
	DateEnd                     *Time       `xmlrpc:"date_end,omitempty"`
	DateStart                   *Time       `xmlrpc:"date_start,omitempty"`
	DateVersion                 *Time       `xmlrpc:"date_version,omitempty"`
	DepartmentColor             *Int        `xmlrpc:"department_color,omitempty"`
	DepartmentId                *Many2One   `xmlrpc:"department_id,omitempty"`
	DepartureDate               *Time       `xmlrpc:"departure_date,omitempty"`
	DepartureDescription        *String     `xmlrpc:"departure_description,omitempty"`
	DepartureReasonId           *Many2One   `xmlrpc:"departure_reason_id,omitempty"`
	DisplayCertificationPage    *Bool       `xmlrpc:"display_certification_page,omitempty"`
	DisplayName                 *String     `xmlrpc:"display_name,omitempty"`
	DistanceHomeWork            *Int        `xmlrpc:"distance_home_work,omitempty"`
	DistanceHomeWorkUnit        *Selection  `xmlrpc:"distance_home_work_unit,omitempty"`
	DrivingLicense              *String     `xmlrpc:"driving_license,omitempty"`
	Email                       *String     `xmlrpc:"email,omitempty"`
	EmergencyContact            *String     `xmlrpc:"emergency_contact,omitempty"`
	EmergencyPhone              *String     `xmlrpc:"emergency_phone,omitempty"`
	EmployeeId                  *Many2One   `xmlrpc:"employee_id,omitempty"`
	EmployeeProperties          interface{} `xmlrpc:"employee_properties,omitempty"`
	EmployeeSkillIds            *Relation   `xmlrpc:"employee_skill_ids,omitempty"`
	EmployeeType                *Selection  `xmlrpc:"employee_type,omitempty"`
	ExceptionalLocationId       *Many2One   `xmlrpc:"exceptional_location_id,omitempty"`
	FridayLocationId            *Many2One   `xmlrpc:"friday_location_id,omitempty"`
	HasMessage                  *Bool       `xmlrpc:"has_message,omitempty"`
	HasMultipleBankAccounts     *Bool       `xmlrpc:"has_multiple_bank_accounts,omitempty"`
	HasTimesheet                *Bool       `xmlrpc:"has_timesheet,omitempty"`
	HasWorkPermit               *String     `xmlrpc:"has_work_permit,omitempty"`
	HourlyCost                  *Float      `xmlrpc:"hourly_cost,omitempty"`
	HrIconDisplay               *Selection  `xmlrpc:"hr_icon_display,omitempty"`
	HrPresenceState             *Selection  `xmlrpc:"hr_presence_state,omitempty"`
	HrResponsibleId             *Many2One   `xmlrpc:"hr_responsible_id,omitempty"`
	Id                          *Int        `xmlrpc:"id,omitempty"`
	IdCard                      *String     `xmlrpc:"id_card,omitempty"`
	IdentificationId            *String     `xmlrpc:"identification_id,omitempty"`
	ImStatus                    *String     `xmlrpc:"im_status,omitempty"`
	Image1024                   *String     `xmlrpc:"image_1024,omitempty"`
	Image128                    *String     `xmlrpc:"image_128,omitempty"`
	Image1920                   *String     `xmlrpc:"image_1920,omitempty"`
	Image256                    *String     `xmlrpc:"image_256,omitempty"`
	Image512                    *String     `xmlrpc:"image_512,omitempty"`
	IsAbsent                    *Bool       `xmlrpc:"is_absent,omitempty"`
	IsCurrent                   *Bool       `xmlrpc:"is_current,omitempty"`
	IsCustomJobTitle            *Bool       `xmlrpc:"is_custom_job_title,omitempty"`
	IsFlexible                  *Bool       `xmlrpc:"is_flexible,omitempty"`
	IsFullyFlexible             *Bool       `xmlrpc:"is_fully_flexible,omitempty"`
	IsFuture                    *Bool       `xmlrpc:"is_future,omitempty"`
	IsInContract                *Bool       `xmlrpc:"is_in_contract,omitempty"`
	IsPast                      *Bool       `xmlrpc:"is_past,omitempty"`
	IsSubordinate               *Bool       `xmlrpc:"is_subordinate,omitempty"`
	IsTrustedBankAccount        *Bool       `xmlrpc:"is_trusted_bank_account,omitempty"`
	IsUserActive                *Bool       `xmlrpc:"is_user_active,omitempty"`
	JobId                       *Many2One   `xmlrpc:"job_id,omitempty"`
	JobTitle                    *String     `xmlrpc:"job_title,omitempty"`
	KmHomeWork                  *Int        `xmlrpc:"km_home_work,omitempty"`
	Lang                        *Selection  `xmlrpc:"lang,omitempty"`
	LastActivity                *Time       `xmlrpc:"last_activity,omitempty"`
	LastActivityTime            *String     `xmlrpc:"last_activity_time,omitempty"`
	LastModifiedDate            *Time       `xmlrpc:"last_modified_date,omitempty"`
	LastModifiedUid             *Many2One   `xmlrpc:"last_modified_uid,omitempty"`
	LeaveDateFrom               *Time       `xmlrpc:"leave_date_from,omitempty"`
	LeaveDateTo                 *Time       `xmlrpc:"leave_date_to,omitempty"`
	LeaveManagerId              *Many2One   `xmlrpc:"leave_manager_id,omitempty"`
	LegalName                   *String     `xmlrpc:"legal_name,omitempty"`
	Marital                     *Selection  `xmlrpc:"marital,omitempty"`
	MemberOfDepartment          *Bool       `xmlrpc:"member_of_department,omitempty"`
	MessageAttachmentCount      *Int        `xmlrpc:"message_attachment_count,omitempty"`
	MessageFollowerIds          *Relation   `xmlrpc:"message_follower_ids,omitempty"`
	MessageHasError             *Bool       `xmlrpc:"message_has_error,omitempty"`
	MessageHasErrorCounter      *Int        `xmlrpc:"message_has_error_counter,omitempty"`
	MessageHasSmsError          *Bool       `xmlrpc:"message_has_sms_error,omitempty"`
	MessageIds                  *Relation   `xmlrpc:"message_ids,omitempty"`
	MessageIsFollower           *Bool       `xmlrpc:"message_is_follower,omitempty"`
	MessageMainAttachmentId     *Many2One   `xmlrpc:"message_main_attachment_id,omitempty"`
	MessageNeedaction           *Bool       `xmlrpc:"message_needaction,omitempty"`
	MessageNeedactionCounter    *Int        `xmlrpc:"message_needaction_counter,omitempty"`
	MessagePartnerIds           *Relation   `xmlrpc:"message_partner_ids,omitempty"`
	MobilePhone                 *String     `xmlrpc:"mobile_phone,omitempty"`
	MondayLocationId            *Many2One   `xmlrpc:"monday_location_id,omitempty"`
	MyActivityDateDeadline      *Time       `xmlrpc:"my_activity_date_deadline,omitempty"`
	Name                        *String     `xmlrpc:"name,omitempty"`
	NewlyHired                  *Bool       `xmlrpc:"newly_hired,omitempty"`
	ParentId                    *Many2One   `xmlrpc:"parent_id,omitempty"`
	PassportExpirationDate      *Time       `xmlrpc:"passport_expiration_date,omitempty"`
	PassportId                  *String     `xmlrpc:"passport_id,omitempty"`
	PermitNo                    *String     `xmlrpc:"permit_no,omitempty"`
	Phone                       *String     `xmlrpc:"phone,omitempty"`
	Pin                         *String     `xmlrpc:"pin,omitempty"`
	PlaceOfBirth                *String     `xmlrpc:"place_of_birth,omitempty"`
	PrimaryBankAccountId        *Many2One   `xmlrpc:"primary_bank_account_id,omitempty"`
	PrivateCarPlate             *String     `xmlrpc:"private_car_plate,omitempty"`
	PrivateCity                 *String     `xmlrpc:"private_city,omitempty"`
	PrivateCountryId            *Many2One   `xmlrpc:"private_country_id,omitempty"`
	PrivateEmail                *String     `xmlrpc:"private_email,omitempty"`
	PrivatePhone                *String     `xmlrpc:"private_phone,omitempty"`
	PrivateStateId              *Many2One   `xmlrpc:"private_state_id,omitempty"`
	PrivateStreet               *String     `xmlrpc:"private_street,omitempty"`
	PrivateStreet2              *String     `xmlrpc:"private_street2,omitempty"`
	PrivateZip                  *String     `xmlrpc:"private_zip,omitempty"`
	RatingIds                   *Relation   `xmlrpc:"rating_ids,omitempty"`
	RelatedPartnersCount        *Int        `xmlrpc:"related_partners_count,omitempty"`
	ResourceCalendarId          *Many2One   `xmlrpc:"resource_calendar_id,omitempty"`
	ResourceId                  *Many2One   `xmlrpc:"resource_id,omitempty"`
	ResumeLineIds               *Relation   `xmlrpc:"resume_line_ids,omitempty"`
	SalaryDistribution          interface{} `xmlrpc:"salary_distribution,omitempty"`
	SaturdayLocationId          *Many2One   `xmlrpc:"saturday_location_id,omitempty"`
	Sex                         *Selection  `xmlrpc:"sex,omitempty"`
	Share                       *Bool       `xmlrpc:"share,omitempty"`
	ShowHrIconDisplay           *Bool       `xmlrpc:"show_hr_icon_display,omitempty"`
	ShowLeaves                  *Bool       `xmlrpc:"show_leaves,omitempty"`
	SkillIds                    *Relation   `xmlrpc:"skill_ids,omitempty"`
	SpouseBirthdate             *Time       `xmlrpc:"spouse_birthdate,omitempty"`
	SpouseCompleteName          *String     `xmlrpc:"spouse_complete_name,omitempty"`
	Ssnid                       *String     `xmlrpc:"ssnid,omitempty"`
	StructureTypeId             *Many2One   `xmlrpc:"structure_type_id,omitempty"`
	StudyField                  *String     `xmlrpc:"study_field,omitempty"`
	StudySchool                 *String     `xmlrpc:"study_school,omitempty"`
	SubordinateIds              *Relation   `xmlrpc:"subordinate_ids,omitempty"`
	SundayLocationId            *Many2One   `xmlrpc:"sunday_location_id,omitempty"`
	ThursdayLocationId          *Many2One   `xmlrpc:"thursday_location_id,omitempty"`
	TodayLocationName           *String     `xmlrpc:"today_location_name,omitempty"`
	TrialDateEnd                *Time       `xmlrpc:"trial_date_end,omitempty"`
	TuesdayLocationId           *Many2One   `xmlrpc:"tuesday_location_id,omitempty"`
	Tz                          *Selection  `xmlrpc:"tz,omitempty"`
	UserId                      *Many2One   `xmlrpc:"user_id,omitempty"`
	UserPartnerId               *Many2One   `xmlrpc:"user_partner_id,omitempty"`
	VersionId                   *Many2One   `xmlrpc:"version_id,omitempty"`
	VersionIds                  *Relation   `xmlrpc:"version_ids,omitempty"`
	VersionRevision             *String     `xmlrpc:"version_revision,omitempty"`
	VersionsCount               *Int        `xmlrpc:"versions_count,omitempty"`
	VisaExpire                  *Time       `xmlrpc:"visa_expire,omitempty"`
	VisaNo                      *String     `xmlrpc:"visa_no,omitempty"`
	Wage                        *Float      `xmlrpc:"wage,omitempty"`
	WebsiteMessageIds           *Relation   `xmlrpc:"website_message_ids,omitempty"`
	WednesdayLocationId         *Many2One   `xmlrpc:"wednesday_location_id,omitempty"`
	WorkContactId               *Many2One   `xmlrpc:"work_contact_id,omitempty"`
	WorkEmail                   *String     `xmlrpc:"work_email,omitempty"`
	WorkLocationId              *Many2One   `xmlrpc:"work_location_id,omitempty"`
	WorkLocationName            *String     `xmlrpc:"work_location_name,omitempty"`
	WorkLocationType            *Selection  `xmlrpc:"work_location_type,omitempty"`
	WorkPermitExpirationDate    *Time       `xmlrpc:"work_permit_expiration_date,omitempty"`
	WorkPermitName              *String     `xmlrpc:"work_permit_name,omitempty"`
	WorkPermitScheduledActivity *Bool       `xmlrpc:"work_permit_scheduled_activity,omitempty"`
	WorkPhone                   *String     `xmlrpc:"work_phone,omitempty"`
	WriteDate                   *Time       `xmlrpc:"write_date,omitempty"`
	WriteUid                    *Many2One   `xmlrpc:"write_uid,omitempty"`
}

// HrEmployees represents array of hr.employee model.
type HrEmployees []HrEmployee

// HrEmployeeModel is the odoo model name.
const HrEmployeeModel = "hr.employee"

// Many2One convert HrEmployee to *Many2One.
func (he *HrEmployee) Many2One() *Many2One {
	return NewMany2One(he.Id.Get(), "")
}

// CreateHrEmployee creates a new hr.employee model and returns its id.
func (c *Client) CreateHrEmployee(he *HrEmployee) (int64, error) {
	ids, err := c.CreateHrEmployees([]*HrEmployee{he})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateHrEmployee creates a new hr.employee model and returns its id.
func (c *Client) CreateHrEmployees(hes []*HrEmployee) ([]int64, error) {
	var vv []interface{}
	for _, v := range hes {
		vv = append(vv, v)
	}
	return c.Create(HrEmployeeModel, vv, nil)
}

// UpdateHrEmployee updates an existing hr.employee record.
func (c *Client) UpdateHrEmployee(he *HrEmployee) error {
	return c.UpdateHrEmployees([]int64{he.Id.Get()}, he)
}

// UpdateHrEmployees updates existing hr.employee records.
// All records (represented by ids) will be updated by he values.
func (c *Client) UpdateHrEmployees(ids []int64, he *HrEmployee) error {
	return c.Update(HrEmployeeModel, ids, he, nil)
}

// DeleteHrEmployee deletes an existing hr.employee record.
func (c *Client) DeleteHrEmployee(id int64) error {
	return c.DeleteHrEmployees([]int64{id})
}

// DeleteHrEmployees deletes existing hr.employee records.
func (c *Client) DeleteHrEmployees(ids []int64) error {
	return c.Delete(HrEmployeeModel, ids)
}

// GetHrEmployee gets hr.employee existing record.
func (c *Client) GetHrEmployee(id int64) (*HrEmployee, error) {
	hes, err := c.GetHrEmployees([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*hes)[0]), nil
}

// GetHrEmployees gets hr.employee existing records.
func (c *Client) GetHrEmployees(ids []int64) (*HrEmployees, error) {
	hes := &HrEmployees{}
	if err := c.Read(HrEmployeeModel, ids, nil, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrEmployee finds hr.employee record by querying it with criteria.
func (c *Client) FindHrEmployee(criteria *Criteria) (*HrEmployee, error) {
	hes := &HrEmployees{}
	if err := c.SearchRead(HrEmployeeModel, criteria, NewOptions().Limit(1), hes); err != nil {
		return nil, err
	}
	return &((*hes)[0]), nil
}

// FindHrEmployees finds hr.employee records by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployees(criteria *Criteria, options *Options) (*HrEmployees, error) {
	hes := &HrEmployees{}
	if err := c.SearchRead(HrEmployeeModel, criteria, options, hes); err != nil {
		return nil, err
	}
	return hes, nil
}

// FindHrEmployeeIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindHrEmployeeIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(HrEmployeeModel, criteria, options)
}

// FindHrEmployeeId finds record id by querying it with criteria.
func (c *Client) FindHrEmployeeId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(HrEmployeeModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
