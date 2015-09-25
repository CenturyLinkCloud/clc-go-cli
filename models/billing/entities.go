package billing

type InvoiceData struct {
	Id                  string
	Terms               string
	CompanyName         string
	AccountAlias        string
	PricingAccountAlias string
	ParentAccountAlias  string
	Address1            string
	Address2            string
	City                string
	StateProvince       string
	PostalCode          string
	BillingContactEmail string
	InvoiceCCEmail      string
	TotalAmount         float64
	InvoiceDate         string
	PoNumber            string
	LineItems           []LineItem
}

type LineItem struct {
	Quantity        int64
	Description     string
	UnitCost        float64
	ItemTotal       float64
	ServiceLocation string
	ItemDetails     []interface{}
}
