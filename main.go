package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/scryinfo/iscap_demo/src/sdk"
	"github.com/scryinfo/iscap_demo/src/sdk/core/chainoperations"
	"github.com/scryinfo/iscap_demo/src/sdk/core/ethereum/events"
	"github.com/scryinfo/iscap_demo/src/sdk/scryclient"
	cif "github.com/scryinfo/iscap_demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/iscap_demo/src/sdk/util/accounts"
	"github.com/spf13/viper"
	"log"
	"math/big"
	"net/http"
	"os"
	"time"
)

var (
	publishId                        = ""
	sellerAddr                       = ""
	txId                    *big.Int = big.NewInt(0)
	metaDataIdEncWithSeller []byte
	metaDataIdEncWithBuyer  []byte
	clientPassword                                 = "888888"
	seller                  *scryclient.ScryClient = nil
	sleepTime               time.Duration          = 15000000000
	appIndex                                       = 0
	appTitle                                       = "scryapp"
	buyer                   *scryclient.ScryClient = nil
)

var (
	ethNodeAddr          = ""
	keyServiceAddr       = ""
	protocolContractAddr = ""
	tokenContractAddr    = ""
	ipfsNodeAddr         = ""
)

// Get the account eth By Address String
func GetEth(s string) *big.Int {
	eth, err := scryclient.ScryClient{}.GetEth(common.HexToAddress(s))
	if err != nil {
		fmt.Println("Get Eth fail", err)
	}
	return eth
}

// Get the account eth By ScryClient
func GetEthByClient(client *scryclient.ScryClient) *big.Int {
	eth, err := scryclient.ScryClient{}.GetEth(common.HexToAddress(client.Account.Address))
	if err != nil {
		fmt.Println("Get Eth fail", err)
	}
	return eth
}

// Get the account eth By ScryClient
func GetTokenByClient(client *scryclient.ScryClient) *big.Int {
	token, err := client.GetScryToken(common.HexToAddress(client.Account.Address))
	if err != nil {
		fmt.Println("Get Eth fail", err)
	}
	return token
}

// Printf the eth and Token
func PrintfEthAndTokenByClient(client *scryclient.ScryClient) {
	fmt.Println(fmt.Sprintf("account %s has eth %d , token %d ", client.Account.Address, GetEthByClient(client), GetTokenByClient(client)))
}
func initScry() {
	// init ScryInfosSDK
	wd, _ := os.Getwd()
	err := sdk.Init(ethNodeAddr, keyServiceAddr, protocolContractAddr, tokenContractAddr, 0, ipfsNodeAddr, wd+"/testconsole.log", "scryapp1")
	if err != nil {
		fmt.Println("sdk init err!", err)
	}
	// 初始化seller
	seller = scryclient.NewScryClient("0x2008cc463061d385d87a294b2f3edce229f74b58")
	// 查询账户比特币和Token余额
	PrintfEthAndTokenByClient(seller)
	// 1. 数据成功发布到IPFS和区块链上
	seller.SubscribeEvent("DataPublish", onPublish)
	// 6. 卖⽅上传待卖数据的metaDataId到区块链上，该ID使⽤买⽅公钥加密
	// 8. 卖⽅收到购买数据的通知，⽣成使⽤买⽅公钥加密的meta data id，发给合约
	seller.SubscribeEvent("Buy", onPurchase)
	seller.SubscribeEvent("TransactionCreate", onTransactionCreate)
	seller.SubscribeEvent("TransactionClose", onClose)

	buyer = scryclient.NewScryClient("0x76c893c10e78fe205cc84489aa65ce29e91ad597")
	// 查询账户比特币和Token余额
	PrintfEthAndTokenByClient(buyer)
	fmt.Println(fmt.Sprintf("account has eth %d , token %d", GetEthByClient(seller), GetTokenByClient(seller)))
	// 2.订阅消息发布
	// 3.准许地址spender从本接⼝调⽤⽅地址进⾏token转账
	buyer.SubscribeEvent("DataPublish", onPublishBuyer)
	// 4.买⽅准备购买数据，希望得到待买数据的证明数据ID
	buyer.SubscribeEvent("Approval", onApprovalBuyerTransfer)
	// 5.买⽅正式购买数据。
	buyer.SubscribeEvent("TransactionCreate", onTransactionCreateBuyer)
	// 7.买⽅确认数据真实性
	// 9.买⽅拿到meta data id，终于可以从IPFS上下载完整数据
	buyer.SubscribeEvent("ReadyForDownload", onReadyForDownload)
	buyer.SubscribeEvent("TransactionClose", onClose)
	// 发送数据
	fmt.Println("Start testing tx without verification...")

}

