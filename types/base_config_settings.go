package types

import (
	"time"
)

type BaseConfigSettings struct {
	LastUpdate               time.Time `xmlrpc:"__last_update"`
	AliasDomain              string    `xmlrpc:"alias_domain"`
	AuthSignupResetPassword  bool      `xmlrpc:"auth_signup_reset_password"`
	AuthSignupTemplateUserId Many2One  `xmlrpc:"auth_signup_template_user_id"`
	AuthSignupUninvited      bool      `xmlrpc:"auth_signup_uninvited"`
	CompanyId                Many2One  `xmlrpc:"company_id"`
	CompanySharePartner      bool      `xmlrpc:"company_share_partner"`
	CompanyShareProduct      bool      `xmlrpc:"company_share_product"`
	CreateDate               time.Time `xmlrpc:"create_date"`
	CreateUid                Many2One  `xmlrpc:"create_uid"`
	CustomFooter             bool      `xmlrpc:"custom_footer"`
	DisplayName              string    `xmlrpc:"display_name"`
	FailCounter              int64     `xmlrpc:"fail_counter"`
	Font                     Many2One  `xmlrpc:"font"`
	GroupMultiCompany        bool      `xmlrpc:"group_multi_company"`
	GroupMultiCurrency       bool      `xmlrpc:"group_multi_currency"`
	GroupProductVariant      string    `xmlrpc:"group_product_variant"`
	Id                       int64     `xmlrpc:"id"`
	Ldaps                    []int64   `xmlrpc:"ldaps"`
	ModuleAuthOauth          bool      `xmlrpc:"module_auth_oauth"`
	ModuleBaseImport         bool      `xmlrpc:"module_base_import"`
	ModuleGoogleCalendar     bool      `xmlrpc:"module_google_calendar"`
	ModuleGoogleDrive        bool      `xmlrpc:"module_google_drive"`
	ModuleInterCompanyRules  bool      `xmlrpc:"module_inter_company_rules"`
	ModulePortal             bool      `xmlrpc:"module_portal"`
	ModuleShare              bool      `xmlrpc:"module_share"`
	PaperformatId            Many2One  `xmlrpc:"paperformat_id"`
	RmlFooter                string    `xmlrpc:"rml_footer"`
	RmlFooterReadonly        string    `xmlrpc:"rml_footer_readonly"`
	RmlHeader                string    `xmlrpc:"rml_header"`
	RmlHeader2               string    `xmlrpc:"rml_header2"`
	RmlHeader3               string    `xmlrpc:"rml_header3"`
	RmlPaperFormat           string    `xmlrpc:"rml_paper_format"`
	WriteDate                time.Time `xmlrpc:"write_date"`
	WriteUid                 Many2One  `xmlrpc:"write_uid"`
}

type BaseConfigSettingsNil struct {
	LastUpdate               interface{} `xmlrpc:"__last_update"`
	AliasDomain              interface{} `xmlrpc:"alias_domain"`
	AuthSignupResetPassword  bool        `xmlrpc:"auth_signup_reset_password"`
	AuthSignupTemplateUserId interface{} `xmlrpc:"auth_signup_template_user_id"`
	AuthSignupUninvited      bool        `xmlrpc:"auth_signup_uninvited"`
	CompanyId                interface{} `xmlrpc:"company_id"`
	CompanySharePartner      bool        `xmlrpc:"company_share_partner"`
	CompanyShareProduct      bool        `xmlrpc:"company_share_product"`
	CreateDate               interface{} `xmlrpc:"create_date"`
	CreateUid                interface{} `xmlrpc:"create_uid"`
	CustomFooter             bool        `xmlrpc:"custom_footer"`
	DisplayName              interface{} `xmlrpc:"display_name"`
	FailCounter              interface{} `xmlrpc:"fail_counter"`
	Font                     interface{} `xmlrpc:"font"`
	GroupMultiCompany        bool        `xmlrpc:"group_multi_company"`
	GroupMultiCurrency       bool        `xmlrpc:"group_multi_currency"`
	GroupProductVariant      interface{} `xmlrpc:"group_product_variant"`
	Id                       interface{} `xmlrpc:"id"`
	Ldaps                    interface{} `xmlrpc:"ldaps"`
	ModuleAuthOauth          bool        `xmlrpc:"module_auth_oauth"`
	ModuleBaseImport         bool        `xmlrpc:"module_base_import"`
	ModuleGoogleCalendar     bool        `xmlrpc:"module_google_calendar"`
	ModuleGoogleDrive        bool        `xmlrpc:"module_google_drive"`
	ModuleInterCompanyRules  bool        `xmlrpc:"module_inter_company_rules"`
	ModulePortal             bool        `xmlrpc:"module_portal"`
	ModuleShare              bool        `xmlrpc:"module_share"`
	PaperformatId            interface{} `xmlrpc:"paperformat_id"`
	RmlFooter                interface{} `xmlrpc:"rml_footer"`
	RmlFooterReadonly        interface{} `xmlrpc:"rml_footer_readonly"`
	RmlHeader                interface{} `xmlrpc:"rml_header"`
	RmlHeader2               interface{} `xmlrpc:"rml_header2"`
	RmlHeader3               interface{} `xmlrpc:"rml_header3"`
	RmlPaperFormat           interface{} `xmlrpc:"rml_paper_format"`
	WriteDate                interface{} `xmlrpc:"write_date"`
	WriteUid                 interface{} `xmlrpc:"write_uid"`
}

var BaseConfigSettingsModel string = "base.config.settings"

type BaseConfigSettingss []BaseConfigSettings

type BaseConfigSettingssNil []BaseConfigSettingsNil

func (s *BaseConfigSettings) NilableType_() interface{} {
	return &BaseConfigSettingsNil{}
}

func (ns *BaseConfigSettingsNil) Type_() interface{} {
	s := &BaseConfigSettings{}
	return load(ns, s)
}

func (s *BaseConfigSettingss) NilableType_() interface{} {
	return &BaseConfigSettingssNil{}
}

func (ns *BaseConfigSettingssNil) Type_() interface{} {
	s := &BaseConfigSettingss{}
	for _, nsi := range *ns {
		*s = append(*s, *nsi.Type_().(*BaseConfigSettings))
	}
	return s
}
