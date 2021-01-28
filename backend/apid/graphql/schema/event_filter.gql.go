// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	errors "errors"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

//
// EventFilterFieldResolvers represents a collection of methods whose products represent the
// response values of the 'EventFilter' type.
type EventFilterFieldResolvers interface {
	// ID implements response to request for 'id' field.
	ID(p graphql.ResolveParams) (string, error)

	// Namespace implements response to request for 'namespace' field.
	Namespace(p graphql.ResolveParams) (string, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)

	// Metadata implements response to request for 'metadata' field.
	Metadata(p graphql.ResolveParams) (interface{}, error)

	// Action implements response to request for 'action' field.
	Action(p graphql.ResolveParams) (EventFilterAction, error)

	// Expressions implements response to request for 'expressions' field.
	Expressions(p graphql.ResolveParams) ([]string, error)

	// RuntimeAssets implements response to request for 'runtimeAssets' field.
	RuntimeAssets(p graphql.ResolveParams) (interface{}, error)

	// ToJSON implements response to request for 'toJSON' field.
	ToJSON(p graphql.ResolveParams) (interface{}, error)
}

// EventFilterAliases implements all methods on EventFilterFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type EventFilterAliases struct{}

// ID implements response to request for 'id' field.
func (_ EventFilterAliases) ID(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'id'")
	}
	return ret, err
}

// Namespace implements response to request for 'namespace' field.
func (_ EventFilterAliases) Namespace(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'namespace'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ EventFilterAliases) Name(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'name'")
	}
	return ret, err
}

// Metadata implements response to request for 'metadata' field.
func (_ EventFilterAliases) Metadata(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Action implements response to request for 'action' field.
func (_ EventFilterAliases) Action(p graphql.ResolveParams) (EventFilterAction, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := EventFilterAction(val.(string)), true
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'action'")
	}
	return ret, err
}

// Expressions implements response to request for 'expressions' field.
func (_ EventFilterAliases) Expressions(p graphql.ResolveParams) ([]string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.([]string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'expressions'")
	}
	return ret, err
}

// RuntimeAssets implements response to request for 'runtimeAssets' field.
func (_ EventFilterAliases) RuntimeAssets(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// ToJSON implements response to request for 'toJSON' field.
func (_ EventFilterAliases) ToJSON(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// EventFilterType A Filter is a filter specification.
var EventFilterType = graphql.NewType("EventFilter", graphql.ObjectKind)

// RegisterEventFilter registers EventFilter object type with given service.
func RegisterEventFilter(svc *graphql.Service, impl EventFilterFieldResolvers) {
	svc.RegisterObject(_ObjectTypeEventFilterDesc, impl)
}
func _ObjTypeEventFilterIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		ID(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ID(frp)
	}
}

func _ObjTypeEventFilterNamespaceHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Namespace(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Namespace(frp)
	}
}

func _ObjTypeEventFilterNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Name(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjTypeEventFilterMetadataHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Metadata(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Metadata(frp)
	}
}

func _ObjTypeEventFilterActionHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Action(p graphql.ResolveParams) (EventFilterAction, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {

		val, err := resolver.Action(frp)
		return string(val), err
	}
}

func _ObjTypeEventFilterExpressionsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Expressions(p graphql.ResolveParams) ([]string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Expressions(frp)
	}
}

func _ObjTypeEventFilterRuntimeAssetsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		RuntimeAssets(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.RuntimeAssets(frp)
	}
}

func _ObjTypeEventFilterToJSONHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		ToJSON(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.ToJSON(frp)
	}
}

func _ObjectTypeEventFilterConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "A Filter is a filter specification.",
		Fields: graphql1.Fields{
			"action": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Command is the command to be executed.",
				Name:              "action",
				Type:              graphql1.NewNonNull(graphql.OutputType("EventFilterAction")),
			},
			"expressions": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Env is a list of environment variables to use with command execution",
				Name:              "expressions",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql1.String))),
			},
			"id": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "The globally unique identifier of the record",
				Name:              "id",
				Type:              graphql1.NewNonNull(graphql1.ID),
			},
			"metadata": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "metadata contains name, namespace, labels and annotations of the record",
				Name:              "metadata",
				Type:              graphql.OutputType("ObjectMeta"),
			},
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "use metadata",
				Description:       "Name is the unique identifier for an event filter.",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"namespace": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "use metadata",
				Description:       "Namespace in which this record resides",
				Name:              "namespace",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"runtimeAssets": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "RuntimeAssets are a list of assets required to execute the event filter",
				Name:              "runtimeAssets",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("Asset")))),
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
			graphql.Interface("Namespaced"),
			graphql.Interface("Resource")},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see EventFilterFieldResolvers.")
		},
		Name: "EventFilter",
	}
}

// describe EventFilter's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeEventFilterDesc = graphql.ObjectDesc{
	Config: _ObjectTypeEventFilterConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"action":        _ObjTypeEventFilterActionHandler,
		"expressions":   _ObjTypeEventFilterExpressionsHandler,
		"id":            _ObjTypeEventFilterIDHandler,
		"metadata":      _ObjTypeEventFilterMetadataHandler,
		"name":          _ObjTypeEventFilterNameHandler,
		"namespace":     _ObjTypeEventFilterNamespaceHandler,
		"runtimeAssets": _ObjTypeEventFilterRuntimeAssetsHandler,
		"toJSON":        _ObjTypeEventFilterToJSONHandler,
	},
}

