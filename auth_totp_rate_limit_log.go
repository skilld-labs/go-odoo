package odoo

// AuthTotpRateLimitLog represents auth.totp.rate.limit.log model.
type AuthTotpRateLimitLog struct {
	CreateDate  *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid   *Many2One  `xmlrpc:"create_uid,omitempty"`
	DisplayName *String    `xmlrpc:"display_name,omitempty"`
	Id          *Int       `xmlrpc:"id,omitempty"`
	Ip          *String    `xmlrpc:"ip,omitempty"`
	LimitType   *Selection `xmlrpc:"limit_type,omitempty"`
	UserId      *Many2One  `xmlrpc:"user_id,omitempty"`
	WriteDate   *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid    *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// AuthTotpRateLimitLogs represents array of auth.totp.rate.limit.log model.
type AuthTotpRateLimitLogs []AuthTotpRateLimitLog

// AuthTotpRateLimitLogModel is the odoo model name.
const AuthTotpRateLimitLogModel = "auth.totp.rate.limit.log"

// Many2One convert AuthTotpRateLimitLog to *Many2One.
func (atrll *AuthTotpRateLimitLog) Many2One() *Many2One {
	return NewMany2One(atrll.Id.Get(), "")
}

// CreateAuthTotpRateLimitLog creates a new auth.totp.rate.limit.log model and returns its id.
func (c *Client) CreateAuthTotpRateLimitLog(atrll *AuthTotpRateLimitLog) (int64, error) {
	ids, err := c.CreateAuthTotpRateLimitLogs([]*AuthTotpRateLimitLog{atrll})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateAuthTotpRateLimitLog creates a new auth.totp.rate.limit.log model and returns its id.
func (c *Client) CreateAuthTotpRateLimitLogs(atrlls []*AuthTotpRateLimitLog) ([]int64, error) {
	var vv []interface{}
	for _, v := range atrlls {
		vv = append(vv, v)
	}
	return c.Create(AuthTotpRateLimitLogModel, vv, nil)
}

// UpdateAuthTotpRateLimitLog updates an existing auth.totp.rate.limit.log record.
func (c *Client) UpdateAuthTotpRateLimitLog(atrll *AuthTotpRateLimitLog) error {
	return c.UpdateAuthTotpRateLimitLogs([]int64{atrll.Id.Get()}, atrll)
}

// UpdateAuthTotpRateLimitLogs updates existing auth.totp.rate.limit.log records.
// All records (represented by ids) will be updated by atrll values.
func (c *Client) UpdateAuthTotpRateLimitLogs(ids []int64, atrll *AuthTotpRateLimitLog) error {
	return c.Update(AuthTotpRateLimitLogModel, ids, atrll, nil)
}

// DeleteAuthTotpRateLimitLog deletes an existing auth.totp.rate.limit.log record.
func (c *Client) DeleteAuthTotpRateLimitLog(id int64) error {
	return c.DeleteAuthTotpRateLimitLogs([]int64{id})
}

// DeleteAuthTotpRateLimitLogs deletes existing auth.totp.rate.limit.log records.
func (c *Client) DeleteAuthTotpRateLimitLogs(ids []int64) error {
	return c.Delete(AuthTotpRateLimitLogModel, ids)
}

// GetAuthTotpRateLimitLog gets auth.totp.rate.limit.log existing record.
func (c *Client) GetAuthTotpRateLimitLog(id int64) (*AuthTotpRateLimitLog, error) {
	atrlls, err := c.GetAuthTotpRateLimitLogs([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*atrlls)[0]), nil
}

// GetAuthTotpRateLimitLogs gets auth.totp.rate.limit.log existing records.
func (c *Client) GetAuthTotpRateLimitLogs(ids []int64) (*AuthTotpRateLimitLogs, error) {
	atrlls := &AuthTotpRateLimitLogs{}
	if err := c.Read(AuthTotpRateLimitLogModel, ids, nil, atrlls); err != nil {
		return nil, err
	}
	return atrlls, nil
}

// FindAuthTotpRateLimitLog finds auth.totp.rate.limit.log record by querying it with criteria.
func (c *Client) FindAuthTotpRateLimitLog(criteria *Criteria) (*AuthTotpRateLimitLog, error) {
	atrlls := &AuthTotpRateLimitLogs{}
	if err := c.SearchRead(AuthTotpRateLimitLogModel, criteria, NewOptions().Limit(1), atrlls); err != nil {
		return nil, err
	}
	return &((*atrlls)[0]), nil
}

// FindAuthTotpRateLimitLogs finds auth.totp.rate.limit.log records by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthTotpRateLimitLogs(criteria *Criteria, options *Options) (*AuthTotpRateLimitLogs, error) {
	atrlls := &AuthTotpRateLimitLogs{}
	if err := c.SearchRead(AuthTotpRateLimitLogModel, criteria, options, atrlls); err != nil {
		return nil, err
	}
	return atrlls, nil
}

// FindAuthTotpRateLimitLogIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindAuthTotpRateLimitLogIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(AuthTotpRateLimitLogModel, criteria, options)
}

// FindAuthTotpRateLimitLogId finds record id by querying it with criteria.
func (c *Client) FindAuthTotpRateLimitLogId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(AuthTotpRateLimitLogModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
