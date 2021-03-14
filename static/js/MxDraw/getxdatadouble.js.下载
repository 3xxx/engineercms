function GetxDataDouble() {
    var selEnt = mxOcx.NewComObject("IMxDrawUiPrEntity");

    selEnt.message = "选择要读取扩展数据的对象";
    if (selEnt.go() != 1)
        return;

    var ent = selEnt.Entity();
    if (ent == null)
        return;

    var val = ent.GetxDataDouble2("ExDataName", 0);
    if (mxOcx.IsOk()) {
        alert(val);
    }
    else {
        alert("没有扩展数!");
    }

}