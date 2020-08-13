package db

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGettingPopulation(t *testing.T) {

	var testCases = []struct {
		city           string
		wantPopulation int
	}{
		{
			city:           "Tokyo",
			wantPopulation: 33200000,
		},
		{
			city:           "New York",
			wantPopulation: 17800000,
		},
		{
			city:           "Sao Paulo",
			wantPopulation: 17700000,
		},
		{
			city:           "Seoul",
			wantPopulation: 17500000,
		},
		{
			city:           "Mexico City",
			wantPopulation: 17400000,
		},
		{
			city:           "Osaka",
			wantPopulation: 16425000,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.city, func(t *testing.T) {
			singleton := GetSingletonDb()
			got := singleton.GetPopulation(tC.city)
			assert.EqualValues(t, tC.wantPopulation, got)
		})
	}
}

var totalPopulationTestCases = []struct {
	cities         []string
	wantPopulation int
}{
	{
		cities:         []string{"Tokyo", "New York"},
		wantPopulation: 33200000 + 17800000,
	},
	{
		cities:         []string{"New York", "Sao Paulo"},
		wantPopulation: 17800000 + 17700000,
	},
	{
		cities:         []string{"Sao Paulo", "Seoul"},
		wantPopulation: 17700000 + 17500000,
	},
	{
		cities:         []string{"Seoul", "Mexico City"},
		wantPopulation: 17500000 + 17400000,
	},
	{
		cities:         []string{"Mexico City", "Osaka"},
		wantPopulation: 17400000 + 16425000,
	},
	{
		cities:         []string{"Osaka", "Tokyo", "New York"},
		wantPopulation: 16425000 + 33200000 + 17800000,
	},
	{
		cities:         []string{},
		wantPopulation: 0,
	},
}

func TestGetTotalPopulation(t *testing.T) {
	for _, tC := range totalPopulationTestCases {
		desc := strings.Join(tC.cities, ", ")
		t.Run(desc, func(t *testing.T) {
			got := GetTotalPopulation(tC.cities)
			assert.EqualValues(t, tC.wantPopulation, got)
		})
	}
}

func TestGetTotalPopulationEx(t *testing.T) {
	for _, tC := range totalPopulationTestCases {
		desc := strings.Join(tC.cities, ", ")
		t.Run(desc, func(t *testing.T) {
			got := GetTotalPopulationEx(GetSingletonDb(), tC.cities)
			assert.EqualValues(t, tC.wantPopulation, got)
		})
	}
}
