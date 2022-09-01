// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

interface IUser {

    // 定义接口的时候 不应该使用public 而是使用external
    function addUser(string memory _name, uint8 _age) external ;
    function getUser(string memory _name) external view returns (string memory, uint8);
    
}