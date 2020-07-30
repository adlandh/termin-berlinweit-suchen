package misc

type MonthMap map[string]string

type MonthsMap map[string]MonthMap

type Date struct {
	Title string
	Url   string
}

type Dates []Date

type Month struct {
	Title string
	Dates Dates
}

type Months []Month
