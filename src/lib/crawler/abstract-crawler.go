package crawler

import "github.com/adlandh/termin-berlinweit-suchen/src/lib/misc"

type AbstractCrawler interface {
	GetTerminUrl(string) (string, error)
	CheckCalendar(string) (misc.MonthsMap, error)
}
