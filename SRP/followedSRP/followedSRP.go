package main

import "fmt"

// Marker struct
type Marker struct {
	Name  string
	Color string
	Year  int
	Price int
}

// Invoice struct
type Invoice struct {
	Marker     Marker
	Quantities int
}

// CalculateTotals calculates the total price
func (i Invoice) CalculateTotals() int {
	price := i.Marker.Price * i.Quantities
	return price
}

// PrintInvoice struct
type PrintInvoice struct {
	Invoice Invoice
}

// PrintInvoice prints the invoice details
func (p PrintInvoice) PrintInvoice() {
	fmt.Printf("Invoice is printed price is :: %d\n", p.Invoice.Marker.Price)
}

// SaveInDB struct
type SaveInDB struct {
	Invoice Invoice
}

// SaveInDB saves the invoice in the database
func (s SaveInDB) SaveInDB() {
	fmt.Printf("Saved in DB is :: %d price:: %d\n", s.Invoice.Quantities, s.Invoice.Marker.Price)
}

func main() {
	// Create a Marker object
	marker := Marker{
		Name:  "ExampleMarker",
		Color: "Red",
		Year:  2023,
		Price: 5,
	}

	// Create an Invoice object with the Marker and quantities
	invoice := Invoice{
		Marker:     marker,
		Quantities: 10,
	}

	// Create a PrintInvoice object and print the invoice
	printInvoice := PrintInvoice{Invoice: invoice}
	printInvoice.PrintInvoice()

	// Create a SaveInDB object and save the invoice in the database
	saveInDB := SaveInDB{Invoice: invoice}
	saveInDB.SaveInDB()

	// Calculate the total price and print it
	totalPrice := invoice.CalculateTotals()
	fmt.Printf("Total Price: %d\n", totalPrice)
}
