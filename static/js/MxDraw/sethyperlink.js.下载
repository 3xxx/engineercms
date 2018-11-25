function DoInputPointToolTipFun(ent) {
    var sHyperlinks = ent.Hyperlinks;
    if(sHyperlinks.length != 0)
    {
        var sClassName = ent.ObjectName;

        var tip = "<b><ct=0x0000FF><al_c>"+sClassName+
            "</b><br><ct=0x00AA00><hr=100%></ct><br><a=\"link\">" + sHyperlinks + "</a>";
        //设置用户自定义事件的结果返回值
        mxOcx.SetEventRetString(tip);
    }
    ent = null;
    CollectGarbage();
}
//设置超连接实体被点击事件回调函数
function DoHyperlinkClickFun(ent, dX, dY) {
    //打开一个网址连接
    mxOcx.GotoURL(ent.Hyperlinks);
    //设置自定义事件的返回值
    mxOcx.SetEventRet(1);
}
function SetHyperlink() {
    //新建一个COM组件对象
    var selEnt = mxOcx.NewComObject("IMxDrawUiPrEntity");
    selEnt.message = "选择要设置的对象";
    if (selEnt.go() != 1)
        return;
    //返回用户选择的实体
    var ent = selEnt.Entity();
    if (ent == null)
        return;
    //设置超链接(用户可以设置所需链接)
    ent.Hyperlinks = "www.mxdraw.com";
    //设置动态提示弹出时间,默认为1000.0毫秒
    mxOcx.DynToolTipTime = 100;
}