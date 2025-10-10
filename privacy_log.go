package odoo

// PrivacyLog represents privacy.log model.
type PrivacyLog struct {
	AdditionalNote     *String   `xmlrpc:"additional_note,omitempty"`
	AnonymizedEmail    *String   `xmlrpc:"anonymized_email,omitempty"`
	AnonymizedName     *String   `xmlrpc:"anonymized_name,omitempty"`
	CreateDate         *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid          *Many2One `xmlrpc:"create_uid,omitempty"`
	Date               *Time     `xmlrpc:"date,omitempty"`
	DisplayName        *String   `xmlrpc:"display_name,omitempty"`
	ExecutionDetails   *String   `xmlrpc:"execution_details,omitempty"`
	Id                 *Int      `xmlrpc:"id,omitempty"`
	RecordsDescription *String   `xmlrpc:"records_description,omitempty"`
	UserId             *Many2One `xmlrpc:"user_id,omitempty"`
	WriteDate          *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid           *Many2One `xmlrpc:"write_uid,omitempty"`
}

// PrivacyLogs represents array of privacy.log model.
type PrivacyLogs []PrivacyLog

// PrivacyLogModel is the odoo model name.
const PrivacyLogModel = "privacy.log"

// Many2One convert PrivacyLog to *Many2One.
func (pl *PrivacyLog) Many2One() *Many2One {
	return NewMany2One(pl.Id.Get(), "")
}

// CreatePrivacyLog creates a new privacy.log model and returns its id.
func (c *Client) CreatePrivacyLog(pl *PrivacyLog) (int64, error) {
	ids, err := c.CreatePrivacyLogs([]*PrivacyLog{pl})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreatePrivacyLog creates a new privacy.log model and returns its id.
func (c *Client) CreatePrivacyLogs(pls []*PrivacyLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range pls {
		vv = append(vv, v)
	}
	return c.Create(PrivacyLogModel, vv, nil)
}

// UpdatePrivacyLog updates an existing privacy.log record.
func (c *Client) UpdatePrivacyLog(pl *PrivacyLog) error {
	return c.UpdatePrivacyLogs([]int64{pl.Id.Get()}, pl)
}

// UpdatePrivacyLogs updates existing privacy.log records.
// All records (represented by ids) will be updated by pl values.
func (c *Client) UpdatePrivacyLogs(ids []int64, pl *PrivacyLog) error {
	return c.Update(PrivacyLogModel, ids, pl, nil)
}

// DeletePrivacyLog deletes an existing privacy.log record.
func (c *Client) DeletePrivacyLog(id int64) error {
	return c.DeletePrivacyLogs([]int64{id})
}

// DeletePrivacyLogs deletes existing privacy.log records.
func (c *Client) DeletePrivacyLogs(ids []int64) error {
	return c.Delete(PrivacyLogModel, ids)
}

// GetPrivacyLog gets privacy.log existing record.
func (c *Client) GetPrivacyLog(id int64) (*PrivacyLog, error) {
	pls, err := c.GetPrivacyLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*pls)[0]), nil
}

// GetPrivacyLogs gets privacy.log existing records.
func (c *Client) GetPrivacyLogs(ids []int64) (*PrivacyLogs, error) {
	pls := &PrivacyLogs{}
	if err := c.Read(PrivacyLogModel, ids, nil, pls); err != nil {
		return nil, err
	}
	return pls, nil
}

// FindPrivacyLog finds privacy.log record by querying it with criteria.
func (c *Client) FindPrivacyLog(criteria *Criteria) (*PrivacyLog, error) {
	pls := &PrivacyLogs{}
	if err := c.SearchRead(PrivacyLogModel, criteria, NewOptions().Limit(1), pls); err != nil {
		return nil, err
	}
	return &((*pls)[0]), nil
}

// FindPrivacyLogs finds privacy.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrivacyLogs(criteria *Criteria, options *Options) (*PrivacyLogs, error) {
	pls := &PrivacyLogs{}
	if err := c.SearchRead(PrivacyLogModel, criteria, options, pls); err != nil {
		return nil, err
	}
	return pls, nil
}

// FindPrivacyLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindPrivacyLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(PrivacyLogModel, criteria, options)
}

// FindPrivacyLogId finds record id by querying it with criteria.
func (c *Client) FindPrivacyLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(PrivacyLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
