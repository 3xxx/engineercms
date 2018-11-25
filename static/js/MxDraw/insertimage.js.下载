function InsertImage() {
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
    //参数一为绘制位置,是图片的中心点X ；参数二为绘制位置,是图片的中心点Y；参数三为缩放比例；参数四为旋转角度；
    //参数五为图片显示文件名；参数六为闪烁文件设置；参数七为图片数据是否随图保存
    mxOcx.DrawImageMark(frstPt.x, frstPt.y, -100.0, 0.0, sImageFile, "", false);
}
