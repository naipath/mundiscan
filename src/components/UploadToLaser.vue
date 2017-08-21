<template>
    <div id="app">

        <div class="file is-pulled-right is-success">
            <label class="file-label">
                <input class="file-input"
                       type="file"
                       accept="image/png,image/jpg,image/jpeg"
                       v-on:change="chooseImage">
                <span class="file-cta">
                      <span class="file-icon">
                        <i class="fa fa-upload"></i>
                      </span>
                      <span class="file-label">
                        Kies een afbeelding ...
                      </span>
                    </span>
            </label>
        </div>

        <div id="container"></div>

        <a id="download" href="#" download="printer.png" style="display: none;"></a>

    </div>
</template>

<script>
    import {dataURItoBlob} from "./../helpers";

    let stage, layer, rect;

    export default {
        name: 'app',
        props: {
            laser: Object
        },
        data() {
            return {}
        },
        methods: {
            chooseImage: function (event) {
                if (window.File && window.FileReader && window.FileList && window.Blob) {
                    const file = event.target.files[0]
                    const fileReader = new FileReader()
                    fileReader.onload = e => {
                        let imageObj = new Image()
                        imageObj.onload = () => {
                            this.addLogo(imageObj)
                        }
                        imageObj.src = e.target.result
                    }
                    fileReader.readAsDataURL(file)
                }
            },
            addLogo(imageObj) {
                let logo = new Konva.Image({
                    x: stage.getWidth() / 2,
                    y: stage.getHeight() / 2,
                    image: imageObj,
                    draggable: true,
                    stroke: 'blue',
                    strokeWidth: 1,
                    strokeEnabled: false
                })

                logo.on('dblclick', () => {
                    logo.destroy()
                    layer.draw()
                })

                logo.on('mouseover', () => {
                    logo.strokeEnabled(true)
                    layer.draw()
                })

                logo.on('mouseout', () => {
                    logo.strokeEnabled(false)
                    layer.draw()
                })

                layer.add(logo)
                logo.moveToBottom()
                layer.draw()
            },
            createBlurEffect(shape) {
                const createRect = (x, y, width, height) => ({
                    x,
                    y,
                    width,
                    height,
                    fill: 'rgba(255, 255, 255, 0.4)',
                    listening: false,
                })
                layer.add(new Konva.Rect(createRect(
                    0, 0,
                    stage.getWidth(), stage.getHeight() / 2 - (shape / 2)
                )))

                layer.add(new Konva.Rect(createRect(
                    stage.getWidth() / 2 + (shape / 2), stage.getHeight() / 2 - (shape / 2),
                    stage.getWidth(), stage.getHeight()
                )))

                layer.add(new Konva.Rect(createRect(
                    0, stage.getHeight() / 2 - (shape / 2),
                    stage.getWidth() / 2 - (shape / 2), stage.getHeight() - (shape / 2)
                )))

                layer.add(new Konva.Rect(createRect(
                    stage.getWidth() / 2 - (shape / 2), stage.getHeight() - (shape / 2),
                    shape, stage.getHeight() - (shape / 2)
                )))
            },
            createPinHole(shape) {
                rect = new Konva.Rect({
                    x: stage.getWidth() / 2 - (shape / 2),
                    y: stage.getHeight() / 2 - (shape / 2),
                    width: shape,
                    height: shape,
                    fill: 'transparent',
                    stroke: 'black',
                    strokeWidth: 1,
                    listening: false,
                })
                layer.add(rect)
            }
        },
        mounted() {
            stage = new Konva.Stage({
                container: 'container',
                width: document.getElementById("container").offsetWidth,
                height: document.getElementById("container").offsetHeight
            })
            layer = new Konva.Layer()

            const laserWidth = 420
            this.createPinHole(laserWidth)
            this.createBlurEffect(laserWidth)

            stage.add(layer)
        }
    }
</script>
<style lang="scss">
    #container {
        width: 100%;
        height: calc(100vh - 105px);
    }

    .file.is-pulled-right.is-success {
        z-index: 10;
    }
</style>
