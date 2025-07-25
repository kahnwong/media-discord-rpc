package cmd

import (
	"fmt"
	"time"

	"github.com/kahnwong/media-discord-rpc/discord"
	"github.com/kahnwong/media-discord-rpc/integrations"
	"github.com/kahnwong/rich-go/client"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
)

var listeningCmd = &cobra.Command{
	Use:   "listening",
	Short: "Display listening activity",
	Run: func(cmd *cobra.Command, args []string) {
		for {
			nowPlaying := integrations.SubsonicNowPlaying()
			fmt.Println(nowPlaying)

			if nowPlaying.Title != "" {
				discord.SetActivity(discord.DiscordApps.Subsonic, client.ActivityTypes.Listening, nowPlaying.Artist, nowPlaying.Title, nowPlaying.CoverArt)
			} else {
				log.Info().Msg("Nothing is currently playing...")
			}

			time.Sleep(15 * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(listeningCmd)
}
