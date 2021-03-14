function SelectUserSelect()
{
    //实例化一个构造选择集进行过滤,该类封装了选择集及其处理函数。
    var ss = mxOcx.NewSelectionSet();
    //构造一个过滤链表
    var spFilte = mxOcx.NewResbuf();
    //用户选择对象  得到用户选取的实体
    ss.Select2(8, null, null, null, spFilte);
    for (var i = 0; i < ss.Count; i++)
    {
        var ent = ss.Item(i);
        alert(ent.ObjectName);
    }
}
