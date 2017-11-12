<!DOCTYPE html>  
<html>  
<head>  
<meta charset="UTF-8">  
<title>Insert title here</title>  
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css"/> 
  <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script> 
  <script type="text/javascript"> 

  $(function(){
    //判断访问终端
    var browser={
      versions:function(){
        var u = navigator.userAgent, app = navigator.appVersion;
        return {
            trident: u.indexOf('Trident') > -1, //IE内核
            presto: u.indexOf('Presto') > -1, //opera内核
            webKit: u.indexOf('AppleWebKit') > -1, //苹果、谷歌内核
            gecko: u.indexOf('Gecko') > -1 && u.indexOf('KHTML') == -1,//火狐内核
            mobile: !!u.match(/AppleWebKit.*Mobile.*/), //是否为移动终端
            ios: !!u.match(/\(i[^;]+;( U;)? CPU.+Mac OS X/), //ios终端//所以两个感叹号的作用就在于，如果明确设置了变量的值（非null/undifined/0/”“等值),结果就会根据变量的实际值来返回，如果没有设置，结果就会返回false。
            android: u.indexOf('Android') > -1 || u.indexOf('Adr') > -1, //android终端
            iPhone: u.indexOf('iPhone') > -1 , //是否为iPhone或者QQHD浏览器
            iPad: u.indexOf('iPad') > -1, //是否iPad
            webApp: u.indexOf('Safari') == -1, //是否web应该程序，没有头部与底部
            //webApp: ("standalone" in window.navigator)&&window.navigator.standalone, //是否web应该程序，没有头部与底部
            weixin: u.indexOf('MicroMessenger') > -1, //是否微信 （2015-01-22新增）
            qq: u.match(/\sQQ/i) == " qq" //是否QQ

        };
      }(),
      language:(navigator.browserLanguage || navigator.language).toLowerCase()
    }

    //判断是否IE内核
    if(browser.versions.trident){ alert("is IE"); }
    //判断是否webKit内核
    if(browser.versions.webKit){ alert("is webKit"); }
    //判断是否移动端
    if(browser.versions.mobile||browser.versions.android||browser.versions.ios){ alert("移动端"); }
    // 检测浏览器语言
    currentLang = navigator.language;   //判断除IE外其他浏览器使用语言
    if(!currentLang){//判断IE浏览器使用语言
        currentLang = navigator.browserLanguage;
    }
    alert(currentLang);
  })



        $(function(){
        var data = 
            [
              {
                text: "欢迎您~{{.Ip}}土豪", 
                selectable: true,
                id: '010',
              },
              {
                text: "系统设置",
                icon: "fa fa-tachometer icon",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                selectable: true,
                id: '01',
                selectable: false,
                // state: {
                  // checked: true,
                  // disabled: true,
                  // expanded: true,
                  // selected: true
                // },
                tags: ['available'],
                nodes: 
                [
                  { 
                    icon: "fa fa-cog",
                    text: "基本设置",
                    id: '011',
                    nodeId: '011'
                  },
                  { 
                    icon: "fa fa-align-right",
                    text: "分级目录",
                    id: '012',
                    nodeId: '012'
                  }, 
                  { 
                    icon: "fa fa-bug",
                    text: "搜索IP",
                    id: '013'
                  }
                ] 
              },
              {
                text: "权限管理",
                icon: "fa fa-balance-scale",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                selectable: true,
                id: '02',
                selectable: false,
                // state: {
                  // checked: true,
                  // disabled: true,
                  // expanded: true,
                  // selected: true
                // },
                tags: ['available'],
                nodes: 
                [
                  { icon: "fa fa-safari",
                    text: '系统权限',
                    id: '021'
                  },
                  { icon: "fa fa-navicon",
                    text: '项目权限',
                    id: '022'
                  }
                ]
              },
              {
                text: "账号管理",
                icon: "fa fa-users icon",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                selectable: true,
                id: '03',
                selectable: false,
                // state: {
                  // checked: true,
                  // disabled: true,
                  // expanded: true,
                  // selected: true
                // },
                tags: ['available'],
                nodes: 
                [
                  { icon: "fa fa-users",
                    text: '用户',
                    id: '031'
                  },
                  { icon: "fa fa-th",
                    text: 'IP地址段',
                    id: '032'
                  },
                  { icon: "fa fa-group",
                    text: '用户组',
                    id: '033'
                  },
                ]
              }, 
              {
                text: "项目设置",
                icon: "fa fa-list-alt icon",
                // selectedIcon: "glyphicon glyphicon-stop",
                href: "#node-1",
                selectable: true,
                id: '04',
                selectable: false,
                tags: ['available'],
                nodes: 
                [
                  { 
                    icon: "fa fa-edit",
                    text: "编辑",
                    id: '041'
                  },
                  { 
                    icon: "fa fa-copy",
                    text: "同步IP",
                    id: '042'
                  },
                  { 
                    icon: "fa fa-lock",
                    text: "项目权限",
                    id: '043'
                  }
                ]
              } 
            ]  
        function init(tree){  
                var $checkableTree = $('#treeview-checkable').treeview({  
                data: data,//tree,  
                showIcon: false,  
                showCheckbox: true,  
                showTags:true,  
                bootstrap2: false,  
                levels:5,  
                onNodeChecked: function(event, node) {  
                  $('#checkable-output').prepend('<p>' + node.text + ' was checked</p>');  
                },  
                onNodeUnchecked: function (event, node) {  
                  $('#checkable-output').prepend('<p>' + node.text + ' was unchecked</p>');  
                }  
              });  
  
              var findCheckableNodess = function() {  
                return $checkableTree.treeview('search', [ $('#input-check-node').val(), { ignoreCase: false, exactMatch: false } ]);  
              };  
              var checkableNodes = findCheckableNodess();  
  
              // Check/uncheck/toggle nodes  
              $('#input-check-node').on('keyup', function (e) {  
                checkableNodes = findCheckableNodess();  
                $('.check-node').prop('disabled', !(checkableNodes.length >= 1));  
              });  
  
              $('#btn-check-node.check-node').on('click', function (e) {  
                $checkableTree.treeview('checkNode', [ checkableNodes, { silent: $('#chk-check-silent').is(':checked') }]);  
              });  
  
              $('#btn-uncheck-node.check-node').on('click', function (e) {  
                $checkableTree.treeview('uncheckNode', [ checkableNodes, { silent: $('#chk-check-silent').is(':checked') }]);  
              });  
  
              $('#btn-toggle-checked.check-node').on('click', function (e) {  
                $checkableTree.treeview('toggleNodeChecked', [ checkableNodes, { silent: $('#chk-check-silent').is(':checked') }]);  
              });  
  
              // Check/uncheck all  
              $('#btn-check-all').on('click', function (e) {  
                $checkableTree.treeview('checkAll', { silent: $('#chk-check-silent').is(':checked') });  
              });  
  
              $('#btn-uncheck-all').on('click', function (e) {  
                $checkableTree.treeview('uncheckAll', { silent: $('#chk-check-silent').is(':checked') });  
              });  
            }  
  
            // var tree;   
            // $.ajax({  
            //       type:'post',  
            //       url:'/init.do',  
            //       success:function(data){  
                 init(data);                    
            //       }  
            //    });  
          
      });    
      
    function submit(){  
        var arr = new Array();  
        var $tree = $('#treeview-checkable');  
        arr = $tree.treeview('getChecked', 0);  
        var fristArr = new Array();  
        var secondArr = new Array();  
        var thirdArr = new Array();  
        var fourthArr = new Array();  
        if(arr.length==0){  
            alert("请选择节点");  
            return;  
        }

//获取节点所在的层次 

        for(var i=0 ; i< arr.length;i++){  
            var node = arr[i];  
            var temp = $('#treeview-checkable').treeview('getNode', node.nodeId);  
            var href = node.href;  
            if(href.indexOf("root")>-1){  
                var v=2;  
                continue;  
            }else{  
                if(node.parentId == 0){  
                    fristArr.push(href);  
                }else{  
                    if (($tree.treeview('getNode', node.parentId)).parentId == 0) {  
                        secondArr.push(href);  
                    } else if($tree.treeview('getNode', ($tree.treeview('getNode', node.parentId)).parentId).parentId == 0) {  
                        thirdArr.push(href);  
                    }else{  
                        fourthArr.push(href);  
                    }  
                }                 
            }  
        }  
        $.ajax({  
            type:'post',  
            traditional :true,   
            url:'gdupWeb/runtask.do',  
            data:{firstSpecIDs:fristArr,secondSpecIDs:secondArr,thirdSpecIDs:thirdArr,fourthSpecIDs:fourthArr,startdate:startdate,enddate:enddate},  
            success:function(data){  
                if (data=="1") {  
  
                }  
            }  
        });  
    }  

        $(function() {

        var defaultData = [
          {
            text: 'Parent 1',
            href: '#parent1',
            tags: ['4'],
            nodes: [
              {
                text: 'Child 1',
                href: '#child1',
                tags: ['2'],
                nodes: [
                  {
                    text: 'Grandchild 1',
                    href: '#grandchild1',
                    tags: ['0']
                  },
                  {
                    text: 'Grandchild 2',
                    href: '#grandchild2',
                    tags: ['0']
                  }
                ]
              },
              {
                text: 'Child 2',
                href: '#child2',
                tags: ['0']
              }
            ]
          },
          {
            text: 'Parent 2',
            href: '#parent2',
            tags: ['0']
          },
          {
            text: 'Parent 3',
            href: '#parent3',
             tags: ['0']
          },
          {
            text: 'Parent 4',
            href: '#parent4',
            tags: ['0']
          },
          {
            text: 'Parent 5',
            href: '#parent5'  ,
            tags: ['0']
          }
        ];

        var alternateData = [
          {
            text: 'Parent 1',
            tags: ['2'],
            nodes: [
              {
                text: 'Child 1',
                tags: ['3'],
                nodes: [
                  {
                    text: 'Grandchild 1',
                    tags: ['6']
                  },
                  {
                    text: 'Grandchild 2',
                    tags: ['3']
                  }
                ]
              },
              {
                text: 'Child 2',
                tags: ['3']
              }
            ]
          },
          {
            text: 'Parent 2',
            tags: ['7']
          },
          {
            text: 'Parent 3',
            icon: 'glyphicon glyphicon-earphone',
            href: '#demo',
            tags: ['11']
          },
          {
            text: 'Parent 4',
            icon: 'glyphicon glyphicon-cloud-download',
            href: '/demo.html',
            tags: ['19'],
            selected: true
          },
          {
            text: 'Parent 5',
            icon: 'glyphicon glyphicon-certificate',
            color: 'pink',
            backColor: 'red',
            href: 'http://www.tesco.com',
            tags: ['available','0']
          }
        ];

        var json = '[' +
          '{' +
            '"text": "Parent 1",' +
            '"nodes": [' +
              '{' +
                '"text": "Child 1",' +
                '"nodes": [' +
                  '{' +
                    '"text": "Grandchild 1"' +
                  '},' +
                  '{' +
                    '"text": "Grandchild 2"' +
                  '}' +
                ']' +
              '},' +
              '{' +
                '"text": "Child 2"' +
              '}' +
            ']' +
          '},' +
          '{' +
            '"text": "Parent 2"' +
          '},' +
          '{' +
            '"text": "Parent 3"' +
          '},' +
          '{' +
            '"text": "Parent 4"' +
          '},' +
          '{' +
            '"text": "Parent 5"' +
          '}' +
        ']';


        $('#treeview1').treeview({
          data: defaultData
        });

        $('#treeview2').treeview({
          levels: 1,
          data: defaultData
        });

        $('#treeview3').treeview({
          levels: 99,
          data: defaultData
        });

        $('#treeview4').treeview({

          color: "#428bca",
          data: defaultData
        });

        $('#treeview5').treeview({
          color: "#428bca",
          expandIcon: 'glyphicon glyphicon-chevron-right',
          collapseIcon: 'glyphicon glyphicon-chevron-down',
          nodeIcon: 'glyphicon glyphicon-bookmark',
          data: defaultData
        });

        $('#treeview6').treeview({
          color: "#428bca",
          expandIcon: "glyphicon glyphicon-stop",
          collapseIcon: "glyphicon glyphicon-unchecked",
          nodeIcon: "glyphicon glyphicon-user",
          showTags: true,
          data: defaultData
        });

        $('#treeview7').treeview({
          color: "#428bca",
          showBorder: false,
          data: defaultData
        });

        $('#treeview8').treeview({
          expandIcon: "glyphicon glyphicon-stop",
          collapseIcon: "glyphicon glyphicon-unchecked",
          nodeIcon: "glyphicon glyphicon-user",
          color: "yellow",
          backColor: "purple",
          onhoverColor: "orange",
          borderColor: "red",
          showBorder: false,
          showTags: true,
          highlightSelected: true,
          selectedColor: "yellow",
          selectedBackColor: "darkorange",
          data: defaultData
        });

        $('#treeview9').treeview({
          expandIcon: "glyphicon glyphicon-stop",
          collapseIcon: "glyphicon glyphicon-unchecked",
          nodeIcon: "glyphicon glyphicon-user",
          color: "yellow",
          backColor: "purple",
          onhoverColor: "orange",
          borderColor: "red",
          showBorder: false,
          showTags: true,
          highlightSelected: true,
          selectedColor: "yellow",
          selectedBackColor: "darkorange",
          data: alternateData
        });

        $('#treeview10').treeview({
          color: "#428bca",
          enableLinks: true,
          data: defaultData
        });



        var $searchableTree = $('#treeview-searchable').treeview({
          data: defaultData,
        });

        var search = function(e) {
          var pattern = $('#input-search').val();
          var options = {
            ignoreCase: $('#chk-ignore-case').is(':checked'),
            exactMatch: $('#chk-exact-match').is(':checked'),
            revealResults: $('#chk-reveal-results').is(':checked')
          };
          var results = $searchableTree.treeview('search', [ pattern, options ]);

          var output = '<p>' + results.length + ' matches found</p>';
          $.each(results, function (index, result) {
            output += '<p>- ' + result.text + '</p>';
          });
          $('#search-output').html(output);
        }

        $('#btn-search').on('click', search);
        $('#input-search').on('keyup', search);

        $('#btn-clear-search').on('click', function (e) {
          $searchableTree.treeview('clearSearch');
          $('#input-search').val('');
          $('#search-output').html('');
        });


        var initSelectableTree = function() {
          return $('#treeview-selectable').treeview({
            data: defaultData,
            multiSelect: $('#chk-select-multi').is(':checked'),
            onNodeSelected: function(event, node) {
              $('#selectable-output').prepend('<p>' + node.text + ' was selected</p>');
            },
            onNodeUnselected: function (event, node) {
              $('#selectable-output').prepend('<p>' + node.text + ' was unselected</p>');
            }
          });
        };
        var $selectableTree = initSelectableTree();

        var findSelectableNodes = function() {
          return $selectableTree.treeview('search', [ $('#input-select-node').val(), { ignoreCase: false, exactMatch: false } ]);
        };
        var selectableNodes = findSelectableNodes();

        $('#chk-select-multi:checkbox').on('change', function () {
          console.log('multi-select change');
          $selectableTree = initSelectableTree();
          selectableNodes = findSelectableNodes();          
        });

        // Select/unselect/toggle nodes
        $('#input-select-node').on('keyup', function (e) {
          selectableNodes = findSelectableNodes();
          $('.select-node').prop('disabled', !(selectableNodes.length >= 1));
        });

        $('#btn-select-node.select-node').on('click', function (e) {
          $selectableTree.treeview('selectNode', [ selectableNodes, { silent: $('#chk-select-silent').is(':checked') }]);
        });

        $('#btn-unselect-node.select-node').on('click', function (e) {
          $selectableTree.treeview('unselectNode', [ selectableNodes, { silent: $('#chk-select-silent').is(':checked') }]);
        });

        $('#btn-toggle-selected.select-node').on('click', function (e) {
          $selectableTree.treeview('toggleNodeSelected', [ selectableNodes, { silent: $('#chk-select-silent').is(':checked') }]);
        });



        var $expandibleTree = $('#treeview-expandible').treeview({
          data: defaultData,
          onNodeCollapsed: function(event, node) {
            $('#expandible-output').prepend('<p>' + node.text + ' was collapsed</p>');
          },
          onNodeExpanded: function (event, node) {
            $('#expandible-output').prepend('<p>' + node.text + ' was expanded</p>');
          }
        });

        var findExpandibleNodess = function() {
          return $expandibleTree.treeview('search', [ $('#input-expand-node').val(), { ignoreCase: false, exactMatch: false } ]);
        };
        var expandibleNodes = findExpandibleNodess();

        // Expand/collapse/toggle nodes
        $('#input-expand-node').on('keyup', function (e) {
          expandibleNodes = findExpandibleNodess();
          $('.expand-node').prop('disabled', !(expandibleNodes.length >= 1));
        });

        $('#btn-expand-node.expand-node').on('click', function (e) {
          var levels = $('#select-expand-node-levels').val();
          $expandibleTree.treeview('expandNode', [ expandibleNodes, { levels: levels, silent: $('#chk-expand-silent').is(':checked') }]);
        });

        $('#btn-collapse-node.expand-node').on('click', function (e) {
          $expandibleTree.treeview('collapseNode', [ expandibleNodes, { silent: $('#chk-expand-silent').is(':checked') }]);
        });

        $('#btn-toggle-expanded.expand-node').on('click', function (e) {
          $expandibleTree.treeview('toggleNodeExpanded', [ expandibleNodes, { silent: $('#chk-expand-silent').is(':checked') }]);
        });

        // Expand/collapse all
        $('#btn-expand-all').on('click', function (e) {
          var levels = $('#select-expand-all-levels').val();
          $expandibleTree.treeview('expandAll', { levels: levels, silent: $('#chk-expand-silent').is(':checked') });
        });

        $('#btn-collapse-all').on('click', function (e) {
          $expandibleTree.treeview('collapseAll', { silent: $('#chk-expand-silent').is(':checked') });
        });



        var $checkableTree = $('#treeview-checkable').treeview({
          data: defaultData,
          showIcon: false,
          showCheckbox: true,
          onNodeChecked: function(event, node) {
            $('#checkable-output').prepend('<p>' + node.text + ' was checked</p>');
          },
          onNodeUnchecked: function (event, node) {
            $('#checkable-output').prepend('<p>' + node.text + ' was unchecked</p>');
          }
        });

        var findCheckableNodess = function() {
          return $checkableTree.treeview('search', [ $('#input-check-node').val(), { ignoreCase: false, exactMatch: false } ]);
        };
        var checkableNodes = findCheckableNodess();

        // Check/uncheck/toggle nodes
        $('#input-check-node').on('keyup', function (e) {
          checkableNodes = findCheckableNodess();
          $('.check-node').prop('disabled', !(checkableNodes.length >= 1));
        });

        $('#btn-check-node.check-node').on('click', function (e) {
          $checkableTree.treeview('checkNode', [ checkableNodes, { silent: $('#chk-check-silent').is(':checked') }]);
        });

        $('#btn-uncheck-node.check-node').on('click', function (e) {
          $checkableTree.treeview('uncheckNode', [ checkableNodes, { silent: $('#chk-check-silent').is(':checked') }]);
        });

        $('#btn-toggle-checked.check-node').on('click', function (e) {
          $checkableTree.treeview('toggleNodeChecked', [ checkableNodes, { silent: $('#chk-check-silent').is(':checked') }]);
        });

        // Check/uncheck all
        $('#btn-check-all').on('click', function (e) {
          $checkableTree.treeview('checkAll', { silent: $('#chk-check-silent').is(':checked') });
        });

        $('#btn-uncheck-all').on('click', function (e) {
          $checkableTree.treeview('uncheckAll', { silent: $('#chk-check-silent').is(':checked') });
        });



        var $disabledTree = $('#treeview-disabled').treeview({
          data: defaultData,
          onNodeDisabled: function(event, node) {
            $('#disabled-output').prepend('<p>' + node.text + ' was disabled</p>');
          },
          onNodeEnabled: function (event, node) {
            $('#disabled-output').prepend('<p>' + node.text + ' was enabled</p>');
          },
          onNodeCollapsed: function(event, node) {
            $('#disabled-output').prepend('<p>' + node.text + ' was collapsed</p>');
          },
          onNodeUnchecked: function (event, node) {
            $('#disabled-output').prepend('<p>' + node.text + ' was unchecked</p>');
          },
          onNodeUnselected: function (event, node) {
            $('#disabled-output').prepend('<p>' + node.text + ' was unselected</p>');
          }
        });

        var findDisabledNodes = function() {
          return $disabledTree.treeview('search', [ $('#input-disable-node').val(), { ignoreCase: false, exactMatch: false } ]);
        };
        var disabledNodes = findDisabledNodes();

        // Expand/collapse/toggle nodes
        $('#input-disable-node').on('keyup', function (e) {
          disabledNodes = findDisabledNodes();
          $('.disable-node').prop('disabled', !(disabledNodes.length >= 1));
        });

        $('#btn-disable-node.disable-node').on('click', function (e) {
          $disabledTree.treeview('disableNode', [ disabledNodes, { silent: $('#chk-disable-silent').is(':checked') }]);
        });

        $('#btn-enable-node.disable-node').on('click', function (e) {
          $disabledTree.treeview('enableNode', [ disabledNodes, { silent: $('#chk-disable-silent').is(':checked') }]);
        });

        $('#btn-toggle-disabled.disable-node').on('click', function (e) {
          $disabledTree.treeview('toggleNodeDisabled', [ disabledNodes, { silent: $('#chk-disable-silent').is(':checked') }]);
        });

        // Expand/collapse all
        $('#btn-disable-all').on('click', function (e) {
          $disabledTree.treeview('disableAll', { silent: $('#chk-disable-silent').is(':checked') });
        });

        $('#btn-enable-all').on('click', function (e) {
          $disabledTree.treeview('enableAll', { silent: $('#chk-disable-silent').is(':checked') });
        });



        var $tree = $('#treeview12').treeview({
          data: json
        });
        });

    </script>  
</head>  
<body>  
    <div id="treeview-checkable" class=""></div>
<!-- </div>   -->
</body>  
</html>