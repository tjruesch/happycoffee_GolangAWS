<template>
    <login-layout>
        <form class="form-signin" v-on:submit.prevent="submit">
            <span style="font-size: 60px;">	&#9749;</span>
            <h1 class="h3 mb-3 font-weight-normal">Please sign in</h1>

            <label for="inputEmail" class="sr-only">Email address</label>
            <input type="email" id="inputEmail" class="form-control" v-model="login" placeholder="Email address"
                   required autofocus>
            <label for="inputPassword" class="sr-only">Password</label>
            <input type="password" id="inputPassword" class="form-control" v-model="password" placeholder="Password"
                   required>
            <div class="checkbox mb-3">
                <label>
                    <input type="checkbox" value="remember-me"> Remember me
                </label>
            </div>
            <button class="btn btn-lg btn-primary btn-block" type="submit">
                Sign in
                <span v-if="showLoader" class="spinner-grow spinner-grow-sm" style="width: 1.3rem; height: 1.3rem;"
                      role="status" aria-hidden="true"></span>
            </button>
            <p class="mt-5 mb-3 text-muted">&copy; 2021 happycoffee</p>
        </form>
    </login-layout>
</template>

<script>
    import LoginLayout from '../layouts/Login'

    export default {
        name: "Login",
        components: {
            LoginLayout,
        },
        mounted() {
        },
        methods: {
            submit: function () {
                // let self = this
                // this.showLoader = true

                // loginUser(this.login, this.password)
                //     .then(function () {
                //         self.$toast.info("Hey buddy!")
                //         self.$router.push({name: 'trainingsList'})
                //     })
                //     .catch(error => {
                //         self.$toast.error("Failed to log in")
                //         console.error(error)
                //         self.showLoader = false
                //     })
            },
        },
        data: function () {
            return {
                'login': '',
                'password': '',
                'showLoader': false,
            }
        }
    }
</script>

<style scoped>
    .form-signin {
        width: 100%;
        width: 330px;
        padding: 15px;
        margin: 0 auto;
    }

    .form-signin .checkbox {
        font-weight: 400;
    }

    .form-signin .form-control {
        position: relative;
        box-sizing: border-box;
        height: auto;
        padding: 10px;
        font-size: 16px;
    }

    .form-signin .form-control:focus {
        z-index: 2;
    }

    .form-signin input[type="email"] {
        margin-bottom: -1px;
        border-bottom-right-radius: 0;
        border-bottom-left-radius: 0;
    }

    .form-signin input[type="password"] {
        margin-bottom: 10px;
        border-top-left-radius: 0;
        border-top-right-radius: 0;
    }
</style>