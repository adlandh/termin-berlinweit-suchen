package app

import (
	"reflect"
	"testing"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/misc"
)

func TestApp_convertAndSortMonths(t *testing.T) {

	tests := []struct {
		name   string
		months misc.MonthsMap
		want   misc.Months
	}{
		{
			name: "Test Case #1",
			months: misc.MonthsMap{
				"November 2020": misc.MonthMap{
					"01": "-",
				},
				"November 2019": misc.MonthMap{
					"02": "-",
				},
				"Februar 2020": misc.MonthMap{
					"03": "-",
				},
			},
			want: misc.Months{
				misc.Month{
					Title: "November 2019",
					Dates: misc.Dates{
						misc.Date{
							Title: "02",
							Url:   "-",
						},
					},
				},
				misc.Month{
					Title: "Februar 2020",
					Dates: misc.Dates{
						misc.Date{
							Title: "03",
							Url:   "-",
						},
					},
				},
				misc.Month{
					Title: "November 2020",
					Dates: misc.Dates{
						misc.Date{
							Title: "01",
							Url:   "-",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{}
			if got := a.convertAndSortMonths(tt.months); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertAndSortMonths() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_convertAndSortDates(t *testing.T) {

	tests := []struct {
		name  string
		month misc.MonthMap
		want  misc.Dates
	}{
		{
			name: "Test Case #1",

			month: misc.MonthMap{
				"01": "-",
				"05": "-",
				"03": "-",
			},
			want: misc.Dates{
				misc.Date{
					Title: "01",
					Url:   "-",
				},
				misc.Date{
					Title: "03",
					Url:   "-",
				},
				misc.Date{
					Title: "05",
					Url:   "-",
				},
			},
		},
		{
			name: "Test Case #2",

			month: misc.MonthMap{
				"12": "-",
				"04": "-",
				"05": "-",
				"01": "-",
			},
			want: misc.Dates{
				misc.Date{
					Title: "01",
					Url:   "-",
				},
				misc.Date{
					Title: "04",
					Url:   "-",
				},
				misc.Date{
					Title: "05",
					Url:   "-",
				},
				misc.Date{
					Title: "12",
					Url:   "-",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &App{}
			if got := a.convertAndSortDates(tt.month); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("convertAndSortDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
