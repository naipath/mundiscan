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
                isLoading: false
            }
        },
        methods: {
            addLaser() {
                this.isLoading = true
                const formData = new FormData(document.getElementById('add-laser'))
                this.handleLaser({
                    Ip: formData.get("Ip"),
                    Port: parseInt(formData.get("Port")),
                    Name: formData.get("Name"),
                })
            }
        },
    }
</script>