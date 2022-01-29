package service

import (
	"context"
	"os"

	"github.com/olekukonko/tablewriter"
)

func Component(ctx context.Context) error {
	// Needs to be in sync with service spec.

	data := [][]string{
		{"postgresql", "postgresql://localhost:5432", "Postgresql", "Service component"},
		{"kafdrop", "http://localhost:9100", "Kafdrop", "Platform component"},
		{"jaeger-ui", "http://localhost:16686", "Jaeger UI", "Platform component"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Name", "Address", "Description", "Type"})

	for _, v := range data {
		table.Append(v)
	}

	table.Render()

	return nil
}
