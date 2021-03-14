function UserSaveJpg() {

    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");

    getPt.message = "点取范围第一点";

    if (getPt.go() != 1) {
        return;
    }
    var frstPt = getPt.value();
    if (frstPt == null)
        return;
    //新创建IMxDrawUtility对象
    var utl = mxOcx.NewUtility();
    //与用户交互得一个矩形框  参数一矩形框一个对角点；参数二命令行提示
    var secondPt = utl.GetCorner(frstPt, "点取范围第二点");
    if (secondPt == null)
        return;
    mxOcx.DrawToJpg("c:\\mxcad.jpg", frstPt.x , frstPt.y,  secondPt.x, secondPt.y, 200, 200, 0);
    alert("图片已经保存");
}
