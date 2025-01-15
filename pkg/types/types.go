package types

// Money type
type Money int64

// Currency type
type Currency string

// Category type: car, chemist, food
type Category string

// Payment status
type Status string

// Payment status variables
const (
	StatusOk         Status = "OK"
	StatusFail       Status = "FAIL"
	StatusInProgress Status = "INPROGRESS"
)

// Currency code
const (
	TJS Currency = "TJS"
	RUB Currency = "RUB"
	USD Currency = "USD"
	EUR Currency = "EUR"
	CAD Currency = "CAD"
)

// PAN is card no.
type PAN string

// Card - information about payment card
type Card struct {
	ID         int
	PAN        PAN
	Balance    Money
	MinBalance Money
	Currency   Currency
	Color      string
	Name       string
	Active     bool
}

type Payment struct {
	ID       int
	Amount   Money
	Category Category
	Status   Status
}

type PaymentSource struct {
	Type    string // 'card'
	Number  PAN    // номер вида '5058 xxxx xxxx 8888'
	Balance Money  // баланс в дирамах
}
