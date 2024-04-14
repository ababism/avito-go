package postgre

import (
	"fmt"
	"strings"
)

const (
	BannerTable        = "bank_accounts"
	TagTable           = "payments"
	FeatureTable       = "feature"
	BannerFeatureTable = "banners_feature"
	BannerTagsTable    = "banners_tags"
)

func formatQuery(q string) string {
	return fmt.Sprintf("SQL Query: %s", strings.ReplaceAll(strings.ReplaceAll(q, "\t", ""), "\n", " "))
}
