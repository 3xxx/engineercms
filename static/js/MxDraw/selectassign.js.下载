function SelectAssign() {
    //实例化一个构造选择集进行过滤,该类封装了选择集及其处理函数。
    var ss = mxOcx.NewSelectionSet();
    //构造一个过滤链表
    var spFilte = mxOcx.NewResbuf();
    // "0"，是图层名，8是DXF码，代表"0"是个图层字符串。
    // 选择 0层上的实体。
    spFilte.AddStringEx("0", 8);
    // 选择直线或圆弧实体。5020是DXF码，代表"LINE"是类型字符串。
    spFilte.AddStringEx("LINE,ARC", 5020);
    // 选择红色的实体。1为颜色索引值，代表红色。 62是DXF码，但1是个颜色类型。
    spFilte.AddLongEx(1,62);
    ss.Select2(5,null, null, null,spFilte);
    for (var i = 0; i < ss.Count; i++)
    {
        var ent = ss.Item(i);
        alert(ent.ObjectName);
    }
}
