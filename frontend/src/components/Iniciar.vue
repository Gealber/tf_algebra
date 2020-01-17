<template>
    <v-container class="mt-12">
        <form action="">
            <v-row justify="center">
                <v-col cols="12" xs="12" md="6" lg="6">
                    <v-text-field
                            v-model="nickname"
                            :rules="[rules.required]"
                            label="Alias"
                    ></v-text-field>
                </v-col>
            </v-row>
            <v-row justify="center">
                <v-col cols="12" xs="12" md="6" lg="6">
                    <v-text-field
                            v-model="password"
                            :append-icon="show1 ? 'mdi-eye' : 'mdi-eye-off'"
                            :rules="[rules.required, rules.min]"
                            :type="show1 ? 'text' : 'password'"
                            name="input-10-1"
                            label="Contraseña"
                            hint="Al menos 8 caracteres"
                            counter
                            @click:append="show1 = !show1"
                    ></v-text-field>
                </v-col>
            </v-row>
            <v-row justify="center">
                <v-btn rounded class="mr-4 indigo" @click="submit">Iniciar</v-btn>
                <v-btn rounded class="red" @click="clear">Borrar</v-btn>
            </v-row>
            <!--            Alert pop up-->
        </form>
        <v-row align="center">
            <v-col cols="12" xs="10" lg="5">
                <div>
                    <v-alert
                            v-model="alert2"
                            dismissible
                            type="error"
                            border="left"
                            elevation="2"
                            colored-border
                    >
                        Ups algo salío bien mal
                    </v-alert>
                </div>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
    import axios from 'axios'

    export default {
        name: "Iniciar",
        data: () => ({
            nickname:'',
            password:'',
            show1: false,
            show2: false,
            alert2: false,
            errorMessages: {
                'name': 'El nombre no debe de pasarse de más de 10 caracteres',
            },
            rules: {
                required: value => !!value || 'Required',
                min: v => v.length >= 8 || 'Min 8 caracteres'
            }
        }),
        methods: {
            submit() {
                const path = 'http://localhost:3000/api/login';
                let data = {
                    nickname:this.nickname,
                    password:this.password
                };
                const headers = {
                    'Content-Type': 'application/x-www-form-urlencoded',
                };
                axios.post(path, data, {
                    headers: headers
                })
                    .then(response => {
                        if (response.status === 200) {
                            if (response.data.token.length > 0) {
                                localStorage.setItem('userAlias', data.nickname);
                                localStorage.setItem('tokenData',response.data.token);
                                this.$router.push("dashboard");   //
                            }
                        }
                    })
                    .catch(() => {
                        this.alert2 = true
                    })
            },
            clear() {
                this.name = '';
                this.nick = '';
                this.password = ''
            },
            errors(data) {
                let nameLength = this.name.length > 10;
                if (nameLength) {
                    return this.errorMessages[data]
                }
            }
        }
    }
</script>

<style scoped>

</style>