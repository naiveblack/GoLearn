// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

import "./IUser.sol";

// 没有实现IUser中的全部方法时 需要使用abstract
// 同时 abstract因为是抽象的 所以无法部署
abstract contract absUser is IUser {

    // function addUser(string memory _name, uint8 _age) override external {

    // }

    function getUser(string memory _name) override external view returns (string memory, uint8) {

    }
}

contract User is absUser {
    function addUser(string memory _name, uint8 _age) override external {

    }
}