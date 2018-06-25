	<!-- <object id="1" visible="true" classid="clsid:B6FCC215-D303-11D1-BC6C-0000C078797F" type="application/x-oleobject" width="800" height="600"> -->
		<!-- <param name="SRC" value="//127.0.0.1/static/img/test.dwg"></object> -->
<HTML>
<HEAD>
<title>DWGViewX Demo-DWG Viewer ActiveX Control</title>
<!-- <LINK REL="stylesheet" TYPE="text/css" HREF="help/ie4.css"/> -->
</HEAD>
<script>
  // function ZoomIn()
  // {
  //      DWGViewX.ZoomIn();
  // }
  // function ZoomOut()
  // {
  // 	DWGViewX.ZoomOut();
  // }
  // function ZoomAll()
  // {
  // 	DWGViewX.ZoomAll();
  // }
  // function HideToolbar()
  // {
  //     DWGViewX.ShowToobar = !DWGViewX.ShowToobar
  // }
  
  // function Pan()
  // {
  // 	DWGViewX.PanByMouse();
  // }
  
  // function ZoomWindow()
  // {
  // 	DWGViewX.ZoomRectByMouse();
  // }
  // function HideLayoutBar()
  // {
  // 	DWGViewX.ShowLayoutBar = !DWGViewX.ShowLayoutBar;
  // }
  // function Background()
  // {
  // 	DWGViewX.Background = DWGViewX.Background ==0? 7:0 
  // }
  // function Print()
  // {
  // 	DWGViewX.Print();
  // }

function OpenMarkup()
{
  var vPath;
   vPath = "c:\\testmarkup.mrk";
   DWGViewX.LoadMarkupFile(vPath);
}
function SaveMarkup()
{
  var vPath;
   vPath = "c:\\testmarkup.mrk";
   DWGViewX.SaveMarkupFile(vPath);
}
function Measure()
{
     DWGViewX.Measure();
}
function Area()
{
  DWGViewX.Area();
}
function MarkerRedLine()
{
     DWGViewX.MarkerRedLine();
}
function MarkerText()
{
  DWGViewX.MarkerText();
}
function MarkerRect()
{
     DWGViewX.MarkerRect();
}
function MarkerCircleRect()
{
  DWGViewX.MarkerCircleRect();
}
function MarkerEllipse()
{
     DWGViewX.MarkerEllipse();
}
function MarkerCound()
{
  DWGViewX.MarkerCound();
}
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
function HideLayout()
{
  DWGViewX.ShowLayoutBar = !DWGViewX.ShowLayoutBar
}

function Pan()
{
  DWGViewX.PanByMouse();
}

function ZoomWindow()
{
  DWGViewX.ZoomRectByMouse();
}
function Layers()
{
     DWGViewX.Layer();
     
}
function OpenState()
{
  DWGViewX.OpenZoomextend = !OpenZoomextend;
}
function SetPrint()
{
  DWGViewX.PrintLandscape = 1;    // 0 use default
            // 1 PORTRAIT
            // 2 LANDSCAPE
}
function SetPrint2()
{
  DWGViewX.PrintLandscape = 2;    // 0 use default
            // 1 PORTRAIT
            // 2 LANDSCAPE
}
function ExtractBlockAttByWindow()
{
  DWGViewX.ExtractBlockAttByWindow();
  var nLong = DWGViewX.ExtractBlocksCount;
  for (var nIndex=0; nIndex<nLong ; nIndex++)
  {
    DWGViewX.SetCurBlockIndex(nIndex);
    var blkname = "block name:";
    blkname += DWGViewX.CurBlkName;
    alert(blkname);
    var nVCount = DWGViewX.CurBlockValuesCount;
    for (var n=0; n<nVCount ; n++)
    {
           var s1 = DWGViewX.CurBlkTagByIndex(n);
           var s2 = DWGViewX.CurBlkValueByIndex(n);
           var ss = "Tag:";

      ss += s1;
      ss += "  Value:";
      ss += s2;
                        alert(ss);          
    }
  }      
}
function Print()
{
  
  DWGViewX.PrintNoWindow = 1;   // 1 not show print setting dlg
            // 0 show print setting dlg

  DWGViewX.UseDefaultPrint = 0;         // 1 use default print, if not find show printsetup dlg
            // 0 show printsetup

  DWGViewX.PrintBlackAndWhite = 1;  // 1 Black and white mode
            // 0 256 color mode

  DWGViewX.UseDefaultPenWidth  = 1; // use default pen width
            // 0 use user define pen width
  
  //DWGViewX.SetPrintArea(0,0, 100,100);  // when DWGViewX.PrintType = 2, set print size
  DWGViewX.SetAllPenWidth(0.5);   // set all pen is 1.2
  DWGViewX.SetPenWidth(1, 0.05);    // set red (1) color's pen width is 0.05
            // pen width value is 
            // 0.00 0.05 0.09 0.13 0.15 0.18 0.20 0.25
            // 0.30 0.35 0.40 0.50 0.53 0.60 0.70 0.80
            // 0.90 1.00 1.06 1.20 1.40 1.58 2.00 2.11

  DWGViewX.Print();
}
function BackGround()
{
  DWGViewX.Background += 1;
}
function Rotate()
{
  DWGViewX.Rotate(90);
}
</script>


