<!-- 这个是使用mxdraw控件 编辑 dwg文件 保存到服务器 -->
<!-- <!DOCTYPE html> 这个会导致高度很小-->
<head>
<meta http-equiv="Content-Type" content="text/html; charset=gb2312">
<meta name="GENERATOR" content="Microsoft FrontPage 4.0">
<meta name="ProgId" content="FrontPage.Editor.Document">
<title>{{.FileName}}</title>
</head>

<SCRIPT language="JavaScript">
  // var mxOcx = document.getElementById("MxDrawXCtrl");
  document.oncontextmenu = new Function('event.returnValue=false;'); //禁用右键
</SCRIPT>
  <!-- onload = "PageInit();" -->
<body topmargin="10" leftmargin="10">

<p>
  <input type="button" value="浏览模式切换" onclick="DoCmd(6)">
  <input type="button" value="保存DWG文件到服务器" onclick="DoCmd(16)">
	<!-- <input type="button" value="动态插入图块" onclick="DoDynInsert()">
	<input type="button" value="绘制自定义实体" onclick="DoCustomEntity()">
	<input type="button" value="插入图片" onclick="DoCmd(3)">
	<input type="button" value="绘制动画" onclick="DoCmd(4)"> -->
</p>

<p align="center">
	<object classid="clsid:74A777F8-7A8F-4e7c-AF47-7074828086E2" id="MxDrawXCtrl"  width=100% height=95% align="left"> 
	<param name="_Version" value="65536">
	<param name="_ExtentX" value="24262">
	<param name="_ExtentY" value="16219">
	<param name="_StockProps" value="0">
	<param name="DwgFilePath" value="">
	<param name="IsRuningAtIE" value="1">
	<param name="EnablePrintCmd" value="0">
	<param name="ShowCommandWindow" value="1">
	<param name="ShowToolBars" value="1">
	<param name="ShowModelBar" value="1">
	<param name="IniFilePath" value="">
	<param name="ToolBarFiles" value="">
	<param name="ShowMenuBar" value="1">
	<param name="EnableUndo" value="1">
	<!-- <param name="Iniset" value=""> -->
	<param name="ShowPropertyWindow" value="1">
  
</object>

<script>
    function savedwgurl(){
      //     mxOcx.Cal("Mx_ShowWeight");
      // var param = mxOcx.NewResbuf();
      // mxOcx.SendStringToExecuteFun("_DrawSpline", param);
      // var retparam = mxOcx.GetEntitysLastCmd();
      // if (retparam == null){
      //     return;
      // }
      // if (retparam.Count == 0){
      //     return;
      // }
      // var spline = retparam.AtObject(0);
      // if (spline == null) {
      //     return;
      // }
      // spline.Lineweight = 40;
      //var dwgName = $("#dwgList option:selected").text();
      // alert("保存DWG文件到服务器");
      //         var hostName = window.location.hostname;
      //         var port = window.location.port;
      //         var path = "http://" + hostName + ":" + port;
      //         alert(hostName+"==="+path);
      //         if (!mxOcx.SaveDwgToURL(path, "/savedwgfile?key=" + Math.random() + "&flag=" + flag, "dwgName", ""))
      //         {
      //             alert(path);
      //             var ret = mxOcx.Call("Mx_GetLastError", "");
      //             alert(ret.AtString(1));
      //         }
      //         else
      //         {
      //             alert("保存成功");
      //         }
      //  MxDrawXCtrl_Obj = document.all.item("MxDrawXCtrl");
      //  MxDrawXCtrl_Obj.focus();
      //  // MxDrawXCtrl_Obj.SaveDwgToURL("http://www.192.168.1.102", "/Save.aspx", "test.dwg", "80");
        // // MxDrawXCtrl_Obj.SaveDwgToURL("", "/Save.aspx", "test.dwg", "80");
        // // return;
        //  /*
        //  if (!MxDrawXCtrl_Obj.SaveDwgToURL(getcurpath(), "/Save.aspx", "test.dwg", "")) {
        //      var ret = MxDrawXCtrl_Obj.Call("Mx_GetLastError","");
        //      alert(ret.AtString(1));
        //  }
        //  else {
        //      alert("成功");
        //  }  
        //  */
        //  var param = MxDrawXCtrl_Obj.Call("Mx_NewResbuf", "");
        //  // param.AddString(getcurpath();             //  服务器网址地址，如：www.mxdraw.com
        //  // param.AddString("http://localhost");             //  服务器网址地址，如：www.mxdraw.com
        //  // param.AddString("https://cnblogs.com");     // 支持https,端口要传成443
        //  param.AddString("");             //  服务器网址地址，如：www.mxdraw.com
        //  param.AddString("/Save.aspx");             //  服务器的文件上传处理程序,如:upload.asp
        // // param.AddString("FileComponentName");      //  HTML组件名称
        //  param.AddString("testsave.dwg");      //  HTML组件名称
        //  //param.AddString("80");                     //  服务处理端口. 如：_T("80")
        //  //param.AddString("443"); 
        //  param.AddString("6046");                     //  服务处理端口. 如：_T("80")
        //  param.AddString("testsave.dwg");               //  文件标志名称值,表单提交事件中filename值
        //  var ret = MxDrawXCtrl_Obj.CallEx("Mx_SaveDwgToURLEx", param);
        //  if (ret.AtString(0) == "Ok") {
        //      alert("成功");
        //  }
        //  else {
        //      alert(ret.AtString(1));
        //  }  
    }

  var mxOcx = document.getElementById("MxDrawXCtrl");
  function DoCustomEntity(){
    document.getElementById("MxDrawXCtrl").DoCommand(1);
  }

