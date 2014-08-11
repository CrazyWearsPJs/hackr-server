angular.module('app', ['ngRoute', 'ui.bootstrap', 'appControllers'])
  .config(function($routeProvider){
    $routeProvider.when('/', {
    controller: 'splashCtrl',
    templateUrl: './templates/splash.html'
    });
    $routeProvider.otherwise({redirectTo: '/'});
  })
  .constant('_', window._);
