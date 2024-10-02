package html

import (
	"fmt"

	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go-samples/config"
	"github.com/groupdocs-conversion-cloud/groupdocs-conversion-cloud-go/models"
)

func ConvertHtmlWithPageNumbering() {

	settings := models.ConvertSettings{
		Format:     "pdf",
		FilePath:   "Html/sample.html",
		OutputPath: "converted",
		LoadOptions: &models.WebLoadOptions{
			Format:        "html",
			PageNumbering: true,
		},
	}

	result, _, err := config.Client.ConvertApi.ConvertDocument(config.Ctx, settings)

	if err != nil {
		fmt.Printf("ConvertHtmlWithPageNumbering error: %v\n", err)
		return
	}

	fmt.Printf("Document converted successfully: %v\n", result[0].Url)
}
