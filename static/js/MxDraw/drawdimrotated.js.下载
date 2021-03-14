function DrawDimRotated (){
    //清空当前显示内容
    mxOcx.NewFile();
    mxOcx.AddDimStyle("MyDimStyle", "41,0.18,141,0.09,40,200", "77,1,271,3", "", "");
    mxOcx.DimStyle = "MyDimStyle";
    //把颜色改回黑白色
    mxOcx.DrawColor = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //绘制一个线型标注，参数一为输入第一条界线的起始点X值，参数二为输入第一条界线的起始点Y值
    //  参数三为输入第二条界线的起始点X值，参数四为输入第二条界线的起始点Y值
    //  参数五为输入标注线自身上的点X值，参数六为输入标注线自身上的点Y值
    //  参数七为输入标注的旋转角
    mxOcx.DrawDimRotated(0, 10, 5, 10, 2, 11, 0);
    /*mxOcx.DrawDimRotated(0, 72, 0, 74, 15, 73, 0.771);*/
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}
