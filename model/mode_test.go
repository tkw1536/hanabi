package model

import (
	"reflect"
	"testing"
)

func TestGameMode_Valid(t *testing.T) {
	tests := []struct {
		name string
		mode GameMode
		want bool
	}{
		{"FiveColor is valid", ModeFiveColor, true},
		{"SixColor is valid", ModeSixColor, true},
		{"Rainbow is valid", ModeRainbow, true},
		{"DarkRainbow is valid", ModeDarkRainbow, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mode.Valid(); got != tt.want {
				t.Errorf("GameMode.Valid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameMode_Count(t *testing.T) {
	type args struct {
		card Card
	}
	tests := []struct {
		name string
		mode GameMode
		args args
		want int
	}{
		{"#{Blue 1} in FiveColor == 3", ModeFiveColor, args{Card{ColorBlue, NumberOne}}, 3},
		{"#{Blue 2} in FiveColor == 2", ModeFiveColor, args{Card{ColorBlue, NumberTwo}}, 2},
		{"#{Blue 5} in FiveColor == 1", ModeFiveColor, args{Card{ColorBlue, NumberFive}}, 1},
		{"#{Rainbow 1} in FiveColor == 0", ModeFiveColor, args{Card{ColorRainbow, NumberOne}}, 0},
		{"#{Rainbow 2} in FiveColor == 0", ModeFiveColor, args{Card{ColorRainbow, NumberTwo}}, 0},
		{"#{Rainbow 5} in FiveColor == 0", ModeFiveColor, args{Card{ColorRainbow, NumberFive}}, 0},

		{"#{Blue 1} in SixColor == 3", ModeSixColor, args{Card{ColorBlue, NumberOne}}, 3},
		{"#{Blue 2} in SixColor == 2", ModeSixColor, args{Card{ColorBlue, NumberTwo}}, 2},
		{"#{Blue 5} in SixColor == 1", ModeSixColor, args{Card{ColorBlue, NumberFive}}, 1},
		{"#{Rainbow 1} in SixColor == 3", ModeSixColor, args{Card{ColorRainbow, NumberOne}}, 3},
		{"#{Rainbow 2} in SixColor == 2", ModeSixColor, args{Card{ColorRainbow, NumberTwo}}, 2},
		{"#{Rainbow 5} in SixColor == 1", ModeSixColor, args{Card{ColorRainbow, NumberFive}}, 1},

		{"#{Blue 1} in Rainbow == 3", ModeRainbow, args{Card{ColorBlue, NumberOne}}, 3},
		{"#{Blue 2} in Rainbow == 2", ModeRainbow, args{Card{ColorBlue, NumberTwo}}, 2},
		{"#{Blue 5} in Rainbow == 1", ModeRainbow, args{Card{ColorBlue, NumberFive}}, 1},
		{"#{Rainbow 1} in Rainbow == 3", ModeRainbow, args{Card{ColorRainbow, NumberOne}}, 3},
		{"#{Rainbow 2} in Rainbow == 2", ModeRainbow, args{Card{ColorRainbow, NumberTwo}}, 2},
		{"#{Rainbow 5} in Rainbow == 1", ModeRainbow, args{Card{ColorRainbow, NumberFive}}, 1},

		{"#{Blue 1} in DarkRainbow == 3", ModeDarkRainbow, args{Card{ColorBlue, NumberOne}}, 3},
		{"#{Blue 2} in DarkRainbow == 2", ModeDarkRainbow, args{Card{ColorBlue, NumberTwo}}, 2},
		{"#{Blue 5} in DarkRainbow == 1", ModeDarkRainbow, args{Card{ColorBlue, NumberFive}}, 1},
		{"#{Rainbow 1} in DarkRainbow == 1", ModeDarkRainbow, args{Card{ColorRainbow, NumberOne}}, 1},
		{"#{Rainbow 2} in DarkRainbow == 1", ModeDarkRainbow, args{Card{ColorRainbow, NumberTwo}}, 1},
		{"#{Rainbow 5} in DarkRainbow == 1", ModeDarkRainbow, args{Card{ColorRainbow, NumberFive}}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mode.Count(tt.args.card); got != tt.want {
				t.Errorf("GameMode.Count() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGameMode_TotalCards(t *testing.T) {
	tests := []struct {
		name      string
		mode      GameMode
		wantTotal int
	}{
		{"FiveColor has 50 cards", ModeFiveColor, 50},
		{"SixColor has 60 cards", ModeSixColor, 60},
		{"Rainbow has 60 cards", ModeRainbow, 60},
		{"DarkRainbow has 55 cards", ModeDarkRainbow, 55},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotTotal := tt.mode.TotalCards(); gotTotal != tt.wantTotal {
				t.Errorf("GameMode.TotalCards() = %v, want %v", gotTotal, tt.wantTotal)
			}
		})
	}
}

func TestGameMode_NewStack(t *testing.T) {
	tests := []struct {
		name string
		mode GameMode
		want []Card
	}{
		{"FiveColor", ModeFiveColor, testFiveColorStack},
		{"SixColor", ModeSixColor, testSixColorStack},
		{"Rainbow", ModeRainbow, testSixColorStack},
		{"DarkRainbow", ModeDarkRainbow, testDarkRainbowStack},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.mode.NewStack(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GameMode.NewStack() = %v, want %v", got, tt.want)
			}
		})
	}
}

// stack for the FiveColor Mode
var testFiveColorStack = []Card{
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberTwo},
	{Color: ColorBlue, Number: NumberTwo},
	{Color: ColorBlue, Number: NumberThree},
	{Color: ColorBlue, Number: NumberThree},
	{Color: ColorBlue, Number: NumberFour},
	{Color: ColorBlue, Number: NumberFour},
	{Color: ColorBlue, Number: NumberFive},

	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberTwo},
	{Color: ColorGreen, Number: NumberTwo},
	{Color: ColorGreen, Number: NumberThree},
	{Color: ColorGreen, Number: NumberThree},
	{Color: ColorGreen, Number: NumberFour},
	{Color: ColorGreen, Number: NumberFour},
	{Color: ColorGreen, Number: NumberFive},

	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberTwo},
	{Color: ColorRed, Number: NumberTwo},
	{Color: ColorRed, Number: NumberThree},
	{Color: ColorRed, Number: NumberThree},
	{Color: ColorRed, Number: NumberFour},
	{Color: ColorRed, Number: NumberFour},
	{Color: ColorRed, Number: NumberFive},

	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberTwo},
	{Color: ColorWhite, Number: NumberTwo},
	{Color: ColorWhite, Number: NumberThree},
	{Color: ColorWhite, Number: NumberThree},
	{Color: ColorWhite, Number: NumberFour},
	{Color: ColorWhite, Number: NumberFour},
	{Color: ColorWhite, Number: NumberFive},

	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberTwo},
	{Color: ColorYellow, Number: NumberTwo},
	{Color: ColorYellow, Number: NumberThree},
	{Color: ColorYellow, Number: NumberThree},
	{Color: ColorYellow, Number: NumberFour},
	{Color: ColorYellow, Number: NumberFour},
	{Color: ColorYellow, Number: NumberFive},
}

