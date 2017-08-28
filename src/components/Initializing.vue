<template>
    <section id="app" class="hero is-fullheight is-light">
        <div class="hero-body">
            <div class="container has-text-centered">
                <div v-if="!hasError">
                    <h1 class="title">Initialiseren</h1>
                    <h1 class="title button is-light is-loading"></h1>
                </div>
                <transition name="fade">
                    <div class="error-message notification is-danger" v-if="hasError">
                        <div>Er is iets misgegaan bij het opstarten van de applicatie. Probeer
                            het later nogmaals.</div>
                        <br/>
                        <button class="button" @click="resetError">Opnieuw proberen</button>
                    </div>
                </transition>
            </div>
        </div>
    </section>
</template>
<script>
export default {
    name: 'app',
    props: {
        lasersRetrieved: Function,
    },
    data() {
        return {
            hasError: false,
        }
    },
    methods: {
        resetError() {
            this.hasError = false
            setTimeout(() => this.retrieveLaserClients(), 1000)
        },
        retrieveLaserClients() {
            fetch("/laserclients")
                .then(result => {
                    if (result.ok) {
                        return result.json()
                    }
                    throw new Error()
                })
                .then(result => this.lasersRetrieved(result))
                .catch(() => this.hasError = true)

        },
    },
    mounted() {
        this.retrieveLaserClients()
    }
}
</script>
<style lang="scss" scoped>
.error-message {
    position: absolute !important;
    top: 0 !important;
    width: 100% !important;
}

.fade-enter-active,
.fade-leave-active {
    transition: opacity .5s
}

.fade-enter,
.fade-leave-to {
    opacity: 0
}
</style>