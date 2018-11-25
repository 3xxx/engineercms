function GetAllLayoutDictionary() {
    var database = mxOcx.GetDatabase();
    // 得到所有图层名
    var sRet = null;
    //返回数据库中的文字样式表对象
    var spLayoutDictionary = database.GetLayoutDictionary();
    //创建一个遍历层表中所有图层的遍历器
    var spIter = spLayoutDictionary.NewIterator();
    //移动当前遍历器位置
    for(;!spIter.Done();spIter.Step())
    {
        //返回遍历器当前位置的记录
        var spLayoutDictionaryRec = spIter.GetObject();
        //符号表记录名属性
        var sName = spLayoutDictionaryRec.LayoutName;
        // 0零层不参加比较
        if(sName != "0")
        {
            if(sRet == null)
                sRet = sName;
            else
            {
                sRet = sRet +","+sName;
            }
        }
    }
    alert(sRet);
}