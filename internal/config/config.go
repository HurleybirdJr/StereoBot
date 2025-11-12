package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	Environment         string            `json:"environment"`
	GuildID             string            `json:"guild_id"`
	GeneralChannelID    string            `json:"general_channel_id"`
	InstantBanChannelId string            `json:"instant_ban_channel_id"`
	Roles               map[string]string `json:"roles"`
	Reactions           Reactions         `json:"reactions"`
}

type Reactions struct {
	InMessage     []map[string]Reaction `json:"inMessage"`
	InStickerName []map[string]Reaction `json:"inStickerName"`
}

type Reaction struct {
	Emoji   *string `json:"emoji"`
	Message *string `json:"message"`
}

func (c *Config) Load() error {
	dir, _ := os.Getwd()

	currentEnv := os.Getenv("ENV")

	dat, err := os.ReadFile(dir + fmt.Sprintf("/config/config.%s.json", currentEnv)) //TODO: DO NOT PUSH WITHOUT CHANGING BACK TO CONFIG.JSON!!!

	if err != nil {
		return err
	}

	err = json.Unmarshal(dat, c)

	return err
}
