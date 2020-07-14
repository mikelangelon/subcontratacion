package company

type Company struct {
	ID            string
	Name          string
	CIF           string
	City          string
	CP            string
	Description   string
	Employees     string
	FundationYear int64
	Quantity      string
	Street        string
	Web           string
	Logo          []byte
	Contact       Contact
	Email         string
}