function DoDynInsert() {
  document.getElementById("MxDrawXCtrl").DoCommand(2);
}

function DoCmd(iCmd) {
  mxOcx.DoCommand(iCmd);
}

function DoCommandEventFunc(iCmd){
  if (iCmd == 1) 
  {
   InsertCustomEntity();
  }
  else if(iCmd == 2) {
  DynInsert();
  }
  else if(iCmd == 3) {
  InsertImage();
  }
  else if (iCmd == 4) {
      DrawGif();
  }
  else if (iCmd == 6) {
      BrownerMode();
  }
  else if (iCmd == 16) {
      DrawSpline();
  }
}

function DrawSpline() {
  // alert("保存DWG文件到服务器");
  // var hostName = window.location.hostname;
  // var port = window.location.port;
  // var path = "http://" + hostName + ":" + port;
  // alert(hostName+"==="+path);
  if (!mxOcx.SaveDwgToURL(path, "/project/product/savedwgfile?id={{.Id}}","file","")){
    // alert(path);
      var ret = mxOcx.Call("Mx_GetLastError", "");
      alert(ret.AtString(1));
  }else{
      alert("保存成功");
  }
}

var isBrowner = false;
function BrownerMode() {
  isBrowner = !isBrowner;
  mxOcx.BrowseMode = isBrowner;
  mxOcx.ShowMenuBar = !isBrowner;
  mxOcx.ShowPropertyWindow = !isBrowner;
}

function DoDynWorldDrawFun(dX,dY,pWorldDraw,pData){
  var sGuid = pData.Guid;
  mxOcx.SetEventRet(0);
  if(sGuid == "TestDynDraw")
  {
    // 动态插入图块。
    var firstPt =  pData.GetPoint("pt1");
    if(firstPt == null)
        return;
    var sBlkName =  pData.GetString("BlkName");
    var secondPt = mxOcx.NewPoint();
    secondPt.x = dX;
    secondPt.y = dY;
    var vec = firstPt.SumVector(secondPt);
    var dAng = vec.Angle();
    pWorldDraw.DrawBlockReference(firstPt.x,firstPt.y,sBlkName,1.0,dAng * 180.0 / 3.14159265 + 90.0);
      mxOcx.SetEventRet(1);
    }
}

function InsertCustomEntity() {
  var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
  getPt.message = "点取第一点";
  if(getPt.go() != 1)
      return;
  var frstPt = getPt.value();
  if(frstPt == null)
      return;
  var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
  getSecondPt.message = "点取第二点";
  getSecondPt.basePoint = frstPt;
  getSecondPt.setUseBasePt(true);
  if(getSecondPt.go() != 1)
      return;
  var secondPt = getSecondPt.value();
  if(secondPt == null)
      return;
  var ent = mxOcx.DrawCustomEntity("TestMxCustomEntity","");
  ent.SetPoint("spt",frstPt);
  ent.SetPoint("ept", secondPt);
}

