<!DOCTYPE html>
<html>

<head>
  <title>FreeCAD Model</title>
  <script crossorigin src="https://unpkg.com/react@17/umd/react.development.js"></script>
  <script crossorigin src="https://unpkg.com/react-dom@17/umd/react-dom.development.js"></script>
  <script src="https://unpkg.com/@babel/standalone/babel.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/moment.js/2.20.1/moment.min.js"></script>
  <script src="https://cdnjs.cloudflare.com/ajax/libs/antd/4.18.2/antd.min.js"></script>
  <link href="https://cdn.jsdelivr.net/npm/ant-design-icons/dist/anticons.min.css" rel="stylesheet">
  <link rel="stylesheet" type="text/css" href="https://cdnjs.cloudflare.com/ajax/libs/antd/4.18.2/antd.compact.min.css">

  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" />
  <!--   <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css" />
  <link href="/static/css/bootstrap-treeview.css" rel="stylesheet">
  <link rel="stylesheet" href="/static/css/jquery.mCustomScrollbar.min.css">
  <link rel="stylesheet" href="/static/css/custom.css">
  <link rel="stylesheet" type="text/css" href="/static/font-awesome-4.7.0/css/font-awesome.min.css" /> -->
  <style type="text/css">
  #app {
    height: 100vh;
  }

  .page-wrapper .sidebar-wrapper {
    position: fixed;
    /*top: 50px;*/
    left: 5;
    /*bottom: 0;*/
    /*width: 200px;*/
    /*background-color: #23262E;*/
  }

  .page-wrapper .container-fluid {
    /*position: fixed;*/
    /*top: 50px;*/
    right: 0;
    /*bottom: 0;*/
    left: 200px;
    /*background-color: #F8F9FA;*/
    overflow: hidden;
    height: 100vh;
  }

  .site-tree-search-value {
    color: #f50;
  }

  .ant-layout-sider {
    position: absolute;
    background-color: #ed5565;
    box-shadow: 4px 6px 14px 0px #383838
  }

  .collapsed-icon {
    position: absolute;
    top: 0px;
    right: -24px;
    font-size: 20px;
  }

  .input-container {
    width: 100%;
    padding: 8px 24px;
  }

  .input-tree {
    min-width: auto;
    width: 100%;
    border-radius: 30px;
  }

  .tree-wrapper {
    background-color: #383838;
    border-bottom: 1px solid #fff;
    border-radius: 1px solid #fff;
  }

  .ant-tree-treenode {
    height: 40px;
    background-color: #383838;
    color: #fff;
    width: 100%;
    border: 1px solid #fff;
    border-left: none;
    border-bottom: 0;
    border-right: 0;
    display: flex;
    align-items: center;
    overflow: auto;
  }

  .ant-tree-treenode:hover {
    background-color: #428bca;
    /*#81B0D9;*/
  }

  .ant-tree-treenode-selected {
    background-color: #428bca;
  }

  .ant-tree .ant-tree-treenode {
    align-items: center;
  }

  .ant-tree-switcher {
    display: flex;
    align-items: center;
  }

  .ant-tree-node-content-wrapper {
    width: 100%;
    height: 100%;
    display: flex;
    align-items: center;
    padding: 0;
  }

  .ant-tree .ant-tree-node-content-wrapper:hover {
    background-color: transparent;
  }

  .ant-tree .ant-tree-node-content-wrapper.ant-tree-node-selected {
    background-color: transparent;
  }

  .ant-tree-title {
    font-size: 16px;
  }

  .link {
    color: #fff;
    display: flex;
    flex-direction: column;
  }

  .link-item {
    height: 40px;
    display: flex;
    align-items: center;
    border-bottom: 1px solid #d1565e;
    padding: 0 20px;
    cursor: pointer;
  }

  .link-item:hover {
    background-color: #428bca;
  }

  .ant-tree.ant-tree-directory .ant-tree-treenode:hover:before {
    background-color: #428bca;
  }
  </style>
</head>

