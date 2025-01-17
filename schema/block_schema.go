package schema

import (
	"errors"
	"fmt"

	"github.com/hashicorp/go-multierror"
	"github.com/hashicorp/hcl-lang/lang"
)

// AttributeSchema describes schema for a block
// e.g. "resource" or "provider" in Terraform
type BlockSchema struct {
	Labels        []*LabelSchema
	Type          BlockType
	Body          *BodySchema
	DependentBody map[SchemaKey]*BodySchema

	Description  lang.MarkupContent
	IsDeprecated bool
	MinItems     uint64
	MaxItems     uint64

	Address *BlockAddrSchema
}

type BlockAddrSchema struct {
	Steps []AddrStep

	FriendlyName string
	ScopeId      lang.ScopeId

	// AsReference defines whether the block itself
	// is addressable as a type-less reference
	AsReference bool

	// BodyAsData defines whether the data in the block body
	// is addressable as cty.Object or cty.List(cty.Object),
	// cty.Set(cty.Object) etc. depending on block type
	BodyAsData bool
	// InferBody defines whether (static) Body's
	// blocks and attributes are also walked
	// and their addresses inferred as data
	InferBody bool

	// DependentBodyAsData defines whether the data in
	// the dependent block body is addressable as cty.Object
	// or cty.List(cty.Object), cty.Set(cty.Object) etc.
	// depending on block type
	DependentBodyAsData bool
	// InferDependentBody defines whether DependentBody's
	// blocks and attributes are also walked
	// and their addresses inferred as data
	InferDependentBody bool
}

func (bas *BlockAddrSchema) Validate() error {
	for i, step := range bas.Steps {
		if _, ok := step.(AttrNameStep); ok {
			return fmt.Errorf("Steps[%d]: AttrNameStep is not valid for attribute", i)
		}
	}

	if bas.InferBody && !bas.BodyAsData {
		return errors.New("InferBody requires BodyAsData")
	}

	if bas.InferDependentBody && !bas.DependentBodyAsData {
		return errors.New("InferDependentBody requires DependentBodyAsData")
	}

	return nil
}

func (bas *BlockAddrSchema) Copy() *BlockAddrSchema {
	if bas == nil {
		return nil
	}

	newBas := &BlockAddrSchema{
		FriendlyName:        bas.FriendlyName,
		ScopeId:             bas.ScopeId,
		AsReference:         bas.AsReference,
		BodyAsData:          bas.BodyAsData,
		InferBody:           bas.InferBody,
		DependentBodyAsData: bas.DependentBodyAsData,
		InferDependentBody:  bas.InferDependentBody,
	}

	newBas.Steps = make([]AddrStep, len(bas.Steps))
	for i, step := range bas.Steps {
		newBas.Steps[i] = step
	}

	return newBas
}

func (*BlockSchema) isSchemaImpl() schemaImplSigil {
	return schemaImplSigil{}
}

func (bSchema *BlockSchema) Validate() error {
	var errs *multierror.Error

	if bSchema.Address != nil {
		err := bSchema.Address.Validate()
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("Address: %w", err))
		}
	}

	if bSchema.Body != nil {
		err := bSchema.Body.Validate()
		if err != nil {
			errs = multierror.Append(errs, fmt.Errorf("Body: %w", err))
		}
	}

	if errs != nil && len(errs.Errors) == 1 {
		return errs.Errors[0]
	}

	return errs.ErrorOrNil()
}

func (bs *BlockSchema) Copy() *BlockSchema {
	if bs == nil {
		return nil
	}

	newBs := &BlockSchema{
		Type:         bs.Type,
		IsDeprecated: bs.IsDeprecated,
		MinItems:     bs.MinItems,
		MaxItems:     bs.MaxItems,
		Description:  bs.Description,
		Body:         bs.Body.Copy(),
		Address:      bs.Address.Copy(),
	}

	if bs.Labels != nil {
		newBs.Labels = make([]*LabelSchema, len(bs.Labels))
		for i, label := range bs.Labels {
			newBs.Labels[i] = label.Copy()
		}
	}

	if bs.DependentBody != nil {
		newBs.DependentBody = make(map[SchemaKey]*BodySchema, 0)
		for key, depSchema := range bs.DependentBody {
			newBs.DependentBody[key] = depSchema.Copy()
		}
	}

	return newBs
}
