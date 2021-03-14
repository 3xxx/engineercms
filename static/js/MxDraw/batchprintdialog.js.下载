function BatchPrintDialog() {
    var pRet =mxOcx.Call("Mx_GetMcDbDatabaseBound", "");
    if (pRet.AtString(0) == "Ok") {

        var pt = mxOcx.NewPoint();
        pt.x =  pRet.AtDouble(1);
        pt.y =  pRet.AtDouble(2);

        var param = mxOcx.NewResbuf();
        param.AddPoint(pt);

        pt.x =  pRet.AtDouble(3);
        pt.y =  pRet.AtDouble(4);
        param.AddPoint(pt);

        var print = mxOcx.NewComObject("IMxDrawPrint");

        print.BatchPrintDialog(param);
    }
    else {
        alert("调用失败");
    }

}