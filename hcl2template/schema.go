package hcl2template

import "github.com/hashicorp/terraform/helper/schema"

type Schema schema.Schema

const (
	TypeInvalid = schema.TypeInvalid
	TypeBool    = schema.TypeBool
	TypeInt     = schema.TypeInt
	TypeFloat   = schema.TypeFloat
	TypeString  = schema.TypeString
	TypeList    = schema.TypeList
	TypeMap     = schema.TypeMap
	TypeSet     = schema.TypeSet
	TypeBlock
)
