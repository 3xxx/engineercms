function GetXData(){
    var ent = mxOcx.GetEntity("选择要读扩展数据的实体:");
    if(ent == null)
    {
        return;
    }
    var exData = ent.GetXData("TestApp");

    if(exData.Count == 0)
    {
        alert("没有扩展数据");

    }
    else
    {
        // 向命令行窗口打印扩展数据。
        exData.PrintData();
        alert(exData.AtString(1));
    }
}
