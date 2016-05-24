<div id="content">
    <md-content class="md-primary" flex layout-padding layout="row">
        <span flex></span>
        <form flex name="userForm" method="post" action="">
            <p class="clean-text text-faint" style="font-size: 32px;">
                - BETA ONLY -
            </p>
            <p class="clean-text">
                Registration Form
            </p>
            <<<if .flash.error>>>
                <div class="alert alert-danger">
                    <strong><<<.flash.error>>></strong>
                </div>
            <<<end>>>
            <<<if .Errors>>>
                <div class="alert alert-danger">
                    <<<range $rec := .Errors>>>
                        <strong><<<$rec>>></strong>
                    <<<end>>>
                </div>
            <<<end>>>
            <div flex layout="row">
                <md-input-container flex>
                    <label>First Name</label>
                    <input name="first_name" ng-model="first_name" />
                </md-input-container>
                <md-input-container flex>
                    <label>Last Name</label>
                    <input name="last_name" ng-model="last_name" />
                </md-input-container>
            </div>
            <md-input-container flex>
                <label>Email</label>
                <input name="email" ng-model="email" />
            </md-input-container>
            <md-input-container flex>
                <label>Password</label>
                <input name="password" type="password" ng-model="password" />
            </md-input-container>
            <md-input-container flex>
                <label>Confirm Password</label>
                <input name="password_2" type="password" ng-model="password_2" />
            </md-input-container>

            <div flex layout="row">
                <a flex href="/user/login/home">
                    Already registered? Login
                </a>
                <span flex></span>
                <md-button flex ng-click="submit()" class="md-raised md-primary">
                    Register
                </md-button>
            </div>
        </form>
        <span flex></span>
    </md-content>
</div>