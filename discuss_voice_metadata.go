package odoo

// DiscussVoiceMetadata represents discuss.voice.metadata model.
type DiscussVoiceMetadata struct {
	AttachmentId *Many2One `xmlrpc:"attachment_id,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// DiscussVoiceMetadatas represents array of discuss.voice.metadata model.
type DiscussVoiceMetadatas []DiscussVoiceMetadata

// DiscussVoiceMetadataModel is the odoo model name.
const DiscussVoiceMetadataModel = "discuss.voice.metadata"

// Many2One convert DiscussVoiceMetadata to *Many2One.
func (dvm *DiscussVoiceMetadata) Many2One() *Many2One {
	return NewMany2One(dvm.Id.Get(), "")
}

// CreateDiscussVoiceMetadata creates a new discuss.voice.metadata model and returns its id.
func (c *Client) CreateDiscussVoiceMetadata(dvm *DiscussVoiceMetadata) (int64, error) {
	ids, err := c.CreateDiscussVoiceMetadatas([]*DiscussVoiceMetadata{dvm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateDiscussVoiceMetadata creates a new discuss.voice.metadata model and returns its id.
func (c *Client) CreateDiscussVoiceMetadatas(dvms []*DiscussVoiceMetadata) ([]int64, error) {
	var vv []interface{}
	for _, v := range dvms {
		vv = append(vv, v)
	}
	return c.Create(DiscussVoiceMetadataModel, vv, nil)
}

// UpdateDiscussVoiceMetadata updates an existing discuss.voice.metadata record.
func (c *Client) UpdateDiscussVoiceMetadata(dvm *DiscussVoiceMetadata) error {
	return c.UpdateDiscussVoiceMetadatas([]int64{dvm.Id.Get()}, dvm)
}

// UpdateDiscussVoiceMetadatas updates existing discuss.voice.metadata records.
// All records (represented by ids) will be updated by dvm values.
func (c *Client) UpdateDiscussVoiceMetadatas(ids []int64, dvm *DiscussVoiceMetadata) error {
	return c.Update(DiscussVoiceMetadataModel, ids, dvm, nil)
}

// DeleteDiscussVoiceMetadata deletes an existing discuss.voice.metadata record.
func (c *Client) DeleteDiscussVoiceMetadata(id int64) error {
	return c.DeleteDiscussVoiceMetadatas([]int64{id})
}

// DeleteDiscussVoiceMetadatas deletes existing discuss.voice.metadata records.
func (c *Client) DeleteDiscussVoiceMetadatas(ids []int64) error {
	return c.Delete(DiscussVoiceMetadataModel, ids)
}

// GetDiscussVoiceMetadata gets discuss.voice.metadata existing record.
func (c *Client) GetDiscussVoiceMetadata(id int64) (*DiscussVoiceMetadata, error) {
	dvms, err := c.GetDiscussVoiceMetadatas([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*dvms)[0]), nil
}

// GetDiscussVoiceMetadatas gets discuss.voice.metadata existing records.
func (c *Client) GetDiscussVoiceMetadatas(ids []int64) (*DiscussVoiceMetadatas, error) {
	dvms := &DiscussVoiceMetadatas{}
	if err := c.Read(DiscussVoiceMetadataModel, ids, nil, dvms); err != nil {
		return nil, err
	}
	return dvms, nil
}

// FindDiscussVoiceMetadata finds discuss.voice.metadata record by querying it with criteria.
func (c *Client) FindDiscussVoiceMetadata(criteria *Criteria) (*DiscussVoiceMetadata, error) {
	dvms := &DiscussVoiceMetadatas{}
	if err := c.SearchRead(DiscussVoiceMetadataModel, criteria, NewOptions().Limit(1), dvms); err != nil {
		return nil, err
	}
	return &((*dvms)[0]), nil
}

// FindDiscussVoiceMetadatas finds discuss.voice.metadata records by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussVoiceMetadatas(criteria *Criteria, options *Options) (*DiscussVoiceMetadatas, error) {
	dvms := &DiscussVoiceMetadatas{}
	if err := c.SearchRead(DiscussVoiceMetadataModel, criteria, options, dvms); err != nil {
		return nil, err
	}
	return dvms, nil
}

// FindDiscussVoiceMetadataIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindDiscussVoiceMetadataIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(DiscussVoiceMetadataModel, criteria, options)
}

// FindDiscussVoiceMetadataId finds record id by querying it with criteria.
func (c *Client) FindDiscussVoiceMetadataId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(DiscussVoiceMetadataModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