func init() {
	// read configs
	viper.SetConfigName("conf")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	ethNodeAddr = viper.GetString("ethNodeAddr")
	keyServiceAddr = viper.GetString("keyServiceAddr")
	protocolContractAddr = viper.GetString("protocolContractAddr")
	tokenContractAddr = viper.GetString("tokenContractAddr")
	ipfsNodeAddr = viper.GetString("ipfsNodeAddr")
	//initScry()
}

type UploadBody struct {
	Address          string `json:"address"`           // user address
	Name             string `json:"name"`              // user email
	Gender           string `json:"gender"`            // user gender
	Country          string `json:"country"`           // user country
	Age              string `json:"age"`               // user age
	ResidencyAddress string `json:"residency_address"` // user residency address
}

type AuthorityCertify struct {
	AuthorityAddress string `json:"authority_address"`
	UserAddress      string `json:"user_address"`
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	// Institute
	r.GET("/institute", func(c *gin.Context) {
		address := c.Query("address")
		log.Println(address)
		c.JSON(http.StatusOK, gin.H{
			"user":   "0x2008cc463061d385d87a294b2f3edce229f74b58",
			"status": "pass",
		})
		return
	})
	user := r.Group("/user")
	{
		// upload
		user.POST("/upload", func(context *gin.Context) {
			var uploadBody UploadBody
			err := context.BindJSON(&uploadBody)
			if err != nil {
				context.JSON(http.StatusBadRequest, gin.H{
					"err_code": 0001,
					"message":  err,
				})
				return
			}
			//SellerPublishData()
			// upload user information
			context.JSON(http.StatusOK, gin.H{
				"status": true,
			})
			return
		})

		user.GET("/grant", func(c *gin.Context) {
			userAddress := c.Query("address")
			instituteAddress := c.Query("institute_address")
			log.Println(userAddress)
			log.Println(instituteAddress)
			c.JSON(http.StatusOK, gin.H{
				"status": true,
			})
			return
		})
		user.GET("/status", func(c *gin.Context) {
			userAddress := c.Query("address")
			log.Println(userAddress)
			c.JSON(http.StatusOK, gin.H{
				"status": "pass",
			})
		})
	}
	authority := r.Group("/authority")
	{
		authority.GET("/users", func(c *gin.Context) {
			authorityAddress := c.Query("address")
			log.Println(authorityAddress)
			c.JSON(http.StatusOK, gin.H{
				"address":           "useraddress",
				"name":              "name",
				"gender":            "male",
				"country":           "china",
				"age":               18,
				"residency_address": "china",
			})
		})
		authority.POST("/certify", func(c *gin.Context) {
			var authorityCertify AuthorityCertify
			err := c.BindJSON(&authorityCertify)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err_code": 0001,
					"message":  err,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": true,
			})
		})
		authority.GET("/institutes", func(c *gin.Context) {
			authorityAddress := c.Query("authority_address")
			log.Println(authorityAddress)
			c.JSON(http.StatusOK, gin.H{
				"institute_address": "0xfsdsd",
				"user_address":      "0xfsdfsd",
			})
			return
		})
		authority.POST("/verify", func(c *gin.Context) {
			var authorityCertify AuthorityCertify
			err := c.BindJSON(&authorityCertify)
			if err != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"err_code": 0001,
					"message":  err,
				})
				return
			}
			c.JSON(http.StatusOK, gin.H{
				"status": true,
			})
		})
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run(":8081") // listen and serve on 0.0.0.0:8080
}

func SellerPublishData() {
	//publish data
	metaData := []byte("meta data test")
	proofData := [][]byte{{'4', '5', '6', '3'}, {'2', '2', '1'}}
	despData := []byte{'7', '8', '9', '3'}

	txParam := chainoperations.TransactParams{
		From:     common.HexToAddress(seller.Account.Address),
		Password: clientPassword,
	}

	cif.Publish(
		&txParam,
		big.NewInt(1000),
		metaData,
		proofData,
		2,
		despData,
	)
}

func onPublish(event events.Event) bool {
	fmt.Println("onpublish: ", event)

	publishId = event.Data.Get("publishId").(string)
	despDataId := event.Data.Get("despDataId").(string)
	price := event.Data.Get("price").(*big.Int)

	fmt.Println("despDataId:", despDataId, ", price:", price)
	return true
}

