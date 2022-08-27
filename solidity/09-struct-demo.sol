// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

// 0.6之前只能在contract内部定义结构体
struct User {
    string name;
    uint8 age;
    string sex;
}

contract struct_demo {
    User user;

    function setUser(string memory _name, uint8 _age, string memory _sex) public {
        user.name = _name;
        user.age  = _age;
        user.sex  = _sex;
    }

    function getUser() public view returns (User memory) {
        return user;
    }
}