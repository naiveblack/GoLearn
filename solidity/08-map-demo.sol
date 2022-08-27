// SPDX-License-Identifier: Apache-2.0
pragma solidity ^0.8.0;

contract map_demo {

    // 定义学生姓名和成绩的mapping
    mapping(string=>uint256) maths;

    constructor() {
        maths["qwq"] = 97;       
    }

    function addScore(string memory _name, uint256 _score) public {
        if (maths[_name] > 0) return; // 表示当前的 key 已经存在了
        maths[_name] = _score;
    }

    function getScore(string memory _name) public view returns(uint256){
        return maths[_name];
    }
}