// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

contract vardefine {

    // public 自动提供一个查看的方法
    uint8 public AuthAge;       // 作者年龄
    bytes32 public AuthHash;    // 作者哈希
    string public AuthName;     // 作者姓名
    uint256 AuthSal;            // 作者薪水

    // 构造函数
    constructor(uint8 _age, string memory _name, uint256 _sal) public {

        AuthAge  = _age;
        AuthName = _name;
        AuthSal  = _sal;

        // keccak256用来计算hash值，需要通过abi.encode进行转码
        // abi.encode可以传入多个不同类型的数据
        // keccak256返回值是bytes32的定长数组，也就是hash值
        AuthHash = keccak256(abi.encode(AuthAge, AuthName, AuthSal));
    }
}