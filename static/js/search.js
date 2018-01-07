$(function () {
    $("#search").keydown(function (event) {
        switch (event.keyCode) {
            case 13:
                if (check()) {                    
                    $(".main_list").empty();
                    $("#search").attr("page", 1);
                    $("#search").attr("isauto", "false");
                    searchstart();
                }
                break;
        }
    });

    if ($("#search").val() != "") {
        searchstart();
    }
    else {
        searchnoresult();
    }

    $("#search_J").click(function () {
        $("#search").focus();
    });
});

function check() {
    if ($("#search").val() == "") {
        alert("请录入要搜索的内容");
        return false;
    }
    return true;
}

function searchstart() {
    var url = location.href;
    if (url.toLowerCase().indexOf("search") > -1) {
        search();
        $("#search_c_J").show();
    } else {
        location.href = "/search/index?keyword=" + $("#search").val();
    }
}
var geting = false;//是否加载信息，默认为false
function search() {    
    if (geting) {
        return;
    }

    var main_list = $(".main_list");
    var search = $("#search");  

    var keyword = search.val();
    if (keyword != "") {

        geting = true;
        var page = 1;
        try {
            page = parseInt(search.attr("page"));
        } catch (e) { }       

        $.post("/search/searchdata?r="+Math.random(), { keyword: keyword, page: page },
         function (res) {
             var data = $.parseJSON(res);
             if (data != "" && data.hits != undefined && data.total > 0 && data.hits.length>0) {
                 $.each(data.hits, function (i) {
                     main_list.css("padding-top", "60px");
                     var item = data.hits[i];

                     $.get("/article/getdetails?id=" + item._id, function (article) {
                         if (article != undefined && article.articleid != undefined) {
                             var arturl = '/article/details?id=' + article.articleid;
                             var blogurl = '/blog/index?username=' + article.username;
                             var html = '<dl class="m_list clearfix">';
                             html += '<dt><a href="' + arturl + '">' + article.title + '</a></dt>';
                             html += '<dd>';
                             html += '<label>';
                             html += '<a href="' + blogurl + '"><img src="' + article.photo + '" alt="@article.UserName"></a>';
                             html += '<a href="' + blogurl + '" class="username">' + article.username + '</a>';
                             html += '</label>';
                             html += '<span><em>' + article.timeago + '</em><em>|</em><a href="@articleUrl">' + article.viewcount + '阅读</a></span>';
                             html += '</dd>';
                             html += '</dl>';

                             main_list.append(html);
                         }
                     });
                 });
             }
             else {
                 if ($("#search").attr("isauto") != "true") {
                     searchnoresult();
                 }
             }

             setTimeout(function () {
                 search.attr("page", page + 1);
                 geting = false;                 
             }, 2000);
         });
    }
}

function searchnoresult() {
    var url = location.href;
    if (url.toLowerCase().indexOf("search") > -1&&url.toLocaleLowerCase().indexOf("details")==-1) {
        var main_list = $(".main_list");
        var html = '<div class="search search_block">';
        html += '</div>';
        html += '<div class="search_no">';
        html += '暂无搜索结果<br/>';
        html += '试试别的关键字<br>';
        html += '<a href="/home/index" style="color:#BB0000">返回首页</a>';
        html += '</div>';
        main_list.html(html);
        $("#search_c_J").show();
        $("#search").focus();
    }
}

function autosearch() {
    $("#search").attr("isauto", "true");
    search();
}
