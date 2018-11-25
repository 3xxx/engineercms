function AllSelect() {
    //实例化一个构造选择集进行过滤,该类封装了选择集及其处理函数。
    var  ss = mxOcx.NewSelectionSet();
    //构造一个过滤链表
    var spFilte = mxOcx.NewResbuf();
    //得到当前空间的所有实体
    ss.AllSelect (spFilte);
    for (var i = 0; i < ss.Count; i++)
    {
        var ent = ss.Item(i);
        alert(ent.ObjectName);
    }
}
