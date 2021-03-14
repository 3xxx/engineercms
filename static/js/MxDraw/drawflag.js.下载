function DrawFlag() {
    //清空当前显示内容
    mxOcx.NewFile();
    //循环
    while(true)
    {
        //新建一个COM组件对象 参数为COM组件类名
        var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
        getPt.message = "点取绘制点";
        if(getPt.go() != 1)
        {
            return;
        }
        var frstPt = getPt.value();
        if(frstPt == null)
            return;
        //返回拾取矩形框的宽度,默认值为6
        var dLen = mxOcx.GetCursorPickRect();
        //视区长度到文档长度变换
        dLen = mxOcx.ViewLongToDocCoord(dLen);
        dLen *= 3.0;
        //绘制失量线
        //参数一为开始点X值；参数二为开始点Y值；参数三为结束点X值；参数四为结束点Y值；参数五为颜色

        mxOcx.DrawVectorLine(frstPt.x - dLen,frstPt.y - dLen,
            frstPt.x + dLen,frstPt.y + dLen,
            255
        );
        mxOcx.DrawVectorLine(frstPt.x - dLen,frstPt.y + dLen,
            frstPt.x + dLen,frstPt.y - dLen,
            255
        );
        //绘制失量圆
        //参数一为失量圆中心点X,文档坐标 ；参数二为失量圆中心点Y,文档坐标；参数三为失量圆半径,文档坐标；参数四为颜色
        mxOcx.DrawVectorCircle(frstPt.x,frstPt.y,
            dLen * 0.5, 65280);
        //更新当前控件的显示
        mxOcx.UpdateDisplay();
    }
}
