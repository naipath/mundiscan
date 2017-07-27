<template>
    <div id="app">

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
</template>

<script>
    import {dataURItoBlob} from "./../helpers";
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
                }
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
        }
    }
</script>
<style lang="scss">
    #c {
        border: 1px solid black;
    }
    .canvas-container {
        margin: 0 auto;
    }
</style>
