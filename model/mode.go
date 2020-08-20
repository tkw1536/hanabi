package model

// GameMode represents a Hanabi Game GameMode
type GameMode string

// The different Game Modes
const (
	// ModeFiveColor represents the "normal" game mode.
	// In ModeFiveColor mode, no rainbow cards exist.
	ModeFiveColor GameMode = "five-color"

	// ModeSixColor represents the Game Mode where Rainbow Cards act as a sixth color.
	ModeSixColor GameMode = "six-color"

	// ModeRainbow represents the GameMode where ModeRainbow Cards act as a sixth color and are additionally hit by every Hint.
	ModeRainbow GameMode = "rainbow"

	// ModeDarkRainbow represents a GameMode that acts like Rainbow except that each GameMode exists only once.
	ModeDarkRainbow GameMode = "dark-rainbow"
)

// Valid checks if this GameMode is valid.
// A GameMode is valid if it is a known gamemode and does not have a different value,
func (mode GameMode) Valid() bool {
	switch mode {
	case ModeFiveColor, ModeSixColor, ModeRainbow, ModeDarkRainbow:
		return true
	}
	return false
}

// Count counts how many times the provided card should occur in a new stack of this GameMode.
// This function assumes that the card and mode passed are both valid, and may invoke panic() if that is not the case.
//
// When a card is not allowed in a specified GameMode, returns 0.
func (mode GameMode) Count(card Card) int {

	// Internally, this function is relied upon as the source of truth for some methods.
	// It should not be reimplemented based on other methods.

	// Only rainbow cards have to be treated special.
	// In FiveColor mode, no rainbow cards occur.
	// In DarkRainbow, each rainbow card occurs exactly once.

	if card.Color == ColorRainbow {
		switch mode {
		case ModeFiveColor: // Five Color has no rainbow cards
			return 0
		case ModeDarkRainbow: // Dark Rainbow has each rainbow color once
			return 1
		}
	}

	// for all the 'regular' colors
	// - a 1 occurs 3 times
	// - a 2, 3 or 4 occur twice
	// - a 5 occurs once

	switch card.Number {
	case NumberOne:
		return 3
	case NumberTwo, NumberThree, NumberFour:
		return 2
	case NumberFive:
		return 1
	}
	panic("mode.CardCount(): precondition failed: card.Valid() is false")
}

// TotalCards counts the total number of cards in a given GameMode.
// This functions assumes that GameMode is valid, and may call panic if this is not the case.
func (mode GameMode) TotalCards() int {
	// A valid implementation of this function might also be:
	//
	//   var total int
	//	 ForEachValidCard(func (c Card){ total += mode.CardCount(c) })
	//   return total
	//
	// But this seems unneccessarily complicated.
	// Instead we just pre-compute the values and return them here.

	switch mode {
	case ModeFiveColor:
		return 50
	case ModeSixColor, ModeRainbow:
		return 60
	case ModeDarkRainbow:
		return 55
	}
	panic("mode.TotalCards(): precondition failed: mode.Valid() is false")
}

// NewStack returns a new stack of cards for the given GameMode
// The order of the returned stack will be the same as in ForEachValidCard.
func (mode GameMode) NewStack() []Card {
	stack := make([]Card, 0, mode.TotalCards())
	ForEachValidCard(func(c Card) {
		repeat := mode.Count(c)
		for i := 0; i < repeat; i++ {
			stack = append(stack, c)
		}
	})
	return stack
}
