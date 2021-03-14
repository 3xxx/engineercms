function DeleteXData() {

    var selEnt = mxOcx.NewComObject("IMxDrawUiPrEntity");
    selEnt.message = "选择要删除扩展数据的对象";
    if (selEnt.go() != 1)
        return;
    var ent = selEnt.Entity();
    if (ent == null)
        return;
    if (ent.DeleteXData("ExDataName")) {
        alert("删除扩展数据成功");
    }
    else {
        alert("删除扩展数据失败");
    }

}
