angular.module('appServices', [])
  .constant('USER_ROLES', {
    all: '*',
    admin: 'admin',
    guest: 'guest'
  })
  .factory('Login', function(Session){

  });
  .service('Session', function(){

  });
