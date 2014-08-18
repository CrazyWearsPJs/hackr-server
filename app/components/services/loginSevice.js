angular.module('appServices', [])
  .constant('AUTH_EVENTS', {
    loginSuccess: 'auth-login-success',
    loginFailed: 'auth-login-failed',
    logoutSuccess: 'auth-logout-success'
  })
  .constant('USER_ROLES', {
      all: '*',
      admin: 'admin',
      guest: 'guest',
      user: 'user'
  })
  .factory('AuthService', function($http, Session){
    var authService = {};

    authService.login = function(credentials) {
        return $http
            .post('/login', credentials)
            .then(function(res) {
                localStorage.setItem('_hackr_session_id', res.id);
                Session.create(res.id, user.email, res.user.role);
                return res.user;
            });
    };
  })
  .service('Session', function(){
    this.create = function(sessionId, userEmail, userRole) {
        this.id = id;
        this.userEmail = email;
        this.userRole = userRole;
    };

    this.destroy = function(){
        this.id = null;
        this.userEmail = null;
        this.userRole = null;
    };

    return this;
  });
