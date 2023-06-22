<!DOCTYPE html>
<html lang="en" xmlns:th="http:www.thymeleaf.org">

<head>
  <meta charset="UTF-8">
  <link rel="stylesheet" href="https://unpkg.com/element-ui/lib/theme-chalk/index.css">
  <script type="text/javascript" src="https://unpkg.com/vtk.js"></script>
  <title>Vue Test</title>
</head>

<body>
  <div id="app">
    <el-container style="border: 1px solid #eee">
      <el-aside width="200px" style="background-color: rgb(238, 241, 246)">
        <el-menu :default-openeds="['1']" :default-active="$route.path" @select="handleSelect">
          <el-menu-item index="">
            <template slot="title"><i class="el-icon-message"></i>首页</template>
          </el-menu-item>
          <el-menu-item index="/menu-1-index">
            <template slot="title"><i class="el-icon-menu"></i>菜单1</template>
          </el-menu-item>
          <el-menu-item index="/menu-2-index">
            <template slot="title"><i class="el-icon-setting"></i>菜单2</template>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-container>
        <router-view></router-view>
      </el-container>
    </el-container>
  </div>
  <script>


    </script>
    <script src="https://unpkg.com/vue@2.5.20/dist/vue.js"></script>
    <script src="https://unpkg.com/vue-router@3.5.0/dist/vue-router.js "></script>
    <script src="https://unpkg.com/element-ui/lib/index.js"></script>
    <script src="https://unpkg.com/axios/dist/axios.min.js"></script>
    <script src="/static/sim/js/util.js"></script>
    <script src="/static/sim/js/route.js"></script>
    <script src="/static/sim/js/main.js"></script>
<script type="module" src="/static/sim/js/aaa.js"></script>

</body>
</html>