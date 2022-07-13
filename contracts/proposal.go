// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// VoteMetaData contains all meta data concerning the Vote contract.
var VoteMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint8\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"submitter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"candidateServiceUrl\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proposalUrl\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"submitBlockNo\",\"type\":\"uint256\"}],\"name\":\"NewProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"result\",\"type\":\"bool\"}],\"name\":\"ProposalResult\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"VoteProposal\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockNo\",\"type\":\"uint256\"}],\"name\":\"WithdrawProposal\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"effectProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllAuthority\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"},{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllProposalId\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"getProposal\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"blockNo\",\"type\":\"uint256\"}],\"name\":\"getProposalId\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"serviceUrl\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"proposalType\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"proposalUrl\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"candidate\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"candidateServiceUrl\",\"type\":\"string\"}],\"name\":\"submitProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"voteProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"withdrawProposal\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5061259e806100206000396000f3fe608060405234801561001057600080fd5b50600436106100cf5760003560e01c8063807896d51161008c578063c7f758a811610066578063c7f758a81461018f578063e1f02ffa146101b4578063f2fde38b146101c7578063f62d1888146101da57600080fd5b8063807896d51461014b5780638da5cb5b1461015e5780638eee15881461017957600080fd5b806351ba9c54146100d45780635944d7af146100f25780636e9960c314610107578063715018a61461011d57806379717e88146101255780637c95ac2714610138575b600080fd5b6100dc6101ed565b6040516100e9919061205a565b60405180910390f35b610105610100366004612158565b610245565b005b61010f6107dd565b6040516100e9929190612232565b610105610889565b61010561013336600461225e565b6108ef565b6100dc61014636600461225e565b6111da565b61010561015936600461225e565b611297565b6033546040516001600160a01b0390911681526020016100e9565b6101816115f5565b6040516100e9929190612277565b6101a261019d36600461225e565b6117ec565b6040516100e996959493929190612318565b6101056101c236600461225e565b611964565b6101056101d536600461236b565b611c87565b6101056101e836600461238d565b611d52565b6060606c80548060200260200160405190810160405280929190818152602001828054801561023b57602002820191906000526020600020905b815481526020019060010190808311610227575b5050505050905090565b60ff84166001148061025a575060ff84166002145b80610268575060ff84166003145b6102b15760405162461bcd60e51b8152602060048201526015602482015274496e76616c69642050726f706f73616c207479706560581b60448201526064015b60405180910390fd5b6067546000908190815b8181101561036557336001600160a01b0316606782815481106102e0576102e06123c2565b60009182526020909120600290910201546001600160a01b03160361030457600193505b856001600160a01b031660678281548110610321576103216123c2565b60009182526020909120600290910201546001600160a01b03160361034557600192505b83801561034f5750835b610365578061035d816123ee565b9150506102bb565b50828061037c57506065546001600160a01b031633145b6103985760405162461bcd60e51b81526004016102a890612407565b60ff87166001036104095781156104045760405162461bcd60e51b815260206004820152602a60248201527f63616e64696461746520697320616c726561647920696e2074686520617574686044820152691bdc9a5d1e481b1a5cdd60b21b60648201526084016102a8565b610465565b816104655760405162461bcd60e51b815260206004820152602660248201527f63616e646964617465206973206e6f7420696e2074686520617574686f7269746044820152651e481b1a5cdd60d21b60648201526084016102a8565b606c5460005b8181101561052d57866001600160a01b0316606d6000606c8481548110610494576104946123c2565b600091825260208083209091015483528201929092526040019020600301546001600160a01b03160361051b5760405162461bcd60e51b815260206004820152602960248201527f63616e64696461746520697320616c726561647920696e206f6e65206f70656e604482015268081c1c9bdc1bdcd85b60ba1b60648201526084016102a8565b80610525816123ee565b91505061046b565b5060ff881660021480610543575060ff88166003145b156105fd5760005b818110156105fb57866001600160a01b0316606d6000606c8481548110610574576105746123c2565b600091825260208083209091015483528201929092526040019020600201546001600160a01b0316036105e95760405162461bcd60e51b815260206004820152601d60248201527f63616e64696461746520686173206f70656e2070726f706f73616c732e00000060448201526064016102a8565b806105f3816123ee565b91505061054b565b505b6040518060e001604052808960ff168152602001888152602001336001600160a01b03168152602001876001600160a01b03168152602001868152602001438152602001600067ffffffffffffffff81111561065b5761065b61209e565b604051908082528060200260200182016040528015610684578160200160208202803683370190505b509052606b546000908152606d602090815260409091208251815460ff191660ff90911617815582820151805191926106c592600185019290910190611e9d565b5060408201516002820180546001600160a01b039283166001600160a01b03199182161790915560608401516003840180549190931691161790556080820151805161071b916004840191602090910190611e9d565b5060a0820151600582015560c08201518051610741916006840191602090910190611f21565b5050606b54606c80546001810182556000919091527f2b4a51ab505fc96a0952efda2ba61bcd3078d4c02c39a186ec16f21883fbe0160181905560405133925060ff8b1691907fe164703263a5127c249f60b9b2536e9e15ff486c85cfe32fd2a128eae337b184906107ba908b908b908e904390612433565b60405180910390a4606b546107d090600161247b565b606b555050505050505050565b606554606680546000926060926001600160a01b0390911691819061080190612493565b80601f016020809104026020016040519081016040528092919081815260200182805461082d90612493565b801561087a5780601f1061084f5761010080835404028352916020019161087a565b820191906000526020600020905b81548152906001019060200180831161085d57829003601f168201915b50505050509050915091509091565b6033546001600160a01b031633146108e35760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a8565b6108ed6000611e4b565b565b606c546000908190815b818110156109435784606c8281548110610915576109156123c2565b9060005260206000200154036109315760019350809250610943565b8061093b816123ee565b9150506108f9565b50826109615760405162461bcd60e51b81526004016102a8906124cd565b6000848152606d60205260409020600201546001600160a01b0316331461099a5760405162461bcd60e51b81526004016102a890612407565b6067546000858152606d602052604090205460ff16600303610c3d57606a546000868152606d60205260409020600501546109d5919061247b565b4311610a2d5760405162461bcd60e51b815260206004820152602160248201527f4e6f74207265616368696e67207468652065666665637469766520706572696f6044820152601960fa1b60648201526084016102a8565b60008060005b83811015610aaa576000888152606d6020526040902060030154606780546001600160a01b039092169183908110610a6d57610a6d6123c2565b60009182526020909120600290910201546001600160a01b031603610a985760019150809250610aaa565b80610aa2816123ee565b915050610a33565b5080610aec5760405162461bcd60e51b8152602060048201526011602482015270696e76616c69642063616e64696461746560781b60448201526064016102a8565b815b610af96001856124fa565b811015610ba3576067610b0d82600161247b565b81548110610b1d57610b1d6123c2565b906000526020600020906002020160678281548110610b3e57610b3e6123c2565b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b039092169190911781556001808301805491830191610b8190612493565b610b8c929190611f76565b509050508080610b9b906123ee565b915050610aee565b506067610bb16001856124fa565b81548110610bc157610bc16123c2565b60009182526020822060029091020180546001600160a01b031916815590610bec6001830182611ff1565b50506067805480610bff57610bff612511565b60008281526020812060026000199093019283020180546001600160a01b031916815590610c306001830182611ff1565b5050905550505050505050565b6069546068546000878152606d6020526040902060050154610c5f919061247b565b610c69919061247b565b4311610cae5760405162461bcd60e51b8152602060048201526014602482015273159bdd1a5b99c81a185cc81b9bdd08195b99195960621b60448201526064016102a8565b6000858152606d6020526040812060060154906003610cce846002612527565b610cd89190612546565b90508082118015611069576000888152606d602052604090205460ff16600103610e5f5760408051808201825260008a8152606d602081815293822060038101546001600160a01b031684528c835290845260040180549193830191610d3d90612493565b80601f0160208091040260200160405190810160405280929190818152602001828054610d6990612493565b8015610db65780601f10610d8b57610100808354040283529160200191610db6565b820191906000526020600020905b815481529060010190602001808311610d9957829003601f168201915b50505091909252505060678054600181018255600091909152815160029091027f9787eeb91fe3101235e4a76063c7023ecb40f923f97916639c598592fa30d6ae810180546001600160a01b039093166001600160a01b031990931692909217825560208084015180519495508594610e56937f9787eeb91fe3101235e4a76063c7023ecb40f923f97916639c598592fa30d6af01929190910190611e9d565b50505050611069565b60008060005b86811015610edc5760008b8152606d6020526040902060030154606780546001600160a01b039092169183908110610e9f57610e9f6123c2565b60009182526020909120600290910201546001600160a01b031603610eca5760019150809250610edc565b80610ed4816123ee565b915050610e65565b5080610f1e5760405162461bcd60e51b8152602060048201526011602482015270696e76616c69642063616e64696461746560781b60448201526064016102a8565b815b610f2b6001886124fa565b811015610fd5576067610f3f82600161247b565b81548110610f4f57610f4f6123c2565b906000526020600020906002020160678281548110610f7057610f706123c2565b60009182526020909120825460029092020180546001600160a01b0319166001600160a01b039092169190911781556001808301805491830191610fb390612493565b610fbe929190611f76565b509050508080610fcd906123ee565b915050610f20565b506067610fe36001886124fa565b81548110610ff357610ff36123c2565b60009182526020822060029091020180546001600160a01b03191681559061101e6001830182611ff1565b5050606780548061103157611031612511565b60008281526020812060026000199093019283020180546001600160a01b0319168155906110626001830182611ff1565b5050905550505b855b6110766001876124fa565b8110156110d657606c61108a82600161247b565b8154811061109a5761109a6123c2565b9060005260206000200154606c82815481106110b8576110b86123c2565b600091825260209091200155806110ce816123ee565b91505061106b565b50606c6110e46001876124fa565b815481106110f4576110f46123c2565b6000918252602082200155606c80548061111057611110612511565b600082815260208082208301600019908101839055909201909255898252606d905260408120805460ff191681559061114c6001830182611ff1565b6002820180546001600160a01b0319908116909155600383018054909116905561117a600483016000611ff1565b6005820160009055600682016000611192919061202b565b5050877f35baaec8d9c9da1c77f27119243a4384b49eae2bf0680aa0eebe3b62d28574a5826040516111c8911515815260200190565b60405180910390a25050505050505050565b60606000821161123c57606c80548060200260200160405190810160405280929190818152602001828054801561123057602002820191906000526020600020905b81548152602001906001019080831161121c575b50505050509050919050565b606c805480602002602001604051908101604052809291908181526020018280548015611230576020028201919060005260206000209081548152602001906001019080831161121c5750505050509050919050565b919050565b606c54600090815b818110156112e65783606c82815481106112bb576112bb6123c2565b9060005260206000200154036112d457600192506112e6565b806112de816123ee565b91505061129f565b50816113045760405162461bcd60e51b81526004016102a8906124cd565b6000838152606d602052604090205460ff166003036113745760405162461bcd60e51b815260206004820152602660248201527f4175746f6d61746963206578697420646f6573206e6f742072657175697265206044820152656120766f746560d01b60648201526084016102a8565b606754600090815b818110156113db57336001600160a01b0316606782815481106113a1576113a16123c2565b60009182526020909120600290910201546001600160a01b0316036113c957600192506113db565b806113d3816123ee565b91505061137c565b5081806113f257506065546001600160a01b031633145b61140e5760405162461bcd60e51b81526004016102a890612407565b6068546000868152606d602052604090206005015461142d919061247b565b431015801561146757506069546068546000878152606d6020526040902060050154611459919061247b565b611463919061247b565b4311155b6114c85760405162461bcd60e51b815260206004820152602c60248201527f566f74696e672073686f756c642062652077697468696e20746865207370656360448201526b1a599a5959081c195c9a5bd960a21b60648201526084016102a8565b6000858152606d6020526040812060060154815b8181101561153f576000888152606d6020526040902060060180543391908390811061150a5761150a6123c2565b6000918252602090912001546001600160a01b03160361152d576001925061153f565b80611537816123ee565b9150506114dc565b5081156115815760405162461bcd60e51b815260206004820152601060248201526f21b0b713ba103b37ba329030b3b0b4b760811b60448201526064016102a8565b6000878152606d6020908152604080832060060180546001810182559084529282902090920180546001600160a01b03191633908117909155915191825288917fbb62908d9b5227a7f81807cfb2651e7167dc90cf52f25cdbd8f232ea7cb7f9c3910160405180910390a250505050505050565b606754606090819060008167ffffffffffffffff8111156116185761161861209e565b604051908082528060200260200182016040528015611641578160200160208202803683370190505b50905060008267ffffffffffffffff81111561165f5761165f61209e565b60405190808252806020026020018201604052801561169257816020015b606081526020019060019003908161167d5790505b50905060005b838110156117e157606781815481106116b3576116b36123c2565b600091825260209091206002909102015483516001600160a01b03909116908490839081106116e4576116e46123c2565b60200260200101906001600160a01b031690816001600160a01b03168152505060678181548110611717576117176123c2565b9060005260206000209060020201600101805461173390612493565b80601f016020809104026020016040519081016040528092919081815260200182805461175f90612493565b80156117ac5780601f10611781576101008083540402835291602001916117ac565b820191906000526020600020905b81548152906001019060200180831161178f57829003601f168201915b50505050508282815181106117c3576117c36123c2565b602002602001018190525080806117d9906123ee565b915050611698565b509094909350915050565b6000818152606d602052604081208054600382015460028301546005840154600185018054606096889688968896879660ff90931695946001600160a01b0392831694600401939190921691859061184390612493565b80601f016020809104026020016040519081016040528092919081815260200182805461186f90612493565b80156118bc5780601f10611891576101008083540402835291602001916118bc565b820191906000526020600020905b81548152906001019060200180831161189f57829003601f168201915b505050505094508280546118cf90612493565b80601f01602080910402602001604051908101604052809291908181526020018280546118fb90612493565b80156119485780601f1061191d57610100808354040283529160200191611948565b820191906000526020600020905b81548152906001019060200180831161192b57829003601f168201915b5050505050925095509550955095509550955091939550919395565b606c546000908190815b818110156119b85784606c828154811061198a5761198a6123c2565b9060005260206000200154036119a657600193508092506119b8565b806119b0816123ee565b91505061196e565b50826119d65760405162461bcd60e51b81526004016102a8906124cd565b6000848152606d60205260409020600201546001600160a01b03163314611a0f5760405162461bcd60e51b81526004016102a890612407565b6000848152606d602052604090205460ff1660021901611a9357606a546000858152606d6020526040902060050154611a48919061247b565b4310611a8e5760405162461bcd60e51b815260206004820152601560248201527470726f706f73616c2069732065666665637469766560581b60448201526064016102a8565b611b1c565b6068546000858152606d6020526040902060050154611ab2919061247b565b4310611b1c5760405162461bcd60e51b815260206004820152603360248201527f566f74696e672068617320616c7265616479207374617274656420616e642063604482015272185b9b9bdd081899481dda5d1a191c985dd959606a1b60648201526084016102a8565b815b611b296001836124fa565b811015611b8957606c611b3d82600161247b565b81548110611b4d57611b4d6123c2565b9060005260206000200154606c8281548110611b6b57611b6b6123c2565b60009182526020909120015580611b81816123ee565b915050611b1e565b50606c611b976001836124fa565b81548110611ba757611ba76123c2565b6000918252602082200155606c805480611bc357611bc3612511565b600082815260208082208301600019908101839055909201909255858252606d905260408120805460ff1916815590611bff6001830182611ff1565b6002820180546001600160a01b03199081169091556003830180549091169055611c2d600483016000611ff1565b6005820160009055600682016000611c45919061202b565b5050837f98397210b92c60979a43a58e8031fa9ec208314be7aadd29a6808e2a67296bbf43604051611c7991815260200190565b60405180910390a250505050565b6033546001600160a01b03163314611ce15760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a8565b6001600160a01b038116611d465760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102a8565b611d4f81611e4b565b50565b600054610100900460ff16611d6d5760005460ff1615611d71565b303b155b611dd45760405162461bcd60e51b815260206004820152602e60248201527f496e697469616c697a61626c653a20636f6e747261637420697320616c72656160448201526d191e481a5b9a5d1a585b1a5e995960921b60648201526084016102a8565b600054610100900460ff16158015611df6576000805461ffff19166101011790555b6000606b55606580546001600160a01b031916331790558151611e20906066906020850190611e9d565b506201518060685562093a8060695561a8c0606a558015611e47576000805461ff00191690555b5050565b603380546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a35050565b828054611ea990612493565b90600052602060002090601f016020900481019282611ecb5760008555611f11565b82601f10611ee457805160ff1916838001178555611f11565b82800160010185558215611f11579182015b82811115611f11578251825591602001919060010190611ef6565b50611f1d929150612045565b5090565b828054828255906000526020600020908101928215611f11579160200282015b82811115611f1157825182546001600160a01b0319166001600160a01b03909116178255602090920191600190910190611f41565b828054611f8290612493565b90600052602060002090601f016020900481019282611fa45760008555611f11565b82601f10611fb55780548555611f11565b82800160010185558215611f1157600052602060002091601f016020900482015b82811115611f11578254825591600101919060010190611fd6565b508054611ffd90612493565b6000825580601f1061200d575050565b601f016020900490600052602060002090810190611d4f9190612045565b5080546000825590600052602060002090810190611d4f91905b5b80821115611f1d5760008155600101612046565b6020808252825182820181905260009190848201906040850190845b8181101561209257835183529284019291840191600101612076565b50909695505050505050565b634e487b7160e01b600052604160045260246000fd5b600082601f8301126120c557600080fd5b813567ffffffffffffffff808211156120e0576120e061209e565b604051601f8301601f19908116603f011681019082821181831017156121085761210861209e565b8160405283815286602085880101111561212157600080fd5b836020870160208301376000602085830101528094505050505092915050565b80356001600160a01b038116811461129257600080fd5b6000806000806080858703121561216e57600080fd5b843560ff8116811461217f57600080fd5b9350602085013567ffffffffffffffff8082111561219c57600080fd5b6121a8888389016120b4565b94506121b660408801612141565b935060608701359150808211156121cc57600080fd5b506121d9878288016120b4565b91505092959194509250565b6000815180845260005b8181101561220b576020818501810151868301820152016121ef565b8181111561221d576000602083870101525b50601f01601f19169290920160200192915050565b6001600160a01b0383168152604060208201819052600090612256908301846121e5565b949350505050565b60006020828403121561227057600080fd5b5035919050565b604080825283519082018190526000906020906060840190828701845b828110156122b95781516001600160a01b031684529284019290840190600101612294565b50505083810382850152845180825282820190600581901b8301840187850160005b8381101561230957601f198684030185526122f78383516121e5565b948701949250908601906001016122db565b50909998505050505050505050565b60ff8716815260c06020820152600061233460c08301886121e5565b6001600160a01b038781166040850152838203606085015261235682886121e5565b95166080840152505060a00152949350505050565b60006020828403121561237d57600080fd5b61238682612141565b9392505050565b60006020828403121561239f57600080fd5b813567ffffffffffffffff8111156123b657600080fd5b612256848285016120b4565b634e487b7160e01b600052603260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b600060018201612400576124006123d8565b5060010190565b60208082526012908201527134b73b30b634b21036b9b39739b2b73232b960711b604082015260600190565b6001600160a01b0385168152608060208201819052600090612457908301866121e5565b828103604084015261246981866121e5565b91505082606083015295945050505050565b6000821982111561248e5761248e6123d8565b500190565b600181811c908216806124a757607f821691505b6020821081036124c757634e487b7160e01b600052602260045260246000fd5b50919050565b6020808252601390820152721a5b9d985b1a59081c1c9bdc1bdcd85b081a59606a1b604082015260600190565b60008282101561250c5761250c6123d8565b500390565b634e487b7160e01b600052603160045260246000fd5b6000816000190483118215151615612541576125416123d8565b500290565b60008261256357634e487b7160e01b600052601260045260246000fd5b50049056fea2646970667358221220ee4f9ebbe09407f1657e78b5988c9aba79f8394a1b4e782a8ff12ffb9506b7cf64736f6c634300080d0033",
}

