// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package foodchain

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// FoodchainMetaData contains all meta data concerning the Foodchain contract.
var FoodchainMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"createdTime\",\"type\":\"string\"}],\"name\":\"Create\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"Delete\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"from\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"to\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AccountDeposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"AccountInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"sender\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"receiver\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AccountTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"AccountWithdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phoneNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"createdTime\",\"type\":\"string\"}],\"name\":\"CreateAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"DeleteAccount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phoneNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"createdTime\",\"type\":\"string\"}],\"name\":\"UpdateAccount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"accountAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"phoneNumber\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"status\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"createdTime\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalAccounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611e59806100206000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063be405fa611610066578063be405fa61461011c578063d695136d14610138578063da6841aa1461016e578063ea3d5e481461018c578063f2a40db8146101c257610093565b806320290b3f14610098578063243fff8f146100c85780636e03dd9a146100e4578063967448cd14610100575b600080fd5b6100b260048036038101906100ad9190611759565b6101f8565b6040516100bf919061187a565b60405180910390f35b6100e260048036038101906100dd9190611895565b610337565b005b6100fe60048036038101906100f99190611895565b6104c0565b005b61011a600480360381019061011591906118d5565b610591565b005b61013660048036038101906101319190611928565b610767565b005b610152600480360381019061014d9190611759565b610928565b60405161016597969594939291906119dd565b60405180910390f35b610176610dd3565b604051610183919061187a565b60405180910390f35b6101a660048036038101906101a19190611928565b610ddf565b6040516101b997969594939291906119dd565b60405180910390f35b6101dc60048036038101906101d79190611928565b611148565b6040516101ef97969594939291906119dd565b60405180910390f35b6000806040518060e001604052808a815260200189815260200188815260200187815260200186815260200185815260200184815250908060018154018082558091505060019003906000526020600020906007020160009091909190915060008201518160000155602082015181600101908051906020019061027d929190611459565b50604082015181600201908051906020019061029a929190611459565b5060608201518160030190805190602001906102b7929190611459565b506080820151816004015560a0820151816005015560c08201518160060190805190602001906102e8929190611459565b5050507fbba37318c149b1db7d49c422b47f2a61ba899714cb9c7377168ec21d7edfaae1888360405161031c929190611a68565b60405180910390a16000805490509050979650505050505050565b6000610342836113ba565b90506000821015610388576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161037f90611ae4565b60405180910390fd5b6000826000838154811061039f5761039e611b04565b5b9060005260206000209060070201600401546103bb9190611b62565b10156103fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f390611be2565b60405180910390fd5b816000828154811061041157610410611b04565b5b906000526020600020906007020160040160008282546104319190611b62565b925050819055507f56ca301a9219608c91e7bcee90e083c19671d2cdcc96752c7af291cee5f9c8c86000828154811061046d5761046c611b04565b5b9060005260206000209060070201600001546000838154811061049357610492611b04565b5b9060005260206000209060070201600401546040516104b3929190611c02565b60405180910390a1505050565b60006104cb836113ba565b905081600082815481106104e2576104e1611b04565b5b906000526020600020906007020160040160008282546105029190611c2b565b925050819055507fa3af609bf46297028ce551832669030f9effef2b02606d02cbbcc40fe6b47c556000828154811061053e5761053d611b04565b5b9060005260206000209060070201600001546000838154811061056457610563611b04565b5b906000526020600020906007020160040154604051610584929190611c02565b60405180910390a1505050565b600061059c846113ba565b905060006105a9846113ba565b905060008310156105ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105e690611ae4565b60405180910390fd5b6000836000848154811061060657610605611b04565b5b9060005260206000209060070201600401546106229190611b62565b1015610663576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065a90611be2565b60405180910390fd5b826000838154811061067857610677611b04565b5b906000526020600020906007020160040160008282546106989190611b62565b9250508190555082600082815481106106b4576106b3611b04565b5b906000526020600020906007020160040160008282546106d49190611c2b565b925050819055507faf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6600083815481106107105761070f611b04565b5b9060005260206000209060070201600001546000838154811061073657610735611b04565b5b9060005260206000209060070201600001548560405161075893929190611c81565b60405180910390a15050505050565b6000610772826113ba565b9050600060016000805490506107889190611b62565b8154811061079957610798611b04565b5b9060005260206000209060070201600082815481106107bb576107ba611b04565b5b90600052602060002090600702016000820154816000015560018201816001019080546107e790611ce7565b6107f29291906114df565b50600282018160020190805461080790611ce7565b6108129291906114df565b50600382018160030190805461082790611ce7565b6108329291906114df565b506004820154816004015560058201548160050155600682018160060190805461085b90611ce7565b6108669291906114df565b50905050600080548061087c5761087b611d19565b5b60019003818190600052602060002090600702016000808201600090556001820160006108a9919061156c565b6002820160006108b9919061156c565b6003820160006108c9919061156c565b600482016000905560058201600090556006820160006108e9919061156c565b505090557f7d815d575a49d6d9a069482c25b0fdb39914ccf1bba8aaea4cfe20322f8f68778260405161091c919061187a565b60405180910390a15050565b600060608060606000806060600061093f8f6113ba565b90508d6000828154811061095657610955611b04565b5b9060005260206000209060070201600101908051906020019061097a929190611459565b508c600082815481106109905761098f611b04565b5b906000526020600020906007020160020190805190602001906109b4929190611459565b508b600082815481106109ca576109c9611b04565b5b906000526020600020906007020160030190805190602001906109ee929190611459565b508a60008281548110610a0457610a03611b04565b5b9060005260206000209060070201600401819055508960008281548110610a2e57610a2d611b04565b5b9060005260206000209060070201600501819055508860008281548110610a5857610a57611b04565b5b90600052602060002090600702016006019080519060200190610a7c929190611459565b5060008181548110610a9157610a90611b04565b5b90600052602060002090600702016000015460008281548110610ab757610ab6611b04565b5b906000526020600020906007020160010160008381548110610adc57610adb611b04565b5b906000526020600020906007020160020160008481548110610b0157610b00611b04565b5b906000526020600020906007020160030160008581548110610b2657610b25611b04565b5b90600052602060002090600702016004015460008681548110610b4c57610b4b611b04565b5b90600052602060002090600702016005015460008781548110610b7257610b71611b04565b5b9060005260206000209060070201600601858054610b8f90611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610bbb90611ce7565b8015610c085780601f10610bdd57610100808354040283529160200191610c08565b820191906000526020600020905b815481529060010190602001808311610beb57829003601f168201915b50505050509550848054610c1b90611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610c4790611ce7565b8015610c945780601f10610c6957610100808354040283529160200191610c94565b820191906000526020600020905b815481529060010190602001808311610c7757829003601f168201915b50505050509450838054610ca790611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610cd390611ce7565b8015610d205780601f10610cf557610100808354040283529160200191610d20565b820191906000526020600020905b815481529060010190602001808311610d0357829003601f168201915b50505050509350808054610d3390611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610d5f90611ce7565b8015610dac5780601f10610d8157610100808354040283529160200191610dac565b820191906000526020600020905b815481529060010190602001808311610d8f57829003601f168201915b50505050509050975097509750975097509750975050975097509750975097509750979050565b60008080549050905090565b6000606080606060008060606000610df6896113ba565b905060008181548110610e0c57610e0b611b04565b5b90600052602060002090600702016000015460008281548110610e3257610e31611b04565b5b906000526020600020906007020160010160008381548110610e5757610e56611b04565b5b906000526020600020906007020160020160008481548110610e7c57610e7b611b04565b5b906000526020600020906007020160030160008581548110610ea157610ea0611b04565b5b90600052602060002090600702016004015460008681548110610ec757610ec6611b04565b5b90600052602060002090600702016005015460008781548110610eed57610eec611b04565b5b9060005260206000209060070201600601858054610f0a90611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610f3690611ce7565b8015610f835780601f10610f5857610100808354040283529160200191610f83565b820191906000526020600020905b815481529060010190602001808311610f6657829003601f168201915b50505050509550848054610f9690611ce7565b80601f0160208091040260200160405190810160405280929190818152602001828054610fc290611ce7565b801561100f5780601f10610fe45761010080835404028352916020019161100f565b820191906000526020600020905b815481529060010190602001808311610ff257829003601f168201915b5050505050945083805461102290611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461104e90611ce7565b801561109b5780601f106110705761010080835404028352916020019161109b565b820191906000526020600020905b81548152906001019060200180831161107e57829003601f168201915b505050505093508080546110ae90611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546110da90611ce7565b80156111275780601f106110fc57610100808354040283529160200191611127565b820191906000526020600020905b81548152906001019060200180831161110a57829003601f168201915b50505050509050975097509750975097509750975050919395979092949650565b6000818154811061115857600080fd5b906000526020600020906007020160009150905080600001549080600101805461118190611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546111ad90611ce7565b80156111fa5780601f106111cf576101008083540402835291602001916111fa565b820191906000526020600020905b8154815290600101906020018083116111dd57829003601f168201915b50505050509080600201805461120f90611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461123b90611ce7565b80156112885780601f1061125d57610100808354040283529160200191611288565b820191906000526020600020905b81548152906001019060200180831161126b57829003601f168201915b50505050509080600301805461129d90611ce7565b80601f01602080910402602001604051908101604052809291908181526020018280546112c990611ce7565b80156113165780601f106112eb57610100808354040283529160200191611316565b820191906000526020600020905b8154815290600101906020018083116112f957829003601f168201915b50505050509080600401549080600501549080600601805461133790611ce7565b80601f016020809104026020016040519081016040528092919081815260200182805461136390611ce7565b80156113b05780601f10611385576101008083540402835291602001916113b0565b820191906000526020600020905b81548152906001019060200180831161139357829003601f168201915b5050505050905087565b600080600090505b6000805490508110156114185782600082815481106113e4576113e3611b04565b5b90600052602060002090600702016000015414156114055780915050611454565b808061141090611d48565b9150506113c2565b506040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161144b90611ddd565b60405180910390fd5b919050565b82805461146590611ce7565b90600052602060002090601f01602090048101928261148757600085556114ce565b82601f106114a057805160ff19168380011785556114ce565b828001600101855582156114ce579182015b828111156114cd5782518255916020019190600101906114b2565b5b5090506114db91906115ac565b5090565b8280546114eb90611ce7565b90600052602060002090601f01602090048101928261150d576000855561155b565b82601f1061151e578054855561155b565b8280016001018555821561155b57600052602060002091601f016020900482015b8281111561155a57825482559160010191906001019061153f565b5b50905061156891906115ac565b5090565b50805461157890611ce7565b6000825580601f1061158a57506115a9565b601f0160209004906000526020600020908101906115a891906115ac565b5b50565b5b808211156115c55760008160009055506001016115ad565b5090565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6115f0816115dd565b81146115fb57600080fd5b50565b60008135905061160d816115e7565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6116668261161d565b810181811067ffffffffffffffff821117156116855761168461162e565b5b80604052505050565b60006116986115c9565b90506116a4828261165d565b919050565b600067ffffffffffffffff8211156116c4576116c361162e565b5b6116cd8261161d565b9050602081019050919050565b82818337600083830152505050565b60006116fc6116f7846116a9565b61168e565b90508281526020810184848401111561171857611717611618565b5b6117238482856116da565b509392505050565b600082601f8301126117405761173f611613565b5b81356117508482602086016116e9565b91505092915050565b600080600080600080600060e0888a031215611778576117776115d3565b5b60006117868a828b016115fe565b975050602088013567ffffffffffffffff8111156117a7576117a66115d8565b5b6117b38a828b0161172b565b965050604088013567ffffffffffffffff8111156117d4576117d36115d8565b5b6117e08a828b0161172b565b955050606088013567ffffffffffffffff811115611801576118006115d8565b5b61180d8a828b0161172b565b945050608061181e8a828b016115fe565b93505060a061182f8a828b016115fe565b92505060c088013567ffffffffffffffff8111156118505761184f6115d8565b5b61185c8a828b0161172b565b91505092959891949750929550565b611874816115dd565b82525050565b600060208201905061188f600083018461186b565b92915050565b600080604083850312156118ac576118ab6115d3565b5b60006118ba858286016115fe565b92505060206118cb858286016115fe565b9150509250929050565b6000806000606084860312156118ee576118ed6115d3565b5b60006118fc868287016115fe565b935050602061190d868287016115fe565b925050604061191e868287016115fe565b9150509250925092565b60006020828403121561193e5761193d6115d3565b5b600061194c848285016115fe565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561198f578082015181840152602081019050611974565b8381111561199e576000848401525b50505050565b60006119af82611955565b6119b98185611960565b93506119c9818560208601611971565b6119d28161161d565b840191505092915050565b600060e0820190506119f2600083018a61186b565b8181036020830152611a0481896119a4565b90508181036040830152611a1881886119a4565b90508181036060830152611a2c81876119a4565b9050611a3b608083018661186b565b611a4860a083018561186b565b81810360c0830152611a5a81846119a4565b905098975050505050505050565b6000604082019050611a7d600083018561186b565b8181036020830152611a8f81846119a4565b90509392505050565b7f706c6561736520696e7365727420612076616c696420616d6f756e7420210000600082015250565b6000611ace601e83611960565b9150611ad982611a98565b602082019050919050565b60006020820190508181036000830152611afd81611ac1565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000611b6d826115dd565b9150611b78836115dd565b925082821015611b8b57611b8a611b33565b5b828203905092915050565b7f796f75722062616c616e6365206973206e6f7420656e6f756768202100000000600082015250565b6000611bcc601c83611960565b9150611bd782611b96565b602082019050919050565b60006020820190508181036000830152611bfb81611bbf565b9050919050565b6000604082019050611c17600083018561186b565b611c24602083018461186b565b9392505050565b6000611c36826115dd565b9150611c41836115dd565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff03821115611c7657611c75611b33565b5b828201905092915050565b6000606082019050611c96600083018661186b565b611ca3602083018561186b565b611cb0604083018461186b565b949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680611cff57607f821691505b60208210811415611d1357611d12611cb8565b5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fd5b6000611d53826115dd565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415611d8657611d85611b33565b5b600182019050919050565b7f4163636f756e7420646f6573206e6f7420657869747300000000000000000000600082015250565b6000611dc7601683611960565b9150611dd282611d91565b602082019050919050565b60006020820190508181036000830152611df681611dba565b905091905056fea2646970667358221220430c2fac0bfef4408ec637d5c92f46cb58a8c45eb895bfb582db77c453bdba0f64736f6c637828302e382e31322d646576656c6f702e323032322e322e31322b636f6d6d69742e31323130633365360059",
}

