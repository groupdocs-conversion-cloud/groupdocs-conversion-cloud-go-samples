package convert

import (
	"fmt"
	"os"

	"github.com/antihax/optional"
	conversion "github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertToPdfDirectOptions() {

	settings := conversion.ConvertApiConvertDocumentDirectOpts{
		LoadOptions: optional.NewInterface(&models.WordProcessingLoadOptions{
			Format:   "docx",
			Password: "password",
		}),
	}

	path := "./Resources/WordProcessing/four-pages.docx"
	localFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", path, err)
		return
	}
	defer localFile.Close()

	result, _, err := config.Client.ConvertApi.ConvertDocumentDirect(config.Ctx, "pdf", localFile, &settings)

	if err != nil {
		fmt.Printf("ConvertToPdfDirectOptions error: %v\n", err)
		return
	}

	fi, _ := result.Stat()

	fmt.Printf("Document converted successfully: %v\n", fi.Size())
}
