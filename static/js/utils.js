/**
 * Created by hzwy23 on 2016/9/28.
 */

// tab 管理模块
var Hutils = {
    rowStyle:function(row,index) {
        return {css:{
            "white-space":"nowrap",
            "text-overflow":"ellipsis"
        }}
    },
    // 隐藏子菜单系统，切换具体页面内容
    hideWrapper:function(){
        $("#wrapper").removeClass("animated slideInUp slideOutDown");
        $("#wrapper").addClass("animated slideOutDown");
        $("#h-main-content").removeClass("animated slideInDown slideOutUp");
        $("#h-main-content").addClass("animated slideInDown");
    },
    // 隐藏内容显示部分，切换到子菜单系统
    showWrapper:function() {
        $("#h-main-content").removeClass("animated slideInDown slideOutUp");
        $("#h-main-content").addClass("animated slideOutUp");
        $("#wrapper").removeClass("animated slideInUp slideOutDown");
        $("#wrapper").addClass("animated slideInUp")
    },
    // 判断子菜单系统显示状态，如果是隐藏，则切换到显示，如果是显示，则隐藏。
    HchangeWrapper:function(){
        if (window.event != undefined){
            window.event.cancelBubble = true;
        } else {
            // firefox
            var event = Hutils.getEvent()
            event.stopPropagation()
        }
        if ($(".H-tabs-index").html()==""){
            $.Notify({
                title:"温馨提示：",
                message:"目前没有已经打开的页面",
                type:"info",
            });
            return
        };

        // 判断子系统菜单也距离底部的位置，如果距离底部的位置是0，则隐藏子菜单系统，否则显示子菜单系统
        if ($("#wrapper").hasClass("slideInUp")){
            Hutils.hideWrapper()
        }else{
            Hutils.showWrapper()
        }
    },
    ShowCannotEditTips:function(obj){
        $(obj).tooltip({
            title:"亲,此处无法编辑哟"
        }).tooltip("show")
    },
    // 跳转到首页系统菜单。
    H_HomePage:function(){
        if (window.event != undefined){
            window.event.cancelBubble = true;
        } else {
            // firefox
            var event = Hutils.getEvent()
            event.stopPropagation()
        }
        window.location.href="/HomePage"
    },
    // 退出登录
    HLogOut:function(){
        if (window.event != undefined){
            window.event.cancelBubble = true;
        } else {
            // firefox
            var event = Hutils.getEvent()
            event.stopPropagation()
        }
        $.Hconfirm({
            callback:function(){
                $.ajax({type:"Get",url:"/signout",cache:!1,async:!1,dataType:"text",
                    error:function(){window.location.href="/"},
                    success:function(a){
                        $.Notify({
                            message:"安全退出",
                            type:"success",
                        });
                        window.location.href="/"}
                })
            },
            body:"点击确定退出系统"
        })
    },
    // 用户信息管理
    UserMgrInfo:function(){
        if (window.event != undefined){
            window.event.cancelBubble = true;
        } else {
            // firefox
            var event = Hutils.getEvent()
            event.stopPropagation()
        }

        $.Hmodal({
            body:$("#mas-passwd-prop").html(),
            footerBtnStatus:false,
            height:"420px",
            width:"720px",
            header:"用户信息",
            preprocess:function () {
                $.getJSON("/v1/auth/user/query",function (data) {
                    $(data).each(function (index, element) {
                        $("#h-user-details-user-id").html(element.user_id)
                        $("#h-user-details-user-name").html(element.user_name)
                        $("#h-user-details-user-email").html(element.user_email)
                        $("#h-user-details-user-phone").html(element.user_phone)

                        $("#h-user-details-user-org-name").html(element.org_unit_desc)
                        $("#h-user-details-user-domain").html(element.domain_id)
                        $("#h-user-details-user-domain-name").html(element.domain_name)
                        $("#h-user-details-user-create").html(element.create_user)
                        $("#h-user-details-user-create-date").html(element.create_date)
                        $("#h-user-details-user-modify").html(element.modify_user)
                        $("#h-user-details-user-modify-date").html(element.modify_date)
                        // 机构编码处理
                        var upcombine = element.org_unit_id.split("_join_")
                        if (upcombine.length==2){
                            $("#h-user-details-user-org").html(upcombine[1])
                        }else{
                            $("#h-user-details-user-org").html(upcombine)
                        }
                    })
                });
            }
        });
    },
    // 子系统中，打开具体页面按钮
    goEntrySubSystem:function(e){
        var flag = false;

        // 资源的url地址
        var url = $(e).attr("data-url");

        // 资源的id
        var data_id = $(e).attr("data-id");

        // 资源的名称
        var name = $(e).find("div:last").html();

        if ($(e).attr("data-openType") == "1") {
            window.open(url,name)
            return
        }
        NProgress.start();

        // 遍历整个tab栏目，查找指定id的资源是否打开，
        // 如果该资源已经打开，则直接切换到该资源，无需从后台获取内容
        // 如果该资源没有打开，则将flag为false，从后台获取资源内容
        $(".H-tabs-index").find("span").each(function(index,element){
            // 如果资源存在，直接切换到这个资源的tab中。
            if (data_id == $(element).attr("data-id")){
                Hutils.__changetab(element)
                flag = true;
                return false;
            }
        });

        // 资源未打开，从后台请求资源信息
        if (flag == false){
            $.HAjaxRequest({
                type:"get",
                url:url,
                cache:false,
                async:true,
                dataType:"text",
                error:function (msg) {
                    console.log(msg);
                    NProgress.done();
                    var m = JSON.parse(msg);
                    $.Notify({
                        title:"温馨提示:",
                        message:m.error_msg,
                        type:"danger",
                    });
                    return
                },
                success: function(data){

                    // 隐藏内容显示区域
                    $("#h-main-content").find("div.active")
                        .removeClass("active").addClass("none");

                    var newContent = document.createElement("div")
                    $(newContent).attr({
                        "data-type":"frame",
                        "data-id":data_id,
                    }).css({
                        "padding":"0px",
                        "margin":"0px",
                    }).addClass("active").html(data);

                    $("#h-main-content").append(newContent);

                    // 隐藏子菜单系统，显示具体的内容
                    Hutils.hideWrapper();

                    // 生成标签页
                    {
                        // 清楚所有的tab选中状态
                        $(".active-tab").removeClass("active-tab");

                        // 获取新tab模板内容
                        var optHtml = Hutils.__genTabUI(data_id,name)

                        // 在tab栏目列表中添加新的tab
                        $(".H-tabs-index").append(optHtml);
                    }

                    NProgress.done();
                }
            });
        }else {
            NProgress.done();
        }
    },
    // 打开指定资源按钮
    openTab:function(param){

        NProgress.start();

        // 判断子元素会否已经被打开，默认设置为未打开
        var flag = false;

        var __DEFAULT = {
            id:"",
            tile:"",
            type: "GET",
            url: "",
            data: {},
            dataType: "text",
        };
        $.extend(true,__DEFAULT,param);

        // 资源url地址
        var url = __DEFAULT.url;

        // 资源id
        var data_id = __DEFAULT.id;

        // 资源名称
        var name = __DEFAULT.title;

        $(".H-tabs-index").find("span").each(function(index,element){
            if (data_id == $(element).attr("data-id")){
                flag = true;
                $.HAjaxRequest({
                    type:__DEFAULT.type,
                    url:url,
                    cache:false,
                    async:true,
                    data:__DEFAULT.data,
                    dataType:__DEFAULT.dataType,
                    success: function(data){
                        Hutils.__changetab(element);
                        $(element).find("hzw").html(name);

                        $("#h-main-content").find("div.active").each(function(index,element){
                            if (data_id == $(element).attr("data-id")){
                                $(element).html(data);
                                return false;
                            }
                        })
                    }
                });
                return false;
            }
        });

        if (flag == false){
            $.HAjaxRequest({
                type:__DEFAULT.type,
                url:url,
                cache:false,
                async:true,
                data:__DEFAULT.data,
                dataType:__DEFAULT.dataType,
                error:function (msg) {
                    var m = JSON.parse(msg.responseText);
                    console.log(m);
                    $.Notify({
                        title:"温馨提示:",
                        message:m.error_msg,
                        type:"danger",
                    });
                    NProgress.done();
                },
                success: function(data){

                    // 隐藏内容显示区域
                    $("#h-main-content").find("div.active")
                        .removeClass("active").addClass("none");

                    var newContent = document.createElement("div")
                    $(newContent).attr({
                        "data-type":"frame",
                        "data-id":data_id,
                    }).css({
                        "padding":"0px",
                        "margin":"0px",
                    }).addClass("active").html(data);

                    $("#h-main-content").append(newContent).hide().fadeIn();

                    // 生成标签页
                    {
                        // 清楚所有的tab选中状态
                        $(".active-tab").removeClass("active-tab");

                        // 获取新tab模板内容
                        var optHtml = Hutils.__genTabUI(data_id,name)

                        // 在tab栏目列表中添加新的tab
                        $(".H-tabs-index").append(optHtml);
                    }
                    NProgress.done();
                }
            });
        } else {
            NProgress.done();
            return
        }
    },

    // 切换tab页面
    __changetab : function(e){

        // 获取新tab的id
        var id = $(e).attr("data-id");
        var flag = true;

        if (window.event != undefined){
            window.event.cancelBubble = true;
        } else {
            // firefox
            var event = Hutils.getEvent()
            event.stopPropagation()
        }

        $(".active-tab").each(function (index, element) {
            var old_id = $(element).attr("data-id")
            // 返现客户点点击的是当前tab也
            // 退出切换tab也操作,仍然显示当前页面
            if (id == old_id) {
                flag = false;
                return false
            }
        });

        // 隐藏子菜单页面
        Hutils.hideWrapper()

        // 如果切换到当前tab也,则直接退出处理.
        if (!flag) {
            return
        }

        // 清除所有tab的激活标签
        $(".active-tab").removeClass("active-tab");

        // 给新的tab加上激活标签
        $(e).addClass("active-tab")


        // 在已经打开的页面中，根据id，寻找到指定的页面，将这个页面显示出来
        $("#h-main-content").find("div.active")
            .removeClass("active")
            .addClass("none");

        $("#h-main-content").find("div[data-id='"+id+"']")
            .removeClass("none")
            .addClass("active")
            .hide()
            .fadeIn(300);
    },
    __genTabUI:function (data_id,name) {
        var mspan = document.createElement("span")
        $(mspan).css({
            "min-width":"120px",
            "border-left":"#6f5499 solid 1px"
        }).attr({
            "data-id":data_id,
            "onclick":"Hutils.__changetab(this)",
        }).addClass("H-left-tab active-tab");

        var hzw = document.createElement("hzw");
        $(hzw).html(name);
        $(hzw).css({
            "font-weight":"600",
            "color":"white",
        });

        var mi = document.createElement("i")
        $(mi).css("font-size","14px")
            .addClass("icon-remove-sign H-gray-close pull-right")
            .attr("onclick","Hutils.__closetab(this)")

        $(mspan).append(hzw);
        $(mspan).append(mi);
        return mspan
    },
    getEvent:function(){
        if(window.event)    {
            return window.event;
        }
        var func = Hutils.getEvent.caller;
        while( func != null ){
            var arg0 = func.arguments[0];
            if(arg0){
                if((arg0.constructor==Event || arg0.constructor ==MouseEvent
                    || arg0.constructor==KeyboardEvent)
                    ||(typeof(arg0)=="object" && arg0.preventDefault
                    && arg0.stopPropagation)){
                    return arg0;
                }
            }
            func = func.caller;
        }
        return null;
    },
    // 关闭tab标签，以及tab标签关联的内容,在__genTabUI中引用了__closetab
    __closetab:function(e){
        // 取消后续事件
        if (window.event != undefined){
            window.event.cancelBubble = true;
        } else {
            var event = Hutils.getEvent()
            event.stopPropagation()
        }

        // 获取被关闭tab的id
        var id = $(e).parent().attr("data-id");

        // 首先判断，这个tab是否被激活，如果是激活状态，则在关闭tab后，
        // 还需要切换到新的tab页面中，切换顺序是，先寻找左侧隐藏的tab，如果没有再寻找右侧
        // 如果两侧都没有，则直接返回子菜单系统。
        if ($(e).parent().hasClass("active-tab")){
            // 获取左侧tab
            var pobj = $(e).parent().prev("span");
            var pid = $(pobj).attr("data-id");

            // 获取右侧tab
            var nobj = $(e).parent().next("span");
            var nid = $(nobj).attr("data-id");

            // 关闭选中的tab,以及这个tab所关联的内容
            $(e).parent().remove();
            $("#h-main-content").find("div[data-id='"+id+"']").remove();

            // 如果pid与nid都为undefined，则直接切换到子菜单系统
            // 如果左侧tab存在，则切换到左侧tab，否则切换到右侧tab
            if (pid == undefined){
                if (nid == undefined){
                    Hutils.showWrapper()
                    return
                } else {
                    id = nid
                }
            } else {
                id = pid
            }

            // 清除左侧tab的隐藏状态，使其显示。
            $("#h-main-content").find("div[data-id='"+id+"']")
                .removeClass("none")
                .addClass("active")
                .hide()
                .fadeIn(500);

            // 遍历整个tab栏，找到匹敌的tab id，
            $(".H-left-tab").each(function(index,element){
                if (id == $(element).attr("data-id")){
                    $(element).addClass("active-tab")
                }
            });

        } else {
            // 当被删除的这个tab没有被激活时，直接将这个tab也从tab栏目中删除，并连同删除这个tab关联的内容即可。
            $(e).parent().remove();
            $("#h-main-content").find("div[data-id='"+id+"']").remove();
        }
    },
    go_entry :function (e){
        var id = $(e).attr("data-id");
        $.HAjaxRequest({
            url:'/v1/auth/index/entry',
            data:{Id:id},
            dataType:'text',
            success:function(d){
                $("#bigdata-platform-subsystem").html(d)
            },
            error:function () {
                $.Notify({
                    title:"温馨提示：",
                    message:"登录连接已经断开，请重新登录系统",
                    type:"info",
                });
                window.location.href="/"
            },
        });
    },
    initMenu:function(TypeId,Id,Group1,Group2,Group3){

        var __genUI = function (name) {
            var mdiv = document.createElement("div")
            $(mdiv).addClass("tile-group")
            var mspan = document.createElement("span")
            $(mspan).addClass("tile-group-title").css("font-size","12px").html(name)
            $(mdiv).append(mspan)

            return mdiv
        };

        var __genDiv = function (res_id,res_class,res_bg_color,res_img,res_name,res_url,open_type) {
            var mdiv = document.createElement("div")
            $(mdiv).attr({
                "data-id":res_id,
                "data-role":"tile",
                "data-url":res_url,
                "data-openType":open_type,
            }).addClass(res_class).addClass("fg-white hzwy23div")
                .css("background-color",res_bg_color);

            var cdiv = document.createElement("div")
            $(cdiv).addClass("tile-content iconic")

            var mspan = document.createElement("span")
            $(mspan).addClass("icon")

            var mimg = document.createElement("img")
            mimg.style.height = "100%";
            mimg.style.width = "100%";
            $(mimg).attr("src",res_img)

            var ccdiv = document.createElement("div");
            $(ccdiv).addClass("tile-label").html(res_name);

            $(mspan).append(mimg)
            $(cdiv).append(mspan)
            $(mdiv).append(cdiv)
            $(mdiv).append(ccdiv)
            return mdiv
        };

        $.HAjaxRequest({
            url:'/v1/auth/main/menu',
            data:{TypeId:TypeId,Id:Id},
            success: function(data){

                var cdiv1 = document.createElement("div");
                $(cdiv1).addClass("tile-container");

                var cdiv2 = document.createElement("div");
                $(cdiv2).addClass("tile-container");

                var cdiv3 = document.createElement("div");
                $(cdiv3).addClass("tile-container");

                var divlist = new Array();

                divlist.push(cdiv1);
                divlist.push(cdiv2);
                divlist.push(cdiv3);

                $(data).each(function(index,element){
                    var gid = parseInt(element.Group_id)-1;
                    var mdiv = divlist[gid];
                    $(mdiv).append(__genDiv(element.Res_id,element.Res_class,element.Res_bg_color,element.Res_img,element.Res_name,element.Res_url,element.Res_open_type));
                });

                if ($(cdiv1).html() != "") {
                    var mdiv1 = __genUI(Group1)
                    $(mdiv1).append(cdiv1)
                    $("#h-system-service").html(mdiv1)
                }else{
                    $("#h-system-service").remove()
                }

                if ($(cdiv2).html() !=""){
                    var mdiv2 = __genUI(Group2)
                    $(mdiv2).append(cdiv2)
                    $("#h-mas-service").html(mdiv2)
                }else{
                    $("#h-mas-service").remove()
                }

                if ($(cdiv3).html() !=""){
                    var mdiv3 = __genUI(Group3)
                    $(mdiv3).append(cdiv3);
                    $("#h-other-service").html(mdiv3)
                }else{
                    $("#h-other-service").remove();
                }

                if (TypeId == 1){
                    $(".hzwy23div").click(function () {
                        Hutils.goEntrySubSystem(this)
                    })
                } else if (TypeId == 0) {
                    $(".hzwy23div").click(function(){
                        Hutils.go_entry(this)
                    })
                }

                $(function() {
                    //取消水平滑动的插件
                    //$.StartScreen();
                    var tiles = $(".tile, .tile-small, .tile-sqaure, .tile-wide, .tile-large, .tile-big, .tile-super");
                    $.each(tiles, function() {
                        var tile = $(this);
                        setTimeout(function() {
                            tile.css({
                                opacity: 1,
                                "-webkit-transform": "scale(1)",
                                "transform": "scale(1)",
                                "-webkit-transition": ".3s",
                                "transition": ".3s"
                            });
                        }, Math.floor(Math.random() * 500));
                    });
                    $(".tile-group").animate({
                        left: 0
                    });
                });
            },
        });
    },
};