// FoodchainABI is the input ABI used to generate the binding from.
// Deprecated: Use FoodchainMetaData.ABI instead.
var FoodchainABI = FoodchainMetaData.ABI

// FoodchainBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FoodchainMetaData.Bin instead.
var FoodchainBin = FoodchainMetaData.Bin

// DeployFoodchain deploys a new Ethereum contract, binding an instance of Foodchain to it.
func DeployFoodchain(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Foodchain, error) {
	parsed, err := FoodchainMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FoodchainBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Foodchain{FoodchainCaller: FoodchainCaller{contract: contract}, FoodchainTransactor: FoodchainTransactor{contract: contract}, FoodchainFilterer: FoodchainFilterer{contract: contract}}, nil
}

// Foodchain is an auto generated Go binding around an Ethereum contract.
type Foodchain struct {
	FoodchainCaller     // Read-only binding to the contract
	FoodchainTransactor // Write-only binding to the contract
	FoodchainFilterer   // Log filterer for contract events
}

// FoodchainCaller is an auto generated read-only Go binding around an Ethereum contract.
type FoodchainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodchainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FoodchainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodchainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FoodchainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FoodchainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FoodchainSession struct {
	Contract     *Foodchain        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FoodchainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FoodchainCallerSession struct {
	Contract *FoodchainCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// FoodchainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FoodchainTransactorSession struct {
	Contract     *FoodchainTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// FoodchainRaw is an auto generated low-level Go binding around an Ethereum contract.
type FoodchainRaw struct {
	Contract *Foodchain // Generic contract binding to access the raw methods on
}

// FoodchainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FoodchainCallerRaw struct {
	Contract *FoodchainCaller // Generic read-only contract binding to access the raw methods on
}

// FoodchainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FoodchainTransactorRaw struct {
	Contract *FoodchainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFoodchain creates a new instance of Foodchain, bound to a specific deployed contract.
func NewFoodchain(address common.Address, backend bind.ContractBackend) (*Foodchain, error) {
	contract, err := bindFoodchain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Foodchain{FoodchainCaller: FoodchainCaller{contract: contract}, FoodchainTransactor: FoodchainTransactor{contract: contract}, FoodchainFilterer: FoodchainFilterer{contract: contract}}, nil
}

// NewFoodchainCaller creates a new read-only instance of Foodchain, bound to a specific deployed contract.
func NewFoodchainCaller(address common.Address, caller bind.ContractCaller) (*FoodchainCaller, error) {
	contract, err := bindFoodchain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FoodchainCaller{contract: contract}, nil
}

// NewFoodchainTransactor creates a new write-only instance of Foodchain, bound to a specific deployed contract.
func NewFoodchainTransactor(address common.Address, transactor bind.ContractTransactor) (*FoodchainTransactor, error) {
	contract, err := bindFoodchain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FoodchainTransactor{contract: contract}, nil
}

// NewFoodchainFilterer creates a new log filterer instance of Foodchain, bound to a specific deployed contract.
func NewFoodchainFilterer(address common.Address, filterer bind.ContractFilterer) (*FoodchainFilterer, error) {
	contract, err := bindFoodchain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FoodchainFilterer{contract: contract}, nil
}

// bindFoodchain binds a generic wrapper to an already deployed contract.
func bindFoodchain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(FoodchainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foodchain *FoodchainRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foodchain.Contract.FoodchainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foodchain *FoodchainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foodchain.Contract.FoodchainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foodchain *FoodchainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foodchain.Contract.FoodchainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Foodchain *FoodchainCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Foodchain.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Foodchain *FoodchainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Foodchain.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Foodchain *FoodchainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Foodchain.Contract.contract.Transact(opts, method, params...)
}

// AccountInfo is a free data retrieval call binding the contract method 0xea3d5e48.
//
// Solidity: function AccountInfo(uint256 id) view returns(uint256, string, string, string, uint256, uint256, string)
func (_Foodchain *FoodchainCaller) AccountInfo(opts *bind.CallOpts, id *big.Int) (*big.Int, string, string, string, *big.Int, *big.Int, string, error) {
	var out []interface{}
	err := _Foodchain.contract.Call(opts, &out, "AccountInfo", id)

	if err != nil {
		return *new(*big.Int), *new(string), *new(string), *new(string), *new(*big.Int), *new(*big.Int), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	out6 := *abi.ConvertType(out[6], new(string)).(*string)

	return out0, out1, out2, out3, out4, out5, out6, err

}

// AccountInfo is a free data retrieval call binding the contract method 0xea3d5e48.
//
// Solidity: function AccountInfo(uint256 id) view returns(uint256, string, string, string, uint256, uint256, string)
func (_Foodchain *FoodchainSession) AccountInfo(id *big.Int) (*big.Int, string, string, string, *big.Int, *big.Int, string, error) {
	return _Foodchain.Contract.AccountInfo(&_Foodchain.CallOpts, id)
}

// AccountInfo is a free data retrieval call binding the contract method 0xea3d5e48.
//
// Solidity: function AccountInfo(uint256 id) view returns(uint256, string, string, string, uint256, uint256, string)
func (_Foodchain *FoodchainCallerSession) AccountInfo(id *big.Int) (*big.Int, string, string, string, *big.Int, *big.Int, string, error) {
	return _Foodchain.Contract.AccountInfo(&_Foodchain.CallOpts, id)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime)
func (_Foodchain *FoodchainCaller) Accounts(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	AccountAddress string
	PhoneNumber    string
	Balance        *big.Int
	Status         *big.Int
	CreatedTime    string
}, error) {
	var out []interface{}
	err := _Foodchain.contract.Call(opts, &out, "accounts", arg0)

	outstruct := new(struct {
		Id             *big.Int
		Name           string
		AccountAddress string
		PhoneNumber    string
		Balance        *big.Int
		Status         *big.Int
		CreatedTime    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.AccountAddress = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.PhoneNumber = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.Balance = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Status = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.CreatedTime = *abi.ConvertType(out[6], new(string)).(*string)

	return *outstruct, err

}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime)
func (_Foodchain *FoodchainSession) Accounts(arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	AccountAddress string
	PhoneNumber    string
	Balance        *big.Int
	Status         *big.Int
	CreatedTime    string
}, error) {
	return _Foodchain.Contract.Accounts(&_Foodchain.CallOpts, arg0)
}

// Accounts is a free data retrieval call binding the contract method 0xf2a40db8.
//
// Solidity: function accounts(uint256 ) view returns(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime)
func (_Foodchain *FoodchainCallerSession) Accounts(arg0 *big.Int) (struct {
	Id             *big.Int
	Name           string
	AccountAddress string
	PhoneNumber    string
	Balance        *big.Int
	Status         *big.Int
	CreatedTime    string
}, error) {
	return _Foodchain.Contract.Accounts(&_Foodchain.CallOpts, arg0)
}

// GetTotalAccounts is a free data retrieval call binding the contract method 0xda6841aa.
//
// Solidity: function getTotalAccounts() view returns(uint256)
func (_Foodchain *FoodchainCaller) GetTotalAccounts(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Foodchain.contract.Call(opts, &out, "getTotalAccounts")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalAccounts is a free data retrieval call binding the contract method 0xda6841aa.
//
// Solidity: function getTotalAccounts() view returns(uint256)
func (_Foodchain *FoodchainSession) GetTotalAccounts() (*big.Int, error) {
	return _Foodchain.Contract.GetTotalAccounts(&_Foodchain.CallOpts)
}

// GetTotalAccounts is a free data retrieval call binding the contract method 0xda6841aa.
//
// Solidity: function getTotalAccounts() view returns(uint256)
func (_Foodchain *FoodchainCallerSession) GetTotalAccounts() (*big.Int, error) {
	return _Foodchain.Contract.GetTotalAccounts(&_Foodchain.CallOpts)
}

// AccountDeposit is a paid mutator transaction binding the contract method 0x6e03dd9a.
//
// Solidity: function AccountDeposit(uint256 id, uint256 amount) returns()
func (_Foodchain *FoodchainTransactor) AccountDeposit(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.contract.Transact(opts, "AccountDeposit", id, amount)
}

// AccountDeposit is a paid mutator transaction binding the contract method 0x6e03dd9a.
//
// Solidity: function AccountDeposit(uint256 id, uint256 amount) returns()
func (_Foodchain *FoodchainSession) AccountDeposit(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.Contract.AccountDeposit(&_Foodchain.TransactOpts, id, amount)
}

// AccountDeposit is a paid mutator transaction binding the contract method 0x6e03dd9a.
//
// Solidity: function AccountDeposit(uint256 id, uint256 amount) returns()
func (_Foodchain *FoodchainTransactorSession) AccountDeposit(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.Contract.AccountDeposit(&_Foodchain.TransactOpts, id, amount)
}

// AccountTransfer is a paid mutator transaction binding the contract method 0x967448cd.
//
// Solidity: function AccountTransfer(uint256 sender, uint256 receiver, uint256 amount) returns()
func (_Foodchain *FoodchainTransactor) AccountTransfer(opts *bind.TransactOpts, sender *big.Int, receiver *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.contract.Transact(opts, "AccountTransfer", sender, receiver, amount)
}

// AccountTransfer is a paid mutator transaction binding the contract method 0x967448cd.
//
// Solidity: function AccountTransfer(uint256 sender, uint256 receiver, uint256 amount) returns()
func (_Foodchain *FoodchainSession) AccountTransfer(sender *big.Int, receiver *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.Contract.AccountTransfer(&_Foodchain.TransactOpts, sender, receiver, amount)
}

// AccountTransfer is a paid mutator transaction binding the contract method 0x967448cd.
//
// Solidity: function AccountTransfer(uint256 sender, uint256 receiver, uint256 amount) returns()
func (_Foodchain *FoodchainTransactorSession) AccountTransfer(sender *big.Int, receiver *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.Contract.AccountTransfer(&_Foodchain.TransactOpts, sender, receiver, amount)
}

// AccountWithdraw is a paid mutator transaction binding the contract method 0x243fff8f.
//
// Solidity: function AccountWithdraw(uint256 id, uint256 amount) returns()
func (_Foodchain *FoodchainTransactor) AccountWithdraw(opts *bind.TransactOpts, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.contract.Transact(opts, "AccountWithdraw", id, amount)
}

// AccountWithdraw is a paid mutator transaction binding the contract method 0x243fff8f.
//
// Solidity: function AccountWithdraw(uint256 id, uint256 amount) returns()
func (_Foodchain *FoodchainSession) AccountWithdraw(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.Contract.AccountWithdraw(&_Foodchain.TransactOpts, id, amount)
}

// AccountWithdraw is a paid mutator transaction binding the contract method 0x243fff8f.
//
// Solidity: function AccountWithdraw(uint256 id, uint256 amount) returns()
func (_Foodchain *FoodchainTransactorSession) AccountWithdraw(id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Foodchain.Contract.AccountWithdraw(&_Foodchain.TransactOpts, id, amount)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x20290b3f.
//
// Solidity: function CreateAccount(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime) returns(uint256)
func (_Foodchain *FoodchainTransactor) CreateAccount(opts *bind.TransactOpts, id *big.Int, name string, accountAddress string, phoneNumber string, balance *big.Int, status *big.Int, createdTime string) (*types.Transaction, error) {
	return _Foodchain.contract.Transact(opts, "CreateAccount", id, name, accountAddress, phoneNumber, balance, status, createdTime)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x20290b3f.
//
// Solidity: function CreateAccount(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime) returns(uint256)
func (_Foodchain *FoodchainSession) CreateAccount(id *big.Int, name string, accountAddress string, phoneNumber string, balance *big.Int, status *big.Int, createdTime string) (*types.Transaction, error) {
	return _Foodchain.Contract.CreateAccount(&_Foodchain.TransactOpts, id, name, accountAddress, phoneNumber, balance, status, createdTime)
}

// CreateAccount is a paid mutator transaction binding the contract method 0x20290b3f.
//
// Solidity: function CreateAccount(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime) returns(uint256)
func (_Foodchain *FoodchainTransactorSession) CreateAccount(id *big.Int, name string, accountAddress string, phoneNumber string, balance *big.Int, status *big.Int, createdTime string) (*types.Transaction, error) {
	return _Foodchain.Contract.CreateAccount(&_Foodchain.TransactOpts, id, name, accountAddress, phoneNumber, balance, status, createdTime)
}

// DeleteAccount is a paid mutator transaction binding the contract method 0xbe405fa6.
//
// Solidity: function DeleteAccount(uint256 id) returns()
func (_Foodchain *FoodchainTransactor) DeleteAccount(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _Foodchain.contract.Transact(opts, "DeleteAccount", id)
}

// DeleteAccount is a paid mutator transaction binding the contract method 0xbe405fa6.
//
// Solidity: function DeleteAccount(uint256 id) returns()
func (_Foodchain *FoodchainSession) DeleteAccount(id *big.Int) (*types.Transaction, error) {
	return _Foodchain.Contract.DeleteAccount(&_Foodchain.TransactOpts, id)
}

// DeleteAccount is a paid mutator transaction binding the contract method 0xbe405fa6.
//
// Solidity: function DeleteAccount(uint256 id) returns()
func (_Foodchain *FoodchainTransactorSession) DeleteAccount(id *big.Int) (*types.Transaction, error) {
	return _Foodchain.Contract.DeleteAccount(&_Foodchain.TransactOpts, id)
}

// UpdateAccount is a paid mutator transaction binding the contract method 0xd695136d.
//
// Solidity: function UpdateAccount(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime) returns(uint256, string, string, string, uint256, uint256, string)
func (_Foodchain *FoodchainTransactor) UpdateAccount(opts *bind.TransactOpts, id *big.Int, name string, accountAddress string, phoneNumber string, balance *big.Int, status *big.Int, createdTime string) (*types.Transaction, error) {
	return _Foodchain.contract.Transact(opts, "UpdateAccount", id, name, accountAddress, phoneNumber, balance, status, createdTime)
}

// UpdateAccount is a paid mutator transaction binding the contract method 0xd695136d.
//
// Solidity: function UpdateAccount(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime) returns(uint256, string, string, string, uint256, uint256, string)
func (_Foodchain *FoodchainSession) UpdateAccount(id *big.Int, name string, accountAddress string, phoneNumber string, balance *big.Int, status *big.Int, createdTime string) (*types.Transaction, error) {
	return _Foodchain.Contract.UpdateAccount(&_Foodchain.TransactOpts, id, name, accountAddress, phoneNumber, balance, status, createdTime)
}

// UpdateAccount is a paid mutator transaction binding the contract method 0xd695136d.
//
// Solidity: function UpdateAccount(uint256 id, string name, string accountAddress, string phoneNumber, uint256 balance, uint256 status, string createdTime) returns(uint256, string, string, string, uint256, uint256, string)
func (_Foodchain *FoodchainTransactorSession) UpdateAccount(id *big.Int, name string, accountAddress string, phoneNumber string, balance *big.Int, status *big.Int, createdTime string) (*types.Transaction, error) {
	return _Foodchain.Contract.UpdateAccount(&_Foodchain.TransactOpts, id, name, accountAddress, phoneNumber, balance, status, createdTime)
}

// FoodchainCreateIterator is returned from FilterCreate and is used to iterate over the raw logs and unpacked data for Create events raised by the Foodchain contract.
type FoodchainCreateIterator struct {
	Event *FoodchainCreate // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoodchainCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodchainCreate)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoodchainCreate)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoodchainCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodchainCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodchainCreate represents a Create event raised by the Foodchain contract.
type FoodchainCreate struct {
	Id          *big.Int
	CreatedTime string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCreate is a free log retrieval operation binding the contract event 0xbba37318c149b1db7d49c422b47f2a61ba899714cb9c7377168ec21d7edfaae1.
//
// Solidity: event Create(uint256 id, string createdTime)
func (_Foodchain *FoodchainFilterer) FilterCreate(opts *bind.FilterOpts) (*FoodchainCreateIterator, error) {

	logs, sub, err := _Foodchain.contract.FilterLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return &FoodchainCreateIterator{contract: _Foodchain.contract, event: "Create", logs: logs, sub: sub}, nil
}

// WatchCreate is a free log subscription operation binding the contract event 0xbba37318c149b1db7d49c422b47f2a61ba899714cb9c7377168ec21d7edfaae1.
//
// Solidity: event Create(uint256 id, string createdTime)
func (_Foodchain *FoodchainFilterer) WatchCreate(opts *bind.WatchOpts, sink chan<- *FoodchainCreate) (event.Subscription, error) {

	logs, sub, err := _Foodchain.contract.WatchLogs(opts, "Create")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodchainCreate)
				if err := _Foodchain.contract.UnpackLog(event, "Create", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCreate is a log parse operation binding the contract event 0xbba37318c149b1db7d49c422b47f2a61ba899714cb9c7377168ec21d7edfaae1.
//
// Solidity: event Create(uint256 id, string createdTime)
func (_Foodchain *FoodchainFilterer) ParseCreate(log types.Log) (*FoodchainCreate, error) {
	event := new(FoodchainCreate)
	if err := _Foodchain.contract.UnpackLog(event, "Create", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoodchainDeleteIterator is returned from FilterDelete and is used to iterate over the raw logs and unpacked data for Delete events raised by the Foodchain contract.
type FoodchainDeleteIterator struct {
	Event *FoodchainDelete // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoodchainDeleteIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodchainDelete)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoodchainDelete)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoodchainDeleteIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodchainDeleteIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodchainDelete represents a Delete event raised by the Foodchain contract.
type FoodchainDelete struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterDelete is a free log retrieval operation binding the contract event 0x7d815d575a49d6d9a069482c25b0fdb39914ccf1bba8aaea4cfe20322f8f6877.
//
// Solidity: event Delete(uint256 id)
func (_Foodchain *FoodchainFilterer) FilterDelete(opts *bind.FilterOpts) (*FoodchainDeleteIterator, error) {

	logs, sub, err := _Foodchain.contract.FilterLogs(opts, "Delete")
	if err != nil {
		return nil, err
	}
	return &FoodchainDeleteIterator{contract: _Foodchain.contract, event: "Delete", logs: logs, sub: sub}, nil
}

// WatchDelete is a free log subscription operation binding the contract event 0x7d815d575a49d6d9a069482c25b0fdb39914ccf1bba8aaea4cfe20322f8f6877.
//
// Solidity: event Delete(uint256 id)
func (_Foodchain *FoodchainFilterer) WatchDelete(opts *bind.WatchOpts, sink chan<- *FoodchainDelete) (event.Subscription, error) {

	logs, sub, err := _Foodchain.contract.WatchLogs(opts, "Delete")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodchainDelete)
				if err := _Foodchain.contract.UnpackLog(event, "Delete", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDelete is a log parse operation binding the contract event 0x7d815d575a49d6d9a069482c25b0fdb39914ccf1bba8aaea4cfe20322f8f6877.
//
// Solidity: event Delete(uint256 id)
func (_Foodchain *FoodchainFilterer) ParseDelete(log types.Log) (*FoodchainDelete, error) {
	event := new(FoodchainDelete)
	if err := _Foodchain.contract.UnpackLog(event, "Delete", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoodchainDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the Foodchain contract.
type FoodchainDepositIterator struct {
	Event *FoodchainDeposit // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoodchainDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodchainDeposit)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoodchainDeposit)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoodchainDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodchainDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodchainDeposit represents a Deposit event raised by the Foodchain contract.
type FoodchainDeposit struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xa3af609bf46297028ce551832669030f9effef2b02606d02cbbcc40fe6b47c55.
//
// Solidity: event Deposit(uint256 id, uint256 amount)
func (_Foodchain *FoodchainFilterer) FilterDeposit(opts *bind.FilterOpts) (*FoodchainDepositIterator, error) {

	logs, sub, err := _Foodchain.contract.FilterLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return &FoodchainDepositIterator{contract: _Foodchain.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xa3af609bf46297028ce551832669030f9effef2b02606d02cbbcc40fe6b47c55.
//
// Solidity: event Deposit(uint256 id, uint256 amount)
func (_Foodchain *FoodchainFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *FoodchainDeposit) (event.Subscription, error) {

	logs, sub, err := _Foodchain.contract.WatchLogs(opts, "Deposit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodchainDeposit)
				if err := _Foodchain.contract.UnpackLog(event, "Deposit", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseDeposit is a log parse operation binding the contract event 0xa3af609bf46297028ce551832669030f9effef2b02606d02cbbcc40fe6b47c55.
//
// Solidity: event Deposit(uint256 id, uint256 amount)
func (_Foodchain *FoodchainFilterer) ParseDeposit(log types.Log) (*FoodchainDeposit, error) {
	event := new(FoodchainDeposit)
	if err := _Foodchain.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoodchainTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Foodchain contract.
type FoodchainTransferIterator struct {
	Event *FoodchainTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoodchainTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodchainTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoodchainTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoodchainTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodchainTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodchainTransfer represents a Transfer event raised by the Foodchain contract.
type FoodchainTransfer struct {
	From   *big.Int
	To     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xaf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6.
//
// Solidity: event Transfer(uint256 from, uint256 to, uint256 amount)
func (_Foodchain *FoodchainFilterer) FilterTransfer(opts *bind.FilterOpts) (*FoodchainTransferIterator, error) {

	logs, sub, err := _Foodchain.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &FoodchainTransferIterator{contract: _Foodchain.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xaf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6.
//
// Solidity: event Transfer(uint256 from, uint256 to, uint256 amount)
func (_Foodchain *FoodchainFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *FoodchainTransfer) (event.Subscription, error) {

	logs, sub, err := _Foodchain.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodchainTransfer)
				if err := _Foodchain.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xaf6151f5085accf2d57e1e7bf7601d3b3982e0de7e9a90f032f8554de9c104f6.
//
// Solidity: event Transfer(uint256 from, uint256 to, uint256 amount)
func (_Foodchain *FoodchainFilterer) ParseTransfer(log types.Log) (*FoodchainTransfer, error) {
	event := new(FoodchainTransfer)
	if err := _Foodchain.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FoodchainWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the Foodchain contract.
type FoodchainWithdrawIterator struct {
	Event *FoodchainWithdraw // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FoodchainWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FoodchainWithdraw)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FoodchainWithdraw)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FoodchainWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FoodchainWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FoodchainWithdraw represents a Withdraw event raised by the Foodchain contract.
type FoodchainWithdraw struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x56ca301a9219608c91e7bcee90e083c19671d2cdcc96752c7af291cee5f9c8c8.
//
// Solidity: event Withdraw(uint256 id, uint256 amount)
func (_Foodchain *FoodchainFilterer) FilterWithdraw(opts *bind.FilterOpts) (*FoodchainWithdrawIterator, error) {

	logs, sub, err := _Foodchain.contract.FilterLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return &FoodchainWithdrawIterator{contract: _Foodchain.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x56ca301a9219608c91e7bcee90e083c19671d2cdcc96752c7af291cee5f9c8c8.
//
// Solidity: event Withdraw(uint256 id, uint256 amount)
func (_Foodchain *FoodchainFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *FoodchainWithdraw) (event.Subscription, error) {

	logs, sub, err := _Foodchain.contract.WatchLogs(opts, "Withdraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FoodchainWithdraw)
				if err := _Foodchain.contract.UnpackLog(event, "Withdraw", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseWithdraw is a log parse operation binding the contract event 0x56ca301a9219608c91e7bcee90e083c19671d2cdcc96752c7af291cee5f9c8c8.
//
// Solidity: event Withdraw(uint256 id, uint256 amount)
func (_Foodchain *FoodchainFilterer) ParseWithdraw(log types.Log) (*FoodchainWithdraw, error) {
	event := new(FoodchainWithdraw)
	if err := _Foodchain.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
