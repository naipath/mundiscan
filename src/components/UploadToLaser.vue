<template>
    <div id="app">

        <div class="file is-pulled-right is-success option">
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

        <a class="is-pulled-right option button is-success" @click="uploadToLaser">Upload naar laser</a>

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

                if (imageObj.width > 410) {
                    imageObj.height = imageObj.height * 410 / imageObj.width
                    imageObj.width = 410
                }


                if (imageObj.height > 410) {
                    imageObj.width = imageObj.width * 410 / imageObj.height
                    imageObj.height = 410
                }

                const xStartLocation = stage.getWidth() / 2 - 420
                const yStartLocation = stage.getHeight() / 2 - 420
                let logo = new Konva.Image({
                    x: xStartLocation,
                    y: yStartLocation,
                    image: imageObj,
                    stroke: 'blue',
                    strokeWidth: 1,
                    strokeEnabled: false
                })

                let logoGroup = new Konva.Group({
                    x: 180,
                    y: 50,
                    draggable: true
                });
                logoGroup.add(logo);
                layer.add(logoGroup);

                logo.moveToBottom()
                logoGroup.moveToBottom()

                let anchors = [
                    addAnchor(logoGroup, xStartLocation, yStartLocation, 'topLeft'),
                    addAnchor(logoGroup, xStartLocation + imageObj.width, yStartLocation, 'topRight'),
                    addAnchor(logoGroup, xStartLocation + imageObj.width, yStartLocation + imageObj.height, 'bottomRight'),
                    addAnchor(logoGroup, xStartLocation, yStartLocation + imageObj.height, 'bottomLeft'),
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
            },
            uploadToLaser() {

            },
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

    .option {
        z-index: 10;
    }
</style>
