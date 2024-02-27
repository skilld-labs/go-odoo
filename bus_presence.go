package odoo

// BusPresence represents bus.presence model.
type BusPresence struct {
	LastUpdate   *Time      `xmlrpc:"__last_update,omitempty"`
	DisplayName  *String    `xmlrpc:"display_name,omitempty"`
	Id           *Int       `xmlrpc:"id,omitempty"`
	LastPoll     *Time      `xmlrpc:"last_poll,omitempty"`
	LastPresence *Time      `xmlrpc:"last_presence,omitempty"`
	Status       *Selection `xmlrpc:"status,omitempty"`
	UserId       *Many2One  `xmlrpc:"user_id,omitempty"`
}

// BusPresences represents array of bus.presence model.
type BusPresences []BusPresence

// BusPresenceModel is the odoo model name.
const BusPresenceModel = "bus.presence"

// Many2One convert BusPresence to *Many2One.
func (bp *BusPresence) Many2One() *Many2One {
	return NewMany2One(bp.Id.Get(), "")
}

// CreateBusPresence creates a new bus.presence model and returns its id.
func (c *Client) CreateBusPresence(bp *BusPresence) (int64, error) {
	ids, err := c.CreateBusPresences([]*BusPresence{bp})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateBusPresence creates a new bus.presence model and returns its id.
func (c *Client) CreateBusPresences(bps []*BusPresence) ([]int64, error) {
	var vv []interface{}
	for _, v := range bps {
		vv = append(vv, v)
	}
	return c.Create(BusPresenceModel, vv, nil)
}

// UpdateBusPresence updates an existing bus.presence record.
func (c *Client) UpdateBusPresence(bp *BusPresence) error {
	return c.UpdateBusPresences([]int64{bp.Id.Get()}, bp)
}

// UpdateBusPresences updates existing bus.presence records.
// All records (represented by ids) will be updated by bp values.
func (c *Client) UpdateBusPresences(ids []int64, bp *BusPresence) error {
	return c.Update(BusPresenceModel, ids, bp, nil)
}

// DeleteBusPresence deletes an existing bus.presence record.
func (c *Client) DeleteBusPresence(id int64) error {
	return c.DeleteBusPresences([]int64{id})
}

// DeleteBusPresences deletes existing bus.presence records.
func (c *Client) DeleteBusPresences(ids []int64) error {
	return c.Delete(BusPresenceModel, ids)
}

// GetBusPresence gets bus.presence existing record.
func (c *Client) GetBusPresence(id int64) (*BusPresence, error) {
	bps, err := c.GetBusPresences([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*bps)[0]), nil
}

// GetBusPresences gets bus.presence existing records.
func (c *Client) GetBusPresences(ids []int64) (*BusPresences, error) {
	bps := &BusPresences{}
	if err := c.Read(BusPresenceModel, ids, nil, bps); err != nil {
		return nil, err
	}
	return bps, nil
}

// FindBusPresence finds bus.presence record by querying it with criteria.
func (c *Client) FindBusPresence(criteria *Criteria) (*BusPresence, error) {
	bps := &BusPresences{}
	if err := c.SearchRead(BusPresenceModel, criteria, NewOptions().Limit(1), bps); err != nil {
		return nil, err
	}
	return &((*bps)[0]), nil
}

// FindBusPresences finds bus.presence records by querying it
// and filtering it with criteria and options.
func (c *Client) FindBusPresences(criteria *Criteria, options *Options) (*BusPresences, error) {
	bps := &BusPresences{}
	if err := c.SearchRead(BusPresenceModel, criteria, options, bps); err != nil {
		return nil, err
	}
	return bps, nil
}

// FindBusPresenceIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindBusPresenceIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(BusPresenceModel, criteria, options)
}

// FindBusPresenceId finds record id by querying it with criteria.
func (c *Client) FindBusPresenceId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(BusPresenceModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
