var routes = [{
    path: '/menu-1-index',
    name: 'menu-1-index',
    component: util.loadComponent('./views/menu1.html')
}, {
    path: '/menu-2-index',
    name: 'menu-2-index',
    component: util.loadComponent('./views/menu2.html')
}];

var router = new VueRouter({
    routes: routes
});
