package main

import "fmt"

type Marker struct {
	Name  string
	Color string
	Year  int
	Price int
}

func NewMarker(name, color string, year, price int) *Marker {
	return &Marker{
		Name:  name,
		Color: color,
		Year:  year,
		Price: price,
	}
}

type Invoice struct {
	Marker     *Marker
	Quantities int
}

func NewInvoice(marker *Marker, quantities int) *Invoice {
	return &Invoice{
		Marker:     marker,
		Quantities: quantities,
	}
}

func (invoice *Invoice) CalculateTotal() int {
	price := invoice.Marker.Price * invoice.Quantities
	return price
}

func (invoice *Invoice) PrintInvoice() {
	fmt.Printf("Invoice is printed price is :: %d\n", invoice.Marker.Price)
}

func (invoice *Invoice) SaveInDB() {
	fmt.Printf("Saved in DB is :: %d price:: %d\n", invoice.Quantities, invoice.Marker.Price)
}

func main() {
	// Create a Marker object
	marker := NewMarker("ExampleMarker", "Red", 2023, 5)

	// Create an Invoice object with the Marker and quantities
	invoice := NewInvoice(marker, 10)

	// Calculate the total price
	totalPrice := invoice.CalculateTotal()

	// Print the invoice
	invoice.PrintInvoice()

	// Save the invoice in the database
	invoice.SaveInDB()

	// Print the details of the Marker
	fmt.Println("Marker Details:")
	fmt.Println("Name:", marker.Name)
	fmt.Println("Color:", marker.Color)
	fmt.Println("Year:", marker.Year)
	fmt.Println("Price:", marker.Price)

	// Print the total price from the Invoice
	fmt.Println("Total Price:", totalPrice)
}
