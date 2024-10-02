package convert

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertToPdfResponseBody() {

	settings := models.ConvertSettings{
		Format:   "pdf",
		FilePath: "WordProcessing/four-pages.docx",
		LoadOptions: &models.WordProcessingLoadOptions{
			Password: "password",
		},
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
			PageSize:             models.PageSizeEnumA4,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocumentDownload(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertToPdfResponseBody error: %v\n", err)
		return
	}

	fi, _ := result.Stat()

	fmt.Printf("Document converted successfully: %v\n", fi.Size())
}
