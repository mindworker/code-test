var app = angular.module('app', []);

app.controller('formCtrl', function($scope, $window, $http) {
  $scope.width = $window.innerWidth;
  $scope.height = $window.innerHeight;

  $http.post("/sessions")
       .then(function(resp){
	 $scope.sessionID = resp.data.session;
       }, function(err) {
	 // @TODO handle error gracefully
	 alert(err);
       });

  
  $scope.paste = function(event) {
    console.log(event.target.id);
  }

  $scope.key = function(event) {
    // ignore ctrl and meta (CMD in Macs) key which effectively
    // ignores copy pasting
    if (event.ctrlKey || event.metaKey) {
      return;
    }

    // Initialize starting time if not done yet
    if (!$scope.timeStarted) {
      console.log("Timer started")
      $scope.timeStarted = new Date();
    }
  }

  // @TODO disable form on submit
  $scope.submit = function() {
    var timeTaken = Math.ceil(((new Date())-$scope.timeStarted)/1000);
    console.log("Time taken:", timeTaken);
  }

});

app.directive('resize', ['$window', function ($window) {
  function link(scope, element, attrs){
    angular.element($window).bind('resize', function(){

      var width = $window.innerWidth;
      var height = $window.innerHeight;

      console.log("old:", scope.width, scope.height)
      console.log("new:", width, height);

      // manuall $digest required as resize event
      // is outside of angular
      scope.$digest();
    });
  }

  return {link: link};
}]);
