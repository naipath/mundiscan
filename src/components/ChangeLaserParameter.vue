<template>
    <Modal id="app" :is-active="isActive">
        <div class="modal-card">
            <header class="modal-card-head">
                <p class="modal-card-title">Laser parameters aanpassen</p>
                <button class="delete" aria-label="close" @click="onClose()"></button>
            </header>
            <form @submit.prevent="onClose()">
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
                    <button class="button is-danger" type="button" @click="onClose">Annuleren</button>
                    <button class="button is-success" type="submit">Wijzigen</button>
                </footer>
            </form>
        </div>
        </div>
    </Modal>
</template>
<script>
import Modal from "./Modal.vue"
export default {
    name: 'app',
    components: {
        Modal
    },
    props: {
        isActive: Boolean,
        onClose: Function,
    }
}
</script>