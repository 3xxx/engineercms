function DelDim() {
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
    //返回数据库中的样式表对象
    var DimStyleTable = database.GetDimStyleTable();
    //得到样式表表中的样式对象
    var DimStyleRec = DimStyleTable.GetAt(txt, false);
    if (DimStyleRec == null)
        return;
    //删除对象
    DimStyleRec.Erase();
    alert("成功删除标注样式");
}