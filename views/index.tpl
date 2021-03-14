<!-- 这个是显示左侧栏，右边index_user显示用户的cms情况 -->
<!DOCTYPE html>
{{template "header"}}
<title>EngineerCMS</title>
<script src="/static/js/echarts.min.js"></script>
<script src="/static/js/china.js"></script>
<!-- <script src="/static/js/world.js"></script>   -->
<!-- <script src="/static/js/jquery.min223.js"></script>   -->
<script type="text/javascript" src="https://api.map.baidu.com/api?v=2.0&ak=hYCENCEx1nXO0Nt46ldexfG9oI49xBGh"></script>
<script src="/static/js/echarts-gl.min.js "></script>
<script src="/static/js/bmap.min.js"></script>
<script type="text/javascript" src="/static/js/bootstrap-treeview.js"></script>
<link rel="stylesheet" type="text/css" href="/static/css/bootstrap-treeview.css" />
</head>
<!-- <div class="navbar navba-default navbar-fixed-top"> -->
<div class="container-fill">{{template "navbar" .}}</div>
<!-- </div> -->

<body>
  <!-- <div id="main" style="height: 400"></div> -->
  <!-- 侧栏 -->
  <div id="treeview" class="col-xs-2 col-sm-2 col-md-2 col-lg-2"></div>
  <!-- 右侧frame -->
  <div class="col-xs-10 col-sm-10 col-md-10 col-lg-10">
    <iframe src="/index/user" name='main' frameborder="0" width="100%" scrolling="no" marginheight="0" marginwidth="0" id="iframepage" onload="this.height=100"></iframe>
  </div>
  <div id="footer">
    <div class="col-lg-12">
      <br>
      <hr />
      <div class="col-xs-6 col-sm-6 col-md-6 col-lg-6">
        <h4>Copyright © 2016-2021 EngineerCMS</h4>
        <p>
          网站由 <i class="user icon"></i>
          <a target="_blank" href="https://github.com/3xxx">@3xxx</a>
          建设，并由
          <a target="_blank" href="http://golang.org">golang</a>
          和
          <a target="_blank" href="http://beego.me">beego</a>
          提供动力。
          <a target="_blank" href="https://github.com/3xxx" target="_blank"><strong><i class="icon-github-sign"></i> Github</strong></a>
          - {{.PageStartTime|loadtimes}}ms.
        </p>
        <p>
          请给 <i class="glyphicon glyphicon-envelope"></i>
          <a class="email" href="mailto:504284@qq.com">我</a>
          发送反馈信息或提交
          <i class="tasks icon"></i>
          <a target="_blank" href="https://github.com/3xxx/engineercms/issues">网站问题</a>
          。
        </p>
      </div>
      <div class="col-xs-6 col-sm-6 col-md-6 col-lg-6">
        <h4>更多项目</h4>
        <div>
          <p>
            <a href="https://github.com/3xxx/hydrows">HydroWS-供水管线设计工具</a>
          </p>
          <p>
            <a href="https://github.com/3xxx/merit">MeritMS-价值成果管理系统</a>
          </p>
        </div>
      </div>
    </div>
  </div>
  <script type="text/javascript">
  function reinitIframe() { //http://caibaojian.com/frame-adjust-content-height.html
    var iframe = document.getElementById("iframepage");
    try {
      var bHeight = iframe.contentWindow.document.body.scrollHeight;
      var dHeight = iframe.contentWindow.document.documentElement.scrollHeight;
      var height = Math.max(bHeight, dHeight);
      iframe.height = height;
      // console.log(height);//这个显示老是在变化
    } catch (ex) {}
  }
  window.setInterval("reinitIframe()", 200);
  </script>
  <script type="text/javascript">
  $(function() {
    $('#treeview').treeview({
      data: {{.json }}, //[{{.json}}]——有时候加个中括号就行了。
      // data:alternateData,
      levels: 2, // expanded to 5 levels
      // enableLinks:true,
      showTags: true,
    });
    $('#treeview').on('nodeSelected', function(event, data) {
      // alert("名称："+data.text);
      // alert("节点id："+data.nodeId);
      // alert("部门id："+data.Id);  
      // alert("部门级别："+data.Level);
      $("#regis").html(data.text); //显示部门名称
      $("#regis").css("color", "black");
      //点击菜单，右侧显示那个人的主页
      // document.getElementById("iframepage").src="/index/main/"+data.Id;
      //点击菜单，打开新标签页个人主页
      window.open('http://' + data.href)
      //当前窗口打开这个个人主页：
      // parent.location.href='/index/main/'+data.Id
    });
  });
  </script>
  <script>
  var myChart = echarts.init(document.getElementById('main'));
  var szRoad = {
    success: true,
    errorCode: 0,
    errorMsg: "成功",
    data: [{
        ROAD_LINE: "113.0665043,22.79910761;113.0683087,22.80351522;113.070613,22.80482283;113.0772174,22.80853043;113.0795217,22.80933804;113.0906261,22.81444565;113.0926304,22.81475326;113.1023348,22.81876087;113.1040391,22.81896848;113.1084435,22.82497609;113.1102478,22.8252837;113.1128522,22.8257913;113.1149565,22.82669891;113.1164609,22.82700652;113.1180652,22.82751413;113.1196696,22.82782174;113.1216739,22.82852935;113.1292783,22.82953696;113.1308826,22.82964457;113.146287,22.83795217;113.1477913,22.83805978;113.1546957,22.83876739;113.1563,22.838875;113.1587043,22.83928261;113.1605087,22.83939022;113.163513,22.83979783;113.1653174,22.83990543;113.1760217,22.83991304;113.1778261,22.84002065;113.1793304,22.84012826;113.1813348,22.84013587;113.1828391,22.84024348;113.1846435,22.84045109;113.2014478,22.8404587;113.2031522,22.8405663;113.2066565,22.84087391;113.2086609,22.84098152;113.2117652,22.84038913;113.2148696,22.84059674;113.2235739,22.84270435;113.2252783,22.84291196;113.2270826,22.84321957;113.228687,22.84342717;113.2303913,22.84353478;113.2323957,22.84384239;113.2353,22.84815;113.2369043,22.84825761;113.2434087,22.84846522;113.245213,22.84867283;113.2493174,22.84868043;113.2517217,22.84408804;113.2533261,22.84399565;113.2551304,22.84360326;113.2570348,22.84341087;113.2632391,22.84321848;113.2655435,22.84332609;113.2677478,22.8432337;113.2709522,22.8426413;113.2725565,22.84264891;113.2761609,22.84175652;113.2776652,22.84186413;113.2870696,22.83987174;113.2898739,22.83987935;113.3093783,22.84218696;113.3120826,22.84219457;113.321287,22.84050217;113.3232913,22.84040978;113.3273957,22.83551739;113.3348,22.835225;113.3378043,22.83523261;113.3400087,22.83534022;113.343013,22.83584783;113.3453174,22.83615543;113.3539217,22.84136304;113.3556261,22.84147065;113.3598304,22.84247826;113.3665348,22.84288587;113.3698391,22.84329348;113.3724435,22.84380109;113.3745478,22.8446087;113.3775522,22.8452163;113.3859565,22.84492391;113.3879609,22.84503152;113.3907652,22.84513913;113.3927696,22.84524674;113.4000739,22.84505435;113.4027783,22.84436196;113.4047826,22.84376957;113.408287,22.83717717;113.4098913,22.83708478;113.4114957,22.83699239;113.413,22.837"
      },
      //高新沙至沙溪
      {
        ROAD_LINE: "113.4217189,22.83728017;113.4250378,22.83826034;113.4273566,22.83884051;113.4310755,22.84002067;113.4331944,22.84050084;113.4354133,22.84088101;113.4373321,22.84126118;113.463251,22.85004135;113.4655699,22.85052152;113.4710888,22.85170169;113.4733076,22.85208185;113.4757265,22.85266202;113.4819454,22.85314219;113.4860643,22.85492236;113.4891831,22.86070253;113.491102,22.8614827;113.4946209,22.86316287;113.5015398,22.86354303;113.5044587,22.8641232;113.5077775,22.86490337;113.5100964,22.86528354;113.5223153,22.87156371;113.5301342,22.87204388;113.533353,22.87222404;113.5553719,22.86750421;113.5662908,22.86688438;113.5682097,22.86716455;113.5709285,22.86734472;113.5734474,22.86732489;113.5758663,22.86680506;113.5783852,22.86658522;113.585204,22.86656539;113.5874229,22.86684556;113.5900418,22.86692573;113.5922607,22.8669059;113.5947796,22.86708607;113.6081984,22.86706624;113.6100173,22.8673464;113.6171362,22.86742657;113.6194551,22.86760674;113.6216739,22.86758691;113.6252928,22.86706708;113.6282117,22.86324725;113.6301306,22.86352742;113.6371494,22.86820758;113.6393683,22.86878775;113.6412872,22.86916792;113.6436061,22.86984809;113.6456249,22.87032826;113.6474438,22.87060843;113.6493627,22.8709886;113.6511816,22.87126876;113.6533004,22.87174893;113.6553193,22.8722291;113.6575382,22.87280927;113.6664571,22.87398944;113.668276,22.87436961;113.6720948,22.87504978;113.6741137,22.87542994;113.6761326,22.87571011"
      },
      //沙溪至罗田
      {
        ROAD_LINE: "113.6761326,22.87571011;113.6771303,22.87654045;113.6989607,22.8842809;113.7015258,22.88480112;113.7047236,22.88453146;113.7062562,22.88574157;113.7067888,22.89095169;113.7113865,22.89208202;113.7117191,22.89209213;113.7218169,22.89782247;113.7270494,22.89773258;113.727682,22.8973427;113.7471146,22.89205281;113.7554472,22.89086292;113.7588427,22.8871236;113.7689056,22.88477416;113.7726708,22.88399438;113.7771663,22.88135506;113.7780247,22.88086629;113.7784899,22.88078652;113.7874876,22.88031685;113.7934202,22.88022697;113.8021854,22.87824719;113.8084483,22.87449775;113.8193809,22.86690787;113.8344461,22.86622809"
      },
      //东莞
      {
        ROAD_LINE: "113.8744059,22.85827647;113.8743118,22.86025294;113.8739176,22.86232941;113.8737235,22.86370588;113.8731294,22.86468235;113.8733353,22.86705882;113.8728412,22.86373529;113.8724471,22.86551176;113.8723529,22.86688824;113.8767588,22.86866471;113.8766647,22.87044118;113.8763706,22.87171765;113.8762765,22.87349412;113.8766824,22.87517059;113.8761882,22.87744706;113.8755941,22.87972353;113.875,22.882"
      },
      //深圳
      {
        ROAD_LINE: "113.8848607,22.83366786;113.8893214,22.83043571;113.8938821,22.82710357;113.8984429,22.82387143;113.9028036,22.82063929;113.8984643,22.82280714;113.902525,22.820475;113.9072857,22.81774286;113.9118464,22.81641071;113.9161071,22.81407857;113.9201679,22.81534643;113.9248286,22.81231429;113.9285893,22.81018214;113.93305,22.80715;113.9316107,22.80511786;113.9364714,22.80178571;113.9412321,22.79855357;113.9452929,22.79622143;113.9498536,22.79408929;113.9545143,22.79555714;113.958575,22.793225;113.9530357,22.79309286;113.9575964,22.78996071;113.9525571,22.78742857;113.9563179,22.78869643;113.9553786,22.78736429;113.9601393,22.78433214;113.965,22.781"
      },
      //南沙支线
      {
        ROAD_LINE: "113.4217625,22.83740625;113.424325,22.8382125;113.4264875,22.83881875;113.42885,22.839625;113.4309125,22.84023125;113.434675,22.8416375;113.4369375,22.84224375;113.4391,22.84275;113.4410625,22.84315625;113.462725,22.8519625;113.4691875,22.85256875;113.47185,22.853175;113.4751125,22.85338125;113.477875,22.8530875;113.4804375,22.85269375;113.483,22.8476"
      }
    ]
  }

  var busLines = [];
  var data = szRoad.data;
  var hStep = 300 / (data.length - 1);

  var i = 0;
  for (var x in data) {
    // i++;
    // if(i<5000)
    //     continue;
    var line = data[x];
    // if(busLines.length>500)
    //     break;
    var pointString = line.ROAD_LINE;
    var pointArr = pointString.split(';');
    var lnglats = [];
    for (var j in pointArr) {
      lnglats.push(pointArr[j].split(','))
    }
    busLines.push({
      coords: lnglats,
      lineStyle: {
        normal: {
          color: echarts.color.modifyHSL('#5A94DF', Math.round(hStep * x))
        }
      }
    })
  }

  option = {
    bmap: {
      center: [113.525, 22.82], //[104.091, 30.639],
      zoom: 11,
      roam: true,
      mapStyle: {
        'styleJson': [{
          'featureType': 'water',
          'elementType': 'all',
          'stylers': {
            'color': '#4169E1'
          }
        }, {
          'featureType': 'land',
          'elementType': 'geometry',
          'stylers': {
            'color': '#3CB371'
          }
        }, {
          'featureType': 'highway',
          'elementType': 'all',
          'stylers': {
            'visibility': 'off'
          }
        }, {
          'featureType': 'arterial',
          'elementType': 'geometry.fill',
          'stylers': {
            'color': '#FFFFFF'
          }
        }, {
          'featureType': 'arterial',
          'elementType': 'geometry.stroke',
          'stylers': {
            'color': '#0b3d51'
          }
        }, {
          'featureType': 'local',
          'elementType': 'geometry',
          'stylers': {
            'color': '#FFFFFF'
          }
        }, {
          'featureType': 'railway',
          'elementType': 'geometry.fill',
          'stylers': {
            'color': '#FFFFFF'
          }
        }, {
          'featureType': 'railway',
          'elementType': 'geometry.stroke',
          'stylers': {
            'color': '#08304b'
          }
        }, {
          'featureType': 'subway',
          'elementType': 'geometry',
          'stylers': {
            'lightness': -70
          }
        }, {
          'featureType': 'building',
          'elementType': 'geometry.fill',
          'stylers': {
            'color': '#FFFFFF'
          }
        }, {
          'featureType': 'all',
          'elementType': 'labels.text.fill',
          'stylers': {
            'color': '#857f7f'
          }
        }, {
          'featureType': 'all',
          'elementType': 'labels.text.stroke',
          'stylers': {
            'color': '#FFFFFF'
          }
        }, {
          'featureType': 'building',
          'elementType': 'geometry',
          'stylers': {
            'color': '#022338'
          }
        }, {
          'featureType': 'green',
          'elementType': 'geometry',
          'stylers': {
            'color': '#32CD32'
          }
        }, {
          'featureType': 'boundary',
          'elementType': 'all',
          'stylers': {
            'color': '#465b6c'
          }
        }, {
          'featureType': 'manmade',
          'elementType': 'all',
          'stylers': {
            'color': '#022338'
          }
        }, {
          'featureType': 'label',
          'elementType': 'all',
          'stylers': {
            'visibility': 'off'
          }
        }]
      }
    },
    series: [{
        type: 'lines',
        coordinateSystem: 'bmap',
        polyline: true,
        data: busLines,
        silent: true,
        lineStyle: {
          normal: {
            // color: '#c23531',
            // color: 'rgb(200, 35, 45)',
            opacity: 0.6,
            width: 5
          }
        },
        progressiveThreshold: 500,
        progressive: 200
      }, {
        type: 'lines',
        coordinateSystem: 'bmap',
        polyline: true,
        data: busLines,
        lineStyle: {
          normal: {
            width: 0
          }
        },
        effect: {
          constantSpeed: 50,
          show: true,
          trailLength: 0.5,
          symbolSize: 8 //1.5
        },
        zlevel: 1
      },
      {
        "name": "1万",
        "type": "effectScatter",
        "coordinateSystem": "bmap",
        "data": [{
          "name": "鲤鱼洲",
          "value": [113.064, 22.8]
        }],
        // "rippleEffect": {
        //     "period": 4,
        //     "scale": 4,
        //     "brushType": "fill"
        // },
        rippleEffect: {
          // brushType: 'stroke'
          "period": 4,
          "scale": 6,
          "brushType": "fill"
        },
        "label": {
          "normal": {
            "formatter": "{b}",
            "position": "right",
            "show": true
          },
          "emphasis": {
            "show": true
          }
        },
        "itemStyle": {
          "normal": {
            "color": "fuchsia"
          }
        }
      },
      {
        "name": "1万",
        "type": "effectScatter",
        "coordinateSystem": "bmap",
        "data": [{
          "name": "高新沙",
          "value": [113.413, 22.837]
        }],
        // "rippleEffect": {
        //     "period": 4,
        //     "scale": 4,
        //     "brushType": "fill"
        // },
        rippleEffect: {
          // brushType: 'stroke'
          "period": 4,
          "scale": 6,
          "brushType": "fill"
        },
        "label": {
          "normal": {
            "formatter": "{b}",
            "position": "right",
            "show": true
          },
          "emphasis": {
            "show": true
          }
        },
        "itemStyle": {
          "normal": {
            "color": "navy"
          }
        }
      },
      {
        "name": "1万",
        "type": "effectScatter",
        "coordinateSystem": "bmap",

        "data": [{
          "name": "沙溪",
          "value": [113.677, 22.876]
        }],
        // "rippleEffect": {
        //     "period": 4,
        //     "scale": 4,
        //     "brushType": "fill"
        // },
        rippleEffect: {
          // brushType: 'stroke'
          "period": 4,
          "scale": 6,
          "brushType": "fill"
        },
        "label": {
          "normal": {
            "formatter": "{b}",
            "position": "right",
            "show": true
          },
          "emphasis": {
            "show": true
          }
        },
        "itemStyle": {
          "normal": {
            "color": "maroon"
          }
        }
      },
      {
        "name": "1万",
        "type": "effectScatter",
        "coordinateSystem": "bmap",
        "data": [{
          "name": "罗田",
          "value": [113.878, 22.841]
        }],
        // "rippleEffect": {
        //     "period": 4,
        //     "scale": 4,
        //     "brushType": "fill"
        // },
        rippleEffect: {
          // brushType: 'stroke'
          "period": 4,
          "scale": 6,
          "brushType": "fill"
        },
        "label": {
          "normal": {
            "formatter": "{b}",
            "position": "right",
            "show": true
          },
          "emphasis": {
            "show": true
          }
        },
        "itemStyle": {
          "normal": {
            "color": "lime"
          }
        }
      },
      {
        "name": "1万",
        "type": "effectScatter",
        "coordinateSystem": "bmap",
        "data": [{
          "name": "松木山",
          "value": [113.878, 22.887]
        }],
        // "rippleEffect": {
        //     "period": 4,
        //     "scale": 4,
        //     "brushType": "fill"
        // },
        rippleEffect: {
          // brushType: 'stroke'
          "period": 4,
          "scale": 6,
          "brushType": "fill"
        },
        "label": {
          "normal": {
            "formatter": "{b}",
            "position": "right",
            "show": true
          },
          "emphasis": {
            "show": true
          }
        },
        "itemStyle": {
          "normal": {
            "color": "#ff3333"
          }
        }
      },
      {
        "name": "1万",
        "type": "effectScatter",
        "coordinateSystem": "bmap",
        "data": [{
          "name": "公明",
          "value": [113.965, 22.781]
        }],
        // "rippleEffect": {
        //     "period": 4,
        //     "scale": 4,
        //     "brushType": "fill"
        // },
        rippleEffect: {
          // brushType: 'stroke'
          "period": 4,
          "scale": 6,
          "brushType": "fill"
        },
        "label": {
          "normal": {
            "formatter": "{b}",
            "position": "right",
            "show": true
          },
          "emphasis": {
            "show": true
          }
        },
        "itemStyle": {
          "normal": {
            "color": "orange"
          }
        }
      },
      {
        "name": "1万",
        "type": "effectScatter",
        "coordinateSystem": "bmap",
        "data": [{
          "name": "黄阁",
          "value": [113.483, 22.8476]
        }],
        // "rippleEffect": {
        //     "period": 4,
        //     "scale": 4,
        //     "brushType": "fill"
        // },
        rippleEffect: {
          // brushType: 'stroke'
          "period": 4,
          "scale": 6,
          "brushType": "fill"
        },
        "label": {
          "normal": {
            "formatter": "{b}",
            "position": "right",
            "show": true
          },
          "emphasis": {
            "show": true
          }
        },
        "itemStyle": {
          "normal": {
            "color": "orange"
          }
        }
      }
    ]
  };
  myChart.setOption(option);
  </script>
</body>

</html>
<!-- <button id="directNextpage" onclick="window.location.reload('/topic/add?id={{.Category.Id}}&mid=1')">Direct Next Page</button> 这个方法跳不出去iframe
  onclick="window.open('/topic/add?id={{.Category.Id}}&mid=1')"新标签页中打开  onclick="parent.location.href='/topic/add?id={{.Category.Id}}&mid=2'跳出iframe重新打开-->