angular.module('myApp', ['cui'])
    .controller('AppController', function ($scope,
                       $http,
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

            var toUtc = function (d){

                return moment.utc(
                    moment(d)
                    .local()
                    .format()
                    .substring(0,19)).toJSON()

            }

            job = {
                customer: $scope.job.customer,
                from: toUtc($scope.job.from),
                to: toUtc($scope.job.to)
            }   

            cuiLoading(
                srv.query(job)
                    .then(function (res) {

                        $scope.job.count = res.count;
                        $scope.job.size = res.size;
                        $scope.job.eta = res.eta;

                    },
                    function (err) {
                        cuiAlertService.warning(err);
                    })
            );
        }


    $scope.submit = function(){

            //todo:make it post
        srv = cuiDataSourceService('/api/job/create');

            cuiLoading(
                srv.query($scope.job)
                    .then(function (res) {
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
