// Package env loads flags and variables from the environment.
package env

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"

	dotenv "github.com/joho/godotenv"
)

// ErrMissingFlag is raised when there a required flag is missing from the
// environment or program options.
var ErrMissingFlag = errors.New("missing required flag or env")

// Env contains all the configuration options for the program.
type Env struct {
	trelloAppKey    string
	trelloToken     string
	trelloBoardID   string
	trelloList      string
	gogoanimeDomain string
}

// nolint: gochecknoglobals
var (
	shared Env

	trelloAppKey    = flag.String("trello_app_key", "", "Trello app key used to authenticate with the Trello REST API")
	trelloToken     = flag.String("trello_token", "", "Trello token used to authenticate as a specific user in the Trello REST API")
	trelloBoardID   = flag.String("trello_board_id", "", "Trello board ID used to identify a particular board")
	trelloList      = flag.String("trello_list", "", "Trello list to check for cards")
	gogoanimeDomain = flag.String("gogoanime_domain", "", "Gogoanime domain for episode URLs")
)

// Load loads all preliminary datum.
func Load() error {
	if err := dotenv.Load(); err != nil {
		log.Printf("[WARN]: %s\n", err)
	}

	var err error

	shared, err = New()

	return err
}

// New creates a new env instance.
func New() (Env, error) {
	var env Env

	env.trelloAppKey = loadStringFlagOrEnv(trelloAppKey, "TRELLO_APP_KEY")
	env.trelloToken = loadStringFlagOrEnv(trelloToken, "TRELLO_TOKEN")
	env.trelloBoardID = loadStringFlagOrEnv(trelloBoardID, "TRELLO_BOARD_ID")
	env.trelloList = loadStringFlagOrEnv(trelloList, "TRELLO_LIST")
	env.gogoanimeDomain = loadStringFlagOrEnv(gogoanimeDomain, "GOGOANIME_DOMAIN")

	switch {
	case env.trelloAppKey == "":
		return env, fmt.Errorf("%w: -trello_app_key or TRELLO_APP_KEY", ErrMissingFlag)
	case env.trelloToken == "":
		return env, fmt.Errorf("%w: -trello_token or TRELLO_TOKEN", ErrMissingFlag)
	case env.trelloBoardID == "":
		return env, fmt.Errorf("%w: -trello_board_id or TRELLO_BOARD_ID", ErrMissingFlag)
	case env.trelloList == "":
		return env, fmt.Errorf("%w: -trello_list or TRELLO_LIST", ErrMissingFlag)
	case env.gogoanimeDomain == "":
		return env, fmt.Errorf("%w: -gogoanime_domain or GOGOANIME_DOMAIN", ErrMissingFlag)
	}

	return env, nil
}

func loadStringFlagOrEnv(f *string, envName string) string {
	if f != nil && *f != "" {
		return *f
	}

	return os.Getenv(envName)
}

// TrelloAppKey returns the app key identifying the authenticated application in Trello.
func TrelloAppKey() string {
	return shared.TrelloAppKey()
}

// TrelloToken returns the token identifying the authenticated user in Trello.
func TrelloToken() string {
	return shared.TrelloToken()
}

// TrelloBoardID returns the board ID to search in Trello.
func TrelloBoardID() string {
	return shared.TrelloBoardID()
}

// TrelloList returns the list of cards to search in Trello.
func TrelloList() string {
	return shared.TrelloList()
}

// GogoanimeDomain returns the address domain for the gogoanime website.
func GogoanimeDomain() string {
	return shared.GogoanimeDomain()
}

// TrelloAppKey returns the app key identifying the authenticated application in Trello.
func (e Env) TrelloAppKey() string {
	return e.trelloAppKey
}

// TrelloToken returns the token identifying the authenticated user in Trello.
func (e Env) TrelloToken() string {
	return e.trelloToken
}

// TrelloBoardID returns the board ID to search in Trello.
func (e Env) TrelloBoardID() string {
	return e.trelloBoardID
}

// TrelloList returns the list of cards to search in Trello.
func (e Env) TrelloList() string {
	return e.trelloList
}

// GogoanimeDomain returns the address domain for the gogoanime website.
func (e Env) GogoanimeDomain() string {
	return e.gogoanimeDomain
}
