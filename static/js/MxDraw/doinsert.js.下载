//动态拖放时的绘制事件
function DoInsert()
{
    //清空当前显示内容
    mxOcx.NewFile();
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message ="点取插入点";
    if(getPt.go() != 1)
    {
        return;
    }
    var frstPt =getPt.value();
    if(frstPt == null)
    {
        return;
    }
    var  getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    var spDrawData =getSecondPt.InitUserDraw("TestDynDraw");
    getSecondPt.message = "点取旋转角度";
    getSecondPt.basePoint = frstPt;
    getSecondPt.setUseBasePt(true);
    spDrawData.SetPoint("pt1",frstPt);
    var sBlkName = "Tree";
    //用户需要出入图块的路径   控件程序在磁盘的文件路径
    var sBlkFile = mxOcx.GetOcxAppPath() + "\\Blk\\树.dwg";
    //向控件数据库中插入一个图块,不用它插入匿名块
    mxOcx.InsertBlock(sBlkFile,"Tree");
    spDrawData.SetString("BlkName","Tree");
    if(getSecondPt.go() != 1)
    {
        return;
    }
    spDrawData.Draw();
}