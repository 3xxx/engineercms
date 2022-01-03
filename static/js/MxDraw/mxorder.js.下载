function MxOrder() {

    var database = mxOcx.GetDatabase();
    var ret = database.GetProp("getMinDrawOrder");
    var lMin = ret.AtLong(0);

    var ret2 = database.GetProp("getMaxDrawOrder");
    var lMax = ret2.AtLong(0);

    var sTemp = mxOcx.DrawImage(0, 0, 0, 0, "D:\\mxcad.jpg");
    var ent = database.ObjectIdToObject(sTemp);

    var res = mxOcx.NewResbuf();
    res.AddLong(lMin + 2);

    ent.SetProp("drawOrder", res);
    mxOcx.ZoomAll();
}