// EventFilterAction self descriptive
type EventFilterAction string

// EventFilterActions holds enum values
var EventFilterActions = _EnumTypeEventFilterActionValues{
	ALLOW: "ALLOW",
	DENY:  "DENY",
}

// EventFilterActionType self descriptive
var EventFilterActionType = graphql.NewType("EventFilterAction", graphql.EnumKind)

// RegisterEventFilterAction registers EventFilterAction object type with given service.
func RegisterEventFilterAction(svc *graphql.Service) {
	svc.RegisterEnum(_EnumTypeEventFilterActionDesc)
}
func _EnumTypeEventFilterActionConfigFn() graphql1.EnumConfig {
	return graphql1.EnumConfig{
		Description: "self descriptive",
		Name:        "EventFilterAction",
		Values: graphql1.EnumValueConfigMap{
			"ALLOW": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "ALLOW",
			},
			"DENY": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "DENY",
			},
		},
	}
}

// describe EventFilterAction's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _EnumTypeEventFilterActionDesc = graphql.EnumDesc{Config: _EnumTypeEventFilterActionConfigFn}

type _EnumTypeEventFilterActionValues struct {
	// ALLOW - self descriptive
	ALLOW EventFilterAction
	// DENY - self descriptive
	DENY EventFilterAction
}

//
// EventFilterConnectionFieldResolvers represents a collection of methods whose products represent the
// response values of the 'EventFilterConnection' type.
type EventFilterConnectionFieldResolvers interface {
	// Nodes implements response to request for 'nodes' field.
	Nodes(p graphql.ResolveParams) (interface{}, error)

	// PageInfo implements response to request for 'pageInfo' field.
	PageInfo(p graphql.ResolveParams) (interface{}, error)
}

// EventFilterConnectionAliases implements all methods on EventFilterConnectionFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type EventFilterConnectionAliases struct{}

// Nodes implements response to request for 'nodes' field.
func (_ EventFilterConnectionAliases) Nodes(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// PageInfo implements response to request for 'pageInfo' field.
func (_ EventFilterConnectionAliases) PageInfo(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// EventFilterConnectionType A connection to a sequence of records.
var EventFilterConnectionType = graphql.NewType("EventFilterConnection", graphql.ObjectKind)

// RegisterEventFilterConnection registers EventFilterConnection object type with given service.
func RegisterEventFilterConnection(svc *graphql.Service, impl EventFilterConnectionFieldResolvers) {
	svc.RegisterObject(_ObjectTypeEventFilterConnectionDesc, impl)
}
func _ObjTypeEventFilterConnectionNodesHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Nodes(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Nodes(frp)
	}
}

func _ObjTypeEventFilterConnectionPageInfoHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		PageInfo(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.PageInfo(frp)
	}
}

func _ObjectTypeEventFilterConnectionConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "A connection to a sequence of records.",
		Fields: graphql1.Fields{
			"nodes": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "nodes",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("EventFilter")))),
			},
			"pageInfo": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "pageInfo",
				Type:              graphql1.NewNonNull(graphql.OutputType("OffsetPageInfo")),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see EventFilterConnectionFieldResolvers.")
		},
		Name: "EventFilterConnection",
	}
}

// describe EventFilterConnection's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeEventFilterConnectionDesc = graphql.ObjectDesc{
	Config: _ObjectTypeEventFilterConnectionConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"nodes":    _ObjTypeEventFilterConnectionNodesHandler,
		"pageInfo": _ObjTypeEventFilterConnectionPageInfoHandler,
	},
}

// EventFilterListOrder Describes ways in which a list of mutators can be ordered.
type EventFilterListOrder string

// EventFilterListOrders holds enum values
var EventFilterListOrders = _EnumTypeEventFilterListOrderValues{
	NAME:      "NAME",
	NAME_DESC: "NAME_DESC",
}

// EventFilterListOrderType Describes ways in which a list of mutators can be ordered.
var EventFilterListOrderType = graphql.NewType("EventFilterListOrder", graphql.EnumKind)

// RegisterEventFilterListOrder registers EventFilterListOrder object type with given service.
func RegisterEventFilterListOrder(svc *graphql.Service) {
	svc.RegisterEnum(_EnumTypeEventFilterListOrderDesc)
}
func _EnumTypeEventFilterListOrderConfigFn() graphql1.EnumConfig {
	return graphql1.EnumConfig{
		Description: "Describes ways in which a list of mutators can be ordered.",
		Name:        "EventFilterListOrder",
		Values: graphql1.EnumValueConfigMap{
			"NAME": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "NAME",
			},
			"NAME_DESC": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "NAME_DESC",
			},
		},
	}
}

// describe EventFilterListOrder's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _EnumTypeEventFilterListOrderDesc = graphql.EnumDesc{Config: _EnumTypeEventFilterListOrderConfigFn}

type _EnumTypeEventFilterListOrderValues struct {
	// NAME - self descriptive
	NAME EventFilterListOrder
	// NAME_DESC - self descriptive
	NAME_DESC EventFilterListOrder
}
