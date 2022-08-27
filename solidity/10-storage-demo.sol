// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

// storage 引用传递 修改的是值的本身
// memory  值传递   修改的是值的副本
contract storage_demo {

    struct User {
        string name;
        uint8 age;
        string sex;
    }
    
     User adminuser;

    function setUser(string memory _name, uint8 _age, string memory _sex) public {
        adminuser.name = _name;
        adminuser.age  = _age;
        adminuser.sex  = _sex;
    }

    function getUser() public view returns (User memory) {
        return adminuser;
    }

    // 无法修改数据
    function setAge1(uint8 _age) public {
        User memory user = adminuser;
        user.age = _age;
    }

    // 可以修改数据
    function setAge2(uint8 _age) public {
        User storage user = adminuser;
        user.age = _age;
    }

    // storage 不能在public里面使用 只能在内部函数中使用 public会报错
    function setAge3(User storage _user, uint8 _age) internal {
        _user.age = _age;
    }

    // User storage 可以修改成功 改为memory后 修改不成功
    function callsetAge3(uint8 _age) public {
        setAge3(adminuser, _age);
    }
}