<body>
  <!-- <div id="container" style="padding: 24px" />
  <script>
    const mountNode = document.getElementById('container');
  </script> -->
  <div class="page-wrapper toggled">
    <!-- <iframe id="iframe" name="frame" src="2.html" width="100%" height="1000px"   frameborder="no" border="0" marginwidth="0" marginheight="0" scrolling="no" allowtransparency="yes"></iframe> -->
    <div id="app"></div>
    <nav id="sidebar" class="sidebar-wrapper">
      <div class="sidebar-content mCustomScrollbar _mCS_1 mCS-autoHide desktop">
        <div id="mCSB_1" class="mCustomScrollBox mCS-light mCSB_vertical mCSB_inside" tabindex="0" style="max-height: none;">
        </div>
      </div>
    </nav>
    <main class="page-content">
    </main>
  </div>
  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <script type="text/javascript" src="/static/js/bootstrap.min.js"></script>
  <script src="/static/js/bootstrap-treeview.js"></script>
  <script src="/static/js/jquery.mCustomScrollbar.concat.min.js"></script>
  <script src="/static/js/custom.js"></script>
  <script type="text/javascript">
  // $(function() {
  // $('#tree').treeview({
  //   data: [{{.json }}], //{{.json }},// 数据结构供bootstrap treeview用
  //   data: getTree(),// 数据结构供antd tree测试用
  //   levels: 1,
  //   showTags: true,
  //   loadingIcon: "fa fa-minus",
  //   lazyLoad: loaddata,
  // });

  // $('#tree').on('nodeSelected', function(event, data) {
  //   document.getElementById("iframepage").src = "/v1/freecad/getfreecad/" + data.id;
  // });
  // })

  function getTree() {
    // Some logic to retrieve, or generate tree structure
    $.ajax({
      type: "get",
      url: "/v1/freecad/freecadmenu",
      success: function(data, status) {
        return data;
      }
    });
  }

  function loaddata(node, func) {
    $.ajax({
      type: "get",
      url: "/v1/freecad/freecadlist",
      data: { id: node.id },
      success: function(data, status) {
        if (data) {
          func(data);
        }
      }
    });
  }

  function reinitIframe() {
    var iframe = document.getElementById("iframepage");
    try {
      var bHeight = iframe.contentWindow.document.body.scrollHeight;
      var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
      // var height = Math.max(bHeight, dHeight, 800);
      var height = '100%';

      iframe.height = height;
    } catch (ex) {}
  }
  window.setInterval("reinitIframe()", 200);
  </script>
  <!--   <script type="text/babel">
    // const { createRoot } = ReactDOM;
    const {  Tree  } = antd;
    const { DirectoryTree } = Tree;
    const treeData = [
      {
          "id": 1,
          "key": "1",
          "tags": [
            ""
          ],
          "lazyLoad": false,
          "title": "0001-0100",
          "children": [
            {
              "id": 2,
              "key": "2",
              "tags": [
                ""
              ],
              "lazyLoad": false,
              "title": "FC0001带法兰岔管",
              "children": null
            },
            {
              "id": 3,
              "key": "3",
              "tags": [
                ""
              ],
              "lazyLoad": false,
              "title": "FC0002相贯线岔管",
              "children": null
            }
          ]
        },
        {
          "id": 4,
          "key": "4",
          "tags": [
            ""
          ],
          "lazyLoad": false,
          "title": "0101-0200",
          "children": [
            {
              "id": 5,
              "key": "5",
              "tags": [
                ""
              ],
              "lazyLoad": false,
              "title": "FC0001带法兰岔管",
              "children": null
            },
            {
              "id": 6,
              "key": "6",
              "tags": [
                ""
              ],
              "lazyLoad": false,
              "title": "FC0002相贯线岔管",
              "children": null
            }
          ]
        },
        {
          "id": 7,
          "key": "7",
          "tags": [
            ""
          ],
          "lazyLoad": false,
          "title": "0101-0200",
          "children": null
        },
        {
          "id": 8,
          "key": "8",
          "tags": [
            ""
          ],
          "lazyLoad": false,
          "title": "0101-0200",
          "children": null
        },
        {
          "id": 9,
          "key": "9",
          "tags": [
            ""
          ],
          "lazyLoad": false,
          "title": "0101-0200",
          "children": null
        }
    ];
    const App = () => {

      const onSelect = (keys, info) => {
        console.log('Trigger Select', keys, info);
      };
      const onExpand = (keys, info) => {
        console.log('Trigger Expand', keys, info);
      };
      return (
        <DirectoryTree
          // multiple
          // defaultExpandAll
          onSelect={onSelect}
          onExpand={onExpand}
          treeData={treeData}
        />
      );

    };
    const ComponentDemo = App;
    
 
    ReactDOM.render(
      <ComponentDemo />,
      document.getElementById("container")
    )

    fetch(`/v1/freecad/freecadmenu`,{
          method:'GET'
        }).then(response => response.json())
          .then(data => {
            defaultData = data;
            console.log(data);
          });
  </script> -->
  <script type="text/babel">
    const SiderBox = () => {
      const { Tree, Layout, Input } = antd;

      const { DirectoryTree } = Tree;
      const { Header, Sider, Content } = antd.Layout;

      const [collapsed, setCollapsed] = React.useState(false);
      const [expandedKeys, setExpandedKeys] = React.useState([]);
      const [searchValue, setSearchValue] = React.useState('');
      const [autoExpandParent, setAutoExpandParent] = React.useState(true);

      const itemSytle ={marginRight: '10px', fontSize: '18px'}
      const userStyle = {fontSize:'16px'}

      var defaultData = []
      $.ajax({
        type: "get",
        async: false,
        url: "/v1/freecad/freecadmenu",
        success: function(data, status) {
          defaultData[0] = data
        }
      });

      const dataList = [];
      const generateList = (data) => {
        data.forEach(function (item) {
          const node = item;
          const { key, title } = node;
          dataList.push({
            key,
            title: title,
          });
          if (node.children) {
            generateList(node.children);
          }
        });
      };

      generateList(defaultData);
      // console.log(dataList);

      const treeData = React.useMemo(() => {
        const loop = (data) =>
          data.map((item) => {
            const strTitle = item.title;
            const index = strTitle.indexOf(searchValue);
            const beforeStr = strTitle.substring(0, index);
            const afterStr = strTitle.slice(index + searchValue.length);
            const title =
              index > -1 ? (
                <span>
                  {beforeStr}
                  <span className="site-tree-search-value">{searchValue}</span>
                  {afterStr}
                </span>
              ) : (
                <span>{strTitle}</span>
              );
            if (item.children) {
              return {
                title,
                key: item.key,
                children: loop(item.children),
              };
            }
            return {
              title,
              key: item.key,
            };
          });
        return loop(defaultData);
      }, [searchValue]);

      const getParentKey = (key, tree) => {
        let parentKey;
        tree.forEach(function (item2) {
          const node = item2;
          if (node.children) {
            if (node.children.some((item) => item.key === key)) {
              parentKey = node.key;
            } else if (getParentKey(key, node.children)) {
              parentKey = getParentKey(key, node.children);
            }
          }

        })
        return parentKey;
      };

      const onExpand = (newExpandedKeys) => {
        setExpandedKeys(newExpandedKeys);
        setAutoExpandParent(false);
      };

      const onChange = (e) => {
        const { value } = e.target;
        const newExpandedKeys = dataList
          .map(item => {
            if (item.title.indexOf(value) > -1) {
              return getParentKey(item.key, defaultData);
            }
            return null;
          })
          .filter((item, i, self) => item && self.indexOf(item) === i);
        setExpandedKeys(newExpandedKeys);
        setSearchValue(value);
        setAutoExpandParent(true);
      };

      const onSelect = (keys, info) => {
        // console.log('Trigger Select', keys, info);
        if (info.node.children){
          if (info.node.children.length==0){
            document.getElementById("iframepage").src = "/v1/freecad/getfreecad/" + keys[0];
          }else{
            console.log('is父级')
          }
        }else{
          console.log('is子级')
          document.getElementById("iframepage").src = "/v1/freecad/getfreecad/" + keys[0];
        }
      };

      const loop = (data) =>
        data
          .map((item) => {
            const index = item.title.indexOf(searchValue);
            const beforeStr = item.title.substr(0, index);
            const afterStr = item.title.substr(index + searchValue.length);
            const title =
              index > -1 ? (
                <span>
                  {beforeStr}
                  <span className="site-tree-search-value">{searchValue}</span>
                  {afterStr}
                </span>
              ) : (
                <span>{item.title}</span>
              );
            if (item.children && item.children.length > 0) {
              const children = loop(item.children);
              return {
                ...item,
                title: index > -1 || children.length ? title : null,
                children: children.length ? children : undefined
              };
            }
            return index > -1
              ? {
                  ...item,
                  title
                }
              : {
                  ...item,
                  title: null
                };
          })
        .filter((item) => !!item.title);

      const openupload = (data) => {
         window.open('/v1/freecad/uploadmodel');
      }
      const openwater = (data) => {
         window.open('https://zsj.itdos.net/docs/freecad');
      }
      return (
        <Layout>
          <Content>
            <Sider width={256}
              height={500}
              collapsible collapsed={collapsed}
              onCollapse={
                function(value){
                  return setCollapsed(value)
                }
              }
              trigger={null}>
              <div className="input-container">
                <Input placeholder=""
                  className="input-tree"
                  onChange={onChange} />
              </div>


              <DirectoryTree
                height={600}
                className='tree-wrapper'
                onExpand={onExpand}
                onSelect={onSelect}
                //treeData={treeData}高亮，但不过滤
                treeData={loop(defaultData)}
                //showLine={true}节点之间带连接线
                expandedKeys={expandedKeys}
                autoExpandParent={autoExpandParent}
                ></DirectoryTree>
        
              <div className="link">
                <span className="link-item" style={userStyle} onClick={
                  function(){
                    return openupload()
                  }
                }><i className="fa fa-upload" style={itemSytle}></i>上传</span>

                <span className="link-item" style={userStyle} onClick={
                  function(){
                    return opendownload()
                  }
                }><i className="fa fa-download" style={itemSytle}></i>下载</span>

                <span className="link-item" style={userStyle} onClick={
                  function(){
                    return openwater()
                  }
                }><i className="fa fa-users" style={itemSytle}></i>水务设计</span>
                
              </div>

              <span className='collapsed-icon' onClick={
                function(){
                  return setCollapsed(!collapsed)
                }
              }>{collapsed ? <i className="ai-menu-unfold"></i> : <i className="ai-menu-fold"></i>}</span>
            </Sider>

            <div className="container-fluid">
              <iframe
                src="/v1/freecad/getfreecad/{{.ModelId}}"
                name='iframepage'
                id="iframepage"
                frameBorder="0"
                width="100%"
                height="100%"
                scrolling="no"
              ></iframe>
            </div>

          </Content>
        </Layout>
      );
    };

    ReactDOM.render(
      <SiderBox />,
      document.getElementById("app")
    )
  </script>
</body>