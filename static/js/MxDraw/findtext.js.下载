function FindText() {
    //对话框
    var rt;
    {
        var winWidth = 440;
        var winHeight = 140;
        var winLeft = (screen.width - winWidth) / 2;
        var winTop = (screen.height - winHeight) / 2 - 20;
        var str = 'dialogHeight:' + winHeight + 'px;dialogWidth:' + winWidth + 'px;dialogTop:' + winTop + 'px;dialogLeft:' + winLeft + 'px;resizable:yes;center:yes;status:no;'
        rt = window.showModalDialog("Gettext.htm?tmp=" + Math.random(), "输入文字", str);
    }
    //目标文字 - 对话框里的文字
    var txt;
    {
        if (typeof (rt) == "undefined") {
            return;
        } else {
            var arr = rt.split(",");
            txt = arr[0];
        }
    }
    var  ss;
    {
        //实例化一个构造选择集进行过滤,该类封装了选择集及其处理函数。
        ss = mxOcx.NewSelectionSet();
        //构造一个过滤链表
        var spFilte = mxOcx.NewResbuf();
        // 把文字对象，当着过滤条件.
        spFilte.AddStringEx("TEXT,MTEXT",5020);
        // 得到图上，所有文字对象.
        ss.Select2(5,null,null,null,spFilte);
    }
    // 遍历每个文字.
    var bFind = false;
    for(var i = 0; i < ss.Count;i++)
    {
        var ent = ss.Item(i);
        if(ent == null)
            continue;
        if (ent.ObjectName == "McDbText") {
            var sTxt = ent.TextString;
            if (sTxt == txt) {
                // 把文字放到视区中间.
                mxOcx.PutEntityInView(ent.ObjectID, 300);
                var dLen = mxOcx.ViewLongToDocCoord(80);
                // 绘制一个标记圆.
                mxOcx.DrawVectorCircle(ent.Position.x,
                    ent.Position.y,
                    dLen, 65280);
                bFind = true;
            }
        }
        else if (ent.ObjectName == "McDbMText") {
            var param = mxOcx.NewResbuf();
            param.AddObjectId(ent.ObjectID);
            var ret = mxOcx.CallEx("Mx_GetMTextContent", param);
            if (ret.AtString(0) == "Ok") {
                if (ret.AtString(1) == txt) {
                    // 把文字放到视区.
                    mxOcx.PutEntityInView(ent.ObjectID, 300);
                    var dLen = mxOcx.ViewLongToDocCoord(80);
                    // 绘制一个标记圆.
                    mxOcx.DrawVectorCircle(ent.Location.x,
                        ent.Location.y,
                        dLen, 65280);
                    bFind = true;
                    break;
                }
            }
        }
        ent = null;
    }

    if (!bFind) {
        alert("没有找到文字对象");
    }
    // 在这里必须显示释放控件的COM对象指针.
    ss = null;
    spFilte = null;
    CollectGarbage();
}
