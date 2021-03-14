function CreateText(){
    //返回控件的数据库对象
    var database =mxOcx.GetDatabase();
    //返回数据库中的文字样式表对象
    var mxtst = database.GetTextStyleTable();
    //增加新的文字样式
    var mxtstr = mxtst.Add("duanceng");
    //设置样式
    //设置文字的TrueType字体,如果文字样式设置了TrueType字体，就不会使用前面设置shx字体
    //参数一为turetype字体名 参数二为是否粗写  参数三为是否斜写  参数四为windows字符集  参数五暂没使用
    mxtstr.SetFont("黑体", false, false, 0, 0);
    //返回文字的高度
    mxtstr.textSize = 5;

    //增加新的文字样式
    var mxtstr = mxtst.Add("duanceng1");
    //设置样式
    //设置文字的TrueType字体,如果文字样式设置了TrueType字体，就不会使用前面设置shx字体
    //参数一为turetype字体名 参数二为是否粗写  参数三为是否斜写  参数四为windows字符集  参数五暂没使用
    mxtstr.SetFont("黑体", false, false, 0, 0);
    //返回文字的高度
    mxtstr.textSize = 9;
    //返回数据库中的文字样式表对象
    //设置为当前文字样式
    database.CurrentlyTextStyle = "duanceng";
}