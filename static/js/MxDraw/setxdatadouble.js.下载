function SetxDataDouble(){
    var selEnt = mxOcx.NewComObject("IMxDrawUiPrEntity");
    selEnt.message = "选择要写扩展数据的对象";
    if (selEnt.go() != 1)
        return;
    var ent = selEnt.Entity();
    if (ent == null)
        return;
    if (ent.SetxDataDouble("ExDataName", 0, 55.5)) {
        alert("写扩展数据成功");
    }
    else {
        alert("写扩展数据失败");
    }
}