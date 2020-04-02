(function(app){
    document.addEventListener('DOMContentLoaded', function() {
        navigator.platformBrowserDynamic
            .platformBrowserDynamic()
            .bootstrapModule(app.AppModule);
    });
})(window.app || (window.app = {}));
