// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// CoreV2PipelineExtensionOverridesFieldResolvers represents a collection of methods whose products represent the
// response values of the 'CoreV2PipelineExtensionOverrides' type.
type CoreV2PipelineExtensionOverridesFieldResolvers interface {
	// ID implements response to request for 'id' field.
	ID(p graphql.ResolveParams) (string, error)

	// ToJSON implements response to request for 'toJSON' field.
	ToJSON(p graphql.ResolveParams) (interface{}, error)
}

// RegisterCoreV2PipelineExtensionOverrides registers CoreV2PipelineExtensionOverrides object type with given service.
func RegisterCoreV2PipelineExtensionOverrides(svc *graphql.Service, impl CoreV2PipelineExtensionOverridesFieldResolvers) {
	svc.RegisterObjectExtension(_ObjectExtensionTypeCoreV2PipelineExtensionOverridesDesc, impl)
}

func _ObjTypeCoreV2PipelineExtensionOverridesIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		ID(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ID(frp)
	}
}

func _ObjTypeCoreV2PipelineExtensionOverridesToJSONHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		ToJSON(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ToJSON(frp)
	}
}

func _ObjectExtensionTypeCoreV2PipelineExtensionOverridesConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "",
		Fields: graphql1.Fields{
			"id": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Unique global identifier used to reference resource.",
				Name:              "id",
				Type:              graphql1.NewNonNull(graphql1.ID),
			},
			"toJSON": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "toJSON returns a REST API compatible representation of the resource. Handy for\nsharing snippets that can then be imported with `sensuctl create`.",
				Name:              "toJSON",
				Type:              graphql1.NewNonNull(graphql.OutputType("JSON")),
			},
		},
		Interfaces: []*graphql1.Interface{
			graphql.Interface("Node"),
			graphql.Interface("Resource")},
		Name: "CoreV2Pipeline",
	}
}

// describe CoreV2PipelineExtensionOverrides's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectExtensionTypeCoreV2PipelineExtensionOverridesDesc = graphql.ObjectDesc{
	Config: _ObjectExtensionTypeCoreV2PipelineExtensionOverridesConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"id":     _ObjTypeCoreV2PipelineExtensionOverridesIDHandler,
		"toJSON": _ObjTypeCoreV2PipelineExtensionOverridesToJSONHandler,
	},
}
