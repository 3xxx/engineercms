function DelText() {
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
    //返回数据库中的文字样式表表对象
    var TextStyleTable = database.GetTextStyleTable();
    //得到层表中的文字样式表对象
    var TextStyleRec = TextStyleTable.GetAt(txt, false);
    if (TextStyleRec == null)
        return;
    //删除对象
    TextStyleRec.Erase();
    alert("成功删除文字样式");
}