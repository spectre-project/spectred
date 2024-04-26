package main

import (
	"github.com/spectre-project/spectred/infrastructure/logger"
	"github.com/spectre-project/spectred/util/panics"
)

var (
	backendLog = logger.NewBackend()
	log        = backendLog.Logger("MNJS")
	spawn      = panics.GoroutineWrapperFunc(log)
)
