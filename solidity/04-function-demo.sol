// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

contract func_demo {

    function getSum() public pure returns (uint256 ) {
        uint256 sum = 0;
        for(uint i = 0; i <= 100; i ++) {
            sum += i;
        }
        return sum;
    }

    function getSum2() public pure returns (uint256 sum) {
        uint i = 0;
        while(i <= 50) {
            sum += i;
            i ++;
        }
        // return sum;
    }

    function isEqual(string memory _a, string memory _b) public pure returns(bool) {
        // solidity中对字符串没有进行 == 的操作
        // return _a == _b;
        
        // 整数类型和hash类型可以进行 == 操作
        // 将a b 求hash后 判断a b是否相等 (理论上存在hash碰撞，但是现实中很难做到)
        bytes32 hashA = keccak256(abi.encode(_a));
        bytes32 hashB = keccak256(abi.encode(_b));
        return hashA == hashB;
    }
}