package model

import (
	"reflect"
	"testing"
)

func TestCard_Valid(t *testing.T) {
	type fields struct {
		Color  CardColor
		Number CardNumber
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Blue / 1 is valid", fields{ColorBlue, NumberOne}, true},
		{"Blue / 2 is valid", fields{ColorBlue, NumberTwo}, true},
		{"Blue / 3 is valid", fields{ColorBlue, NumberThree}, true},
		{"Blue / 4 is valid", fields{ColorBlue, NumberFour}, true},
		{"Blue / 5 is valid", fields{ColorBlue, NumberFive}, true},
		{"Blue / ? is invalid", fields{ColorBlue, NumberUnspecified}, false},

		{"Green / 1 is valid", fields{ColorGreen, NumberOne}, true},
		{"Green / 2 is valid", fields{ColorGreen, NumberTwo}, true},
		{"Green / 3 is valid", fields{ColorGreen, NumberThree}, true},
		{"Green / 4 is valid", fields{ColorGreen, NumberFour}, true},
		{"Green / 5 is valid", fields{ColorGreen, NumberFive}, true},
		{"Green / ? is invalid", fields{ColorGreen, NumberUnspecified}, false},

		{"Red / 1 is valid", fields{ColorRed, NumberOne}, true},
		{"Red / 2 is valid", fields{ColorRed, NumberTwo}, true},
		{"Red / 3 is valid", fields{ColorRed, NumberThree}, true},
		{"Red / 4 is valid", fields{ColorRed, NumberFour}, true},
		{"Red / 5 is valid", fields{ColorRed, NumberFive}, true},
		{"Red / ? is invalid", fields{ColorRed, NumberUnspecified}, false},

		{"White / 1 is valid", fields{ColorWhite, NumberOne}, true},
		{"White / 2 is valid", fields{ColorWhite, NumberTwo}, true},
		{"White / 3 is valid", fields{ColorWhite, NumberThree}, true},
		{"White / 4 is valid", fields{ColorWhite, NumberFour}, true},
		{"White / 5 is valid", fields{ColorWhite, NumberFive}, true},
		{"White / ? is invalid", fields{ColorWhite, NumberUnspecified}, false},

		{"Yellow / 1 is valid", fields{ColorYellow, NumberOne}, true},
		{"Yellow / 2 is valid", fields{ColorYellow, NumberTwo}, true},
		{"Yellow / 3 is valid", fields{ColorYellow, NumberThree}, true},
		{"Yellow / 4 is valid", fields{ColorYellow, NumberFour}, true},
		{"Yellow / 5 is valid", fields{ColorYellow, NumberFive}, true},
		{"Yellow / ? is invalid", fields{ColorYellow, NumberUnspecified}, false},

		{"Rainbow / 1 is valid", fields{ColorRainbow, NumberOne}, true},
		{"Rainbow / 2 is valid", fields{ColorRainbow, NumberTwo}, true},
		{"Rainbow / 3 is valid", fields{ColorRainbow, NumberThree}, true},
		{"Rainbow / 4 is valid", fields{ColorRainbow, NumberFour}, true},
		{"Rainbow / 5 is valid", fields{ColorRainbow, NumberFive}, true},
		{"Rainbow / ? is invalid", fields{ColorRainbow, NumberUnspecified}, false},

		{"? / 1 is invalid", fields{ColorUnspecified, NumberOne}, false},
		{"? / 2 is invalid", fields{ColorUnspecified, NumberTwo}, false},
		{"? / 3 is invalid", fields{ColorUnspecified, NumberThree}, false},
		{"? / 4 is invalid", fields{ColorUnspecified, NumberFour}, false},
		{"? / 5 is invalid", fields{ColorUnspecified, NumberFive}, false},
		{"? / ? is invalid", fields{ColorUnspecified, NumberUnspecified}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				Color:  tt.fields.Color,
				Number: tt.fields.Number,
			}
			if got := c.Valid(); got != tt.want {
				t.Errorf("Card.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHint_Valid(t *testing.T) {
	type fields struct {
		Color  CardColor
		Number CardNumber
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Blue / 1 is invalid", fields{ColorBlue, NumberOne}, false},
		{"Blue / 2 is invalid", fields{ColorBlue, NumberTwo}, false},
		{"Blue / 3 is invalid", fields{ColorBlue, NumberThree}, false},
		{"Blue / 4 is invalid", fields{ColorBlue, NumberFour}, false},
		{"Blue / 5 is invalid", fields{ColorBlue, NumberFive}, false},
		{"Blue / ? is valid", fields{ColorBlue, NumberUnspecified}, true},

		{"Green / 1 is invalid", fields{ColorGreen, NumberOne}, false},
		{"Green / 2 is invalid", fields{ColorGreen, NumberTwo}, false},
		{"Green / 3 is invalid", fields{ColorGreen, NumberThree}, false},
		{"Green / 4 is invalid", fields{ColorGreen, NumberFour}, false},
		{"Green / 5 is invalid", fields{ColorGreen, NumberFive}, false},
		{"Green / ? is valid", fields{ColorGreen, NumberUnspecified}, true},

		{"Red / 1 is invalid", fields{ColorRed, NumberOne}, false},
		{"Red / 2 is invalid", fields{ColorRed, NumberTwo}, false},
		{"Red / 3 is invalid", fields{ColorRed, NumberThree}, false},
		{"Red / 4 is invalid", fields{ColorRed, NumberFour}, false},
		{"Red / 5 is invalid", fields{ColorRed, NumberFive}, false},
		{"Red / ? is valid", fields{ColorRed, NumberUnspecified}, true},

		{"White / 1 is invalid", fields{ColorWhite, NumberOne}, false},
		{"White / 2 is invalid", fields{ColorWhite, NumberTwo}, false},
		{"White / 3 is invalid", fields{ColorWhite, NumberThree}, false},
		{"White / 4 is invalid", fields{ColorWhite, NumberFour}, false},
		{"White / 5 is invalid", fields{ColorWhite, NumberFive}, false},
		{"White / ? is valid", fields{ColorWhite, NumberUnspecified}, true},

		{"Yellow / 1 is invalid", fields{ColorYellow, NumberOne}, false},
		{"Yellow / 2 is invalid", fields{ColorYellow, NumberTwo}, false},
		{"Yellow / 3 is invalid", fields{ColorYellow, NumberThree}, false},
		{"Yellow / 4 is invalid", fields{ColorYellow, NumberFour}, false},
		{"Yellow / 5 is invalid", fields{ColorYellow, NumberFive}, false},
		{"Yellow / ? is valid", fields{ColorYellow, NumberUnspecified}, true},

		{"Rainbow / 1 is invalid", fields{ColorRainbow, NumberOne}, false},
		{"Rainbow / 2 is invalid", fields{ColorRainbow, NumberTwo}, false},
		{"Rainbow / 3 is invalid", fields{ColorRainbow, NumberThree}, false},
		{"Rainbow / 4 is invalid", fields{ColorRainbow, NumberFour}, false},
		{"Rainbow / 5 is invalid", fields{ColorRainbow, NumberFive}, false},
		{"Rainbow / ? is valid", fields{ColorRainbow, NumberUnspecified}, true},

		{"? / 1 is valid", fields{ColorUnspecified, NumberOne}, true},
		{"? / 2 is valid", fields{ColorUnspecified, NumberTwo}, true},
		{"? / 3 is valid", fields{ColorUnspecified, NumberThree}, true},
		{"? / 4 is valid", fields{ColorUnspecified, NumberFour}, true},
		{"? / 5 is valid", fields{ColorUnspecified, NumberFive}, true},
		{"? / ? is invalid", fields{ColorUnspecified, NumberUnspecified}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hint{
				Color:  tt.fields.Color,
				Number: tt.fields.Number,
			}
			if got := h.Valid(); got != tt.want {
				t.Errorf("Hint.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardColor_Valid(t *testing.T) {
	tests := []struct {
		name string
		c    CardColor
		want bool
	}{
		{"? is invalid", ColorUnspecified, false},
		{"Blue is valid", ColorBlue, true},
		{"Green is valid", ColorGreen, true},
		{"Red is valid", ColorRed, true},
		{"White is valid", ColorWhite, true},
		{"Yellow is valid", ColorYellow, true},
		{"Rainbow is valid", ColorRainbow, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Valid(); got != tt.want {
				t.Errorf("CardColor.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCardNumber_Valid(t *testing.T) {
	tests := []struct {
		name string
		c    CardNumber
		want bool
	}{
		{"One is invalid", NumberUnspecified, false},
		{"One is valid", NumberOne, true},
		{"Two is valid", NumberTwo, true},
		{"Three is valid", NumberThree, true},
		{"Four is valid", NumberFour, true},
		{"Five is valid", NumberFive, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.c.Valid(); got != tt.want {
				t.Errorf("CardNumber.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestForEachValidCard(t *testing.T) {
	var count int
	ForEachValidCard(func(got Card) {
		if count >= len(testForeachValidCardCards) {
			t.Fatal("ForeachValidCard: Too many cards")
		}

		want := testForeachValidCardCards[count]

		if !reflect.DeepEqual(got, want) {
			t.Errorf("ForeachValidCard: got %d = %v, want = %v", count, got, want)
		}

		count++
	})
}

// all the cards that are expected to be valid
// Used by TestForEachValidCard.
var testForeachValidCardCards = []Card{
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberTwo},
	{Color: ColorBlue, Number: NumberThree},
	{Color: ColorBlue, Number: NumberFour},
	{Color: ColorBlue, Number: NumberFive},

	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberTwo},
	{Color: ColorGreen, Number: NumberThree},
	{Color: ColorGreen, Number: NumberFour},
	{Color: ColorGreen, Number: NumberFive},

	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberTwo},
	{Color: ColorRed, Number: NumberThree},
	{Color: ColorRed, Number: NumberFour},
	{Color: ColorRed, Number: NumberFive},

	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberTwo},
	{Color: ColorWhite, Number: NumberThree},
	{Color: ColorWhite, Number: NumberFour},
	{Color: ColorWhite, Number: NumberFive},

	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberTwo},
	{Color: ColorYellow, Number: NumberThree},
	{Color: ColorYellow, Number: NumberFour},
	{Color: ColorYellow, Number: NumberFive},

	{Color: ColorRainbow, Number: NumberOne},
	{Color: ColorRainbow, Number: NumberTwo},
	{Color: ColorRainbow, Number: NumberThree},
	{Color: ColorRainbow, Number: NumberFour},
	{Color: ColorRainbow, Number: NumberFive},
}

func TestCard_Legal(t *testing.T) {
	type fields struct {
		Color  CardColor
		Number CardNumber
	}
	type args struct {
		mode GameMode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Blue 1 is legal in FiveColor", fields{ColorBlue, NumberOne}, args{ModeFiveColor}, true},
		{"Blue 1 is legal in SixColor", fields{ColorBlue, NumberOne}, args{ModeSixColor}, true},
		{"Blue 1 is legal in Rainbow", fields{ColorBlue, NumberOne}, args{ModeRainbow}, true},
		{"Blue 1 is legal in DarkRainbow", fields{ColorBlue, NumberOne}, args{ModeDarkRainbow}, true},

		{"Rainbow 1 is illegal in FiveColor", fields{ColorRainbow, NumberOne}, args{ModeFiveColor}, false},
		{"Rainbow 1 is legal in SixColor", fields{ColorRainbow, NumberOne}, args{ModeSixColor}, true},
		{"Rainbow 1 is legal in Rainbow", fields{ColorRainbow, NumberOne}, args{ModeRainbow}, true},
		{"Rainbow 1 is legal in DarkRainbow", fields{ColorRainbow, NumberOne}, args{ModeDarkRainbow}, true},

		{"? 1 is is illegal in FiveColor", fields{ColorUnspecified, NumberOne}, args{ModeFiveColor}, false},
		{"? 1 is is illegal in SixColor", fields{ColorUnspecified, NumberOne}, args{ModeSixColor}, false},
		{"? 1 is is illegal in Rainbow", fields{ColorUnspecified, NumberOne}, args{ModeRainbow}, false},
		{"? 1 is is illegal in DarkRainbow", fields{ColorUnspecified, NumberOne}, args{ModeDarkRainbow}, false},

		{"Blue ? is is illegal in FiveColor", fields{ColorBlue, NumberUnspecified}, args{ModeFiveColor}, false},
		{"Blue ? is is illegal in SixColor", fields{ColorBlue, NumberUnspecified}, args{ModeSixColor}, false},
		{"Blue ? is is illegal in Rainbow", fields{ColorBlue, NumberUnspecified}, args{ModeRainbow}, false},
		{"Blue ? is is illegal in DarkRainbow", fields{ColorBlue, NumberUnspecified}, args{ModeDarkRainbow}, false},

		{"? ? is is illegal in FiveColor", fields{ColorUnspecified, NumberUnspecified}, args{ModeFiveColor}, false},
		{"? ? is is illegal in SixColor", fields{ColorUnspecified, NumberUnspecified}, args{ModeSixColor}, false},
		{"? ? is is illegal in Rainbow", fields{ColorUnspecified, NumberUnspecified}, args{ModeRainbow}, false},
		{"? ? is is illegal in DarkRainbow", fields{ColorUnspecified, NumberUnspecified}, args{ModeDarkRainbow}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := Card{
				Color:  tt.fields.Color,
				Number: tt.fields.Number,
			}
			if got := c.Legal(tt.args.mode); got != tt.want {
				t.Errorf("Card.Legal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHint_IsNumberHint(t *testing.T) {
	type fields struct {
		Color  CardColor
		Number CardNumber
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Blue / 1 is not a number hint", fields{ColorBlue, NumberOne}, false},
		{"Blue / 2 is not a number hint", fields{ColorBlue, NumberTwo}, false},
		{"Blue / 3 is not a number hint", fields{ColorBlue, NumberThree}, false},
		{"Blue / 4 is not a number hint", fields{ColorBlue, NumberFour}, false},
		{"Blue / 5 is not a number hint", fields{ColorBlue, NumberFive}, false},
		{"Blue / ? is not a number hint", fields{ColorBlue, NumberUnspecified}, false},

		{"Green / 1 is not a number hint", fields{ColorGreen, NumberOne}, false},
		{"Green / 2 is not a number hint", fields{ColorGreen, NumberTwo}, false},
		{"Green / 3 is not a number hint", fields{ColorGreen, NumberThree}, false},
		{"Green / 4 is not a number hint", fields{ColorGreen, NumberFour}, false},
		{"Green / 5 is not a number hint", fields{ColorGreen, NumberFive}, false},
		{"Green / ? is not a number hint", fields{ColorGreen, NumberUnspecified}, false},

		{"Red / 1 is not a number hint", fields{ColorRed, NumberOne}, false},
		{"Red / 2 is not a number hint", fields{ColorRed, NumberTwo}, false},
		{"Red / 3 is not a number hint", fields{ColorRed, NumberThree}, false},
		{"Red / 4 is not a number hint", fields{ColorRed, NumberFour}, false},
		{"Red / 5 is not a number hint", fields{ColorRed, NumberFive}, false},
		{"Red / ? is not a number hint", fields{ColorRed, NumberUnspecified}, false},

		{"White / 1 is not a number hint", fields{ColorWhite, NumberOne}, false},
		{"White / 2 is not a number hint", fields{ColorWhite, NumberTwo}, false},
		{"White / 3 is not a number hint", fields{ColorWhite, NumberThree}, false},
		{"White / 4 is not a number hint", fields{ColorWhite, NumberFour}, false},
		{"White / 5 is not a number hint", fields{ColorWhite, NumberFive}, false},
		{"White / ? is not a number hint", fields{ColorWhite, NumberUnspecified}, false},

		{"Yellow / 1 is not a number hint ", fields{ColorYellow, NumberOne}, false},
		{"Yellow / 2 is not a number hint ", fields{ColorYellow, NumberTwo}, false},
		{"Yellow / 3 is not a number hint ", fields{ColorYellow, NumberThree}, false},
		{"Yellow / 4 is not a number hint ", fields{ColorYellow, NumberFour}, false},
		{"Yellow / 5 is not a number hint ", fields{ColorYellow, NumberFive}, false},
		{"Yellow / ? is not a number hint ", fields{ColorYellow, NumberUnspecified}, false},

		{"Rainbow / 1 is not a number hint", fields{ColorRainbow, NumberOne}, false},
		{"Rainbow / 2 is not a number hint", fields{ColorRainbow, NumberTwo}, false},
		{"Rainbow / 3 is not a number hint", fields{ColorRainbow, NumberThree}, false},
		{"Rainbow / 4 is not a number hint", fields{ColorRainbow, NumberFour}, false},
		{"Rainbow / 5 is not a number hint", fields{ColorRainbow, NumberFive}, false},
		{"Rainbow / ? is not a number hint", fields{ColorRainbow, NumberUnspecified}, false},

		{"? / 1 is a number hint", fields{ColorUnspecified, NumberOne}, true},
		{"? / 2 is a number hint", fields{ColorUnspecified, NumberTwo}, true},
		{"? / 3 is a number hint", fields{ColorUnspecified, NumberThree}, true},
		{"? / 4 is a number hint", fields{ColorUnspecified, NumberFour}, true},
		{"? / 5 is a number hint", fields{ColorUnspecified, NumberFive}, true},
		{"? / ? is not a number hint", fields{ColorUnspecified, NumberUnspecified}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hint{
				Color:  tt.fields.Color,
				Number: tt.fields.Number,
			}
			if got := h.IsNumberHint(); got != tt.want {
				t.Errorf("Hint.IsNumberHint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHint_IsColorHint(t *testing.T) {
	type fields struct {
		Color  CardColor
		Number CardNumber
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{"Blue / 1 is not a color hint", fields{ColorBlue, NumberOne}, false},
		{"Blue / 2 is not a color hint", fields{ColorBlue, NumberTwo}, false},
		{"Blue / 3 is not a color hint", fields{ColorBlue, NumberThree}, false},
		{"Blue / 4 is not a color hint", fields{ColorBlue, NumberFour}, false},
		{"Blue / 5 is not a color hint", fields{ColorBlue, NumberFive}, false},
		{"Blue / ? is a color hint", fields{ColorBlue, NumberUnspecified}, true},

		{"Green / 1 is not a color hint", fields{ColorGreen, NumberOne}, false},
		{"Green / 2 is not a color hint", fields{ColorGreen, NumberTwo}, false},
		{"Green / 3 is not a color hint", fields{ColorGreen, NumberThree}, false},
		{"Green / 4 is not a color hint", fields{ColorGreen, NumberFour}, false},
		{"Green / 5 is not a color hint", fields{ColorGreen, NumberFive}, false},
		{"Green / ? is a color hint", fields{ColorGreen, NumberUnspecified}, true},

		{"Red / 1 is not a color hint", fields{ColorRed, NumberOne}, false},
		{"Red / 2 is not a color hint", fields{ColorRed, NumberTwo}, false},
		{"Red / 3 is not a color hint", fields{ColorRed, NumberThree}, false},
		{"Red / 4 is not a color hint", fields{ColorRed, NumberFour}, false},
		{"Red / 5 is not a color hint", fields{ColorRed, NumberFive}, false},
		{"Red / ? is a color hint", fields{ColorRed, NumberUnspecified}, true},

		{"White / 1 is not a color hint", fields{ColorWhite, NumberOne}, false},
		{"White / 2 is not a color hint", fields{ColorWhite, NumberTwo}, false},
		{"White / 3 is not a color hint", fields{ColorWhite, NumberThree}, false},
		{"White / 4 is not a color hint", fields{ColorWhite, NumberFour}, false},
		{"White / 5 is not a color hint", fields{ColorWhite, NumberFive}, false},
		{"White / ? is a color hint", fields{ColorWhite, NumberUnspecified}, true},

		{"Yellow / 1 is not a color hint ", fields{ColorYellow, NumberOne}, false},
		{"Yellow / 2 is not a color hint ", fields{ColorYellow, NumberTwo}, false},
		{"Yellow / 3 is not a color hint ", fields{ColorYellow, NumberThree}, false},
		{"Yellow / 4 is not a color hint ", fields{ColorYellow, NumberFour}, false},
		{"Yellow / 5 is not a color hint ", fields{ColorYellow, NumberFive}, false},
		{"Yellow / ? is a color hint ", fields{ColorYellow, NumberUnspecified}, true},

		{"Rainbow / 1 is not a color hint", fields{ColorRainbow, NumberOne}, false},
		{"Rainbow / 2 is not a color hint", fields{ColorRainbow, NumberTwo}, false},
		{"Rainbow / 3 is not a color hint", fields{ColorRainbow, NumberThree}, false},
		{"Rainbow / 4 is not a color hint", fields{ColorRainbow, NumberFour}, false},
		{"Rainbow / 5 is not a color hint", fields{ColorRainbow, NumberFive}, false},
		{"Rainbow / ? is a color hint", fields{ColorRainbow, NumberUnspecified}, true},

		{"? / 1 is not a color hint", fields{ColorUnspecified, NumberOne}, false},
		{"? / 2 is not a color hint", fields{ColorUnspecified, NumberTwo}, false},
		{"? / 3 is not a color hint", fields{ColorUnspecified, NumberThree}, false},
		{"? / 4 is not a color hint", fields{ColorUnspecified, NumberFour}, false},
		{"? / 5 is not a color hint", fields{ColorUnspecified, NumberFive}, false},
		{"? / ? is not a color hint", fields{ColorUnspecified, NumberUnspecified}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hint{
				Color:  tt.fields.Color,
				Number: tt.fields.Number,
			}
			if got := h.IsColorHint(); got != tt.want {
				t.Errorf("Hint.IsColorHint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHint_Legal(t *testing.T) {
	type fields struct {
		Color  CardColor
		Number CardNumber
	}
	type args struct {
		mode GameMode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{"Blue 1 is illegal in FiveColor", fields{ColorBlue, NumberOne}, args{ModeFiveColor}, false},
		{"Blue 1 is illegal in SixColor", fields{ColorBlue, NumberOne}, args{ModeSixColor}, false},
		{"Blue 1 is illegal in Rainbow", fields{ColorBlue, NumberOne}, args{ModeRainbow}, false},
		{"Blue 1 is illegal in DarkRainbow", fields{ColorBlue, NumberOne}, args{ModeDarkRainbow}, false},

		{"Rainbow 1 is illegal in FiveColor", fields{ColorRainbow, NumberOne}, args{ModeFiveColor}, false},
		{"Rainbow 1 is illegal in SixColor", fields{ColorRainbow, NumberOne}, args{ModeSixColor}, false},
		{"Rainbow 1 is illegal in Rainbow", fields{ColorRainbow, NumberOne}, args{ModeRainbow}, false},
		{"Rainbow 1 is illegal in DarkRainbow", fields{ColorRainbow, NumberOne}, args{ModeDarkRainbow}, false},

		{"? 1 is is legal in FiveColor", fields{ColorUnspecified, NumberOne}, args{ModeFiveColor}, true},
		{"? 1 is is legal in SixColor", fields{ColorUnspecified, NumberOne}, args{ModeSixColor}, true},
		{"? 1 is is legal in Rainbow", fields{ColorUnspecified, NumberOne}, args{ModeRainbow}, true},
		{"? 1 is is legal in DarkRainbow", fields{ColorUnspecified, NumberOne}, args{ModeDarkRainbow}, true},

		{"Blue ? is is legal in FiveColor", fields{ColorBlue, NumberUnspecified}, args{ModeFiveColor}, true},
		{"Blue ? is is legal in SixColor", fields{ColorBlue, NumberUnspecified}, args{ModeSixColor}, true},
		{"Blue ? is is legal in Rainbow", fields{ColorBlue, NumberUnspecified}, args{ModeRainbow}, true},
		{"Blue ? is is legal in DarkRainbow", fields{ColorBlue, NumberUnspecified}, args{ModeDarkRainbow}, true},

		{"Rainbow ? is is illegal in FiveColor", fields{ColorRainbow, NumberUnspecified}, args{ModeFiveColor}, false},
		{"Rainbow ? is is legal in SixColor", fields{ColorRainbow, NumberUnspecified}, args{ModeSixColor}, true},
		{"Rainbow ? is is illegal in Rainbow", fields{ColorRainbow, NumberUnspecified}, args{ModeRainbow}, false},
		{"Rainbow ? is is illegal in DarkRainbow", fields{ColorRainbow, NumberUnspecified}, args{ModeDarkRainbow}, false},

		{"? ? is is illegal in FiveColor", fields{ColorUnspecified, NumberUnspecified}, args{ModeFiveColor}, false},
		{"? ? is is illegal in SixColor", fields{ColorUnspecified, NumberUnspecified}, args{ModeSixColor}, false},
		{"? ? is is illegal in Rainbow", fields{ColorUnspecified, NumberUnspecified}, args{ModeRainbow}, false},
		{"? ? is is illegal in DarkRainbow", fields{ColorUnspecified, NumberUnspecified}, args{ModeDarkRainbow}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hint{
				Color:  tt.fields.Color,
				Number: tt.fields.Number,
			}
			if got := h.Legal(tt.args.mode); got != tt.want {
				t.Errorf("Hint.Legal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHint_Matches(t *testing.T) {
	type fields struct {
		Color  CardColor
		Number CardNumber
	}
	type args struct {
		c    Card
		mode GameMode
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// non-rainbow color hints

		{"FiveColor Blue hint applies to blue card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorBlue, NumberOne}, ModeFiveColor}, true},
		{"FiveColor Blue hint does not apply to red card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorRed, NumberOne}, ModeFiveColor}, false},
		{"FiveColor Blue hint does not apply to rainbow card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorRainbow, NumberOne}, ModeFiveColor}, false},

		{"SixColor Blue hint applies to blue card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorBlue, NumberOne}, ModeSixColor}, true},
		{"SixColor Blue hint does not apply to red card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorRed, NumberOne}, ModeSixColor}, false},
		{"SixColor Blue hint does not apply to rainbow card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorRainbow, NumberOne}, ModeSixColor}, false},

		{"Rainbow Blue hint applies to blue card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorBlue, NumberOne}, ModeRainbow}, true},
		{"Rainbow Blue hint does not apply to red card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorRed, NumberOne}, ModeRainbow}, false},
		{"Rainbow Blue hint applies to rainbow card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorRainbow, NumberOne}, ModeRainbow}, true},

		{"DarkRainbow Blue hint applies to blue card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorBlue, NumberOne}, ModeDarkRainbow}, true},
		{"DarkRainbow Blue hint does not apply to red card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorRed, NumberOne}, ModeDarkRainbow}, false},
		{"DarkRainbow Blue hint applies to rainbow card", fields{ColorBlue, NumberUnspecified}, args{Card{ColorRainbow, NumberOne}, ModeDarkRainbow}, true},

		// rainbow color hint, only legal in SixColor

		{"SixColor Rainbow hint does not apply to blue card", fields{ColorRainbow, NumberUnspecified}, args{Card{ColorBlue, NumberOne}, ModeSixColor}, false},
		{"SixColor Rainbow hint does not apply to red card", fields{ColorRainbow, NumberUnspecified}, args{Card{ColorRed, NumberOne}, ModeSixColor}, false},
		{"SixColor Rainbow hint applies to rainbow card", fields{ColorRainbow, NumberUnspecified}, args{Card{ColorRainbow, NumberOne}, ModeSixColor}, true},

		// number hint
		{"FiveColor One hint applies to 1", fields{ColorUnspecified, NumberOne}, args{Card{ColorBlue, NumberOne}, ModeFiveColor}, true},
		{"FiveColor One hint does not apply to 2", fields{ColorUnspecified, NumberOne}, args{Card{ColorBlue, NumberTwo}, ModeFiveColor}, false},

		{"SixColor One hint applies to 1", fields{ColorUnspecified, NumberOne}, args{Card{ColorBlue, NumberOne}, ModeSixColor}, true},
		{"SixColor One hint does not apply to 2", fields{ColorUnspecified, NumberOne}, args{Card{ColorBlue, NumberTwo}, ModeSixColor}, false},

		{"Rainbow One hint applies to 1", fields{ColorUnspecified, NumberOne}, args{Card{ColorBlue, NumberOne}, ModeRainbow}, true},
		{"Rainbow One hint does not apply to 2", fields{ColorUnspecified, NumberOne}, args{Card{ColorBlue, NumberTwo}, ModeRainbow}, false},

		{"DarkRainbow One hint applies to 1", fields{ColorUnspecified, NumberOne}, args{Card{ColorBlue, NumberOne}, ModeDarkRainbow}, true},
		{"DarkRainbow One hint does not apply to 2", fields{ColorUnspecified, NumberOne}, args{Card{ColorBlue, NumberTwo}, ModeDarkRainbow}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := Hint{
				Color:  tt.fields.Color,
				Number: tt.fields.Number,
			}
			if got := h.Matches(tt.args.c, tt.args.mode); got != tt.want {
				t.Errorf("Hint.Matches() = %v, want %v", got, tt.want)
			}
		})
	}
}
