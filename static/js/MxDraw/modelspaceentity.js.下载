function ModelSpaceEntity() {
    var database = mxOcx.GetDatabase();
    var blkRec = database.GetBlockTable().GetAt("*Model_Space");
    var iter = blkRec.NewIterator();

    if (iter == null)
        return;
    var iNum = 0;
    // 循环得到所有实体

    for (; !iter.Done(); iter.Step(true, false))
    {
        // 得到遍历器当前的实体
        iNum++;

    }
    alert(iNum.toString());

}