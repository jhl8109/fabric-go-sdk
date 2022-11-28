# fabric-go-sdk
Hyperledger Fabric 2.x 기반
## 순서
### 설정
`GOPATH` = `home/jeho/go`

```
cd $GOPATH/src && git clone https://github.com/sxguan/fabric-go-sdk.git
```

### 네트워크 시작 

```
cd ./fabric-go-sdk/fixtures/ && docker-compose up -d
```

### SDK 실행

```
cd .. && go build && ./fabric-go-sdk
```
```
>> start creating channel...
 >>>>Update anchor node configuration with each org's admin identity... 
 >>>>Update anchor node configuration with each org's admin identity completed 
>> channel created successfully
 >> Join channel...
>> joined the channel successfully
>> Start packaging chaincode...
>> packaged chaincode successfully
>> Start installing chaincode...
>> chaincode installed successfully
>> organization approved smart contract definition...
>>> chaincode approved by Org1 peers:
        grpcs://127.0.0.1:7051
        grpcs://127.0.0.1:9051
>> organization approved smart contract definition is complete
>> Check if smart contract is ready...
LifecycleCheckCCCommitReadiness cc = simplecc, = {map[Org1MSP:true]}
LifecycleCheckCCCommitReadiness cc = simplecc, = {map[Org1MSP:true]}
>> smart contract is ready
>> commit chaincode ......
>> smart contract definition submitted
>> call smart contract initialization method...
>> complete smart contract initialization
>> set status by invoking chaincode......
>> chaincode set up finished
2022/11/28 10:34:52 Registered block event
<--- add row1　--->： fe986763b2b887e65372c8228638455e82821501a670c38843337bada51ce415
2022/11/28 10:34:55 Receive cc event, ccid: simplecc 
eventName: chaincode-event
payload: {"EventName":"set"} 
txid: fe986763b2b887e65372c8228638455e82821501a670c38843337bada51ce415 
block: 5 
sourceURL: grpcs://127.0.0.1:7051
2022/11/28 10:34:55 Receive block event:
SourceURL: grpcs://127.0.0.1:7051
Number: 5
Hash: 8689e4914bba342e7a4273e529d2dce9d10cc66fd2818a57282714bfb65352c4
PreviousHash: 2ffde7229dc9ec1d55ff0615cc5cea774fa9557512f7b8e68e28b163e386b27a

<--- add row2　--->： 0da650f38af9a36a694a310a3c6612dca603a3a47e8849f4de17db4909ada965
2022/11/28 10:34:58 Receive cc event, ccid: simplecc 
eventName: chaincode-event
payload: {"EventName":"set"} 
txid: 0da650f38af9a36a694a310a3c6612dca603a3a47e8849f4de17db4909ada965 
block: 6 
sourceURL: grpcs://127.0.0.1:7051
2022/11/28 10:34:58 Receive block event:
SourceURL: grpcs://127.0.0.1:7051
Number: 6
Hash: 1f627b10905b1acae8722400c7a09839ce4573b65fc205e2e3f8ca0bde2e0ec3
PreviousHash: 910985f9715d406c444163448f7df09307d7e8dae5292548c069f1a2f2fd9485

<--- add row3　--->： be26f09564a3e6ef49ab96381354e59b95995b7dc143e2d25e6ee1ba15be39f2
2022/11/28 10:35:00 Receive cc event, ccid: simplecc 
eventName: chaincode-event
payload: {"EventName":"set"} 
txid: be26f09564a3e6ef49ab96381354e59b95995b7dc143e2d25e6ee1ba15be39f2 
block: 7 
sourceURL: grpcs://127.0.0.1:7051
2022/11/28 10:35:00 Receive block event:
SourceURL: grpcs://127.0.0.1:7051
Number: 7
Hash: e4760d772420b6bde1cb0af77c725c0feb36dfdc18118e4a57b7e46ef61fdafa
PreviousHash: 231d9f5b4c26299feeaa3cad09280424cc012df1b6441258c3c979648341bc73

<--- get row3　--->： 789
<--- get row2　--->： 456
<--- get row1　--->： 123
```

