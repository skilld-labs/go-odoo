package odoo

// ResUsersLog represents res.users.log model.
type ResUsersLog struct {
	LastUpdate  *Time     `xmlrpc:"__last_update,omitempty"`
	CreateDate  *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName *String   `xmlrpc:"display_name,omitempty"`
	Id          *Int      `xmlrpc:"id,omitempty"`
	WriteDate   *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One `xmlrpc:"write_uid,omitempty"`
}

// ResUsersLogs represents array of res.users.log model.
type ResUsersLogs []ResUsersLog

// ResUsersLogModel is the odoo model name.
const ResUsersLogModel = "res.users.log"

// Many2One convert ResUsersLog to *Many2One.
func (rul *ResUsersLog) Many2One() *Many2One {
	return NewMany2One(rul.Id.Get(), "")
}

// CreateResUsersLog creates a new res.users.log model and returns its id.
func (c *Client) CreateResUsersLog(rul *ResUsersLog) (int64, error) {
	ids, err := c.CreateResUsersLogs([]*ResUsersLog{rul})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateResUsersLog creates a new res.users.log model and returns its id.
func (c *Client) CreateResUsersLogs(ruls []*ResUsersLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range ruls {
		vv = append(vv, v)
	}
	return c.Create(ResUsersLogModel, vv, nil)
}

// UpdateResUsersLog updates an existing res.users.log record.
func (c *Client) UpdateResUsersLog(rul *ResUsersLog) error {
	return c.UpdateResUsersLogs([]int64{rul.Id.Get()}, rul)
}

// UpdateResUsersLogs updates existing res.users.log records.
// All records (represented by ids) will be updated by rul values.
func (c *Client) UpdateResUsersLogs(ids []int64, rul *ResUsersLog) error {
	return c.Update(ResUsersLogModel, ids, rul, nil)
}

// DeleteResUsersLog deletes an existing res.users.log record.
func (c *Client) DeleteResUsersLog(id int64) error {
	return c.DeleteResUsersLogs([]int64{id})
}

// DeleteResUsersLogs deletes existing res.users.log records.
func (c *Client) DeleteResUsersLogs(ids []int64) error {
	return c.Delete(ResUsersLogModel, ids)
}

// GetResUsersLog gets res.users.log existing record.
func (c *Client) GetResUsersLog(id int64) (*ResUsersLog, error) {
	ruls, err := c.GetResUsersLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ruls)[0]), nil
}

// GetResUsersLogs gets res.users.log existing records.
func (c *Client) GetResUsersLogs(ids []int64) (*ResUsersLogs, error) {
	ruls := &ResUsersLogs{}
	if err := c.Read(ResUsersLogModel, ids, nil, ruls); err != nil {
		return nil, err
	}
	return ruls, nil
}

// FindResUsersLog finds res.users.log record by querying it with criteria.
func (c *Client) FindResUsersLog(criteria *Criteria) (*ResUsersLog, error) {
	ruls := &ResUsersLogs{}
	if err := c.SearchRead(ResUsersLogModel, criteria, NewOptions().Limit(1), ruls); err != nil {
		return nil, err
	}
	return &((*ruls)[0]), nil
}

// FindResUsersLogs finds res.users.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersLogs(criteria *Criteria, options *Options) (*ResUsersLogs, error) {
	ruls := &ResUsersLogs{}
	if err := c.SearchRead(ResUsersLogModel, criteria, options, ruls); err != nil {
		return nil, err
	}
	return ruls, nil
}

// FindResUsersLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindResUsersLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(ResUsersLogModel, criteria, options)
}

// FindResUsersLogId finds record id by querying it with criteria.
func (c *Client) FindResUsersLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(ResUsersLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
