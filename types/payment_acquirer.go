package types

import (
	"time"
)

type PaymentAcquirer struct {
	AuthorizeImplemented       bool      `xmlrpc:"authorize_implemented"`
	CancelMsg                  string    `xmlrpc:"cancel_msg"`
	CaptureManually            bool      `xmlrpc:"capture_manually"`
	CompanyId                  Many2One  `xmlrpc:"company_id"`
	CountryIds                 []int64   `xmlrpc:"country_ids"`
	CreateDate                 time.Time `xmlrpc:"create_date"`
	CreateUid                  Many2One  `xmlrpc:"create_uid"`
	Description                string    `xmlrpc:"description"`
	DisplayName                string    `xmlrpc:"display_name"`
	DoneMsg                    string    `xmlrpc:"done_msg"`
	Environment                string    `xmlrpc:"environment"`
	ErrorMsg                   string    `xmlrpc:"error_msg"`
	FeesActive                 bool      `xmlrpc:"fees_active"`
	FeesDomFixed               float64   `xmlrpc:"fees_dom_fixed"`
	FeesDomVar                 float64   `xmlrpc:"fees_dom_var"`
	FeesImplemented            bool      `xmlrpc:"fees_implemented"`
	FeesIntFixed               float64   `xmlrpc:"fees_int_fixed"`
	FeesIntVar                 float64   `xmlrpc:"fees_int_var"`
	Id                         int64     `xmlrpc:"id"`
	Image                      string    `xmlrpc:"image"`
	ImageMedium                string    `xmlrpc:"image_medium"`
	ImageSmall                 string    `xmlrpc:"image_small"`
	JournalId                  Many2One  `xmlrpc:"journal_id"`
	LastUpdate                 time.Time `xmlrpc:"__last_update"`
	ModuleId                   Many2One  `xmlrpc:"module_id"`
	ModuleState                string    `xmlrpc:"module_state"`
	Name                       string    `xmlrpc:"name"`
	PaymentFlow                string    `xmlrpc:"payment_flow"`
	PaymentIconIds             []int64   `xmlrpc:"payment_icon_ids"`
	PendingMsg                 string    `xmlrpc:"pending_msg"`
	PostMsg                    string    `xmlrpc:"post_msg"`
	PreMsg                     string    `xmlrpc:"pre_msg"`
	Provider                   string    `xmlrpc:"provider"`
	RegistrationViewTemplateId Many2One  `xmlrpc:"registration_view_template_id"`
	SaveToken                  string    `xmlrpc:"save_token"`
	Sequence                   int64     `xmlrpc:"sequence"`
	SpecificCountries          bool      `xmlrpc:"specific_countries"`
	TokenImplemented           bool      `xmlrpc:"token_implemented"`
	ViewTemplateId             Many2One  `xmlrpc:"view_template_id"`
	WebsitePublished           bool      `xmlrpc:"website_published"`
	WriteDate                  time.Time `xmlrpc:"write_date"`
	WriteUid                   Many2One  `xmlrpc:"write_uid"`
}

type PaymentAcquirerNil struct {
	AuthorizeImplemented       bool        `xmlrpc:"authorize_implemented"`
	CancelMsg                  interface{} `xmlrpc:"cancel_msg"`
	CaptureManually            bool        `xmlrpc:"capture_manually"`
	CompanyId                  interface{} `xmlrpc:"company_id"`
	CountryIds                 interface{} `xmlrpc:"country_ids"`
	CreateDate                 interface{} `xmlrpc:"create_date"`
	CreateUid                  interface{} `xmlrpc:"create_uid"`
	Description                interface{} `xmlrpc:"description"`
	DisplayName                interface{} `xmlrpc:"display_name"`
	DoneMsg                    interface{} `xmlrpc:"done_msg"`
	Environment                interface{} `xmlrpc:"environment"`
	ErrorMsg                   interface{} `xmlrpc:"error_msg"`
	FeesActive                 bool        `xmlrpc:"fees_active"`
	FeesDomFixed               interface{} `xmlrpc:"fees_dom_fixed"`
	FeesDomVar                 interface{} `xmlrpc:"fees_dom_var"`
	FeesImplemented            bool        `xmlrpc:"fees_implemented"`
	FeesIntFixed               interface{} `xmlrpc:"fees_int_fixed"`
	FeesIntVar                 interface{} `xmlrpc:"fees_int_var"`
	Id                         interface{} `xmlrpc:"id"`
	Image                      interface{} `xmlrpc:"image"`
	ImageMedium                interface{} `xmlrpc:"image_medium"`
	ImageSmall                 interface{} `xmlrpc:"image_small"`
	JournalId                  interface{} `xmlrpc:"journal_id"`
	LastUpdate                 interface{} `xmlrpc:"__last_update"`
	ModuleId                   interface{} `xmlrpc:"module_id"`
	ModuleState                interface{} `xmlrpc:"module_state"`
	Name                       interface{} `xmlrpc:"name"`
	PaymentFlow                interface{} `xmlrpc:"payment_flow"`
	PaymentIconIds             interface{} `xmlrpc:"payment_icon_ids"`
	PendingMsg                 interface{} `xmlrpc:"pending_msg"`
	PostMsg                    interface{} `xmlrpc:"post_msg"`
	PreMsg                     interface{} `xmlrpc:"pre_msg"`
	Provider                   interface{} `xmlrpc:"provider"`
	RegistrationViewTemplateId interface{} `xmlrpc:"registration_view_template_id"`
	SaveToken                  interface{} `xmlrpc:"save_token"`
	Sequence                   interface{} `xmlrpc:"sequence"`
	SpecificCountries          bool        `xmlrpc:"specific_countries"`
	TokenImplemented           bool        `xmlrpc:"token_implemented"`
	ViewTemplateId             interface{} `xmlrpc:"view_template_id"`
	WebsitePublished           bool        `xmlrpc:"website_published"`
	WriteDate                  interface{} `xmlrpc:"write_date"`
	WriteUid                   interface{} `xmlrpc:"write_uid"`
}

var PaymentAcquirerModel string = "payment.acquirer"

type PaymentAcquirers []PaymentAcquirer

type PaymentAcquirersNil []PaymentAcquirerNil

func (s *PaymentAcquirer) NilableType_() interface{} {
	return &PaymentAcquirerNil{}
}

func (ns *PaymentAcquirerNil) Type_() interface{} {
	s := &PaymentAcquirer{}
	return load(ns, s)
}

func (s *PaymentAcquirers) NilableType_() interface{} {
	return &PaymentAcquirersNil{}
}

func (ns *PaymentAcquirersNil) Type_() interface{} {
	s := &PaymentAcquirers{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*PaymentAcquirer))
	}
	return s
}
