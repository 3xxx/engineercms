function DrawDimRadial(){
    //清空当前显示内容
    mxOcx.NewFile();
    mxOcx.AddDimStyle("MyDimStyle", "41,0.18,141,0.09,40,20", "77,1,271,2", "", "");
    mxOcx.DimStyle = "MyDimStyle";
    //把颜色改回黑白色
    mxOcx.DrawColor = 0;
    //把线型改成实线
    mxOcx.LineType = "";
    //设置线宽 4
    mxOcx.LineWidth = 0;
    //绘制一个半径标注
    //参数一为被标注的曲线的中点X值 ，参数二为被标注的曲线的中点Y值
    //参数三为被标注的曲线上的点X值，参数四为被标注的曲线上的点Y值
    //参数五为输入箭头长度
    mxOcx.DrawDimRadial (10, 50, 10, 55, 0);
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}
