<template>
    <div>
        <WebMenu/>
        <div class="container">
            <div>
                <slot></slot>
            </div>
            <WebFooter/>
        </div>
    </div>
</template>

<script>
    import WebFooter from "../components/WebFooter.vue";
    import WebMenu from "../components/WebMenu.vue";

    export default {
        name: 'Main',
        components: {
            WebFooter,
            WebMenu,
        },
        beforeCreate: function () {
            document.body.className = 'app';
        },
    }
</script>

<style scoped>
.container {
    min-height: calc(100vh - 70px);  /* Height of Menu is 70px */
    display: flex;
    flex-direction: column;
    justify-content: space-between;
}
</style>