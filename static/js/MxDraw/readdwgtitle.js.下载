function ReadDwgTitle() {

    var database = mxOcx.GetDatabase();
    var param = mxOcx.NewResbuf();

    param.AddString("MyTitleVal");

    // 设置dwg文件说明信息标题为"MyTitleVal"
    database.SetProp("setTitle", param);

    // 取dwg文件说明信息标题.
    var ret = database.GetProp("getTitle");
    alert(ret.AtString(0));

}