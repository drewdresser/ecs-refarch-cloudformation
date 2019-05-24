package main

import "github.com/gin-gonic/gin"
import "github.com/satori/go.uuid"

type account struct {
	AccountId        string
	AccountName      string
	Description      string
	AccountType      string
}
type postResponse struct {
	AccountId				string
}

type healthCheck struct {
	status string
}
// {
//   "accountId": {},
//   "parentAccountId": {},
//   "accountName": {},
//   "accountType": "SYSTEM",
//   "mfaRequirement": "OPTIONAL"
// }

func main() {

	router := gin.Default()

	// Respond to GET requests
	// router.GET("/account", func(c *gin.Context) {
	//
	// 	accounts := []account{
	// 		account{
	// 			AccountId:       "ea9ef5f9-107b-4a4e-9295-57d701d85a92",
	// 			AccountName:     "TestSystemAccount",
	// 			Description:     "Account is used for a system user.",
	// 			AccountType:     "SYSTEM",
	// 		},
	// 		account{
	// 			AccountId:       "e980f21c-b6e3-4e3a-bf8e-9f2c3074637a",
	// 			AccountName:     "CustomerAAccount",
	// 			Description:     "My customer account",
	// 			AccountType:     "CUSTOMER",
	// 		},
	// 		account{
	// 			AccountId:       "04f7b1c0-a7ea-472b-88d6-a977eb94f4da",
	// 			AccountName:     "NewUserAccount",
	// 			Description:     "Account used to sign up for rides!",
	// 			AccountType:     "CUSTOMER",
	// 		},
	// 	}
	//
	// 	c.IndentedJSON(200, accounts)
	//
	// })

	router.GET("/healthcheck", func(c *gin.Context) {
		health := healthCheck{
			status: "healthly",
		}
		c.IndentedJSON(200, health)
	})

	router.POST("/account", func(c *gin.Context) {
		id, _ := uuid.NewV4()
		response := postResponse{
			AccountId: id.String(),
		}
		c.IndentedJSON(201, response)
	})

	router.GET("/account/:accountId", func(c *gin.Context) {
		account := account{
			AccountId: c.Param("accountId"),
			AccountName: "TestAccount",
			Description: "This is an example of an account",
			AccountType: "CUSTOMER",
		}
		c.IndentedJSON(200, account)
	})

	// Serve all of the things..
	router.Run(":8001")

}
