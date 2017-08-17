<template>
    <section id="app" class="hero">
        <div class="hero-body">
            <h1 class="subtitle">
                Weet je zeker dat je {{laser.Name}} wilt verwijderen?
            </h1>
            <a class="button is-danger" @click.prevent="removeLaser"
               v-bind:class="{'is-loading': isLoading}">
                <span>Verwijder</span>
                <span class="icon is-small">
                  <i class="fa fa-times"></i>
                </span>
            </a>
        </div>
    </section>
</template>
<script>
    export default {
        name: 'app',
        props: {
            laser: Object,
            removed: Function
        },
        data() {
            return {
                isLoading: false
            }
        },
        methods: {
            removeLaser() {
                this.isLoading = true
                fetch("/laserclients/" + this.laser.Name, {method: "DELETE"})
                    .then(() => {
                        this.isLoading = false
                        this.removed(this.laser.Name)
                    })
            }
        }
    }
</script>