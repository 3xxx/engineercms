function CurrentSelect() {
    //实例化一个构造选择集进行过滤,该类封装了选择集及其处理函数。
    var ss = mxOcx.NewSelectionSet();
    //构造一个过滤链表
    var spFilte = mxOcx.NewResbuf();
    ss.CurrentSelect(spFilte);
    var i = 0;
    for (; i < ss.Count; i++) {
        var ent = ss.Item(i);
        if (ent.ObjectName == "McDbLine") {
            // 是个直线。
            var spt = ent.StartPoint;
            var ept = ent.EndPoint;
            alert(spt.x);
        }
        else if (ent.ObjectName == "McDbPolyline") {
            // 是个多义线.
            var j = 0;
            // 循环，得到端点坐标
            for (j = 0; j < ent.NumVerts; j++) {
                var pt = ent.GetPointAt(j);
                alert(pt.x);
            }
        }
    }
}