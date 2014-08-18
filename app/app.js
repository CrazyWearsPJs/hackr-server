angular.module('app', ['ui.router', 'ui.bootstrap', 'appControllers', 'appDirectives"])
  .config(function($stateProvider, $urlRouterProvider){
    $stateProvider
    .state('splash', {
        url: "",
        controller: 'splashCtrl',
        templateUrl: './templates/splash.html'
    });

    $urlRouterProvider.otherwise("");
  })
  .constant('_', window._);
