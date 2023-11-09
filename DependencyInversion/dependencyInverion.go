package main

import "fmt"

// KeyboardI interface
type KeyboardI interface {
	CanType()
	WriteFasterWiredOrBluetooth()
	EasyToUseWiredOrBluetooth()
}

// MouseI interface
type MouseI interface {
	Cursor()
}

// BluetoothKeyboard struct
type BluetoothKeyboard struct{}

func (bk BluetoothKeyboard) CanType() {
	fmt.Println("We can type in BluetoothKeyboard")
}

func (bk BluetoothKeyboard) WriteFasterWiredOrBluetooth() {
	fmt.Println("BluetoothKeyboard is lack in writing")
}

func (bk BluetoothKeyboard) EasyToUseWiredOrBluetooth() {
	fmt.Println("BluetoothKeyboard is easy to use")
}

// BluetoothMouse struct
type BluetoothMouse struct{}

func (bm BluetoothMouse) Cursor() {
	fmt.Println("BluetoothMouse cursor is slow")
}

// WiredKeyboard struct
type WiredKeyboard struct{}

func (wk WiredKeyboard) CanType() {
	fmt.Println("We can type in wiredKeyboard")
}

func (wk WiredKeyboard) WriteFasterWiredOrBluetooth() {
	fmt.Println("WiredKeyboard is faster to write")
}

func (wk WiredKeyboard) EasyToUseWiredOrBluetooth() {
	fmt.Println("WiredKeyboard is not easy to use")
}

// WiredMouse struct
type WiredMouse struct{}

func (wm WiredMouse) Cursor() {
	fmt.Println("WiredMouse cursor is fast")
}

// MackBook struct
type MackBook struct {
	KeyboardI KeyboardI
	MouseI    MouseI
}

// UseMackBook function
func UseMackBook(mackBook MackBook) {
	fmt.Println("Using MackBook with:")
	mackBook.KeyboardI.CanType()
	mackBook.KeyboardI.WriteFasterWiredOrBluetooth()
	mackBook.KeyboardI.EasyToUseWiredOrBluetooth()
	mackBook.MouseI.Cursor()
	fmt.Println("--------------")
}

func main() {
	bluetoothKeyboard := BluetoothKeyboard{}
	bluetoothMouse := BluetoothMouse{}

	wiredKeyboard := WiredKeyboard{}
	wiredMouse := WiredMouse{}

	mackBookWithBluetooth := MackBook{KeyboardI: bluetoothKeyboard, MouseI: bluetoothMouse}
	mackBookWithWired := MackBook{KeyboardI: wiredKeyboard, MouseI: wiredMouse}

	UseMackBook(mackBookWithBluetooth)
	UseMackBook(mackBookWithWired)
}
