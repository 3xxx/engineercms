<!-- 这个是显示左侧栏，右边index_user显示用户的cms情况 -->
<!DOCTYPE html>
<head>
  <title>珠三角水资源配置工程设代</title>
  <!-- <script src="https://cdn.jsdelivr.net/npm/vue/dist/vue.js"></script> -->
  <script src="/static/vue.js/vue.js"></script>
  <script src="/static/vue.js/vue-router.js"></script>
  <!-- 引入样式 -->
  <link rel="stylesheet" href="https://unpkg.com/element-ui@2.4.11/lib/theme-chalk/index.css">
  <!-- <link rel="stylesheet" href="/static/vue.js/index.css"> -->
  <!-- <script src="https://unpkg.com/element-ui@2.4.11/lib/index.js"></script> -->
  <script src="/static/vue.js/index.js"></script>
  <!-- 引入组件库 -->
  <!-- <script src="https://unpkg.com/element-ui/lib/index.js"></script> -->
  <!-- <script src="https://unpkg.com/axios/dist/axios.min.js"></script> -->
  <script src="/static/vue.js/axios.js"></script>
  <!-- <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css"> -->

<!-- <link href="/static/index/css/framework.css" rel="stylesheet" type="text/css">
<link href="/static/index/css/owl.theme.css" rel="stylesheet" type="text/css">
<script type="text/javascript" src="/static/index/js/jquery.js"></script>
<script type="text/javascript" src="/static/index/js/framework.plugins.js"></script> -->

  <script type="text/javascript" src="/static/js/jquery-3.3.1.min.js"></script>
  <!-- <script src="https://unpkg.com/vue/dist/vue.js"></script> -->
