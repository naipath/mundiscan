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
                       v-bind:class="{ ['is-active']: routes.settings }"
                       href="#"
                       @click.prevent="reroute('settings')">
                        Settings
                    </a>
                </div>
            </div>
        </nav>

        <div v-if="routes.laser">
            <laser/>
        </div>

        <div v-if="routes.settings">
            <status/>
        </div>
    </div>
</template>

<script>
    import Status from "./components/Status.vue";
    import Laser from "./components/Laser.vue";

    export default {
        name: 'app',
        components: {
            'status': Status,
            'laser': Laser
        },
        data () {
            return {
                routes: {
                    laser: true,
                    settings: false,
                }
            }
        },
        methods: {
            reroute: function (route) {
                switch (route) {
                    case "laser":
                        this.routes.laser = true
                        this.routes.settings = false
                        break;

                    case "settings":
                        this.routes.laser = false
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
