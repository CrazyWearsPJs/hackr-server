angular.module('appControllers', ['loginController'])
  .controller('splashCtrl', function($scope, $modal){
    $scope.open = function(size){
      modalInstance = $modal.open({
        templateUrl: 'templates/login.html',
        controller: 'loginCtrl',
        size: size
      });
    };
});

angular.module('loginController', [])
.controller('loginCtrl', function($scope, $rootScope, $modalInstance, AuthService){
    
    $scope.credentials = {
        email: "",
        password: ""
    };
  
    $scope.login = function(credentials) {
        AuthService.login(credentials);
    };



  $scope.cancel = function(){
    $modalInstance.dismiss('cancel');
  };
});
