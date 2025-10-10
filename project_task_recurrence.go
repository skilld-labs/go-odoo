package odoo

// ProjectTaskRecurrence represents project.task.recurrence model.
type ProjectTaskRecurrence struct {
	CreateDate     *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid      *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName    *String    `xmlrpc:"display_name,omitempty"`
	Id             *Int       `xmlrpc:"id,omitempty"`
	RepeatInterval *Int       `xmlrpc:"repeat_interval,omitempty"`
	RepeatType     *Selection `xmlrpc:"repeat_type,omitempty"`
	RepeatUnit     *Selection `xmlrpc:"repeat_unit,omitempty"`
	RepeatUntil    *Time      `xmlrpc:"repeat_until,omitempty"`
	TaskIds        *Relation  `xmlrpc:"task_ids,omitempty"`
	WriteDate      *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid       *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// ProjectTaskRecurrences represents array of project.task.recurrence model.
type ProjectTaskRecurrences []ProjectTaskRecurrence

// ProjectTaskRecurrenceModel is the odoo model name.
const ProjectTaskRecurrenceModel = "project.task.recurrence"

// Many2One convert ProjectTaskRecurrence to *Many2One.
func (ptr *ProjectTaskRecurrence) Many2One() *Many2One {
	return NewMany2One(ptr.Id.Get(), "")
}

// CreateProjectTaskRecurrence creates a new project.task.recurrence model and returns its id.
func (c *Client) CreateProjectTaskRecurrence(ptr *ProjectTaskRecurrence) (int64, error) {
	ids, err := c.CreateProjectTaskRecurrences([]*ProjectTaskRecurrence{ptr})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateProjectTaskRecurrence creates a new project.task.recurrence model and returns its id.
func (c *Client) CreateProjectTaskRecurrences(ptrs []*ProjectTaskRecurrence) ([]int64, error) {
	var vv []interface{}
	for _, v := range ptrs {
		vv = append(vv, v)
	}
	return c.Create(ProjectTaskRecurrenceModel, vv, nil)
}

// UpdateProjectTaskRecurrence updates an existing project.task.recurrence record.
func (c *Client) UpdateProjectTaskRecurrence(ptr *ProjectTaskRecurrence) error {
	return c.UpdateProjectTaskRecurrences([]int64{ptr.Id.Get()}, ptr)
}

// UpdateProjectTaskRecurrences updates existing project.task.recurrence records.
// All records (represented by ids) will be updated by ptr values.
func (c *Client) UpdateProjectTaskRecurrences(ids []int64, ptr *ProjectTaskRecurrence) error {
	return c.Update(ProjectTaskRecurrenceModel, ids, ptr, nil)
}

// DeleteProjectTaskRecurrence deletes an existing project.task.recurrence record.
func (c *Client) DeleteProjectTaskRecurrence(id int64) error {
	return c.DeleteProjectTaskRecurrences([]int64{id})
}

// DeleteProjectTaskRecurrences deletes existing project.task.recurrence records.
func (c *Client) DeleteProjectTaskRecurrences(ids []int64) error {
	return c.Delete(ProjectTaskRecurrenceModel, ids)
}

// GetProjectTaskRecurrence gets project.task.recurrence existing record.
func (c *Client) GetProjectTaskRecurrence(id int64) (*ProjectTaskRecurrence, error) {
	ptrs, err := c.GetProjectTaskRecurrences([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ptrs)[0]), nil
}

// GetProjectTaskRecurrences gets project.task.recurrence existing records.
func (c *Client) GetProjectTaskRecurrences(ids []int64) (*ProjectTaskRecurrences, error) {
	ptrs := &ProjectTaskRecurrences{}
	if err := c.Read(ProjectTaskRecurrenceModel, ids, nil, ptrs); err != nil {
		return nil, err
	}
	return ptrs, nil
}

// FindProjectTaskRecurrence finds project.task.recurrence record by querying it with criteria.
func (c *Client) FindProjectTaskRecurrence(criteria *Criteria) (*ProjectTaskRecurrence, error) {
	ptrs := &ProjectTaskRecurrences{}
	if err := c.SearchRead(ProjectTaskRecurrenceModel, criteria, NewOptions().Limit(1), ptrs); err != nil {
		return nil, err
	}
	return &((*ptrs)[0]), nil
}

// FindProjectTaskRecurrences finds project.task.recurrence records by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskRecurrences(criteria *Criteria, options *Options) (*ProjectTaskRecurrences, error) {
	ptrs := &ProjectTaskRecurrences{}
	if err := c.SearchRead(ProjectTaskRecurrenceModel, criteria, options, ptrs); err != nil {
		return nil, err
	}
	return ptrs, nil
}

// FindProjectTaskRecurrenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindProjectTaskRecurrenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ProjectTaskRecurrenceModel, criteria, options)
}

// FindProjectTaskRecurrenceId finds record id by querying it with criteria.
func (c *Client) FindProjectTaskRecurrenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ProjectTaskRecurrenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
