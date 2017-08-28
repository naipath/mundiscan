<template>
    <div id="app">

        <nav class="panel" style="position: absolute; z-index: 10; right: 20px; background: white; top: 70px">
            <p class="panel-heading">
                Acties
            </p>
            <a class="panel-block" @click="showTextModal = true">
                <span class="panel-icon">
                    <i class="fa fa-font"></i>
                </span>
                Tekst
            </a>
            <label class="panel-block file-label">
                <input class="file-input" type="file" accept="image/png,image/jpg,image/jpeg" v-on:change="chooseImage">

                <span class="panel-icon">
                    <i class="fa fa-image"></i>
                </span>
                Afbeelding
            </label>
            <a class="panel-block" @click="showLaserParameters = true">
                <span class="panel-icon">
                    <i class="fa fa-barcode"></i>
                </span>
                Parameters
            </a>
            <a class="panel-block" @click="initialize">
                <span class="panel-icon">
                    <i class="fa fa-refresh"></i>
                </span>
                Reset
            </a>
            <div class="panel-block">
                <button class="button is-success is-outlined is-fullwidth" @click="uploadToLaser">
                    Upload
                </button>
            </div>
        </nav>

        <div id="container"></div>

        <a id="download" href="#" download="printer.png" style="display: none;"></a>

        <div class="modal" v-bind:class="{'is-active': showTextModal}">
            <div class="modal-background" @click="showTextModal = false"></div>
            <div class="modal-card">
                <header class="modal-card-head">
                    <p class="modal-card-title">Tekst toevoegen</p>
                    <button class="delete" aria-label="close" @click="showTextModal = false"></button>
                </header>
                <form @submit.prevent="addText">
                    <section class="modal-card-body">

                        <div class="field">
                            <div class="control">
                                <input class="input" placeholder="Tekst" v-model="textInput" required>
                            </div>
                        </div>

                        <div class="field">
                            <label class="label">Tekst grootte: {{textSize}}</label>
                            <div class="control">
                                <input type="range" min="2" max="100" step="1" v-model="textSize" style="width: 100%">
                            </div>
                        </div>

                    </section>
                    <footer class="modal-card-foot">
                        <button class="button is-success" type="submit">Toevoegen</button>
                    </footer>
                </form>
            </div>
        </div>

        <div class="modal" v-bind:class="{'is-active': showLaserParameters}">
            <div class="modal-background" @click="showLaserParameters = false"></div>
            <div class="modal-card">
                <header class="modal-card-head">
                    <p class="modal-card-title">Laser parameters aanpassen</p>
                    <button class="delete" aria-label="close" @click="showLaserParameters = false"></button>
                </header>
                <form @submit.prevent="showLaserParameters = false">
                    <section class="modal-card-body">
                        <div class="field">
                            <label class="label">Frequency 500-50000hz</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">Duty 1-85%</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">Mark speed 10-20.000mm/s</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">Jump speed 10-20.000mm/s</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">Off delay 2-8000 ûseconds</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>
                        <div class="field">
                            <label class="label">On delay -8000-8000üseconds</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>

                        <div class="field">
                            <label class="label">Jump delay 0-32000ûseconds</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>

                        <div class="field">
                            <label class="label">Mark delay 0-32000ûseconds</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>

                        <div class="field">
                            <label class="label">Polygon delay 0-32000useconds</label>
                            <div class="control">
                                <input class="input" required>
                            </div>
                        </div>
                    </section>
                    <footer class="modal-card-foot">
                        <button class="button is-danger" type="submit">Annuleren</button>
                        <button class="button is-success" type="submit">Wijzigen</button>
                    </footer>
                </form>
            </div>
        </div>

        <div class="modal" v-bind:class="{'is-active': isUploading || uploadError}">
            <div class="modal-background"></div>
            <div class="modal-content hero is-light has-text-centered">
                <div class="hero-body">
                    <h1 class="title">Uploaden</h1>
                    <h1 class="title button is-light is-loading" v-if="!uploadError"></h1>

                    <div class="error-message notification is-danger" v-if="uploadError">
                        Er is iets misgegaan bij het uploaden naar de laser
                        <br/>
                        <button class="button" @click="uploadToLaser">Opnieuw proberen</button>
                    </div>

                </div>
            </div>
        </div>

    </div>
</template>

<script>
import { dataURItoBlob } from "./../helpers";

let stage, layer, rect, backgroundRect;
let rasterLines = []
const laserShape = 420

