//利用图片制作动画
function  DrawGif()
{
    //新建一个COM组件对象 参数为COM组件类名
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message = "点取图片的插入中点";
    if (getPt.go() != 1) {
        return;
    }
    var frstPt = getPt.value();
    if (frstPt == null) {
        return;
    }
    //控件程序在磁盘的文件路径
    var sImageFile1 = mxOcx.GetOcxAppPath() + "\\1.png";
    var sImageFile2 = mxOcx.GetOcxAppPath() + "\\2.png";
    var sImageFile3 = mxOcx.GetOcxAppPath() + "\\3.png";
    // 绘图制一个图象标记对象
    //参数一为绘制位置,是图片的中心点X ；参数二为绘制位置,是图片的中心点Y；参数三为缩放比例；参数四为旋转角度；
    //参数五为图片显示文件名；参数六为闪烁文件设置；参数七为图片数据是否随图保存
    var lId = mxOcx.DrawImageMark(frstPt.x, frstPt.y, -20, 0, sImageFile1,
        sImageFile1 + "," + sImageFile2 + "," + sImageFile3, true);
    //闪烁实体
    //参数一为被闪烁的实体对象id；参数二为闪烁次数，默认为-1表示不限闪烁次数，成功返回true
    mxOcx.TwinkeEnt(lId);
}
