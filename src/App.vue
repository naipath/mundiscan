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
                             :routes="laserclients"
                             :laserChanged="laserChanged"
                             :activeLaser="activeLaser"/>

                    <About v-if="activeRoute === 'about'"/>

                    <AddLaser v-if="activeRoute === 'add-laser'" :handleLaser="addLaser"/>

                    <RemoveLaser v-if="activeRoute === 'remove'" :removed="removeLaser" :laser="activeLaser"/>

                    <UploadToLaser v-if="activeRoute === 'manage-laser'" :laser="activeLaser"/>

                    <Status v-if="activeRoute === 'status'" :laser="activeLaser"/>

                    <Settings v-if="activeRoute === 'settings'" :laser="activeLaser"/>

                </div>
            </div>
        </transition>
    </div>
</template>

<script>
    import Heading from "./components/Heading.vue"
    import About from "./components/About.vue"
    import UploadToLaser from "./components/UploadToLaser.vue"
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
            UploadToLaser,
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
                activeLaser: {Id: '', Ip: '', Port: 0, Name: '-'}
            }
        },
        methods: {
            routeChanged(route) {
                this.activeRoute = route
            },
            laserChanged(laser) {
                this.activeLaser = laser
                this.activeRoute = 'manage-laser'
            },
            addLaser(laser) {
                this.laserclients.push(laser)
                this.laserChanged(laser)
            },
            removeLaser(laser) {
                this.laserclients = this.laserclients.filter(client => client.Id !== laser.Id)
                this.activeLaser = ''
                this.activeRoute = 'add-laser'
            },
            initializeLaser(lasers) {
                this.laserclients = lasers
                if (this.laserclients.length > 0) {
                    this.laserChanged(this.laserclients[0])
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
