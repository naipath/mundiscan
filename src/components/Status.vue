<template>
    <div id="app" class="section">

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
            <div class="column">
                <div v-for="metric in getRowData('PrintStatus', 'FailureStatus', 'WarningStatus', 'MaintenanceStatus')"
                    :key="metric[0]">
                    <label class="checkbox">
                        <input type="checkbox" onclick="return false;" v-model="metric[1]">{{metric[0]}}
                    </label>
                </div>
            </div>
            <div class="column">
                <div v-for="metric in getRowData('DetectorFault', 'MarkNotLoaded', 'CodeParametersNotLoaded', 'TestingLaser', 'DisableShutter', 'ExternalUpdateSingleShotNotUpdated', 'ComPortDisconnected', 'BarcodeError')"
                    :key="metric[0]">
                    <label class="checkbox">
                        <input type="checkbox" onclick="return false;" v-model="metric[1]">{{metric[0]}}
                    </label>
                </div>
            </div>
            <div class="column">
                <div v-for="metric in getRowData('EStop', 'ExternalInterlocks', 'CoolantTemperature', 'DC24Volts', 'ShutterKeyswitch', 'Keyswitch', 'DC48Volts')"
                    :key="metric[0]">
                    <label class="checkbox">
                        <input type="checkbox" onclick="return false;" v-model="metric[1]">{{metric[0]}}
                    </label>
                </div>
            </div>
            <div class="column">
                <div v-for="metric in getRowData('CoolantFlow', 'EmissionIndicator', 'ShutterPrint', 'ShutterStandby', 'VSWRError', 'OverModulation', 'TachoFault', 'RFStatus')"
                    :key="metric[0]">
                    <label class="checkbox">
                        <input type="checkbox" onclick="return false;" v-model="metric[1]">{{metric[0]}}
                    </label>
                </div>
            </div>
            <div class="column">
                <div v-for="metric in getRowData('SystemEnable', 'VapourExtractorFault', 'GalvoPower', 'GalvoTemperature', 'GalvoCableDisconnected')"
                    :key="metric[0]">
                    <label class="checkbox">
                        <input type="checkbox" onclick="return false;" v-model="metric[1]">{{metric[0]}}
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
        return {
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
    },
    methods: {
        getRowData(...args) {
            return Object.entries(this.statusData).filter(entry => args.includes(entry[0]))
        },
        retrieveMundiStatus() {
            fetch("/laserclients/" + this.laser.Id + "/status")
                .then(response => {
                    if (response.ok) {
                        return response.json()
                    }
                    throw new Error()
                })
                .then(response => response.json())
                .then(msg => {
                    this.statusMessage = msg.Message
                    this.statusData = msg.Status
                })
                .catch(() => this.statusMessage = "Something went wrong with retrieving the status")
        },
    },
    mounted() {
        this.retrieveMundiStatus()
    }
}
</script>
