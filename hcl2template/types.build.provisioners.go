package hcl2template

import (
	"github.com/hashicorp/hcl2/gohcl"
	"github.com/hashicorp/hcl2/hcl"
)

type ProvisionerGroup struct {
	CommunicatorRef CommunicatorRef

	Provisioners []Provisioner
	HCL2Ref      HCL2Ref
}

type Provisioner struct {
	*hcl.Block
}

var provisionerGroupSchema = hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{},
	Attributes: []hcl.AttributeSchema{
		{"communicator", false},
	},
}

type ProvisionerGroups []*ProvisionerGroup

func (pgs ProvisionerGroups) FirstCommunicatorRef() CommunicatorRef {
	if len(pgs) == 0 {
		return NoCommunicator
	}
	return pgs[0].CommunicatorRef
}

func (p *Parser) decodeProvisionerGroup(block *hcl.Block) (*ProvisionerGroup, hcl.Diagnostics) {

	var b struct {
		Communicator string   `hcl:"communicator"`
		Remain       hcl.Body `hcl:",remain"`
	}

	diags := gohcl.DecodeBody(block.Body, nil, &b)

	pg := &ProvisionerGroup{}
	pg.CommunicatorRef = communicatorRefFromString(b.Communicator)
	pg.HCL2Ref.DeclRange = block.DefRange
	pg.HCL2Ref.Remain = b.Remain

	s := provisionerGroupSchema
	s.Attributes = append(s.Attributes, p.ProvisionersSchema.Attributes...)
	s.Blocks = append(s.Blocks, p.ProvisionersSchema.Blocks...)

	content, moreDiags := pg.HCL2Ref.Remain.Content(&s)
	diags = append(diags, moreDiags...)

	for _, block := range content.Blocks {
		p := Provisioner{block}
		pg.Provisioners = append(pg.Provisioners, p)
	}

	return pg, diags
}

var postProvisionerGroupSchema = hcl.BodySchema{
	Blocks: []hcl.BlockHeaderSchema{},
	Attributes: []hcl.AttributeSchema{
		{"communicator", false},
	},
}

func (p *Parser) decodePostProvisionerGroup(block *hcl.Block) (*ProvisionerGroup, hcl.Diagnostics) {

	var b struct {
		Communicator string
		Remain       hcl.Body `hcl:",remain"`
	}

	diags := gohcl.DecodeBody(block.Body, nil, &b)

	pg := &ProvisionerGroup{}
	pg.CommunicatorRef = communicatorRefFromString(b.Communicator)
	pg.HCL2Ref.DeclRange = block.DefRange
	pg.HCL2Ref.Remain = b.Remain

	s := postProvisionerGroupSchema
	s.Attributes = append(s.Attributes, p.PostProvisionersSchema.Attributes...)
	s.Blocks = append(s.Blocks, p.PostProvisionersSchema.Blocks...)

	content, moreDiags := pg.HCL2Ref.Remain.Content(&s)
	diags = append(diags, moreDiags...)

	for _, block := range content.Blocks {
		p := Provisioner{block}
		pg.Provisioners = append(pg.Provisioners, p)
	}

	return pg, diags
}

// func (pgs ProvisionerGroups) FlatProvisioners() []Provisioner {
// 	res := []Provisioner{}
// 	for _, provisionerGroup := range pgs {
// 		p := *provisionerGroup
// 		for _, provisioner := range p {
// 			res = append(provisioner, provisioner)
// 		}
// 	}
// 	return res
// }
