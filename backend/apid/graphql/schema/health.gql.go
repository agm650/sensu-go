// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	errors "errors"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

//
// EtcdClusterMemberHealthFieldResolvers represents a collection of methods whose products represent the
// response values of the 'EtcdClusterMemberHealth' type.
type EtcdClusterMemberHealthFieldResolvers interface {
	// MemberID implements response to request for 'memberID' field.
	MemberID(p graphql.ResolveParams) (string, error)

	// Name implements response to request for 'name' field.
	Name(p graphql.ResolveParams) (string, error)

	// Err implements response to request for 'err' field.
	Err(p graphql.ResolveParams) (string, error)

	// Healthy implements response to request for 'healthy' field.
	Healthy(p graphql.ResolveParams) (bool, error)
}

// EtcdClusterMemberHealthAliases implements all methods on EtcdClusterMemberHealthFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type EtcdClusterMemberHealthAliases struct{}

// MemberID implements response to request for 'memberID' field.
func (_ EtcdClusterMemberHealthAliases) MemberID(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'memberID'")
	}
	return ret, err
}

// Name implements response to request for 'name' field.
func (_ EtcdClusterMemberHealthAliases) Name(p graphql.ResolveParams) (string, error) {
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

// Err implements response to request for 'err' field.
func (_ EtcdClusterMemberHealthAliases) Err(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'err'")
	}
	return ret, err
}

// Healthy implements response to request for 'healthy' field.
func (_ EtcdClusterMemberHealthAliases) Healthy(p graphql.ResolveParams) (bool, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(bool)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'healthy'")
	}
	return ret, err
}

// EtcdClusterMemberHealthType Describes the health of an Etcd cluster member
var EtcdClusterMemberHealthType = graphql.NewType("EtcdClusterMemberHealth", graphql.ObjectKind)

// RegisterEtcdClusterMemberHealth registers EtcdClusterMemberHealth object type with given service.
func RegisterEtcdClusterMemberHealth(svc *graphql.Service, impl EtcdClusterMemberHealthFieldResolvers) {
	svc.RegisterObject(_ObjectTypeEtcdClusterMemberHealthDesc, impl)
}
func _ObjTypeEtcdClusterMemberHealthMemberIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		MemberID(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.MemberID(frp)
	}
}

func _ObjTypeEtcdClusterMemberHealthNameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Name(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Name(frp)
	}
}

func _ObjTypeEtcdClusterMemberHealthErrHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Err(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Err(frp)
	}
}

func _ObjTypeEtcdClusterMemberHealthHealthyHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Healthy(p graphql.ResolveParams) (bool, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Healthy(frp)
	}
}

func _ObjectTypeEtcdClusterMemberHealthConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Describes the health of an Etcd cluster member",
		Fields: graphql1.Fields{
			"err": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Err holds the string representation of any errors encountered while checking the member's health.",
				Name:              "err",
				Type:              graphql1.String,
			},
			"healthy": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Healthy describes the health of the cluster member.",
				Name:              "healthy",
				Type:              graphql1.NewNonNull(graphql1.Boolean),
			},
			"memberID": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "MemberID is the etcd cluster member's ID.",
				Name:              "memberID",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
			"name": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Name is the cluster member's name.",
				Name:              "name",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see EtcdClusterMemberHealthFieldResolvers.")
		},
		Name: "EtcdClusterMemberHealth",
	}
}

// describe EtcdClusterMemberHealth's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeEtcdClusterMemberHealthDesc = graphql.ObjectDesc{
	Config: _ObjectTypeEtcdClusterMemberHealthConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"err":      _ObjTypeEtcdClusterMemberHealthErrHandler,
		"healthy":  _ObjTypeEtcdClusterMemberHealthHealthyHandler,
		"memberID": _ObjTypeEtcdClusterMemberHealthMemberIDHandler,
		"name":     _ObjTypeEtcdClusterMemberHealthNameHandler,
	},
}

//
// EtcdClusterHealthFieldResolvers represents a collection of methods whose products represent the
// response values of the 'EtcdClusterHealth' type.
type EtcdClusterHealthFieldResolvers interface {
	// Alarms implements response to request for 'alarms' field.
	Alarms(p graphql.ResolveParams) (interface{}, error)

	// Members implements response to request for 'members' field.
	Members(p graphql.ResolveParams) (interface{}, error)
}

// EtcdClusterHealthAliases implements all methods on EtcdClusterHealthFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type EtcdClusterHealthAliases struct{}

