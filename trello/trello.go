// Package trello wraps the Trello API.
package trello

import (
	"errors"

	"github.com/adlio/trello"
)

// ErrListNotFound is raised when a board contains no lists matching the provided name or ID.
var ErrListNotFound = errors.New("no list found with the name or ID")

// CustomFields stores custom field IDs used to identify custom field items in Trello cards.
type CustomFields struct {
	episodeNumberID string
	gogoanimeID     string
}

// Client wraps the Trello API.
type Client struct {
	client *trello.Client
}

// NewClient creates a new instance of Client.
func NewClient(appKey, token string) *Client {
	return &Client{
		client: trello.NewClient(appKey, token),
	}
}

// Board returns the board with the given ID.
//
// Returns an error if an issue is encountered fetching the board.
func (c *Client) Board(id string) (*trello.Board, error) {
	return c.client.GetBoard(id, trello.Defaults())
}

// CardsInList returns all the cards in the given list in a board.
//
// Returns an error when the board ID is invalid, if there is an issue fetching cards on a list,
// or if no lists are found on the board that match the given list name or ID.
func (c *Client) CardsInList(list string, board *trello.Board) ([]*trello.Card, error) {
	lists, err := board.GetLists(trello.Defaults())
	if err != nil {
		return nil, err
	}

	var target *trello.List

	for _, l := range lists {
		if l.ID != list && l.Name != list {
			continue
		}

		target = l

		break
	}

	if target == nil {
		return nil, ErrListNotFound
	}

	cards, err := target.GetCards(trello.Arguments{"customFieldItems": "true"})
	if err != nil {
		return nil, err
	}

	return cards, nil
}

// CustomFields returns all the custom fields with supported implementations registered to a board.
//
// Currently supported are:
// *  "Episode" with the episode number (int)
// *  "Gogoanime ID" with the gogoanime show ID (string)
//
// Returns an error if an issue is encountered fetching fields.
func (c *Client) CustomFields(board *trello.Board) (*CustomFields, error) {
	fields, err := board.GetCustomFields(trello.Defaults())
	if err != nil {
		return nil, err
	}

	ids := CustomFields{}

	for _, f := range fields {
		switch f.Name {
		case "Episode":
			ids.episodeNumberID = f.ID
		case "Gogoanime ID":
			ids.gogoanimeID = f.ID
		}
	}

	return &ids, nil
}

// EpisodeNumber returns the episode number custom field on a card.
func (f *CustomFields) EpisodeNumber(card *trello.Card) int {
	if card == nil {
		return -1
	}

	for _, field := range card.CustomFieldItems {
		if field == nil {
			continue
		} else if field.IDCustomField != f.episodeNumberID {
			continue
		}

		if val, ok := field.Value.Get().(int); ok {
			return val
		}
	}

	return -1
}

// Gogoanime returns the gogoanime show ID custom field on a card.
func (f *CustomFields) Gogoanime(card *trello.Card) string {
	if card == nil {
		return ""
	}

	for _, field := range card.CustomFieldItems {
		if field == nil {
			continue
		} else if field.IDCustomField != f.gogoanimeID {
			continue
		}

		if val, ok := field.Value.Get().(string); ok {
			return val
		}
	}

	return ""
}
