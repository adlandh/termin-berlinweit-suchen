package crawler

import "github.com/adlandh/termin-berlinweit-suchen/src/lib/misc"

type AbstractCrawler interface {
	GetTerminUrl(string) string
	CheckCalendar(string) misc.MonthsMap
}
