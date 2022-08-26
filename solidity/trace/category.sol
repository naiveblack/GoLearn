pragma solidity^0.8.7;

import "./goods.sol";

// 商品种类合约
contract Category {
    // 记录商品的相关信息
    struct GoodsData {
        Goods traceData;
        bool isExists;
    }
    bytes32 currentCategory; // 记录商品的种类 可以用int或者hash
    // 商品id对应商品信息 goodsID => Goods
    mapping(uint256 => GoodsData) goods;
    uint8 constant STATUS_INVAILD = 255;

    // 新创建一个商品
    event NewGoods(address _operator, uint256 _goodsID);

    constructor(bytes32 _category) {
        currentCategory = _category;
    }

    function createGoods(uint256 _goodsID) public returns (Goods) {
        // 判断商品ID是否存在
        require(!goods[_goodsID].isExists, "goodsID already exists");
        goods[_goodsID].isExists = true;
        Goods _goods = new Goods(_goodsID);
        goods[_goodsID].traceData = _goods;

        emit NewGoods(msg.sender, _goodsID);
        return _goods;
    }

    function getStatus(uint256 _goodsID) public view returns (uint8) {
        if(!goods[_goodsID].isExists){
            return STATUS_INVAILD;
        }
        return goods[_goodsID].traceData.getStatus();
    }

    function changeStatus(uint256 _goodsID, uint8 _status, string memory _remark) public returns (bool) {

        return goods[_goodsID].traceData.changeStatus(_status, _remark);

    }

    function getGoods(uint256 _goodsID) public view returns (bool, Goods) {

        return (goods[_goodsID].isExists, goods[_goodsID].traceData);
    }

    function getTraceInfo(uint256 _goodsID) public view returns (Goods.TraceData[] memory) {
        return goods[_goodsID].traceData.getTraceInfo();
    } 

    // function getGoodsAddressInfo(uint256 _goodsID) public view returns (Goods) {
    //     return goods[_goodsID].traceData;
    // } 
}