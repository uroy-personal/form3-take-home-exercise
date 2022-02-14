package form3

import (
	"context"
	"fmt"
	"os"
	"testing"

	"log"

	"github.com/stretchr/testify/assert"
)

var uniqueAccountId string
var uniqueOrganizationId string

func init() {
	// accountBaseURL := "http://accountapi:8080"
	fmt.Printf("AccountBaseURL: %+v\n", os.Getenv("ACCOUNT_BASE_URL"))
	c := NewClient(os.Getenv("ACCOUNT_BASE_URL"))
	var err error
	uniqueAccountId, err = c.GenerateUniqueID()
	uniqueOrganizationId, err = c.GenerateUniqueID()
	fmt.Printf("baseURL: %+v\n", c.BaseURL)
	fmt.Printf("uniqueAccountId: %+v\n", uniqueAccountId)
	fmt.Printf("uniqueOrganizationId: %+v\n", uniqueOrganizationId)
	fmt.Printf("error: %+v\n", err)
}
func TestAll(t *testing.T) {
	t.Run("A=Create", func(t *testing.T) {
		//test data
		country := "GB"
		baseCurrency := "GBP"
		bankId := "400300"
		bankIdCode := "GBDSC"
		bic := "NWBKGB22"
		name := []string{"UmeshRoy"}
		fmt.Printf("AccountBaseURL2: %+v\n", os.Getenv("ACCOUNT_BASE_URL"))
		c := NewClient(os.Getenv("ACCOUNT_BASE_URL"))
		//fmt.Printf("uniqueAccountId: %+v\n", uniqueAccountId)
		//fmt.Printf("uniqueOrganizationId: %+v\n", uniqueOrganizationId)
		accountMessage, err := c.PopulateAccountMesage(bankId, bankIdCode, baseCurrency, bic, country, name, uniqueOrganizationId, uniqueAccountId)
		//fmt.Printf("accountMessage: %+v\n", accountMessage)
		ctx := context.Background()

		// res := AccountMessage{}
		res, err := c.CreateAccount(ctx, accountMessage)

		assert.Nil(t, err, "expecting nil error")

		assert.NotNil(t, res, "expecting non-nil result")

	})

	t.Run("A=Fetch", func(t *testing.T) {

		c := NewClient(os.Getenv("ACCOUNT_BASE_URL"))

		ctx := context.Background()

		res, err := c.FetchAccount(ctx, uniqueAccountId)

		assert.Nil(t, err, "expecting nil error")

		assert.NotNil(t, res, "expecting non-nil result")

	})

	t.Run("A=Delete", func(t *testing.T) {

		c := NewClient(os.Getenv("ACCOUNT_BASE_URL"))

		ctx := context.Background()

		res, err := c.DeleteAccount(ctx, uniqueAccountId)

		fmt.Printf("Delete-Response %+v\n", res)
		fmt.Printf("Delete-error %+v\n", err)

		assert.Nil(t, res, "expecting nil error")

		//assert.NotNil(t, res, "expecting non-nil result")

		//assert.Equal(t, 1, res.Count, "expecting 1 face found")
		//assert.Equal(t, 1, res.PagesCount, "expecting 1 PAGE found")
		//assert.Equal(t, "integration_face_id", res.Faces[0].FaceID, "expecting correct face_id")
		//assert.NotEmpty(t, res.Faces[0].FaceToken, "expecting non-empty face_token")
		//assert.Greater(t, len(res.Faces[0].FaceImages), 0, "expecting non-empty face_images")
	})
	// <tear-down code>
}
