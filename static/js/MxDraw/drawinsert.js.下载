function DrawInsert() {
    //清空当前显示内容
    mxOcx.NewFile();
    mxOcx.DrawColor = 0;
    //创建一个图层,名为"BlockLayer"
    mxOcx.AddLayer ("BlockLayer");
    //设置当前图层为"BlockLayer"
    mxOcx.LayerName = "BlockLayer";
    var sBlkName = "attribblock";
    //用户需要出入图块的路径   控件程序在磁盘的文件路径
    var sBlkFile = mxOcx.GetOcxAppPath() + "\\Blk\\attribblock.dwg";
    //向控件数据库中插入一个图块,不用它插入匿名块
    mxOcx.InsertBlock(sBlkFile, "attribblock");
    //绘制块引用对象
    //参数一插入点的X坐标 ，参数二插入点的Y坐标
    //参数三图块名，参数四图块缩放比例，参数五图块的旋转角度
    mxOcx.DrawBlockReference(0, 3600, "attribblock", 1, 0) ;
    mxOcx.ZoomAll();
    mxOcx.UpdateDisplay();
}
