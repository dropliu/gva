package starkbank

import (
	"fmt"
	"time"

	"github.com/starkbank/sdk-go/starkbank/dynamicbrcode"
	"github.com/starkinfra/core-go/starkcore/user/organization"

	"github.com/starkinfra/core-go/starkcore/user/project"
	"github.com/starkinfra/core-go/starkcore/user/user"
)

// This is only an example of a private key content. You should use your own key.
var privateKeyContent = `   
-----BEGIN EC PRIVATE KEY-----
MHQCAQEEIMV/xiBoc8UJVqHpgTAvD1isln8IIU3iYwQJ/ws5+KXIoAcGBSuBBAAKoUQDQgAEekAFMl8sX8MHcwUybjpzpklFuBxPtNk2ReiEaLViLWdAjzqdpCfRKrFAbsQPcoefekG5BXCBYj5FhkICXvXI0g==
-----END EC PRIVATE KEY-----`

// for project:
var projectInfo = project.Project{
	Id:          "5012474017349632",
	PrivateKey:  privateKeyContent,
	Environment: "production",
}

// or, for organization users:
var usr = organization.Organization{
	Id:          "4886161570922496",
	PrivateKey:  privateKeyContent,
	Environment: "production",
	WorkspaceId: "5585470964629504",
}

func main() {

	// privateKey, publicKey := key.Create("")

	// log.Println(privateKey)
	// log.Println(publicKey)

	// starkbank.User = user

	// workspaces := workspace.Query(nil, nil)

	// for workspace := range workspaces {
	//  fmt.Printf("%+v", workspace)
	// }

	due := time.Now().Add(time.Hour * 24)

	fmt.Println(due.Unix())
	brcodes, err := dynamicbrcode.Create(
		[]dynamicbrcode.DynamicBrcode{
			{
				Amount:     4000,
				Expiration: 123456789,
				Tags:       []string{"New sword", "DynamicBrcode #1234"},
			},
		}, usr)
	if err.Errors != nil {
		for _, e := range err.Errors {
			fmt.Printf("code: %s, message: %s\n", e.Code, e.Message)
		}
		return
	}

	for _, brcode := range brcodes {
		fmt.Printf("%+v", brcode)
	}
}

// CreateDynamicBrCode 创建一个二维码，返回二维码链接
func CreateDynamicBrcode(u user.User, amount, expiration int, tags ...string) (dynamicbrcode.DynamicBrcode, error) {
	brcodes := []dynamicbrcode.DynamicBrcode{
		{
			Amount:     amount,
			Expiration: expiration,
			Tags:       tags,
		},
	}
	list, errs := dynamicbrcode.Create(brcodes, u)
	if len(errs.Errors) > 0 {
		err := errs.Errors[0]
		return dynamicbrcode.DynamicBrcode{}, fmt.Errorf("code: %s, msg: %s", err.Code, err.Message)
	}

	return list[0], nil
}
