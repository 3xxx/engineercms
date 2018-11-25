//打开全部图层
function OpenAllLayer() {
    var database = mxOcx.GetDatabase();
    //返回数据库中的层表对象
    var spLayerTable = database.GetLayerTable();
    //创建一个遍历层表中所有图层的遍历器
    var spIter = spLayerTable.NewIterator();
    //移动当前遍历器位置
    for(;!spIter.Done();spIter.Step(true,true))
    {
        //返回遍历器当前位置的记录
        var spLayerRec = spIter.GetRecord();
        var sName = spLayerRec.Name;
        var layerRec=spLayerTable.GetAt(sName,false);
        if(layerRec)
        {
            layerRec.IsOff = false;
            alert("成功显示图层");
        }

    }

}