package command

import "context"

//go:generate mockery --case=snake --outpkg=commandmocks --output=commandmocks --name=Bus

// Bus defines the expected behaviour from a command bus.
type Bus interface {
	// Dispatch is the method used to dispatch new commands.
	Dispatch(context.Context, Command) error
	// Register is the method used to register a new command handler.
	Register(Type, Handler)
}

// Handler defines the expected behaviour from a command handler.
type Handler interface {
	Handle(context.Context, Command) error
}

// Type represents an application command type.
type Type string

// Command represents an application command.
type Command interface {
	Type() Type
}
