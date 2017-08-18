<template>
    <div id="app">

        <div class="columns">
            <h1 class="title column">Status</h1>
            <div class="column">
                <button class="button is-info is-pulled-right" @click="retrieveMundiStatus">Refresh</button>
            </div>
        </div>
        <input class="input" v-model="statusMessage" readonly>

        <hr/>

        <h1 class="title">Status data</h1>

        <div class="columns">
            <div class="column" v-for="metricRow in metricRows">
                <div v-for="metric in metricRow">
                    <label class="checkbox">
                        <input type="checkbox" @click.prevent v-model="metric.value">
                        {{metric.label}}
                    </label>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    export default {
        name: 'app',
        props: {
            laser: Object,
        },
        data() {
            const status = {
                statusMessage: "Press refresh to obtain a new status",
                statusData: {
                    PrintStatus: false,
                    FailureStatus: false,
                    WarningStatus: false,
                    MaintenanceStatus: false,
                    DetectorFault: false,
                    MarkNotLoaded: false,
                    CodeParametersNotLoaded: false,
                    TestingLaser: false,
                    DisableShutter: false,
                    ExternalUpdateSingleShotNotUpdated: false,
                    ComPortDisconnected: false,
                    BarcodeError: false,
                    EStop: false,
                    ExternalInterlocks: false,
                    CoolantTemperature: false,
                    DC24Volts: false,
                    ShutterKeyswitch: false,
                    Keyswitch: false,
                    DC48Volts: false,
                    CoolantFlow: false,
                    EmissionIndicator: false,
                    ShutterPrint: false,
                    ShutterStandby: false,
                    VSWRError: false,
                    OverModulation: false,
                    TachoFault: false,
                    RFStatus: false,
                    SystemEnable: false,
                    VapourExtractorFault: false,
                    GalvoPower: false,
                    GalvoTemperature: false,
                    GalvoCableDisconnected: false,
                }
            }
            return {
                statusMessage: status.statusMessage,
                statusData: status.statusData,
                metricRows: [
                    [
                        {label: 'Print status', value: status.statusData.PrintStatus},
                        {label: 'Failure status', value: status.statusData.FailureStatus},
                        {label: 'Warning status', value: status.statusData.WarningStatus},
                        {label: 'Maintenance status', value: status.statusData.MaintenanceStatus},
                    ],
                    [
                        {label: 'Detector fault', value: status.statusData.DetectorFault},
                        {label: 'Mark not loaded', value: status.statusData.MarkNotLoaded},
                        {label: 'Code parameters not loaded', value: status.statusData.CodeParametersNotLoaded},
                        {label: 'Testing laser', value: status.statusData.TestingLaser},
                        {label: 'Disable shutter', value: status.statusData.DisableShutter},
                        {label: 'External update, single shot not updated', value: status.statusData.ExternalUpdateSingleShotNotUpdated},
                        {label: 'Com port disconnected', value: status.statusData.ComPortDisconnected},
                        {label: 'Barcode error', value: status.statusData.BarcodeError},
                    ],
                    [
                        {label: 'E stop', value: status.statusData.EStop},
                        {label: 'External interlocks', value: status.statusData.ExternalInterlocks},
                        {label: 'Coolant temperature', value: status.statusData.CoolantTemperature},
                        {label: 'DC 24 volts', value: status.statusData.DC24Volts},
                        {label: 'Shutter keyswitch', value: status.statusData.ShutterKeyswitch},
                        {label: 'Keyswitch', value: status.statusData.Keyswitch},
                        {label: 'DC 48 volts', value: status.statusData.DC48Volts},
                    ],
                    [
                        {label: 'Coolant flow', value: status.statusData.CoolantFlow},
                        {label: 'Emission indicator', value: status.statusData.EmissionIndicator},
                        {label: 'Shutter print', value: status.statusData.ShutterPrint},
                        {label: 'Shutter standby', value: status.statusData.ShutterStandby},
                        {label: 'VSWR error', value: status.statusData.VSWRError},
                        {label: 'Over modulation', value: status.statusData.OverModulation},
                        {label: 'Tacho fault', value: status.statusData.TachoFault},
                        {label: 'RF status', value: status.statusData.RFStatus},
                    ],
                    [
                        {label: 'System enable', value: status.statusData.SystemEnable},
                        {label: 'Vapour extractor fault', value: status.statusData.VapourExtractorFault},
                        {label: 'Galvo power', value: status.statusData.GalvoPower},
                        {label: 'Galvo temperature', value: status.statusData.GalvoTemperature},
                        {label: 'Galvo cable disconnected', value: status.statusData.GalvoCableDisconnected},
                    ],
                ]
            }
        },
        methods: {
            retrieveMundiStatus: function () {
                fetch("/laserclients/" + this.laser.Id + "/status")
                    .then(response => response.json())
                    .then(msg => {
                        this.statusMessage = msg.Message
                        this.statusData = msg.Status
                    })
                    .catch(err => this.statusData = "Something went wrong with retrieving the status")
            },
        },
        mounted() {
            this.retrieveMundiStatus()
        }
    }
</script>
