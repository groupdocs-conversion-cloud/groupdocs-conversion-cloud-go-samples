package spreadsheet

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertSpreadsheetWithHiddenSheetsIncluded() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Spreadsheet/with_hidden_sheet.xlsx",
		OutputPath: "converted",
		LoadOptions: &models.SpreadsheetLoadOptions{
			ShowHiddenSheets: true,
			OnePagePerSheet:  true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertSpreadsheetWithHiddenSheetsIncluded error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
