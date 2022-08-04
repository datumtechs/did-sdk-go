// Copyright metis network contributors
// SPDX-License-Identifier: (Apache-2.0 AND CC-BY-4.0)
// Code is Apache-2.0 and docs are CC-BY-4.0

pragma solidity ^0.8.0;

import "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

contract Vote is Initializable, OwnableUpgradeable{

    function initialize(string memory serviceUrl) public initializer {
        _proposalId = 0;
        _admin.addr = msg.sender;
        _admin.serviceUrl = serviceUrl;

        RANGE_BEGIN_VOTE = 24 * 60 * 60;

        RANGE_VOTE = 7 * 24 * 60 * 60;

        RANGE_QUIT = 12 * 60 * 60;
    }

    struct AuthorityInfo{
        address addr;
        string serviceUrl;
    }

    AuthorityInfo _admin;
    AuthorityInfo[] private _authorityList;


    uint256 constant ADD_AUTHORITY= 1;
    uint256 constant KICK_OUT_AUTHORITY= 2;
    uint256 constant AUTO_QUIT_AUTHORITY= 3;


    uint256 RANGE_BEGIN_VOTE;
    function setRangeBeginVote(uint256 range)  public{
        require(msg.sender == _admin.addr, "Only admin can do this");
        RANGE_BEGIN_VOTE = range;
    }

    uint256 RANGE_VOTE;
    function setRangeVote(uint256 range)  public{
        require(msg.sender == _admin.addr, "Only admin can do this");
        RANGE_VOTE = range;
    }


    uint256 RANGE_QUIT;
    function setRangeQuit(uint256 range)  public{
        require(msg.sender == _admin.addr, "Only admin can do this");
        RANGE_QUIT = range;
    }

    struct ProposalInfo {
        uint256 proposalType;
        string publicInfoUrl;
        address senderAddress;
        address operatorAddress;
        string operatorServiceUrl;
        uint256 proposalBlockNumber;
        address[] voter;
    }

    uint256 private _proposalId;
    mapping(uint256 => ProposalInfo) private _proposalInfo;
    uint256[] private _proposalList;

    event NewProposal(
        uint256 indexed proposalId,
        uint256 indexed proposalType,
        address indexed sender,
        address operator,
        string operatorServiceUrl,
        string publicInfoUrl,
        uint256 proposalBlockNumber
    );

    event RemoveProposal(
        uint256 indexed proposalId,
        uint256 blockNumber
    );

    event VoteProposal(
        uint256 indexed proposalId,
        address voter
    );

    event ProposalResult(
        uint256 indexed proposalId,
        bool result
    );


    function submitProposal(uint256 proposalType, string memory publicInfoUrl, address operator, string memory serviceUrl) public{

        require(proposalType == ADD_AUTHORITY || proposalType == KICK_OUT_AUTHORITY || proposalType == AUTO_QUIT_AUTHORITY,  "Invalid Proposal type");

        bool findSender = false;
        bool findOperator = false;

        uint256 authorityLength = _authorityList.length;
        for (uint256 i = 0; i < authorityLength; i++){
            if (_authorityList[i].addr == msg.sender){
                findSender = true;
            }

            if (_authorityList[i].addr == operator){
                findOperator = true;
            }

            if(findSender && findOperator){
                break;
            }
        }

        require(findSender || msg.sender == _admin.addr, "invalid msg.sender");

        if(ADD_AUTHORITY == proposalType){
            require(!findOperator, "operator is already in the authority list");
        } else {
            require(findOperator, "operator is not in the authority list");
        }


        uint256 proposalLength = _proposalList.length;
        for (uint256 i = 0; i < proposalLength; i++){
            require(_proposalInfo[_proposalList[i]].operatorAddress != operator, "operator is already in one open proposal");
        }
    

        if(proposalType == KICK_OUT_AUTHORITY || proposalType == AUTO_QUIT_AUTHORITY){
            for (uint256 i = 0; i < proposalLength; i++){   
                require(_proposalInfo[_proposalList[i]].senderAddress != operator, "operator has open proposals.");
            }
        }
 
        _proposalInfo[_proposalId] = ProposalInfo({
            proposalType: proposalType,
            publicInfoUrl: publicInfoUrl,
            senderAddress: msg.sender,
            operatorAddress: operator,
            operatorServiceUrl: serviceUrl,
            proposalBlockNumber: block.number,
            voter: new address[](0)
        });

        _proposalList.push(_proposalId);

        emit NewProposal(
            _proposalId,
            proposalType,
            msg.sender,
            operator,
            serviceUrl,
            publicInfoUrl,
            block.number
        );

        _proposalId = _proposalId + 1;
    }


    function removeProposal(uint256 proposalId)  public{
        bool find = false;
        uint256 proposalIndex = 0;
        uint256 proposalLength = _proposalList.length;
        for (uint256 i = 0; i < proposalLength; i++){
            if (_proposalList[i] == proposalId){
                find = true;
                proposalIndex = i;
                break;
            }
        }

        require(find, "invalid proposal id");
        require(_proposalInfo[proposalId].senderAddress == msg.sender, "invalid msg.sender");
        if(_proposalInfo[proposalId].proposalType == AUTO_QUIT_AUTHORITY){
            require(block.number < _proposalInfo[proposalId].proposalBlockNumber + RANGE_QUIT, 
                "Cancellation can only take place before the withdrawal application becomes effective");
        }else{
            require(block.number < _proposalInfo[proposalId].proposalBlockNumber + RANGE_BEGIN_VOTE, "Voting has already started and cannot be removed");
        }

        for(uint256 i = proposalIndex; i < proposalLength - 1; i++){
            _proposalList[i] = _proposalList[i+1];
        }

        delete _proposalList[proposalLength-1];
        _proposalList.pop();

        delete _proposalInfo[proposalId];

        emit RemoveProposal(
            proposalId,
            block.number
        );
    }

    function voteProposal(uint256 proposalId)  public{
        bool find = false;
        uint256 proposalLength = _proposalList.length;
        for (uint256 i = 0; i < proposalLength; i++){
            if (_proposalList[i] == proposalId){
                find = true;
                break;
            }
        }

        require(find, "invalid proposal id");

        require(AUTO_QUIT_AUTHORITY != _proposalInfo[proposalId].proposalType, "Automatic exit does not require a vote");

        bool findSender = false;
        uint256 authorityLength = _authorityList.length;
        for (uint256 i = 0; i < authorityLength; i++){
            if (_authorityList[i].addr == msg.sender){
                findSender = true;
                break;
            }
        }

        require(findSender || msg.sender == _admin.addr, "invalid msg.sender");

        require(_proposalInfo[proposalId].proposalBlockNumber + RANGE_BEGIN_VOTE <= block.number && 
            block.number <= _proposalInfo[proposalId].proposalBlockNumber + RANGE_BEGIN_VOTE + RANGE_VOTE, 
            "Voting has already started and cannot be removed");

        bool findVoter = false;
        uint256 voterLength = _proposalInfo[proposalId].voter.length;
        for (uint256 i = 0; i < voterLength; i++){
            if (_proposalInfo[proposalId].voter[i] == msg.sender){
                findVoter = true;
                break;
            }
        }

        require(!findVoter, "Can't vote again");

        _proposalInfo[proposalId].voter.push(msg.sender);

        emit VoteProposal(
            proposalId,
            msg.sender
        );
    }

    function effectProposal(uint256 proposalId)  public{
        bool find = false;
        uint256 proposalIndex = 0;
        uint256 proposalLength = _proposalList.length;
        for (uint256 i = 0; i < proposalLength; i++){
            if (_proposalList[i] == proposalId){
                find = true;
                proposalIndex = i;
                break;
            }
        }

        require(find, "invalid proposal id");
        require(_proposalInfo[proposalId].senderAddress == msg.sender, "invalid msg.sender");

        uint256 authorityLength = _authorityList.length;
        if(AUTO_QUIT_AUTHORITY == _proposalInfo[proposalId].proposalType){
            require(block.number > _proposalInfo[proposalId].proposalBlockNumber + RANGE_QUIT, "The automatic exit has not yet taken effect");
            uint256 authorityIndex = 0;
            bool findAuthority = false;
            for (uint256 i = 0; i < authorityLength; i++){
                if (_authorityList[i].addr == _proposalInfo[proposalId].operatorAddress){
                    findAuthority = true;
                    authorityIndex = i;
                    break;
                }
            }
            
            require(findAuthority, "invalid operator");

            for(uint256 i = authorityIndex; i < authorityLength - 1; i++){
                _authorityList[i] = _authorityList[i+1];
            }

            delete _authorityList[authorityLength-1];
            _authorityList.pop();

            return;
        }

        require(block.number > _proposalInfo[proposalId].proposalBlockNumber + RANGE_BEGIN_VOTE + RANGE_VOTE, "Voting has not ended, the results cannot be counted");

        uint256 voterLength = _proposalInfo[proposalId].voter.length;

        uint256 waterLevel = (authorityLength * 2) / 3;
        bool voteResult = voterLength > waterLevel;

        if(voteResult){
            if(ADD_AUTHORITY == _proposalInfo[proposalId].proposalType){
                AuthorityInfo memory newAuthority = AuthorityInfo(_proposalInfo[proposalId].operatorAddress, _proposalInfo[proposalId].operatorServiceUrl);
               _authorityList.push(newAuthority);
            } else {
                uint256 authorityIndex = 0;
                bool findAuthority = false;
                for (uint256 i = 0; i < authorityLength; i++){
                    if (_authorityList[i].addr == _proposalInfo[proposalId].operatorAddress){
                        findAuthority = true;
                        authorityIndex = i;
                        break;
                    }
                }
                
                require(findAuthority, "invalid operator");

                for(uint256 i = authorityIndex; i < authorityLength - 1; i++){
                    _authorityList[i] = _authorityList[i+1];
                }

                delete _authorityList[authorityLength-1];
                _authorityList.pop();
            }
        }

        for(uint256 i = proposalIndex; i < proposalLength - 1; i++){
            _proposalList[i] = _proposalList[i+1];
        }

        delete _proposalList[proposalLength-1];
        _proposalList.pop();

        delete _proposalInfo[proposalId];

        emit ProposalResult(
            proposalId,
            voteResult
        );
    }

    function getAdmin() public view returns (address, string memory){
        return (
            _admin.addr,
            _admin.serviceUrl
        );
    }

    function getAllAuthority() public view returns (address[] memory, string[] memory){
        uint256 authorityLength = _authorityList.length;
        address[] memory allAddress = new address[](authorityLength);
        string[] memory allUrl = new string[](authorityLength);
        for (uint256 i = 0; i < authorityLength; i++){
            allAddress[i] = _authorityList[i].addr;
            allUrl[i] = _authorityList[i].serviceUrl;
        }

        return (
            allAddress,
            allUrl
        );
    }

    function getAllProposal() public view returns (uint256[] memory){
        return _proposalList;
    }

    function getProposalInfo(uint256 proposalId)  public view returns (uint256,
        string memory,
        address,
        address,
        uint256,
        address[] memory){
        return (
            _proposalInfo[proposalId].proposalType,
            _proposalInfo[proposalId].publicInfoUrl,
            _proposalInfo[proposalId].senderAddress,
            _proposalInfo[proposalId].operatorAddress,
            _proposalInfo[proposalId].proposalBlockNumber,
            _proposalInfo[proposalId].voter
        );
    }
}