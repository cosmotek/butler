package service

import (
	"github.com/itchio/butler/buse"
	"github.com/itchio/butler/buse/messages"
	"github.com/itchio/butler/endpoints/cleandownloads"
	"github.com/itchio/butler/endpoints/install"
	"github.com/itchio/butler/endpoints/launch"
	"github.com/itchio/butler/endpoints/tests"
	"github.com/itchio/butler/endpoints/update"
	"github.com/itchio/butler/endpoints/utilities"
	"github.com/itchio/butler/mansion"
)

var mainRouter *buse.Router

func getRouter(mansionContext *mansion.Context) *buse.Router {
	if mainRouter != nil {
		return mainRouter
	}

	mainRouter = buse.NewRouter(mansionContext)
	utilities.Register(mainRouter)
	tests.Register(mainRouter)
	update.Register(mainRouter)
	install.Register(mainRouter)
	launch.Register(mainRouter)
	cleandownloads.Register(mainRouter)

	messages.EnsureAllRequests(mainRouter)

	return mainRouter
}