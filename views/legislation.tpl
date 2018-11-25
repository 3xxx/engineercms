<!DOCTYPE html>
{{template "header"}}
<title>首页 - 水利设计CMS系统</title>
   <!-- <meta charset="UTF-8"> -->
   <link type="text/css" href="/static/youdao/g3.css" rel="stylesheet">
   <link type="text/css" href="/static/youdao/fanyi.css" rel="stylesheet">
   <script type="text/javascript" src="/static/js/jquery-2.1.3.min.js"></script>
   <script type="text/javascript" src="/static/youdao/openapi.do" charset="utf-8"></script>
   <script type="text/javascript" src="/static/youdao/dict_req_web_1.1.js"></script>
   <script type="text/javascript" src="/static/youdao/fanyi.js" data-main="/fanyi" ></script>
   <script type="text/javascript" src="/static/youdao/ntes.js"></script>
</head>
<body class="open result-default">
  <div class="navbar navba-default navbar-fixed-top">
    <div class="container-fill">{{template "navbar" .}}</div>
  </div>
  <h1 > <i class="glyphicon glyphicon-chevron-right"></i> <i class="glyphicon glyphicon-minus"></i>
  </h1>
  <div id="w" class="cf">
    <div id="transBackground">
        <div id="main" class="cf show-translate">
            <div id="transBtnTip">
                <div id="transBtnTipInner">
                    点击对标按钮继续，查看网页对标结果。
                    <p class="ar">
                        <a href="/#" id="transBtnTipOK">我知道了</a>
                    </p> <b id="transBtnTipArrow"></b>
                </div>
            </div>
            <div id="inputMod" class="column fl">
                <div class="wrapper">
                    <form action="" method="post" id="transForm" name="transForm">
                      <div class="clearall">
                        <div id="clearInput" class="clog-js" style="display: block;">
                          <div class="clearnotice">清空内容</div>
                        </div>
                      </div>
                      <div class="row border content" id="inputContent">
                        <textarea id="inputText" class="text" dir="ltr" tabindex="1" wrap="SOFT" name="name" placeholder="将报告中的法规复制粘贴到这里，一行一条，《法规名称必须放在书名号中》；自动记录未查询到的法规，以便改进。"></textarea>
                                <!-- <div class="typo-suggest" style="display: none;">
                                    您是不是要对标：
                                    <a class="spell-corrected" href="/#"></a>
                                </div> -->
                      </div>
                            <!-- <div class="row">
                                <a type="submit" id="translateBtn" href="/#" title="Enter自动对标" value="自动对标" name="action" class="button translateBtn"></a>
                            </div> -->

                      <span class="input-group-btn">
                        <button class="button translateBtn" type="button" id="checklist">
                        <i class="glyphicon glyphicon-search"></i>
                                    Checklist!
                        </button>
                      </span>
                    </form>
          </div>
                    <!-- end of wrapper --> </div>
                <!-- end of div inputMod -->
                <div id="outputMod" class="column fr">
                    <div class="wrapper">
                        <!-- end of entryList -->
                        <div class="row-hidden" id="outputHidden"></div>
                        <div id="translated" style="display: block;">

                            <div class="row">
                                <div class="row" id="outputText">
                                    <div class="translated_result">
                                        <!-- <p class="tgt">Austin</p> -->
                                    </div>
                                </div>
                            </div>
                            <div id="modeWrapper" class="read-mode" style="display: block;">

                                <a class="open-reading-mode title" href="h/#" title="" hidefocus=""></a>
                                <a class="close-reading-mode title" href="/#" title="" hidefocus=""></a>
                                <div class="opennotice">
                                    <div class="arrow"> <em></em>
                                        <span></span>
                                    </div>
                                    全屏阅读
                                </div>
                                <div class="closenotice">
                                    <div class="arrow">
                                        <em></em>
                                        <span></span>
                                    </div>
                                    关闭全屏阅读
                                </div>
                            </div>

                            <div class="row desc">

                                <div id="selectorSwitcher" class="selector-sprite selector-enable" style="display: block;">
                                    <span id="selectorStatus">显示执行日期</span>
                                </div>
                                
                                <div class="read-mode" id="compareMode" style="display: block;">
                                    <label class="compare-mode compare-disable" for="compare">
                                        <input id="compare" name="compare"  type="checkbox"><!-- class="clog-js" data-clog="COMPARE_CLICK" data-pos="web.o.righttop" -->
                                        <span class="compare-message">法规显示大会通过</span>
                                    </label>

                                </div>

                                <div class="tool">
                                    <a href="/#" id="speech" title=""></a>
                                    <a href="/#" id="resultScore" style="background-position: 0px -27px;"></a>
                                    <div class="speechnotice" style="display: none;">
                                        <div class="arrow">
                                            <em></em>
                                            <span></span>
                                        </div>
                                        朗读
                                    </div>
                                    <div class="copynotice" style="display: none;">
                                        <div class="arrow">
                                            <em></em>
                                            <span></span>
                                        </div>

                                        <span class="copy-notice">复制</span>
                                    </div>
                                    <a class="actions copyIt-js">
                                        <span id="copyOutput" class="copy-init"></span>
                                    </a>
                                </div>
                            </div>

                        </div>
                        <!-- end translated --> </div>
                    <!-- end of wrapper -->
                    <div class="row cf" id="addons" style="display: none;">

                </div>
                <div class="suggest" style="display: none;">
                    <div id="suggestYou"></div>
                </div>
            </div>

        </div>
        <!-- end of main --> </div>
</div>

<div id="c_footer">
    <a href="/">首页</a>
    <span class="c_fnl">|</span>
    <a href="https://github.com/3xxx">源码托管</a>
    <p class="c_fcopyright">© 2016~2018 <a href="https://github.com/3xxx" target="_blank" rel="nofollow">3xxx</a> QQ504284</p>
</div>

<script type="text/javascript">
    var global = {};
    var abtest = "0";
    global.abTest = "0";
    global.sessionFrom = "https://www.baidu.com/link";
</script>


<div id="custheme"></div>
<!-- START NetEase Devilfish 2006 -->

<script>
// <button class="btn btn-primary" id="export">导出excel</button>
$(document).ready(function(){
$("#checklist").click(function(){//这里应该用button的id来区分按钮的哪一个,因为本页有好几个button
    $("#translated")[0].style.display = 'block';
    $('#translated').css('display','block');
    $("#translated").css({ display: "block"});
        $.ajax({
        type:"post",//这里是否一定要用post？？？
        url:"/legislation/checklist",
        data: {name: $("#inputText").val()},
        success:function(data,status){//数据提交成功时返回数据
            // alert(data[0].Title);
            $.each(data,function(i,d){
                // alert(data[i].Title);
                $(".translated_result").append('<p>('+data[i].Id+')&nbsp;《'+data[i].LibraryTitle+'》&nbsp;('+data[i].LibraryNumber+')</p>');
                    }); 

        }       
        });
    });
});
</script> 

</body>
</html>