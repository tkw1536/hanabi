package model

import (
	"math/rand"
	"time"

	"github.com/pkg/errors"

	"github.com/google/uuid"
)

// TODO: All functions in this package are not yet tested.

// MaxHintCount is the maximal number of hints a game of hanabi contains.
const MaxHintCount = 8

// MaxMisplayCount is the maximal number of misplays before a game is considered lost
const MaxMisplayCount = 3

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

	// MovesLeft indicates how many turns are left. 
	// If MovesLeft <= -1 implies that there are infinitly many moves left. 
	MovesLeft int

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

// Applicable checks if a hint is applicable to this player. 
// A hint is applicable to this player if it matches at least one card
func (p *Player) Applicable(hint Hint) bool {
	// In order for a hint to be applicable, it must be legal
	// Furthermore, the player must not be nil. 
	if !hint.Legal(state.Mode) || p == nil {
		return false
	}
	for _, card := range player.Hand {
		if hint.Matches(card, state.Mode) {
			return true
		}
	}
	return false
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
func (state *GameState) /home/twiesing/Projects/github.com/tkw1536/hanabiStart(seed int64) error {

	// This function has to initialize the game, i.e:

	// - set MovesLeft to the right number
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

	state.MovesLeft = -1

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

// Player returns the current player
func (state *GameState) Player() *Player {
	return state.Players[state.CurrentPlayer]
}

// GetPlayer returns a player by UUID
func (state *GameState) GetPlayer(uuid uuid.UUID) *Player {
	for _, p := range state.Players {
		if p.ID == uuid {
			return p
		}
	}
	return nil
}

// Finished returns if a game is finished
func (state *GameState) Finished() bool {

	// a started game can not be finished
	if (!state.Started) {
		return false
	}

	// maximal number of misplays reached => failure
	if state.Misplays == MaxMisplayCount {
		return true
	}

	// if there are no moves left the game has finished
	if state.MovesLeft == 0 {
		return true
	}

	// if all piles are at #5, the game is finished
	for _, number := range state.ColorPiles {
		if number != NumberFive {
			return false
		}
	}
	return true
}


// Legal checks if the provided move is Legal in the current game state
// Assumes that the game has been started and is not yet finished. 
// See .Started and .Finished()
func (state *GameState) Legal(move Move) bool {
	player := state.Player()
	switch move.Kind {
	case MovePlay:
		// a player can play if the index is in the range of the hand
		return 0 <= move.Index < len(player.Hand)
	case MoveDiscard:
		// a player can discard if:
		// - the index is in range of the hand
		// - the hints are not full
		return state.Hints < MaxHintCount && (0 <= move.Index < len(player.Hand))
	case MoveHint:
		// a player can hint if:
		// - there is at least one hint left
		// - the player that is being hinted exists and is not the player itself
		// - there is at least one hint
		return state.Hints > 0 && state.GetPlayer(move.ID).Applicable(move.Hint) && hinted.ID !== player.ID
	}
	return false
}

const errMoveIllegal = errors.New("Move is illegal")

// Apply applies a move
// If a move is not legal returns an error. 
func (state *GameState) ApplyMove(move Move) error {
	if (!state.Legal(move)) {
		return errMoveIllegal
	}

	switch move.Kind {
	case MovePlay:
		
	}
}