package main

import (
	"bytes"
	"html"

	"golang.org/x/perf/benchstat"
)

func colorizeHTML(buf *bytes.Buffer, change int, isBiggerIsBetter bool) {
	if change > 0 {
		if isBiggerIsBetter {
			buf.WriteString(" bgcolor='red'")
		} else {
			buf.WriteString(" bgcolor='green'")
		}
	} else if change < 0 {
		if isBiggerIsBetter {
			buf.WriteString(" bgcolor='green'")
		} else {
			buf.WriteString(" bgcolor='red'")
		}
	}
}

// FormatHTML appends an HTML formatting of the tables to buf.
func FormatHTML(buf *bytes.Buffer, tables []*benchstat.Table, colorized bool, increasing StringSet) {
	if len(tables) == 0 {
		return
	}

	if tables[0].OldNewDelta {
		buf.WriteString("<table border='1' class='benchstat'>\n")
	} else {
		buf.WriteString("<table class='benchstat oldnew'>\n")
	}

	buf.WriteString("<tr class='configs'><th/>")
	for _, c := range tables[0].Configs {
		buf.WriteString("<th>")
		buf.WriteString(html.EscapeString(c))
		buf.WriteString("</th>")
	}
	buf.WriteString("</tr>\n")

	for _, table := range tables {
		buf.WriteString("<tbody>\n")

		// metrics
		buf.WriteString("<tr><th/>")
		buf.WriteString("<th colspan='2' class='metric'>")
		buf.WriteString(html.EscapeString(table.Metric))
		buf.WriteString("</th><th>delta</th>\n</tr>\n")

		_, isBiggerIsBetter := increasing[table.Metric]

		var group string
		// rows
		for _, row := range table.Rows {
			if row.Group != group {
				// tests group
				group = row.Group
				buf.WriteString("<tr><th class='group' colspan='4'>")
				buf.WriteString(html.EscapeString(group))
				buf.WriteString("</th></tr>\n")
			}
			buf.WriteString("<tr><td")
			// color class
			colorizeHTML(buf, row.Change, isBiggerIsBetter)
			buf.WriteString(">")
			buf.WriteString(html.EscapeString(row.Benchmark))
			for _, metric := range row.Metrics {
				buf.WriteString("</td><td>")
				buf.WriteString(html.EscapeString(metric.FormatMean(row.Scaler)))
				buf.WriteString(" ± ")
				buf.WriteString(html.EscapeString(metric.FormatDiff()))
			}
			// delta
			buf.WriteString("</td><td")
			colorizeHTML(buf, row.Change, isBiggerIsBetter)
			buf.WriteString(">")
			buf.WriteString(html.EscapeString(row.Delta))
			buf.WriteString(" ")
			buf.WriteString(html.EscapeString(row.Note))
			buf.WriteString("</td></tr>\n")
		}
		buf.WriteString("</tbody>\n")
	}
	buf.WriteString("</table>\n")
}
