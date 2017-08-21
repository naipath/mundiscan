<template>
    <div id="app">

        <nav class="panel" style="position: absolute; z-index: 10000; right: 0;">
            <p class="panel-heading">
                Acties
            </p>
            <a class="panel-block" @click="addText">
                <span class="panel-icon">
                  <i class="fa fa-font"></i>
                </span>
                Tekst
            </a>
            <a class="panel-block" @click="initialize">
                <span class="panel-icon">
                  <i class="fa fa-refresh"></i>
                </span>
                Reset
            </a>
            <label class="panel-block file-label">
                <input class="file-input"
                       type="file"
                       accept="image/png,image/jpg,image/jpeg"
                       v-on:change="chooseImage">

                <span class="panel-icon">
                    <i class="fa fa-upload"></i>
                </span>
                Afbeelding
            </label>
            <a class="panel-block" @click="uploadToLaser">
                <span class="panel-icon">
                  <i class="fa fa-cloud-upload"></i>
                </span>
                Upload
            </a>
        </nav>

        <div id="container"></div>

        <a id="download" href="#" download="printer.png" style="display: none;"></a>

    </div>
</template>

<script>
    import {dataURItoBlob} from "./../helpers";

    let stage, layer, rect;
    const laserShape = 420

    export default {
        name: 'app',
        props: {
            laser: Object
        },
        data() {
            return {}
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
                const newText = new Konva.Text({
                    x: rect.x(),
                    y: rect.y(),
                    fontSize: 30,
                    text: "testing...",
                    draggable: true,
                })
                newText.on('mouseover', () => {
                    document.body.style.cursor = 'pointer'
                    layer.draw()
                })
                newText.on('mouseout', () => {
                    document.body.style.cursor = 'default'
                    layer.draw()
                })
                newText.on('dblclick', () => {
                    newText.destroy()
                    layer.draw()
                })
                layer.add(newText)
                newText.moveToBottom()
                layer.draw()
            },
            addLogo(imageObj) {

                const update = activeAnchor => {
                    let group = activeAnchor.getParent();
                    let topLeft = group.get('.topLeft')[0];
                    let topRight = group.get('.topRight')[0];
                    let bottomRight = group.get('.bottomRight')[0];
                    let bottomLeft = group.get('.bottomLeft')[0];
                    let image = group.get('Image')[0];
                    let anchorX = activeAnchor.getX();
                    let anchorY = activeAnchor.getY();
                    // update anchor positions
                    switch (activeAnchor.getName()) {
                        case 'topLeft':
                            topRight.setY(anchorY);
                            bottomLeft.setX(anchorX);
                            break;
                        case 'topRight':
                            topLeft.setY(anchorY);
                            bottomRight.setX(anchorX);
                            break;
                        case 'bottomRight':
                            bottomLeft.setY(anchorY);
                            topRight.setX(anchorX);
                            break;
                        case 'bottomLeft':
                            bottomRight.setY(anchorY);
                            topLeft.setX(anchorX);
                            break;
                    }
                    image.position(topLeft.position());
                    let width = topRight.getX() - topLeft.getX();
                    let height = bottomLeft.getY() - topLeft.getY();
                    if (width && height) {
                        image.width(width);
                        image.height(height);
                    }
                }

                const addAnchor = (group, x, y, name) => {
                    let layer = group.getLayer();
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
                    });
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
                });
                logoGroup.add(logo);
                layer.add(logoGroup);

                logo.moveToBottom()
                logoGroup.moveToBottom()

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
            },
            uploadToLaser() {
                const croppedStage = stage.toCanvas({
                    width: stage.getWidth(),
                    height: stage.getHeight(),
                })

                let hiddenCanvas = document.createElement('canvas');
                hiddenCanvas.style.display = 'none';
                document.body.appendChild(hiddenCanvas);
                hiddenCanvas.width = stage.getWidth();
                hiddenCanvas.height = stage.getHeight();

                let hiddenContext = hiddenCanvas.getContext('2d');
                hiddenContext.drawImage(
                    croppedStage,
                    stage.getWidth() / 2 - laserShape / 2,
                    stage.getHeight() / 2 - laserShape / 2,
                    laserShape, laserShape,
                    0, 0,
                    hiddenCanvas.width, hiddenCanvas.height
                );
                document.getElementById("download").href = hiddenCanvas.toDataURL("image/png")
                document.getElementById("download").click();
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
</style>
