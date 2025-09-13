// SPDX-License-Identifier: MIT
pragma solidity ^0.8.20;
import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";
import "@openzeppelin/contracts/token/ERC721/extensions/ERC721URIStorage.sol";

contract vericredNFT is ERC721URIStorage, Ownable {
    address public owner;
    uint256 public _tokenIds;
    mapping (address => bool) public verifiedOrgs;
    address[] public Orgs;

    constructor() ERC721("document", "VCRD") Ownable(msg.sender) {
    }

    function mintDoc(address user, string memory tokenURI) public returns (uint256) {
        require(verifiedOrgs[msg.sender], "University is not verified/listed");
        
        _tokenIds += 1;
        uint256 newItemId = _tokenIds;
        _mint(user, newItemId);
        _setTokenURI(newItemId, tokenURI);

        return newItemId;
    }

    function currentTokenId() public view returns (uint256){
        return _tokenIds;
    }

    function allOrgs() public view returns (address[] memory) {
        return Orgs;
    }

    function isVerfiedOrg(address adr) public view returns (bool) {
        bool status = verifiedOrgs[adr];
        return status;
    }

    function newOrg(address _org) public {
        require(msg.sender == owner, "Only the owner can push new orgs.");
        require(_org != address(0), "Invalid organization address");
        require(!verifiedOrgs[_org], "Organization already registered"); // Add duplicate check
        
        verifiedOrgs[_org] = true;
        Orgs.push(_org);
    
        emit OrganizationAdded(_org); // Add event for transparency
    }
}