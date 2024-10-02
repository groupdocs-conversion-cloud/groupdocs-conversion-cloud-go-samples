package txt

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertTxtByControllingTrailingSpacesBehavior() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Text/sample.txt",
		OutputPath: "converted",
		LoadOptions: &models.TxtLoadOptions{
			TrailingSpacesOptions:          models.TrailingSpacesOptionsEnumTrim,
			DetectNumberingWithWhitespaces: true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertTxtByControllingTrailingSpacesBehavior error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
