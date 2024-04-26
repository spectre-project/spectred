package prefixmanager

import (
	"github.com/spectre-project/spectred/infrastructure/logger"
	"github.com/spectre-project/spectred/util/panics"
)

var log = logger.RegisterSubSystem("PRFX")
var spawn = panics.GoroutineWrapperFunc(log)
