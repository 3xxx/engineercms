//绘制样条线
function DyDrawSpline() {
    //清空当前显示内容
    mxOcx.NewFile();
    var param = mxOcx.NewResbuf();
    mxOcx.SendStringToExecuteFun("_DrawSpline", param);
    var retparam = mxOcx.GetEntitysLastCmd();
    if (retparam == null)
        return;
    if (retparam.Count == 0)
        return;
    var spline = retparam.AtObject(0);
    if (spline == null) {
        return;
    }
    //把实体绘制到图上。一般用于动态拖放时，在拖放结束时，绘制到图上；返回新绘制的实体的id
    spline.Draw();
}
