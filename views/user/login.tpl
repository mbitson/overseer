<div id="content">
    <md-content class="md-primary" flex layout-padding layout="row">
        <span flex></span>
        <form flex name="userForm" method="post" action="">
            <p class="clean-text text-faint" style="font-size: 32px;">
                - BETA ONLY -
            </p>
            <p class="clean-text">
                Please login to manage your assets.
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
            <md-input-container flex>
                <label>Email</label>
                <input name="email" ng-model="email" />
            </md-input-container>
            <md-input-container flex>
                <label>Password</label>
                <input name="password" type="password" ng-model="password" />
            </md-input-container>
            <div flex layout="row">
                <md-button flex href="/user/register" class="md-raised md-accent">
                    Register
                </md-button>
                <span flex></span>
                <md-button flex ng-click="submit()" class="md-raised md-primary">
                    Login
                </md-button>
            </div>
        </form>
        <span flex></span>
    </md-content>
</div>
