package crawler

import "github.com/adlandh/termin-berlinweit-suchen/src/lib/misc"

type AbstractCrawler interface {
	GetAppointmentURL(string) string
	CheckCalendar(string) misc.MonthsMap
}
