package booking

import "go.mongodb.org/mongo-driver/bson/primitive"

type Booking struct {
	ID           primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	CustomerName string             `json:"customerName"`
	Room         uint8              `json:"room"`
	StartDate    primitive.DateTime `json:"startDate"`
	DueDate      primitive.DateTime `json:"dueDate"`
	BookingType  string             `json:"bookingType"`
	Contact      string             `json:"contact"`
	People       uint8              `json:"people"`
}
