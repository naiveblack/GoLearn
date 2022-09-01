// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "./IUser.sol";
import "./11-require.sol";

// 继承的时候需要导入相应的文件
// 导入后，需要实现对应的方法
// 继承合约的时候 直接把合约里面的方法都继承过来了
contract User is IUser, money_demo(msg.sender) {

    // 继承的时候需要加上override表示重载相应的方法
    function addUser(string memory _name, uint8 _age) override external {

    }

    function getUser(string memory _name) override external view returns (string memory, uint8) {

    }
}