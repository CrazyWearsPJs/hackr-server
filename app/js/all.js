angular.module('appControllers', ['loginController'])
  .controller('splashCtrl', ["$scope", "$modal", function($scope, $modal){
    $scope.open = function(size){
      modalInstance = $modal.open({
        templateUrl: 'templates/login.html',
        controller: 'loginCtrl',
        size: size
      });
    };
}]);

angular.module('loginController', [])
.controller('loginCtrl', ["$scope", "$modalInstance", function($scope, $modalInstance){

  $scope.ok = function() {
    $modalInstance.close();
  };

  $scope.cancel = function(){
    $modalInstance.dismiss('cancel');
  };
}]);

angular.module('app', ['ngRoute', 'ui.bootstrap', 'appControllers'])
  .config(["$routeProvider", function($routeProvider){
    $routeProvider.when('/', {
    controller: 'splashCtrl',
    templateUrl: './templates/splash.html'
    });
    $routeProvider.otherwise({redirectTo: '/'});
  }])
  .constant('_', window._);
