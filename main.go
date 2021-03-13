package main

import (
	"log"

	"github.com/aaronsky/animenight/browser"
	"github.com/aaronsky/animenight/env"
	"github.com/aaronsky/animenight/gogoanime"
	"github.com/aaronsky/animenight/trello"
)

func main() {
	env.Load()

	trelloClient := trello.NewClient(env.TrelloAppKey(), env.TrelloToken())

	board, err := trelloClient.Board(env.TrelloBoardID())
	if err != nil {
		log.Fatal(err)
		return
	}

	cards, err := trelloClient.CardsInList(env.TrelloList(), board)
	if err != nil {
		log.Fatal(err)
		return
	}

	fields, err := trelloClient.CustomFields(board)
	if err != nil {
		log.Fatal(err)
		return
	}

	urls := make([]string, len(cards))
	for i, c := range cards {
		urls[i] = gogoanime.FindEpisodeURL(fields.Gogoanime(c), fields.EpisodeNumber(c))
	}

	browser.OpenURLs(urls...)
}
