package services

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type PrintService struct {
}

func NewPrintService() *PrintService {
	return &PrintService{}
}

func (ps *PrintService) RenderTable(data [][]string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.Header(data[0])
	table.Bulk(data[1:])
	table.Render()
}

func (ps *PrintService) Println(value string) {
	fmt.Println(value)
}