<style type="text/css">
  .overlay{
    pointer-events:none;
    position:absolute;
    width:100%;
    height:100%;
    background-color:rgba(0,0,0,0.5);
    z-index:9999;
  }
  /*  .snap-content {
    position: absolute;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    width: auto;
    height: auto;
    z-index: 2;
    overflow: auto;
    -webkit-overflow-scrolling: touch;
    -webkit-transform: translate3d(0, 0, 0);
       -moz-transform: translate3d(0, 0, 0);
        -ms-transform: translate3d(0, 0, 0);
         -o-transform: translate3d(0, 0, 0);
            transform: translate3d(0, 0, 0);
  }*/

  /* Header*/
  .header{
    background-image:url(/static/index/ky_img/menu-bg.png);
    background-size:100px 100px;
    height:61px;
  }
  .header-logo{
    margin-top:18px;
    margin-left:25px;
    background-image:url(/static/index/ky_img/logo.png);
    width:80px;
    height:24px;
    background-size:80px 24px;
    float:right;
    margin-right:25px;
  }
  .header-clear{
      height:30px;
  }
  .open-menu{
    color:#FF5722;/*bcbcbc*/
    position:absolute;
    width:55px;
    height:60px;
    left:0px;
    top:0px;
    transition:all 250ms ease;
  }
  .open-menu:hover{
      color:#000;/*FFFFFF*/
      background-color:rgba(255,255,255,0.02);
      transition:all 250ms ease;
  }
  .open-menu i{
      width:55px;
      height:60px;
      line-height:60px;
      text-align:center;
      font-size:14px;
  }
  
  /*Content Heading*/
  /*///////////////*/
  .content-heading{
    margin-bottom:30px; 
  }
  .content-heading h4{
    color:#FFFFFF;
    position:absolute;  
    z-index:9999;
    text-transform:uppercase;
    margin-top:27px;
    padding-left:30px;
      pointer-events:none;
      font-weight:800;
  }
  .content-heading{
    max-height:100px; 
  }
  .content-heading p{
    color:#FFFFFF;
    position:absolute;  
    z-index:999;
    margin-top:48px;
    padding-left:30px;
    opacity:0.5;
      pointer-events:none;
  }
  .content-heading .overlay{
    z-index:99;
    background-color:rgba(233,233,233,0.4);/*rgba(0,0,0,0.8)*/
  }
  .content-heading i{
    font-size:32px;
    position:absolute;
    color:#FFFFFF;
    right:30px; 
    z-index:999;
    margin-top:36px;
      pointer-events:none;
  }
  .content-heading img{
    width:100%;
    display:block;
    position:relative;
    z-index:2;
      transition:all 300ms ease;
  }
  .content-heading img:hover{
      filter: blur(3px);  
      -webkit-filter:blur(3px);
      transition:all 300ms ease;
  }
  @media (min-width:768px){
    .content-heading{
      max-height:140px; 
    }
    
    .content-heading h4{
      font-size:20px; 
      margin-top:45px;
      padding-left:50px;
    }
    
    .content-heading p{
      font-size:13px; 
      margin-top:75px;
      padding-left:50px;
    }
    
    .content-heading i{
      font-size:40px;
      margin-top:53px;
      right:50px; 
    }
  }
  
  .overlay{
    pointer-events:none;
    position:absolute;
    width:100%;
    height:100%;
    background-color:rgba(0,0,0,0.5);
    z-index:9999;
  }
  
  .content{
    clear:both;
    margin-left:30px;
    margin-right:30px;
  }
  @media (min-width:768px){
    .content{
      margin-left:70px;
      margin-right:70px;  
    }
  }
  
  .decoration{ 
    height:1px;
    background-color:rgba(0,0,0,0.1);
    margin-bottom:30px;
    display:block;
    clear:both;
  }
  
  .container{
    margin-bottom:30px;
  }
  .no-bottom{
    margin-bottom:0px;
    padding-bottom:0px;
  }
  
  @media (min-width:760px) { 
    .one-third-responsive{
      width:30%;
      float:left;
      margin-right:5%;  
    }
    
    .one-half-responsive{
      width:46%;
      float:left;
      margin-right:8%;
    }
    
    .sidebar-left-big{
      width:70%;
      float:left;
      margin-right:5% 
    }
    
    .sidebar-right-small{
      width:25%;
      float:right;  
    }
      
    .sidebar-right-big{
      width:70%;
      float:right;  
    }
    
    .sidebar-left-small{
      width:25%;
      float:left;
      margin-right:5%;  
    }
    
    .hide-if-responsive{
      display:none;
    }
  }
  
  /*Thumbnails Columns*/
  .thumb-clear{
    height:40px;
    display:block;
    width:100%;
  }
  .thumb-left{
    line-height:24px;
    display:block;
    padding-bottom:10px;
  }
  .thumb-left a{
    display:block;
    text-align:right;
  }
  .thumb-left img{
    width:100px;
    height:100px;
    border-radius:100px;
    float:left;
    margin-right:20px;
      transition:all 500ms ease;
  }
  .thumb-left img:hover{
      transform:scale(0.9, 0.9);
      transition:all 500ms ease;
  }
  .thumb-left strong{
    color:#1a1a1a;
    display:inline-block;
    padding-bottom:5px;
    font-size:13px;
  }
  .thumb-left em{
    font-style:normal;
  }
  .thumb-right{
    line-height:24px;
    display:block;
    padding-bottom:10px;
  }
  .thumb-right img{
    width:100px;
    height:100px;
    border-radius:100px;
    float:right;
    margin-left:20px;
      transition:all 500ms ease;
  }
  .thumb-right img:hover{
      transform:scale(0.9, 0.9);
      transition:all 500ms ease;
  }
  .thumb-right strong{
    color:#1a1a1a;
    display:inline-block;
    padding-top:5px;
    padding-bottom:5px;
    font-size:13px;
  }
  .thumb-right em{
    font-style:normal;
  }
  @media (min-width:600px){
    .thumb-left img{
      width:140px;
      height:140px;
      border-radius:150px;  
    }
    
    .thumb-left em{
      line-height:28px; 
    }
    
    .thumb-left strong{
      padding-top:10px; 
    }
    
    .thumb-right img{
      width:140px;
      height:140px;
      border-radius:150px;  
    }
    
    .thumb-right em{
      line-height:28px; 
    }
    
    .thumb-right strong{
      padding-top:10px; 
    } 
  }
  
  .last-column{
    margin-right:0%!important;
  }

  .prev-staff, .prev-quote{
    z-index:99999;
    background-image:url(../ky_img/prev1.png);
    background-repeat:no-repeat;
    width:45px;
    height:45px;
    background-size:16px 16px;
      background-position: 13px 15px;
    position:absolute;
    left:0px;
    margin-top:20%; 
    border-radius:45px;
    transition:all 400ms ease;
  }
  /*
  *   Owl Carousel Owl  Theme 
  * v1.24
  */
  .slider-controls{
    margin-bottom:-23px;
  }
  .third-thumb{
    width:29%;
    float:left;
    margin-left:2%;
    margin-right:2%;
  }
  .half-thumb{
    width:48%;
    float:left;
    margin-left:1%;
    margin-right:1%;
  }
  .next-slider{
    background-image:url(../ky_img/next.png);
    background-repeat:no-repeat;
    width:27px;
    height:27px;
    background-size:27px;
    height:27px;
    position:absolute;
    right:20px;
    margin-top:-45%;
  }
  .prev-slider{
    background-image:url(../ky_img/prev.png);
    background-repeat:no-repeat;
    width:27px;
    height:27px;
    background-size:27px;
    height:27px;
    position:absolute;
    left:20px;
    margin-top:-45%;
  }
  .prev-staff, .prev-quote{
    z-index:99999;
    background-image:url(/static/index/ky_img/prev1.png);
    background-repeat:no-repeat;
    width:45px;
    height:45px;
    background-size:16px 16px;
      background-position: 13px 15px;
    position:absolute;
    left:0px;
    margin-top:20%; 
      border-radius:45px;
      transition:all 400ms ease;
  }
  .prev-staff:hover, .prev-quote:hover{
      background-color:rgba(0,0,0,0.1);
      transition:all 400ms ease;
  }
  .next-staff, .next-quote{
    z-index:99999;
    background-image:url(/static/index/ky_img/next1.png);
    background-repeat:no-repeat;
    width:45px;
    height:45px;
    background-size:16px 16px;
      background-position: 16px 15px;
    position:absolute;
    right:0px;
    margin-top:20%;   
      border-radius:45px;
      transition:all 400ms ease;
  }
  .next-staff:hover, .next-quote:hover{
       background-color:rgba(0,0,0,0.1);
      transition:all 400ms ease;   
  }
  @media screen and (orientation:landscape) {
    .next-staff{
      margin-top:24%;
    }
    
    .prev-staff{
      margin-top:24%; 
    }
    
    .next-quote{
      margin-top:4%;  
    }
    
    .prev-quote{
      margin-top:4%;  
    }
    
    .staff-item strong{
      width:180px;
      margin-left:auto;
      margin-right:auto;
    }
  }
  @media only screen 
  and (min-device-width : 768px) 
  and (max-device-width : 1024px) 
  and (orientation : landscape) { 
    .next-staff{
      margin-top:12.1%;
    }
    .prev-staff{
      margin-top:12.1%; 
    }
    
    .next-quote{
      margin-top:4%;  
    }
    
    .prev-quote{
      margin-top:4%;  
    }
    
    .staff-item strong{
      width:180px;
      margin-left:auto;
      margin-right:auto;
    }
  }
  @media only screen 
  and (min-device-width : 768px) 
  and (max-device-width : 1024px) 
  and (orientation : portrait) { 
    .next-staff{
      margin-top:17.3%;
    }
    .prev-staff{
      margin-top:17.3%; 
    }
    
    .next-quote{
      margin-top:8%;  
    }
    
    .prev-quote{
      margin-top:8%;  
    }
    
    .staff-item strong{
      width:180px;
      margin-left:auto;
      margin-right:auto;
    }
  }
  
  /*////////////////////*/
  /*Quote & Staff Slider*/
  /*////////////////////*/
  .staff-item img{
    border-radius:300px;  
    max-width:150px;
    margin-left:auto;
    margin-right:auto;
    transition:all 500ms ease;
    margin-top:10px;
    text-align:center;
    /*width: 400px;*/
    /*height: 300px;*/
    /*border: 1px dashed #ccc;*/
    display: table-cell; /*//主要是这个属性*/
    /*vertical-align: middle;*/
    /*text-align: center;*/
  }
  .staff-item img:hover{
      transform:scale(1.1, 1.1);
      transition:all 500ms ease;
  }
  .staff-item h4{
    text-align:center;
    margin-top:20px;  
    font-size:16px;
  }
  .staff-item em{
    text-align:center;
    font-size:12px;
    color:#e34e47;
    display:block;
    margin-top:-10px;
    margin-bottom:10px;
    font-style:normal;
  }
  .staff-item strong{
    display:block;  
    font-weight:400;
    text-align:center;
    padding-left:20px;
    padding-right:20px;
    line-height:28px;
  }
  /* 8. Quote Slider */
  .quote-slider h4{
    font-family:'Source Sans Pro', sans-serif;
    font-size:18px;
    font-weight:300;  
    width:87%;
    text-align:center;
    margin-left:auto;
    margin-right:auto;
    line-height:36px;
  }
  .quote-slider a{
    text-align:center;
    margin-top:20px;  
  }
