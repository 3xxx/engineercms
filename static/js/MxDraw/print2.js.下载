//按给定的中心点，缩放比例打印
function Print2() {
    //新建一个COM组件对象 参数为COM组件类名
    var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");

    getPt.message = "点取打印中心点";

    if (getPt.go() != 1) {
        return;
    }
    var frstPt = getPt.value();
    if (frstPt == null)
        return;

    //新建一个COM组件对象 参数为COM组件类名
    var print = mxOcx.NewComObject("IMxDrawPrint");
    //开始打印  参数一为打印的范围左下角X1；参数二为打印的范围左下角Y1
    //参数三为打印的范围右上角X2 ；参数四为打印的范围右上角Y2
    if (print.Print2(frstPt.x, frstPt.y,2)) {
        alert("打印成功");
    }
    else {
        alert("打印失败");
    }
}
