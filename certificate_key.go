package odoo

// CertificateKey represents certificate.key model.
type CertificateKey struct {
	Active       *Bool     `xmlrpc:"active,omitempty"`
	CompanyId    *Many2One `xmlrpc:"company_id,omitempty"`
	Content      *String   `xmlrpc:"content,omitempty"`
	CreateDate   *Time     `xmlrpc:"create_date,omitempty"`
	CreateUid    *Many2One `xmlrpc:"create_uid,omitempty"`
	DisplayName  *String   `xmlrpc:"display_name,omitempty"`
	Id           *Int      `xmlrpc:"id,omitempty"`
	LoadingError *String   `xmlrpc:"loading_error,omitempty"`
	Name         *String   `xmlrpc:"name,omitempty"`
	Password     *String   `xmlrpc:"password,omitempty"`
	PemKey       *String   `xmlrpc:"pem_key,omitempty"`
	Public       *Bool     `xmlrpc:"public,omitempty"`
	WriteDate    *Time     `xmlrpc:"write_date,omitempty"`
	WriteUid     *Many2One `xmlrpc:"write_uid,omitempty"`
}

// CertificateKeys represents array of certificate.key model.
type CertificateKeys []CertificateKey

// CertificateKeyModel is the odoo model name.
const CertificateKeyModel = "certificate.key"

// Many2One convert CertificateKey to *Many2One.
func (ck *CertificateKey) Many2One() *Many2One {
	return NewMany2One(ck.Id.Get(), "")
}

// CreateCertificateKey creates a new certificate.key model and returns its id.
func (c *Client) CreateCertificateKey(ck *CertificateKey) (int64, error) {
	ids, err := c.CreateCertificateKeys([]*CertificateKey{ck})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCertificateKey creates a new certificate.key model and returns its id.
func (c *Client) CreateCertificateKeys(cks []*CertificateKey) ([]int64, error) {
	var vv []interface{}
	for _, v := range cks {
		vv = append(vv, v)
	}
	return c.Create(CertificateKeyModel, vv, nil)
}

// UpdateCertificateKey updates an existing certificate.key record.
func (c *Client) UpdateCertificateKey(ck *CertificateKey) error {
	return c.UpdateCertificateKeys([]int64{ck.Id.Get()}, ck)
}

// UpdateCertificateKeys updates existing certificate.key records.
// All records (represented by ids) will be updated by ck values.
func (c *Client) UpdateCertificateKeys(ids []int64, ck *CertificateKey) error {
	return c.Update(CertificateKeyModel, ids, ck, nil)
}

// DeleteCertificateKey deletes an existing certificate.key record.
func (c *Client) DeleteCertificateKey(id int64) error {
	return c.DeleteCertificateKeys([]int64{id})
}

// DeleteCertificateKeys deletes existing certificate.key records.
func (c *Client) DeleteCertificateKeys(ids []int64) error {
	return c.Delete(CertificateKeyModel, ids)
}

// GetCertificateKey gets certificate.key existing record.
func (c *Client) GetCertificateKey(id int64) (*CertificateKey, error) {
	cks, err := c.GetCertificateKeys([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*cks)[0]), nil
}

// GetCertificateKeys gets certificate.key existing records.
func (c *Client) GetCertificateKeys(ids []int64) (*CertificateKeys, error) {
	cks := &CertificateKeys{}
	if err := c.Read(CertificateKeyModel, ids, nil, cks); err != nil {
		return nil, err
	}
	return cks, nil
}

// FindCertificateKey finds certificate.key record by querying it with criteria.
func (c *Client) FindCertificateKey(criteria *Criteria) (*CertificateKey, error) {
	cks := &CertificateKeys{}
	if err := c.SearchRead(CertificateKeyModel, criteria, NewOptions().Limit(1), cks); err != nil {
		return nil, err
	}
	return &((*cks)[0]), nil
}

// FindCertificateKeys finds certificate.key records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCertificateKeys(criteria *Criteria, options *Options) (*CertificateKeys, error) {
	cks := &CertificateKeys{}
	if err := c.SearchRead(CertificateKeyModel, criteria, options, cks); err != nil {
		return nil, err
	}
	return cks, nil
}

// FindCertificateKeyIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCertificateKeyIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CertificateKeyModel, criteria, options)
}

// FindCertificateKeyId finds record id by querying it with criteria.
func (c *Client) FindCertificateKeyId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CertificateKeyModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
