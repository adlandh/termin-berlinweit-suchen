package crawler

import (
	"fmt"
	"sync"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/misc"
)
import "github.com/gocolly/colly"

type CollyCrawler struct {
	mutex     *sync.RWMutex
	collector *colly.Collector
	months    misc.MonthsMap
}

func NewCollyCrawler(verbose bool) *CollyCrawler {
	c := &CollyCrawler{collector: colly.NewCollector(), mutex: &sync.RWMutex{}, months: make(misc.MonthsMap)}
	if verbose {
		c.collector.OnRequest(func(r *colly.Request) {
			fmt.Println("Visiting: ", r.URL.String())
		})
	}
	return c
}

func (c *CollyCrawler) GetTerminUrl(mainUrl string) (terminUrl string, err error) {
	c.collector.OnHTML(misc.TerminUrlElement, func(element *colly.HTMLElement) {
		terminUrl = element.Request.AbsoluteURL(element.ChildAttr(misc.TerminButtonElement, "href"))
	})

	err = c.collector.Visit(mainUrl)

	return
}

func (c *CollyCrawler) CheckCalendar(terminUrl string) (misc.MonthsMap, error) {
	var err error
	c.collector.OnHTML(misc.TerminMonthTable, func(element *colly.HTMLElement) {
		month := element.ChildText(misc.TerminMonthHeader)
		c.mutex.Lock()
		if _, ok := c.months[month]; !ok {
			c.months[month] = make(misc.MonthMap)
			element.ForEach(misc.TerminDateAvailable, func(i int, element *colly.HTMLElement) {
				c.months[month][element.Text] = c.getAbsoluteUrl(element)
			})
		}
		c.mutex.Unlock()
	})

	c.collector.OnHTML(misc.TerminNextLinkElement, func(element *colly.HTMLElement) {
		err = c.collector.Visit(c.getAbsoluteUrl(element))
	})

	err = c.collector.Visit(terminUrl)

	return c.months, err
}

func (c *CollyCrawler) getAbsoluteUrl(element *colly.HTMLElement) string {
	url := element.ChildAttr("a", "href")
	if url != "" {
		url = element.Request.AbsoluteURL(url)
	}

	return url
}
