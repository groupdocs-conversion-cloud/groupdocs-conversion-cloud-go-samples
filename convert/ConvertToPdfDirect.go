package convert

import (
	"fmt"
	"os"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
)

func ConvertToPdfDirect() {

	path := "./Resources/WordProcessing/four-pages.docx"
	localFile, err := os.Open(path)
	if err != nil {
		fmt.Printf("Failed to open file %s: %v\n", path, err)
		return
	}
	defer localFile.Close()

	result, _, err := config.Client.ConvertApi.ConvertDocumentDirect(config.Ctx, "pdf", localFile, nil)

	if err != nil {
		fmt.Printf("ConvertToPdfDirect error: %v\n", err)
		return
	}

	fi, _ := result.Stat()

	fmt.Printf("Document converted successfully: %v\n", fi.Size())
}
