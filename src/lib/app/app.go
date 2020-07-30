package app

import (
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/config"
	"github.com/adlandh/termin-berlinweit-suchen/src/lib/crawler"
	"github.com/adlandh/termin-berlinweit-suchen/src/lib/misc"
)

type App struct {
	config  config.AbstractConfigProvider
	crawler crawler.AbstractCrawler
	err     chan error
	done    chan struct{}
}

func NewApp(config config.AbstractConfigProvider, crawler crawler.AbstractCrawler, err chan error,
	done chan struct{}) *App {
	return &App{
		config:  config,
		crawler: crawler,
		err:     err,
		done:    done,
	}
}

func (a *App) Run() {
	terminUrl := a.crawler.GetTerminUrl(a.config.GetMainUrl())
	months := a.crawler.CheckCalendar(terminUrl)
	a.printDates(a.convertAndSortMonths(months))
}

func (a *App) convertAndSortMonths(months misc.MonthsMap) misc.Months {
	germanMonthNames := map[string]int{
		"Januar":    1,
		"Februar":   2,
		"März":      3,
		"April":     4,
		"Mai":       5,
		"Juni":      6,
		"Juli":      7,
		"August":    8,
		"September": 9,
		"Oktober":   10,
		"November":  11,
		"Dezember":  12,
	}
	var newMonths misc.Months
	var monthsSlice []string
	for month := range months {
		monthsSlice = append(monthsSlice, month)
	}
	sort.Slice(monthsSlice, func(i, j int) bool {
		monthI := strings.Split(monthsSlice[i], " ")
		monthJ := strings.Split(monthsSlice[j], " ")
		if monthI[1] < monthJ[1] || (monthI[1] == monthJ[1] && germanMonthNames[monthI[0]] < germanMonthNames[monthJ[0]]) {
			return true
		}

		return false
	})

	for _, month := range monthsSlice {
		if len(months[month]) > 0 {
			newMonths = append(newMonths, misc.Month{
				Title: month,
				Dates: a.convertAndSortDates(months[month]),
			})
		}
	}

	return newMonths
}

func (a *App) convertAndSortDates(month misc.MonthMap) misc.Dates {
	var newMonth misc.Dates
	var datesSlice []string
	for date := range month {
		datesSlice = append(datesSlice, date)
	}

	sort.Slice(datesSlice, func(i, j int) bool {
		return datesSlice[i] < datesSlice[j]
	})

	for _, date := range datesSlice {
		newMonth = append(newMonth, misc.Date{
			Title: date,
			Url:   month[date],
		})
	}

	return newMonth
}

func (a *App) printDates(months misc.Months) {
	if len(months) == 0 {
		if a.config.GetVerbose() {
			log.Println("No available dates found")
		}
		return
	}
	fmt.Println("Available dates by months: ")
	for _, month := range months {
		fmt.Println(month.Title + ":")
		for _, date := range month.Dates {
			fmt.Println(date.Title, "-", date.Url)
		}
	}
	a.done <- struct{}{}
}
