package jobdb

import "fmt"

func ValidJobTitle(query string) string {
	return fmt.Sprintf(`
	%s where jobTitle is not null
	and jobTitle <> ""
	and companyName is not null
	and companyName <> ""
	`, query)
}

func ConditionalSource(query, source string) string {
	return fmt.Sprintf(`
	%s and source = "%s"`, query, source)
}

func ConditionalCrawledAt(query, date string) string {
	return fmt.Sprintf(`
	%s and crawledAt = "%s"`, query, date)
}
