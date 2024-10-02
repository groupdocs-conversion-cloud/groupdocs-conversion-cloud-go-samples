package common

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func AddWatermark() {

	watermark := models.WatermarkOptions{
		Text:       "Sample watermark",
		Color:      "Red",
		Width:      100,
		Height:     100,
		Background: true,
	}

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "WordProcessing/four-pages.docx",
		OutputPath: "converted",
		ConvertOptions: &models.PdfConvertOptions{
			WatermarkOptions: &watermark,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("AddWatermark error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