</style>
<style lang="scss">
  .el-row {
    margin-bottom: 20px;
    &:last-child {
      margin-bottom: 0;
    }
  }
  .el-col {
    border-radius: 4px;
  }
  .bg-purple-dark {
    background: #99a9bf;
  }
  .bg-purple {
    background: #d3dce6;
  }
  .bg-purple-light {
    background: #e5e9f2;
  }
  .grid-content {
    border-radius: 4px;
    min-height: 36px;
    line-height:36px;
    text-indent : 25px;
  }

  .grid-content .cusheader {
    min-height: 60px;
    line-height:60px;
    text-indent : 5px;
  }

  .row-bg {
    padding: 10px 0;
    background-color: #f9fafc;
  }

  /*走马灯高度*/
  .el-carousel__container {
    height: 100% !important;
  }
 
  /*img {
    display: inline-block;
    height: auto;
    max-width: 100%;
  }*/

  /*a:link,a:visited{*/
    /*text-decoration:none;*/  /*超链接无下划线*/
  /*color:#000;
  }*/
  /*div span a:link,a:visited{*/
    /*text-decoration:none;*/  /*超链接无下划线*/
    /*color:#FFFFFF;*/
  /*}*/
  /*  .el-carousel__item h3 {
    color: #475669;
    font-size: 14px;
    opacity: 0.75;
    line-height: 200px;
    margin: 0;
  }
  
  .el-carousel__item:nth-child(2n) {
    background-color: #99a9bf;
  }
  
  .el-carousel__item:nth-child(2n+1) {
    background-color: #d3dce6;
  }

  .bannerImg{
    width: 100%;
    height: inherit;
    min-height: 360px;
    min-width: 1400px;
  }*/

  /*  .el-header, .el-footer {
    background-color: #B3C0D1;
    color: #333;
    margin: 0;
    text-align: center;
    line-height: 60px;
  }*/
  
  /*  .el-aside {
    background-color: #D3DCE6;
    color: #333;
    text-align: center;
    line-height: 200px;
  }
  
  .el-main {
    background-color: #E9EEF3;
    color: #333;
    text-align: center;
    line-height: 160px;
  }
  
  body > .el-container {
    margin-bottom: 40px;
  }*/
  
  /*  .el-container:nth-child(5) .el-aside,
  .el-container:nth-child(6) .el-aside {
    line-height: 260px;
  }
  
  .el-container:nth-child(7) .el-aside {
    line-height: 320px;
  }*/
</style>
</head>

<body>
<div id="app">
  <el-container>
  <!-- <el-aside> -->
  <!-- <el-radio-group v-model="isCollapse" style="yg-radio-group"> <i class="el-icon-d-arrow-right"></i><div @click="changeWidthR" class="ab-r"></div><i class="el-icon-d-arrow-left"></i><div @click="changeWidthL" class="ab-l"></div>-->

  <!-- <el-radio-group v-model="isCollapse" style="margin-bottom: 20px;">
  <el-radio-button :label="false">展开</el-radio-button>
  <el-radio-button :label="true">收起</el-radio-button>
  </el-radio-group> -->
