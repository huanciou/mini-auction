package schemas

import "google.golang.org/genproto/googleapis/type/date"

type Lot struct {
	Id      int
	Name    string
	Price   int
	Bidder  string
	StartAt date.Date
	EndAt   date.Date
}
