package odoo

// CertificateCertificate represents certificate.certificate model.
type CertificateCertificate struct {
	Active            *Bool      `xmlrpc:"active,omitempty"`
	CompanyId         *Many2One  `xmlrpc:"company_id,omitempty"`
	Content           *String    `xmlrpc:"content,omitempty"`
	ContentFormat     *Selection `xmlrpc:"content_format,omitempty"`
	CountryCode       *String    `xmlrpc:"country_code,omitempty"`
	CreateDate        *Time      `xmlrpc:"create_date,omitempty"`
	CreateUid         *Many2One  `xmlrpc:"create_uid,omitempty"`
	DateEnd           *Time      `xmlrpc:"date_end,omitempty"`
	DateStart         *Time      `xmlrpc:"date_start,omitempty"`
	DisplayName       *String    `xmlrpc:"display_name,omitempty"`
	Id                *Int       `xmlrpc:"id,omitempty"`
	IsValid           *Bool      `xmlrpc:"is_valid,omitempty"`
	IssuerCertId      *Many2One  `xmlrpc:"issuer_cert_id,omitempty"`
	LoadingError      *String    `xmlrpc:"loading_error,omitempty"`
	Name              *String    `xmlrpc:"name,omitempty"`
	PemCertificate    *String    `xmlrpc:"pem_certificate,omitempty"`
	Pkcs12Password    *String    `xmlrpc:"pkcs12_password,omitempty"`
	PrivateKeyId      *Many2One  `xmlrpc:"private_key_id,omitempty"`
	PublicKeyId       *Many2One  `xmlrpc:"public_key_id,omitempty"`
	Scope             *Selection `xmlrpc:"scope,omitempty"`
	SerialNumber      *String    `xmlrpc:"serial_number,omitempty"`
	SubjectCommonName *String    `xmlrpc:"subject_common_name,omitempty"`
	WriteDate         *Time      `xmlrpc:"write_date,omitempty"`
	WriteUid          *Many2One  `xmlrpc:"write_uid,omitempty"`
}

// CertificateCertificates represents array of certificate.certificate model.
type CertificateCertificates []CertificateCertificate

// CertificateCertificateModel is the odoo model name.
const CertificateCertificateModel = "certificate.certificate"

// Many2One convert CertificateCertificate to *Many2One.
func (cc *CertificateCertificate) Many2One() *Many2One {
	return NewMany2One(cc.Id.Get(), "")
}

// CreateCertificateCertificate creates a new certificate.certificate model and returns its id.
func (c *Client) CreateCertificateCertificate(cc *CertificateCertificate) (int64, error) {
	ids, err := c.CreateCertificateCertificates([]*CertificateCertificate{cc})
	if err != nil {
		return -1, err
	}
	if len(ids) == 0 {
		return -1, nil
	}
	return ids[0], nil
}

// CreateCertificateCertificate creates a new certificate.certificate model and returns its id.
func (c *Client) CreateCertificateCertificates(ccs []*CertificateCertificate) ([]int64, error) {
	var vv []interface{}
	for _, v := range ccs {
		vv = append(vv, v)
	}
	return c.Create(CertificateCertificateModel, vv, nil)
}

// UpdateCertificateCertificate updates an existing certificate.certificate record.
func (c *Client) UpdateCertificateCertificate(cc *CertificateCertificate) error {
	return c.UpdateCertificateCertificates([]int64{cc.Id.Get()}, cc)
}

// UpdateCertificateCertificates updates existing certificate.certificate records.
// All records (represented by ids) will be updated by cc values.
func (c *Client) UpdateCertificateCertificates(ids []int64, cc *CertificateCertificate) error {
	return c.Update(CertificateCertificateModel, ids, cc, nil)
}

// DeleteCertificateCertificate deletes an existing certificate.certificate record.
func (c *Client) DeleteCertificateCertificate(id int64) error {
	return c.DeleteCertificateCertificates([]int64{id})
}

// DeleteCertificateCertificates deletes existing certificate.certificate records.
func (c *Client) DeleteCertificateCertificates(ids []int64) error {
	return c.Delete(CertificateCertificateModel, ids)
}

// GetCertificateCertificate gets certificate.certificate existing record.
func (c *Client) GetCertificateCertificate(id int64) (*CertificateCertificate, error) {
	ccs, err := c.GetCertificateCertificates([]int64{id})
	if err != nil {
		return nil, err
	}
	return &((*ccs)[0]), nil
}

// GetCertificateCertificates gets certificate.certificate existing records.
func (c *Client) GetCertificateCertificates(ids []int64) (*CertificateCertificates, error) {
	ccs := &CertificateCertificates{}
	if err := c.Read(CertificateCertificateModel, ids, nil, ccs); err != nil {
		return nil, err
	}
	return ccs, nil
}

// FindCertificateCertificate finds certificate.certificate record by querying it with criteria.
func (c *Client) FindCertificateCertificate(criteria *Criteria) (*CertificateCertificate, error) {
	ccs := &CertificateCertificates{}
	if err := c.SearchRead(CertificateCertificateModel, criteria, NewOptions().Limit(1), ccs); err != nil {
		return nil, err
	}
	return &((*ccs)[0]), nil
}

// FindCertificateCertificates finds certificate.certificate records by querying it
// and filtering it with criteria and options.
func (c *Client) FindCertificateCertificates(criteria *Criteria, options *Options) (*CertificateCertificates, error) {
	ccs := &CertificateCertificates{}
	if err := c.SearchRead(CertificateCertificateModel, criteria, options, ccs); err != nil {
		return nil, err
	}
	return ccs, nil
}

// FindCertificateCertificateIds finds records ids by querying it
// and filtering it with criteria and options.
func (c *Client) FindCertificateCertificateIds(criteria *Criteria, options *Options) ([]int64, error) {
	return c.Search(CertificateCertificateModel, criteria, options)
}

// FindCertificateCertificateId finds record id by querying it with criteria.
func (c *Client) FindCertificateCertificateId(criteria *Criteria, options *Options) (int64, error) {
	ids, err := c.Search(CertificateCertificateModel, criteria, options)
	if err != nil {
		return -1, err
	}
	return ids[0], nil
}
