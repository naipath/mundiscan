<template>
    <div id="app" class="section">
        <form @submit.prevent="addLaser" id="add-laser">

            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">Naam</label>
                </div>
                <div class="field-body">
                    <div class="field">
                        <div class="control has-icons-left has-icons-right">
                            <input class="input" name="Name" placeholder="MundiLaser" required>
                            <span class="icon is-left">
                                <i class="fa fa-gamepad"></i>
                            </span>
                        </div>
                    </div>
                </div>
            </div>

            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">IP adres</label>
                </div>
                <div class="field-body">
                    <div class="field is-horizontal">
                        <div class="field-body control has-icons-left">
                            <input class="input" name="Ip1" placeholder="10" type="number" required min="0" max="255">
                            <span class="icon is-left">
                                <i class="fa fa-cloud"></i>
                            </span>
                        </div>
                        <div class="field-body">
                            <input type="number" class="input" name="Ip2" placeholder="0" required min="0" max="255">
                        </div>
                        <div class="field-body">
                            <input type="number" class="input" name="Ip3" placeholder="195" required min="0" max="255">
                        </div>
                        <div class="field-body">
                            <input type="number" class="input" name="Ip4" placeholder="10" required min="0" max="255">
                        </div>
                    </div>
                </div>
            </div>

            <div class="field is-horizontal">
                <div class="field-label is-normal">
                    <label class="label">Poort</label>
                </div>
                <div class="field-body">
                    <div class="field">
                        <div class="control has-icons-left has-icons-right">
                            <input class="input" name="Port" type="number" placeholder="1470" required>
                            <span class="icon is-left">
                                <i class="fa fa-fort-awesome"></i>
                            </span>
                        </div>
                    </div>
                </div>
            </div>

            <button class="button is-success is-pulled-right" v-bind:class="{'is-loading': isLoading}">
                Voeg laser toe
            </button>
        </form>

        <transition name="fade">
            <div class="error-message custom-message notification is-danger" v-if="hasError">
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
                    Ip: formData.get("Ip1") + "." + formData.get("Ip2") + "." + formData.get("Ip3") + "." + formData.get("Ip4"),
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
.custom-message {
    position: absolute !important;
    top: 0!important;
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