function ExplodeFun(pCustomEntity, pWorldDraw) {
  var sGuid = pCustomEntity.Guid;
  if(sGuid == "TestMxCustomEntity"){
    if (!pCustomEntity.IsHave("ept"))
        return;
    var stp =  pCustomEntity.GetPoint("spt");
    if(stp == null)
        return;
    var ept =  pCustomEntity.GetPoint("ept");
    if(ept == null)
        return;
    pWorldDraw.DrawLine(stp.x, stp.y, ept.x, ept.y);
    mxOcx.SetEventRet(1);
  }
}

function GetGripPointsFun(pCustomEntity) {
  var sGuid = pCustomEntity.Guid;
  if(sGuid == "TestMxCustomEntity")
  {
    if(!pCustomEntity.IsHave("ept") )
        return;
    var stp =  pCustomEntity.GetPoint("spt");
    if(stp == null)
        return;
    var ept =  pCustomEntity.GetPoint("ept");
    if(ept == null)
        return;
    var ret = mxOcx.NewResbuf();
    ret.AddPoint(stp);
    ret.AddPoint(ept);
    mxOcx.SetEventRetEx(ret);
  }
}

function MoveGripPointsFun(pCustomEntity, lGridIndex, dOffsetX, dOffsetY) {
  var sGuid = pCustomEntity.Guid;
  if (sGuid == "TestMxCustomEntity") {
    if (!pCustomEntity.IsHave("ept"))
        return;
    var stp = pCustomEntity.GetPoint("spt");
    if (stp == null)
        return;
    var ept = pCustomEntity.GetPoint("ept");
    if (ept == null)
        return;
    if(lGridIndex == 0)
    {
        stp.x=stp.x + dOffsetX;
        stp.y=stp.y + dOffsetY;
        pCustomEntity.SetPoint("spt",stp);
    }
    else
    {
        ept.x=ept.x + dOffsetX;
        ept.y=ept.y + dOffsetY;
        pCustomEntity.SetPoint("ept",ept);
    }
      mxOcx.SetEventRet(1);
    }
}

function TransformByFun(pCustomEntity, pMatXform) {
  var sGuid = pCustomEntity.Guid;
  if (sGuid == "TestMxCustomEntity") {
    if (!pCustomEntity.IsHave("ept"))
        return;
    var stp = pCustomEntity.GetPoint("spt");
    if (stp == null)
        return;
    var ept = pCustomEntity.GetPoint("ept");
    if (ept == null)
        return;
    stp.TransformBy(pMatXform);
    ept.TransformBy(pMatXform);
    pCustomEntity.SetPoint("spt", stp);
    pCustomEntity.SetPoint("ept", ept);
    mxOcx.SetEventRet(1);
  }
}

function GetGeomExtentsFun(pCustomEntity) { 
  var sGuid = pCustomEntity.Guid;
  if (sGuid == "TestMxCustomEntity") 
  {
    if (!pCustomEntity.IsHave("ept"))
        return;
    var stp = pCustomEntity.GetPoint("spt");
    if (stp == null)
        return;
    var ept = pCustomEntity.GetPoint("ept");
    if (ept == null)
        return;
    var ret = mxOcx.NewResbuf();
    ret.AddPoint(stp);
    ret.AddPoint(ept);
    mxOcx.SetEventRetEx(ret);
  }
}

function DynInsert() {
  var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
  getPt.message = "点取插入点";
  if (getPt.go() != 1) {
      return;
  }
  var frstPt = getPt.value();
  if (frstPt == null) {
      return;
  }
  var getSecondPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
  var spDrawData = getSecondPt.InitUserDraw("TestDynDraw");
  getSecondPt.message = "点取旋转角度";
  getSecondPt.basePoint = frstPt;
  getSecondPt.setUseBasePt(true);
  spDrawData.SetPoint("pt1", frstPt);
  var sBlkName = "Tree";
  var sBlkFile = mxOcx.GetOcxAppPath() + "\\Blk\\树.dwg";
  mxOcx.InsertBlock(sBlkFile, "Tree");
  spDrawData.SetString("BlkName", "Tree");
  if (getSecondPt.go() != 1) {
      return;
  }
  spDrawData.Draw();
}

