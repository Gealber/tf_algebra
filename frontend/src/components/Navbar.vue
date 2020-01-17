<template>
    <nav>
        <v-app-bar app class="indigo darken-1">
            <v-toolbar-title>
                <span class="font-weight-light white--text">Become a Ninja in </span>
                <span class="white--text">Algebra</span>
            </v-toolbar-title>
            <v-avatar>
                <img
                    :class="loginActive()"
                    src="https://i.ya-webdesign.com/images/ninja-head-png.png"
                    alt="Ninja"
                >                
            </v-avatar>
            <v-spacer></v-spacer>
            <v-btn depressed :class="loginActive()" href="login">
                <span class="font-weight-light white--text">Login</span>
                <v-icon color="white">face</v-icon>
            </v-btn>
            <v-btn depressed :class="exitActive()" @click="exit">
                <span class="font-weight-light white--text">Salir</span>
                <v-icon color="white">exit_to_app</v-icon>
            </v-btn>
        </v-app-bar>
    </nav>
</template>

<script>
    import router from "../router";
    export default {
        name: "Navbar",
        data: () => ({
            routeName: '',
        }),
        methods: {
            loginActive() {
                if (this.$route.name !== "home") {
                    return "hide"
                } else {
                    return "indigo darken-1"
                }
            },
            exitActive() {
                if (["dashboard","rank"].includes(this.$route.name)) {
                    return "indigo darken-1"
                } else {
                    return "hide"
                }
            },
            exit() {
                localStorage.clear();
                router.push("/");
            }
        }
    }
</script>

<style scoped>
    .hide {
        display: none;
    }
</style>