func onPurchase(event events.Event) bool {
	fmt.Println("onPurchase:", event)

	metaDataIdEncWithSeller = event.Data.Get("metaDataIdEncSeller").([]byte)
	buyerAddr := event.Data.Get("buyer").(common.Address)

	var err error
	metaDataIdEncWithBuyer, err = accounts.GetAMInstance().ReEncrypt(
		metaDataIdEncWithSeller,
		seller.Account.Address,
		buyerAddr.String(),
		clientPassword,
	)

	if err != nil {
		fmt.Println("failed to ReEncrypt meta data id with buyer's public key")
		return false
	}

	SubmitMetaDataIdEncWithBuyer(txId)
	return true
}
func SubmitMetaDataIdEncWithBuyer(txId *big.Int) {
	txParam := chainoperations.TransactParams{
		From:     common.HexToAddress(seller.Account.Address),
		Password: clientPassword,
	}
	err := cif.SubmitMetaDataIdEncWithBuyer(
		&txParam,
		txId,
		metaDataIdEncWithBuyer)
	fmt.Println("SubmitData:", txId, txParam)
	if err != nil {
		fmt.Println("failed to SubmitMetaDataIdEncWithBuyer, error:", err)
	}
}
func onClose(event events.Event) bool {
	fmt.Println("onClose:", event)
	return true
}

func onTransactionCreate(event events.Event) bool {
	fmt.Println("onTransactionCreated:", event)

	txId = event.Data.Get("transactionId").(*big.Int)

	return true
}

func onApprovalBuyerTransfer(event events.Event) bool {
	fmt.Println("onApprovalBuyerTransfer:", event)

	PrepareToBuy(publishId)
	return true
}

// 买方准备购买
func PrepareToBuy(publishId string) {
	txParam := chainoperations.TransactParams{
		From:     common.HexToAddress(buyer.Account.Address),
		Password: clientPassword,
	}
	err := cif.PrepareToBuy(&txParam, publishId)
	if err != nil {
		fmt.Println("failed to prepareToBuy, error:", err)
	}
}

func onPublishBuyer(event events.Event) bool {
	fmt.Println("onpublish: ", event)

	publishId = event.Data.Get("publishId").(string)
	despDataId := event.Data.Get("despDataId").(string)
	price := event.Data.Get("price").(*big.Int)

	fmt.Println("publishId", publishId, ", despDataId:", despDataId, ", price:", price)
	BuyerApproveTransfer()
	return true
}
func BuyerApproveTransfer() {
	txParam := chainoperations.TransactParams{
		From:     common.HexToAddress(buyer.Account.Address),
		Password: clientPassword,
	}
	err := cif.ApproveTransfer(&txParam,
		common.HexToAddress(protocolContractAddr),
		big.NewInt(1001))
	if err != nil {
		fmt.Println("BuyerApproveTransfer:", err)
	}
}
func onTransactionCreateBuyer(event events.Event) bool {
	fmt.Println("onTransactionCreated:", event)

	txId = event.Data.Get("transactionId").(*big.Int)
	proofIDs := event.Data.Get("proofIds").([][32]byte)

	for _, proofId := range proofIDs {
		proofId, err := cif.Bytes32ToIpfsHash(proofId)
		if err != nil {
			panic(err)
		}

		fmt.Println("proof id:", proofId)
	}

	Buy(txId)
	return true
}
func Buy(txId *big.Int) {
	txParam := chainoperations.TransactParams{
		From:     common.HexToAddress(buyer.Account.Address),
		Password: clientPassword,
	}
	err := cif.BuyData(&txParam, txId)
	if err != nil {
		fmt.Println("failed to buyData, error:", err)
	}
}

func onReadyForDownload(event events.Event) bool {
	fmt.Println("onReadyForDownload:", event)
	metaDataIdEncWithBuyer = event.Data.Get("metaDataIdEncBuyer").([]byte)

	metaDataId, err := accounts.GetAMInstance().Decrypt(
		metaDataIdEncWithBuyer,
		buyer.Account.Address,
		clientPassword)

	if err != nil {
		fmt.Println("failed to decrypt meta data id with buyer's private key", err)
		return false
	}

	fmt.Println("meta data id:", string(metaDataId))

	ConfirmDataTruth(txId)
	return true
}

func ConfirmDataTruth(txId *big.Int) {
	txParam := chainoperations.TransactParams{
		From:     common.HexToAddress(buyer.Account.Address),
		Password: clientPassword,
	}
	err := cif.ConfirmDataTruth(
		&txParam,
		txId,
		true)
	fmt.Println("ConfirmDataTruth:", txId)
	if err != nil {
		fmt.Println("failed to ConfirmDataTruth, error:", err)
	}
}
