function LockeAllLayer() {
    //返回控件的数据库对象
    var database = mxOcx.GetDatabase();
    //返回数据库中的层表对象
     var layerTable = database.GetLayerTable();
    //创建一个遍历层表中所有图层的遍历器
    var iter = layerTable.NewIterator();
    for (; !iter.Done(); iter.Step())
    {
        //返回遍历器当前位置的记录
        var layerRec = iter.GetRecord();
        layerRec.IsLocked = true;
        alert("成功锁定图层");

    }
}
