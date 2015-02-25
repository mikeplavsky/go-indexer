angular.module('myApp', ['cui'])
    .controller('AppController', function ($scope,
                                           cuiDataSourceService,
                                           cuiAlertService,
                                           cuiLoading) {

        $scope.applicationframe = {};

        $scope.job = {}

        $scope.job.customers = []

        customerSvc = cuiDataSourceService('api/customers');

        cuiLoading(customerSvc.query()
            .then(function (res) {
                $scope.job.customers = res.result.map(function (value, key) {
                    return {
                        label: value,
                        description: value
                    };
                });
            },
		function (err) {
                   cuiAlertService.warning(err);
	    }));

        $scope.calculate = function () {

            srv = cuiDataSourceService('/api/job');

            cuiLoading(
                srv.query($scope.job)
                    .then(function (res) {

                        $scope.job.count = res.count;
                        $scope.job.size = res.size;

                    },
                    function (err) {
                        cuiAlertService.warning(err);
                    })
            );
        }
    })
    .controller('AboutBoxCtrl', function ($scope, cuiAboutBox) {
        var aboutBox = cuiAboutBox({
            applicationName: 'Logs Indexer'
        });
        $scope.showAboutBox = aboutBox.modal.show;
    });
