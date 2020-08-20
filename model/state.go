package model

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"

	"github.com/google/uuid"
)

// TODO: All functions in this package are not yet tested.

// GameState represents the state of a Hanabi Game.
// This state is represented openly, i.e. every single card can be seen.
// This state is not goroutine safe, and excepts that only one goroutine accesses the game at any point.
type GameState struct {
	Mode GameMode

	// Stack is the stack new cards are dran from
	Stack []Card

	// Discarded is the stack of cards that have been discarded
	Discarded []Card

	// ColorPiles represents the current number for each color that has been played.
	// When a card has not yet been played, it will be NumberUnspecified.
	ColorPiles map[CardColor]CardNumber

	Hints    uint8 // current number of hints available
	Misplays uint8 // number of misplays so far

	// Players is the list of players
	// We use a pointer so that we can modify the player.
	Players []*Player

	// Stated returns if the game has already been started
	Started bool
	// CurrentPlayer is the player who has to make a move next
	CurrentPlayer int
}

// Player represents a player in Hanabi
type Player struct {
	// ID represents a unique ID for the player
	ID uuid.UUID

	// Hand is the Hand of the Player
	Hand []Card
}

// MoveKind represents the kind of moves a player can make.
type MoveKind string

// MovePlay, MoveHint and MoveDiscard represent the play, hint and discard moves respectively.
const (
	MovePlay    MoveKind = "play"
	MoveHint    MoveKind = "hint"
	MoveDiscard MoveKind = "discard"
)

// Move represents a move a player can make
type Move struct {
	// Kind is the kind of move the player makes
	Kind MoveKind

	// ID is the player making this move.
	// In most cases, this field may be omitted and the GameState wil fill it automatically.
	ID uuid.UUID

	// Index is the index into their hand that the player plays or discards
	Index int

	// Hint represents the hint that is being given.
	Hint Hint
	// ToPlayerID represents the player that is being hinted.
	ToPlayerID uuid.UUID
}

// ErrGameStarted represents an error that an action cannot be performed because the game has already been started
var ErrGameStarted = errors.New("Game Already started")

// AddPlayer adds a new player to the Game.
// When the Game has already started, returns ErrGameStarted
func (state *GameState) AddPlayer() (*Player, error) {
	if state.Started {
		return nil, ErrGameStarted
	}

	player := &Player{}

	// generate a new UUID for the player and make sure that it is unique for the current set of players.
	// In most cases this loop will be only one iteration.

	var err error
	idTaken := true
	for idTaken {

		// generate a new UUID or bail out
		player.ID, err = uuid.NewRandom()
		if err != nil {
			return nil, errors.Wrap(err, "Unable to generate new Player UUID")
		}

		// check that it is not already taken
		idTaken = false
		for _, p := range state.Players {
			if player.ID == p.ID {
				idTaken = true
			}
		}
	}

	state.Players = append(state.Players, player)
	return player, nil
}

// ErrModeInvalid is an error that indicates that the GameMode selected is not valid.
var ErrModeInvalid = errors.New("GameState: Mode is invalid")

// ErrInvalidPlayerCount is an error that is returned if there is the wrong number of players
var ErrInvalidPlayerCount = errors.New("GameState: There must be between 2 and 5 players")

// Start sets up this Game by initializing all internal
// data structures.
// The seed is used to shuffle the stack, and thus determines all the randomness in the game.
// If seed is 0, a random seed is picked.
func (state *GameState) Start(seed int64) error {

	// This function has to initialize the game, i.e:

	// - set Hints to the right number
	// - set Misplays to the right number
	// - initialize the color and discard piles.
	// - create and shuffle the stack
	// - distribute cards to all the players
	// - determine the first player to play

	// - setup Stack to contain all the cards

	if state.Started {
		return ErrGameStarted
	}

	if !state.Mode.Valid() {
		return ErrModeInvalid
	}

	// setup hints and misplays
	state.Hints = 8
	state.Misplays = 0

	// setup the color piles
	state.ColorPiles = make(map[CardColor]CardNumber)
	state.ColorPiles[ColorBlue] = NumberUnspecified
	state.ColorPiles[ColorGreen] = NumberUnspecified
	state.ColorPiles[ColorRed] = NumberUnspecified
	state.ColorPiles[ColorWhite] = NumberUnspecified
	state.ColorPiles[ColorYellow] = NumberUnspecified
	if state.Mode == ModeRainbow || state.Mode == ModeDarkRainbow {
		state.ColorPiles[ColorRainbow] = NumberUnspecified
	}

	// setup the stack
	state.Stack = state.Mode.NewStack()

	// Create a new random source
	// When the seed is zero, use the current time.
	source := rand.NewSource(seed)
	if seed == 0 {
		source.Seed(time.Now().UnixNano())
	}
	random := rand.New(source)

	// Shuffle the stack with it
	random.Shuffle(len(state.Stack), func(i, j int) {
		state.Stack[i], state.Stack[j] = state.Stack[j], state.Stack[i]
	})

	// setup the discard pile
	state.Discarded = make([]Card, 0, len(state.Stack))

	// determine the number of cards per player
	cardsPerPlayer := 0
	switch len(state.Players) {
	case 2, 3:
		cardsPerPlayer = 5
	case 4, 5:
		cardsPerPlayer = 4
	default:
		return ErrInvalidPlayerCount
	}

	cardsInStack := len(state.Stack)
	for _, p := range state.Players {
		// make a hand for the player
		p.Hand = make([]Card, 0, cardsPerPlayer)

		// pick the right amount of cards from the deck
		// and put them into the hand of the player
		cardsInStack -= cardsPerPlayer
		p.Hand = append(p.Hand, state.Stack[cardsInStack:]...)
		state.Stack = state.Stack[:cardsInStack]
	}

	// the first player starts
	state.CurrentPlayer = 0

	return nil
}
