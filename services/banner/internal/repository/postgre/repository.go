package postgre

import (
	"fmt"
	"strings"
)

const (
	BankAccountsTable     = "bank_accounts"
	PaymentsTable         = "payments"
	ProductsTable         = "products"
	DiscountsTable        = "discounts"
	PayoutRequestsTable   = "payout_requests"
	PurchaseRequestsTable = "purchase_requests"
)

func formatQuery(q string) string {
	return fmt.Sprintf("SQL Query: %s", strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " "))
}
