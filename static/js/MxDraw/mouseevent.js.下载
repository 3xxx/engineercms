//控件中的鼠标事件
function MouseEvent(dX,dY,lType)
{
    //事件类型,1鼠标移动，2是鼠标左键按下，3是鼠标右键按下，
    // 4是鼠标左键双击 5是鼠标左键释放 6是鼠标右键释放
    // 7是鼠标中键按下 8是鼠标中键释放 9是鼠标中键双击 10是鼠标中键滚动
    if(lType == 4)
    {
        var ss = mxOcx.Call("Mx_NewSelectionSet","");
        var pt = mxOcx.Call("Mx_NewPoint","");
        var fil = mxOcx.Call("Mx_NewResbuf","");
        pt.x = dX;
        pt.y = dY;
        ss.SelectAtPoint(pt,fil);
        if(ss.Count > 0)
        {
            var ent = ss.Item(0);
            alert("点击了" + ent.ObjectName + "对象");
        }
        // 取消后面的命令。
        mxOcx.SendStringToExecute("");
        return 1;
    }
    return 0;
}
