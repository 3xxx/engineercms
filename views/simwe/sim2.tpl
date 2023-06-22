<!DOCTYPE html>
<!-- saved from url=(0124)http://run.simapps.com/sim/simdroid.jsp?num=150804d9a76945de82f4ba249dabc925&downloadId=NzEyMiwyNjU4LDQxNzI1LDQxNDk2LDM1NDM2 -->
<html lang="zh-cn">

<head>
  <meta http-equiv="Content-Type" content="text/html; charset=GBK">
  <title>混凝土剪力墙均布载荷</title>
  <meta name="description" content="">
  <meta name="keywords" content="混凝土剪力墙均布载荷">
  <meta name="wap-font-scale" content="no">
  <link rel="shortcut icon" href="http://run.simapps.com/sim/images/favicon.ico" type="image/x-icon">
  <link rel="stylesheet" type="text/css" href="/static/sim/css/easyui.css">
  <link rel="stylesheet" type="text/css" href="/static/sim/css/icon.css">
  <link rel="stylesheet" type="text/css" href="/static/sim/css/ribbon.css">
  <link rel="stylesheet" type="text/css" href="/static/sim/css/ribbon-icon.css">
  <link rel="stylesheet" type="text/css" href="/static/sim/css/ion.rangeSlider.min.css">
  <link rel="stylesheet" type="text/css" href="/static/sim/css/index.css">
  <script type="text/javascript" src="/static/sim/js/jquery.min.js" charset="utf-8"></script>
  <script type="text/javascript" src="/static/sim/js/jquery.easyui.min.js"></script>
  <script type="text/javascript" src="/static/sim/js/jquery.ribbon.js"></script>
  <script src="/static/sim/js/ion.rangeSlider.min.js" charset="utf-8"></script>
  <script src="/static/sim/js/sim.js" charset="utf-8"></script>
  <script src="/static/sim/js/js-big-decimal.min.js" charset="utf-8"></script>
  <script src="/static/sim/js/wxLogin.js"></script>
  <script src="/static/sim/js/jweixin-1.6.0.js"></script>
  <meta name="viewport" id="viewport" content="width=1170, initial-scale = 1 ,minimum-scale = 1, maximum-scale =1 , user-scalable=yes">
</head>

