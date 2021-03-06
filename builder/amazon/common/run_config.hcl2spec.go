// Code generated by "mapstructure-to-hcl2 -type AmiFilterOptions,SecurityGroupFilterOptions,SubnetFilterOptions,VpcFilterOptions"; DO NOT EDIT.
package common

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

// FlatAmiFilterOptions is an auto-generated flat version of AmiFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatAmiFilterOptions struct {
	Filters    map[string]string `cty:"filters"`
	Owners     []string          `cty:"owners"`
	MostRecent *bool             `mapstructure:"most_recent" cty:"most_recent"`
}

// FlatMapstructure returns a new FlatAmiFilterOptions.
// FlatAmiFilterOptions is an auto-generated flat version of AmiFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*AmiFilterOptions) FlatMapstructure() interface{} { return new(FlatAmiFilterOptions) }

// HCL2Spec returns the hcldec.Spec of a FlatAmiFilterOptions.
// This spec is used by HCL to read the fields of FlatAmiFilterOptions.
func (*FlatAmiFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters":     &hcldec.BlockAttrsSpec{TypeName: "filters", ElementType: cty.String, Required: false},
		"owners":      &hcldec.AttrSpec{Name: "owners", Type: cty.List(cty.String), Required: false},
		"most_recent": &hcldec.AttrSpec{Name: "most_recent", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatSecurityGroupFilterOptions is an auto-generated flat version of SecurityGroupFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSecurityGroupFilterOptions struct {
	Filters map[string]string `cty:"filters"`
}

// FlatMapstructure returns a new FlatSecurityGroupFilterOptions.
// FlatSecurityGroupFilterOptions is an auto-generated flat version of SecurityGroupFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SecurityGroupFilterOptions) FlatMapstructure() interface{} {
	return new(FlatSecurityGroupFilterOptions)
}

// HCL2Spec returns the hcldec.Spec of a FlatSecurityGroupFilterOptions.
// This spec is used by HCL to read the fields of FlatSecurityGroupFilterOptions.
func (*FlatSecurityGroupFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters": &hcldec.BlockAttrsSpec{TypeName: "filters", ElementType: cty.String, Required: false},
	}
	return s
}

// FlatSubnetFilterOptions is an auto-generated flat version of SubnetFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatSubnetFilterOptions struct {
	Filters  map[string]string `cty:"filters"`
	MostFree *bool             `mapstructure:"most_free" cty:"most_free"`
	Random   *bool             `mapstructure:"random" cty:"random"`
}

// FlatMapstructure returns a new FlatSubnetFilterOptions.
// FlatSubnetFilterOptions is an auto-generated flat version of SubnetFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*SubnetFilterOptions) FlatMapstructure() interface{} { return new(FlatSubnetFilterOptions) }

// HCL2Spec returns the hcldec.Spec of a FlatSubnetFilterOptions.
// This spec is used by HCL to read the fields of FlatSubnetFilterOptions.
func (*FlatSubnetFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters":   &hcldec.BlockAttrsSpec{TypeName: "filters", ElementType: cty.String, Required: false},
		"most_free": &hcldec.AttrSpec{Name: "most_free", Type: cty.Bool, Required: false},
		"random":    &hcldec.AttrSpec{Name: "random", Type: cty.Bool, Required: false},
	}
	return s
}

// FlatVpcFilterOptions is an auto-generated flat version of VpcFilterOptions.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatVpcFilterOptions struct {
	Filters map[string]string `cty:"filters"`
}

// FlatMapstructure returns a new FlatVpcFilterOptions.
// FlatVpcFilterOptions is an auto-generated flat version of VpcFilterOptions.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*VpcFilterOptions) FlatMapstructure() interface{} { return new(FlatVpcFilterOptions) }

// HCL2Spec returns the hcldec.Spec of a FlatVpcFilterOptions.
// This spec is used by HCL to read the fields of FlatVpcFilterOptions.
func (*FlatVpcFilterOptions) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"filters": &hcldec.BlockAttrsSpec{TypeName: "filters", ElementType: cty.String, Required: false},
	}
	return s
}
