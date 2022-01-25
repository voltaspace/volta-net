// +build wireinject

// The build tag makes sure the stub is not built in the final build.
//go:generate go run github.com/google/wire/cmd/wire

package voltanet

import (
	"github.com/google/wire"
)

// initApp init gateway application.
func initApp() (*App, func(), error) {
	panic(
		wire.Build(
			OptionsProvider,
			EventsProvider,
			SignalProvider,
			WebSocketProvider,
			WsHeartBeatProvider,
			newApp,
			),
		)
}

