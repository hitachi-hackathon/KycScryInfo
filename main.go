package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/scryinfo/iscap_demo/src/sdk"
	"github.com/scryinfo/iscap_demo/src/sdk/core/chainoperations"
	"github.com/scryinfo/iscap_demo/src/sdk/core/ethereum/events"
	"github.com/scryinfo/iscap_demo/src/sdk/scryclient"
	cif "github.com/scryinfo/iscap_demo/src/sdk/scryclient/chaininterfacewrapper"
	"github.com/scryinfo/iscap_demo/src/sdk/util/accounts"
	"github.com/spf13/viper"
	"math/big"
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
	fromBlock            = 0
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

func PrintfEthAndTokenByClient(client *scryclient.ScryClient) {
	fmt.Println(fmt.Sprintf("account %s has eth %d , token %d ", client.Account.Address, GetEthByClient(client), GetTokenByClient(client)))
}

func init()  {
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
	fromBlock = viper.GetInt("fromBlock")
	ipfsNodeAddr = viper.GetString("ipfsNodeAddr")
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.Run() // listen and serve on 0.0.0.0:8080
	
	wd, _ := os.Getwd()
	err := sdk.Init(ethNodeAddr, keyServiceAddr, protocolContractAddr, tokenContractAddr, fromBlock, ipfsNodeAddr, wd+"/testconsole.log", "scryapp1")
	if err != nil {
		fmt.Println("sdk init err!", err)
	}
	// 初始化seller
	seller = scryclient.NewScryClient("0x2008cc463061d385d87a294b2f3edce229f74b58")
	// 查询账户比特币和Token余额
	PrintfEthAndTokenByClient(seller)
	// 1.数据成功发布到IPFS和区块链上
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
	SellerPublishData()

	PrintfEthAndTokenByClient(seller)
	PrintfEthAndTokenByClient(buyer)
	time.Sleep(100000000000000)
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
	fmt.Println("Testing end")

	os.Exit(0)
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