<el-menu default-active="activeIndex" class="el-menu-vertical-demo" @open="handleOpen" @close="handleClose" :collapse="isCollapse" router>
  <el-menu-item index="index">
    <i class="el-icon-menu"></i>
    <span slot="title">首页</span>
  </el-menu-item>
  <el-submenu index="1">
    <template slot="title">
      <i class="el-icon-location"></i>
      <span slot="title" @click="jump('project')">项目</span>
    </template>
    <!-- <el-menu-item-group> -->
      <el-menu-item index="bar" @click="jump('design')">设计</el-menu-item>
      <el-menu-item index="1-2">建设</el-menu-item>
      <el-menu-item index="1-3">监理</el-menu-item>
      <!-- <el-menu-item index="1-4">施工</el-menu-item> -->
    <!-- </el-menu-item-group> -->
  </el-submenu>
  <el-submenu index="2">
    <template slot="title">
      <i class="el-icon-printer"></i>
      <span slot="title">标段</span>
    </template>
    <!-- <el-menu-item-group title="A标"> -->
      <el-submenu index="2-1">
        <span slot="title">A标</span>
        <el-menu-item index="2-1-1">A1标</el-menu-item>
        <el-menu-item index="2-1-2">A2标</el-menu-item>
        <el-menu-item index="2-1-3">A3标</el-menu-item>
        <el-menu-item index="2-1-4">A4标</el-menu-item>
        <el-menu-item index="2-1-5">A5标</el-menu-item>
        <el-menu-item index="2-1-6">A6标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-2">
        <span slot="title">B标</span>
        <el-menu-item index="2-2-1">B1标</el-menu-item>
        <el-menu-item index="2-2-2">B2标</el-menu-item>
        <el-menu-item index="2-2-3">B3标</el-menu-item>
        <el-menu-item index="2-2-4">B4标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-3">
        <span slot="title">C标</span>
        <el-menu-item index="2-3-1">C1标</el-menu-item>
        <el-menu-item index="2-3-2">C2标</el-menu-item>
        <el-menu-item index="2-3-3">C3标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group> -->
    <!-- <el-menu-item-group title="B标"> -->
      <el-submenu index="2-4">
        <span slot="title">D标</span>
        <el-menu-item index="2-4-1">D1标</el-menu-item>
        <el-menu-item index="2-4-2">D2标</el-menu-item>
      </el-submenu>
    <!-- </el-menu-item-group>     -->
  </el-submenu>

  <el-menu-item index="3" @click="jump('onlyoffice')">
    <i class="el-icon-upload"></i>
    <span slot="title">onlyoffice</span>
  </el-menu-item>
  <el-menu-item index="4" disabled>
    <i class="el-icon-document"></i>
    <span slot="title">视频</span>
  </el-menu-item>
  <el-menu-item index="5">
    <i class="el-icon-setting"></i>
    <span slot="title">日记</span>
  </el-menu-item>
</el-menu>
    <!-- </el-aside> v-model语法糖，相当于v-bind:value="isCollapse" v-on:input="isCollapse = $event.target.value-即label的true或false"-->
