<!DOCTYPE html>
<html>
<head>
  <meta charset="utf-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Demo</title>
  <!-- 请勿在项目正式环境中引用该 layui.css 地址 -->
  <link href="//unpkg.com/layui@2.9.8/dist/css/layui.css" rel="stylesheet">
</head>
<body>
<table class="layui-hide" id="ID-treeTable-demo"></table>
<script type="text/html" id="TPL-treeTable-demo">
  <div class="layui-btn-container">
    <button class="layui-btn layui-btn-sm" lay-event="getChecked">获取选中数据</button>
  </div>
</script>
<script type="text/html" id="TPL-treeTable-demo-tools">
  <div class="layui-btn-container">
    <a class="layui-btn layui-btn-primary layui-btn-xs" lay-event="detail">查看</a>
    <a class="layui-btn layui-btn-warm layui-btn-xs" lay-event="addChild">新增</a>
    <a class="layui-btn layui-btn-xs" lay-event="more">更多 <i class="layui-icon layui-icon-down"></i></a>
  </div>
</script>
  
<!-- 请勿在项目正式环境中引用该 layui.js 地址 -->
<script src="//unpkg.com/layui@2.9.8/dist/layui.js"></script>
<script>
layui.use(function(){
  var treeTable = layui.treeTable;
  var layer = layui.layer;
  var dropdown = layui.dropdown;
  // 渲染
  var inst = treeTable.render({
    elem: '#ID-treeTable-demo',
    url: '/v1/estimate/getestimatecostmetaldata/{{.PhaseID}}', // 此处为静态模拟数据，实际使用时需换成真实接口
    id: 'ID', // 自定义 id 索引
    tree: {
      /*
      // 异步加载子节点
      async: {
        enable: true,
        url: '/static/json/2/treeTable/demo-async.json', // 此处为静态模拟数据，实际使用时需换成真实接口
        autoParam: ["parentId=id"]
      }
      */
      // id: 'ID', // 自定义 id 索引
    },
    // maxHeight: '800px',
    toolbar: '#TPL-treeTable-demo',
    cols: [[
      {type: 'checkbox', fixed: 'left'},
      {field: 'ID', title: 'ID', width: 20, sort: true, fixed: 'left'},
      {
        field: 'number',
        title: '编号',
        halign: "center",
        align: "center",
        valign: "middle",
        width: '20',
      }, {
        field: 'name',
        title: '工程或费用名称',
        halign: "center",
        align: "left",
        valign: "middle",
        width: '600'
      }, {
        field: 'unit',
        title: '单位',
        halign: "center",
        align: "center",
        valign: "middle",
        width: '20',
      }, {
        field: 'quantity',
        title: '数量',
        halign: "center",
        align: "center",
        valign: "middle",
        width: '120',
      }, {
        field: 'unitpriceequipment',
        title: '设备单价(元)',
        halign: "center",
        align: "center",
        valign: "middle",
        width: '120',
      }, {
        field: 'unitpriceinstallation',
        title: '安装单价(元)',
        halign: "center",
        align: "center",
        valign: "middle",
        width: '120',
      }, {
        field: 'totalequipment',
        title: '设备合计(万元)',
        halign: "center",
        align: "center",
        valign: "middle",
        width: '120',
      }, {
        field: 'totalinstallation',
        title: '安装合计(万元)',
        halign: "center",
        align: "center",
        valign: "middle",
        width: '120',
      },
        // {
        //   field: 'parentId',
        //   title: '节点',
        //   halign: "center",
        //   align: "left",
        //   valign: "middle",
        //   width: '100'
        // },
      // {field: 'name', title: '用户名', width: 180, fixed: 'left'},
      // {field: 'sex', title: '性别', width: 80, sort: true},
      // {field: 'experience', title: '积分', width: 90, sort: true},
      // {field: 'city', title: '城市', width: 100},
      { fixed: "right", title: "操作", width: 190, align: "center", toolbar: "#TPL-treeTable-demo-tools"} 
    ]],
    page: true,
    done: function(res, curr, count, origin){
      console.log(res); // 得到当前渲染的数据
      console.log(curr);  // 得到当前页码
      console.log(count); // 得到数据总量
      console.log(origin); // 回调函数所执行的来源 --- 2.8.7+
      // 展开或关闭全部节点
      treeTable.expandAll('ID', true); // 关闭全部节点
    },
    // done: function(){treeTable.expandAll('tet', true);}
    
  });
  // 展开或关闭对应节点
  // treeTable.expandNode('name', {
  //   index: 0, // 第一行
  //   expandFlag: true // 展开
  // })
  
  // 表头工具栏工具事件
  treeTable.on("toolbar(ID-treeTable-demo)", function (obj) {
    var config = obj.config;
    var tableId = config.id;
    var status = treeTable.checkStatus(tableId);
    // 获取选中行
    if (obj.event === "getChecked") {
      if(!status.data.length) return layer.msg('无选中数据');
      console.log(status);
      layer.alert("当前数据选中已经输出到控制台，<br>您可按 F12 从控制台中查看结果。");
    }
  });
  // 单元格工具事件
  treeTable.on('tool('+ inst.config.id +')', function (obj) {
    var layEvent = obj.event; // 获得 lay-event 对应的值
    var trElem = obj.tr;
    var trData = obj.data;
    var tableId = obj.config.id;
    if (layEvent === "detail") {
      layer.msg("查看操作：" + trData.name);
    } else if (layEvent === "addChild") {
      var data = { id: Date.now(), name: "新节点" };
      var newNode2 = treeTable.addNodes(tableId, {
        parentIndex: trData["LAY_DATA_INDEX"], 
        index: -1, 
        data: data
      });
    } else if (layEvent === "more") {
      // 下拉菜单
      dropdown.render({
        elem: this, // 触发事件的 DOM 对象
        show: true, // 外部事件触发即显示
        align: "right", // 右对齐弹出
        data: [
          {
            title: "修改积分",
            id: "edit"
          },
          {
            title: "删除",
            id: "del"
          }
        ],
        click: function (menudata) {
          if (menudata.id === "del") {
            layer.confirm("真的删除行么", function (index) {
              obj.del(); // 等效如下
              // treeTable.removeNode(tableId, trElem.attr('data-index'))
              layer.close(index);
            });
          } else if (menudata.id === "edit") {
            layer.prompt({
                value: trData.experience,
                title: "输入新的积分"
            }, function (value, index) {
              obj.update({ experience: value }); // 等效如下
              // treeTable.updateNode(tableId, trElem.attr('data-index'), {experience: value});
              layer.close(index);
            });
          }
        }
      });
    }
  });
});
</script>

</body>
</html>