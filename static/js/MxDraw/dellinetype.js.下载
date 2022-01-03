function DelLineType() {
    var winWidth = 440;
    var winHeight = 140;
    var winLeft = (screen.width - winWidth) / 2;
    var winTop = (screen.height - winHeight) / 2 - 20;
    var str = 'dialogHeight:' + winHeight + 'px;dialogWidth:' + winWidth + 'px;dialogTop:' + winTop + 'px;dialogLeft:' + winLeft + 'px;resizable:yes;center:yes;status:no;'
    var rt = window.showModalDialog("Gettext.htm?tmp=" + Math.random(), "输入图层名", str);
    var txt;
    if (typeof (rt) == "undefined") {
        return;
    } else {
        var arr = rt.split(",");
        txt = arr[0];
    }
    var database = mxOcx.GetDatabase();
    //返回数据库中的线型表对象
    var LinetypeStyleTable = database.GetLinetypeTable();
    //得到线型表中的线型对象
    var LinetypeStyleRec = LinetypeStyleTable.GetAt(txt, false);
    if (LinetypeStyleRec == null)
        return;
    //删除对象
    LinetypeStyleRec.Erase();
    alert("成功删除线型");
    mxOcx.Regen();
}
