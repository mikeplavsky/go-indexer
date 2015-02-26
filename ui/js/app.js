angular.module('myApp', ['cui'])
    .controller('AppController', function ($scope,
                                           $http,
                                           cuiDataSourceService,
                                           cuiAlertService,
                                           cuiLoading,
                                           $interval) {

        $scope.applicationframe = {};
        $scope.job = {}
        $scope.eta = {}
        $scope.job.customers = []


        etaSvc = cuiDataSourceService('api/eta');

        $scope.refreshEta = function() {

            etaSvc.query()
            .then(
                function(eta){

                    $scope.eta.files = eta.files;
                    $scope.eta.time = eta.time;

                    $scope.queue = eta.queue

                },
                function(err){
                }
            );
             
            console.log($scope.eta)

        }

        $scope.refreshEta()
        $interval($scope.refreshEta, 10000)

        $scope.toUtc = function (d){

            return moment.utc(
                moment(d)
                .local()
                .format()
                .substring(0,19)).toJSON()

        }

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

            job = {
                customer: $scope.job.customer,
                from: $scope.toUtc($scope.job.from),
                to: $scope.toUtc($scope.job.to)
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

            job = {
                customer: $scope.job.customer,
                from: $scope.toUtc($scope.job.from),
                to: $scope.toUtc($scope.job.to)
            }   

            cuiLoading(
                srv.query(job)
                    .then(function (res) {
                        $scope.refreshEta();
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
