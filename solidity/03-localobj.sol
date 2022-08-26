// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

// 计算随机数
contract localobj {
    address public admin;
    bytes32 public hash;
    uint256 public randnum;

    constructor() public {
        admin = msg.sender;     // msg.sender是调用者
        hash  = blockhash(0);   // 返回0块的hash值
        // 利用时间戳、调用者账号、hash值共同模拟随机值，得到100以内的数
        // 都是伪随机值，实现安全的随机数一般通过预言机来实现
        randnum = uint256(keccak256(abi.encode(block.timestamp, msg.sender, hash))) % 100;
    }
}