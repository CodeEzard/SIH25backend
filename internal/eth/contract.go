package eth

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"strings"

	"vericred/internal/eth/build"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Contract struct {
	privateKeyHex string
	rpcURL string
	cAddress string
}

type InitializeContractVars struct {
	instance *build.Build
	auth *bind.TransactOpts
}

type ContractFunctions struct {}

var C = &Contract{
	privateKeyHex: "e27fc11f6468ca7b5916ebfd853b0e92aca9a5899af26414bae8d22bfd55c5d4",
	rpcURL: "https://sepolia.infura.io/v3/2a159ca7304a4df4ade0d0ccf9ce70ef",
	cAddress: "0xDE5C084a7959533893954BA072895B53fE1E7486",
}

func (c Contract) Deploy() string {

	privateKeyHex := C.privateKeyHex
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatal("Invalid private key")
	}

	client, err := ethclient.Dial(C.rpcURL)
	if err != nil {
		log.Fatal("❌ Failed to connect to the sepolia network.", err)
	}

	log.Println("Successfully logged in to sepolia network")

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)

	log.Println("contex.Background():", context.Background())
	chainID, _ := client.NetworkID(context.Background())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.From = fromAddress
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)


	address, tx, _, err := build.DeployBuild(auth, client)
	if err != nil {
		log.Fatal("Failed to deploy:", err)
	}

	fmt.Println("Contract deployed at: ", address.Hex())
	fmt.Println("Deployment TX:", tx.Hash().Hex())

	log.Println("Waiting for the for Contract deployment transaction...")
	_, err = bind.WaitDeployed(context.Background(), client, tx)
	if err != nil {
		log.Fatal("Wait failed: ", err)
	}
	log.Println("Successfully waited for the deployment.")	
	// cAddress := common.HexToAddress("0x2a67320e4fc9d8E072021575e5D0450ED5ae316e")

	// instance, err = build.NewBuild(cAddress, client)
	// if err != nil {
	// 	log.Fatal("❌ Failed to bind to Contract:", err)
	// }
	// log.Println("Instance: ", instance)
	return address.Hex()

	// userAddr := common.HexToAddress("0x2a67320e4fc9d8E072021575e5D0450ED5ae316e")
	// _, err = instance.MintDoc(auth, userAddr, "https://chocolate-electoral-parrot-267.mypinata.cloud/ipfs/bafkreifgophqzx5pw2uvukbozh7fxcpo5fdx2rprh73t3l72pgszekeita")
	// if err != nil {
	// 	log.Fatal("❌ Failed to call AllOrgs:", err)
	// }

	// tokenId, err := instance.CurrentTokenId(nil)
	// if err != nil {
	// 	log.Fatal("Failed to call currenttokenid: ", err)
	// }

	// fmt.Println("TokenID:", tokenId)
}

func (icv *InitializeContractVars) InitializeViewContracts() {

	log.Println("InitializeContract function called.")
	client, err := ethclient.Dial(C.rpcURL)
	if err != nil {
		log.Fatal("❌ Failed to connect to the sepolia network.", err)
	}
	log.Println("Connected to the sepolia network.")

	cAddress := common.HexToAddress(C.cAddress)
	instance, err := build.NewBuild(cAddress, client)
	if err != nil {
		log.Fatal("Some error occurred when calling NewBuild: ", err)
	}
	icv.instance = instance
	//	userAddress := common.HexToAddress(uAddress)
//	fmt.Println(userAddress)

	// check, err := instance.AllOrgs(nil)
	// if err != nil {
	// 	log.Fatal("Failed to call currentTokenId: ", err)
	// }
	
	// fmt.Println("token ID: ", check)
}

func (cf ContractFunctions) CurrentTokenID() *big.Int {
	var icv = &InitializeContractVars{}
	icv.InitializeViewContracts()

	log.Println("CurrentTokenID function called.")
	tokenID, err := icv.instance.TokenIds(nil)
	if err != nil {
		log.Fatal("Failed to call tokenID: ", err)
	}
	return tokenID
}


func (cf ContractFunctions) AllOrgs() []common.Address {
	var icv = &InitializeContractVars{}
	icv.InitializeViewContracts()

	log.Println("AllOrgs() function called.")
	allOrgs, err := icv.instance.AllOrgs(nil)
	if err != nil {
		log.Fatal("Some error occurred when calling the AllOrgs function: ", err)
	}

	for _, org := range allOrgs {
		fmt.Println("Org Address:", org.Hex())
	}
	return allOrgs
}

func (cf ContractFunctions) IsVerfiedOrg(orgAddress string) {
	var icv = &InitializeContractVars{}
	icv.InitializeViewContracts()

	log.Println("AllOrgs() function called.")
	orgHexToAddress := common.HexToAddress(orgAddress)
	check, err := icv.instance.IsVerfiedOrg(nil, orgHexToAddress)
	if err != nil {
		log.Fatal("Some error occurred when calling the AllOrgs function: ", err)
	}
	if (check) {
		log.Printf("Org %s is verified.", orgAddress)
	} else {
		log.Printf("Org %s is not verified.", orgAddress)
	}
}

func (cf ContractFunctions) TokenURI(id *big.Int) {
	var icv = &InitializeContractVars{}
	icv.InitializeViewContracts()

	log.Println("TokenURI function called.")
	url, err := icv.instance.TokenURI(nil, id)
	if err != nil {
		fmt.Println("Error calling TokenURI: ", err)
	}
	fmt.Println("URL: ", url)
}
 
func (icv *InitializeContractVars) InitializeTrxnContracts() {
	privateKeyHex := C.privateKeyHex
	privateKey, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		log.Fatal("Invalid private key")
	}
	client, err := ethclient.Dial(C.rpcURL)
	if err != nil {
		log.Fatal("❌ Failed to connect to the sepolia network.", err)
	}
	log.Println("Successfully logged in to sepolia network")

	publicKey := privateKey.Public().(*ecdsa.PublicKey)
	fromAddress := crypto.PubkeyToAddress(*publicKey)
	chainID, _ := client.NetworkID(context.Background())

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}
	auth.From = fromAddress
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(3000000)

	icv.auth = auth
	cAddress := common.HexToAddress(C.cAddress)
	instance, err := build.NewBuild(cAddress, client)
	if err != nil {
		log.Fatal("Some error occurred when calling NewBuild: ", err)
	}
	icv.instance = instance
	log.Println("Instance has been set.")
}


func (cf ContractFunctions) NewOrg(uAddress string) {
	var icv = InitializeContractVars{}
	icv.InitializeTrxnContracts()

	userAddr := common.HexToAddress(uAddress)
	_, err := icv.instance.NewOrg(icv.auth, userAddr)
	if err != nil {
		log.Fatal("Failed to call NewOrg: ", err)
	}
	log.Println("New org added successfully.")
}

func (cf ContractFunctions) MintDoc(uAddress string, tokenURI string) {
	var icv = InitializeContractVars{}
	icv.InitializeTrxnContracts()
	
	userAddr := common.HexToAddress(uAddress)
	_, err := icv.instance.MintDoc(icv.auth, userAddr, tokenURI)
	if err != nil {
		log.Fatal("Failed to call MintDoc: ", err)
	}
	log.Println("Minted successfully")
}

