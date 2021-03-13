// Package gogoanime scrapes and navigates the gogoanime streaming site.
package gogoanime

import (
	"fmt"

	"github.com/aaronsky/animenight/env"
)

// FindEpisodeURL returns the canonical Gogoanime URL for the episode in question.
func FindEpisodeURL(id string, episode int) string {
	// since we can't reliably scrape (we don't want to execute arbitrary JS), we have to fake it.
	return episodeURL(id, episode)
}

func episodeURL(id string, episode int) string {
	return fmt.Sprintf("https://%s/%s-episode-%d", env.GogoanimeDomain(), id, episode)
}
