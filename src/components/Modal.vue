<template>
    <transition id="app" name="fade">
        <div class="modal is-active" v-if="isActive">
            <div class="modal-background"></div>
            <div class="modal-content">
                <slot></slot>
            </div>
        </div>
    </transition>
</template>
<script>
export default {
    name: 'app',
    props: {
        isActive: Boolean,
    }
}
</script>
<style>
.fade-enter-active,
.fade-leave-active {
    transition: opacity .5s
}

.fade-enter,
.fade-leave-to {
    opacity: 0
}
</style>