<el-container>
  <el-header class="grid-content bg-purple-light">
    <el-row>
      <el-col :span="12" style="text-align: left;" class="cusheader">
        <!-- <el-switch v-model="isCollapse" active-color="#13ce66" inactive-color="#ff4949">
        </el-switch> -->
        <i v-model="isCollapse" @click="changeCollapse(isCollapse)" :class="{'el-icon-d-arrow-right':isCollapse,'el-icon-d-arrow-left':!isCollapse}"></i>
      </el-col>
      <el-col :span="12" style="text-align: right;" class="cusheader">
        <el-dropdown>
          <i class="el-icon-setting" style="margin-right: 5px"></i>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item @click.native="dialogFormVisible = true">登陆</el-dropdown-item>
            <el-dropdown-item>大事记</el-dropdown-item>
            <el-dropdown-item>查阅</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
        <span>设代处</span>
      </el-col>
    </el-row>
  </el-header>

  <el-main>
  <!-- <div class="header">
    <a href="index.html" class="header-logo">
      <el-dropdown>
        <i class="el-icon-setting" style="margin-right: 15px"></i>
        <el-dropdown-menu slot="dropdown">
          <el-dropdown-item @click.native="dialogFormVisible = true">登陆</el-dropdown-item>
          <el-dropdown-item>大事记</el-dropdown-item>
          <el-dropdown-item>查阅</el-dropdown-item>
        </el-dropdown-menu>
      </el-dropdown>
    </a>
    <a href="#" class="open-menu">
      <i class="fa fa-navicon"></i> transform:rotate(180deg);
      <i v-model="isCollapse" @click="changeCollapse(isCollapse)" :class="{'el-icon-d-arrow-right':isCollapse,'el-icon-d-arrow-left':!isCollapse}"></i>  
    </a>
  </div>-->


 <!-- </div> -->
  <!-- <el-radio-group v-model="isCollapse" style="margin-bottom: 20px;">
    <el-radio-button :label="false">展开</el-radio-button>
    <el-radio-button :label="true">收起</el-radio-button>
  </el-radio-group> -->
    <el-row  style="height: 300px">
    <el-col :span="24">
  <template>
    <el-carousel :interval="4000" style="height: 300px">
    <!-- <el-carousel indicator-position="outside" :height="bannerHeight + 'px'"> -->
      <!-- <el-carousel-item v-for="(item,index) in superurl" :key="item.id"> -->
      <el-carousel-item v-for="(item,index) in posts" :key="item.id">
        <!-- <h3>[[ item ]]</h3>  v-if="index == item.id" class="bannerImg"-->
        <img :src=item.imgUrl style="height: 100%">
      </el-carousel-item>
    </el-carousel>
  </template>
  </el-col>
  </el-row>

    <!-- <template>
      <el-carousel :interval="4000" type="card" height="200px">
        <el-carousel-item v-for="item in card_img" :key="item">
          <img :src=item.url >
        </el-carousel-item>
      </el-carousel>
    </template> -->
  <!-- <el-row type="flex" class="row-bg" justify="space-around" style="height:200px"> -->
    <!-- <i class="el-icon-edit">设计</i> -->
  <!-- </el-row> -->
  
  <!-- <el-row type="flex" class="row-bg" justify="space-around">
    <el-col :span="11">
      <div class="grid-content bg-purple">      
        <h1>文章组</h1>
        <div v-for="(post,index) in posts" :key="post.id" >
          <router-link :to="'/blog/' + post.id">
            <figure>
              <img v-if="post.imgUrl" :src="post.imgUrl" alt="" style="width: 100%">
              <img v-else src="http://via.placeholder.com/250x250" alt="">
            </figure>
            <h2>[[post.title]]</h2>
              <p></p>
              <div v-html=post.html style="font-size:32px;"></div>
          </router-link>
        </div>
      </div>
    </el-col>
  </el-row> -->
  <!-- <div class="header">
    <a href="index.html" class="header-logo"></a>
    <a href="#" class="open-menu">
      <i class="fa fa-navicon"></i>
      <i class="el-icon-d-arrow-right"></i>  
    </a>
  </div> -->
  <!-- Content Heding -->
  <!-- <div class="content-heading">
    <h4>Services</h4>
      <p>The things we love to do, for you!</p>
      <i class="fa fa-cog"></i>
      <div class="overlay"></div>
      <img src="/static/images/5w.jpg" alt="img">
  </div> -->
  <el-row>
    <el-col :span="24" style="text-align: left;">
      <div class="grid-content bg-purple-light">最新消息</div>
    </el-col>
  </el-row>

  <!-- Page Content-->
  <div class="content">
    <div class="decoration"></div>
    <div class="container no-bottom">
      <!-- Page Columns-->
      <cus v-for="(post,index) in posts" :key="post.id">
        <div v-if="!even1(index)" class="decoration hide-if-responsive"></div>
        <div v-if="even(index)" class="one-half-responsive">
          <p class="thumb-left no-bottom">
            <img v-if="post.imgUrl" :src="post.imgUrl" alt="img">
            <!-- <img v-if="even(index)" :src="post.imgUrl" alt="img"> -->
            <strong>[[post.title]]</strong>
            <em v-html=getSimpleText(post.html) style="font-size:16px;"><br></em> 
          </p>
          <div class="thumb-clear"></div>
        </div>

        <!-- <div v-if="even(index)" class="decoration hide-if-responsive"></div> -->

        <div v-else-if="!even(index)" class="one-half-responsive last-column">
            <p class="thumb-right no-bottom">
              <img v-if="post.imgUrl" :src="post.imgUrl" alt="img">
              <!-- <img v-if="!even(index)" :src="post.imgUrl" alt="img"> -->
              <strong>[[post.title]]</strong>
              <em v-html=getSimpleText(post.html) style="font-size:16px;"><br></em></p>
              <div class="thumb-clear"></div>
        </div> 
      </cus>
    </div>

    <div class="decoration"></div>

  </div>



    
      <!-- Content Heding -->
    <!-- <div class="content-heading">
      <h4>Services</h4>
        <p>The things we love to do, for you!</p>
        <i class="fa fa-cog"></i>
        <div class="overlay"></div>
        <img src="/static/index/images/5w.jpg" alt="img">
    </div> -->

        <!-- Content Heading -->
        <!-- <div class="content-heading">
          <h4>The Staff</h4>
            <p>Meet our awesome staff!</p>
            <i class="fa fa-user"></i>
            <div class="overlay"></div>
            <img src="/static/images/6w.jpg" alt="img">
        </div>  -->
        
        <!-- Page Content-->
        <!-- <div class="content"> -->
          <!-- <div class="decoration"></div> -->
            <!-- Staff Slider-->
            <!-- <div class="container">
                <a href="#" class="next-staff"></a>
                <a href="#" class="prev-staff"></a>
                <div class="staff-slider" data-snap-ignore="true">
                    <div>
                        <div class="staff-item">
                            <img src="/static/images/1s.jpg" alt="img">
                            <h4>John Doe</h4>
                            <em>Web Designer</em>
                            <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                            <a href="#" class="button button-red center-button">Call</a>
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="/static/images/2s.jpg" alt="img">
                            <h4>Jane Hidden</h4>
                            <em>Front End Developer</em>
                            <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                            <a href="#" class="button button-green center-button">Text</a>
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="/static/images/3s.jpg" alt="img">
                            <h4>Johanna Pear</h4>
                            <em>Business Manager</em>
                            <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                            <a href="#" class="button button-blue center-button">Mail</a>
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="/static/images/4s.jpg" alt="img">
                            <h4>Mike Grape</h4>
                            <em>Web Designer</em>
                            <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                            <a href="#" class="button button-dark center-button">Read More</a>
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="/static/images/5s.jpg" alt="img">
                            <h4>Victor Leaf</h4>
                            <em>Front End Developer</em>
                            <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                            <a href="#" class="button button-orange center-button">Facebook</a>
                        </div>
                    </div>
                    <div>
                        <div class="staff-item">
                            <img src="/static/images/6s.jpg" alt="img">
                            <h4>Snow White</h4>
                            <em>Business Manager</em>
                            <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                            <a href="#" class="button button-yellow center-button">Twitter</a>
                        </div>
                    </div>
                </div>
            </div> -->  
            <!-- <div class="decoration"></div>
        </div> -->



  <el-row>
    <el-col :span="24" style="text-align: left;">
      <div class="grid-content bg-purple-light">人员简介</div>
    </el-col>
  </el-row>

  <template>
    <el-carousel :interval="5000" arrow="always" style="height: 500px">
      <el-carousel-item v-for="(post,index) in posts" :key="post.id">
        <el-row>
          <el-col :span="12">
            <div class="grid-content bg-purple">
              <div class="content">
                <div class="container">
                  <div class="staff-item">
                    <img src="/static/images/1s.jpg" alt="img">
                    <h4>[[post.id]]John Doe</h4>
                    <em>Web Designer</em>
                    <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                    <!-- <a href="#" class="button button-red center-button"></a> -->
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Call</el-button>
                  </div>
                </div>
              </div>
            </div>
          </el-col>
          <el-col :span="12">
            <div class="grid-content bg-purple-light">
              <div class="content">
                <div class="container"> 
                  <div class="staff-item">
                    <img src="/static/images/3s.jpg" alt="img">
                    <h4>Johanna Pear</h4>
                    <em>Business Manager</em>
                    <strong style="font-size:16px;">We love quotes, and sometimes it's annoying to see tons of them that you need to scroll to!</strong>
                    <!-- <a href="#" class="button button-blue center-button">Mail</a> -->
                    <el-button type="primary" style="display:block;margin:0 auto;width:200px;">Mail</el-button>
                  </div>
                </div>
              </div>
            </div>
          </el-col>
        </el-row>
      </el-carousel-item>
    </el-carousel>
  </template>
