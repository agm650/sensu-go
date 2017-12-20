package generator

import (
	"fmt"
	"strings"

	"github.com/dave/jennifer/jen"
	"github.com/jamesdphillips/graphql/language/ast"
)

func genObjectType(node *ast.ObjectDefinition) jen.Code {
	code := newGroup()
	name := node.GetName().Value
	resolverName := fmt.Sprintf("%sResolver", name)

	//
	// Generate resolver interface
	//
	// ... comment: Describe resolver interface and usage
	// ... method:  [one method for each field]
	//

	code.Commentf(`//
// %s represents a collection of methods whose products represent the
// response values of the '%s' type.
//
//  == Example SDL
//
//    """
//    Dog's are not hooman.
//    """
//    type Dog implements Pet {
//      "name of this fine beast."
//      name:  String!
//
//      "breed of this silly animal; probably shibe."
//      breed: [Breed]
//    }
//
//  == Example generated interface
//
//   // DogResolver ...
//   type DogResolver interface {
//     // Name implements response to request for name field.
//     Name(context.Context, interface{}, graphql.Params) interface{}
//     // Breed implements response to request for breed field.
//     Breed(context.Context, interface{}, graphql.Params) interface{}
//     // IsTypeOf is used to determine if a given value is associated with the Dog type
//     IsTypeOf(interface{}, graphql.IsTypeOfParams) bool
//   }
//
//  == Example implementation ...
//
//   // MyDogResolver implements DogResolver interface
//   type MyDogResolver struct {
//     logger logrus.LogEntry
//     store interface{
//       store.BreedStore
//       store.DogStore
//     }
//   }
//
//   // Name implements response to request for name field.
//   func (r *MyDogResolver) Name(ctx context.Context, r interface{}, p graphql.Params) interface{} {
//     // ... implementation details ...
//     dog := r.(DogGetter)
//     return dog.GetName()
//   }
//
//   // Breed implements response to request for breed field.
//   func (r *MyDogResolver) Name(ctx context.Context, r interface{}, p graphql.Params) interface{} {
//     // ... implementation details ...
//     dog := r.(DogGetter)
//     breed := r.store.GetBreed(dog.GetBreedName())
//     return breed
//   }
//
//   // IsTypeOf is used to determine if a given value is associated with the Dog type
//   func (r *MyDogResolver) IsTypeOf(r interface{}, p graphql.IsTypeOfParams) interface{} {
//     // ... implementation details ...
//     _, ok := r.(DogGetter)
//     return ok
//   }`,
		resolverName,
		name,
	)
	// Generate resolver interface.
	code.Type().Id(resolverName).InterfaceFunc(func(g *jen.Group) {
		for _, field := range node.Fields {
			// Define method for each field in object type
			name := field.Name.Value
			titleizedName := strings.Title(field.Name.Value)

			// intended interface is (context.Context, record interface{}, params graphql.Params).
			// while this differs from graphql packages I feel it is more conventional.
			g.Commentf("%s implements response to request for '%s' field.", name, titleizedName)
			g.Id(titleizedName).Params(
				jen.Qual("context", "Context"),
				jen.Id("interface{}"),
				jen.Qual(graphqlPkg, "Params"),
			).Interface()
		}

		// Satisfy IsTypeOf() callback
		g.Commentf("IsTypeOf is used to determine if a given value is associated with the %s type", name)
		g.Id("IsTypeOf").Params( // IsTypeOf(context.Context, graphql.IsTypeOfParams) bool
			jen.Qual("context", "Context"),
			jen.Qual(graphqlPkg, "IsTypeOfParams"),
		).Bool()
	})

	//
	// Generate type definition
	//
	// ... comment: Include description in comment
	// ... panic callbacks panic if not configured
	//

	// Object ype description
	typeDesc := fetchDescription(node)

	// To appease the linter ensure that the the description of the object type
	// begins with the name of the resulting method.
	desc := typeDesc
	if hasPrefix := strings.HasPrefix(typeDesc, name); !hasPrefix {
		desc = name + " " + desc
	}

	// Generate interface references
	ints := jen.Index().Op("*").Qual(graphqlPkg, "Interface").ValuesFunc(
		func(g *jen.Group) {
			for _, n := range node.Interfaces {
				g.Line().Add(genMockInterfaceReference(n))
			}
		},
	)

	//
	// Generates thunk that returns new instance of object config
	//
	//  == Example input SDL
	//
	//    """
	//    Dogs are not hooman.
	//    """
	//    type Dog implements Pet {
	//      "name of this fine beast."
	//      name:  String!
	//
	//      "breed of this silly animal; probably shibe."
	//      breed: [Breed]
	//    }
	//
	//  == Example output
	//
	//   // Dogs are not hooman
	//   func Dog() graphql.ObjectConfig { // implements TypeThunk
	//     return graphql.ObjectConfig{
	//       Name:        "Dog",
	//       Description: "are not hooman",
	//       Interfaces:  // ...
	//       Fields:      // ...
	//       IsKindOf:    // ...
	//     }
	//   }
	//
	code.Comment(desc)
	code.Func().Id(name).Params().Qual(graphqlPkg, "ObjectConfig").Block(
		jen.Return(jen.Qual(graphqlPkg, "ObjectConfig").Values(jen.Dict{
			jen.Id("Name"):        jen.Lit(name),
			jen.Id("Description"): jen.Lit(typeDesc),
			jen.Id("Interfaces"):  ints,
			jen.Id("Fields"):      genFields(node.Fields),
			jen.Id("IsTypeOf"): jen.Func().Params(jen.Id("_").Qual(graphqlPkg, "IsTypeOfParams")).Bool().Block(
				jen.Comment(missingResolverNote),
				jen.Panic(jen.Lit("Unimplemented; see "+resolverName+".")),
			),
		})),
	)

	return code
}
