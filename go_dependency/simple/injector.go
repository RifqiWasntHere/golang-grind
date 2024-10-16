//go:build wireinject
// +build wireinject

package simple

import "github.com/google/wire"

func InitializeService(isError bool) (*SimpleService, error) {
	wire.Build(
		NewSimpleRepository, NewSimpleService,
	)
	return nil, nil
}

// Multiple binding
func InitializeDatabase() *DatabaseRepository {
	wire.Build(
		NewDatabasePostgreSQL,
		NewDatabaseMongoDB,
		NewDatabaseRepository,
	)
	return nil
}

// Provider set
var fooSet = wire.NewSet(NewFooRepository, NewFooService)
var barSet = wire.NewSet(NewBarRepository, NewBarService)
var fooBarSet = wire.NewSet(NewFooBarService)

func InitializeFooBarService() *FooBarService {
	wire.Build(fooSet, barSet, fooBarSet) //fooBarSet is best kept hardcoded in the build for efficiency

	return nil
}

// Binding Interface
// SayHello sets
var helloSet = wire.NewSet(
	NewSayHelloImpl,
	wire.Bind(new(SayHello), new(*SayHelloImpl)), //new() is used to return pointer of it's parameter (equals to pointer type conversion [(SayHello)(nil)])
)

func InitializeSayHelloService() *SayHelloService {
	wire.Build(helloSet, NewSayHelloService)
	return nil
}

// Struct Provider
var fooSet2 = wire.NewSet(NewFooRepository, NewFooService)
var barSet2 = wire.NewSet(NewBarRepository, NewBarService)

func InitializeFooBarStruct() *FooBarService {
	wire.Build(fooSet2, barSet2, wire.Struct(new(FooBarService), "FooService", "BarService")) // Wilcard is available to assign all struct's field "*"
	return nil
}

// Binding Values

var fooValue = &FooService{}
var barValue = &BarService{}

func InitializeFoobarStructWithValues() *FooBarService {
	wire.Build(wire.Value(fooValue), wire.Value(barValue), wire.Struct(new(FooBarService), "*"))
	return nil
}
