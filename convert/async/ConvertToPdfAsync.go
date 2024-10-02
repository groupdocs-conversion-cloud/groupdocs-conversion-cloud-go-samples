package async

import (
	"fmt"
	"time"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertToPdfAsync() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "WordProcessing/password-protected.docx",
		OutputPath: "converted",
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
		},
	}

	operationId, _, err := config.Client.AsyncApi.StartConvertAndSave(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertToPdfAsync error: %v\n", err)
		return
	}

	fmt.Printf("Operation ID: %v\n", operationId)

	for {
		time.Sleep(1 * time.Second)
		result, _, err := config.Client.AsyncApi.GetOperationStatus(config.Ctx, operationId)
		if err != nil {
			fmt.Printf("ConvertToPdfAsync error: %v\n", err)
			return
		}
		if result.Status == models.StatusEnumFinished {
			fmt.Printf("Document converted successfully: %v\n", result.Result[0].Url)
			break
		}
		if result.Status == models.StatusEnumFailed {
			fmt.Printf("Document converted failed: %v\n", result.Error_)
			break
		}
		fmt.Printf("Operation status: %v\n", result.Status)
	}
}