<div class="decoration"></div>
  </el-main>


  </el-container>
</el-container>
  <el-footer style="text-align: center;">Copyright © 2016-2020 EngineerCMS</el-footer>
    <!-- <el-carousel indicator-position="outside" :height="bannerHeight + 'px'">
     <el-carousel-item v-for="(item,index) in BannerImg">
       <img src="../../assets/images/banner1.jpg" v-if="index == 0" class="bannerImg" />
       <img src="../../assets/images/banner2.jpg" v-if="index == 1" class="bannerImg" />
       <img src="../../assets/images/banner3.jpg" v-if="index == 2" class="bannerImg" />
     </el-carousel-item>
   </el-carousel> -->

<!-- <el-button round @click="dialogFormVisible = true">登录</el-button> -->
<el-dialog title="系统登录" :visible.sync="dialogFormVisible" center>
  <!-- 插入测试 -->
  <el-form :model="ruleForm2" status-icon :rules="rules2" ref="ruleForm2" label-width="100px" class="demo-ruleForm">
    <el-form-item label="账号" prop="num">
      <el-input v-model.number="ruleForm2.num" auto-complete="off" placeholder="账号"></el-input>
    </el-form-item>
    <el-form-item label="密码" prop="pass">
      <el-input type="password" v-model="ruleForm2.pass" auto-complete="off" placeholder="密码"></el-input>
    </el-form-item>
    <el-checkbox v-model="checked" checked class="remember">记住密码</el-checkbox>
    <el-form-item label="记住密码" prop="delivery">
      <el-switch v-model="ruleForm2.delivery"></el-switch>
    </el-form-item> 
    <span><a>忘记密码？</a></span>
  </el-form>
   <!-- 插入测试 -->
  <div slot="footer" class="dialog-footer">
    <el-button @click="dialogFormVisible = false; resetForm('ruleForm2')">取 消</el-button>
    <el-button type="primary" @click="submitForm('ruleForm2')">登 录</el-button>
  </div>
</el-dialog>

  <transition>
    <router-view>afadf</router-view>
  </transition>

</div>

<!-- [[ info ]]
  <input type="text" :name="users.name" v-model="users.name" placeholder="姓名">
        <input type="text" :age="users.age" v-model="users.age" placeholder="年龄">
        <button @click="sendGetByObj">发送get请求</button>
</div> -->

<script type="text/javascript">
  // var Main = {
  //   delimiters: ['[[', ']]'],
  //   data() {
  //     return {
  //       isCollapse: true,
  //       // item:"哈哈",
  //       // calleft:0
  //     };
  //   },
  //   mounted:function () {
  //     // this.sendGetByObj();
  //     this.setSize();
  //     const that = this;
  //     window.addEventListener('resize', function() {
  //       that.screenWidth = $(window).width();
  //       that.setSize();
  //     }, false);
  //   },
    
  //   methods: {
  //     handleOpen(key, keyPath) {
  //       console.log(key, keyPath);
  //     },
  //     handleClose(key, keyPath) {
  //       console.log(key, keyPath);
  //     }
  //   }
  // }
