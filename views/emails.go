package views

type Messages struct {
	// Name     string
	// Email       string
	// ForgotPasswordMessage string

	// Name        string `json:"name"`
	// Email       string `json:"email"`
	// PhoneNumber string `json:"phonenumber"`
	// Quantity    string `json:"quantity"`
	// Flavour     string `json:"flavor"`
	// Address     string `json:"address"`
	OrderMessage      string
	OrderResponse     string
	NoneReplyResponse string
	MessageSubject    string
}

/*New ... */
func (r *Messages) New() *Messages {
	t := new(Messages)

	t.OrderMessage = `Hello You Have A New Order From {FullName} \n\n\n
	Order Details\n
	=============================\n
	1. Product Type {Product}\n
	2. Quantity: {Quantity}\n
	3. Flavour: {Flavour}\n
	
	Client Details\n
	===============================\n
	1. Client Name: {FullName}\n
	2. Client Email Address: {Email}\n
	3. Client Phone Number: {PhoneNumber}\n
	4. Client Address: {Address}\n
	
	Regards
	Support Team @ KreationsByKgola`

	t.MessageSubject = `Order From {FullName}`

	t.OrderResponse = `Thank You For Placing Your Order With KreationsByKgola
	
	We Will Be In Contact With You As Soon As Possible
	
	Regards
	Support Team @ KreationsByKgola`

	t.NoneReplyResponse = `Order Confirmed, Please Do Not Respond To This Email`

	return t
}