// stack for the SixColor and Rainbow modes
var testSixColorStack = []Card{
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberTwo},
	{Color: ColorBlue, Number: NumberTwo},
	{Color: ColorBlue, Number: NumberThree},
	{Color: ColorBlue, Number: NumberThree},
	{Color: ColorBlue, Number: NumberFour},
	{Color: ColorBlue, Number: NumberFour},
	{Color: ColorBlue, Number: NumberFive},

	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberTwo},
	{Color: ColorGreen, Number: NumberTwo},
	{Color: ColorGreen, Number: NumberThree},
	{Color: ColorGreen, Number: NumberThree},
	{Color: ColorGreen, Number: NumberFour},
	{Color: ColorGreen, Number: NumberFour},
	{Color: ColorGreen, Number: NumberFive},

	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberTwo},
	{Color: ColorRed, Number: NumberTwo},
	{Color: ColorRed, Number: NumberThree},
	{Color: ColorRed, Number: NumberThree},
	{Color: ColorRed, Number: NumberFour},
	{Color: ColorRed, Number: NumberFour},
	{Color: ColorRed, Number: NumberFive},

	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberTwo},
	{Color: ColorWhite, Number: NumberTwo},
	{Color: ColorWhite, Number: NumberThree},
	{Color: ColorWhite, Number: NumberThree},
	{Color: ColorWhite, Number: NumberFour},
	{Color: ColorWhite, Number: NumberFour},
	{Color: ColorWhite, Number: NumberFive},

	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberTwo},
	{Color: ColorYellow, Number: NumberTwo},
	{Color: ColorYellow, Number: NumberThree},
	{Color: ColorYellow, Number: NumberThree},
	{Color: ColorYellow, Number: NumberFour},
	{Color: ColorYellow, Number: NumberFour},
	{Color: ColorYellow, Number: NumberFive},

	{Color: ColorRainbow, Number: NumberOne},
	{Color: ColorRainbow, Number: NumberOne},
	{Color: ColorRainbow, Number: NumberOne},
	{Color: ColorRainbow, Number: NumberTwo},
	{Color: ColorRainbow, Number: NumberTwo},
	{Color: ColorRainbow, Number: NumberThree},
	{Color: ColorRainbow, Number: NumberThree},
	{Color: ColorRainbow, Number: NumberFour},
	{Color: ColorRainbow, Number: NumberFour},
	{Color: ColorRainbow, Number: NumberFive},
}

