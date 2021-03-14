function DrawDimDiametric(){
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
    //绘制一个直径标注
    //参数一为在被标注的曲线上的第一个点X值 ，参数二为在被标注的曲线上的第一个点Y值
    //参数三为在被标注的曲线上的第二个点X值，这个点与chordPoint直径相对
    //参数四为在被标注的曲线上的第二个点Y值，这个点与chordPoint直径相对
    //参数五为输入箭头长度
    mxOcx.DrawDimDiametric (50, 570, 60, 600, 100);
    //把所有的实体都放到当前显示视区
    mxOcx.ZoomAll();
    //更新视区显示
    mxOcx.UpdateDisplay();
}