// Alarms implements response to request for 'alarms' field.
func (_ EtcdClusterHealthAliases) Alarms(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Members implements response to request for 'members' field.
func (_ EtcdClusterHealthAliases) Members(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// EtcdClusterHealthType Describes the health of an Etcd cluster
var EtcdClusterHealthType = graphql.NewType("EtcdClusterHealth", graphql.ObjectKind)

// RegisterEtcdClusterHealth registers EtcdClusterHealth object type with given service.
func RegisterEtcdClusterHealth(svc *graphql.Service, impl EtcdClusterHealthFieldResolvers) {
	svc.RegisterObject(_ObjectTypeEtcdClusterHealthDesc, impl)
}
func _ObjTypeEtcdClusterHealthAlarmsHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Alarms(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Alarms(frp)
	}
}

func _ObjTypeEtcdClusterHealthMembersHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Members(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Members(frp)
	}
}

func _ObjectTypeEtcdClusterHealthConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Describes the health of an Etcd cluster",
		Fields: graphql1.Fields{
			"alarms": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Alarms is the list of active etcd alarms.",
				Name:              "alarms",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("EtcdAlarmMember")))),
			},
			"members": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "Returns list of health status for every cluster member.",
				Name:              "members",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("EtcdClusterMemberHealth")))),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see EtcdClusterHealthFieldResolvers.")
		},
		Name: "EtcdClusterHealth",
	}
}

// describe EtcdClusterHealth's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeEtcdClusterHealthDesc = graphql.ObjectDesc{
	Config: _ObjectTypeEtcdClusterHealthConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"alarms":  _ObjTypeEtcdClusterHealthAlarmsHandler,
		"members": _ObjTypeEtcdClusterHealthMembersHandler,
	},
}

//
// EtcdAlarmMemberFieldResolvers represents a collection of methods whose products represent the
// response values of the 'EtcdAlarmMember' type.
type EtcdAlarmMemberFieldResolvers interface {
	// MemberID implements response to request for 'memberID' field.
	MemberID(p graphql.ResolveParams) (string, error)

	// Alarm implements response to request for 'alarm' field.
	Alarm(p graphql.ResolveParams) (EtcdAlarmType, error)
}

// EtcdAlarmMemberAliases implements all methods on EtcdAlarmMemberFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type EtcdAlarmMemberAliases struct{}

// MemberID implements response to request for 'memberID' field.
func (_ EtcdAlarmMemberAliases) MemberID(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := val.(string)
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'memberID'")
	}
	return ret, err
}

// Alarm implements response to request for 'alarm' field.
func (_ EtcdAlarmMemberAliases) Alarm(p graphql.ResolveParams) (EtcdAlarmType, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret, ok := EtcdAlarmType(val.(string)), true
	if err != nil {
		return ret, err
	}
	if !ok {
		return ret, errors.New("unable to coerce value for field 'alarm'")
	}
	return ret, err
}

// EtcdAlarmMemberType Describes the state of an Etcd alarm
var EtcdAlarmMemberType = graphql.NewType("EtcdAlarmMember", graphql.ObjectKind)

// RegisterEtcdAlarmMember registers EtcdAlarmMember object type with given service.
func RegisterEtcdAlarmMember(svc *graphql.Service, impl EtcdAlarmMemberFieldResolvers) {
	svc.RegisterObject(_ObjectTypeEtcdAlarmMemberDesc, impl)
}
func _ObjTypeEtcdAlarmMemberMemberIDHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		MemberID(p graphql.ResolveParams) (string, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.MemberID(frp)
	}
}

func _ObjTypeEtcdAlarmMemberAlarmHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Alarm(p graphql.ResolveParams) (EtcdAlarmType, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {

		val, err := resolver.Alarm(frp)
		return string(val), err
	}
}

func _ObjectTypeEtcdAlarmMemberConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Describes the state of an Etcd alarm",
		Fields: graphql1.Fields{
			"alarm": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "The type of alarm which has been raised.",
				Name:              "alarm",
				Type:              graphql1.NewNonNull(graphql.OutputType("EtcdAlarmType")),
			},
			"memberID": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "ID of the member associated with the raised alarm.",
				Name:              "memberID",
				Type:              graphql1.NewNonNull(graphql1.String),
			},
		},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see EtcdAlarmMemberFieldResolvers.")
		},
		Name: "EtcdAlarmMember",
	}
}

