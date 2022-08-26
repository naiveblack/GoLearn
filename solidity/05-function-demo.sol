// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

contract func_demo2 {
    uint256 count;

    constructor(uint256 _c) {
        count = _c;
    }

    function setCount(uint256 _c) public {
        count = _c;
    }

    function getCount() external view returns (uint256){
        return count;
    }

    function deposit() external payable {

    }

    function isEqual(string memory _a, string memory _b) internal pure returns(bool) {
        bytes32 hashA = keccak256(abi.encode(_a));
        bytes32 hashB = keccak256(abi.encode(_b));
        return hashA == hashB;
    }
}