//设置隐藏图层
function HideLayer() {
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
    //返回数据库中的层表对象
    var layerTable = database.GetLayerTable();
    //得到层表中的图层对象
    var layerRec = layerTable.GetAt(txt, false);
    if (layerRec == null)
        return;
    //返回图层是否处于关闭状态
    layerRec.IsOff = true;
    alert("成功隐藏图层");
}
