// Code generated by scripts/gengraphql.go. DO NOT EDIT.

package schema

import (
	fmt "fmt"
	graphql1 "github.com/graphql-go/graphql"
	graphql "github.com/sensu/sensu-go/graphql"
)

// UserUsernameFieldResolver implement to resolve requests for the User's username field.
type UserUsernameFieldResolver interface {
	// Username implements response to request for username field.
	Username(p graphql.ResolveParams) (string, error)
}

// UserRolesFieldResolver implement to resolve requests for the User's roles field.
type UserRolesFieldResolver interface {
	// Roles implements response to request for roles field.
	Roles(p graphql.ResolveParams) (interface{}, error)
}

// UserDisabledFieldResolver implement to resolve requests for the User's disabled field.
type UserDisabledFieldResolver interface {
	// Disabled implements response to request for disabled field.
	Disabled(p graphql.ResolveParams) (bool, error)
}

// UserHasPasswordFieldResolver implement to resolve requests for the User's hasPassword field.
type UserHasPasswordFieldResolver interface {
	// HasPassword implements response to request for hasPassword field.
	HasPassword(p graphql.ResolveParams) (bool, error)
}

//
// UserFieldResolvers represents a collection of methods whose products represent the
// response values of the 'User' type.
//
// == Example SDL
//
//   """
//   Dog's are not hooman.
//   """
//   type Dog implements Pet {
//     "name of this fine beast."
//     name:  String!
//
//     "breed of this silly animal; probably shibe."
//     breed: [Breed]
//   }
//
// == Example generated interface
//
//   // DogResolver ...
//   type DogFieldResolvers interface {
//     DogNameFieldResolver
//     DogBreedFieldResolver
//
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
// == Example implementation ...
//
//   // DogResolver implements DogFieldResolvers interface
//   type DogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *DogResolver) Name(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // ... implementation details ...
//     dog := p.Source.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *DogResolver) IsTypeOf(p graphql.IsTypeOfParams) bool {
//     // ... implementation details ...
//     _, ok := p.Value.(DogGetter)
//     return ok
//   }
//
type UserFieldResolvers interface {
	UserUsernameFieldResolver
	UserRolesFieldResolver
	UserDisabledFieldResolver
	UserHasPasswordFieldResolver
}

// UserAliases implements all methods on UserFieldResolvers interface by using reflection to
// match name of field to a field on the given value. Intent is reduce friction
// of writing new resolvers by removing all the instances where you would simply
// have the resolvers method return a field.
//
// == Example SDL
//
//    type Dog {
//      name:   String!
//      weight: Float!
//      dob:    DateTime
//      breed:  [Breed]
//    }
//
// == Example generated aliases
//
//   type DogAliases struct {}
//   func (_ DogAliases) Name(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Weight(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Dob(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//   func (_ DogAliases) Breed(p graphql.ResolveParams) (interface{}, error) {
//     // reflect...
//   }
//
// == Example Implementation
//
//   type DogResolver struct { // Implements DogResolver
//     DogAliases
//     store store.BreedStore
//   }
//
//   // NOTE:
//   // All other fields are satisified by DogAliases but since this one
//   // requires hitting the store we implement it in our resolver.
//   func (r *DogResolver) Breed(p graphql.ResolveParams) interface{} {
//     dog := v.(*Dog)
//     return r.BreedsById(dog.BreedIDs)
//   }
//
type UserAliases struct{}

// Username implements response to request for 'username' field.
func (_ UserAliases) Username(p graphql.ResolveParams) (string, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := fmt.Sprint(val)
	return ret, err
}

// Roles implements response to request for 'roles' field.
func (_ UserAliases) Roles(p graphql.ResolveParams) (interface{}, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	return val, err
}

// Disabled implements response to request for 'disabled' field.
func (_ UserAliases) Disabled(p graphql.ResolveParams) (bool, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(bool)
	return ret, err
}

// HasPassword implements response to request for 'hasPassword' field.
func (_ UserAliases) HasPassword(p graphql.ResolveParams) (bool, error) {
	val, err := graphql.DefaultResolver(p.Source, p.Info.FieldName)
	ret := val.(bool)
	return ret, err
}

// UserType User describes an operator in the system
var UserType = graphql.NewType("User", graphql.ObjectKind)

// RegisterUser registers User object type with given service.
func RegisterUser(svc *graphql.Service, impl UserFieldResolvers) {
	svc.RegisterObject(_ObjectTypeUserDesc, impl)
}
func _ObjTypeUserUsernameHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(UserUsernameFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Username(p)
	}
}

func _ObjTypeUserRolesHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(UserRolesFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Roles(p)
	}
}

func _ObjTypeUserDisabledHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(UserDisabledFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.Disabled(p)
	}
}

func _ObjTypeUserHasPasswordHandler(impl interface{}) graphql1.FieldResolveFn {
	resolver := impl.(UserHasPasswordFieldResolver)
	return func(p graphql1.ResolveParams) (interface{}, error) {
		return resolver.HasPassword(p)
	}
}

func _ObjectTypeUserConfigFn() graphql1.ObjectConfig {
	return graphql1.ObjectConfig{
		Description: "User describes an operator in the system",
		Fields: graphql1.Fields{
			"disabled": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "disabled",
				Type:              graphql1.NewNonNull(graphql1.Boolean),
			},
			"hasPassword": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "hasPassword",
				Type:              graphql1.NewNonNull(graphql1.Boolean),
			},
			"roles": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "roles",
				Type:              graphql1.NewNonNull(graphql1.NewList(graphql1.NewNonNull(graphql.OutputType("Role")))),
			},
			"username": &graphql1.Field{
				Args:              graphql1.FieldConfigArgument{},
				DeprecationReason: "",
				Description:       "self descriptive",
				Name:              "username",
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
			panic("Unimplemented; see UserFieldResolvers.")
		},
		Name: "User",
	}
}

// describe User's configuration; kept private to avoid unintentional tampering of configuration at runtime.
var _ObjectTypeUserDesc = graphql.ObjectDesc{
	Config: _ObjectTypeUserConfigFn,
	FieldHandlers: map[string]graphql.FieldHandler{
		"disabled":    _ObjTypeUserDisabledHandler,
		"hasPassword": _ObjTypeUserHasPasswordHandler,
		"roles":       _ObjTypeUserRolesHandler,
		"username":    _ObjTypeUserUsernameHandler,
	},
}
