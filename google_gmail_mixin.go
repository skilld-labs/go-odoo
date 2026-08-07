package odoo

// GoogleGmailMixin represents google.gmail.mixin model.
type GoogleGmailMixin struct {
	Active                           *Bool   `xmlrpc:"active,omitempty"`
	DisplayName                      *String `xmlrpc:"display_name,omitempty"`
	GoogleGmailAccessToken           *String `xmlrpc:"google_gmail_access_token,omitempty"`
	GoogleGmailAccessTokenExpiration *Int    `xmlrpc:"google_gmail_access_token_expiration,omitempty"`
	GoogleGmailRefreshToken          *String `xmlrpc:"google_gmail_refresh_token,omitempty"`
	GoogleGmailUri                   *String `xmlrpc:"google_gmail_uri,omitempty"`
	Id                               *Int    `xmlrpc:"id,omitempty"`
}

// GoogleGmailMixins represents array of google.gmail.mixin model.
type GoogleGmailMixins []GoogleGmailMixin

// GoogleGmailMixinModel is the odoo model name.
const GoogleGmailMixinModel = "google.gmail.mixin"

// Many2One convert GoogleGmailMixin to *Many2One.
func (ggm *GoogleGmailMixin) Many2One() *Many2One {
	return NewMany2One(ggm.Id.Get(), "")
}

// CreateGoogleGmailMixin creates a new google.gmail.mixin model and returns its id.
func (c *Client) CreateGoogleGmailMixin(ggm *GoogleGmailMixin) (int64, error) {
	ids, err := c.CreateGoogleGmailMixins([]*GoogleGmailMixin{ggm})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateGoogleGmailMixin creates a new google.gmail.mixin model and returns its id.
func (c *Client) CreateGoogleGmailMixins(ggms []*GoogleGmailMixin) ([]int64, error) {
	var vv []interface{}
	for _, v := range ggms {
		vv = append(vv, v)
	}
	return c.Create(GoogleGmailMixinModel, vv, nil)
}

// UpdateGoogleGmailMixin updates an existing google.gmail.mixin record.
func (c *Client) UpdateGoogleGmailMixin(ggm *GoogleGmailMixin) error {
	return c.UpdateGoogleGmailMixins([]int64{ggm.Id.Get()}, ggm)
}

// UpdateGoogleGmailMixins updates existing google.gmail.mixin records.
// All records (represented by ids) will be updated by ggm values.
func (c *Client) UpdateGoogleGmailMixins(ids []int64, ggm *GoogleGmailMixin) error {
	return c.Update(GoogleGmailMixinModel, ids, ggm, nil)
}

// DeleteGoogleGmailMixin deletes an existing google.gmail.mixin record.
func (c *Client) DeleteGoogleGmailMixin(id int64) error {
	return c.DeleteGoogleGmailMixins([]int64{id})
}

// DeleteGoogleGmailMixins deletes existing google.gmail.mixin records.
func (c *Client) DeleteGoogleGmailMixins(ids []int64) error {
	return c.Delete(GoogleGmailMixinModel, ids)
}

// GetGoogleGmailMixin gets google.gmail.mixin existing record.
func (c *Client) GetGoogleGmailMixin(id int64) (*GoogleGmailMixin, error) {
	ggms, err := c.GetGoogleGmailMixins([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ggms)[0]), nil
}

// GetGoogleGmailMixins gets google.gmail.mixin existing records.
func (c *Client) GetGoogleGmailMixins(ids []int64) (*GoogleGmailMixins, error) {
	ggms := &GoogleGmailMixins{}
	if err := c.Read(GoogleGmailMixinModel, ids, nil, ggms); err != nil {
		return nil, err
	}
	return ggms, nil
}

// FindGoogleGmailMixin finds google.gmail.mixin record by querying it with criteria.
func (c *Client) FindGoogleGmailMixin(criteria *Criteria) (*GoogleGmailMixin, error) {
	ggms := &GoogleGmailMixins{}
	if err := c.SearchRead(GoogleGmailMixinModel, criteria, NewOptions().Limit(1), ggms); err != nil {
		return nil, err
	}
	return &((*ggms)[0]), nil
}

// FindGoogleGmailMixins finds google.gmail.mixin records by querying it
// and filtering it with criteria and options.
func (c *Client) FindGoogleGmailMixins(criteria *Criteria, options *Options) (*GoogleGmailMixins, error) {
	ggms := &GoogleGmailMixins{}
	if err := c.SearchRead(GoogleGmailMixinModel, criteria, options, ggms); err != nil {
		return nil, err
	}
	return ggms, nil
}

// FindGoogleGmailMixinIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindGoogleGmailMixinIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(GoogleGmailMixinModel, criteria, options)
}

// FindGoogleGmailMixinId finds record id by querying it with criteria.
func (c *Client) FindGoogleGmailMixinId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(GoogleGmailMixinModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
