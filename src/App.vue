<template>
    <div id="app" class="container">

        <nav class="navbar">
            <div class="navbar-menu">
                <div class="navbar-start">
                    <a class="navbar-item" v-bind:class="{ ['is-active']: routes.laser }" href="#"
                       @click.prevent="reroute('laser')">
                        Laser
                    </a>
                </div>

                <div class="navbar-end">
                    <a class="navbar-item" v-bind:class="{ ['is-active']: routes.settings }" href="#"
                       @click.prevent="reroute('settings')">
                        Settings
                    </a>
                </div>
            </div>
        </nav>


        <div v-if="routes.laser">
            <div v-if="showFileChooser">
                <input type="file" accept="image/png,image/jpg,image/jpeg" v-on:change="chooseImage"/>
            </div>

            <hr v-if="showFileChooser"/>

            <div class="columns" v-if="!showFileChooser">
                <div class="column is-half">
                    <button v-if="noTextAdded" v-on:click="addText" class="button is-info">Text toevoegen</button>
                </div>
                <div class="column is-half">
                    <button class="button is-success is-pulled-right" v-on:click="onFinished">Klaar!</button>
                </div>
            </div>

            <div class="column" v-if="!showFileChooser">
                <canvas id="c" width="800" height="500"></canvas>
            </div>

            <a id="download" href="#" download="printer.png" style="display: none;"></a>
        </div>

        <div v-if="routes.settings">

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
                            <input type="checkbox" @click.prevent v-model="statusData.ExternalUpdateSingleShotNotUpdated">
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
    </div>
</template>

<script>
    import {dataURItoBlob} from "./helpers";
    let canvas;

    export default {
        name: 'app',
        data () {
            return {
                showFileChooser: true,
                noTextAdded: true,

                routes: {
                    laser: true,
                    settings: false,
                },

                statusMessage: "Press refresh to obtain a new status",
                statusData: {}
            }
        },
        methods: {
            onFinished: function () {
                document.getElementById("download").href = canvas.toDataURL('image/jpeg');
                document.getElementById("download").click();

                const data = new FormData()
                data.append('uploadfile', dataURItoBlob(canvas.toDataURL()), 'mundiscan-' + Date.now() + '.png')
                fetch('/upload', {
                    method: 'POST',
                    body: data,
                })
                    .then(result => console.log(result))
                    .catch(err => console.log(err))
            },
            addText: function () {
                this.noTextAdded = false;
                const text = new fabric.IText('Typ hier text');
                canvas.add(text);
            },
            chooseImage: function (event) {
                this.showFileChooser = false
                if (window.File && window.FileReader && window.FileList && window.Blob) {
                    const file = event.target.files[0]
                    const fileReader = new FileReader()
                    fileReader.onload = (e) => {
                        canvas = new fabric.Canvas('c');
                        canvas.selection = false;

                        const img = new Image();
                        img.src = e.target.result;
                        img.onload = () => {
                            const imageSelected = new fabric.Image(img, {left: 0, top: 0, angle: 0});
                            imageSelected.filters.push(new fabric.Image.filters.Grayscale());
                            imageSelected.applyFilters(canvas.renderAll.bind(canvas));
                            canvas.add(imageSelected);
                        }
                    }
                    fileReader.readAsDataURL(file)
                } else {
                    alert('The File APIs are not fully supported in this browser.');
                }
            },

            retrieveMundiStatus: function () {
                fetch("/statusMessage")
                    .then(response => response.json())
                    .then(msg => this.statusMessage = msg.StatusMessage)

                setTimeout(() => fetch("/statusData")
                    .then(response => response.json())
                    .then(msg => this.statusData = msg), 400)
            },

            reroute: function (route) {
                switch (route) {
                    case "laser":
                        this.routes.laser = true
                        this.routes.settings = false
                        break;

                    case "settings":
                        this.routes.laser = false
                        this.routes.settings = true
                        this.retrieveMundiStatus()
                        break;
                }
            }
        },
        mounted() {

        }
    }
</script>

<style lang="scss">
    #app {
        font-family: 'Avenir', Helvetica, Arial, sans-serif;
        -webkit-font-smoothing: antialiased;
        -moz-osx-font-smoothing: grayscale;
    }

    .navbar {
        border-bottom: 1px solid #eee;
        margin-bottom: 50px;
    }

    #c {
        border: 1px solid black;
    }

    .canvas-container {
        margin: 0 auto;
    }
</style>
