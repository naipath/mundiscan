<template>
    <div id="app">
        <transition name="fade">
            <Initializing v-if="initializing" :lasersRetrieved="initializeLaser"/>

            <div v-if="!initializing">

                <div v-if="laserclients.length <= 0">
                    <section class="hero is-fullheight is-light">
                        <div class="hero-body">
                            <div class="container">

                                <h1 class="title">Er is geen Mundi laser gevonden</h1>
                                <h1 class="subtitle">Voeg een laser toe</h1>

                                <AddLaser :handleLaser="addLaser"/>
                            </div>
                        </div>
                    </section>
                </div>

                <div class="container" v-if="laserclients.length > 0">
                    <Heading :activeRoute="activeRoute"
                             :routeChanged="routeChanged"
                             :routes="laserclients.map(laser => laser.Name)"
                             :laserChanged="laserChanged"
                             :activeLaser="activeLaser"/>

                    <About v-if="activeRoute === 'about'"/>

                    <AddLaser v-if="activeRoute === 'add-laser'" :handleLaser="addLaser"/>

                    <RemoveLaser v-if="activeRoute === 'remove'" :removed="removeLaser" :laser="getActiveLaser()"/>

                    <ManageLaser v-if="activeRoute === 'manage-laser'" :laser="getActiveLaser()"/>

                    <Status v-if="activeRoute === 'status'" :laser="getActiveLaser()"/>

                    <Settings v-if="activeRoute === 'settings'" :laser="getActiveLaser()"/>

                </div>
            </div>
        </transition>
    </div>
</template>

<script>
    import Heading from "./components/Heading.vue"
    import About from "./components/About.vue"
    import ManageLaser from "./components/ManageLaser.vue"
    import AddLaser from "./components/AddLaser.vue"
    import RemoveLaser from "./components/RemoveLaser.vue"
    import Initializing from "./components/Initializing.vue"
    import Status from "./components/Status.vue"
    import Settings from "./components/Settings.vue"

    export default {
        name: 'app',
        components: {
            Heading,
            About,
            ManageLaser,
            AddLaser,
            RemoveLaser,
            Initializing,
            Status,
            Settings,
        },
        data() {
            return {
                initializing: true,
                laserclients: [],
                activeRoute: '',
                activeLaser: ''
            }
        },
        methods: {
            routes() {
                return this.laserclients.map(laser => laser.Name)
            },
            routeChanged(route) {
                this.activeRoute = route
            },
            laserChanged(laserName) {
                this.activeLaser = laserName
                this.activeRoute = 'manage-laser'
            },
            getActiveLaser() {
                return this.laserclients.find(client => client.Name === this.activeLaser)
            },
            addLaser(laser) {
                console.log(laser)
                this.activeLaser = laser.Name
                this.laserclients.push(laser)
                this.activeRoute = 'manage-laser'
            },
            removeLaser(laserName) {
                this.laserclients = this.laserclients.filter(client => client.Name !== laserName)
                this.activeLaser = ''
                this.activeRoute = 'add-laser'
            },
            initializeLaser(lasers) {
                this.laserclients = lasers
                if (this.laserclients.length > 0) {
                    this.activeRoute = "manage-laser"
                    this.activeLaser = this.laserclients[0].Name
                }
                setTimeout(() => this.initializing = false, 500)
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

    .fade-leave-active {
        transition: opacity .5s
    }

    .fade-leave-to {
        opacity: 0
    }
</style>