export default {
    name: 'app',
    props: {
        laser: Object
    },
    data() {
        return {
            showTextModal: false,
            textInput: null,
            textSize: 30,
            showLaserParameters: false,
            isUploading: false,
            uploadError: true,
        }
    },
    methods: {
        chooseImage(event) {
            if (window.File && window.FileReader && window.FileList && window.Blob) {
                const file = event.target.files[0]
                const fileReader = new FileReader()
                fileReader.onload = e => {
                    let imageObj = new Image()
                    imageObj.onload = () => this.addLogo(imageObj)
                    imageObj.src = e.target.result
                }
                fileReader.readAsDataURL(file)
            }
        },
        addText() {
            this.showTextModal = false
            const textNode = new Konva.Text({
                x: 0,
                y: 0,
                fontSize: this.textSize,
                text: this.textInput,
            })

            let textGroup = new Konva.Group({
                x: rect.x(),
                y: rect.y(),
                draggable: true,

            });
            textGroup.add(textNode);

            let rectNode = new Konva.Rect({
                x: 0,
                y: 0,
                width: textNode.getWidth(),
                height: textNode.getHeight(),
                strokeEnabled: false,
                stroke: 'blue',
                strokeWidth: 1,
            })

            textGroup.add(rectNode);

            textGroup.on('mouseover', () => {
                document.body.style.cursor = 'pointer'
                rectNode.strokeEnabled(true)
                layer.draw()
            })
            textGroup.on('mouseout', () => {
                document.body.style.cursor = 'default'
                rectNode.strokeEnabled(false)
                layer.draw()
            })
            textGroup.on('dblclick', () => {
                textGroup.destroy()
                layer.draw()
            })

            layer.add(textGroup)
            textGroup.moveToBottom()
            backgroundRect.moveToBottom()
            layer.draw()
        },
        addLogo(imageObj) {
            const update = activeAnchor => {
                let group = activeAnchor.getParent()
                let topLeft = group.get('.topLeft')[0]
                let topRight = group.get('.topRight')[0]
                let bottomRight = group.get('.bottomRight')[0]
                let bottomLeft = group.get('.bottomLeft')[0]
                let image = group.get('Image')[0]
                let anchorX = activeAnchor.getX()
                let anchorY = activeAnchor.getY()
                // update anchor positions
                switch (activeAnchor.getName()) {
                    case 'topLeft':
                        topRight.setY(anchorY)
                        bottomLeft.setX(anchorX)
                        break
                    case 'topRight':
                        topLeft.setY(anchorY)
                        bottomRight.setX(anchorX)
                        break
                    case 'bottomRight':
                        bottomLeft.setY(anchorY)
                        topRight.setX(anchorX)
                        break
                    case 'bottomLeft':
                        bottomRight.setY(anchorY)
                        topLeft.setX(anchorX)
                        break
                }
                image.position(topLeft.position())
                let width = topRight.getX() - topLeft.getX()
                let height = bottomLeft.getY() - topLeft.getY()
                if (width && height) {
                    image.width(width)
                    image.height(height)
                }
            }

            const addAnchor = (group, x, y, name) => {
                let layer = group.getLayer()
                let anchor = new Konva.Circle({
                    x, y,
                    stroke: '#666',
                    fill: '#ddd',
                    strokeWidth: 1,
                    radius: 6,
                    name: name,
                    visible: false,
                    draggable: true,
                    dragOnTop: false
                })
                anchor.on('dragmove', () => {
                    update(anchor)
                    layer.draw()
                })
                anchor.on('mousedown touchstart', () => group.setDraggable(false))

                anchor.on('dragend', () => {
                    group.setDraggable(true)
                    layer.draw()
                })

                anchor.on('mouseover', () => {
                    let layer = anchor.getLayer()
                    document.body.style.cursor = 'pointer'
                    anchor.setStrokeWidth(3)
                    layer.draw()
                })
                anchor.on('mouseout', () => {
                    let layer = anchor.getLayer()
                    document.body.style.cursor = 'default'
                    anchor.setStrokeWidth(1)
                    layer.draw()
                })
                group.add(anchor)
                return anchor
            }

            if (imageObj.width > laserShape + 10) {
                imageObj.height = imageObj.height * laserShape / imageObj.width
                imageObj.width = laserShape
            }


            if (imageObj.height > laserShape + 10) {
                imageObj.width = imageObj.width * laserShape / imageObj.height
                imageObj.height = laserShape
            }

            const xStartLocation = rect.x()
            const yStartLocation = rect.y()

            let logo = new Konva.Image({
                x: 0,
                y: 0,
                image: imageObj,
                stroke: 'blue',
                strokeWidth: 1,
                strokeEnabled: false
            })

            let logoGroup = new Konva.Group({
                x: xStartLocation,
                y: yStartLocation,
                draggable: true
            })
            logoGroup.add(logo)
            layer.add(logoGroup)

            logoGroup.moveToBottom()
            backgroundRect.moveToBottom()

            let anchors = [
                addAnchor(logoGroup, 0, 0, 'topLeft'),
                addAnchor(logoGroup, imageObj.width, 0, 'topRight'),
                addAnchor(logoGroup, imageObj.width, imageObj.height, 'bottomRight'),
                addAnchor(logoGroup, 0, imageObj.height, 'bottomLeft'),
            ]

            logoGroup.on('dblclick', () => {
                logoGroup.destroy()
                layer.draw()
            })

            logoGroup.on('mouseover', () => {
                logo.strokeEnabled(true)
                anchors.forEach(anchor => anchor.show())
                layer.draw()
            })

            logoGroup.on('mouseout', () => {
                logo.strokeEnabled(false)
                anchors.forEach(anchor => anchor.hide())
                layer.draw()
            })

            layer.draw()
        },
        createBlurEffect(shape) {
            const createRect = (x, y, width, height) => ({
                x, y,
                width, height,
                fill: 'rgba(255, 255, 255, 0.5)',
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
                stage.getWidth() / 2 - (shape / 2), stage.getHeight()
            )))

            layer.add(new Konva.Rect(createRect(
                stage.getWidth() / 2 - (shape / 2), stage.getHeight() / 2 + (shape / 2),
                shape, stage.getHeight()
            )))
        },
        createPinHole(shape) {
            rect = new Konva.Rect({
                x: stage.getWidth() / 2 - (shape / 2) - 1,
                y: stage.getHeight() / 2 - (shape / 2) - 1,
                width: shape + 2,
                height: shape + 2,
                fill: 'transparent',
                stroke: 'black',
                strokeWidth: 1,
                listening: false,
            })
            layer.add(rect)

            backgroundRect = new Konva.Rect({
                x: stage.getWidth() / 2 - (shape / 2) - 1,
                y: stage.getHeight() / 2 - (shape / 2) - 1,
                width: shape + 2,
                height: shape + 2,
                fill: 'white',
                strokeEnabled: 'false',
                listening: false,
            })

            layer.add(backgroundRect)
            backgroundRect.moveToBottom()
        },
        createRaster(shape) {
            Array.from(Array(9).keys())
                .forEach(i => {
                    let verticalRasterLine = new Konva.Line({
                        points: [
                            stage.getWidth() / 2 - (shape / 2) + (i + 1) * (shape / 10),
                            stage.getHeight() / 2 - (shape / 2),

                            stage.getWidth() / 2 - (shape / 2) + (i + 1) * (shape / 10),
                            stage.getHeight() / 2 - (shape / 2) + shape
                        ],
                        stroke: 'rgba(140, 140, 140, 0.4)',
                        strokeWidth: 1,
                        listening: false,
                    })
                    layer.add(verticalRasterLine)
                    rasterLines.push(verticalRasterLine)

                    let horizontalRasterLine = new Konva.Line({
                        points: [
                            stage.getWidth() / 2 - (shape / 2),
                            stage.getHeight() / 2 - (shape / 2) + (i + 1) * (shape / 10),

                            stage.getWidth() / 2 - (shape / 2) + shape,
                            stage.getHeight() / 2 - (shape / 2) + (i + 1) * (shape / 10)
                        ],
                        stroke: 'rgba(140, 140, 140, 0.4)',
                        strokeWidth: 1,
                        listening: false,
                    })
                    layer.add(horizontalRasterLine)
                    rasterLines.push(horizontalRasterLine)
                })

        },
        uploadToLaser() {
            this.uploadError = false
            this.isUploading = true
            rasterLines.forEach(line => line.hide())
            layer.draw()

            const croppedStage = stage.toCanvas({
                width: stage.getWidth(),
                height: stage.getHeight(),
            })

            let hiddenCanvas = document.createElement('canvas')
            hiddenCanvas.style.display = 'none'
            document.body.appendChild(hiddenCanvas)
            hiddenCanvas.width = laserShape
            hiddenCanvas.height = laserShape

            let hiddenContext = hiddenCanvas.getContext('2d')
            hiddenContext.drawImage(
                croppedStage,
                stage.getWidth() / 2 - laserShape / 2,
                stage.getHeight() / 2 - laserShape / 2,
                laserShape, laserShape,
                0, 0,
                hiddenCanvas.width, hiddenCanvas.height
            )
            document.getElementById("download").href = hiddenCanvas.toDataURL("image/png")
            document.getElementById("download").click()

            const data = new FormData()
            data.append('uploadfile', dataURItoBlob(hiddenCanvas.toDataURL("image/png")), 'mundiscan-' + Date.now() + '.png')
            fetch('/laserclients/' + this.laser.Id + '/upload', {
                method: 'POST',
                body: data,
            })
                .then(result => {
                    if (result.ok) {
                        this.isUploading = false
                        this.uploadError = false
                    } else {
                        this.uploadError = true
                        this.isUploading = false

                    }
                })

            rasterLines.forEach(line => line.show())
            layer.draw()
        },
        initialize() {
            stage = new Konva.Stage({
                container: 'container',
                width: document.getElementById("container").offsetWidth,
                height: document.getElementById("container").offsetHeight
            })
            layer = new Konva.Layer()

            this.createPinHole(laserShape)
            this.createBlurEffect(laserShape)
            this.createRaster(laserShape)

            stage.add(layer)
        }
    },
    mounted() {
        this.initialize()
    }
}
</script>
<style lang="scss">
#container {
    width: 100%;
    height: calc(100vh - 105px);
}

.option {
    z-index: 10;
}

.modal-card-foot {
    justify-content: flex-end;
}

.panel {
    position: absolute;
    z-index: 10;
    right: 20px;
    background: white;
    top: 70px
}
</style>
