function DrawDimAngular(){
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
    //绘制一个角度标注
    //参数一为角度标注的顶点的X值，参数二为角度标注的顶点的Y值
    //参数三为角度标注第一个参考线上的点X值，参数四为角度标注第一个参考线上的点Y值
    //参数五为角度标注第二个参考线上的点X值
    //参数六为角度标注第二个参考线上的点Y值
    //参数七为文字位置点X值，参数八为文字位置点Y值
    mxOcx.DrawDimAngular(500, 5000, 0, 5500, 1000, 5500, 500, 5500);
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}