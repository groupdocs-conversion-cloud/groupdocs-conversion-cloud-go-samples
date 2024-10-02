package async

import (
	"fmt"
	"os"
	"time"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertToPdfDirectAsync() {

	path := "./Resources/WordProcessing/four-pages.docx"
	localFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", path, err)
		return
	}
	defer localFile.Close()

	operationId, _, err := config.Client.AsyncApi.StartConvert(config.Ctx, "pdf", localFile, nil)

	if err != nil {
		fmt.Printf("ConvertToPdfDirectAsync error: %v\n", err)
		return
	}

	fmt.Printf("Operation ID: %v\n", operationId)

	for {
		time.Sleep(1 * time.Second)
		result, _, err := config.Client.AsyncApi.GetOperationStatus(config.Ctx, operationId)
		if err != nil {
			fmt.Printf("ConvertToPdfDirectAsync error: %v\n", err)
			return
		}
		if result.Status == models.StatusEnumFinished {
			result, _, _ := config.Client.AsyncApi.GetOperationResult(config.Ctx, operationId)
			fi, _ := result.Stat()
			fmt.Printf("Document converted successfully: %v\n", fi.Size())
			break
		}
		if result.Status == models.StatusEnumFailed {
			fmt.Printf("Document converted failed: %v\n", result.Error_)
			break
		}
		fmt.Printf("Operation status: %v\n", result.Status)
	}
}
