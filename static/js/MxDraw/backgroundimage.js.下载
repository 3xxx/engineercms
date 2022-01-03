function BackGroundImage() {
    //清空当前显示内容
    mxOcx.NewFile();
    // 新建一个COM组件对象
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
    getPt.message = "点取图片的插入中点";
    if (getPt.go() != 1) {
        return;
    }
    var frstPt = getPt.value();
    if (frstPt == null) {
        return;
    }
    // 控件程序在磁盘的文件路径
    var sImageFile = mxOcx.GetOcxAppPath() + "\\mxcad.jpg";
    // 绘图制一个图象标记对象
    //参数一为光栅图文件.DWGBMPJPG文件路径；参数二为栅图显示的在控件视区显示的左上角位置X，视区窗口坐标 ；
    // 参数三为栅图显示的在控件视区显示的左上角位置Y，视区窗口坐标 ；
    mxOcx.DrawImageToBackground(sImageFile,frstPt.x, frstPt.y );
}
