<template>
    <div class="d-flex flex-column flex-md-row align-items-center p-3 px-md-4 mb-3 bg-white border-bottom shadow-sm">
        <h5 class="my-0 mr-md-auto font-weight-normal">
            happycoffee - Coffee and Tea Specialities
        </h5>
        <a class="btn btn-outline-primary" v-if="isAuthenticated()" href="/login">Logout</a>
        <a class="btn btn-outline-primary" v-else href="/login">Login</a>
    </div>
</template>

<script>
    export default {
        name: "WebMenu",
        methods: {
            isAuthenticated: function() {
                return false
            }
        }
    }
</script>

<style scoped>
.no-underline:hover {
    text-decoration: none;
}
</style>