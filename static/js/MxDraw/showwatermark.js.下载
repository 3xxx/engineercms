//设置水印
//首先定义iShowWatermark为false
var iShowWatermark = false;
function ShowWatermark() {
    //若是有水印，点击按钮水印不显示，反之，则反
    iShowWatermark = !iShowWatermark;
    if (iShowWatermark) {
        //设置控件水印图片显示,字符串用逗号隔开，分为: “文件名,透明度，x方向距离，
        // y方向距离，是否居中”, 是否居中0代表在上角定位，1表示自动居中.默认为居中，透明度为50.
        // GetOcxAppPath为控件程序在磁盘的文件路径
        mxOcx.Watermark = mxOcx.GetOcxAppPath() + "\\Watermark.png" + ",20,5,5,1";
        // 背景色改成白色（用户可以根据需要改变背景色）
        mxOcx.ViewColor = 16777215;
    }
    else {
        mxOcx.Watermark = "";
        // 背景色改成白色（用户可以根据需要改变背景色）
        mxOcx.ViewColor = 0;
    }
}