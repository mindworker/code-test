var app = angular.module('app', []);

app.controller('formCtrl', function($scope, $window, $http) {
  $scope.websiteUrl = window.location.href
  $scope.windowSize = new Dimension($window.innerWidth, $window.innerHeight)

  $http.post("/sessions").then(function(resp){
    $scope.session = resp.data.sessionId;
    console.log($scope.session)
    $scope.active = true;
  }, function(err) {
    $scope.error = "Error on getting valid session. Is the backend running? Is this file served by the backend?";
    console.error(err);
  });

  $scope.sendAction = function(data) {
    var body = angular.extend({
      sessionId: $scope.session,
      websiteUrl: $scope.websiteUrl
    }, data);

    $http.post("/actions", body).then(function(){}, function(err) {
      $scope.error = "Send action of " + data.eventType + "failed";
      console.error(err.data);
    });
  }
  
  $scope.paste = function(event) {
    $scope.sendAction({
      eventType: "copyAndPaste",
      pasted: true,
      formId: event.target.id
    });
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

    $scope.sendAction({
      eventType: "timeTaken",
      time: timeTaken
    });

    $scope.active = false;
  }

});

app.directive('resize', ['$window', function ($window) {
  function link(scope, element, attrs){
    angular.element($window).bind('resize', function(){

      var newSize = new Dimension($window.innerWidth, $window.innerHeight)
      console.log("old:", scope.windowSize, "\nnew:", newSize);

      scope.sendAction({
	eventType: "resizeWindow",
	resizeFrom: scope.windowSize,
	resizeTo: newSize
      });

      // manuall $digest required as resize event
      // is outside of angular
      scope.$digest();

    });
  }

  return {link: link};
}]);

function Dimension(width, height) {
  // Somehow the code-test requires the dimensions to be strings
  this.width = String(width);
  this.height = String(height);
}
