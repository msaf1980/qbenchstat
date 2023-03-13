package main

import (
	"bytes"

	"golang.org/x/perf/benchstat"
)

func writeMDTableDelimiter(buf *bytes.Buffer, n int) {
	buf.WriteString("|")
	for i := 0; i < n; i++ {
		buf.WriteString("--|")
	}
	buf.WriteString("\n")
}

// FormatMarkdown formats benchstat output as markdown
func FormatMarkdown(buf *bytes.Buffer, tables []*benchstat.Table, colorized bool, increasing StringSet) {
	if len(tables) == 0 {
		return
	}

	buf.WriteString("| |")
	for _, c := range tables[0].Configs {
		buf.WriteString(" ")
		buf.WriteString(c)
		buf.WriteString(" |")
	}
	buf.WriteString(" |\n")
	writeMDTableDelimiter(buf, len(tables[0].Configs)+2)

	for _, table := range tables {
		// metrics
		buf.WriteString("| | ")
		buf.WriteString(table.Metric)
		buf.WriteString(" | | delta |\n")

		// _, isBiggerIsBetter := increasing[table.Metric]

		var group string
		// rows
		for _, row := range table.Rows {
			if row.Group != group {
				// tests group
				group = row.Group
				buf.WriteString("| ")
				buf.WriteString(group)
				buf.WriteString(" |\n")
			}
			// color class
			// 		colorizeHTML(buf, row.Change, isBiggerIsBetter)
			buf.WriteString(row.Benchmark)
			for _, metric := range row.Metrics {
				buf.WriteString("| ")
				buf.WriteString(metric.FormatMean(row.Scaler))
				buf.WriteString(" ± ")
				buf.WriteString(metric.FormatDiff())
			}
			buf.WriteString(" | ")
			// delta
			// colorizeHTML(buf, row.Change, isBiggerIsBetter)
			buf.WriteString(row.Delta)
			buf.WriteString(" ")
			buf.WriteString(row.Note)
			buf.WriteString(" |\n")
		}
	}
}
