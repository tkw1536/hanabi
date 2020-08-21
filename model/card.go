package model

// Card represents a single Hanabi Card.
//
// A Card has only two fields, a color and a number.
type Card struct {
	Color  CardColor  `json:"color,omitempty"`
	Number CardNumber `json:"number,omitempty"`
}

func (c Card) String() string {
	return c.Color.String() + " " + c.Number.String()
}

// Valid checks if Card is valid.
//
// A Card is valid if it contains both a valid color and valid number.
func (c Card) Valid() bool {
	return c.Color.Valid() && c.Number.Valid()
}

// Legal checks if a card is legal in a specific GameMode.
// This function assumes that mode is valid.
func (c Card) Legal(mode GameMode) bool {
	// need validity check because Count pre-condition
	return c.Valid() && mode.Count(c) != 0
}

// ForEachValidCard calls f exactly once for each card that is considered valid.
// The order in which f is called on the cards is the following.
//
// Each of the colors are hit in the order Blue,Green,Red,White,Yellow,Rainbow
// Within each color, cards are hit in ascending order.
func ForEachValidCard(f func(Card)) {
	for _, color := range []CardColor{
		ColorBlue,
		ColorGreen,
		ColorRed,
		ColorWhite,
		ColorYellow,
		ColorRainbow,
	} {
		for _, number := range []CardNumber{
			NumberOne,
			NumberTwo,
			NumberThree,
			NumberFour,
			NumberFive,
		} {
			f(Card{Color: color, Number: number})
		}
	}
}

// Hint represents a Hint on a set of cards.
//
// A Hint uses the same struct as a card, except that it expects
// exactly one of the fields (Color or Number) to be invalid, i.e. have the zero value.
// See the Valid function.
type Hint Card

// Valid checks if this hint is valid.
// For a hint to be valid exactly one of the fields (Color or Number) needs to be invalid.
// This function does not determine if a hint actually applies to a card in a game.
func (h Hint) Valid() bool {
	return h.Color.Valid() != h.Number.Valid()
}

func (h Hint) String() string {
	switch {
	case h.IsNumberHint():
		return "Number Hint on " + h.Number.String()
	case h.IsColorHint():
		return "Color Hint on " + h.Color.String()
	}
	return "Invalid Hint on " + Card(h).String()
}

// Legal checks if a hint is legal in a given GameMode.
// This assumes that the GameMode is valid, and may panic if not.
func (h Hint) Legal(mode GameMode) bool {
	// Valid hints are legal as follows:
	// - Number hints are legal in every mode.
	// - Non-rainbow color hints are legal in every mode
	// - Rainbow color hints are legal only in SixColor mode
	//
	// For valid hints, this compresses down to a single 'or' condition.
	// This works because for number hints h.Color == ColorUnspecified, that means h.Color != ColorRainbow will be true.

	if !h.Valid() {
		return false
	}

	return h.Color != ColorRainbow || mode == ModeSixColor
}

// IsNumberHint checks if this hint represents a valid number hint.
func (h Hint) IsNumberHint() bool {
	return h.Valid() && !h.Color.Valid()
}

// IsColorHint checks if a hint represents a valid color hint.
func (h Hint) IsColorHint() bool {
	return h.Valid() && !h.Number.Valid()
}

// Matches checks if a hint matches a card in this GameMode.
// Assumes that h.Legal(mode) and c.Legal() are true.
func (h Hint) Matches(c Card, mode GameMode) bool {
	// - number hints match all cards of that number
	// - in non-Rainbow + non-DarkRainbow color hints match all cards of that color
	// - in Rainbow + non-DarkRainbow, color hints match the color and rainbow cards

	if (mode == ModeRainbow || mode == ModeDarkRainbow) && h.IsColorHint() {
		return h.Color == c.Color || c.Color == ColorRainbow
	}

	// regular color matching
	return h.Number == c.Number || h.Color == c.Color
}

// CardColor represents the color of a card in hanabi
type CardColor string

// The six different colors of the game.
// These are represented as strings so that they can be read in JSON by external libraries
const (
	ColorUnspecified CardColor = ""
	ColorBlue        CardColor = "blue"
	ColorGreen       CardColor = "green"
	ColorRed         CardColor = "red"
	ColorWhite       CardColor = "white"
	ColorYellow      CardColor = "yellow"
	ColorRainbow     CardColor = "rainbow"
)

// Hint returns a new Color Hint of this color
func (c CardColor) Hint() Hint {
	return Hint{
		Color: c,
	}
}

func (c CardColor) String() string {
	switch c {
	case ColorBlue:
		return "Blue"
	case ColorGreen:
		return "Green"
	case ColorRed:
		return "Red"
	case ColorWhite:
		return "White"
	case ColorYellow:
		return "Yellow"
	case ColorRainbow:
		return "Rainbow"
	}
	return "?"
}

// Valid checks if the given CardColor is valid
func (c CardColor) Valid() bool {
	switch c {
	case ColorBlue, ColorGreen, ColorRed, ColorWhite, ColorYellow, ColorRainbow:
		return true
	}
	return false
}

// CardNumber represents the number of a card in the game
type CardNumber uint8

// The different numbers of cards in the game -- from 1 to 5 and an unspecefied number
// Note that NumberOne - NumberFive correspond to their uint8 counterparts
const (
	NumberUnspecified CardNumber = iota
	NumberOne
	NumberTwo
	NumberThree
	NumberFour
	NumberFive
)

func (n CardNumber) String() string {
	switch n {
	case NumberOne:
		return "1"
	case NumberTwo:
		return "2"
	case NumberThree:
		return "3"
	case NumberFour:
		return "4"
	case NumberFive:
		return "5"
	}
	return "?"
}

// Hint returns a new Number Hint of this Number
func (n CardNumber) Hint() Hint {
	return Hint{
		Number: n,
	}
}

// Valid checks if this CardNumber is valid
func (n CardNumber) Valid() bool {
	switch n {
	case NumberOne, NumberTwo, NumberThree, NumberFour, NumberFive:
		return true
	}
	return false
}
