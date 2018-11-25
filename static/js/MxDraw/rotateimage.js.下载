//旋转图片
function RotateImage() {
    // 新建一个COM组件对象
    var selEnt = mxOcx.NewComObject("IMxDrawUiPrEntity");
    selEnt.message = "选择图像对象";
    if(selEnt.go() != 1)
        return;
    var image = selEnt.Entity();
    if(image == null)
        return;
    if (image.ObjectName != "McDbMxImageMark") {
        alert("选择对象不是图像");
        return;
    }
    //新创建IMxDrawPoint对象
    var point = mxOcx.NewPoint();
    point.x = 0;
    point.y = 0;
    //旋转一个对象，参数一为旋转基点，参数二为旋转角度
    image.Rotate(point, 45 * 3.14159265 / 180.0);
    mxOcx.ZoomAll();

}
