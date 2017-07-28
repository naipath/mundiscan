<template>
    <div id="settings">

        <div class="columns">
            <h1 class="title column">Laser client</h1>
            <div class="column">
                <button class="button is-info is-pulled-right" @click="getSettings">Refresh</button>
            </div>
        </div>
        <hr/>

        <div class="field is-horizontal">
            <div class="field-label is-normal">
                <label class="label">IP</label>
            </div>
            <div class="field-body">
                <div class="field">
                    <input class="input" type="text" v-model="client.ip" readonly>
                </div>
            </div>
        </div>

        <div class="field is-horizontal">
            <div class="field-label is-normal">
                <label class="label">Port</label>
            </div>
            <div class="field-body">
                <div class="field">
                    <input class="input" type="text" v-model="client.port" readonly>
                </div>
            </div>
        </div>

        <hr />

        <div class="columns">
            <h1 class="title column">Counters</h1>
            <div class="column">
                <button class="button is-warning is-pulled-right" @click="resetCount">Reset current count</button>
            </div>
        </div>
        <hr/>
        <div class="field is-horizontal">
            <div class="field-label is-normal">
                <label class="label">Recent</label>
            </div>
            <div class="field-body">
                <div class="field">
                    <input class="input" type="text" v-model="counters.Recent" readonly>
                </div>
            </div>
        </div>
        <div class="field is-horizontal">
            <div class="field-label is-normal">
                <label class="label">Life time</label>
            </div>
            <div class="field-body">
                <div class="field">
                    <input class="input" type="text" v-model="counters.Lifetime" readonly>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'settings',
        data () {
            return {
                client: {
                    ip: '-',
                    port: '-'
                },
                counters: {
                    Lifetime: '-',
                    Recent: '-'
                }
            }
        },
        methods: {
            getSettings: function () {
                fetch('/counters')
                    .then(resp => resp.json())
                    .then(msg => this.counters = msg)
                fetch('/clientSettings')
                    .then(resp => resp.json())
                    .then(msg => this.client = msg)
            },
            resetCount: function() {
                fetch('/resetCurrentCount').then(() => this.getSettings())
            }
        },
        mounted() {
            this.getSettings()
        }
    }
</script>