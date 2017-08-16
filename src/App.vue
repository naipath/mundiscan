<template>
    <div id="app">
        <transition name="fade">
            <Initializing v-if="initializing"/>

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
                    <Heading :activeRoute="activeRoute" :routeChanged="routeChanged" :routes="laserclients.map(laser => laser.Name)"/>

                    <About v-if="activeRoute === 'about'"/>

                    <AddLaser v-if="activeRoute === 'add-laser'" :handleLaser="addLaser"/>

                    <ManageLaser v-if="activeRoute.indexOf('manage-laser') === 0" :laser="laserclients.find(client => client.Name == 'testing')"/>

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
    import Initializing from "./components/Initializing.vue"

    export default {
        name: 'app',
        components: {
            Heading,
            About,
            ManageLaser,
            AddLaser,
            Initializing,
        },
        data() {
            return {
                initializing: true,
                laserclients: [],
                activeRoute: ''
            }
        },
        methods: {
            routes() {
                return this.laserclients.map(laser => laser.Name)
            },
            routeChanged(route) {
                this.activeRoute = route
            },
            addLaser(laser) {
                fetch("/laserclients", {
                    method: 'POST',
                    body: JSON.stringify(laser)
                })
                    .then(result => result.json())
                    .then(result => this.laserclients.push(result))
            }
        },
        mounted: function () {
            fetch("/laserclients")
                .then(result => result.json())
                .then(result => {
                    this.laserclients = result
                    if (this.laserclients.length > 0) {
                        this.activeRoute = "manage-laser/" + this.laserclients[0].Name
                    }
                    setTimeout(() => this.initializing = false, 500)
                })
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
