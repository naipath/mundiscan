<template>
    <nav class="navbar" id="app">

        <div class="navbar-brand">
            <a class="navbar-item"
               v-bind:class="{'is-active': activeRoute === 'about'}"
               href="#"
               @click="routeChanged('about')">
                <img src="https://www.mundi-technology.nl/files/_thumbnail/2/logo_1500x1500.png"
                     width="112"
                     height="28">
            </a>
        </div>

        <div class="navbar-menu">

            <div class="navbar-start" v-for="route in routes">
                <a class="navbar-item" href="#" @click="routeChanged('laser/' + route)">{{route}}</a>
            </div>

            <div class="navbar-end">
                <a class="navbar-item"
                   v-bind:class="{'is-active': activeRoute === 'add-laser'}"
                   href="#"
                   @click="routeChanged('add-laser')">
                    <i class="fa fa-plus-square"></i>&nbsp;&nbsp;&nbsp;Voeg toe
                </a>
            </div>
        </div>
    </nav>
</template>
<script>
    export default {
        name: "app",
        props: {
            routes: Array,
            routeChanged: Function,
            activeRoute: String,
        }
    }
</script>
<style lang="scss">
    .navbar {
        border-bottom: 1px solid #eee;
        margin-bottom: 50px;
    }

    .navbar-item.is-active {
        background-color: #f5f5f5 !important;
    }
</style>