<body>
  <div style="display:none;"><img src="/static/sim/images/30d2206c-4516-4501-9bc2-3c80b430792d_sm.PNG"></div>
  <!-- S:\ibestorefile2\app//developer//cfa0fd61-ceb0-4ea6-af01-92ee6f1a31c2.xml -->
  <!--  -->
  <!-- 114.116.222.254 -->
  <script type="text/javascript">
  if (browser.versions.mobile) {
    //window.location.href='mobile.jsp?downloadId=NzEyMiwyNjU4LDQxNzI1LDQxNDk2LDM1NDM2&simType=0&num=150804d9a76945de82f4ba249dabc925';
  }
  </script>
  <!-- 遮罩层DIV -->
  <div id="overlay" class="overlay" style="display: none;"></div>
  <!-- Loading提示 DIV -->
  <div id="loadingTip" class="loading-tip" style="top: 408px; left: 960px; display: none;">
    <div class="loader"></div>
  </div>
  <div class="card" id="card"></div>
  <div id="tt">
  </div>
  <input id="token" type="hidden" value="eyJhbGciOiJIUzI1NiJ9.eyJ1c2VyTmFtZSI6ImhvdHFpbjg4OCIsInB3ZCI6Ijk1Nzg3MyIsImV4cCI6MTY0MTMxMTIzNSwiaWF0IjoxNjQxMzA5NDM1fQ.Xr9c1p-_JTrl2qkZfQaFKvc0TbFP-Qxss9Lvh0itxc4">
  <input id="balance" type="hidden" value="1.00">
  <input id="appurl" type="hidden" value="http://run.simapps.com/sim/simdroid.jsp?num=150804d9a76945de82f4ba249dabc925&amp;downloadId=NzEyMiwyNjU4LDQxNzI1LDQxNDk2LDM1NDM2">
  <input id="applogo" type="hidden" value="http://run.simapps.com/sim/images/logo.png">
  <script type="text/javascript">
  var mainFormData = eval([{ "Name": "LabelWidget", "Label": "杨氏模量", "Icon": "", "Type": "8", "x": "30", "y": "30", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "1", "Buddy": "", "ToolTip": "杨氏模量", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit", "Label": "LineEdit", "Icon": "", "Type": "10", "x": "110", "y": "30", "w": "150", "h": "30", "Source": "P_E", "SourceLabel": "P_E", "Value": "2.75e+10 Pa", "Kind": "1", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "", "MinValue": "2.75e+09", "MaxValue": "2.75e+11", "SignalSources": "", "FileCount": "0", "OldValue": "2.75e+10 Pa" }, { "Name": "LabelWidget2", "Label": "泊松比", "Icon": "", "Type": "8", "x": "30", "y": "80", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "477144833", "Buddy": "", "ToolTip": "泊松比", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit1", "Label": "LineEdit1", "Icon": "", "Type": "10", "x": "110", "y": "80", "w": "150", "h": "30", "Source": "P_v", "SourceLabel": "P_v", "Value": "0.2", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.1", "MaxValue": "0.3", "SignalSources": "", "FileCount": "0", "OldValue": "0.2 " }, { "Name": "LabelWidget4", "Label": "内摩擦角", "Icon": "", "Type": "8", "x": "30", "y": "130", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "内摩擦角", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit2", "Label": "LineEdit2", "Icon": "", "Type": "10", "x": "110", "y": "130", "w": "150", "h": "30", "Source": "P_b", "SourceLabel": "P_b", "Value": "68.2283 °", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "50", "MaxValue": "80", "SignalSources": "", "FileCount": "0", "OldValue": "68.2283 &deg" }, { "Name": "LabelWidget6", "Label": "剪胀角", "Icon": "", "Type": "8", "x": "30", "y": "180", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "剪胀角", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit3", "Label": "LineEdit3", "Icon": "", "Type": "10", "x": "110", "y": "180", "w": "150", "h": "30", "Source": "P_a", "SourceLabel": "P_a", "Value": "68.2283 °", "Kind": "1", "Forms": "", "BuddyControl": "586827009", "Buddy": "", "ToolTip": "", "MinValue": "50", "MaxValue": "80", "SignalSources": "", "FileCount": "0", "OldValue": "68.2283 &deg" }, { "Name": "LabelWidget8", "Label": "长度", "Icon": "", "Type": "8", "x": "30", "y": "230", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "长度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit4", "Label": "LineEdit4", "Icon": "", "Type": "10", "x": "110", "y": "230", "w": "150", "h": "30", "Source": "S_L1", "SourceLabel": "S_L1", "Value": "0.59 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.5", "MaxValue": "2", "SignalSources": "", "FileCount": "0", "OldValue": "0.59 m" }, { "Name": "LabelWidget10", "Label": "高度", "Icon": "", "Type": "8", "x": "30", "y": "280", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "高度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit5", "Label": "LineEdit5", "Icon": "", "Type": "10", "x": "110", "y": "280", "w": "150", "h": "30", "Source": "S_H", "SourceLabel": "S_H", "Value": "1.2 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.5", "MaxValue": "3", "SignalSources": "", "FileCount": "0", "OldValue": "1.2 m" }, { "Name": "LabelWidget12", "Label": "厚度", "Icon": "", "Type": "8", "x": "30", "y": "330", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "厚度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit6", "Label": "LineEdit6", "Icon": "", "Type": "10", "x": "110", "y": "330", "w": "150", "h": "30", "Source": "S_T", "SourceLabel": "S_T", "Value": "0.1 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.05", "MaxValue": "0.5", "SignalSources": "", "FileCount": "0", "OldValue": "0.1 m" }, { "Name": "LabelWidget16", "Label": "载荷", "Icon": "", "Type": "8", "x": "30", "y": "380", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "载荷", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit8", "Label": "LineEdit8", "Icon": "", "Type": "10", "x": "110", "y": "380", "w": "150", "h": "30", "Source": "L_F1", "SourceLabel": "L_F1", "Value": "-5900000 N", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "-7e+06", "MaxValue": "-4e+06", "SignalSources": "", "FileCount": "0", "OldValue": "-5900000 N" }, { "Name": "TabWidget", "Label": "表单集合", "Icon": "", "Type": "14", "x": "280", "y": "10", "w": "822", "h": "657", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "Form1,Form2,Form3,Form4", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "ButtonWidget", "Label": "计算", "Icon": "", "Type": "13", "x": "30", "y": "430", "w": "60", "h": "30", "Source": "App_Recomputer", "SourceLabel": "计算", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "586827009", "Buddy": "", "ToolTip": "", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }]);
  var mainFormInfoData = eval([{ "Name": "Form", "Label": "参数1", "Width": "1134", "Height": "679" }]);
  var allFormData = eval([{ "Name": "Form", "Label": "参数1", "Width": "1134", "Height": "679", "Widget": [{ "Name": "LabelWidget", "Label": "杨氏模量", "Icon": "", "Type": "8", "x": "30", "y": "30", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "1", "Buddy": "", "ToolTip": "杨氏模量", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit", "Label": "LineEdit", "Icon": "", "Type": "10", "x": "110", "y": "30", "w": "150", "h": "30", "Source": "P_E", "SourceLabel": "P_E", "Value": "2.75e+10 Pa", "Kind": "1", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "", "MinValue": "2.75e+09", "MaxValue": "2.75e+11", "SignalSources": "", "FileCount": "0", "OldValue": "2.75e+10 Pa" }, { "Name": "LabelWidget2", "Label": "泊松比", "Icon": "", "Type": "8", "x": "30", "y": "80", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "477144833", "Buddy": "", "ToolTip": "泊松比", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit1", "Label": "LineEdit1", "Icon": "", "Type": "10", "x": "110", "y": "80", "w": "150", "h": "30", "Source": "P_v", "SourceLabel": "P_v", "Value": "0.2", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.1", "MaxValue": "0.3", "SignalSources": "", "FileCount": "0", "OldValue": "0.2 " }, { "Name": "LabelWidget4", "Label": "内摩擦角", "Icon": "", "Type": "8", "x": "30", "y": "130", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "内摩擦角", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit2", "Label": "LineEdit2", "Icon": "", "Type": "10", "x": "110", "y": "130", "w": "150", "h": "30", "Source": "P_b", "SourceLabel": "P_b", "Value": "68.2283 °", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "50", "MaxValue": "80", "SignalSources": "", "FileCount": "0", "OldValue": "68.2283 &deg" }, { "Name": "LabelWidget6", "Label": "剪胀角", "Icon": "", "Type": "8", "x": "30", "y": "180", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "剪胀角", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit3", "Label": "LineEdit3", "Icon": "", "Type": "10", "x": "110", "y": "180", "w": "150", "h": "30", "Source": "P_a", "SourceLabel": "P_a", "Value": "68.2283 °", "Kind": "1", "Forms": "", "BuddyControl": "586827009", "Buddy": "", "ToolTip": "", "MinValue": "50", "MaxValue": "80", "SignalSources": "", "FileCount": "0", "OldValue": "68.2283 &deg" }, { "Name": "LabelWidget8", "Label": "长度", "Icon": "", "Type": "8", "x": "30", "y": "230", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "长度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit4", "Label": "LineEdit4", "Icon": "", "Type": "10", "x": "110", "y": "230", "w": "150", "h": "30", "Source": "S_L1", "SourceLabel": "S_L1", "Value": "0.59 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.5", "MaxValue": "2", "SignalSources": "", "FileCount": "0", "OldValue": "0.59 m" }, { "Name": "LabelWidget10", "Label": "高度", "Icon": "", "Type": "8", "x": "30", "y": "280", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "高度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit5", "Label": "LineEdit5", "Icon": "", "Type": "10", "x": "110", "y": "280", "w": "150", "h": "30", "Source": "S_H", "SourceLabel": "S_H", "Value": "1.2 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.5", "MaxValue": "3", "SignalSources": "", "FileCount": "0", "OldValue": "1.2 m" }, { "Name": "LabelWidget12", "Label": "厚度", "Icon": "", "Type": "8", "x": "30", "y": "330", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "厚度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit6", "Label": "LineEdit6", "Icon": "", "Type": "10", "x": "110", "y": "330", "w": "150", "h": "30", "Source": "S_T", "SourceLabel": "S_T", "Value": "0.1 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.05", "MaxValue": "0.5", "SignalSources": "", "FileCount": "0", "OldValue": "0.1 m" }, { "Name": "LabelWidget16", "Label": "载荷", "Icon": "", "Type": "8", "x": "30", "y": "380", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "载荷", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit8", "Label": "LineEdit8", "Icon": "", "Type": "10", "x": "110", "y": "380", "w": "150", "h": "30", "Source": "L_F1", "SourceLabel": "L_F1", "Value": "-5900000 N", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "-7e+06", "MaxValue": "-4e+06", "SignalSources": "", "FileCount": "0", "OldValue": "-5900000 N" }, { "Name": "TabWidget", "Label": "表单集合", "Icon": "", "Type": "14", "x": "280", "y": "10", "w": "822", "h": "657", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "Form1,Form2,Form3,Form4", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "ButtonWidget", "Label": "计算", "Icon": "", "Type": "13", "x": "30", "y": "430", "w": "60", "h": "30", "Source": "App_Recomputer", "SourceLabel": "计算", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "586827009", "Buddy": "", "ToolTip": "", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }] }, { "Name": "Form1", "Label": "几何", "Width": "822", "Height": "622", "Widget": [{ "Name": "Geometry", "Label": "几何1", "Icon": "", "Type": "12", "x": "10", "y": "10", "w": "800", "h": "600", "Source": "Split", "SourceLabel": "几何", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507116801", "Buddy": "", "ToolTip": "几何", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }] }, { "Name": "Form2", "Label": "网格", "Width": "822", "Height": "622", "Widget": [{ "Name": "Mesh", "Label": "网格1", "Icon": "", "Type": "11", "x": "10", "y": "10", "w": "800", "h": "600", "Source": "MeshDomain", "SourceLabel": "网格", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "477139201", "Buddy": "", "ToolTip": "网格", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }] }, { "Name": "Form3", "Label": "等效塑性应变云图", "Width": "822", "Height": "622", "Widget": [{ "Name": "Nephogram", "Label": "结果", "Icon": "", "Type": "11", "x": "10", "y": "10", "w": "800", "h": "600", "Source": "Nephogram:{\"ObjectName\":\"Nephogram\",\"ResultModelName\":\"ResultModel\",\"CaseFile\":\"Solving/SolvingDomain/result/1-General.case\",\"ArrayName\":\"PEEQ\",\"ArrayComponent\":-1,\"FieldAssociation\":0,\"PartName\":[\"Split-Face1\"],\"Warp\":false,\"WarpFactor\":1.0,\"SourceType\":\"ResultFile\",\"GeometryType\":\"Edge\"};{\"CameraInfor\":[{\"CameraFocalPoint\":[0.29499998688697817,0.6000000238418579,0.0]},{\"CameraPosition\":[-0.45020858165868057,1.239957474391526,2.3892256091609549]},{\"CameraViewUp\":[-0.6529783811788842,0.6555750851259989,-0.37926315596831086]},{\"CameraViewAngle\":30.0},{\"CameraClippingRange\":[2.09709412648625,3.201494714929422]}]}", "SourceLabel": "云图", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "477139201", "Buddy": "", "ToolTip": "云图", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }] }, { "Name": "Form4", "Label": "位移矢量图", "Width": "822", "Height": "622", "Widget": [{ "Name": "GlyphArrow1", "Label": "结果1", "Icon": "", "Type": "11", "x": "10", "y": "10", "w": "800", "h": "600", "Source": "GlyphArrow:{\"ObjectName\":\"GlyphArrow\",\"ResultModelName\":\"ResultModel\",\"CaseFile\":\"Solving/SolvingDomain/result/1-General.case\",\"ArrayName\":\"Disp\",\"ArrayComponent\":-1,\"FieldAssociation\":0,\"PartName\":[\"Split-Face1\"],\"Warp\":true,\"WarpFactor\":0.1,\"SourceType\":\"ResultFile\",\"GeometryType\":\"Edge\",\"ScaledByArray\":true,\"OrientatedByArray\":true,\"GlyphType\":0,\"GlyphFactor\":0.05008778206,\"GlyphMode\":0};{\"CameraInfor\":[{\"CameraFocalPoint\":[0.40410139550668,0.4871888954659533,-0.08111811883641061]},{\"CameraPosition\":[0.13338656589793869,0.7640899228282316,-2.63519714319999]},{\"CameraViewUp\":[0.46799605191049678,-0.8718978419949927,-0.14413135856108012]},{\"CameraViewAngle\":30.0},{\"CameraClippingRange\":[2.4054057427067239,2.9840228960551817]}]}", "SourceLabel": "矢量图", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "矢量图", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }] }]);
  var allPageData = eval([]);
  var groupInputData = eval([{ "Name": "LabelWidget", "Label": "杨氏模量", "Icon": "", "Type": "8", "x": "30", "y": "30", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "1", "Buddy": "", "ToolTip": "杨氏模量", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit", "Label": "LineEdit", "Icon": "", "Type": "10", "x": "110", "y": "30", "w": "150", "h": "30", "Source": "P_E", "SourceLabel": "P_E", "Value": "2.75e+10 Pa", "Kind": "1", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "", "MinValue": "2.75e+09", "MaxValue": "2.75e+11", "SignalSources": "", "FileCount": "0" }, { "Name": "LabelWidget2", "Label": "泊松比", "Icon": "", "Type": "8", "x": "30", "y": "80", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "477144833", "Buddy": "", "ToolTip": "泊松比", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit1", "Label": "LineEdit1", "Icon": "", "Type": "10", "x": "110", "y": "80", "w": "150", "h": "30", "Source": "P_v", "SourceLabel": "P_v", "Value": "0.2 ", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.1", "MaxValue": "0.3", "SignalSources": "", "FileCount": "0" }, { "Name": "LabelWidget4", "Label": "内摩擦角", "Icon": "", "Type": "8", "x": "30", "y": "130", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "内摩擦角", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit2", "Label": "LineEdit2", "Icon": "", "Type": "10", "x": "110", "y": "130", "w": "150", "h": "30", "Source": "P_b", "SourceLabel": "P_b", "Value": "68.2283 °", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "50", "MaxValue": "80", "SignalSources": "", "FileCount": "0" }, { "Name": "LabelWidget6", "Label": "剪胀角", "Icon": "", "Type": "8", "x": "30", "y": "180", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "剪胀角", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit3", "Label": "LineEdit3", "Icon": "", "Type": "10", "x": "110", "y": "180", "w": "150", "h": "30", "Source": "P_a", "SourceLabel": "P_a", "Value": "68.2283 °", "Kind": "1", "Forms": "", "BuddyControl": "586827009", "Buddy": "", "ToolTip": "", "MinValue": "50", "MaxValue": "80", "SignalSources": "", "FileCount": "0" }, { "Name": "LabelWidget8", "Label": "长度", "Icon": "", "Type": "8", "x": "30", "y": "230", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "长度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit4", "Label": "LineEdit4", "Icon": "", "Type": "10", "x": "110", "y": "230", "w": "150", "h": "30", "Source": "S_L1", "SourceLabel": "S_L1", "Value": "0.59 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.5", "MaxValue": "2", "SignalSources": "", "FileCount": "0" }, { "Name": "LabelWidget10", "Label": "高度", "Icon": "", "Type": "8", "x": "30", "y": "280", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "高度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit5", "Label": "LineEdit5", "Icon": "", "Type": "10", "x": "110", "y": "280", "w": "150", "h": "30", "Source": "S_H", "SourceLabel": "S_H", "Value": "1.2 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.5", "MaxValue": "3", "SignalSources": "", "FileCount": "0" }, { "Name": "LabelWidget12", "Label": "厚度", "Icon": "", "Type": "8", "x": "30", "y": "330", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507113729", "Buddy": "", "ToolTip": "厚度", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit6", "Label": "LineEdit6", "Icon": "", "Type": "10", "x": "110", "y": "330", "w": "150", "h": "30", "Source": "S_T", "SourceLabel": "S_T", "Value": "0.1 m", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "0.05", "MaxValue": "0.5", "SignalSources": "", "FileCount": "0" }, { "Name": "LabelWidget16", "Label": "载荷", "Icon": "", "Type": "8", "x": "30", "y": "380", "w": "60", "h": "30", "Source": "", "SourceLabel": "", "Value": "", "Kind": "0", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "载荷", "MinValue": "", "MaxValue": "", "SignalSources": "", "FileCount": "0" }, { "Name": "LineEdit8", "Label": "LineEdit8", "Icon": "", "Type": "10", "x": "110", "y": "380", "w": "150", "h": "30", "Source": "L_F1", "SourceLabel": "L_F1", "Value": "-5900000 N", "Kind": "1", "Forms": "", "BuddyControl": "507110145", "Buddy": "", "ToolTip": "", "MinValue": "-7e+06", "MaxValue": "-4e+06", "SignalSources": "", "FileCount": "0" }]);
  var num = "150804d9a76945de82f4ba249dabc925";
  var downloadId = "NzEyMiwyNjU4LDQxNzI1LDQxNDk2LDM1NDM2";
  var simType = "0";
  var simNode = "";
  var act_value = "";
  var allDatName = "bak_Geometry.dat,Nephogram.dat,Mesh.dat,bak_GlyphArrow1.dat,bak_Mesh.dat,bak_Nephogram.dat,Geometry.dat,GlyphArrow1.dat,";
  var appParamTxtData = eval({ "P_E": "2.75e+010", "P_v": "0.2", "P_b": "68.2283", "P_a": "68.2283", "L_F1": "-5.9e+006", "S_L1": "590", "S_H": "1200", "S_T": "100", "S_L2": "147.5" });
  var appstoreUrl = "https://www.simapps.com/";
  var websocketUrl = "ws://run.simapps.com/websocket/";
  var webservletUrl = "http://run.simapps.com/SimdroidServlet";
  var appstoreSimFreeUrl = "https://www.simapps.com/jsondata/simfreeurl.action";
  var calculationNode = "0";
  var appVersion = "4";
  var isLan = "0";

  $("#w").width(parseInt(mainFormInfoData[0]["Width"]) + 50);
  $("#w").height(parseInt(mainFormInfoData[0]["Height"]) + 50);

  if (allPageData != '') {
    $("#page").css("top", "170px");
    $("#w").height(parseInt(mainFormInfoData[0]["Height"]) + 200);
    $(".easyui-ribbon").show();
  }

  for (var i = 0; i < groupInputData.length; i++) {
    var _type = groupInputData[i]['Type'];
    if (_type == "10") {
      var _Source = groupInputData[i]['Source'];
      var _Value = groupInputData[i]['Value'];
      d[_Source] = _Value;
    }
  }

  drawWidgetPC($("#page"), eval(mainFormData));
  </script>
  <script type="text/javascript">
  //登录切换
  $(function() {
    var pc = $("#pc"),
      wx = $("#wx"),
      login_box = $("#loginbox"),
      wx_box = $("#wxsm");
    wx.on('click', function(event) {
      $(this).hide();
      pc.show();
      login_box.hide();
      wx_box.show();
    });
    pc.on('click', function(event) {
      $(this).hide();
      wx.show();
      login_box.show();
      wx_box.hide();
    });
  })
  $("#simdroid_login").show();
  </script>
  <script type="text/javascript">
  //加载验证登录状态
  checkUserFun();

  function submitForm() {
    submitFun();
  }

  function clearForm() {
    $('#ff').form('clear');
  }

  function submitFun() {

    if ($('#username').val() == "") {
      $.messager.alert('提示', "请输入用户名");
      return false;
    }
    if ($('#password').val() == "") {
      $.messager.alert('提示', "请输入密码");
      return false;
    }
    var json = {
      "method": "accountLogin",
      "parameters": {
        "account": $('#username').val(),
        "password": $('#password').val(),
        "deviceID": "simdroid",
        "deviceName": "simdroid",
      }
    }
    $.ajax({
      url: appstoreUrl + 'client_windows_pc.action',
      data: 'data=' + JSON.stringify(json),
      type: 'post',
      cache: false,
      dataType: 'json',
      success: function(data) {

        if (data.result == "0") {
          console.log("登录成功");
          //console.log(data.data);
          localStorage.setItem("ibe_token", data.data.token);
          localStorage.setItem("ibe_username", $('#username').val());
          localStorage.setItem("ibe_password", $('#password').val());
          $('#token').val(data.data.token);
          $('#balance').val(data.data.balance);
          $('#login_win').window('close');
          $.messager.show({
            title: '登录成功',
            msg: data.data.nickname + '，你好<br>账户余额：' + data.data.balance + "元",
            showType: 'show'
          });
        } else {
          $.messager.alert('提示', data.message);
        }
      },
      error: function(data) {
        console.log(data);
        $.messager.alert('提示', "操作异常");
      }
    });
    return false;
  };

  function checkUserFun() {

    var token = localStorage.getItem('ibe_token');
    if (!token) {
      console.log("token未保存");
      return;
    } else {
      console.log("token已保存");
    }

    var json = {
      "method": "accountLoginToken",
      "parameters": {
        "token": token,
        "type": "ibe",
        "deviceID": "simdroid",
        "deviceName": "simdroid",
      }
    }
    $.ajax({
      url: appstoreUrl + 'client_windows_pc.action',
      data: 'data=' + JSON.stringify(json),
      type: 'post',
      cache: false,
      dataType: 'json',
      success: function(data) {
        if (data.result == "0") {
          console.log("token登录成功");
          localStorage.setItem("ibe_username", data.data.account);
          localStorage.setItem("ibe_password", data.data.password);
          $('#token').val(token);
          $('#balance').val(data.data.balance);
          $.messager.show({
            title: '登录成功',
            msg: data.data.nickname + '，你好<br>账户余额：' + data.data.balance + "元",
            showType: 'show'
          });

        } else {
          console.log("token无效");
          localStorage.removeItem('ibe_token');
          $('#token').val("");
          $('#balance').val("");
        }
      },
      error: function(data) {
        console.log(data);
      }
    });
  };
  </script>
  <script language="JavaScript">
    //微信登录
$(document).ready(function()
{
    var obj = new WxLogin
    ({
        id:"login_container",
        appid: "wx1f7a11e1fe8191df",
        scope: "snsapi_login",
        redirect_uri:encodeURI(appstoreUrl+"servlet/WeixinWebServlet") ,
        state: encodeURIComponent(encodeURIComponent($('#appurl').val()))
    });

});


//微信分享
var link = location.href;
$.ajax({
　　url: "../servlet/WXSignServlet",
　　type: "post",
　　data:{url:link},
　　cache:false,
　　dataType:'json',
　　success: function (data) {
　　　　var datad = data;
　　　　wx.config({
　　　　　　debug: false,
　　　　　　appId: datad.appid,
　　　　　　timestamp: datad.timestamp,
　　　　　　nonceStr: datad.noncestr,
　　　　　　signature: datad.signature,
　　　　　　jsApiList: [
　　　　　　　　"updateAppMessageShareData","updateTimelineShareData"
　　　　　　]
　　　　});
　　　　wx.error(function (res) {
　　　　　　console.log(res);
　　　　});
　　},
　　error: function (error) {
　　　　console.log(error)
　　}
});

wx.ready(function () {
  wx.updateAppMessageShareData({
    title: '混凝土剪力墙均布载荷',
    desc: '',
    link: window.location.href,
    imgUrl: 'https://cdnwww.simapps.com/upload/image/20200713/30d2206c-4516-4501-9bc2-3c80b430792d_sm.PNG',
    success: function () {
      // 设置成功
    }
  })
});

wx.ready(function () {
  wx.updateTimelineShareData({
    title: '混凝土剪力墙均布载荷',
    link: window.location.href,
    imgUrl: 'https://cdnwww.simapps.com/upload/image/20200713/30d2206c-4516-4501-9bc2-3c80b430792d_sm.PNG',
    success: function () {
      // 设置成功
    }
  })
});
</script>
  <script src="/static/sim/js/loader.js" charset="UTF-8"></script>
  <script src="/static/sim/js/phone.js" charset="UTF-8"></script>
  <div class="panel window panel-htop" style="display: block; width: 1194px; left: 355px; top: -253px; position: absolute; z-index: 9001;">
    <div class="panel-header panel-header-noborder window-header" style="width: 1182px;">
      <div class="panel-title panel-with-icon">混凝土剪力墙均布载荷</div>
      <div class="panel-icon icon-app"></div>
      <div class="panel-tool"><a href="javascript:void(0)" class="icon-share panel-tool-a" onclick="shareDialog();"></a></div>
    </div>
    <div id="w" class="easyui-window panel-body panel-body-noborder window-body" title="" data-options="iconCls:&#39;icon-app&#39;,collapsible:false,closable:false,minimizable:false,maximizable:false,draggable:false,resizable:false,tools:&#39;#tt&#39;" style="padding: 5px; width: 1182px; height: 700px;">
      <div class="easyui-ribbon ribbon tabs-container easyui-fluid" style="width: 1168px; display: none;">
        <div class="tabs-header" style="width: 1168px;">
          <div class="tabs-scroller-left" style="display: none;"></div>
          <div class="tabs-scroller-right" style="display: none;"></div>
          <div class="tabs-wrap" style="margin-left: 0px; margin-right: 0px; width: 1166px;">
            <ul class="tabs" style="height: 32px;">
              <li class="tabs-first tabs-last tabs-selected"><a href="javascript:;" class="tabs-inner" style="height: 32px; line-height: 30px;"><span class="tabs-title">操作</span><span class="tabs-icon"></span></a></li>
            </ul>
          </div>
        </div>
        <div class="tabs-panels" style="width: 1168px;">
          <div class="panel panel-htop" style="width: 1166px;">
            <div title="" class="panel-body panel-body-noheader panel-body-noborder ribbon-tab" id="" style="width: 1166px;">
              <div class="ribbon-group">
                <!--
					<div class="ribbon-toolbar">
						<a href="javascript:void(0)" class="easyui-menubutton"
							data-options="name:'model',iconCls:'icon-model-large',iconAlign:'top',size:'large'"
							name="Model_bt">生成几何</a>
					</div>
					<div class="ribbon-toolbar">
						<a href="javascript:void(0)" class="easyui-menubutton"
							data-options="name:'mesh',iconCls:'icon-mesh-large',iconAlign:'top',size:'large'"
							name="Mesh_bt">网格剖分</a>
					</div>
					 -->
                <div class="ribbon-toolbar">
                  <a href="http://run.simapps.com/sim/simdroid.jsp?num=150804d9a76945de82f4ba249dabc925&amp;downloadId=NzEyMiwyNjU4LDQxNzI1LDQxNDk2LDM1NDM2#" class="easyui-menubutton l-btn l-btn-large l-btn-plain m-btn m-btn-large" data-options="name:&#39;recomputer&#39;,iconCls:&#39;icon-recomputer-large&#39;,iconAlign:&#39;top&#39;,size:&#39;large&#39;" name="Recompute_bt" group="" id=""><span class="l-btn-left l-btn-icon-top"><span class="l-btn-text">计算</span><span class="l-btn-icon icon-recomputer-large">&nbsp;</span><span class="m-btn-downarrow"></span><span class="m-btn-line"></span></span></a>
                </div>
                <div class="ribbon-toolbar">
                  <a href="http://run.simapps.com/sim/simdroid.jsp?num=150804d9a76945de82f4ba249dabc925&amp;downloadId=NzEyMiwyNjU4LDQxNzI1LDQxNDk2LDM1NDM2#" class="easyui-menubutton l-btn l-btn-large l-btn-plain m-btn m-btn-large" data-options="name:&#39;share&#39;,iconCls:&#39;icon-share-large&#39;,iconAlign:&#39;top&#39;,size:&#39;large&#39;" name="Share" group="" id=""><span class="l-btn-left l-btn-icon-top"><span class="l-btn-text">分享</span><span class="l-btn-icon icon-share-large">&nbsp;</span><span class="m-btn-downarrow"></span><span class="m-btn-line"></span></span></a>
                </div>
              </div>
              <div class="ribbon-group-sep"></div>
            </div>
          </div>
        </div>
      </div>
      <div id="page" class="page">
        <div class="wcss sim_label" title="杨氏模量" id="LabelWidget_0" name="LabelWidget" style="top: 30px; left: 30px; width: 60px; height: 30px;">杨氏模量</div><input class="wcss tooltip-f" type="text" style="z-index: 1; top: 30px; left: 110px; width: 154px; height: 30px;" id="LineEdit_1" name="P_E" value="2.75e+10 Pa" title="">
        <div class="wcss sim_label" title="泊松比" id="LabelWidget2_2" name="LabelWidget2" style="top: 80px; left: 30px; width: 60px; height: 30px;">泊松比</div><input class="wcss tooltip-f" type="text" style="z-index: 1; top: 80px; left: 110px; width: 154px; height: 30px;" id="LineEdit1_3" name="P_v" value="0.2" title="">
        <div class="wcss sim_label" title="内摩擦角" id="LabelWidget4_4" name="LabelWidget4" style="top: 130px; left: 30px; width: 60px; height: 30px;">内摩擦角</div><input class="wcss tooltip-f" type="text" style="z-index: 1; top: 130px; left: 110px; width: 154px; height: 30px;" id="LineEdit2_5" name="P_b" value="68.2283 °" title="">
        <div class="wcss sim_label" title="剪胀角" id="LabelWidget6_6" name="LabelWidget6" style="top: 180px; left: 30px; width: 60px; height: 30px;">剪胀角</div><input class="wcss tooltip-f" type="text" style="z-index: 1; top: 180px; left: 110px; width: 154px; height: 30px;" id="LineEdit3_7" name="P_a" value="68.2283 °" title="">
        <div class="wcss sim_label" title="长度" id="LabelWidget8_8" name="LabelWidget8" style="top: 230px; left: 30px; width: 60px; height: 30px;">长度</div><input class="wcss tooltip-f" type="text" style="z-index: 1; top: 230px; left: 110px; width: 154px; height: 30px;" id="LineEdit4_9" name="S_L1" value="0.59 m" title="">
        <div class="wcss sim_label" title="高度" id="LabelWidget10_10" name="LabelWidget10" style="top: 280px; left: 30px; width: 60px; height: 30px;">高度</div><input class="wcss tooltip-f" type="text" style="z-index: 1; top: 280px; left: 110px; width: 154px; height: 30px;" id="LineEdit5_11" name="S_H" value="1.2 m" title="">
        <div class="wcss sim_label" title="厚度" id="LabelWidget12_12" name="LabelWidget12" style="top: 330px; left: 30px; width: 60px; height: 30px;">厚度</div><input class="wcss tooltip-f" type="text" style="z-index: 1; top: 330px; left: 110px; width: 154px; height: 30px;" id="LineEdit6_13" name="S_T" value="0.1 m" title="">
        <div class="wcss sim_label" title="载荷" id="LabelWidget16_14" name="LabelWidget16" style="top: 380px; left: 30px; width: 60px; height: 30px;">载荷</div><input class="wcss tooltip-f" type="text" style="z-index: 1; top: 380px; left: 110px; width: 154px; height: 30px;" id="LineEdit8_15" name="L_F1" value="-5900000 N" title="">
        <div class="wcss tabs-container" id="TabWidget" type="14" style="top: 10px; left: 280px; width: 822px; height: 657px;">
          <div class="tabs-header tabs-header-plain" style="width: 822px;">
            <div class="tabs-scroller-left" style="display: none;"></div>
            <div class="tabs-scroller-right" style="display: none;"></div>
            <div class="tabs-wrap" style="margin-left: 0px; margin-right: 0px; width: 822px;">
              <ul class="tabs">
                <li class="tabs-first"><a href="javascript:;" class="tabs-inner" style="line-height: 23px;"><span class="tabs-title">几何</span><span class="tabs-icon"></span></a></li>
                <li class=""><a href="javascript:;" class="tabs-inner" style="line-height: 23px;"><span class="tabs-title">网格</span><span class="tabs-icon"></span></a></li>
                <li class="tabs-selected"><a href="javascript:;" class="tabs-inner" style="line-height: 23px;"><span class="tabs-title">等效塑性应变云图</span><span class="tabs-icon"></span></a></li>
                <li class="tabs-last"><a href="javascript:;" class="tabs-inner" style="line-height: 23px;"><span class="tabs-title">位移矢量图</span><span class="tabs-icon"></span></a></li>
              </ul>
            </div>
          </div>
          <div class="tabs-panels" style="height: 629px; width: 822px;">
            <div class="panel panel-htop" style="width: 820px; display: none;">
              <div title="" id="tab_TabWidget_panel_0" style="overflow: hidden; width: 820px; height: 628px;" class="panel-body panel-body-noheader panel-body-noborder">
                <div style="position:relative;padding:0px;height:100%;widht:100%" id="TabWidget_panel_0" type="panel"><iframe frameborder="0" height="100%" width="100%" scrolling="no" id="iframe_Geometry" src="/static/sim/html/prep.html">

                    <head>
                      <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
                    </head>
                  </iframe></div>
              </div>
            </div>
            <div class="panel panel-htop" style="width: 820px; display: none;">
              <div title="" id="tab_TabWidget_panel_1" style="overflow: hidden; width: 820px; height: 628px;" class="panel-body panel-body-noheader panel-body-noborder">
                <div style="position:relative;padding:0px;height:100%;widht:100%" id="TabWidget_panel_1" type="panel"><iframe frameborder="0" height="100%" width="100%" scrolling="no" id="iframe_Mesh" src="/static/sim/html/mesh.html">

                    <head>
                      <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
                    </head>
                  </iframe></div>
              </div>
            </div>
            <div class="panel panel-htop" style="width: 820px;">
              <div title="" id="tab_TabWidget_panel_2" style="overflow: hidden; width: 820px; height: 628px;" class="panel-body panel-body-noheader panel-body-noborder">
                <div style="position:relative;padding:0px;height:100%;widht:100%" id="TabWidget_panel_2" type="panel"><iframe frameborder="0" height="100%" width="100%" scrolling="no" id="iframe_Nephogram" src="/static/sim/html/post.html">

                    <head>
                      <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
                    </head>
                  </iframe></div>
              </div>
            </div>
            <div class="panel panel-htop" style="width: 820px; display: none;">
              <div title="" id="tab_TabWidget_panel_3" style="overflow: hidden; width: 820px; height: 628px;" class="panel-body panel-body-noheader panel-body-noborder">
                <div style="position:relative;padding:0px;height:100%;widht:100%" id="TabWidget_panel_3" type="panel"><iframe frameborder="0" height="100%" width="100%" scrolling="no" id="iframe_GlyphArrow1" src="/static/sim/html/post(1).html">

                    <head>
                      <meta http-equiv="Content-Type" content="text/html; charset=utf-8" />
                    </head>
                  </iframe></div>
              </div>
            </div>
          </div>
        </div><button class="wcss sim_button" type="button" source="App_Recomputer" id="ButtonWidget17" name="ButtonWidget" style="top: 430px; left: 30px; width: 60px; height: 34px; z-index: 1;">计算</button>
      </div>
      <div id="dd"></div>
      <!--模态分析窗口 -->
    </div>
  </div>
  <div class="window-shadow" style="display: block; left: 355px; top: -253px; position: absolute; z-index: 9000; width: 1194px; height: 739px;"></div>
  <div class="panel window panel-htop" style="display: none; width: 400px; left: 760px; top: 208px;">
    <div class="panel-header panel-header-noborder window-header" style="width: 388px;">
      <div class="panel-title">请先登录</div>
      <div class="panel-tool"><a href="javascript:;" class="panel-tool-close"></a></div>
    </div>
    <div id="login_win" class="easyui-window panel-body panel-body-noborder window-body" title="" data-options="modal:true,resizable:false,minimizable:false,maximizable:false,closable:true,collapsible:false,shadow:false,closed:true" style="padding: 10px; width: 388px; height: 361px;">
      <div class="simdroid_login" id="simdroid_login" style="">
        <div class="mod_login">
          <div class="tab_qrcode">
            <a href="javascript:void(0);" class="wx" id="wx" title="二维码" style="display: block;"></a>
            <a href="javascript:void(0);" class="pc" id="pc" title="电脑登录" style="display: none;"></a>
          </div>
          <div class="loginbox" id="loginbox">
            <div class="login-tip">
              <div class="poptip">
                <div class="poptip-arrow">
                  <em></em>
                  <span></span>
                </div>
                <div class="poptip-content">
                  扫码登录更安全
                </div>
              </div>
            </div>
            <div class="l_form">
              <form id="ff" method="post">
                <div style="margin-bottom:20px">
                  <input class="easyui-textbox textbox-f" id="username" prompt="用户名" iconwidth="28" style="width: 100%; height: 34px; padding: 10px; display: none;" textboxname="username"><span class="textbox easyui-fluid" style="width: 240px; height: 34px;"><input id="_easyui_textbox_input1" type="text" class="textbox-text validatebox-text textbox-prompt" autocomplete="off" tabindex="" placeholder="用户名" style="padding: 0px 10px; margin: 0px; height: 32px; line-height: 32px; width: 238px;"><input type="hidden" class="textbox-value" name="username" value=""></span>
                </div>
                <div style="margin-bottom:20px">
                  <input class="easyui-passwordbox passwordbox-f textbox-f" id="password" prompt="密码" iconwidth="28" style="width: 100%; height: 34px; padding: 10px; display: none;" textboxname="password"><span class="textbox easyui-fluid" style="width: 240px; height: 34px;"><span class="textbox-addon textbox-addon-right" style="right: 0px; top: 0px;"><a href="javascript:;" class="textbox-icon passwordbox-open" icon-index="0" tabindex="-1" style="width: 28px; height: 32px;"></a></span><input id="_easyui_textbox_input2" type="text" class="textbox-text validatebox-text textbox-prompt" autocomplete="off" tabindex="" placeholder="密码" style="padding: 0px 10px; margin: 0px 28px 0px 0px; height: 32px; line-height: 32px; width: 210px;"><input type="hidden" class="textbox-value" name="password" value=""></span>
                </div>
                <div style="text-align:center;padding:5px 0">
                  <a href="javascript:void(0)" class="easyui-linkbutton l-btn l-btn-small" onclick="submitForm()" style="width:80px" group="" id=""><span class="l-btn-left" style="margin-top: 0px;"><span class="l-btn-text">登录</span></span></a>
                  <a href="https://www.simapps.com/reg/create.action" target="_blank" class="easyui-linkbutton l-btn l-btn-small" style="width:80px" group="" id=""><span class="l-btn-left" style="margin-top: 0px;"><span class="l-btn-text">免费注册</span></span></a>
                </div>
              </form>
            </div>
          </div>
          <div class="wxsm" id="wxsm" style="display:none;">
            <div id="login_container"><iframe src="/static/sim/html/qrconnect.html" frameborder="0" scrolling="no" width="300px" height="400px"></iframe></div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div class="window-mask" style="display:none"></div>
  <div class="panel window panel-htop" style="display: none; width: 500px; left: 710px; top: 298px;">
    <div class="panel-header panel-header-noborder window-header" style="width: 488px;">
      <div class="panel-title">请稍候...</div>
      <div class="panel-tool"></div>
    </div>
    <div id="win" title="" class="panel-body panel-body-noborder window-body" style="width: 488px; height: 181px;"></div>
  </div>
  <div class="window-mask" style="display:none"></div>
</body>

</html>