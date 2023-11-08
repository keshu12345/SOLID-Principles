package main

import "fmt"

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

func (i *Invoice) CalculateTotal() int {
	return i.Marker.Price * i.Quantities
}

type InvoiceDAO struct {
	Invoice Invoice
}

func (dao *InvoiceDAO) SaveInDB() {
	fmt.Printf("Saved in DB is :: quantities: %d, price: %d\n", dao.Invoice.Quantities, dao.Invoice.Marker.Price)
}

func (dao *InvoiceDAO) SaveInFile() {
	fmt.Printf("Saved in File is :: quantities: %d, price: %d\n", dao.Invoice.Quantities, dao.Invoice.Marker.Price)
}

func main() {
	// Create a Marker object
	marker := Marker{Name: "ExampleMarker", Color: "Red", Year: 2023, Price: 5}
	// Create an Invoice object with the Marker and quantities
	invoice := Invoice{Marker: marker, Quantities: 10}
	// Calculate the total price
	totalPrice := invoice.CalculateTotal()
	invoiceDAO := InvoiceDAO{Invoice: invoice}
	invoiceDAO.SaveInFile()
	invoiceDAO.SaveInDB()
	fmt.Println("Total Price:", totalPrice)
}