// expected stack for dark rainbow
var testDarkRainbowStack = []Card{
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberOne},
	{Color: ColorBlue, Number: NumberTwo},
	{Color: ColorBlue, Number: NumberTwo},
	{Color: ColorBlue, Number: NumberThree},
	{Color: ColorBlue, Number: NumberThree},
	{Color: ColorBlue, Number: NumberFour},
	{Color: ColorBlue, Number: NumberFour},
	{Color: ColorBlue, Number: NumberFive},

	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberOne},
	{Color: ColorGreen, Number: NumberTwo},
	{Color: ColorGreen, Number: NumberTwo},
	{Color: ColorGreen, Number: NumberThree},
	{Color: ColorGreen, Number: NumberThree},
	{Color: ColorGreen, Number: NumberFour},
	{Color: ColorGreen, Number: NumberFour},
	{Color: ColorGreen, Number: NumberFive},

	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberOne},
	{Color: ColorRed, Number: NumberTwo},
	{Color: ColorRed, Number: NumberTwo},
	{Color: ColorRed, Number: NumberThree},
	{Color: ColorRed, Number: NumberThree},
	{Color: ColorRed, Number: NumberFour},
	{Color: ColorRed, Number: NumberFour},
	{Color: ColorRed, Number: NumberFive},

	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberOne},
	{Color: ColorWhite, Number: NumberTwo},
	{Color: ColorWhite, Number: NumberTwo},
	{Color: ColorWhite, Number: NumberThree},
	{Color: ColorWhite, Number: NumberThree},
	{Color: ColorWhite, Number: NumberFour},
	{Color: ColorWhite, Number: NumberFour},
	{Color: ColorWhite, Number: NumberFive},

	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberOne},
	{Color: ColorYellow, Number: NumberTwo},
	{Color: ColorYellow, Number: NumberTwo},
	{Color: ColorYellow, Number: NumberThree},
	{Color: ColorYellow, Number: NumberThree},
	{Color: ColorYellow, Number: NumberFour},
	{Color: ColorYellow, Number: NumberFour},
	{Color: ColorYellow, Number: NumberFive},

	{Color: ColorRainbow, Number: NumberOne},
	{Color: ColorRainbow, Number: NumberTwo},
	{Color: ColorRainbow, Number: NumberThree},
	{Color: ColorRainbow, Number: NumberFour},
	{Color: ColorRainbow, Number: NumberFive},
}