// 树形插件
(function($){
    $.fn.Htree = function(param){
        // 1. 获取top节点
        // 2. 获取id，text，upId，将其他属性设置成data-属性
        // 3. 生成tree。
        // 4. 绑定单击按钮
        // 5. 伸缩按钮图标
        // function list:
        // 6. 删除节点
        // 7. 新增节点
        // 8. 更新节点
        // attr 为0 表示叶子,为1 表示节点

        // 保留节点索引
        var $this = this;

        /*
         * 插件默认参数列表。
         * */
        var __DEFAULT = {
            data: "",
            fontSize:"13px",
            showLiHeight:"30px",
            showFontSize:"14px",
            iconColor:"#030202",
            showLevel:3,

            onChange:function (obj) {
                console.log("没有注册点击函数")
            },
        };
        var __autoAttr = false;

        $.extend(true,__DEFAULT,param);

        if (__DEFAULT.data.length > 0 && __DEFAULT.data[0].attr == undefined) {
            __autoAttr = true;
        }

        var getEvent = function(){

            if(window.event)    {
                return window.event;
            }
            var func = getEvent.caller;

            while( func != null ){
                var arg0 = func.arguments[0];
                if(arg0){
                    if((arg0.constructor==Event || arg0.constructor ==MouseEvent
                        || arg0.constructor==KeyboardEvent)
                        ||(typeof(arg0)=="object" && arg0.preventDefault
                        && arg0.stopPropagation)){
                        return arg0;
                    }
                }
                func = func.caller;
            }
            return null;
        };

        // 1.get top node, and sort array
        function sortTree(a){

            // load result sorted
            var list = [];

            // get select's options
            // append it to new select which simulate by ul li
            if (Object.prototype.toString.call(a) == '[object Array]'){

            } else {
                return [];
            }

            //set max dept val
            var MAXDEPT = 8;

            var INDEX = 1;

            function getRoots(arr){
                var Roots = [];
                for(var i = 0; i < arr.length;i++){
                    var rootFlag = true
                    for ( var j = 0; j < arr.length;j++){
                        if (arr[i].upId == arr[j].id){
                            rootFlag = false
                            break
                        }
                    }
                    if (rootFlag == true){
                        Roots.push(arr[i])
                    }
                }
                return Roots
            }

            function traversed(node,arr){
                if (++INDEX > MAXDEPT){
                    console.log("递归超过8层,为保护计算机,退出递归");
                    return
                }
                for (var i = 0; i < arr.length; i++){
                    if (node == arr[i].upId){
                        arr[i].dept = INDEX
                        list.push(arr[i])
                        traversed(arr[i].id,arr)
                    }
                }
                INDEX--;
            }

            function listElem(roots,arr){
                for (var i = 0; i < roots.length; i++){
                    roots[i].dept = INDEX
                    list.push(roots[i])
                    traversed(roots[i].id,arr)
                }
            }

            listElem(getRoots(a),a)

            return list
        }
        // 2. set data-*
        // 3. genUI
        function genTreeUI(a){
            var opt = "<ul class='col-sm-12 col-md-12 col-lg-12'>"
            for(var i = 0; i < a.length; i++){
                var pd = parseInt(a[i].dept)*20 - 10
                if (isNaN(pd)){
                    pd = 10
                }

                if (a[i].attr == "0") {
                    if  ( parseInt(a[i].dept) <= __DEFAULT.showLevel) {
                        // 叶子信息
                        opt += '<li data-id="'+a[i].id+'" data-dept="'+a[i].dept+'" style="margin:0px; text-align: left;font-weight:500;padding-left:'+pd+'px; height:'+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; font-size: '+__DEFAULT.showFontSize+'; cursor: pointer;position: relative;">' +
                            '<hzw style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; width: 20px;cursor: pointer ;display: inline-block">' +
                            '<i class="icon-leaf" style="color: green;"></i>' +
                            '</hzw>' +
                            '<span class="HTreeLi" style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; position: absolute;">'+a[i].text+'</span></li>'

                    } else {
                        // 叶子信息
                        opt += '<li data-id="'+a[i].id+'" data-dept="'+a[i].dept+'" style="margin:0px; text-align: left;font-weight:500;padding-left:'+pd+'px; height:'+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; font-size: '+__DEFAULT.showFontSize+'; cursor: pointer;position: relative; display: none">' +
                            '<hzw style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; width: 20px;cursor: pointer ;display: inline-block">' +
                            '<i class="icon-leaf" style="color: green;"></i>' +
                            '</hzw>' +
                            '<span class="HTreeLi" style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; position: absolute;">'+a[i].text+'</span></li>'

                    }

                } else {
                    if (parseInt(a[i].dept) < __DEFAULT.showLevel){
                        // 节点信息
                        opt += '<li data-id="'+a[i].id+'" data-dept="'+a[i].dept+'" style="margin:0px; text-align: left;font-weight:500;padding-left:'+pd+'px; height:'+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; font-size: '+__DEFAULT.showFontSize+'; cursor: pointer;position: relative;">' +
                            '<hzw class="HTreeshowOrHideIconHzw" style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; width: 20px;cursor: cell ;display: inline-block">' +
                            '<i style="color: #ccb008;" class="icon-folder-open"> </i>' +
                            '</hzw>' +
                            '<span class="HTreeLi" style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; position: absolute;">'+a[i].text+'</span></li>'
                    } else  if (parseInt(a[i].dept) == __DEFAULT.showLevel) {
                        // 节点信息
                        opt += '<li data-id="'+a[i].id+'" data-dept="'+a[i].dept+'" style="margin:0px; text-align: left;font-weight:500;padding-left:'+pd+'px; height:'+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; font-size: '+__DEFAULT.showFontSize+'; cursor: pointer;position: relative;">' +
                            '<hzw class="HTreeshowOrHideIconHzw" style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; width: 20px;cursor: cell ;display: inline-block">' +
                            '<i style="color: #ccb008;" class="icon-folder-close"> </i>' +
                            '</hzw>' +
                            '<span class="HTreeLi" style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; position: absolute;">'+a[i].text+'</span></li>'
                    } else {
                        // 节点信息
                        opt += '<li data-id="'+a[i].id+'" data-dept="'+a[i].dept+'" style="margin:0px; text-align: left;font-weight:500;padding-left:'+pd+'px; height:'+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; font-size: '+__DEFAULT.showFontSize+'; cursor: pointer;position: relative;display: none;">' +
                            '<hzw class="HTreeshowOrHideIconHzw" style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; width: 20px;cursor: cell ;display: inline-block">' +
                            '<i style="color: #ccb008;" class="icon-folder-open"> </i>' +
                            '</hzw>' +
                            '<span class="HTreeLi" style="height: '+__DEFAULT.showLiHeight+'; line-height: '+__DEFAULT.showLiHeight+'; position: absolute;">'+a[i].text+'</span></li>'
                    }

                }
            }
            opt +='</ul>'
            return opt;
        }

        // 绑定伸缩按钮
        function showOrHide(e){
            $.notifyClose()
            var dept = $(e).attr("data-dept")
            var nextObj = $(e).next()
            var nextDept = $(nextObj).attr("data-dept")
            var nextDisplay = $(nextObj).css("display")

            if (nextDisplay == "none" && parseInt(nextDept)>parseInt(dept)){

                $(e).find("i").removeClass("icon-folder-close").addClass("icon-folder-open")

                $(e).nextAll().each(function(index,element){
                    if (parseInt(dept)+1==parseInt($(element).attr("data-dept"))){
                        // 显示
                        if ($(element).find("i").hasClass("icon-folder-open")) {
                            $(element).find("i").addClass("icon-folder-close").removeClass("icon-folder-open")
                        }
                        $(element).fadeIn(400);
                    }else if (parseInt(dept)+1 < parseInt($(element).attr("data-dept"))){
                        // 隐藏
                        if ($(element).find("i").hasClass("icon-folder-open")) {
                            $(element).find("i").addClass("icon-folder-close").removeClass("icon-folder-open")
                        }
                        $(element).fadeOut(200);
                    } else {
                        return false
                    }
                })
            }else if (nextDisplay == "none" && parseInt(nextDept)<=parseInt(dept)){
                return
            }else if (nextDisplay != "none" && parseInt(nextDept)>parseInt(dept)){

                $(e).find("i").removeClass("icon-folder-open").addClass("icon-folder-close")

                $(e).nextAll().each(function(index,element){
                    if (parseInt(dept)<parseInt($(element).attr("data-dept"))){
                        if ($(element).find("i").hasClass("icon-folder-open")) {
                            $(element).find("i").addClass("icon-folder-open").removeClass("icon-folder-close")
                        }
                        $(element).fadeOut(200);
                    }else if (parseInt(dept)>=parseInt($(element).attr("data-dept"))){
                        return false
                    }
                })
            } else {
                $.Notify({
                    title:"温馨提示:",
                    message:"这个节点下边没有叶子信息",
                    type:"info",
                });

                if ($(e).find("i").hasClass("icon-folder-open")) {
                    $(e).find("i").removeClass("icon-folder-open").addClass("icon-folder-close")
                } else if ($(e).find("i").hasClass("icon-folder-close")) {
                    $(e).find("i").removeClass("icon-folder-close").addClass("icon-folder-open")
                }
                return
            }
        }

        var li = sortTree(__DEFAULT.data)
        var opt = genTreeUI(li)

        $this.html(opt)

        /*
         * 如果这个节点没有下层信息，则将这个层级的伸缩按钮去掉。
         * */
        if (__autoAttr){
            $this.find("ul li").each(function(index,element){
                var curDept = parseInt($(element).attr("data-dept"));
                var nextDept = parseInt($(element).next().attr("data-dept"));
                if (curDept>=nextDept || isNaN(nextDept)){
                    $(element).find("hzw").html("<i class='icon-leaf' style='color: green;'></i>").removeClass("HTreeshowOrHideIconHzw").css("cursor","pointer");
                }
            });
        }

        /*
         * 给ul中每一行li绑定点击事件
         * */
        $this.find("ul li").on("click",function(){
            $.notifyClose()
            $this.find(".HTreeLi").css("color","")
            $(this).find("span").css("color","red")
            $this.attr("data-selected",$(this).attr("data-id"))
            __DEFAULT.onChange(this)
        });

        $this.find("ul li").on("mouseover",function () {
            $(this).css({
                "background-color":"#cccccc",
            })
        });

        $this.find("ul li").on("mouseout",function () {
            $(this).css({
                "background-color":"",
            })
        });

        /*
         * 给伸缩按钮绑定单击事件
         * */
        $this.find(".HTreeshowOrHideIconHzw").on("click",function () {
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }
            showOrHide($(this).parent())
        });

    };

    $.fn.Hselect = function(param){
        var sel = this
        var obj = document.createElement("div")

        if ( $(sel).attr("hselect") == "true"){
            // 重复初始化Hselect
            var hselect = $(sel).next()
            var displaycss = $(hselect).css("display")
            $(obj).attr("style",$(sel).attr("style"));
            $(obj).css("display",displaycss)

            $(hselect).remove()
            $(sel).html("");
        } else {
            // 第一次初始化Hselect
            $(obj).attr("style",$(sel).attr("style"));
        }
        //init div css
        //get parent class to it
        //get parent css to it
        $(obj).addClass($(sel).attr("class"));
        $(obj).css({"padding":"0px","border":"none"});

        $(sel).attr("hselect","true");
        // default parameters
        var __DEFAULT = {
            data: "",
            height:"26px",
            width:"100%",
            border:"#ccc solid 1px",
            fontSize:"13px",
            borderRadius:"5px",
            bgColor:"white",
            placeholder:"<i style='color: #959595;font-size: 12px;'>--点击树形多选框,设置选项--</i>",

            showLiHeight:"30px",
            showHeight:"230px",
            showBorder:"",
            showFontSize:"14px",
            iconColor:"#ff5763",

            // 值改变触发事件
            onclick:"",

            // select中默认值
            value:"hselectdefault",

            // 是否禁止选择
            disabled:false,

            // checkbox选择框,
            // 默认禁用选择框,如果设置成true,则显示考勤选怎狂
            checkbox:false,

            // 是否启动自动化节点/叶子处理
            // 如果开启,则将会把末级节点全部处理成叶子
            autoAttr: true,

            // 默认开启节点可以选中
            // 如果设置成false, 则不允许选择节点,在单击节点时,将会显示或隐藏子节点
            nodeSelect: true,

            dashed:true,

        };

        // 记录是否为层级结构
        // true表示是层级结构
        // false表示非层级结构
        var level_flag = false;

        $.extend(true,__DEFAULT,param);

        if (__DEFAULT.data.length > 0 && __DEFAULT.data[0].attr != undefined) {
            __DEFAULT.autoAttr = false;
        }

        // set showBorder to border style
        if (__DEFAULT.showBorder==""){
            __DEFAULT.showBorder = __DEFAULT.border
        };

        var getEvent = function(){

            if(window.event)    {
                return window.event;
            }
            var func = getEvent.caller;

            while( func != null ){
                var arg0 = func.arguments[0];
                if(arg0){
                    if((arg0.constructor==Event || arg0.constructor ==MouseEvent
                        || arg0.constructor==KeyboardEvent)
                        ||(typeof(arg0)=="object" && arg0.preventDefault
                        && arg0.stopPropagation)){
                        return arg0;
                    }
                }
                func = func.caller;
            }
            return null;
        };
        /*
         * This function sort array.
         * Accept One Array Variable.
         * */
        function sortTree(a){

            // load result sorted
            var list = [];

            // get select's options
            // append it to new select which simulate by ul li
            if (Object.prototype.toString.call(a) == '[object Array]'){
                $(sel).find("option").each(function(index,element){
                    var ijs = {}
                    ijs.id = $(element).val();
                    ijs.text = $(element).text()
                    a.push(ijs)
                })
            } else {
                $(sel).find("option").each(function(index,element){
                    var ijs = {}
                    ijs.id = $(element).val();
                    ijs.text = $(element).text()
                    list.push(ijs)
                })
                return list
            }

            //set max dept val
            var MAXDEPT = 8;

            var INDEX = 1;

            function getRoots(arr){
                var Roots = [];
                for(var i = 0; i < arr.length;i++){
                    var rootFlag = true
                    for ( var j = 0; j < arr.length;j++){
                        if (arr[i].upId == arr[j].id){
                            rootFlag = false
                            break
                        }
                    }
                    if (rootFlag == true){
                        Roots.push(arr[i])
                    }
                }
                return Roots
            }

            function traversed(node,arr){
                if (++INDEX > MAXDEPT){
                    console.log("递归超过8层,为保护计算机,退出递归");
                    return
                }
                for (var i = 0; i < arr.length; i++){

                    if (node == arr[i].upId){
                        arr[i].dept = INDEX
                        list.push(arr[i])
                        traversed(arr[i].id,arr)
                    }
                }
                INDEX--;
            }

            function listElem(roots,arr){
                for (var i = 0; i < roots.length; i++){
                    roots[i].dept = INDEX
                    list.push(roots[i])
                    traversed(roots[i].id,arr)
                }
            }

            listElem(getRoots(a),a)

            return list
        };

        // 生成树形线
        // 辅助查看
        function genImgLine(dept,attr) {
            var i = 0;
            var img_src = document.createElement("yph")

            if (dept == 0 || dept == undefined){
                var img = document.createElement("img")
                img.style.height = __DEFAULT.showLiHeight;
                img.style.width = "20px";
                img.style.lineHeight = __DEFAULT.showLiHeight;
                if (__DEFAULT.dashed) {
                    img.setAttribute("src","/images/icon_select/line.gif");
                    img_src.appendChild(img);
                } else {
                    img.setAttribute("src","/images/icon_select/empty.gif");
                    img.style.width = "8px";
                    img_src.appendChild(img);
                }
                return img
            }

            for (;i<dept-1;i++) {
                var img = document.createElement("img")
                img.style.height = __DEFAULT.showLiHeight;
                img.style.width = "20px";
                img.style.lineHeight = __DEFAULT.showLiHeight;
                if (__DEFAULT.dashed) {
                    img.setAttribute("src","/images/icon_select/line.gif");
                    img_src.appendChild(img);
                } else {
                    if (i == 0) {
                        img.style.width = "8px";
                    }
                    img.setAttribute("src","/images/icon_select/empty.gif");
                    img_src.appendChild(img);
                }
            }

            if (attr == "1") {
                // 节点
                var img = document.createElement("img")
                img.style.height = __DEFAULT.showLiHeight;
                img.style.width = "20px";
                img.style.lineHeight = __DEFAULT.showLiHeight;
                if (__DEFAULT.dashed) {
                    img.setAttribute("src","/images/icon_select/minus.gif");
                    img_src.appendChild(img);
                } else {
                    if (i == 0) {
                        img.style.width = "8px";
                    }
                    img.setAttribute("src","/images/icon_select/empty.gif");
                    img_src.appendChild(img);
                }
            } else if (attr == "0") {
                // 叶子
                var img = document.createElement("img")
                img.style.height = __DEFAULT.showLiHeight;
                img.style.width = "20px";
                img.style.lineHeight = __DEFAULT.showLiHeight;
                if (__DEFAULT.dashed) {
                    img.setAttribute("src","/images/icon_select/join.gif");
                    img_src.appendChild(img);
                } else {
                    if (i == 0) {
                        img.style.width = "8px";
                    }
                    img.setAttribute("src","/images/icon_select/empty.gif");
                    img_src.appendChild(img);
                }
            } else {
                // 没有指定叶子,节点标识
                // 节点
                var img = document.createElement("img")
                img.style.height = __DEFAULT.showLiHeight;
                img.style.width = "20px";
                img.style.lineHeight = __DEFAULT.showLiHeight;
                if (__DEFAULT.dashed) {
                    img.setAttribute("src","/images/icon_select/minus.gif");
                    img_src.appendChild(img);
                } else {
                    if (i == 0) {
                        img.style.width = "8px";
                    }
                    img.setAttribute("src","/images/icon_select/empty.gif");
                    img_src.appendChild(img);
                }
            }
            return img_src
        }

        // 产生一行li
        function genULli(attr,id,dept,text) {
            var li = document.createElement("li")
            li.setAttribute("data-attr",attr);
            li.setAttribute("data-id",id);
            li.setAttribute("data-dept",dept);
            li.style.margin = "0px";
            li.style.textAlign = "left";
            li.style.fontSize = "500";
            li.style.height = __DEFAULT.showLiHeight;
            li.style.lineHeight = __DEFAULT.showLiHeight;
            li.style.fontSize = __DEFAULT.showFontSize;
            li.style.cursor = "pointer";
            li.style.position = "relative";

            var hzw = document.createElement("hzw")
            hzw.setAttribute("class","HshowOrHideIconHzw");
            hzw.style.height = __DEFAULT.showLiHeight;
            hzw.style.lineHeight = __DEFAULT.showLiHeight;
            hzw.style.width = "20px";
            hzw.style.display = "inline-block";
            var i = document.createElement("i")
            if (attr == "0") {
                i.setAttribute("class","icon-leaf")
                i.style.color = "green";
            } else {
                i.setAttribute("class","icon-folder-open");
                i.style.color = "#ccb008";
            }
            hzw.appendChild(i)

            var span = document.createElement("span");
            span.style.height = __DEFAULT.showLiHeight;
            span.style.lineHeight = __DEFAULT.showLiHeight;
            span.style.position = "absolute";
            span.innerHTML = text;


            var img_src = genImgLine(dept,attr)
            li.appendChild(img_src)


            // 如果selectBox开启多选,则显示勾选框
            if (__DEFAULT.checkbox) {
                var input = document.createElement("input")
                input.setAttribute("name","select_check")
                input.setAttribute("type","checkbox")
                input.style.width = "20px";
                input.style.marginTop = "0px";
                li.appendChild(input);
            }

            li.appendChild(hzw);
            li.appendChild(span)
            return li
        }

        // 替换select显示区域.
        // 重新构建选择框
        function genDiv() {
            var div = document.createElement("div")
            div.style.cursor = "pointer";
            div.style.backgroundColor = __DEFAULT.bgColor;
            div.style.textAlign = "left !important";
            div.style.width = __DEFAULT.width;
            div.style.height = __DEFAULT.height;
            div.style.lineHeight = __DEFAULT.height;
            div.style.paddingLeft = "10px";
            div.style.display = "inline-block"
            div.style.border = "#ccc solid 1px";
            div.style.borderRadius = __DEFAULT.borderRadius;
            div.setAttribute("class","HshowSelectValue")
            div.style.position = "relative";

            var span = document.createElement("span")
            span.style.height = __DEFAULT.height;
            span.style.fontSize = __DEFAULT.fontSize;
            span.style.overflow = "hidden";
            span.style.wordBreak = "keep-all";
            span.style.whitespace = "nowrap";
            span.style.textOverflow = "ellipsis";
            span.style.display = "inline-block";
            span.innerHTML = __DEFAULT.placeholder;

            var hzw = document.createElement("hzw")
            hzw.style.position = "absolute";
            hzw.style.width = "20px";
            hzw.style.right = "0px";
            // hzw.style.float = "right";
            hzw.style.height = __DEFAULT.height;
            hzw.style.lineHeight = __DEFAULT.height;

            var i = document.createElement("i")
            i.style.borderColor = "#888 transparent transparent transparent";
            i.style.borderWidth = "5px 4px 0px 4px"
            i.style.height = "0";
            i.style.left = "50%";
            i.style.marginLeft = "-4px";
            i.style.marginTop = "-3px";
            i.style.position = "absolute"
            i.style.top = "50%";
            i.style.width = "0px";
            i.style.borderStyle = "solid";
            hzw.appendChild(i)

            div.appendChild(span)
            div.appendChild(hzw)
            return div;
        }

        // 根据数据,
        // 生成需要展示的数据
        function genTreeUI(a){

            var div = document.createElement("div")
            div.setAttribute("class","HselectShowAreaHuangZhanWei");
            div.style.whitespace = "nowrap";
            div.style.backgroundColor = "#fefefe";
            div.style.border = __DEFAULT.showBorder
            div.style.borderRadius = " 3px";
            div.style.display = "none";
            div.style.position = "fixed";
            div.style.zIndex  = 9999;

            var input = document.createElement("input")
            input.style.border = "#6699CC solid 1px";
            input.style.paddingLeft = "5px";
            input.style.margin = "5px 5px";
            input.style.height = __DEFAULT.showLiHeight;
            div.appendChild(input)

            var opt = document.createElement("ul")
            opt.style.zIndex = 9999;
            opt.style.padding = "0px";
            opt.style.listStyle = "none";
            opt.style.margin = "0px";
            opt.style.maxHeight = __DEFAULT.showHeight;
            opt.style.overflow = "auto";

            for(var i = 0; i < a.length; i++){
                var pd = parseInt(a[i].dept)
                if (pd > 1) {
                    level_flag = true;
                }
                if (isNaN(pd)){
                    pd = 1
                }
                opt.appendChild(genULli(a[i].attr,a[i].id,a[i].dept,a[i].text));
            }
            div.appendChild(opt)
            return div
        };

        function adjustImg() {
            if (__DEFAULT.dashed && level_flag) {
                var ul = ui.getElementsByTagName("ul")
                if (ul.length > 0) {
                    var yph = ul[0].firstElementChild.getElementsByTagName("yph");
                    if (yph.length > 0) {
                        if (yph[0].firstElementChild.getAttribute("src") == "/images/icon_select/minus.gif"){
                            yph[0].firstElementChild.setAttribute("src","/images/icon_select/minusbottom.gif")
                        }
                        if (yph[0].firstElementChild.getAttribute("src") == "/images/icon_select/join.gif"){
                            yph[0].firstElementChild.setAttribute("src","/images/icon_select/upjoinbutton.gif")
                        }
                    }
                }
                var children = ul[0].childNodes;
                var length = children.length;
                var img_index = children[length-1].getAttribute("data-dept");
                for (var i = 1; i < length; i++){
                    var curli = children[i]
                    var nextli = children[i+1]
                    if (nextli == undefined) {
                        var curdept = curli.getAttribute("data-dept")
                        var curNode = curli.getAttribute("data-attr")
                        if (curNode == "0") {
                            var yph = curli.getElementsByTagName("yph")
                            if (yph.length >= 0 && !isNaN(curdept)){
                                yph[0].childNodes[parseInt(curdept)-1].setAttribute("src","/images/icon_select/joinbottom.gif");
                            }
                        }
                    } else {
                        var curdept = curli.getAttribute("data-dept")
                        var nextdept = nextli.getAttribute("data-dept")

                        var curNode = curli.getAttribute("data-attr")
                        var nextNode = nextli.getAttribute("data-attr")
                        if (parseInt(curdept) > parseInt(nextdept)){
                            // 下级节点是节点，当前选中想层级大于下级接节点
                            if (curNode == "0") {
                                var yph = curli.getElementsByTagName("yph")
                                if (yph.length >= 0 && !isNaN(curdept)){
                                    yph[0].childNodes[parseInt(curdept)-1].setAttribute("src","/images/icon_select/joinbottom.gif");
                                }
                            }
                        }
                    }

                    var yph = children[length-i].getElementsByTagName("yph")
                    if (img_index > 0) {
                        if (children[length-i].getAttribute("data-attr") == "1") {
                            var dept = children[length-i].getAttribute("data-dept");
                            if (parseInt(dept) < img_index) {
                                img_index = parseInt(dept);
                                yph[0].lastElementChild.setAttribute("src","/images/icon_select/minustop.gif")
                            }
                        }
                        for (var j = 0; j < img_index - 1; j++){
                            yph[0].childNodes[j].setAttribute("src","/images/icon_select/empty.gif")
                        }
                    }

                    if (children[length-i].getAttribute("data-attr") == "1") {
                        var imgs = children[length-i].getElementsByTagName("img")
                        if (imgs.length > 0) {
                            imgs[imgs.length-1].setAttribute("class","hzwy23-images-click")
                        }
                    }
                }


                // 处理第一个
                var yph = children[0].getElementsByTagName("yph")
                if (img_index > 0) {
                    if (children[0].getAttribute("data-attr") == "1") {
                        var dept = children[0].getAttribute("data-dept");
                        if (parseInt(dept) < img_index) {
                            img_index = parseInt(dept);
                            yph[0].lastElementChild.setAttribute("src","/images/icon_select/minusright.gif")
                        }
                    }
                    for (var j = 0; j < img_index - 1; j++){
                        yph[0].childNodes[j].setAttribute("src","/images/icon_select/empty.gif")
                    }
                }

                if (children[0].getAttribute("data-attr") == "1") {
                    var imgs = children[0].getElementsByTagName("img")
                    if (imgs.length > 0) {
                        imgs[imgs.length-1].setAttribute("class","hzwy23-images-click")
                    }
                }
            }
        }

        function showUp(e){
            var dept = $(e).attr("data-dept")
            $(e).prevAll().each(function(index,element){
                if (parseInt(dept)>parseInt($(element).attr("data-dept"))){
                    $(element).show();
                    dept = $(element).attr("data-dept")
                }
            })
        };

        // 初始化select框
        // 然后隐藏select框
        function initSelect(selObj,arr){
            for (var i = 0; i < arr.length; i++){
                var opt = document.createElement("option")
                opt.setAttribute("value",arr[i].id)
                opt.innerHTML = arr[i].text;
                selObj.append(opt)
            }
            $(selObj).hide();
        };


        function modifyIcon(e,showOrHide) {
            if (showOrHide) {
                // 显示
                if ($(e).find("i").hasClass("icon-folder-close")) {
                    $(e).find("i").addClass("icon-folder-open").css({"color":"#ccb008"}).removeClass("icon-folder-close");
                }
            } else {
                // 隐藏
                if ($(e).find("i").hasClass("icon-folder-open")) {
                    $(e).find("i").addClass("icon-folder-close").removeClass("icon-folder-open");
                }
            }
        }

        function modifyImgLine(e,showOrHide) {
            if (showOrHide && __DEFAULT.dashed && $(e).find("img").last() != undefined) {
                // 显示
                if ($(e).find("img").last().attr("src") == "/images/icon_select/plus.gif") {
                    $(e).find("img").last().attr("src","/images/icon_select/minus.gif")
                    return
                }
                if ($(e).find("img").last().attr("src") == "/images/icon_select/plusbottom.gif") {
                    $(e).find("img").last().attr("src","/images/icon_select/minusbottom.gif")
                    return
                }
                if ($(e).find("img").last().attr("src") == "/images/icon_select/plustop.gif") {
                    $(e).find("img").last().attr("src","/images/icon_select/minustop.gif")
                    return
                }
                if ($(e).find("img").last().attr("src") == "/images/icon_select/plusright.gif") {
                    $(e).find("img").last().attr("src","/images/icon_select/minusright.gif")
                    return
                }
            } else if ( __DEFAULT.dashed && $(e).find("img").last() != undefined) {
                // 隐藏
                if ($(e).find("img").last().attr("src") == "/images/icon_select/minus.gif") {
                    $(e).find("img").last().attr("src","/images/icon_select/plus.gif")
                    return
                }
                if ($(e).find("img").last().attr("src") == "/images/icon_select/minusbottom.gif") {
                    $(e).find("img").last().attr("src","/images/icon_select/plusbottom.gif")
                    return
                }
                if ($(e).find("img").last().attr("src") == "/images/icon_select/minustop.gif") {
                    $(e).find("img").last().attr("src","/images/icon_select/plustop.gif")
                    return
                }
                if ($(e).find("img").last().attr("src") == "/images/icon_select/minusright.gif") {
                    $(e).find("img").last().attr("src","/images/icon_select/plusright.gif")
                    return
                }
            }
        }

        // 隐藏过现实下级信息
        function showOrHide(e){
            // 参数是被点击的哪一行

            // 获取当前被点击的哪一行的层级
            var dept = $(e).attr("data-dept")

            // 获取下一行对象
            var nextObj = $(e).next()

            // 获取下一行层级
            var nextDept = $(nextObj).attr("data-dept")

            // 获取下一行状态
            var nextDisplay = $(nextObj).css("display")

            // 如果下一行是隐藏状态, 则修改成现实
            // 如果下一行是显示状态,则修改为隐藏
            if (nextDisplay == "none" && parseInt(nextDept) > parseInt(dept)){

                // 如果检查到图标是关闭状态状态
                // 则修改为打开状态
                modifyIcon(e,true)

                //修改线条
                //如果展开,则替换成 减号
                //如果隐藏,则替换成 加号
                modifyImgLine(e,true)

                $(e).nextAll().each(function(index,element) {
                    if (parseInt(dept)+1==parseInt($(element).attr("data-dept"))){

                        $(element).show();
                        modifyIcon(element,false)
                        modifyImgLine(element,false)

                    } else if (parseInt(dept)+1 < parseInt($(element).attr("data-dept"))) {
                        // 切换成关闭状态
                        modifyIcon(element,false)

                        $(element).hide();

                        modifyImgLine(element,false)
                    } else {
                        // 如果遇到下一层级的高度小于被选中的层级高度
                        // 则退出执行过程
                        return false
                    }
                })
            }else if (nextDisplay == "none" && parseInt(nextDept)<=parseInt(dept)){
                // 如果下一行被隐藏
                // 且下一行的层级下雨当前层级
                // 则表示没有下一层极,直接退出
                return
            }else if (nextDisplay != "none" && parseInt(nextDept)>parseInt(dept)){
                modifyImgLine(e,false);
                modifyIcon(e,false);

                $(e).nextAll().each(function(index,element){
                    if (parseInt(dept)<parseInt($(element).attr("data-dept"))){
                        modifyImgLine(e,false)
                        modifyIcon(e,false)
                        $(element).hide();
                    }else if (parseInt(dept)>=parseInt($(element).attr("data-dept"))){
                        return false
                    }
                })
            }else {
                return
            }
        };

        // 如果非层级结构
        // 则删除虚线和图标
        function adjustnormaltree() {
            if (!level_flag) {
                $(ui).find("li").each(function (index, element) {
                    $(element).find("hzw").remove();
                    $(element).find("img").remove();
                    $(element).css("padding-left","8px");
                })
            }
        };

        function adjustAttr() {
            // 修正叶子节点
            // 如果设置了成员属性 attr == false, 则执行这个操作
            if (__DEFAULT.autoAttr) {
                $(ui).find("ul li").each(function(index,element){
                    var curDept = parseInt($(element).attr("data-dept"))
                    var nextDept = parseInt($(element).next().attr("data-dept"))
                    if (curDept>=nextDept || isNaN(nextDept)){
                        $(element).attr("data-attr","0");
                        $(element).find("hzw i").removeClass("icon-folder-open icon-folder-close").addClass("icon-leaf").css("color","green");
                        if (__DEFAULT.dashed && level_flag) {
                            $(element).find("img").last().attr("src","/images/icon_select/join.gif");
                        }
                    } else {
                        $(element).attr("data-attr","1")
                    }
                });
            }
        }

        function initDefaultValue() {
            var flag = false;
            $(sel).find("option").each(function (index, element) {
                if ($(element).attr("value") == __DEFAULT.value) {
                    flag = true;
                }
            })

            if (__DEFAULT.value != "hselectdefault" && flag == true){
                $(sel).val(__DEFAULT.value);
                var text = $(sel).find("option:selected").text()
                $(obj).find(".HshowSelectValue span").html(text)
            } else {
                $(sel).val("");
            }
        };

        var ui = genTreeUI(sortTree(__DEFAULT.data));
        // 如果没有层级,则清除第一层级的叶子
        adjustnormaltree();
        // 调整li属性,
        // 如果没有设置attr这个字段,则自动匹配
        adjustAttr();

        // 设置虚线
        adjustImg()

        var pdiv = genDiv();

        initSelect(sel,__DEFAULT.data);
        obj.appendChild(pdiv)
        obj.appendChild(ui);

        $(sel).after(obj);

        // 清除select的默认选中状态，确保select初始化后，没有任何值被选中
        // 如果在初始化Hselect时，指定了初始值，则使用初始值
        initDefaultValue();

        if (__DEFAULT.disabled){
            // 如果select禁止选择,
            // 不需要绑定触发事件
            $(obj).find(".HshowSelectValue").css("background-color","#f5f5f5");
            return
        }

        // 让搜索框获取焦点
        $(obj).find("input:eq(0)").focus();

        // input 框中输入事件，当用户在Hselect的下拉框中搜索时，触发这个事件
        $(obj).find("input:eq(0)").on('input',function(){
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }

            var inpText = $(this).val();
            if (inpText == ""){
                $(obj).find("ul li").show();
                return
            }
            $(obj).find("ul li").each(function(index,element){
                if ($(element).find("span").html().indexOf(inpText)>=0){
                    $(element).show()
                    showUp(element)
                }else{
                    $(element).hide()
                }
            })
        });

        // 当用户在搜索框中点击鼠标左键时，触发这个事件。
        $(obj).find("input:eq(0)").on('click',function(){
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }
            $(this).focus();
        });


        // checkbox选择框,绑定事件
        $(obj).find("input[name='select_check']").on("click",function () {
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }

            // 获取当前层级的li对象
            var curli = $(this).parent();

            // 判断当前li的层级
            var dept = $(curli).attr("data-dept");

            // 获取下一个li对象层级
            var nextDept = $(curli).next().attr("data-dept");

            // 如果下一个层级的高度大于当前层级高度,
            // 则选中所有的下级
            if (parseInt(dept) < parseInt(nextDept)){
                if ($(this).is(":checked")) {
                    $(curli).nextAll().each(function (index, element) {
                        if (parseInt($(element).attr("data-dept")) > parseInt(dept)) {
                            $(element).find("input").prop("checked",true)
                        } else {
                            return false
                        }
                    })
                } else {
                    $(curli).nextAll().each(function (index, element) {
                        if (parseInt($(element).attr("data-dept")) > parseInt(dept)) {
                            $(element).find("input").prop("checked",false)
                        } else {
                            return false
                        }
                    })
                }
            }
        });

        // 图标, 点击隐藏或显示子菜单
        $(obj).find(".HshowOrHideIconHzw").on("click",function(){
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }

            // 当选中的是叶子的时候,不触发事件
            if ($(this).find("i").hasClass("icon-leaf")) {
                return
            }

            showOrHide($(this).parent())
        });

        $(obj).find("li").on('mouseover',function(){
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }

            var ul = $(this).closest("ul")

            $(ul).find("li").css({
                "background-color":"",
                "color":""
            })

            $(this).css({
                "background-color":"#6699CC",
                "color":"white"
            })
        });

        // 绑定虚线点击事件
        $(obj).find(".hzwy23-images-click").on('click',function(){
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }
            showOrHide($(this).closest("li"))
        });

        $(obj).find("li").on('click',function(){
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }

            if (__DEFAULT.nodeSelect == false && $(this).attr("data-attr") == "1"){
                showOrHide($(this))
                return
            }

            if (__DEFAULT.checkbox == true) {
                if ($(this).attr("data-attr") == "0"){
                    // 如果开启了checkbox, 则表示seledct不允填写值
                    // 则提供多选,如果这个叶子节点已经被选中
                    // 则取消选中状态
                    // 否则选中这个叶子
                    if ($(this).find("input").is(":checked")) {
                        $(this).find("input").prop("checked",false)
                    } else {
                        $(this).find("input").prop("checked",true)
                    }
                } else if ($(this).attr("data-attr") == "1") {
                    showOrHide($(this))
                    return
                }
            } else {
                // 如果是节点
                // 则不管是否开启checkbox
                // 都一样的处理方式
                var text = $(this).find("span").html();
                var id = $(this).attr("data-id");
                $(sel).val(id);
                $(this).closest("div").prev().find("span").html(text);
                $(this).closest("div").hide();

                $(obj).find(".HshowSelectValue i").css({
                    "border-color":"#888 transparent transparent transparent",
                    "border-width":"5px 4px 0px 4px"
                });

                if (typeof __DEFAULT.onclick == "function"){
                    __DEFAULT.onclick();
                };
            }
        });

        $(obj).find("ul").on('mousewheel',function(){
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }
        });

        $(obj).find(".HshowSelectValue").on('click',function(){
            var showUiStatus = $(obj).find(".HselectShowAreaHuangZhanWei").css("display")
            // 取消后续事件
            if (window.event != undefined){
                window.event.cancelBubble = true;
            } else {
                var event = getEvent()
                event.stopPropagation()
            }

            if (showUiStatus == "none"){
                $(".HselectShowAreaHuangZhanWei").hide()
                $(obj).find("ul li").show();
                var w = $(obj).width();
                $(obj).find(".HselectShowAreaHuangZhanWei").css("min-width",w);
                $(obj).find(".HselectShowAreaHuangZhanWei input:eq(0)").css("min-width",w-12);

                var nextObj = $(this).next();
                $(nextObj).find("input:eq(0)").val("");
                $(nextObj).show();
                $(nextObj).find("input:eq(0)").focus();
                $(nextObj).find("ul").scrollTop(0);
                $(nextObj).find("ul").scrollLeft(0);

                $(obj).find(".HshowSelectValue i").css({
                    "border-color":"transparent transparent #888 transparent",
                    "border-width":"0px 4px 5px 4px"
                });

            }else{
                $(obj).find(".HshowSelectValue i").css({
                    "border-color":"#888 transparent transparent transparent",
                    "border-width":"5px 4px 0px 4px"
                });

                $(obj).find("ul").closest("div").hide();
            }
        });

        $(document).on('click',function(){
            $(obj).find("ul").closest("div").hide();
            $(obj).find(".HshowSelectValue i").css({
                "border-color":"#888 transparent transparent transparent",
                "border-width":"5px 4px 0px 4px"
            });
        });

        //when select was change
        //change show values
        $(sel).on('change',function(){
            var text = $(this).find("option:selected").text()
            $(obj).find(".HshowSelectValue span").html(text)
            if (typeof __DEFAULT.onclick == "function"){
                __DEFAULT.onclick();
            }
        });

    };
}(jQuery));

