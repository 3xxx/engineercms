<!doctype html>
<html>

<head>
  <meta charset="utf-8">
  <title>viewer</title>
  <link rel="stylesheet" href="https://developer.api.autodesk.com/modelderivative/v2/viewers/6.0/style.min.css" type="text/css">
  <!-- <script src="js/jquery.min.js" type="text/javascript" charset="utf-8"></script> -->
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script src="https://cdn.bootcss.com/three.js/r83/three.min.js" type="text/javascript" charset="utf-8"></script>
  <script src="https://developer.api.autodesk.com/modelderivative/v2/viewers/6.0/viewer3D.min.js"></script>
  <style type="text/css">
  body {
    margin: 0;
  }

  #MyViewerDiv {
    width: 100%;
    height: 100%;
    margin: 0;
    background-color: #f0f8ff;
  }
  </style>
</head>

<body>
  <div id="MyViewerDiv"></div>
  <script>
  var viewer;

  var model_viewer = null;

  var options = {
    env: 'Local',
    // model_src: 'http://192.168.1.25/mbcadmin/public/uploads/20180612/9b5d987e06c92c0ee8c6d78a0847ac76/3d.svf' // 一个svf的本地地址，这个地址不可为相对路径，不然是找不到这个路径的
    model_src: 'http://127.0.0.1/static/download/drawing1.dwg'
  };

  model_viewer = ViewerWithTool("#MyViewerDiv", true) // With toolbar

  Autodesk.Viewing.Initializer(options, function() {
    model_viewer.initialize()

    let btnNew = new Autodesk.Viewing.UI.Button('btnNew'); //新增一个按钮
    btnNew.onClick = function(e) {
      //给按钮添加点击事件
    }
    btnNew.setToolTip('新的按钮'); //鼠标放置按钮上方时的显示
    btnNew.container.children[0].innerHTML = '若'; //按钮上显示的内容

    let subToolbar = new Autodesk.Viewing.UI.ControlGroup('custom-toolbar'); //按钮组

    subToolbar.addControl(btnNew); //将新的按钮放入到新增的按钮组中
    model_viewer.toolbar.addControl(subToolbar); //将按钮组添加至控制栏中


    model_viewer.loadModel(options.model_src)
  })

  // 是否包含工具栏加载
  function ViewerWithTool(el, hasTool) {
    if (hasTool) {
      return (new Autodesk.Viewing.Private.GuiViewer3D($(el)[0], {}))
    } else {
      return (new Autodesk.Viewing.Viewer3D($(el)[0], {}))
    }
  }
  </script>
</body>

</html>