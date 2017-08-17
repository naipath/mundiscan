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

            <div class="navbar-start">
                <a class="navbar-item" v-for="route in routes"
                   href="#"
                   @click="routeChanged('manage-laser/' + route)"
                   v-bind:class="{'is-active': activeRoute === 'manage-laser/' + route}">{{route}}</a>

                <a class="navbar-item"
                   v-bind:class="{'is-active': activeRoute === 'add-laser'}"
                   href="#"
                   @click="routeChanged('add-laser')">
                    <i class="fa fa-plus-square"></i>&nbsp;&nbsp;&nbsp;Toevoegen
                </a>
            </div>

            <div class="navbar-end">
                <a class="navbar-item"
                   v-bind:class="{ 'is-active': activeRoute.indexOf('/status') >= 0 }"
                   href="#"
                   @click.prevent="routeChanged('manage-laser/' + activeRoute + '/status')">
                    <i class="fa fa-stethoscope"></i>&nbsp;Status
                </a>

                <a class="navbar-item"
                   v-bind:class="{ ['is-active']: activeRoute.indexOf('/settings') >= 0  }"
                   href="#"
                   @click.prevent="routeChanged('manage-laser/' + activeRoute + '/settings')">
                    <i class="fa fa-cog"></i>&nbsp;Instellingen
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
        font-weight: bold;
    }
</style>