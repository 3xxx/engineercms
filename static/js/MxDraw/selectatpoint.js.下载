function SelectAtPoint()
{
    //实例化一个构造选择集进行过滤,该类封装了选择集及其处理函数。
    var ss = mxOcx.NewSelectionSet();
    //构造一个过滤链表
    var spFilte = mxOcx.NewResbuf();
    //在一个点构造选择集
    var vPt = mxOcx.NewPoint();
    vPt.x = -20;
    vPt.y = -20;
    ss.SelectAtPoint (vPt, spFilte);
    for (var i = 0; i < ss.Count; i++)
    {
        var ent = ss.Item(i);
        alert(ent.ObjectName);
    }
}