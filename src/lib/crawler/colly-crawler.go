package crawler

import (
	"log"
	"sync"

	"github.com/adlandh/termin-berlinweit-suchen/src/lib/misc"
)
import "github.com/gocolly/colly"

type CollyCrawler struct {
	mutex     *sync.RWMutex
	collector *colly.Collector
	months    misc.MonthsMap
	err       chan error
	verbose   bool
}

func NewCollyCrawler(verbose bool, err chan error) *CollyCrawler {
	c := &CollyCrawler{collector: colly.NewCollector(), mutex: &sync.RWMutex{}, months: make(misc.MonthsMap), err: err, verbose: verbose}
	if c.verbose {
		c.collector.OnRequest(func(r *colly.Request) {
			log.Println("Visiting: ", r.URL.String())
		})
	}
	return c
}

func (c *CollyCrawler) GetTerminUrl(mainUrl string) string {
	var terminUrl string
	c.collector.OnHTML(misc.TerminUrlElement, func(element *colly.HTMLElement) {
		terminUrl = element.Request.AbsoluteURL(element.ChildAttr(misc.TerminButtonElement, "href"))
	})

	err := c.collector.Visit(mainUrl)
	if err != nil {
		c.err <- err
	}

	return terminUrl
}

func (c *CollyCrawler) CheckCalendar(terminUrl string) misc.MonthsMap {
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
		url := c.getAbsoluteUrl(element)
		if url == "" {
			return
		}
		err := c.collector.Visit(url)
		if err != nil {
			c.err <- err
		}
	})

	err := c.collector.Visit(terminUrl)
	if err != nil {
		c.err <- err
	}

	return c.months
}

func (c *CollyCrawler) getAbsoluteUrl(element *colly.HTMLElement) string {
	url := element.ChildAttr("a", "href")
	if url != "" {
		url = element.Request.AbsoluteURL(url)
	}

	return url
}
