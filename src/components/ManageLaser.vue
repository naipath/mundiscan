<template>
    <div id="app">
        {{laser.Name}}
        {{laser.Ip}}
        {{laser.Port}}
    </div>
</template>
<script>
    export default {
        name: "app",
        props: {
            laser: Object
        }
    }
</script>