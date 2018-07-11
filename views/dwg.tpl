<HTML>

<HEAD>
<title>DWGView</title>
<meta charset="UTF-8">
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0">
  <meta name="renderer" content="webkit">
  <!-- 加上这句，360等浏览器就会默认使用google内核，而不是IE内核 。 -->
  <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1">
  <!-- 加上这一句，如果被用户强行使用IE浏览器，就会使用IE的最高版本渲染内核 -->
<!-- <LINK REL="stylesheet" TYPE="text/css" HREF="help/ie4.css"/> -->
</HEAD>
<script>
function ZoomIn()
{
     DWGViewX.ZoomIn();
}
function ZoomOut()
{
  DWGViewX.ZoomOut();
}
function ZoomAll()
{
  DWGViewX.ZoomAll();
}
function HideToolbar()
{
    DWGViewX.ShowToobar = !DWGViewX.ShowToobar
}

function Pan()
{
  DWGViewX.PanByMouse();
}

function ZoomWindow()
{
  DWGViewX.ZoomRectByMouse();
}
function HideLayoutBar()
{
  DWGViewX.ShowLayoutBar = !DWGViewX.ShowLayoutBar;
}
function Background()
{
  DWGViewX.Background = DWGViewX.Background ==0? 7:0 
}
function Print()
{
  DWGViewX.Print();
}

</script>
<BODY >

<!-- <a href="javascript:ZoomIn()">Zoom In</a> | <a href="javascript:ZoomOut()">Zoom Out</a> | <a href="javascript:ZoomAll()">Zoom All</a> 
| <a href="javascript:ZoomWindow()">Zoom Window</a>| <a href="javascript:Pan()">Pan</a>| <a href="javascript:HideToolbar()">Show/Hide Toolbar</a> 
 | <a href="javascript:HideLayoutBar()">Show/Hide LayoutBar</a>
 |<a href="javascript:Print()">Print</a>                           
 |<a href="javascript:Background()">Background</a>  -->                          
<table border="0" style="width: 1800;margin:auto">
  <tr style="width: 100%">
    <td style="width: 100%"> 
      <OBJECT id=DWGViewX classid="clsid:AC53EFE4-94A7-47E6-BBFC-E9B9CF322299" codebase="http://www.autodwg.com/dwg-viewer/dwgviewx.cab" width="1200" height="820">
        <param name="_Version" value="65536">
        <param name="_ExtentX" value="18521">
        <param name="_ExtentY" value="13758">
        <param name="_StockProps" value="0">
        <param name="FontPath" value="http://127.0.0.1/static/img/Fonts/">
        <param name="DrawingFile" value="{{.DwgLink}}">
        <param name="ShowToobar" value="-1">
        <param name="ShowLayoutBar" value="1">  
      </OBJECT>
    </td>
  </tr>
    <!-- <td width="50%" valign="top">
      <table border="0" width="100%">-->
        <tr>
          <td width="100%" style="border-bottom:1px dashed">  <img border="0" src="/static/img/loading.gif" width="12" height="11"> 
            If the control can't display correctly, please download and install the control first.  
            <p align="left"><a href="/static/download/dwgviewx.dll">Download DWGViewX</a>
            </td>  
        </tr>
         <!--<tr>
          <td width="100%"></td>
        </tr>
        <tr>
          <td width="100%">
            <p align="right"><a href="DWGViewX.html">Help for Developers..</a>.</td>
        </tr>
        <tr>
          <td width="100%"></td>
        </tr>
      </table>
    </td> --> 
  
</table>

</BODY>
</HTML>