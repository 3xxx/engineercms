function GetAllAppName() {

    var selEnt = mxOcx.NewComObject("IMxDrawUiPrEntity");
    selEnt.message = "选择需要读取的对象";
    if (selEnt.go() != 1)
        return;
    var ent = selEnt.Entity();
    if (ent == null)
        return;
    //得到所有扩展数据名称
    var val = ent.GetAllAppName();
    if (val == null) {
        return;
    }
    for(var i = 0;i < val.Count; i++ ){
        //返回链表中的指定位置字符串变量值
        alert(val.AtString(i));
    }

}