/*
 * 弹出框效果
 * */
(function($){

    $.extend({
        Notify:function(param){
            var DEFAULT = {
                icon:"icon-ok",
                caption:"",
                title:$.i18n.prop("notify_header_title"),
                message:$.i18n.prop("notify_header_success"),
                content:"",
                type:"success",
                position:null,
                placement: {
                    from: "top",
                    align: "center"
                },
            };

            $.extend(true,DEFAULT,param);
            switch (DEFAULT.type){
                case "success":DEFAULT.icon = "icon-ok";break;
                case "danger":DEFAULT.icon = "icon-remove" ; break;
                case "info" : DEFAULT.icon = "icon-bullhorn";break;
                case "primary": DEFAULT.icon = "icon-bell" ; break;
                case "warning": DEFAULT.icon = "icon-warning-sign"; break;
                default :
                    DEFAULT.icon = "icon-bullhorn"
            }

            $.notify({
                // options
                icon: DEFAULT.icon,
                title: DEFAULT.title,
                message:DEFAULT.message,
                url: '',
                target: '_blank'
            },{
                // settings
                element: 'body',
                position: DEFAULT.position,
                type: DEFAULT.type,
                allow_dismiss: true,
                newest_on_top: true,
                showProgressbar: false,
                placement:DEFAULT.placement,
                offset: 20,
                spacing: 10,
                z_index: 2147483647,
                delay: 3000,
                timer: 1000,
                url_target: '_blank',
                mouse_over: null,
                animate: {
                    enter: 'animated fadeInDown',
                    exit: 'animated fadeOutUp'
                },
                onShow: null,
                onShown: null,
                onClose: null,
                onClosed: null,
                icon_type: 'class',
            });
        },
        HAjaxRequest:function(a){
            var b = {
                type:"get",
                url:"",
                data:"",
                cache:false,
                async:false,
                dataType:"json",
                beforeSend:function () {

                },
                complete:function () {

                },
                error:function(m) {
                    var msg = JSON.parse(m.responseText);
                    jQuery.Notify({
                        title: $.i18n.prop("notify_header_title"),
                        message: msg.error_msg,
                        type: "danger",
                    });
                    console.log("return message is :",msg);
                    console.log("return code is :",msg.error_code);
                    console.log("return details error info:",msg.error_details);
                    console.log("return version: ",msg.version);
                },
                success:function(b){
                }
            };

            $.extend(!0,b,a);

            "delete"==b.type.toLowerCase()?(
                b.data._method="DELETE",
                    $.ajax({
                        type:"post",
                        url:b.url,
                        cache:b.cache,
                        async:b.async,
                        data:b.data,
                        dataType:b.dataType,
                        error:b.error,
                        beforeSend:b.beforeSend,
                        complete:b.complete,
                        success:function(a){
                            b.success(a)}
                    })
            ):$.ajax({
                type:b.type,
                url:b.url,
                cache:b.cache,
                async:b.async,
                data:b.data,
                dataType:b.dataType,
                beforeSend:b.beforeSend,
                complete:b.complete,
                error:b.error,
                success: function(da) {
                    b.success(da)
                },
            })
        },
        Hdownload:function (params) {
            var __DEFAULT = {
                url:"",
                name:"导出数据",
            };
            $.extend(true,__DEFAULT,params);

            if (__DEFAULT.name == "" || __DEFAULT.name==undefined) {
                __DEFAULT.name = "导出参数信息"
            }

            var expname = __DEFAULT.name+".xlsx";

            var x=new XMLHttpRequest();
            x.open("GET",__DEFAULT.url, true);
            x.responseType = 'blob';
            x.onload=function(e){
                download(x.response, expname, "application/vnd.ms-excel" );
            };
            x.send();
        },
        Hupload:function (param) {
            var __DEFAULT = {
                header:"数据导入",
                height:"360px",
                url:"",
                callback:function () {
                },
            };

            $.extend(true,__DEFAULT,param)

            var uploader;
            var upload_modual = '<div class="row"><div class="col-sm-12 col-md-12 col-lg-12"><div class="pull-left"><div id="h-upload-select-file-btn">选择文件</div></div></div><div class="col-sm-12 col-md-12 col-lg-12 uploader-list" style="margin-top: 15px;"><table id="h-selected-files-list" class="table table-responsive table-bordered"><thead><tr><th>文件名</th><th>大小</th></tr></thead><tbody></tbody></table></div></div>'

            $.Hmodal({
                header:__DEFAULT.header,
                body:upload_modual,
                height:__DEFAULT.height,
                submitDesc:"上传",
                cancelDesc:"关闭",
                preprocess:function () {
                    uploader = WebUploader.create({

                        // swf文件路径
                        swf: '/webuploader/dist/Uploader.swf',

                        // 文件接收服务端。
                        server: __DEFAULT.url,

                        // 选择文件的按钮。可选。
                        // 内部根据当前运行是创建，可能是input元素，也可能是flash.
                        pick: '#h-upload-select-file-btn',

                        // 不压缩image, 默认如果是jpeg，文件上传前会压缩一把再上传！
                        resize: false
                    });
                    uploader.on('beforeFileQueued',function () {
                        uploader.reset();
                    });
                    uploader.on( 'fileQueued', function( file ) {

                        var unit = "KB"
                        var size = (file.size/1024).toFixed(0)
                        if (size > 1024) {
                            size = (size/1024).toFixed(2)
                            if (size > 256) {
                                $.Notify({
                                    title:"温馨提示:",
                                    message:"上传文件大于256M,无法上传",
                                    type:"danger",
                                })
                                uploader.reset();
                                return
                            }
                            unit = "MB"
                        }
                        var v = size +" "+ unit

                        var optHtml = "<tr><td>"+file.name+"</td><td>"+v+"</td></tr>"

                        $("#h-selected-files-list").find("tbody").html(optHtml)

                    });

                    uploader.on("uploadError",function (file) {
                        $.Notify({
                            title:"温馨提示:",
                            message:"上传失败,请联系管理员",
                            type:"info",
                        });
                        $("#h-hauth-org-upload-progress").css("width","100%");
                        $("#h-hauth-org-upload-progress").removeClass("progress-bar-info progress-bar-striped").addClass("progress-bar-danger")
                        $("#h-hauth-org-upload-progress span").html("导入失败")
                    });

                },
                callback:function () {
                    uploader.on("uploadSuccess",function (file,response ) {
                        $.Notify({
                            title:"温馨提示:",
                            message:response.data,
                            type:"success",
                        });
                        $("#h-hauth-org-upload-progress span").html("100%")
                        $("#h-hauth-org-upload-progress").css("width","100%");

                        $("#h-hauth-org-upload-progress").removeClass("progress-bar-info progress-bar-striped").addClass("progress-bar-success")

                        if (typeof __DEFAULT.callback == "function"){
                            __DEFAULT.callback()
                        }
                    });
                    uploader.upload();
                    if ($("#h-hauth-org-upload-progress").html() == undefined) {
                        $("#h-selected-files-list").find("tbody").append('<tr><td colspan="2"><span>上传进度:</span><div class="progress" style="margin-top: 5px;"><div id="h-hauth-org-upload-progress" class="progress-bar progress-bar-info progress-bar-striped active" role="progressbar" aria-valuenow="60" aria-valuemin="0" aria-valuemax="100" style="width: 0%;"><span>0%</span></div></div></td></tr>')
                    } else {
                        $.Notify({
                            title:"温馨提示:",
                            message:"已经上传完成",
                            type:"info",
                        });
                        return
                    }
                    uploader.on("uploadProgress",function (file, percentage) {
                        if (percentage==1) {
                            percentage = 0.99
                        }
                        var p = percentage*100+"%"
                        $("#h-hauth-org-upload-progress span").html(p)
                        $("#h-hauth-org-upload-progress").css("width",p);
                    });
                },
            })
        },
        Hmodal:function(param){
            var __DEFAULT = {
                callback : "",
                preprocess: "",
                width:"800px",
                height:"494px ",

                header:"弹框信息",
                headerHeight:"40px",
                headerFontSize:"14px",
                headerFontColor:"",

                body:"",

                footer:"",

                footerBtnStatus:true,

                submitBtn:true,
                submitIcon:"icon-ok",
                submitDesc:"提交",

                cancelBtn:true,
                cancelIcon:"icon-remove",
                cancelDesc:"取消",
            }
            $.extend(true,__DEFAULT,param)

            //初始化弹框主体
            function init(){
                var mframe='<div class="modal-dialog">'+
                    '<div class="modal-content h-modal-content" style="width: '+__DEFAULT.width+'; height: '+__DEFAULT.height+';">'+
                    '<div class="modal-header h-modal-header" style="height: '+__DEFAULT.headerHeight+'; line-height: '+__DEFAULT.headerHeight+'; padding: 0px;">'+
                    '<button type="button" class="close" data-dismiss="modal" aria-hidden="true" style="height: '+__DEFAULT.headerHeight+'; line-height: '+__DEFAULT.headerHeight+'; width: 30px; padding-top: 2px;">×</button>'+
                    '<h4 class="modal-title" style="margin-left: 15px;height: '+__DEFAULT.headerFontSize+';color: '+__DEFAULT.headerFontColor+'; line-height: '+__DEFAULT.headerHeight+';font-weight: 600; font-size: '+__DEFAULT.headerFontSize+'">'+__DEFAULT.header+'</h4>'+
                    '</div>'+
                    '<div class="modal-body" style="width: '+__DEFAULT.width+'; overflow-y: auto">'+__DEFAULT.body+'</div>'+
                    '<div class="modal-footer btn-group-sm">'+
                    '<button type="button" class="btn btn-danger cancel" data-dismiss="modal"><i class="'+__DEFAULT.cancelIcon+'"></i>&nbsp;'+__DEFAULT.cancelDesc+'</button>'+
                    '<button type="button" class="btn btn-primary submit"><i class="'+__DEFAULT.submitIcon+'"></i>&nbsp;'+__DEFAULT.submitDesc+'</button>'+
                    '</div>' +
                    '</div>' +
                    '</div>';
                return mframe;
            }
            //显示弹出框
            function showModal(mframe){
                var hmod=document.createElement("div");
                $(hmod).addClass("modal fade").attr({
                    "tabindex":"-1",
                    "role":"dialog",
                    "aria-labelledby":"myModalLabel",
                    "aria-hidden":"true",
                })
                hmod.innerHTML=mframe;
                document.body.appendChild(hmod);
                $(hmod).modal({backdrop:false});
                $(hmod).modal("show");
                return hmod
            }

            //根据类获取对象实例
            function getObj(mod,className,typeObj){
                if (typeof typeObj == "undefined"){
                    typeObj = "div"
                }
                var obj = {}
                $(mod).find(typeObj).each(function(index,element){
                    if ($(element).hasClass(className)){
                        obj = element
                    }
                })
                return obj
            }

            //调节body高度和宽度
            function modifyBodyHeightAndWidth(mod){
                var headerObj = getObj(mod,"modal-header")
                var contentObj = getObj(mod,"modal-content")
                var bodyObj = getObj(hmode,"modal-body")
                var headHeight = $(headerObj).height()
                var contentHeight = $(contentObj).height()

                $(bodyObj).css("height",contentHeight-headHeight-65)
                $(bodyObj).css("width","-=4")
            }

            //modify location
            function modifyLocation(mod){
                var ww = $(window).width()
                var wh = $(window).height();
                var mw = $(getObj(mod,"modal-content")).width()
                var mh = $(getObj(mod,"modal-content")).height()
                //var modifyY = (wh - 2*mh)/2
                var modifyX = (ww - mw)/2
                $(getObj(mod,"modal-content")).offset({
                    left:modifyX
                })
            }

            function initfooter(mode){
                if (!__DEFAULT.cancelBtn){
                    $(getObj(mode,"cancel","button")).remove();
                }
                if (!__DEFAULT.submitBtn){
                    $(getObj(mode,"submit","button")).remove();
                }
            }

            //
            var mframe =  init()
            var hmode = showModal(mframe)
            modifyBodyHeightAndWidth(hmode)
            modifyLocation(hmode)
            //close modal when click close button in right header
            $(getObj(hmode,"modal-header")).find("button").on("click",function(){
                $(hmode).remove();
            })

            // init footer
            //
            if (__DEFAULT.footerBtnStatus){
                var footer = $(getObj(hmode,"modal-body")).find(".h-modal-footer")
                if ($(footer).find("button").html()==""){
                    console.log("can not found button in modal body content")
                    $(getObj(getObj(hmode,"modal-footer"),"submit","button")).on("click",function(){
                        console.log("no button found, default submit")
                        $(hmode).remove()
                    })
                    $(getObj(getObj(hmode,"modal-footer"),"cancel","button")).on("click",function(){
                        console.log("no button found, default cancel")
                        $(hmode).remove()
                    })
                }else{
                    $(getObj(hmode,"modal-footer")).html($(footer).html())
                    $(footer).remove()
                    if (__DEFAULT.callback == "") {
                        $(getObj(getObj(hmode,"modal-footer"),"submit","button")).on("click",function(){
                            console.log("no callback found, default submit")
                            $(hmode).remove()
                        })
                        $(getObj(getObj(hmode,"modal-footer"),"cancel","button")).on("click",function(){
                            console.log("no callback found, default cancel")
                            $(hmode).remove()
                        })
                    } else if (typeof __DEFAULT.callback == "function"){
                        $(getObj(getObj(hmode,"modal-footer"),"cancel","button")).on("click",function(){
                            console.log("defined callback, cancel");
                            $(hmode).remove()
                        })
                        $(getObj(getObj(hmode,"modal-footer"),"submit","button")).on("click",function(){
                            console.log("defined callback, submit");
                            __DEFAULT.callback(hmode)
                        })
                    }
                }
            }else{
                $(getObj(hmode,"modal-footer")).remove();
                var h = $(getObj(hmode,"modal-body")).height();
                $(getObj(hmode,"modal-body")).height(h+57);
            }

            initfooter(hmode);

            // preprocess function
            if (typeof  __DEFAULT.preprocess == "function"){
                __DEFAULT.preprocess(hmode)
            }


            // 拖动绑定
            var d = "getSelection" in window?function(){
                window.getSelection().removeAllRanges()
            }:function(){
                document.selection.empty()
            };

            var f=0,c=0,e=0,b=0,a=0;
            $(getObj(hmode,"modal-header")).bind("mousemove",function(h){
                if(a==1){
                    f=h.pageX-e;
                    c=h.pageY-b;
                    if(c<=0){
                        c=0
                    }
                    $(this).parent().offset({left:f,top:c})
                }
            }).bind("mousedown",function(h){
                d();
                e=h.pageX-$(this).parent().offset().left;
                b=h.pageY-$(this).parent().offset().top;
                a=1;
                $(getObj(hmode,"modal-header")).css({"cursor":"move"})}
            ).bind("mouseup",function(h){
                $(getObj(hmode,"modal-header")).css({"cursor":"default"});
                a=0;
                e=0;
                b=0
            }).bind("mouseleave",function(h){
                a=0;
                $(getObj(hmode,"modal-header")).css({"cursor":"default"})
            })
        },
        Hconfirm:function(param){
            var __DEFAULT = {
                callback : "",
                preprocess: "",
                width:"420px",
                height:"240px ",

                header:"",
                headerHeight:"30px",
                headerColor :"white",
                headerFontSize:"14px",
                headerFontColor:"#0c0c0c",

                body:"",
                footer:"",
                iconClass:"icon-3x icon-question-sign",
                cancelBtn:true,
                submitBtn:true,
            }
            $.extend(true,__DEFAULT,param)

            //初始化弹框主体
            function init(){
                var mframe='<div class="modal-dialog">'+
                    '<div class="modal-content" style="border: '+__DEFAULT.headerColor+' solid 2px; width: '+__DEFAULT.width+'; height: '+__DEFAULT.height+';">'+
                    '<div class="modal-header h-modal-header" style="border: none !important;background-color: '+__DEFAULT.headerColor+'; height: '+__DEFAULT.headerHeight+'; line-height: '+__DEFAULT.headerHeight+';">'+
                    '<h4 class="modal-title" style="margin-left: 15px;height: '+__DEFAULT.headerFontSize+';color: '+__DEFAULT.headerFontColor+'; line-height: '+__DEFAULT.headerHeight+';font-weight: 600; font-size: '+__DEFAULT.headerFontSize+'; margin-right: 30px;">'+__DEFAULT.header+'</h4>'+
                    '<button type="button" class="close" data-dismiss="modal" aria-hidden="true" style="width: 30px;">×</button>'+
                    '</div>'+
                    '<div class="modal-body" style="width: '+__DEFAULT.width+'; overflow-y: auto;text-align: center"><i style="color: red;" class="'+__DEFAULT.iconClass+'"></i><br/><span style="font-size: 16px;display: block; font-weight: 600; margin-top: 15px;">'+__DEFAULT.body+'</span></div>'+
                    '<div class="modal-footer btn-group-sm" style="text-align: center; border: none;">'+
                    '<button type="button" class="btn btn-danger cancel" style="width: 120px;" data-dismiss="modal"><i class="icon-remove"></i>&nbsp;&nbsp;&nbsp;&nbsp;取消</button>'+
                    '<button type="button" class="btn btn-primary submit" style="width: 120px;margin-left: 50px;"><i class="icon-ok"></i>&nbsp;&nbsp;&nbsp;&nbsp;确定</button>'+
                    '</div>' +
                    '</div>' +
                    '</div>';
                return mframe;
            }

            //显示弹出框
            function showModal(mframe){
                var hmod=document.createElement("div");
                $(hmod).addClass("modal animated rubberBand").attr({
                    "tabindex":"-1",
                    "role":"dialog",
                    "aria-labelledby":"myModalLabel",
                    "aria-hidden":"true",
                })
                hmod.innerHTML=mframe;
                document.body.appendChild(hmod);
                $(hmod).modal({backdrop:false});
                $(hmod).modal("show");
                return hmod
            }

            //根据类获取对象实例
            function getObj(mod,className,typeObj){
                if (typeof typeObj == "undefined"){
                    typeObj = "div"
                }
                var obj = {}
                $(mod).find(typeObj).each(function(index,element){
                    if ($(element).hasClass(className)){
                        obj = element
                    }
                })
                return obj
            }

            //调节body高度和宽度
            function modifyBodyHeightAndWidth(mod){
                var headerObj = getObj(mod,"modal-header")
                var contentObj = getObj(mod,"modal-content")
                var bodyObj = getObj(hmode,"modal-body")
                var headHeight = $(headerObj).height()
                var contentHeight = $(contentObj).height()

                $(bodyObj).css("height",contentHeight-headHeight-65)
                $(bodyObj).css("width","-=4")
            }

            //modify location
            function modifyLocation(mod){
                var ww = $(window).width()
                var wh = document.documentElement.clientHeight;
                var mw = $(getObj(mod,"modal-content")).width()
                var mh = $(getObj(mod,"modal-content")).height()
                var modifyY = (wh - 1.5*mh)/2
                var modifyX = (ww - mw)/2
                if (modifyY < 0){
                    modifyY = 20
                }
                $(getObj(mod,"modal-content")).offset({
                    left:modifyX,
                    top:modifyY
                })
            }

            function initfooter(mode){
                if (!__DEFAULT.cancelBtn){
                    $(getObj(mode,"cancel","button")).remove();
                }
                if (!__DEFAULT.submitBtn){
                    $(getObj(mode,"submit","button")).remove();
                }
            }

            //
            var mframe =  init()
            var hmode = showModal(mframe)
            modifyBodyHeightAndWidth(hmode)
            modifyLocation(hmode);

            //close modal when click close button in right header
            $(getObj(hmode,"modal-header")).find("button").on("click",function(){
                $(hmode).remove();
            })

            // init footer
            var footer = $(getObj(hmode,"modal-body")).find(".h-modal-footer")
            if ($(footer).find("button").html()==""){
                console.log("can not found button in modal body content")
                $(getObj(getObj(hmode,"modal-footer"),"submit","button")).on("click",function(){
                    console.log("no button found, default submit")
                    $(hmode).remove()
                })
                $(getObj(getObj(hmode,"modal-footer"),"cancel","button")).on("click",function(){
                    console.log("no button found, default cancel")
                    $(hmode).remove()
                })
            }else{
                $(getObj(hmode,"modal-footer")).html($(footer).html())
                $(footer).remove()
                if (__DEFAULT.callback == ""){
                    $(getObj(getObj(hmode,"modal-footer"),"submit","button")).on("click",function(){
                        console.log("no callback found, default submit")
                        $(hmode).remove()
                    })
                    $(getObj(getObj(hmode,"modal-footer"),"cancel","button")).on("click",function(){
                        console.log("no callback found, default cancel")
                        $(hmode).remove()
                    })
                }else if (typeof __DEFAULT.callback == "function"){
                    $(getObj(getObj(hmode,"modal-footer"),"cancel","button")).on("click",function(){
                        console.log("defined callback, cancel")
                        $(hmode).remove()
                    })
                    $(getObj(getObj(hmode,"modal-footer"),"submit","button")).on("click",function(){
                        console.log("defined callback, submit")
                        __DEFAULT.callback()
                        $(hmode).remove()
                    })
                }
            }
            initfooter(hmode);
            // preprocess function
            if (typeof  __DEFAULT.preprocess == "function"){
                __DEFAULT.preprocess()
            }
            // 拖动绑定
            var d = "getSelection" in window?function(){
                window.getSelection().removeAllRanges()
            }:function(){
                document.selection.empty()
            };

            var f=0,c=0,e=0,b=0,a=0;
            $(getObj(hmode,"modal-header")).bind("mousemove",function(h){
                if(a==1){
                    f=h.pageX-e;
                    c=h.pageY-b;
                    if(c<=0){
                        c=0
                    }
                    $(this).parent().offset({left:f,top:c})
                }
            }).bind("mousedown",function(h){
                d();
                e=h.pageX-$(this).parent().offset().left;
                b=h.pageY-$(this).parent().offset().top;
                a=1;
                $(getObj(hmode,"modal-header")).css({"cursor":"move"})}
            ).bind("mouseup",function(h){
                $(getObj(hmode,"modal-header")).css({"cursor":"default"});
                a=0;
                e=0;
                b=0
            }).bind("mouseleave",function(h){
                a=0;
                $(getObj(hmode,"modal-header")).css({"cursor":"default"})
            })
        },
        Hworkflow:function (param) {
            var __DEFAULT = {
                callback : "",
                preprocess: "",
                width:"800px",
                height:"494px ",

                header:"弹框信息",
                headerHeight:"40px",
                headerFontSize:"14px",
                headerFontColor:"",

                submitDesc:"提交",
                body:"",

                footer: "",
            };

            $.extend(true,__DEFAULT,param);
            if (__DEFAULT.footer=="" || __DEFAULT.footer == undefined) {
                $(__DEFAULT.body).find(".hzwy23-page").each(function (index, element) {
                    index = index + 1
                    if (index > 1) {
                        __DEFAULT.footer += "&nbsp;&nbsp;--&nbsp;<button class='btn btn-info' style='border-radius: 15px;'>"+index+"</button>"
                    } else {
                        __DEFAULT.footer += "<button class='btn btn-info' style='border-radius: 15px;'>"+index+"</button>"
                    }
                });
            }
            //初始化弹框主体
            function init(){
                var mframe='<div class="modal-dialog">'+
                    '<div class="modal-content h-modal-content" style="width: '+__DEFAULT.width+'; height: '+__DEFAULT.height+';">'+
                    '<div class="modal-header h-modal-header" style="height: '+__DEFAULT.headerHeight+'; line-height: '+__DEFAULT.headerHeight+'; padding: 0px;">'+
                    '<button type="button" class="close" data-dismiss="modal" aria-hidden="true" style="height: '+__DEFAULT.headerHeight+'; line-height: '+__DEFAULT.headerHeight+'; width: 30px; padding-top: 2px;">×</button>'+
                    '<h4 class="modal-title" style="margin-left: 15px;height: '+__DEFAULT.headerFontSize+';color: '+__DEFAULT.headerFontColor+'; line-height: '+__DEFAULT.headerHeight+';font-weight: 600; font-size: '+__DEFAULT.headerFontSize+'">'+__DEFAULT.header+'</h4>'+
                    '</div>'+
                    '<div class="modal-body" style="width: '+__DEFAULT.width+'; overflow-y: auto">'+__DEFAULT.body+'</div>'+
                    '<div class="modal-footer btn-group-sm" style="text-align: center;">'+__DEFAULT.footer+
                    '<button type="button" class="btn btn-primary submit pull-right"><i class="icon-ok"></i>&nbsp;'+__DEFAULT.submitDesc+'</button>'+
                    '</div>' +
                    '</div>' +
                    '</div>';
                return mframe;
            };

            //显示弹出框
            function showModal(mframe){
                var hmod=document.createElement("div");
                $(hmod).addClass("modal fade").attr({
                    "tabindex":"-1",
                    "role":"dialog",
                    "aria-labelledby":"myModalLabel",
                    "aria-hidden":"true",
                })
                hmod.innerHTML=mframe;
                document.body.appendChild(hmod);
                $(hmod).modal({backdrop:false});
                $(hmod).modal("show");
                return hmod
            };

            //根据类获取对象实例
            function getObj(mod,className,typeObj){
                if (typeof typeObj == "undefined"){
                    typeObj = "div"
                }
                var obj = {}
                $(mod).find(typeObj).each(function(index,element){
                    if ($(element).hasClass(className)){
                        obj = element
                    }
                })
                return obj
            };

            //调节body高度和宽度
            function modifyBodyHeightAndWidth(mod){
                var headerObj = getObj(mod,"modal-header")
                var contentObj = getObj(mod,"modal-content")
                var bodyObj = getObj(hmode,"modal-body")
                var headHeight = $(headerObj).height()
                var contentHeight = $(contentObj).height()

                $(bodyObj).css("height",contentHeight-headHeight-65)
                $(bodyObj).css("width","-=4")
            };

            //modify location
            function modifyLocation(mod){
                var ww = $(window).width()
                var wh = $(window).height();
                var mw = $(getObj(mod,"modal-content")).width()
                var mh = $(getObj(mod,"modal-content")).height()
                //var modifyY = (wh - 2*mh)/2
                var modifyX = (ww - mw)/2
                $(getObj(mod,"modal-content")).offset({
                    left:modifyX
                })
            };

            //
            var mframe =  init();
            var hmode = showModal(mframe);
            modifyBodyHeightAndWidth(hmode);
            modifyLocation(hmode);
            //close modal when click close button in right header
            $(getObj(hmode,"modal-header")).find("button").on("click",function(){
                $(hmode).remove();
            });


            // preprocess function
            if (typeof  __DEFAULT.preprocess == "function"){
                __DEFAULT.preprocess(hmode)
            }

            $(getObj(hmode,"modal-footer")).find("button").on("click",function(){
                if (!$(this).hasClass("submit")) {
                    var text = $(this).html();
                    $(getObj(hmode,"modal-body")).find(".hzwy23-page").each(function (index, element) {
                        var index = index + 1;
                        if (index == text) {
                            $(element).show();
                        } else {
                            $(element).hide();
                        }
                    });
                } else {
                    __DEFAULT.callback(hmode );
                }
            });


            // 移除文本选中状态
            var d = "getSelection" in window?function(){
                window.getSelection().removeAllRanges()
            }:function(){
                document.selection.empty()
            };

            // 绑定拖动效果
            var f=0,c=0,e=0,b=0,a=0;
            $(getObj(hmode,"modal-header")).bind("mousemove",function(h){
                if(a==1){
                    f=h.pageX-e;
                    c=h.pageY-b;
                    if(c<=0){
                        c=0
                    }
                    $(this).parent().offset({left:f,top:c})
                }
            }).bind("mousedown",function(h){
                d();
                e=h.pageX-$(this).parent().offset().left;
                b=h.pageY-$(this).parent().offset().top;
                a=1;
                $(getObj(hmode,"modal-header")).css({"cursor":"move"})}
            ).bind("mouseup",function(h){
                $(getObj(hmode,"modal-header")).css({"cursor":"default"});
                a=0;
                e=0;
                b=0
            }).bind("mouseleave",function(h){
                a=0;
                $(getObj(hmode,"modal-header")).css({"cursor":"default"})
            })
        }
    })
}(jQuery));