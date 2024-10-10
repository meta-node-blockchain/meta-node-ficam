package model
type EmailOrder struct {
	ID       []byte
	HexID    string
	Customer string
	Products []struct {
		ID               []byte
		HexIdProduct     string
		Quantity         uint
		SUBCRIPTION_TYPE uint
		ProductName      string
		ImgUrl           string
	}
	CreateAt uint
	CreateAtDate string
	ShipInfo struct {
		FirstName  string
		LastName   string
		Email      string
		Country    string
		City       string
		State      string
		PostalCode string
		Phone      string
		Address    string
	}
	ShippingFee uint
}
type Data struct {
	Order EmailOrder
	PaymentOrder uint
}
type Sender	struct {
		Address string
		Subject string
} 
type Recipient struct {
	ToEmails  []string
	CcEmails  []string
	BccEmails []string
}
