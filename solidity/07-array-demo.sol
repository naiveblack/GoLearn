// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

contract array_demo {

    string[5] public names;
    uint8[] public ages;

    constructor(){
        names[0] = "zhangsan";
        names[1] = "lisi";
        names[2] = "wangwu";

        // 动态数组 没有初始化 直接访问会发生越界问题 应该使用push添加数据
        // ages[0] = 20;
        ages.push(20);
    }

    function addAge(uint8 _age) public {
        ages.push(_age);
    }

    function getLength() public view returns (uint256, uint256) {
        return (names.length, ages.length);
    }

    function addName(string memory _name) public {
        // 对定长数组不能使用push操作
        // names.push(_name);
    }
    /*
        很少用delete去删除元素
        删除后元素的空间不会被释放
    */
}