//替换图片
function ModifyImage() {
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
    // 控件程序在磁盘的文件路径
    var sImageFile = mxOcx.GetOcxAppPath() + "\\mxdraw.png";
    //替换目标文件
    image.ImageFile = sImageFile;
}