// VoteABI is the input ABI used to generate the binding from.
// Deprecated: Use VoteMetaData.ABI instead.
var VoteABI = VoteMetaData.ABI

// VoteBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use VoteMetaData.Bin instead.
var VoteBin = VoteMetaData.Bin

// DeployVote deploys a new Ethereum contract, binding an instance of Vote to it.
func DeployVote(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Vote, error) {
	parsed, err := VoteMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(VoteBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// Vote is an auto generated Go binding around an Ethereum contract.
type Vote struct {
	VoteCaller     // Read-only binding to the contract
	VoteTransactor // Write-only binding to the contract
	VoteFilterer   // Log filterer for contract events
}

// VoteCaller is an auto generated read-only Go binding around an Ethereum contract.
type VoteCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VoteTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VoteFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VoteSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VoteSession struct {
	Contract     *Vote             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VoteCallerSession struct {
	Contract *VoteCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VoteTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VoteTransactorSession struct {
	Contract     *VoteTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VoteRaw is an auto generated low-level Go binding around an Ethereum contract.
type VoteRaw struct {
	Contract *Vote // Generic contract binding to access the raw methods on
}

// VoteCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VoteCallerRaw struct {
	Contract *VoteCaller // Generic read-only contract binding to access the raw methods on
}

// VoteTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VoteTransactorRaw struct {
	Contract *VoteTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVote creates a new instance of Vote, bound to a specific deployed contract.
func NewVote(address common.Address, backend bind.ContractBackend) (*Vote, error) {
	contract, err := bindVote(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vote{VoteCaller: VoteCaller{contract: contract}, VoteTransactor: VoteTransactor{contract: contract}, VoteFilterer: VoteFilterer{contract: contract}}, nil
}

// NewVoteCaller creates a new read-only instance of Vote, bound to a specific deployed contract.
func NewVoteCaller(address common.Address, caller bind.ContractCaller) (*VoteCaller, error) {
	contract, err := bindVote(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VoteCaller{contract: contract}, nil
}

// NewVoteTransactor creates a new write-only instance of Vote, bound to a specific deployed contract.
func NewVoteTransactor(address common.Address, transactor bind.ContractTransactor) (*VoteTransactor, error) {
	contract, err := bindVote(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VoteTransactor{contract: contract}, nil
}

// NewVoteFilterer creates a new log filterer instance of Vote, bound to a specific deployed contract.
func NewVoteFilterer(address common.Address, filterer bind.ContractFilterer) (*VoteFilterer, error) {
	contract, err := bindVote(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VoteFilterer{contract: contract}, nil
}

// bindVote binds a generic wrapper to an already deployed contract.
func bindVote(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(VoteABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.VoteCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.VoteTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vote *VoteCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vote.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vote *VoteTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vote *VoteTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vote.Contract.contract.Transact(opts, method, params...)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address, string)
func (_Vote *VoteCaller) GetAdmin(opts *bind.CallOpts) (common.Address, string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getAdmin")

	if err != nil {
		return *new(common.Address), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)

	return out0, out1, err

}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address, string)
func (_Vote *VoteSession) GetAdmin() (common.Address, string, error) {
	return _Vote.Contract.GetAdmin(&_Vote.CallOpts)
}

// GetAdmin is a free data retrieval call binding the contract method 0x6e9960c3.
//
// Solidity: function getAdmin() view returns(address, string)
func (_Vote *VoteCallerSession) GetAdmin() (common.Address, string, error) {
	return _Vote.Contract.GetAdmin(&_Vote.CallOpts)
}

// GetAllAuthority is a free data retrieval call binding the contract method 0x8eee1588.
//
// Solidity: function getAllAuthority() view returns(address[], string[])
func (_Vote *VoteCaller) GetAllAuthority(opts *bind.CallOpts) ([]common.Address, []string, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getAllAuthority")

	if err != nil {
		return *new([]common.Address), *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	out1 := *abi.ConvertType(out[1], new([]string)).(*[]string)

	return out0, out1, err

}

// GetAllAuthority is a free data retrieval call binding the contract method 0x8eee1588.
//
// Solidity: function getAllAuthority() view returns(address[], string[])
func (_Vote *VoteSession) GetAllAuthority() ([]common.Address, []string, error) {
	return _Vote.Contract.GetAllAuthority(&_Vote.CallOpts)
}

// GetAllAuthority is a free data retrieval call binding the contract method 0x8eee1588.
//
// Solidity: function getAllAuthority() view returns(address[], string[])
func (_Vote *VoteCallerSession) GetAllAuthority() ([]common.Address, []string, error) {
	return _Vote.Contract.GetAllAuthority(&_Vote.CallOpts)
}

// GetAllProposalId is a free data retrieval call binding the contract method 0x51ba9c54.
//
// Solidity: function getAllProposalId() view returns(uint256[])
func (_Vote *VoteCaller) GetAllProposalId(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getAllProposalId")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllProposalId is a free data retrieval call binding the contract method 0x51ba9c54.
//
// Solidity: function getAllProposalId() view returns(uint256[])
func (_Vote *VoteSession) GetAllProposalId() ([]*big.Int, error) {
	return _Vote.Contract.GetAllProposalId(&_Vote.CallOpts)
}

// GetAllProposalId is a free data retrieval call binding the contract method 0x51ba9c54.
//
// Solidity: function getAllProposalId() view returns(uint256[])
func (_Vote *VoteCallerSession) GetAllProposalId() ([]*big.Int, error) {
	return _Vote.Contract.GetAllProposalId(&_Vote.CallOpts)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns(uint8, string, address, string, address, uint256)
func (_Vote *VoteCaller) GetProposal(opts *bind.CallOpts, proposalId *big.Int) (uint8, string, common.Address, string, common.Address, *big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getProposal", proposalId)

	if err != nil {
		return *new(uint8), *new(string), *new(common.Address), *new(string), *new(common.Address), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	out3 := *abi.ConvertType(out[3], new(string)).(*string)
	out4 := *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	out5 := *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, out4, out5, err

}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns(uint8, string, address, string, address, uint256)
func (_Vote *VoteSession) GetProposal(proposalId *big.Int) (uint8, string, common.Address, string, common.Address, *big.Int, error) {
	return _Vote.Contract.GetProposal(&_Vote.CallOpts, proposalId)
}

// GetProposal is a free data retrieval call binding the contract method 0xc7f758a8.
//
// Solidity: function getProposal(uint256 proposalId) view returns(uint8, string, address, string, address, uint256)
func (_Vote *VoteCallerSession) GetProposal(proposalId *big.Int) (uint8, string, common.Address, string, common.Address, *big.Int, error) {
	return _Vote.Contract.GetProposal(&_Vote.CallOpts, proposalId)
}

// GetProposalId is a free data retrieval call binding the contract method 0x7c95ac27.
//
// Solidity: function getProposalId(uint256 blockNo) view returns(uint256[])
func (_Vote *VoteCaller) GetProposalId(opts *bind.CallOpts, blockNo *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "getProposalId", blockNo)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetProposalId is a free data retrieval call binding the contract method 0x7c95ac27.
//
// Solidity: function getProposalId(uint256 blockNo) view returns(uint256[])
func (_Vote *VoteSession) GetProposalId(blockNo *big.Int) ([]*big.Int, error) {
	return _Vote.Contract.GetProposalId(&_Vote.CallOpts, blockNo)
}

// GetProposalId is a free data retrieval call binding the contract method 0x7c95ac27.
//
// Solidity: function getProposalId(uint256 blockNo) view returns(uint256[])
func (_Vote *VoteCallerSession) GetProposalId(blockNo *big.Int) ([]*big.Int, error) {
	return _Vote.Contract.GetProposalId(&_Vote.CallOpts, blockNo)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vote.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Vote *VoteCallerSession) Owner() (common.Address, error) {
	return _Vote.Contract.Owner(&_Vote.CallOpts)
}

// EffectProposal is a paid mutator transaction binding the contract method 0x79717e88.
//
// Solidity: function effectProposal(uint256 proposalId) returns()
func (_Vote *VoteTransactor) EffectProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "effectProposal", proposalId)
}

// EffectProposal is a paid mutator transaction binding the contract method 0x79717e88.
//
// Solidity: function effectProposal(uint256 proposalId) returns()
func (_Vote *VoteSession) EffectProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.EffectProposal(&_Vote.TransactOpts, proposalId)
}

// EffectProposal is a paid mutator transaction binding the contract method 0x79717e88.
//
// Solidity: function effectProposal(uint256 proposalId) returns()
func (_Vote *VoteTransactorSession) EffectProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.EffectProposal(&_Vote.TransactOpts, proposalId)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string serviceUrl) returns()
func (_Vote *VoteTransactor) Initialize(opts *bind.TransactOpts, serviceUrl string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "initialize", serviceUrl)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string serviceUrl) returns()
func (_Vote *VoteSession) Initialize(serviceUrl string) (*types.Transaction, error) {
	return _Vote.Contract.Initialize(&_Vote.TransactOpts, serviceUrl)
}

// Initialize is a paid mutator transaction binding the contract method 0xf62d1888.
//
// Solidity: function initialize(string serviceUrl) returns()
func (_Vote *VoteTransactorSession) Initialize(serviceUrl string) (*types.Transaction, error) {
	return _Vote.Contract.Initialize(&_Vote.TransactOpts, serviceUrl)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vote *VoteTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vote *VoteSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vote.Contract.RenounceOwnership(&_Vote.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Vote *VoteTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Vote.Contract.RenounceOwnership(&_Vote.TransactOpts)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x5944d7af.
//
// Solidity: function submitProposal(uint8 proposalType, string proposalUrl, address candidate, string candidateServiceUrl) returns()
func (_Vote *VoteTransactor) SubmitProposal(opts *bind.TransactOpts, proposalType uint8, proposalUrl string, candidate common.Address, candidateServiceUrl string) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "submitProposal", proposalType, proposalUrl, candidate, candidateServiceUrl)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x5944d7af.
//
// Solidity: function submitProposal(uint8 proposalType, string proposalUrl, address candidate, string candidateServiceUrl) returns()
func (_Vote *VoteSession) SubmitProposal(proposalType uint8, proposalUrl string, candidate common.Address, candidateServiceUrl string) (*types.Transaction, error) {
	return _Vote.Contract.SubmitProposal(&_Vote.TransactOpts, proposalType, proposalUrl, candidate, candidateServiceUrl)
}

// SubmitProposal is a paid mutator transaction binding the contract method 0x5944d7af.
//
// Solidity: function submitProposal(uint8 proposalType, string proposalUrl, address candidate, string candidateServiceUrl) returns()
func (_Vote *VoteTransactorSession) SubmitProposal(proposalType uint8, proposalUrl string, candidate common.Address, candidateServiceUrl string) (*types.Transaction, error) {
	return _Vote.Contract.SubmitProposal(&_Vote.TransactOpts, proposalType, proposalUrl, candidate, candidateServiceUrl)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vote *VoteTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vote *VoteSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vote.Contract.TransferOwnership(&_Vote.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Vote *VoteTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Vote.Contract.TransferOwnership(&_Vote.TransactOpts, newOwner)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x807896d5.
//
// Solidity: function voteProposal(uint256 proposalId) returns()
func (_Vote *VoteTransactor) VoteProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "voteProposal", proposalId)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x807896d5.
//
// Solidity: function voteProposal(uint256 proposalId) returns()
func (_Vote *VoteSession) VoteProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.VoteProposal(&_Vote.TransactOpts, proposalId)
}

// VoteProposal is a paid mutator transaction binding the contract method 0x807896d5.
//
// Solidity: function voteProposal(uint256 proposalId) returns()
func (_Vote *VoteTransactorSession) VoteProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.VoteProposal(&_Vote.TransactOpts, proposalId)
}

// WithdrawProposal is a paid mutator transaction binding the contract method 0xe1f02ffa.
//
// Solidity: function withdrawProposal(uint256 proposalId) returns()
func (_Vote *VoteTransactor) WithdrawProposal(opts *bind.TransactOpts, proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.contract.Transact(opts, "withdrawProposal", proposalId)
}

// WithdrawProposal is a paid mutator transaction binding the contract method 0xe1f02ffa.
//
// Solidity: function withdrawProposal(uint256 proposalId) returns()
func (_Vote *VoteSession) WithdrawProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.WithdrawProposal(&_Vote.TransactOpts, proposalId)
}

// WithdrawProposal is a paid mutator transaction binding the contract method 0xe1f02ffa.
//
// Solidity: function withdrawProposal(uint256 proposalId) returns()
func (_Vote *VoteTransactorSession) WithdrawProposal(proposalId *big.Int) (*types.Transaction, error) {
	return _Vote.Contract.WithdrawProposal(&_Vote.TransactOpts, proposalId)
}

// VoteNewProposalIterator is returned from FilterNewProposal and is used to iterate over the raw logs and unpacked data for NewProposal events raised by the Vote contract.
type VoteNewProposalIterator struct {
	Event *VoteNewProposal // Event containing the contract specifics and raw log

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
func (it *VoteNewProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteNewProposal)
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
		it.Event = new(VoteNewProposal)
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
func (it *VoteNewProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteNewProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteNewProposal represents a NewProposal event raised by the Vote contract.
type VoteNewProposal struct {
	ProposalId          *big.Int
	ProposalType        uint8
	Submitter           common.Address
	Candidate           common.Address
	CandidateServiceUrl string
	ProposalUrl         string
	SubmitBlockNo       *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterNewProposal is a free log retrieval operation binding the contract event 0xe164703263a5127c249f60b9b2536e9e15ff486c85cfe32fd2a128eae337b184.
//
// Solidity: event NewProposal(uint256 indexed proposalId, uint8 indexed proposalType, address indexed submitter, address candidate, string candidateServiceUrl, string proposalUrl, uint256 submitBlockNo)
func (_Vote *VoteFilterer) FilterNewProposal(opts *bind.FilterOpts, proposalId []*big.Int, proposalType []uint8, submitter []common.Address) (*VoteNewProposalIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalTypeRule []interface{}
	for _, proposalTypeItem := range proposalType {
		proposalTypeRule = append(proposalTypeRule, proposalTypeItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "NewProposal", proposalIdRule, proposalTypeRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return &VoteNewProposalIterator{contract: _Vote.contract, event: "NewProposal", logs: logs, sub: sub}, nil
}

// WatchNewProposal is a free log subscription operation binding the contract event 0xe164703263a5127c249f60b9b2536e9e15ff486c85cfe32fd2a128eae337b184.
//
// Solidity: event NewProposal(uint256 indexed proposalId, uint8 indexed proposalType, address indexed submitter, address candidate, string candidateServiceUrl, string proposalUrl, uint256 submitBlockNo)
func (_Vote *VoteFilterer) WatchNewProposal(opts *bind.WatchOpts, sink chan<- *VoteNewProposal, proposalId []*big.Int, proposalType []uint8, submitter []common.Address) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}
	var proposalTypeRule []interface{}
	for _, proposalTypeItem := range proposalType {
		proposalTypeRule = append(proposalTypeRule, proposalTypeItem)
	}
	var submitterRule []interface{}
	for _, submitterItem := range submitter {
		submitterRule = append(submitterRule, submitterItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "NewProposal", proposalIdRule, proposalTypeRule, submitterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteNewProposal)
				if err := _Vote.contract.UnpackLog(event, "NewProposal", log); err != nil {
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

// ParseNewProposal is a log parse operation binding the contract event 0xe164703263a5127c249f60b9b2536e9e15ff486c85cfe32fd2a128eae337b184.
//
// Solidity: event NewProposal(uint256 indexed proposalId, uint8 indexed proposalType, address indexed submitter, address candidate, string candidateServiceUrl, string proposalUrl, uint256 submitBlockNo)
func (_Vote *VoteFilterer) ParseNewProposal(log types.Log) (*VoteNewProposal, error) {
	event := new(VoteNewProposal)
	if err := _Vote.contract.UnpackLog(event, "NewProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Vote contract.
type VoteOwnershipTransferredIterator struct {
	Event *VoteOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *VoteOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteOwnershipTransferred)
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
		it.Event = new(VoteOwnershipTransferred)
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
func (it *VoteOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteOwnershipTransferred represents a OwnershipTransferred event raised by the Vote contract.
type VoteOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vote *VoteFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*VoteOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &VoteOwnershipTransferredIterator{contract: _Vote.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vote *VoteFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *VoteOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteOwnershipTransferred)
				if err := _Vote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Vote *VoteFilterer) ParseOwnershipTransferred(log types.Log) (*VoteOwnershipTransferred, error) {
	event := new(VoteOwnershipTransferred)
	if err := _Vote.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteProposalResultIterator is returned from FilterProposalResult and is used to iterate over the raw logs and unpacked data for ProposalResult events raised by the Vote contract.
type VoteProposalResultIterator struct {
	Event *VoteProposalResult // Event containing the contract specifics and raw log

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
func (it *VoteProposalResultIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteProposalResult)
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
		it.Event = new(VoteProposalResult)
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
func (it *VoteProposalResultIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteProposalResultIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteProposalResult represents a ProposalResult event raised by the Vote contract.
type VoteProposalResult struct {
	ProposalId *big.Int
	Result     bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalResult is a free log retrieval operation binding the contract event 0x35baaec8d9c9da1c77f27119243a4384b49eae2bf0680aa0eebe3b62d28574a5.
//
// Solidity: event ProposalResult(uint256 indexed proposalId, bool result)
func (_Vote *VoteFilterer) FilterProposalResult(opts *bind.FilterOpts, proposalId []*big.Int) (*VoteProposalResultIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "ProposalResult", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteProposalResultIterator{contract: _Vote.contract, event: "ProposalResult", logs: logs, sub: sub}, nil
}

// WatchProposalResult is a free log subscription operation binding the contract event 0x35baaec8d9c9da1c77f27119243a4384b49eae2bf0680aa0eebe3b62d28574a5.
//
// Solidity: event ProposalResult(uint256 indexed proposalId, bool result)
func (_Vote *VoteFilterer) WatchProposalResult(opts *bind.WatchOpts, sink chan<- *VoteProposalResult, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "ProposalResult", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteProposalResult)
				if err := _Vote.contract.UnpackLog(event, "ProposalResult", log); err != nil {
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

// ParseProposalResult is a log parse operation binding the contract event 0x35baaec8d9c9da1c77f27119243a4384b49eae2bf0680aa0eebe3b62d28574a5.
//
// Solidity: event ProposalResult(uint256 indexed proposalId, bool result)
func (_Vote *VoteFilterer) ParseProposalResult(log types.Log) (*VoteProposalResult, error) {
	event := new(VoteProposalResult)
	if err := _Vote.contract.UnpackLog(event, "ProposalResult", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteVoteProposalIterator is returned from FilterVoteProposal and is used to iterate over the raw logs and unpacked data for VoteProposal events raised by the Vote contract.
type VoteVoteProposalIterator struct {
	Event *VoteVoteProposal // Event containing the contract specifics and raw log

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
func (it *VoteVoteProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteVoteProposal)
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
		it.Event = new(VoteVoteProposal)
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
func (it *VoteVoteProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteVoteProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteVoteProposal represents a VoteProposal event raised by the Vote contract.
type VoteVoteProposal struct {
	ProposalId *big.Int
	Voter      common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteProposal is a free log retrieval operation binding the contract event 0xbb62908d9b5227a7f81807cfb2651e7167dc90cf52f25cdbd8f232ea7cb7f9c3.
//
// Solidity: event VoteProposal(uint256 indexed proposalId, address voter)
func (_Vote *VoteFilterer) FilterVoteProposal(opts *bind.FilterOpts, proposalId []*big.Int) (*VoteVoteProposalIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "VoteProposal", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteVoteProposalIterator{contract: _Vote.contract, event: "VoteProposal", logs: logs, sub: sub}, nil
}

// WatchVoteProposal is a free log subscription operation binding the contract event 0xbb62908d9b5227a7f81807cfb2651e7167dc90cf52f25cdbd8f232ea7cb7f9c3.
//
// Solidity: event VoteProposal(uint256 indexed proposalId, address voter)
func (_Vote *VoteFilterer) WatchVoteProposal(opts *bind.WatchOpts, sink chan<- *VoteVoteProposal, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "VoteProposal", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteVoteProposal)
				if err := _Vote.contract.UnpackLog(event, "VoteProposal", log); err != nil {
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

// ParseVoteProposal is a log parse operation binding the contract event 0xbb62908d9b5227a7f81807cfb2651e7167dc90cf52f25cdbd8f232ea7cb7f9c3.
//
// Solidity: event VoteProposal(uint256 indexed proposalId, address voter)
func (_Vote *VoteFilterer) ParseVoteProposal(log types.Log) (*VoteVoteProposal, error) {
	event := new(VoteVoteProposal)
	if err := _Vote.contract.UnpackLog(event, "VoteProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VoteWithdrawProposalIterator is returned from FilterWithdrawProposal and is used to iterate over the raw logs and unpacked data for WithdrawProposal events raised by the Vote contract.
type VoteWithdrawProposalIterator struct {
	Event *VoteWithdrawProposal // Event containing the contract specifics and raw log

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
func (it *VoteWithdrawProposalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VoteWithdrawProposal)
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
		it.Event = new(VoteWithdrawProposal)
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
func (it *VoteWithdrawProposalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VoteWithdrawProposalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VoteWithdrawProposal represents a WithdrawProposal event raised by the Vote contract.
type VoteWithdrawProposal struct {
	ProposalId *big.Int
	BlockNo    *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterWithdrawProposal is a free log retrieval operation binding the contract event 0x98397210b92c60979a43a58e8031fa9ec208314be7aadd29a6808e2a67296bbf.
//
// Solidity: event WithdrawProposal(uint256 indexed proposalId, uint256 blockNo)
func (_Vote *VoteFilterer) FilterWithdrawProposal(opts *bind.FilterOpts, proposalId []*big.Int) (*VoteWithdrawProposalIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Vote.contract.FilterLogs(opts, "WithdrawProposal", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &VoteWithdrawProposalIterator{contract: _Vote.contract, event: "WithdrawProposal", logs: logs, sub: sub}, nil
}

// WatchWithdrawProposal is a free log subscription operation binding the contract event 0x98397210b92c60979a43a58e8031fa9ec208314be7aadd29a6808e2a67296bbf.
//
// Solidity: event WithdrawProposal(uint256 indexed proposalId, uint256 blockNo)
func (_Vote *VoteFilterer) WatchWithdrawProposal(opts *bind.WatchOpts, sink chan<- *VoteWithdrawProposal, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _Vote.contract.WatchLogs(opts, "WithdrawProposal", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VoteWithdrawProposal)
				if err := _Vote.contract.UnpackLog(event, "WithdrawProposal", log); err != nil {
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

// ParseWithdrawProposal is a log parse operation binding the contract event 0x98397210b92c60979a43a58e8031fa9ec208314be7aadd29a6808e2a67296bbf.
//
// Solidity: event WithdrawProposal(uint256 indexed proposalId, uint256 blockNo)
func (_Vote *VoteFilterer) ParseWithdrawProposal(log types.Log) (*VoteWithdrawProposal, error) {
	event := new(VoteWithdrawProposal)
	if err := _Vote.contract.UnpackLog(event, "WithdrawProposal", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
