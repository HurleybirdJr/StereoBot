package bot

import (
	"stereobot/internal/config"
	"stereobot/internal/logging"
)

type CommonDeps struct {
	Bot    *Bot
	Logger *logging.Logger
	Config *config.Config
}

type Module interface {
	Load(deps *CommonDeps) error
	GetName() string
}
