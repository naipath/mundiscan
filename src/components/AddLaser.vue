<template>
    <div id="app">
        <form @submit.prevent="addLaser" id="add-laser">
            <div class="field">
                <div class="control has-icons-left has-icons-right">
                    <input class="input" name="Name" placeholder="Naam" required>
                    <span class="icon is-left"><i class="fa fa-gamepad"></i></span>
                </div>
            </div>
            <div class="field">
                <div class="control has-icons-left has-icons-right">
                    <input class="input" name="Ip" placeholder="IP-adres" required>
                    <span class="icon is-left"><i class="fa fa-cloud"></i></span>
                </div>
            </div>
            <div class="field">
                <div class="control has-icons-left has-icons-right">
                    <input class="input" name="Port" type="number" placeholder="poort" required>
                    <span class="icon is-left"><i class="fa fa-fort-awesome"></i></span>
                </div>
            </div>
            <button class="button is-success is-pulled-right"
                    v-bind:class="{'is-loading': isLoading}">
                Voeg laser toe
            </button>
        </form>

        <transition name="fade">
            <div class="error-message notification is-danger" v-if="hasError">
                <button class="delete" @click="resetError"></button>
                Er kan geen connectie worden opgebouwd met de mundi laser. Controleer of de laser correct is aangesloten.
            </div>
        </transition>
    </div>
</template>

<script>
    export default {
        name: 'app',
        props: {
            handleLaser: Function
        },
        data() {
            return {
                isLoading: false,
                hasError: false,
            }
        },
        methods: {
            addLaser() {
                this.isLoading = true
                const formData = new FormData(document.getElementById('add-laser'))
                fetch("/laserclients", {
                    method: 'POST',
                    body: JSON.stringify({
                        Ip: formData.get("Ip"),
                        Port: parseInt(formData.get("Port")),
                        Name: formData.get("Name"),
                    })
                })
                    .then(response => {
                        this.isLoading = false
                        if (response.ok) {
                            response.json().then(result => this.handleLaser(result))
                        } else {
                            this.hasError = true
                        }
                    })
            },
            resetError() {
                this.hasError = false
            }
        },
    }
</script>
<style lang="scss">
    .error-message {
        position: absolute;
        top: 0;
        width: 100%;
    }

    .fade-enter-active, .fade-leave-active {
        transition: opacity .5s
    }

    .fade-enter, .fade-leave-to {
        opacity: 0
    }
</style>