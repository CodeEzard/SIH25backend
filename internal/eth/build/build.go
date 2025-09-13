// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package build

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
	_ = abi.ConvertType
)

// BuildMetaData contains all meta data concerning the Build contract.
var BuildMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721IncorrectOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721InsufficientApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOperator\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"ERC721InvalidOwner\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC721InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC721InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ERC721NonexistentToken\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_toTokenId\",\"type\":\"uint256\"}],\"name\":\"BatchMetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MetadataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"tokenURI\",\"type\":\"string\"}],\"name\":\"mintDoc\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_org\",\"type\":\"address\"}],\"name\":\"newOrg\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"_tokenIds\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"allOrgs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"currentTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"adr\",\"type\":\"address\"}],\"name\":\"isVerfiedOrg\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"Orgs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"verifiedOrgs\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561000f575f5ffd5b506040518060400160405280600881526020017f646f63756d656e740000000000000000000000000000000000000000000000008152506040518060400160405280600481526020017f5643524400000000000000000000000000000000000000000000000000000000815250815f908161008a919061031f565b50806001908161009a919061031f565b5050503360075f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506103ee565b5f81519050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061015d57607f821691505b6020821081036101705761016f610119565b5b50919050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026101d27fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610197565b6101dc8683610197565b95508019841693508086168417925050509392505050565b5f819050919050565b5f819050919050565b5f61022061021b610216846101f4565b6101fd565b6101f4565b9050919050565b5f819050919050565b61023983610206565b61024d61024582610227565b8484546101a3565b825550505050565b5f5f905090565b610264610255565b61026f818484610230565b505050565b5b81811015610292576102875f8261025c565b600181019050610275565b5050565b601f8211156102d7576102a881610176565b6102b184610188565b810160208510156102c0578190505b6102d46102cc85610188565b830182610274565b50505b505050565b5f82821c905092915050565b5f6102f75f19846008026102dc565b1980831691505092915050565b5f61030f83836102e8565b9150826002028217905092915050565b610328826100e2565b67ffffffffffffffff811115610341576103406100ec565b5b61034b8254610146565b610356828285610296565b5f60209050601f831160018114610387575f8415610375578287015190505b61037f8582610304565b8655506103e6565b601f19841661039586610176565b5f5b828110156103bc57848901518255600182019150602085019450602081019050610397565b868310156103d957848901516103d5601f8916826102e8565b8355505b6001600288020188555050505b505050505050565b61297a806103fb5f395ff3fe608060405234801561000f575f5ffd5b506004361061013f575f3560e01c80636352211e116100b657806395d89b411161007a57806395d89b41146103ab578063a22cb465146103c9578063aa46a400146103e5578063b88d4fde14610403578063c87b56dd1461041f578063e985e9c51461044f5761013f565b80636352211e146102cd5780636fae464c146102fd57806370a082311461032d57806376c77fff1461035d5780638da5cb5b1461038d5761013f565b806323b872dd1161010857806323b872dd146101fb5780632712ea3f1461021757806342842e0e1461024757806345050d41146102635780634ec6e937146102815780635a696f371461029d5761013f565b80629a9b7b1461014357806301ffc9a71461016157806306fdde0314610191578063081812fc146101af578063095ea7b3146101df575b5f5ffd5b61014b61047f565b6040516101589190611c1d565b60405180910390f35b61017b60048036038101906101769190611c9c565b610488565b6040516101889190611ce1565b60405180910390f35b6101996104e8565b6040516101a69190611d6a565b60405180910390f35b6101c960048036038101906101c49190611db4565b610577565b6040516101d69190611e1e565b60405180910390f35b6101f960048036038101906101f49190611e61565b610592565b005b61021560048036038101906102109190611e9f565b6105a8565b005b610231600480360381019061022c9190611eef565b6106a7565b60405161023e9190611ce1565b60405180910390f35b610261600480360381019061025c9190611e9f565b6106fe565b005b61026b61071d565b6040516102789190611fd1565b60405180910390f35b61029b60048036038101906102969190611eef565b6107a8565b005b6102b760048036038101906102b29190611db4565b6108ef565b6040516102c49190611e1e565b60405180910390f35b6102e760048036038101906102e29190611db4565b61092a565b6040516102f49190611e1e565b60405180910390f35b61031760048036038101906103129190611eef565b61093b565b6040516103249190611ce1565b60405180910390f35b61034760048036038101906103429190611eef565b610958565b6040516103549190611c1d565b60405180910390f35b6103776004803603810190610372919061211d565b610a0e565b6040516103849190611c1d565b60405180910390f35b610395610ad5565b6040516103a29190611e1e565b60405180910390f35b6103b3610afa565b6040516103c09190611d6a565b60405180910390f35b6103e360048036038101906103de91906121a1565b610b8a565b005b6103ed610ba0565b6040516103fa9190611c1d565b60405180910390f35b61041d6004803603810190610418919061227d565b610ba6565b005b61043960048036038101906104349190611db4565b610bcb565b6040516104469190611d6a565b60405180910390f35b610469600480360381019061046491906122fd565b610cd6565b6040516104769190611ce1565b60405180910390f35b5f600854905090565b5f634906490660e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614806104e157506104e082610d64565b5b9050919050565b60605f80546104f690612368565b80601f016020809104026020016040519081016040528092919081815260200182805461052290612368565b801561056d5780601f106105445761010080835404028352916020019161056d565b820191905f5260205f20905b81548152906001019060200180831161055057829003601f168201915b5050505050905090565b5f61058182610e45565b5061058b82610ecb565b9050919050565b6105a4828261059f610f04565b610f0b565b5050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603610618575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161060f9190611e1e565b60405180910390fd5b5f61062b8383610626610f04565b610f1d565b90508373ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146106a1578382826040517f64283d7b00000000000000000000000000000000000000000000000000000000815260040161069893929190612398565b60405180910390fd5b50505050565b5f5f60095f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905080915050919050565b61071883838360405180602001604052805f815250610ba6565b505050565b6060600a80548060200260200160405190810160405280929190818152602001828054801561079e57602002820191905f5260205f20905b815f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610755575b5050505050905090565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610837576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082e9061243d565b60405180910390fd5b600160095f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff021916908315150217905550600a81908060018154018082558091505060019003905f5260205f20015f9091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600a81815481106108fe575f80fd5b905f5260205f20015f915054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b5f61093482610e45565b9050919050565b6009602052805f5260405f205f915054906101000a900460ff1681565b5f5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036109c9575f6040517f89c62b640000000000000000000000000000000000000000000000000000000081526004016109c09190611e1e565b60405180910390fd5b60035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f20549050919050565b5f60095f3373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16610a98576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610a8f906124cb565b60405180910390fd5b600160085f828254610aaa9190612516565b925050819055505f6008549050610ac18482611128565b610acb818461121b565b8091505092915050565b60075f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b606060018054610b0990612368565b80601f0160208091040260200160405190810160405280929190818152602001828054610b3590612368565b8015610b805780601f10610b5757610100808354040283529160200191610b80565b820191905f5260205f20905b815481529060010190602001808311610b6357829003601f168201915b5050505050905090565b610b9c610b95610f04565b8383611275565b5050565b60085481565b610bb18484846105a8565b610bc5610bbc610f04565b858585856113de565b50505050565b6060610bd682610e45565b505f60065f8481526020019081526020015f208054610bf490612368565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2090612368565b8015610c6b5780601f10610c4257610100808354040283529160200191610c6b565b820191905f5260205f20905b815481529060010190602001808311610c4e57829003601f168201915b505050505090505f610c7b61158a565b90505f815103610c8f578192505050610cd1565b5f82511115610cc3578082604051602001610cab929190612583565b60405160208183030381529060405292505050610cd1565b610ccc846115a0565b925050505b919050565b5f60055f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f9054906101000a900460ff16905092915050565b5f7f80ac58cd000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161480610e2e57507f5b5e139f000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916145b80610e3e5750610e3d82611606565b5b9050919050565b5f5f610e508361166f565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1603610ec257826040517f7e273289000000000000000000000000000000000000000000000000000000008152600401610eb99190611c1d565b60405180910390fd5b80915050919050565b5f60045f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b5f33905090565b610f1883838360016116a8565b505050565b5f5f610f288461166f565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614610f6957610f68818486611867565b5b5f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614610ff457610fa85f855f5f6116a8565b600160035f8373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825403925050819055505b5f73ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff161461107357600160035f8773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f82825401925050819055505b8460025f8681526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4809150509392505050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1603611198575f6040517f64a0ae9200000000000000000000000000000000000000000000000000000000815260040161118f9190611e1e565b60405180910390fd5b5f6111a483835f610f1d565b90505f73ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614611216575f6040517f73c6ac6e00000000000000000000000000000000000000000000000000000000815260040161120d9190611e1e565b60405180910390fd5b505050565b8060065f8481526020019081526020015f2090816112399190612746565b507ff8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7826040516112699190611c1d565b60405180910390a15050565b5f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16036112e557816040517f5b08ba180000000000000000000000000000000000000000000000000000000081526004016112dc9190611e1e565b60405180910390fd5b8060055f8573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f8473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020015f205f6101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516113d19190611ce1565b60405180910390a3505050565b5f8373ffffffffffffffffffffffffffffffffffffffff163b1115611583578273ffffffffffffffffffffffffffffffffffffffff1663150b7a02868685856040518563ffffffff1660e01b815260040161143c9493929190612867565b6020604051808303815f875af192505050801561147757506040513d601f19601f8201168201806040525081019061147491906128c5565b60015b6114f8573d805f81146114a5576040519150601f19603f3d011682016040523d82523d5f602084013e6114aa565b606091505b505f8151036114f057836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016114e79190611e1e565b60405180910390fd5b805181602001fd5b63150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161461158157836040517f64a0ae920000000000000000000000000000000000000000000000000000000081526004016115789190611e1e565b60405180910390fd5b505b5050505050565b606060405180602001604052805f815250905090565b60606115ab82610e45565b505f6115b561158a565b90505f8151116115d35760405180602001604052805f8152506115fe565b806115dd8461192a565b6040516020016115ee929190612583565b6040516020818303038152906040525b915050919050565b5f7f01ffc9a7000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916827bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916149050919050565b5f60025f8381526020019081526020015f205f9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050919050565b80806116e057505f73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614155b15611812575f6116ef84610e45565b90505f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415801561175957508273ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614155b801561176c575061176a8184610cd6565b155b156117ae57826040517fa9fbf51f0000000000000000000000000000000000000000000000000000000081526004016117a59190611e1e565b60405180910390fd5b811561181057838573ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b505b8360045f8581526020019081526020015f205f6101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050505050565b6118728383836119f4565b611925575f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16036118e657806040517f7e2732890000000000000000000000000000000000000000000000000000000081526004016118dd9190611c1d565b60405180910390fd5b81816040517f177e802f00000000000000000000000000000000000000000000000000000000815260040161191c9291906128f0565b60405180910390fd5b505050565b60605f600161193884611ab4565b0190505f8167ffffffffffffffff81111561195657611955611ff9565b5b6040519080825280601f01601f1916602001820160405280156119885781602001600182028036833780820191505090505b5090505f82602001820190505b6001156119e9578080600190039150507f3031323334353637383961626364656600000000000000000000000000000000600a86061a8153600a85816119de576119dd612917565b5b0494505f8503611995575b819350505050919050565b5f5f73ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614158015611aab57508273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff161480611a6c5750611a6b8484610cd6565b5b80611aaa57508273ffffffffffffffffffffffffffffffffffffffff16611a9283610ecb565b73ffffffffffffffffffffffffffffffffffffffff16145b5b90509392505050565b5f5f5f90507a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008310611b10577a184f03e93ff9f4daa797ed6e38ed64bf6a1f0100000000000000008381611b0657611b05612917565b5b0492506040810190505b6d04ee2d6d415b85acef81000000008310611b4d576d04ee2d6d415b85acef81000000008381611b4357611b42612917565b5b0492506020810190505b662386f26fc100008310611b7c57662386f26fc100008381611b7257611b71612917565b5b0492506010810190505b6305f5e1008310611ba5576305f5e1008381611b9b57611b9a612917565b5b0492506008810190505b6127108310611bca576127108381611bc057611bbf612917565b5b0492506004810190505b60648310611bed5760648381611be357611be2612917565b5b0492506002810190505b600a8310611bfc576001810190505b80915050919050565b5f819050919050565b611c1781611c05565b82525050565b5f602082019050611c305f830184611c0e565b92915050565b5f604051905090565b5f5ffd5b5f5ffd5b5f7fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b611c7b81611c47565b8114611c85575f5ffd5b50565b5f81359050611c9681611c72565b92915050565b5f60208284031215611cb157611cb0611c3f565b5b5f611cbe84828501611c88565b91505092915050565b5f8115159050919050565b611cdb81611cc7565b82525050565b5f602082019050611cf45f830184611cd2565b92915050565b5f81519050919050565b5f82825260208201905092915050565b8281835e5f83830152505050565b5f601f19601f8301169050919050565b5f611d3c82611cfa565b611d468185611d04565b9350611d56818560208601611d14565b611d5f81611d22565b840191505092915050565b5f6020820190508181035f830152611d828184611d32565b905092915050565b611d9381611c05565b8114611d9d575f5ffd5b50565b5f81359050611dae81611d8a565b92915050565b5f60208284031215611dc957611dc8611c3f565b5b5f611dd684828501611da0565b91505092915050565b5f73ffffffffffffffffffffffffffffffffffffffff82169050919050565b5f611e0882611ddf565b9050919050565b611e1881611dfe565b82525050565b5f602082019050611e315f830184611e0f565b92915050565b611e4081611dfe565b8114611e4a575f5ffd5b50565b5f81359050611e5b81611e37565b92915050565b5f5f60408385031215611e7757611e76611c3f565b5b5f611e8485828601611e4d565b9250506020611e9585828601611da0565b9150509250929050565b5f5f5f60608486031215611eb657611eb5611c3f565b5b5f611ec386828701611e4d565b9350506020611ed486828701611e4d565b9250506040611ee586828701611da0565b9150509250925092565b5f60208284031215611f0457611f03611c3f565b5b5f611f1184828501611e4d565b91505092915050565b5f81519050919050565b5f82825260208201905092915050565b5f819050602082019050919050565b611f4c81611dfe565b82525050565b5f611f5d8383611f43565b60208301905092915050565b5f602082019050919050565b5f611f7f82611f1a565b611f898185611f24565b9350611f9483611f34565b805f5b83811015611fc4578151611fab8882611f52565b9750611fb683611f69565b925050600181019050611f97565b5085935050505092915050565b5f6020820190508181035f830152611fe98184611f75565b905092915050565b5f5ffd5b5f5ffd5b7f4e487b71000000000000000000000000000000000000000000000000000000005f52604160045260245ffd5b61202f82611d22565b810181811067ffffffffffffffff8211171561204e5761204d611ff9565b5b80604052505050565b5f612060611c36565b905061206c8282612026565b919050565b5f67ffffffffffffffff82111561208b5761208a611ff9565b5b61209482611d22565b9050602081019050919050565b828183375f83830152505050565b5f6120c16120bc84612071565b612057565b9050828152602081018484840111156120dd576120dc611ff5565b5b6120e88482856120a1565b509392505050565b5f82601f83011261210457612103611ff1565b5b81356121148482602086016120af565b91505092915050565b5f5f6040838503121561213357612132611c3f565b5b5f61214085828601611e4d565b925050602083013567ffffffffffffffff81111561216157612160611c43565b5b61216d858286016120f0565b9150509250929050565b61218081611cc7565b811461218a575f5ffd5b50565b5f8135905061219b81612177565b92915050565b5f5f604083850312156121b7576121b6611c3f565b5b5f6121c485828601611e4d565b92505060206121d58582860161218d565b9150509250929050565b5f67ffffffffffffffff8211156121f9576121f8611ff9565b5b61220282611d22565b9050602081019050919050565b5f61222161221c846121df565b612057565b90508281526020810184848401111561223d5761223c611ff5565b5b6122488482856120a1565b509392505050565b5f82601f83011261226457612263611ff1565b5b813561227484826020860161220f565b91505092915050565b5f5f5f5f6080858703121561229557612294611c3f565b5b5f6122a287828801611e4d565b94505060206122b387828801611e4d565b93505060406122c487828801611da0565b925050606085013567ffffffffffffffff8111156122e5576122e4611c43565b5b6122f187828801612250565b91505092959194509250565b5f5f6040838503121561231357612312611c3f565b5b5f61232085828601611e4d565b925050602061233185828601611e4d565b9150509250929050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52602260045260245ffd5b5f600282049050600182168061237f57607f821691505b6020821081036123925761239161233b565b5b50919050565b5f6060820190506123ab5f830186611e0f565b6123b86020830185611c0e565b6123c56040830184611e0f565b949350505050565b7f4f6e6c7920746865206f776e65722063616e2070757368206e6577206f7267735f8201527f2e00000000000000000000000000000000000000000000000000000000000000602082015250565b5f612427602183611d04565b9150612432826123cd565b604082019050919050565b5f6020820190508181035f8301526124548161241b565b9050919050565b7f556e6976657273697479206973206e6f742076657269666965642f6c697374655f8201527f6400000000000000000000000000000000000000000000000000000000000000602082015250565b5f6124b5602183611d04565b91506124c08261245b565b604082019050919050565b5f6020820190508181035f8301526124e2816124a9565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61252082611c05565b915061252b83611c05565b9250828201905080821115612543576125426124e9565b5b92915050565b5f81905092915050565b5f61255d82611cfa565b6125678185612549565b9350612577818560208601611d14565b80840191505092915050565b5f61258e8285612553565b915061259a8284612553565b91508190509392505050565b5f819050815f5260205f209050919050565b5f6020601f8301049050919050565b5f82821b905092915050565b5f600883026126027fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff826125c7565b61260c86836125c7565b95508019841693508086168417925050509392505050565b5f819050919050565b5f61264761264261263d84611c05565b612624565b611c05565b9050919050565b5f819050919050565b6126608361262d565b61267461266c8261264e565b8484546125d3565b825550505050565b5f5f905090565b61268b61267c565b612696818484612657565b505050565b5b818110156126b9576126ae5f82612683565b60018101905061269c565b5050565b601f8211156126fe576126cf816125a6565b6126d8846125b8565b810160208510156126e7578190505b6126fb6126f3856125b8565b83018261269b565b50505b505050565b5f82821c905092915050565b5f61271e5f1984600802612703565b1980831691505092915050565b5f612736838361270f565b9150826002028217905092915050565b61274f82611cfa565b67ffffffffffffffff81111561276857612767611ff9565b5b6127728254612368565b61277d8282856126bd565b5f60209050601f8311600181146127ae575f841561279c578287015190505b6127a6858261272b565b86555061280d565b601f1984166127bc866125a6565b5f5b828110156127e3578489015182556001820191506020850194506020810190506127be565b8683101561280057848901516127fc601f89168261270f565b8355505b6001600288020188555050505b505050505050565b5f81519050919050565b5f82825260208201905092915050565b5f61283982612815565b612843818561281f565b9350612853818560208601611d14565b61285c81611d22565b840191505092915050565b5f60808201905061287a5f830187611e0f565b6128876020830186611e0f565b6128946040830185611c0e565b81810360608301526128a6818461282f565b905095945050505050565b5f815190506128bf81611c72565b92915050565b5f602082840312156128da576128d9611c3f565b5b5f6128e7848285016128b1565b91505092915050565b5f6040820190506129035f830185611e0f565b6129106020830184611c0e565b9392505050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601260045260245ffdfea2646970667358221220ab5f6c34d189af958542d085ef63c5042f060d9b6db29b3891f3095a5bd1c3e464736f6c634300081e0033",
}

