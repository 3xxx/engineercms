function TextOffsetPosition(){


    var sImageFile1 = mxOcx.GetOcxAppPath() + "\\mxcad.jpg";
    var lId = mxOcx.DrawImageMark(0, 0, 1, 0, sImageFile1, "", true);
    //mxOcx.TwinkeEnt(lId);
    // 给一个很大的数据，可以代表把它放到最上面去。
    mxOcx.SetEntityDrawOrder(lId,999999);
    var imageMark = mxOcx.ObjectIdToObject(lId);

    var textOffset = mxOcx.NewPoint();
    textOffset.x = 0.0;
    textOffset.y = -70;

    var param = mxOcx.NewResbuf();


    param.AddPoint(textOffset);
    imageMark.SetProp("textOffsetPosition",param);

    imageMark.Text = "合格";
}