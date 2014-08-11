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
.controller('loginCtrl', function($scope, $modalInstance){

  $scope.ok = function() {
    $modalInstance.close();
  };

  $scope.cancel = function(){
    $modalInstance.dismiss('cancel');
  };
});
