package booking

type Booking struct {
	ID           string `json:"id"`
	CustomerName string `json:"customerName"`
	Room         uint8  `json:"room"`
	StartDate    string `json:"startDate"`
	DueDate      string `json:"dueDate"`
	BookingType  string `json:"bookingType"`
	Contact      string `json:"contact"`
}

var books []Booking
