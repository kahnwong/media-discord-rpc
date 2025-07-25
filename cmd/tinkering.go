package cmd

import (
	"time"

	"github.com/kahnwong/media-discord-rpc/discord"
	"github.com/kahnwong/rich-go/client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var tinkeringCmd = &cobra.Command{
	Use:   "tinkering",
	Short: "Display tinkering activity",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			discord.SetActivity(discord.DiscordApps.Intellij, client.ActivityTypes.Playing, "Tinkering", "", "")

			log.Info().Msg("Tinkering...")
			time.Sleep(60 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(tinkeringCmd)
}
