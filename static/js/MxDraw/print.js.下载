//选择范围打印
function Print() {
    //新建一个COM组件对象 参数为COM组件类名
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");

    getPt.message = "点取打印范围第一点";

    if (getPt.go() != 1) {
        return;
    }
    var frstPt = getPt.value();
    if (frstPt == null)
        return;
    //新创建IMxDrawUtility对象
    var utl = mxOcx.NewUtility();
    //与用户交互得一个矩形框  参数一矩形框一个对角点；参数二命令行提示
    var secondPt = utl.GetCorner(frstPt, "点取打印范围第二点");
    if (secondPt == null)
        return;
    //新建一个COM组件对象 参数为COM组件类名
    var print = mxOcx.NewComObject("IMxDrawPrint");
    //开始打印  参数一为打印的范围左下角X1；参数二为打印的范围左下角Y1
    //参数三为打印的范围右上角X2 ；参数四为打印的范围右上角Y2
    if (print.Print(frstPt.x, frstPt.y, secondPt.x, secondPt.y)) {
        alert("打印成功");
    }
    else {
        alert("打印失败");
    }
}
