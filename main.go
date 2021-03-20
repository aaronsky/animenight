package main

import (
	"fmt"
	"log"

	"github.com/aaronsky/animenight/env"
	"github.com/aaronsky/animenight/gogoanime"
	"github.com/aaronsky/animenight/trello"
)

func main() {
	if err := env.Load(); err != nil {
		log.Fatal(err)
		return
	}

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

	for _, c := range cards {
		url := gogoanime.FindEpisodeURL(fields.Gogoanime(c), fields.EpisodeNumber(c))
		fmt.Println(url)
	}
}
