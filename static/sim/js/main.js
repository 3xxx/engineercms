var app = new Vue({
    el: '#app',
    router : router,
    data: function () {
        return {
            activeIndex:1,
        }
    },
    methods: {
        handleSelect : function(key, path) {
            console.log(key, path)
            this.$router.push(key)
        }
    }
});