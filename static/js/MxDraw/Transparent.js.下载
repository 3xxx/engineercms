function Transparent() {

    var sImageFile1 = mxOcx.GetOcxAppPath() + "\\mxcad.jpg";
    var lId = mxOcx.DrawImageMark(0, 0, 1, 0, sImageFile1, "", true);
    // 设置图片的透明度

    var imageMark=mxOcx.ObjectIdToObject(lId);
    var param = mxOcx.NewResbuf();

    param.AddLong(50);
    imageMark.SetProp("setTransparent",param);
    mxOcx.ZoomAll();
    mxOcx.UpdateDisplay();

}