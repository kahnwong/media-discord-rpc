package integrations

import (
	"context"
	"fmt"

	jellyfin "github.com/sj14/jellyfin-go/api"
)

var client *jellyfin.APIClient

type JellyfinNowPlaying struct {
	Title    string
	Episode  string
	CoverArt string
}

func JellyfinGetNowPlaying() (JellyfinNowPlaying, error) {
	sessions, _, err := client.SessionAPI.GetSessions(context.Background()).Execute()
	if err != nil {
		return JellyfinNowPlaying{}, err
	}

	var r JellyfinNowPlaying
	if len(sessions) > 0 {
		nowPlaying := sessions[0].NowPlayingItem.Get()

		r = JellyfinNowPlaying{
			Title:    *nowPlaying.Name.Get(),
			Episode:  "",
			CoverArt: fmt.Sprintf("%s/Items/%s/Images/Primary?fillHeight=100&tag=%s", AppConfig.JellyfinApiEndpoint, *nowPlaying.Id, nowPlaying.ImageTags["Primary"]),
		}

		if *nowPlaying.Type == "Episode" {
			r = JellyfinNowPlaying{
				Title:    *nowPlaying.SeriesName.Get(),
				Episode:  fmt.Sprintf("S%vE%v", *nowPlaying.ParentIndexNumber.Get(), *nowPlaying.IndexNumber.Get()),
				CoverArt: fmt.Sprintf("%s/Items/%s/Images/Primary?fillHeight=100&tag=%s", AppConfig.JellyfinApiEndpoint, *nowPlaying.SeriesId.Get(), *nowPlaying.SeriesPrimaryImageTag.Get()),
			}
		}
	}

	return r, err
}

func init() {
	config := &jellyfin.Configuration{
		Servers: jellyfin.ServerConfigurations{{URL: AppConfig.JellyfinApiEndpoint}},
		DefaultHeader: map[string]string{
			"Authorization": fmt.Sprintf("MediaBrowser Token=\"%s\"", AppConfig.JelllyfinApiKey),
		},
	}

	client = jellyfin.NewAPIClient(config)
}
