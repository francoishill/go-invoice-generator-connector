package invoicegenerator

import (
	"testing"
)

func TestCreate(t *testing.T) {

	r := new(Invoice)

	r.From = "parag@invoiced.com"
	r.To = "jared@invoiced.com"
	r.Date = "January 2nd 2014"
	r.Logo = "https://invoiced.com/img/header-logo.png"
	r.PaymentTerms = "NET 30"
	r.Number = "INV-001"
	r.PurchaseOrder = "PO-32422"

	items := make([]Item, 1)
	item := Item{}
	item.Name = "Marketing"
	item.Quantity = 221
	item.Unit_cost = 50.45

	items[0] = item

	r.Items = items

	r.Discounts = 12.21
	r.Tax = 13.11
	r.Shipping = 12.20

	r.TaxTitle = "Austin Tax"

	r.Create("bill-5", "/usr/local/pdfstash/")
}
