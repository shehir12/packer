package hcl2template

import (
	"github.com/hashicorp/hcl2/hcl"
	"github.com/hashicorp/packer/template"
)

const (
	buildFromLabel = "from"

	buildProvisionnersLabel = "provision"

	buildPostProvisionnersLabel = "post_provision"
)

var buildSchema = &hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{
		{Type: buildFromLabel, LabelNames: []string{"src"}},
		{Type: buildProvisionnersLabel},
		{Type: buildPostProvisionnersLabel},
	},
}

type Build struct {
	// Ordered list of provisioner groups
	ProvisionerGroups ProvisionerGroups

	// Ordered list of post-provisioner groups
	PostProvisionerGroups ProvisionerGroups

	// Ordered list of output stanzas
	Froms BuildFromList

	HCL2Ref HCL2Ref
}

type Builds []*Build

func (p *Parser) decodeBuildConfig(block *hcl.Block) (*Build, hcl.Diagnostics) {
	build := &Build{}

	content, diags := block.Body.Content(buildSchema)
	for _, block := range content.Blocks {
		switch block.Type {
		case buildFromLabel:
			bf := BuildFrom{}
			moreDiags := bf.decodeConfig(block)
			diags = append(diags, moreDiags...)
			build.Froms = append(build.Froms, bf)
		case buildProvisionnersLabel:
			pg, moreDiags := p.decodeProvisionerGroup(block)
			diags = append(diags, moreDiags...)
			build.ProvisionerGroups = append(build.ProvisionerGroups, pg)
		case buildPostProvisionnersLabel:
			pg, moreDiags := p.decodePostProvisionerGroup(block)
			diags = append(diags, moreDiags...)
			build.PostProvisionerGroups = append(build.PostProvisionerGroups, pg)
		}
	}

	return build, diags
}

type PackerV1Build struct {
	Builders       []*template.Builder
	Provisioners   []*template.Provisioner
	PostProcessors []*template.PostProcessor
}

func (builds Builds) ToV1Build() PackerV1Build {
	res := PackerV1Build{}

	for _, build := range builds {
		for _, from := range build.Froms {
			for _, provisionerGroup := range build.ProvisionerGroups {
				for _, provisioner := range provisionerGroup {
				}
			}
		}
	}
	return res
}
