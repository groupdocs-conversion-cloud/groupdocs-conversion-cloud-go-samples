package txt

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertTxtBySpecifyingEncoding() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Text/sample.txt",
		OutputPath: "converted",
		LoadOptions: &models.TxtLoadOptions{
			Encoding: "shift_jis",
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertTxtBySpecifyingEncoding error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
