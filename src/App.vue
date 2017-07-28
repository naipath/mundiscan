<template>
    <div id="app" class="container">

        <nav class="navbar">
            <div class="navbar-menu">
                <div class="navbar-start">
                    <a class="navbar-item"
                       v-bind:class="{ ['is-active']: routes.laser }"
                       href="#"
                       @click.prevent="reroute('laser')">
                        Laser
                    </a>
                </div>

                <div class="navbar-end">
                    <a class="navbar-item"
                       v-bind:class="{ ['is-active']: routes.status }"
                       href="#"
                       @click.prevent="reroute('status')">
                        <i class="fa fa-stethoscope"></i>&nbsp;Status
                    </a>

                    <a class="navbar-item"
                       v-bind:class="{ ['is-active']: routes.settings }"
                       href="#"
                       @click.prevent="reroute('settings')">
                        <i class="fa fa-cog"></i>&nbsp;Settings
                    </a>
                </div>
            </div>
        </nav>

        <div v-if="routes.laser">
            <laser/>
        </div>

        <div v-if="routes.status">
            <status/>
        </div>

        <div v-if="routes.settings">
            <settings/>
        </div>
    </div>
</template>

<script>
    import Status from "./components/Status.vue";
    import Laser from "./components/Laser.vue";
    import Settings from "./components/Settings.vue";

    export default {
        name: 'app',
        components: {
            'status': Status,
            'laser': Laser,
            'settings': Settings
        },
        data () {
            return {
                routes: {
                    laser: true,
                    status: false,
                    settings: false,
                }
            }
        },
        methods: {
            reroute: function (route) {
                switch (route) {
                    case "laser":
                        this.routes.laser = true
                        this.routes.status = false
                        this.routes.settings = false
                        break;

                    case "status":
                        this.routes.laser = false
                        this.routes.status = true
                        this.routes.settings = false
                        break;

                    case "settings":
                        this.routes.laser = false
                        this.routes.status = false
                        this.routes.settings = true
                        break;
                }
            }
        }
    }
</script>

<style lang="scss">
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    .navbar {
        border-bottom: 1px solid #eee;
        margin-bottom: 50px;
    }
</style>
