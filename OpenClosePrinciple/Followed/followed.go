package main

import "fmt"

type InvoiceDAO interface {
	Save(invoice Invoice)
}

type FileInvoiceDAO struct{}

func (dao FileInvoiceDAO) Save(invoice Invoice) {
	fmt.Printf("Saved in File is :: quantities: %d price: %d\n", invoice.Quantities, invoice.Marker.Price)
}

type DatabaseInvoiceDAO struct{}

func (dao DatabaseInvoiceDAO) Save(invoice Invoice) {
	fmt.Printf("Saved in DB is :: quantities: %d price: %d\n", invoice.Quantities, invoice.Marker.Price)
}

type Marker struct {
	Name  string
	Color string
	Year  int
	Price int
}

type Invoice struct {
	Marker     Marker
	Quantities int
}

func (invoice Invoice) CalculateTotals() int {
	price := invoice.Marker.Price * invoice.Quantities
	return price
}

func main() {
	// Create a marker and an invoice
	marker := Marker{Name: "Marker A", Color: "Red", Year: 2023, Price: 5}
	invoice := Invoice{Marker: marker, Quantities: 10}

	// Create instances of DAO implementations
	var fileInvoiceDAO InvoiceDAO
	var databaseInvoiceDAO InvoiceDAO

	fileInvoiceDAO = FileInvoiceDAO{}
	databaseInvoiceDAO = DatabaseInvoiceDAO{}

	// Save the invoice using FileInvoiceDAO
	fileInvoiceDAO.Save(invoice)

	// Save the invoice using DatabaseInvoiceDAO
	databaseInvoiceDAO.Save(invoice)

	// Calculate and print the total price
	totalPrice := invoice.CalculateTotals()
	fmt.Printf("Total Price: %d\n", totalPrice)
}
