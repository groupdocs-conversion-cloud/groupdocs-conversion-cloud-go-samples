package pdf

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertPdfAndRemoveEmbeddedFiles() {

	settings := models.ConvertSettings{
		Format:     "docx",
		FilePath:   "Pdf/sample.pdf",
		OutputPath: "converted",
		LoadOptions: &models.PdfLoadOptions{
			Format:              "pdf",
			RemoveEmbeddedFiles: true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertPdfAndRemoveEmbeddedFiles error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