// describe EtcdAlarmMember's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeEtcdAlarmMemberDesc = graphql.ObjectDesc{
	Config: _ObjectTypeEtcdAlarmMemberConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"alarm":    _ObjTypeEtcdAlarmMemberAlarmHandler,
		"memberID": _ObjTypeEtcdAlarmMemberMemberIDHandler,
	},
}

// EtcdAlarmType Alarm describes the type of alarm which has been raised.
type EtcdAlarmType string

// EtcdAlarmTypes holds enum values
var EtcdAlarmTypes = _EnumTypeEtcdAlarmTypeValues{
	CORRUPT: "CORRUPT",
	NONE:    "NONE",
	NOSPACE: "NOSPACE",
}

// EtcdAlarmTypeType Alarm describes the type of alarm which has been raised.
var EtcdAlarmTypeType = graphql.NewType("EtcdAlarmType", graphql.EnumKind)

// RegisterEtcdAlarmType registers EtcdAlarmType object type with given service.
func RegisterEtcdAlarmType(svc *graphql.Service) {
	svc.RegisterEnum(_EnumTypeEtcdAlarmTypeDesc)
}
func _EnumTypeEtcdAlarmTypeConfigFn() graphql1.EnumConfig {
	return graphql1.EnumConfig{
		Description: "Alarm describes the type of alarm which has been raised.",
		Name:        "EtcdAlarmType",
		Values: graphql1.EnumValueConfigMap{
			"CORRUPT": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "CORRUPT",
			},
			"NONE": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "NONE",
			},
			"NOSPACE": &graphql1.EnumValueConfig{
				DeprecationReason: "",
				Description:       "self descriptive",
				Value:             "NOSPACE",
			},
		},
	}
}

// describe EtcdAlarmType's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _EnumTypeEtcdAlarmTypeDesc = graphql.EnumDesc{Config: _EnumTypeEtcdAlarmTypeConfigFn}

type _EnumTypeEtcdAlarmTypeValues struct {
	// NONE - self descriptive
	NONE EtcdAlarmType
	// NOSPACE - self descriptive
	NOSPACE EtcdAlarmType
	// CORRUPT - self descriptive
	CORRUPT EtcdAlarmType
}

//
// ClusterHealthFieldResolvers represents a collection of methods whose products represent the
// response values of the 'ClusterHealth' type.
type ClusterHealthFieldResolvers interface {
	// Etcd implements response to request for 'etcd' field.
	Etcd(p graphql.ResolveParams) (interface{}, error)
}

// ClusterHealthAliases implements all methods on ClusterHealthFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
type ClusterHealthAliases struct{}

// Etcd implements response to request for 'etcd' field.
func (_ ClusterHealthAliases) Etcd(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// ClusterHealthType Describes the health of the Sensu backend and it's components
var ClusterHealthType = graphql.NewType("ClusterHealth", graphql.ObjectKind)

// RegisterClusterHealth registers ClusterHealth object type with given service.
func RegisterClusterHealth(svc *graphql.Service, impl ClusterHealthFieldResolvers) {
	svc.RegisterObject(_ObjectTypeClusterHealthDesc, impl)
}
func _ObjTypeClusterHealthEtcdHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(interface {
		Etcd(p graphql.ResolveParams) (interface{}, error)
	})
	return func(frp graphql1.ResolveParams) (interface{}, error) {
		return resolver.Etcd(frp)
	}
}

func _ObjectTypeClusterHealthConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "Describes the health of the Sensu backend and it's components",
		Fields: graphql1.Fields{"etcd": &graphql1.Field{
			Args:              graphql1.FieldConfigArgument{},
			DeprecationReason: "",
			Description:       "Returns health of the etcd cluster.",
			Name:              "etcd",
			Type:              graphql.OutputType("EtcdClusterHealth"),
		}},
		Interfaces: []*graphql1.Interface{},
		IsTypeOf: func(_ graphql1.IsTypeOfParams) bool {
			// NOTE:
			// Panic by default. Intent is that when Service is invoked, values of
			// these fields are updated with instantiated resolvers. If these
			// defaults are called it is most certainly programmer err.
			// If you're see this comment then: 'Whoops! Sorry, my bad.'
			panic("Unimplemented; see ClusterHealthFieldResolvers.")
		},
		Name: "ClusterHealth",
	}
}

// describe ClusterHealth's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeClusterHealthDesc = graphql.ObjectDesc{
	Config:        _ObjectTypeClusterHealthConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{"etcd": _ObjTypeClusterHealthEtcdHandler},
}
