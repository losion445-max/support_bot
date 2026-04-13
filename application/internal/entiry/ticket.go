package entiry

type Ticket struct {
	AdminID        string
	UserID         string
	AdminMessageID string
}

func NewTicket(adminID, userID, AdminMessageID string) Ticket {
	return Ticket{
		AdminID:        adminID,
		UserID:         userID,
		AdminMessageID: adminID,
	}
}
