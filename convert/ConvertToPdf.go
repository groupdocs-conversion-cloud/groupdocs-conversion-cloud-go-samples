package convert

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertToPdf() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "WordProcessing/four-pages.docx",
		OutputPath: "converted",
		ConvertOptions: &models.PdfConvertOptions{
			CenterWindow:         true,
			CompressImages:       false,
			DisplayDocTitle:      true,
			Dpi:                  1024,
			FitWindow:            false,
			FromPage:             1,
			Grayscale:            false,
			ImageQuality:         100,
			Linearize:            false,
			MarginTop:            5,
			MarginLeft:           5,
			Password:             "password",
			UnembedFonts:         true,
			RemoveUnusedStreams:  true,
			RemoveUnusedObjects:  true,
			RemovePdfaCompliance: false,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertToPdf error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
