package config

import (
	cliBase "github.com/kahnwong/cli-base-sops"
)

type Config struct {
	DiscordAppIDSubsonic string `yaml:"DISCORD_APP_ID_SUBSONIC"`
	DiscordAppIDJellyfin string `yaml:"DISCORD_APP_ID_JELLYFIN"`
	DiscordAppIDIntellij string `yaml:"DISCORD_APP_ID_INTELLIJ"`
	SubsonicUsername     string `yaml:"SUBSONIC_USERNAME"`
	SubsonicToken        string `yaml:"SUBSONIC_TOKEN"`
	SubsonicSalt         string `yaml:"SUBSONIC_SALT"`
	SubsonicApiEndpoint  string `yaml:"SUBSONIC_API_ENDPOINT"`
	JelllyfinApiKey      string `yaml:"JELLYFIN_API_KEY"`
	JellyfinApiEndpoint  string `yaml:"JELLYFIN_API_ENDPOINT"`
}

var AppConfig = cliBase.ReadYamlSops[Config]("~/.config/media-discord-rpc/config.sops.yaml")