// BuildABI is the input ABI used to generate the binding from.
// Deprecated: Use BuildMetaData.ABI instead.
var BuildABI = BuildMetaData.ABI

// BuildBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BuildMetaData.Bin instead.
var BuildBin = BuildMetaData.Bin

// DeployBuild deploys a new Ethereum contract, binding an instance of Build to it.
func DeployBuild(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Build, error) {
	parsed, err := BuildMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BuildBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Build{BuildCaller: BuildCaller{contract: contract}, BuildTransactor: BuildTransactor{contract: contract}, BuildFilterer: BuildFilterer{contract: contract}}, nil
}

// Build is an auto generated Go binding around an Ethereum contract.
type Build struct {
	BuildCaller     // Read-only binding to the contract
	BuildTransactor // Write-only binding to the contract
	BuildFilterer   // Log filterer for contract events
}

// BuildCaller is an auto generated read-only Go binding around an Ethereum contract.
type BuildCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuildTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BuildTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuildFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BuildFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BuildSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BuildSession struct {
	Contract     *Build            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuildCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BuildCallerSession struct {
	Contract *BuildCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BuildTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BuildTransactorSession struct {
	Contract     *BuildTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BuildRaw is an auto generated low-level Go binding around an Ethereum contract.
type BuildRaw struct {
	Contract *Build // Generic contract binding to access the raw methods on
}

// BuildCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BuildCallerRaw struct {
	Contract *BuildCaller // Generic read-only contract binding to access the raw methods on
}

// BuildTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BuildTransactorRaw struct {
	Contract *BuildTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBuild creates a new instance of Build, bound to a specific deployed contract.
func NewBuild(address common.Address, backend bind.ContractBackend) (*Build, error) {
	contract, err := bindBuild(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Build{BuildCaller: BuildCaller{contract: contract}, BuildTransactor: BuildTransactor{contract: contract}, BuildFilterer: BuildFilterer{contract: contract}}, nil
}

// NewBuildCaller creates a new read-only instance of Build, bound to a specific deployed contract.
func NewBuildCaller(address common.Address, caller bind.ContractCaller) (*BuildCaller, error) {
	contract, err := bindBuild(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BuildCaller{contract: contract}, nil
}

// NewBuildTransactor creates a new write-only instance of Build, bound to a specific deployed contract.
func NewBuildTransactor(address common.Address, transactor bind.ContractTransactor) (*BuildTransactor, error) {
	contract, err := bindBuild(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BuildTransactor{contract: contract}, nil
}

// NewBuildFilterer creates a new log filterer instance of Build, bound to a specific deployed contract.
func NewBuildFilterer(address common.Address, filterer bind.ContractFilterer) (*BuildFilterer, error) {
	contract, err := bindBuild(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BuildFilterer{contract: contract}, nil
}

// bindBuild binds a generic wrapper to an already deployed contract.
func bindBuild(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BuildMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Build *BuildRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Build.Contract.BuildCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Build *BuildRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Build.Contract.BuildTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Build *BuildRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Build.Contract.BuildTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Build *BuildCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Build.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Build *BuildTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Build.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Build *BuildTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Build.Contract.contract.Transact(opts, method, params...)
}

// Orgs is a free data retrieval call binding the contract method 0x5a696f37.
//
// Solidity: function Orgs(uint256 ) view returns(address)
func (_Build *BuildCaller) Orgs(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "Orgs", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Orgs is a free data retrieval call binding the contract method 0x5a696f37.
//
// Solidity: function Orgs(uint256 ) view returns(address)
func (_Build *BuildSession) Orgs(arg0 *big.Int) (common.Address, error) {
	return _Build.Contract.Orgs(&_Build.CallOpts, arg0)
}

// Orgs is a free data retrieval call binding the contract method 0x5a696f37.
//
// Solidity: function Orgs(uint256 ) view returns(address)
func (_Build *BuildCallerSession) Orgs(arg0 *big.Int) (common.Address, error) {
	return _Build.Contract.Orgs(&_Build.CallOpts, arg0)
}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256)
func (_Build *BuildCaller) TokenIds(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "_tokenIds")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256)
func (_Build *BuildSession) TokenIds() (*big.Int, error) {
	return _Build.Contract.TokenIds(&_Build.CallOpts)
}

// TokenIds is a free data retrieval call binding the contract method 0xaa46a400.
//
// Solidity: function _tokenIds() view returns(uint256)
func (_Build *BuildCallerSession) TokenIds() (*big.Int, error) {
	return _Build.Contract.TokenIds(&_Build.CallOpts)
}

// AllOrgs is a free data retrieval call binding the contract method 0x45050d41.
//
// Solidity: function allOrgs() view returns(address[])
func (_Build *BuildCaller) AllOrgs(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "allOrgs")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// AllOrgs is a free data retrieval call binding the contract method 0x45050d41.
//
// Solidity: function allOrgs() view returns(address[])
func (_Build *BuildSession) AllOrgs() ([]common.Address, error) {
	return _Build.Contract.AllOrgs(&_Build.CallOpts)
}

// AllOrgs is a free data retrieval call binding the contract method 0x45050d41.
//
// Solidity: function allOrgs() view returns(address[])
func (_Build *BuildCallerSession) AllOrgs() ([]common.Address, error) {
	return _Build.Contract.AllOrgs(&_Build.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Build *BuildCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Build *BuildSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Build.Contract.BalanceOf(&_Build.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Build *BuildCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Build.Contract.BalanceOf(&_Build.CallOpts, owner)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Build *BuildCaller) CurrentTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "currentTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Build *BuildSession) CurrentTokenId() (*big.Int, error) {
	return _Build.Contract.CurrentTokenId(&_Build.CallOpts)
}

// CurrentTokenId is a free data retrieval call binding the contract method 0x009a9b7b.
//
// Solidity: function currentTokenId() view returns(uint256)
func (_Build *BuildCallerSession) CurrentTokenId() (*big.Int, error) {
	return _Build.Contract.CurrentTokenId(&_Build.CallOpts)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Build *BuildCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Build *BuildSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Build.Contract.GetApproved(&_Build.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Build *BuildCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Build.Contract.GetApproved(&_Build.CallOpts, tokenId)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Build *BuildCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Build *BuildSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Build.Contract.IsApprovedForAll(&_Build.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Build *BuildCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Build.Contract.IsApprovedForAll(&_Build.CallOpts, owner, operator)
}

// IsVerfiedOrg is a free data retrieval call binding the contract method 0x2712ea3f.
//
// Solidity: function isVerfiedOrg(address adr) view returns(bool)
func (_Build *BuildCaller) IsVerfiedOrg(opts *bind.CallOpts, adr common.Address) (bool, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "isVerfiedOrg", adr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsVerfiedOrg is a free data retrieval call binding the contract method 0x2712ea3f.
//
// Solidity: function isVerfiedOrg(address adr) view returns(bool)
func (_Build *BuildSession) IsVerfiedOrg(adr common.Address) (bool, error) {
	return _Build.Contract.IsVerfiedOrg(&_Build.CallOpts, adr)
}

// IsVerfiedOrg is a free data retrieval call binding the contract method 0x2712ea3f.
//
// Solidity: function isVerfiedOrg(address adr) view returns(bool)
func (_Build *BuildCallerSession) IsVerfiedOrg(adr common.Address) (bool, error) {
	return _Build.Contract.IsVerfiedOrg(&_Build.CallOpts, adr)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Build *BuildCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Build *BuildSession) Name() (string, error) {
	return _Build.Contract.Name(&_Build.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Build *BuildCallerSession) Name() (string, error) {
	return _Build.Contract.Name(&_Build.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Build *BuildCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Build *BuildSession) Owner() (common.Address, error) {
	return _Build.Contract.Owner(&_Build.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Build *BuildCallerSession) Owner() (common.Address, error) {
	return _Build.Contract.Owner(&_Build.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Build *BuildCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Build *BuildSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Build.Contract.OwnerOf(&_Build.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Build *BuildCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Build.Contract.OwnerOf(&_Build.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Build *BuildCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Build *BuildSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Build.Contract.SupportsInterface(&_Build.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Build *BuildCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Build.Contract.SupportsInterface(&_Build.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Build *BuildCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Build *BuildSession) Symbol() (string, error) {
	return _Build.Contract.Symbol(&_Build.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Build *BuildCallerSession) Symbol() (string, error) {
	return _Build.Contract.Symbol(&_Build.CallOpts)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Build *BuildCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Build *BuildSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Build.Contract.TokenURI(&_Build.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Build *BuildCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Build.Contract.TokenURI(&_Build.CallOpts, tokenId)
}

// VerifiedOrgs is a free data retrieval call binding the contract method 0x6fae464c.
//
// Solidity: function verifiedOrgs(address ) view returns(bool)
func (_Build *BuildCaller) VerifiedOrgs(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Build.contract.Call(opts, &out, "verifiedOrgs", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifiedOrgs is a free data retrieval call binding the contract method 0x6fae464c.
//
// Solidity: function verifiedOrgs(address ) view returns(bool)
func (_Build *BuildSession) VerifiedOrgs(arg0 common.Address) (bool, error) {
	return _Build.Contract.VerifiedOrgs(&_Build.CallOpts, arg0)
}

// VerifiedOrgs is a free data retrieval call binding the contract method 0x6fae464c.
//
// Solidity: function verifiedOrgs(address ) view returns(bool)
func (_Build *BuildCallerSession) VerifiedOrgs(arg0 common.Address) (bool, error) {
	return _Build.Contract.VerifiedOrgs(&_Build.CallOpts, arg0)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Build *BuildTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Build *BuildSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.Approve(&_Build.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Build *BuildTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.Approve(&_Build.TransactOpts, to, tokenId)
}

// MintDoc is a paid mutator transaction binding the contract method 0x76c77fff.
//
// Solidity: function mintDoc(address user, string tokenURI) returns(uint256)
func (_Build *BuildTransactor) MintDoc(opts *bind.TransactOpts, user common.Address, tokenURI string) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "mintDoc", user, tokenURI)
}

// MintDoc is a paid mutator transaction binding the contract method 0x76c77fff.
//
// Solidity: function mintDoc(address user, string tokenURI) returns(uint256)
func (_Build *BuildSession) MintDoc(user common.Address, tokenURI string) (*types.Transaction, error) {
	return _Build.Contract.MintDoc(&_Build.TransactOpts, user, tokenURI)
}

// MintDoc is a paid mutator transaction binding the contract method 0x76c77fff.
//
// Solidity: function mintDoc(address user, string tokenURI) returns(uint256)
func (_Build *BuildTransactorSession) MintDoc(user common.Address, tokenURI string) (*types.Transaction, error) {
	return _Build.Contract.MintDoc(&_Build.TransactOpts, user, tokenURI)
}

// NewOrg is a paid mutator transaction binding the contract method 0x4ec6e937.
//
// Solidity: function newOrg(address _org) returns()
func (_Build *BuildTransactor) NewOrg(opts *bind.TransactOpts, _org common.Address) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "newOrg", _org)
}

// NewOrg is a paid mutator transaction binding the contract method 0x4ec6e937.
//
// Solidity: function newOrg(address _org) returns()
func (_Build *BuildSession) NewOrg(_org common.Address) (*types.Transaction, error) {
	return _Build.Contract.NewOrg(&_Build.TransactOpts, _org)
}

// NewOrg is a paid mutator transaction binding the contract method 0x4ec6e937.
//
// Solidity: function newOrg(address _org) returns()
func (_Build *BuildTransactorSession) NewOrg(_org common.Address) (*types.Transaction, error) {
	return _Build.Contract.NewOrg(&_Build.TransactOpts, _org)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.SafeTransferFrom(&_Build.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.SafeTransferFrom(&_Build.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Build *BuildTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Build *BuildSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Build.Contract.SafeTransferFrom0(&_Build.TransactOpts, from, to, tokenId, data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes data) returns()
func (_Build *BuildTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _Build.Contract.SafeTransferFrom0(&_Build.TransactOpts, from, to, tokenId, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Build *BuildTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Build *BuildSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Build.Contract.SetApprovalForAll(&_Build.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Build *BuildTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Build.Contract.SetApprovalForAll(&_Build.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.TransferFrom(&_Build.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Build *BuildTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Build.Contract.TransferFrom(&_Build.TransactOpts, from, to, tokenId)
}

// BuildApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Build contract.
type BuildApprovalIterator struct {
	Event *BuildApproval // Event containing the contract specifics and raw log

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
func (it *BuildApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildApproval)
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
		it.Event = new(BuildApproval)
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
func (it *BuildApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildApproval represents a Approval event raised by the Build contract.
type BuildApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Build *BuildFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BuildApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Build.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BuildApprovalIterator{contract: _Build.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Build *BuildFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BuildApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Build.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildApproval)
				if err := _Build.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Build *BuildFilterer) ParseApproval(log types.Log) (*BuildApproval, error) {
	event := new(BuildApproval)
	if err := _Build.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuildApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Build contract.
type BuildApprovalForAllIterator struct {
	Event *BuildApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BuildApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildApprovalForAll)
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
		it.Event = new(BuildApprovalForAll)
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
func (it *BuildApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildApprovalForAll represents a ApprovalForAll event raised by the Build contract.
type BuildApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Build *BuildFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BuildApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Build.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BuildApprovalForAllIterator{contract: _Build.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Build *BuildFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BuildApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Build.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildApprovalForAll)
				if err := _Build.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Build *BuildFilterer) ParseApprovalForAll(log types.Log) (*BuildApprovalForAll, error) {
	event := new(BuildApprovalForAll)
	if err := _Build.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuildBatchMetadataUpdateIterator is returned from FilterBatchMetadataUpdate and is used to iterate over the raw logs and unpacked data for BatchMetadataUpdate events raised by the Build contract.
type BuildBatchMetadataUpdateIterator struct {
	Event *BuildBatchMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *BuildBatchMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildBatchMetadataUpdate)
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
		it.Event = new(BuildBatchMetadataUpdate)
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
func (it *BuildBatchMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildBatchMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildBatchMetadataUpdate represents a BatchMetadataUpdate event raised by the Build contract.
type BuildBatchMetadataUpdate struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBatchMetadataUpdate is a free log retrieval operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Build *BuildFilterer) FilterBatchMetadataUpdate(opts *bind.FilterOpts) (*BuildBatchMetadataUpdateIterator, error) {

	logs, sub, err := _Build.contract.FilterLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &BuildBatchMetadataUpdateIterator{contract: _Build.contract, event: "BatchMetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchBatchMetadataUpdate is a free log subscription operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Build *BuildFilterer) WatchBatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *BuildBatchMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Build.contract.WatchLogs(opts, "BatchMetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildBatchMetadataUpdate)
				if err := _Build.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
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

// ParseBatchMetadataUpdate is a log parse operation binding the contract event 0x6bd5c950a8d8df17f772f5af37cb3655737899cbf903264b9795592da439661c.
//
// Solidity: event BatchMetadataUpdate(uint256 _fromTokenId, uint256 _toTokenId)
func (_Build *BuildFilterer) ParseBatchMetadataUpdate(log types.Log) (*BuildBatchMetadataUpdate, error) {
	event := new(BuildBatchMetadataUpdate)
	if err := _Build.contract.UnpackLog(event, "BatchMetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuildMetadataUpdateIterator is returned from FilterMetadataUpdate and is used to iterate over the raw logs and unpacked data for MetadataUpdate events raised by the Build contract.
type BuildMetadataUpdateIterator struct {
	Event *BuildMetadataUpdate // Event containing the contract specifics and raw log

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
func (it *BuildMetadataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildMetadataUpdate)
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
		it.Event = new(BuildMetadataUpdate)
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
func (it *BuildMetadataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildMetadataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildMetadataUpdate represents a MetadataUpdate event raised by the Build contract.
type BuildMetadataUpdate struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMetadataUpdate is a free log retrieval operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Build *BuildFilterer) FilterMetadataUpdate(opts *bind.FilterOpts) (*BuildMetadataUpdateIterator, error) {

	logs, sub, err := _Build.contract.FilterLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return &BuildMetadataUpdateIterator{contract: _Build.contract, event: "MetadataUpdate", logs: logs, sub: sub}, nil
}

// WatchMetadataUpdate is a free log subscription operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Build *BuildFilterer) WatchMetadataUpdate(opts *bind.WatchOpts, sink chan<- *BuildMetadataUpdate) (event.Subscription, error) {

	logs, sub, err := _Build.contract.WatchLogs(opts, "MetadataUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildMetadataUpdate)
				if err := _Build.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
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

// ParseMetadataUpdate is a log parse operation binding the contract event 0xf8e1a15aba9398e019f0b49df1a4fde98ee17ae345cb5f6b5e2c27f5033e8ce7.
//
// Solidity: event MetadataUpdate(uint256 _tokenId)
func (_Build *BuildFilterer) ParseMetadataUpdate(log types.Log) (*BuildMetadataUpdate, error) {
	event := new(BuildMetadataUpdate)
	if err := _Build.contract.UnpackLog(event, "MetadataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BuildTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Build contract.
type BuildTransferIterator struct {
	Event *BuildTransfer // Event containing the contract specifics and raw log

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
func (it *BuildTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BuildTransfer)
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
		it.Event = new(BuildTransfer)
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
func (it *BuildTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BuildTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BuildTransfer represents a Transfer event raised by the Build contract.
type BuildTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Build *BuildFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BuildTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Build.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BuildTransferIterator{contract: _Build.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Build *BuildFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BuildTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Build.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BuildTransfer)
				if err := _Build.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Build *BuildFilterer) ParseTransfer(log types.Log) (*BuildTransfer, error) {
	event := new(BuildTransfer)
	if err := _Build.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
