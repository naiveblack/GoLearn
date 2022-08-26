pragma solidity^0.8.7;

import "./category.sol";

contract Factory {
    struct GoodsCategory {
        Category categoryData;
        bool isExists;
    }

    mapping(bytes32=>GoodsCategory) goodsCategory;

    event NewCategory(address _operator, bytes32 _category);

    function newGoods(bytes32 _category, uint256 _goodsID) public returns (Goods) {
        Category category = getCategory(_category);
        return category.createGoods(_goodsID);
    }

    function newCategory(bytes32 _category) public returns (Category) {

        require(!goodsCategory[_category].isExists, "category already exists");
        Category category = new Category(_category);    //新建了一个合约地址 category是个address
        goodsCategory[_category].isExists = true;
        goodsCategory[_category].categoryData = category;
        // 此时goodsCategory下面存储了一个address（新合约）和一个bool
        emit NewCategory(msg.sender, _category);
        return category;

    }

    // 获取当前产品类型下对应的产品合约的地址
    function getCategory(bytes32 _category) public view returns (Category) {
        return goodsCategory[_category].categoryData;
    }

    function getStatus(bytes32 _category, uint256 _goodsID) public view returns (uint8) {
        Category category = getCategory(_category);
        return category.getStatus(_goodsID);

    }

    function changeStatus(bytes32 _category, uint256 _goodsID, uint8 _status, string memory _remark) public returns (bool) {
        Category category = getCategory(_category);
        return category.changeStatus(_goodsID, _status, _remark);
    }

    function calCategory(string memory _name ) public pure returns (bytes32) {
        return keccak256(abi.encode(_name));
    }

    function getTraceInfo(bytes32 _category, uint256 _goodsID) public view returns (Goods.TraceData[] memory) {
        Category category = getCategory(_category);
        return category.getTraceInfo(_goodsID);
    } 

    // function getAllInfo(bytes32 _category) public view returns (GoodsCategory memory) {
    //     return goodsCategory[_category];    // 地址和bool
    //     /*返回值：
    //         tuple(address,bool): 0x5C9eb5D6a6C2c1B3EFc52255C0b356f116f6f66D,true
    //     */
    // }

    // // 得到商品的地址
    // function getGoodsAddressInfo(bytes32 _category, uint256 _goodsID) public view returns (Goods) {
    //     Category category = getCategory(_category);
    //    return category.getGoodsAddressInfo(_goodsID);
    // }
}