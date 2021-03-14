function DrawDimAligned(){
    //清空当前显示内容
    mxOcx.NewFile();
    //增加一个新的标注样式,如果当前已经有指定名的标注样式，就直接失败返回。
    mxOcx.AddDimStyle("MyDimStyle", "41,0.18,141,0.09,40,200", "77,1,271,3", "", "");
    mxOcx.DimStyle = "MyDimStyle";
    //把颜色改回黑白色
    mxOcx.DrawColor = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //绘制一个对齐标注
    //参数一为第一条界线开始点X值，参数二为 第一条界线开始点Y值
    //参数三为第二条界线开始点X值
    //参数四为第二条界线开始点Y值
    //参数五为标注使用的指定标注线位置的定义点X值
    //参数六为标注使用的指定标注线位置的定义点Y值
    mxOcx.DrawDimAligned(0, 40, 3, 45, 15, 46);
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}
