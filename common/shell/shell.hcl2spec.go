// Code generated by "hcl2-schema"; DO NOT EDIT.\n

package shell

import (
	"github.com/hashicorp/hcl2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

func (*Provisioner) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"ExecuteCommand": &hcldec.AttrSpec{Name: "execute_command", Type: cty.String, Required: false},
		"RemotePath":     &hcldec.AttrSpec{Name: "remote_path", Type: cty.String, Required: false},
		"ValidExitCodes": nil, /* TODO */
		"Vars":           &hcldec.AttrSpec{Name: "environment_vars", Type: cty.List(cty.String), Required: false},
	}
	for k, v := range (*Provisioner)(nil).PackerConfig.HCL2Spec() {
		s[k] = v
	}
	return s
}