<BODY >
  <!-- <a href="javascript:ZoomIn()">Zoom In</a> 
  | <a href="javascript:ZoomOut()">Zoom Out</a> 
  | <a href="javascript:ZoomAll()">Zoom All</a>
  | <a href="javascript:ZoomWindow()">Zoom Window</a>| <a href="javascript:Pan()">Pan</a>| <a href="javascript:HideToolbar()">Show/Hide Toolbar</a>                             
  | <a href="javascript:HideLayoutBar()">Show/Hide LayoutBar</a>
  | <a href="javascript:Print()">Print</a>                           
  | <a href="javascript:Background()">Background</a> -->
  <!-- <a href="javascript:ZoomIn()">Zoom In</a> | <a href="javascript:ZoomOut()">Zoom Out</a> | <a href="javascript:ZoomAll()">Zoom All</a>          
| <a href="javascript:ZoomWindow()">Zoom Window</a>| <a href="javascript:Pan()">Pan</a>| <a href="javascript:HideToolbar()">Show/Hide Toolbar</a>  
| <a href="javascript:HideLayout()">Show/HideLayoutbar</a>     
| <a href="javascript:BackGround()">BackGround</a>  
| <a href="javascript:Print()">Print</a>   
| <a href="javascript:Rotate()">Rotate</a>
| <a href="javascript:OpenMarkup()">OpenMarkup</a>
| <a href="javascript:SaveMarkup()">SaveMarkup</a>
| <a href="javascript:Measure()">Measure</a>
| <a href="javascript:MarkerCound()">MarkerCound</a>   
| <a href="javascript:MarkerRedLine()">MarkerRedLine</a>   
| <a href="javascript:MarkerText()">MarkerText</a> 
| <a href="javascript:MarkerRect()">MarkerRect</a>   
| <a href="javascript:MarkerCircleRect()">MarkerCircleRect</a>  
| <a href="javascript:MarkerEllipse()">MarkerEllipse</a>   
| <a href="javascript:Layers()">Layers</a>
| <a href="javascript:Area()">Area</a> 
| <a href="javascript:SetPrint()">SetPrint</a>
| <a href="javascript:SetPrint2()">SetPrint2</a>
| <a href="javascript:ExtractBlockAttByWindow()">ExtractBlockAttByWindow</a> 
 <a href="javascript:OpenState()">OpenState</a> -->                      
<table border="0" style="width: 100%;margin:auto">
  <tr>
    <td> 
      <!-- AC53EFE4-94A7-47E6-BBFC-E9B9CF322299_dwgviewx   -->
      <!-- B6FCC215-D303-11D1-BC6C-0000C078797F_autovue  style="margin-left:auto;margin-right:auto;"        -->
      <OBJECT id=DWGViewX classid="clsid:AC53EFE4-94A7-47E6-BBFC-E9B9CF322299" codebase="http://127.0.0.1/static/img/dwgviewx.cab" width="1224" height="850">
        <param name="_Version" value="65536">
        <param name="_ExtentX" value="18521">
        <param name="_ExtentY" value="13758">
        <param name="_StockProps" value="0">
        <!-- <param name="FontPath" value="http://127.0.0.1/static/img/Fonts/"> "http://127.0.0.1/static/img/06.dwg"-->
        <param name="DrawingFile" value="{{.DwgLink}}">
        <param name="ShowToobar" value="-1">
        <param name="ShowLayoutBar" value="1">
      </OBJECT>
    </td>
    <!-- <td width="50%" valign="top">
      <table border="0" width="100%">
        <tr>
          <td width="100%" style="border-bottom:1px dashed">	<img border="0" src="Help/tips.gif" width="12" height="11"> 
            If the control can't display correctly,please download and install the control first.  
            <p align="right"><a href="http://www.autodwg.com/download/dwgviewx.exe">Download DWGViewX</a></td>  
        </tr>
        <tr>
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
  </tr>
</table>

</BODY>
</HTML>