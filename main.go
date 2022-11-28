package main

import (
	"fabric-go-sdk/sdkInit"
	"fmt"
	"os"
	"time"
)

const (
	cc_name    = "simplecc"
	cc_version = "1.0.0"
)

var App sdkInit.Application

func main() {
	// init orgs information
	GOPATH := os.Getenv("GOPATH")
	orgs := []*sdkInit.OrgInfo{
		{
			OrgAdminUser:  "Admin",
			OrgName:       "Org1",
			OrgMspId:      "Org1MSP",
			OrgUser:       "User1",
			OrgPeerNum:    2,
			OrgAnchorFile: fmt.Sprintf("%s/src/fabric-go-sdk/fixtures/channel-artifacts/Org1MSPanchors.tx", GOPATH),
		},
	}

	// init sdk env info
	info := sdkInit.SdkEnvInfo{
		ChannelID:        "mychannel",
		ChannelConfig:    fmt.Sprintf("%s/src/fabric-go-sdk/fixtures/channel-artifacts/channel.tx", GOPATH),
		Orgs:             orgs,
		OrdererAdminUser: "Admin",
		OrdererOrgName:   "OrdererOrg",
		OrdererEndpoint:  "orderer.example.com",
		ChaincodeID:      cc_name,
		ChaincodePath:    fmt.Sprintf("%s/src/fabric-go-sdk/chaincode/", GOPATH),
		ChaincodeVersion: cc_version,
	}

	// sdk setup
	sdk, err := sdkInit.Setup("config.yaml", &info)
	if err != nil {
		fmt.Println(">> SDK setup error:", err)
		os.Exit(-1)
	}

	// create channel and join
	if err := sdkInit.CreateAndJoinChannel(&info); err != nil {
		fmt.Println(">> Create channel and join error:", err)
		os.Exit(-1)
	}

	// create chaincode lifecycle
	if err := sdkInit.CreateCCLifecycle(&info, 1, false, sdk); err != nil {
		fmt.Println(">> create chaincode lifecycle error: %v", err)
		os.Exit(-1)
	}

	// invoke chaincode set status
	fmt.Println(">> set status by invoking chaincode......")

	if err := info.InitService(info.ChaincodeID, info.ChannelID, info.Orgs[0], sdk); err != nil {

		fmt.Println("InitService successful")
		os.Exit(-1)
	}

	App = sdkInit.Application{
		SdkEnvInfo: &info,
	}
	fmt.Println(">> chaincode set up finished")

	defer info.EvClient.Unregister(sdkInit.BlockListener(info.EvClient))
	defer info.EvClient.Unregister(sdkInit.ChainCodeEventListener(info.EvClient, info.ChaincodeID))

	a := []string{"set", "ID1", "123"}
	ret, err := App.Set(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- add row1　--->：", ret)

	a = []string{"set", "ID2", "456"}
	ret, err = App.Set(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- add row2　--->：", ret)

	a = []string{"set", "ID3", "789"}
	ret, err = App.Set(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- add row3　--->：", ret)

	a = []string{"get", "ID3"}
	response, err := App.Get(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- get row3　--->：", response)

	a = []string{"get", "ID2"}
	response, err = App.Get(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- get row2　--->：", response)
	a = []string{"get", "ID1"}
	response, err = App.Get(a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("<--- get row1　--->：", response)

	time.Sleep(time.Second * 10)

}