//初始化
function InitMxDrawX() {
  if (mxOcx) {
    if (!mxOcx.IsIniting()) 
    {
      clearInterval(mxtime);
      // 控件初始化完成，需要在启动做的事，在这里做
      
      // 启动时打开文件
      //var sDwgFile = mxOcx.GetOcxAppPath() + "\\管道安装大样图.dwg";
      //mxOcx.OpenDwgFile(sDwgFile);
      //....
      mxOcx.OpenWebDwgFile("{{.DwgLink}}?hotqinsessionid={{.Sessionid}}");
      // mxOcx.HideToolBarControl(_T("绘图工具"),_T("绘线,绘矩形框"),true,true);
      document.getElementById("MxDrawXCtrl").EnableToolBarButton("另存为dwg文件", false);
      document.getElementById("MxDrawXCtrl").EnableToolBarButton("保存", false);
      document.getElementById("MxDrawXCtrl").EnableToolBarButton("打开dwg文件", false);
      document.getElementById("MxDrawXCtrl").EnableToolBarButton("保存为mxg文件", false);
      document.getElementById("MxDrawXCtrl").EnableToolBarButton("打开网上dwg文件", false);
      document.getElementById("MxDrawXCtrl").EnableToolBarButton("打开mxg文件", false);
      document.getElementById("MxDrawXCtrl").EnableToolBarButton("新建", false);
      mxOcx.HideMenuBarControl("新建(&N),打开(&O),打开DWG文件(&M),保存(&S),另存为(&A),另存为DWG(&F)",true);
        }
    }
}

mxtime = setInterval(InitMxDrawX, 100);
document.getElementById("MxDrawXCtrl").ImplementCommandEventFun = DoCommandEventFunc;
document.getElementById("MxDrawXCtrl").ImpDynWorldDrawFun = DoDynWorldDrawFun;
document.getElementById("MxDrawXCtrl").ImpExplodeFun = ExplodeFun;
document.getElementById("MxDrawXCtrl").ImpGetGripPointsFun = GetGripPointsFun;
document.getElementById("MxDrawXCtrl").ImpMoveGripPointsAtFun = MoveGripPointsFun;
document.getElementById("MxDrawXCtrl").ImpTransformByFun = TransformByFun;
document.getElementById("MxDrawXCtrl").ImpGetGeomExtentsFun = GetGeomExtentsFun;

function InsertImage() {
  var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
  getPt.message = "点取图片的插入中点";
  if (getPt.go() != 1) {
      return;
  }
  var frstPt = getPt.value();
  if (frstPt == null) {
      return;
  }
  var sImageFile = mxOcx.GetOcxAppPath() + "\\mxcad.jpg";
  mxOcx.DrawImageMark(frstPt.x, frstPt.y, -100.0, 0.0, sImageFile, "", true);
}

function  DrawGif(){
  var getPt = mxOcx.NewComObject("IMxDrawUiPrPoint");
  getPt.message = "点取图片的插入中点";
  if (getPt.go() != 1) {
      return;
  }
  var frstPt = getPt.value();
  if (frstPt == null) {
      return;
  }
  //     var sImageFile1 = mxOcx.GetOcxAppPath() + "\\1.png";
  //     var sImageFile2 = mxOcx.GetOcxAppPath() + "\\2.png";
  //     var sImageFile3 = mxOcx.GetOcxAppPath() + "\\3.png";
  //     
  var sImageFile1 = "./1.png";
  var sImageFile2 = "./2.png";
  var sImageFile3 = "./3.png";
  //    var lId = mxOcx.DrawImageMark(frstPt.x, frstPt.y, -20, 45.0 * 3.14159265 / 180.0, sImageFile1,
  //       sImageFile1 + "," + sImageFile2 + "," + sImageFile3, true);
  var lId = mxOcx.DrawImageMark(frstPt.x, frstPt.y, -20, 0, sImageFile1,
      sImageFile1 + "," + sImageFile2 + "," + sImageFile3, true);
  mxOcx.TwinkeEnt(lId);
}
</script>

</p>

</body>
</html>

