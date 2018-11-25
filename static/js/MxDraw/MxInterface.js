
var isShow = false;
//隐藏/显示工具栏
function HideToolbar()
{
    mxOcx.ShowToolBar("常用工具",isShow);
    mxOcx.ShowToolBar("绘图工具",isShow);
    mxOcx.ShowToolBar("编辑工具",isShow);
    mxOcx.ShowToolBar("特性",isShow);
    mxOcx.ShowToolBar("ET工具",isShow);

    isShow = !isShow;
}

//隐藏/显示菜单栏
function HideMenuBar()
{

    mxOcx.ShowMenuBar(isShow);
    isShow = !isShow;
}


//隐藏/显示标尺栏
function HideRulerWindow()
{

    mxOcx.ShowRulerWindow(isShow);
    isShow = !isShow;
}

//隐藏/显示属性栏
function HidePropertyWindow()
{

    mxOcx.ShowPropertyWindow(isShow);
    isShow = !isShow;
}

//隐藏/显示命令栏
function HideCommandWindow()
{

    mxOcx.ShowCommandWindow(isShow);
    isShow = !isShow;
}

//隐藏/显示模型栏
function HideModelBar()
{

    mxOcx.ShowModelBar(isShow);
    isShow = !isShow;
}


//隐藏/显示状态栏
function HideStatusBar()
{

    mxOcx.ShowStatusBar(isShow);
    isShow = !isShow;
}






//模式切换
var isBrowner = false;
function BrownerMode() {
    isBrowner = !isBrowner;
    mxOcx.BrowseMode = isBrowner;
    mxOcx.ShowMenuBar = !isBrowner;
    mxOcx.ShowPropertyWindow = !isBrowner;
}