// var Ctor = Vue.extend(Main)
// new Ctor().$mount('#app')
// new Vue().$mount('#app')
  const Foo = { template: '<div>foo</div>' };
  const Bar = resolve => require(['../views/login.tpl'],resolve);
  const Main = { template: '<div></div>' };
  const routes = [
    { path: '/foo', component: Foo },
    { path: '/bar', component: Bar },
    { path: '/', component: Main }
  ];
  const router = new VueRouter({
    mode: 'history',
    routes // (缩写) 相当于 routes: routes
  });
  var app = new Vue({
      router,
      delimiters: ['[[', ']]'],
      el:'#app',
      data:{
        users:{
            name:'',
            age:''
        },
        message: 'Hello Vue!',
        info: null,
        superurl: [
          {
              id:'1',
              url: '',
              img: '/static/img/100126522.jpg',
          },
          {
              id:'2',
              url: '',
              img: '/static/img/100165705.jpg',
          },
          {
              id:'3',
              url: '',
              img: '/static/img/100268041.jpg',
          },
          {
              id:'4',
              url: '',
              img: '/static/img/16pic_1227574_b.jpg',
          },
          {
              id:'5',
              url: '',
              img: '/static/img/254247.jpg',
          },
          {
              id:'6',
              url: '',
              img: '/static/img/500171023.jpg',
          },
        ],
        posts:[],
      },
      data() {
        var checkNum = (rule, value, callback) => {
          if (!value) {
            return callback(new Error('账号不能为空'));
          }
          setTimeout(() => {
            if (!Number.isInteger(value)) {
              callback(new Error('请输入数字值'));
            } else {
              var myreg=/^[1][3,4,5,7,8][0-9]{9}$/;
              if (!myreg.test(value) ) {
                callback(new Error('请输入正确的手机号码'));
              } else {
                callback();
              }
  
            }
          }, 1000);
        };
        var validatePass = (rule, value, callback) => {
          if (value === '') {
            callback(new Error('请输入密码'));
          } else {
  
            callback();
          }
        };

        return {
          loginPower:false,
          checked: true,
          /*插入form方法*/
          /*设定规则指向*/
          ruleForm2: {
            pass: '',
            num: '',
             delivery: false,
          },
          rules2: {
            pass: [
              { validator: validatePass, trigger: 'blur' }
            ],
  
            num: [
              { validator: checkNum, trigger: 'blur' }
            ]
          },
          /*插入form方法*/
          dialogTableVisible: false,
          dialogFormVisible: false,
          form: {
            name: '',
            type: [],
            resource: '',
            desc: ''
          },
          formLabelWidth: '120px',

          isCollapse: true,
          bannerHeight:200,
          // item:"哈哈",
          superurl: [
              {
                  id:'1',
                  url: '',
                  img: '/static/img/100126522.jpg',
              },
              {
                  id:'2',
                  url: '',
                  img: '/static/img/100165705.jpg',
              },
              {
                  id:'3',
                  url: '',
                  img: '/static/img/100268041.jpg',
              },
              {
                  id:'4',
                  url: '',
                  img: '/static/img/16pic_1227574_b.jpg',
              },
              {
                  id:'5',
                  url: '',
                  img: '/static/img/254247.jpg',
              },
              {
                  id:'6',
                  url: '',
                  img: '/static/img/500171023.jpg',
              },
          ],
          clientHeight:'',
          // calleft:0
          posts:[],
          numbers:0,

          dialogFormVisible: false,
          form: {
            name: '',
            region: '',
            date1: '',
            date2: '',
            delivery: false,
            type: [],
            resource: '',
            desc: ''
          },
          formLabelWidth: '120px'
        };
      },
      //注意这里我是将它的改为这样的，效果是一样的，但使用这个可以在页面任何部位设置跳转时头部导航条部分样式会跟这变化即为选中，而不是不变的例如：你设置跳转到产品页导航条的样式也会跟着选中产品项
      computed:{
        activeIndex(){
          return this.$route.path.replace('/','')
        }
      },

      mounted:function () {
        this.sendGetByObj();
        this.setSize();
        const that = this;
        window.addEventListener('resize', function() {
          that.screenWidth = $(window).width();
          that.setSize();
        }, false);

        // 获取浏览器可视区域高度
        this.clientHeight =   `${document.documentElement.clientHeight}`          //document.body.clientWidth;
        console.log(this.clientHeight);
        window.onresize = function temp() {
          this.clientHeight = `${document.documentElement.clientHeight}`;
        };
      },
      watch: {
        initData () {
          let H = document.querySelector('.boxShadow')
          H.style.height = ''
          setTimeout(() => {
            console.log(H.offsetHeight)
            if (H.offsetHeight < window.innerHeight) {
              document.body.style.height = window.innerHeight + 'px'
              H.style.height = window.innerHeight - 30 + 'px'
            } else {
              document.body.style.height = H.offsetHeight + 'px'
              H.style.height = ''
            }
          }, 300)
        },

        // 如果 `clientHeight` 发生改变，这个函数就会运行
        // clientHeight: function () {
        //   this.changeFixed(this.clientHeight)
        // },
        // 如果 `clientWidth` 发生改变，这个函数就会运行
        // clientWidth: function () {
        //   this.changeFixed2(this.clientWidth)
        // }
      },
      proxyTable: {
        '/api':{//此处并非一定和url一致。
          target:'https://zsj.itdos.com',
          changeOrigin:true,//允许跨域
          pathRewrite:{
            '^/api': ''
          }
        }
      },      
      methods:{
        submitForm(formName) {
          this.$refs[formName].validate((valid) => {
            if (valid) {
            //提交成功做的动作
            dialogFormVisible = false;
              /* alert('submit!') ; */
              this.$message({
                type: 'success',
                message: '提交成功' 
              });
            } else {
              console.log('error submit!!');
              return false;
            }
          }); 
        },

        resetForm(formName) {
          this.$refs[formName].resetFields();
        },

        setSize: function () {
          this.bannerHeight = 740 / 2560 * this.screenWidth
          if(this.bannerHeight > 740) this.bannerHeight = 740
          if(this.bannerHeight < 360) this.bannerHeight = 360
        },

        // changeFixed(clientHeight){//动态修改样式
        //   console.log(clientHeight);
        //   this.$refs.homePage.style.height = clientHeight+'px';
        // },

        // changeFixed2(clientWidth){//动态修改样式
        //   console.log(clientWidth);
        //   this.$refs.homePage.style.width = clientWidth+'px';
        // },
        changeWidthL(key, keyPath) {
          console.log(key, keyPath);
        },

        handleOpen(key, keyPath) {
          console.log(key, keyPath);
        },
        handleClose(key, keyPath) {
          console.log(key, keyPath);
        },
        sendGetByStr(){
          axios.get(`api/v1/wx/getlistarticles?page=1`)//1.get通过直接发字符串拼接
            // .then(function (response) {
            //   console.log(response);
            //   console.log(response.data.info);
            // }
              .then(response => (this.info = response.data.info)
              )
            .catch(function (error) {
              console.log(error);
            });
        },
        sendGetByObj(){
          axios.get(`/v1/wx/getlistarticles`,{//2.get通过params选项
            params:{
                page:1
            }
          })
          // axios({
          //  headers: {
          //   'X-Requested-With': 'XMLHttpRequest',
          //   'Content-Type': 'application/json; charset=UTF-8',
          //   'Access-Control-Allow-Origin': '*'
          //   },//设置跨域请求头
          //   method: "GET",//请求方式
          //   url: "https://zsj.itdos.com/v1/wx/getlistarticles",//请求地址
          //   params:{
          //       page:1
          //   }
            // data: {
            //   "menu_id":1,
            //   "thirdapp_id":1//请求参数
            // }
          // })
            .then(response => (this.posts = response.data.articles))
            .catch(function (error) {
              console.log(error);
          });
        },
        //html剔除富文本标签，留下纯文本
        getSimpleText(html){
          var re1 = new RegExp("<.+?>","g");//匹配html标签的正则表达式，"g"是搜索匹配多个符合的内容
          var msg = html.replace(re1,'');//执行替换成空字符
          return '<br>'+msg.substring(0,20);
        },
        even: function (numbers) {
          // console.log(numbers);
          // console.log(numbers % 2 === 0);
          // return numbers.filter(function (number) {
            return numbers % 2 === 0
          // })
        },
        even1: function (numbers) {
          // console.log(numbers);
          // console.log(numbers % 2 === 0);
          // return numbers.filter(function (number) {
            return numbers === 0
          // })
        },
        changeCollapse:function(isCollapse){
          if (isCollapse==true) {
            this.isCollapse=false
          }else{
            this.isCollapse=true
          }
        },
        skip(a){
          this.$router.push(a)
        },
        jump(select){
          console.log(select);
          let routeUrl = this.$router.resolve({
            path: "/onlyoffice",
            // query: {id:96}
          });
          switch (select) {
            case ("onlyoffice"):
              routeUrl = this.$router.resolve({
                path: "/onlyoffice",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            case ("project"):
              routeUrl = this.$router.resolve({
                path: "/project",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            case ("design"):
              routeUrl = this.$router.resolve({
                path: "/design",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
            default:
              routeUrl = this.$router.resolve({
                path: "/index",
                // query: {id:96}
              });
              window.open(routeUrl .href);
              break;
          }
          //this.$router.push({path: '/cart?goodsId=12'})
          //this.$router.go(-2)
          //后退两步
          // let routeUrl = this.$router.resolve({
          //   path: "/onlyoffice",
          //   query: {id:96}
          // });
          // window.open(routeUrl .href, '_blank');
          // window.open(routeUrl .href);
        }

      }
  });

  //用delimiters
  // Vue.config.delimiters = ['[[', ']]']
  // const axios = require('axios');
  //   var app = new Vue({
  //   delimiters: ['[[', ']]'],
  // el: '#app',
  // data: {
  //   message: 'Hello Vue!'
  // },
  // data () {
  //   return {
  //     info: null
  //   }
  // },
  // mounted () {
  //   axios
  //     .get('/v1/wx/getlistarticles?page=1')
  //     .then(response => (this.info = response))
  // }
// })

</script>
</body>

<!-- 注：对于路径的写法： ./ 当前目录 ../ 父级目录 / 根目录
<el-carousel trigger="click" :height="bannerH +'px'">
   <el-carousel-item v-for="(item,index) in bannerImgLst" :key="index">
       <img :src="'https://mirror198829.github.io/static/cloud/'+item" class="bannerImg"/>
   </el-carousel-item>
</el-carous>

export default {
  name: 'homePage',
  data () {
    return {
      bannerH:200,
      bannerImgLst:['navBg1.png','navBg2.png','navBg3.jpg']
    }
  },
  methods:{
    setBannerH(){
      this.bannerH = document.body.clientWidth / 4
    }
  },
  mounted(){
    this.setBannerH()
    window.addEventListener('resize', () => {
      this.setBannerH()
    }, false)
  },
  created(){}
} -->