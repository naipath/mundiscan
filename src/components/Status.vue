<template>
    <div id="app">

        <div class="columns">
            <h1 class="title column">Status message</h1>
            <div class="column">
                <button class="button is-info is-pulled-right" @click="retrieveMundiStatus">Refresh</button>
            </div>
        </div>
        <input class="input" type="text" v-model="statusMessage" readonly>

        <hr/>

        <h1 class="title">Status data</h1>

        <div class="columns">
            <div class="column">
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.PrintStatus">
                        Print status
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.FailureStatus">
                        Failure status
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.WarningStatus">
                        Warning status
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.MaintenanceStatus">
                        Maintenance status
                    </label>
                </div>
            </div>
            <div class="column">
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.DetectorFault">
                        Detector fault
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.MarkNotLoaded">
                        Mark not loaded
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.CodeParametersNotLoaded">
                        Code parameters not loaded
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.TestingLaser">
                        Testing laser
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.DisableShutter">
                        Disable shutter
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent
                               v-model="statusData.ExternalUpdateSingleShotNotUpdated">
                        External update, single shot not updated
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.ComPortDisconnected">
                        Com port disconnected
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.BarcodeError">
                        Barcode error
                    </label>
                </div>
            </div>
            <div class="column">
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.EStop">
                        E stop
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.ExternalInterlocks">
                        External interlocks
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.CoolantTemperature">
                        Coolant temperature
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.DC24Volts">
                        DC 24 volts
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.ShutterKeyswitch">
                        Shutter keyswitch
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.Keyswitch">
                        Keyswitch
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.DC48Volts">
                        DC 48 volts
                    </label>
                </div>


            </div>
            <div class="column">
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.CoolantFlow">
                        Coolant Flow
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.EmissionIndicator">
                        Emission indicator
                    </label>
                </div>
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.ShutterPrint">
                        Shutter print
                    </label>
                </div>
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.ShutterStandby">
                        Shutter standby
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.VSWRError">
                        VSWR error
                    </label>
                </div>

                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.OverModulation">
                        Over modulation
                    </label>
                </div>
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.TachoFault">
                        Tacho fault
                    </label>
                </div>
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.RFStatus">
                        RF status
                    </label>
                </div>
            </div>

            <div class="column">
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.SystemEnable">
                        System enable
                    </label>
                </div>
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.VapourExtractorFault">
                        Vapour extractor fault
                    </label>
                </div>
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.GalvoPower">
                        Galvo power
                    </label>
                </div>
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.GalvoTemperature">
                        Galvo temperature
                    </label>
                </div>
                <div>
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="statusData.GalvoCableDisconnected">
                        Galvo cable disconnected
                    </label>
                </div>
            </div>

        </div>
    </div>
</template>

<script>
    export default {
        name: 'app',
        data() {
            return {
                statusMessage: "Press refresh to obtain a new status",
                statusData: {}
            }
        },
        methods: {
            retrieveMundiStatus: function () {
                fetch("/statusMessage")
                    .then(response => response.json())
                    .then(msg => this.statusMessage = msg.StatusMessage)
                    .catch(err => this.statusData = "Something went wrong with retrieving the status")

                setTimeout(() => fetch("/statusData")
                    .then(response => response.json())
                    .then(msg => this.statusData = msg)
                    .catch(err => this.statusData = "Something went wrong with retrieving the status"), 400)

            },
        },
        mounted() {
            this.retrieveMundiStatus()
        }
    }
</script>
