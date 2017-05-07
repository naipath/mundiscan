<template>
    <div id="app" class="container">

        <h1 class="title">Stap 1: Bovenkant</h1>

        <div v-if="showFileChooser">
            <input type="file"
                   accept="image/png,image/jpg,image/jpeg"
                   v-on:change="chooseImage"/>
        </div>

        <hr/>

        <div class="columns" v-if="!showFileChooser">
            <div class="column is-half">
                <button v-if="noTextAdded" v-on:click="addText" class="button is-info">Text toevoegen</button>
            </div>
            <div class="column is-half">
                <button class="button is-success" v-on:click="onFinished">Klaar!</button>
            </div>
        </div>

        <div class="column" v-if="!showFileChooser">
            <canvas id="c" width="300" height="300"></canvas>
        </div>

        <!--<a id="download" href="#" download="printer.png" style="display: none;"></a>-->

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
            }
        },
        methods: {
            onFinished: function () {
//                document.getElementById("download").href = canvas.toDataURL('image/jpeg');
//                document.getElementById("download").click();

                const data = new FormData()
                data.append('uploadfile', dataURItoBlob(canvas.toDataURL()), 'canvas.png')
                fetch('http://localhost:8100/upload', {
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
        text-align: center;
        color: #2c3e50;
        margin-top: 60px;
    }

    h1, h2 {
        font-weight: normal;
    }

    #c {
        border: 1px solid black;
    }

    .canvas-container {
        margin: 0 auto;
    }
</style>
