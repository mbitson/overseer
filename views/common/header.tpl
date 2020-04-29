<body ng-app="overseerApp" ng-controller="mainCtrl" layout="column">

<div ng-include="'/static/views/common/header.html'"></div>

<div class="main" layout-fill layout="row" flex>
    <div ng-include="'/static/views/common/navigation-sites.html'" style="z-index:2 overflow:visible !important;"></div>
    <div flex id="page" style="z-index:1">
        <div ui-view></div>