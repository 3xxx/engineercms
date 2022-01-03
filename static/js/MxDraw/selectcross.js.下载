function SelectCross()
{
    //实例化一个构造选择集进行过滤,该类封装了选择集及其处理函数。
    var ss = mxOcx.NewSelectionSet();
    //构造一个过滤链表
    var spFilte = mxOcx.NewResbuf();
    //定义两个点的位置
    //定义第一个点的位置
    var vPt1 = mxOcx.NewPoint();
    vPt1.x = -20;
    vPt1.y = -20;
    //定义第二个点的位置
    var vPt2 = mxOcx.NewPoint();
    vPt2.x = 20;
    vPt2.y = 20;
    //选取跨范围的实体
    ss.Select(1, vPt1, vPt2,spFilte);
    for (var i = 0; i < ss.Count; i++)
    {
        var ent = ss.Item(i);
        alert(ent.ObjectName);
    }
}
