//动态插入图块
function DyDynInsert() {
    //清空当前显示内容
    mxOcx.NewFile();
    var sBlkName = "Tree";
    //用户需要出入图块的路径   控件程序在磁盘的文件路径
    var sBlkFile = mxOcx.GetOcxAppPath() + "\\Blk\\树.dwg";
    //向控件数据库中插入一个图块,不用它插入匿名块
    mxOcx.InsertBlock(sBlkFile, "Tree");
    //新建一个COM组件对象
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message = "点取插入点";
    //初始动态绘制
    var spDrawData1 = getPt.InitUserDraw("TestDynDraw1");
    //设置string类型的属性
    spDrawData1.SetString("BlkName","Tree");
    if (getPt.go() != 1) {
        return;
    }
    var frstPt = getPt.value();
    if (frstPt == null) {
        return;
    }
    var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    var spDrawData = getSecondPt.InitUserDraw("TestDynDraw");
    getSecondPt.message = "点取旋转角度";
    getSecondPt.basePoint = frstPt;
    getSecondPt.setUseBasePt(true);
    spDrawData.SetPoint("pt1", frstPt);
    spDrawData.SetString("BlkName", "Tree");
    if (getSecondPt.go() != 1) {
        return;
    }
    var ret = spDrawData.Draw();
    // 给属性块写文字.
    var blkRef = ret.AtObject(0);
    if (blkRef != null)
    {
        var param = mxOcx.NewResbuf();
        blkRef.SetProp("createAttribute",param);
    }
}
