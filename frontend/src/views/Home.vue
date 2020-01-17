<template>
    <div class="home">
        <v-spacer></v-spacer>
        <v-spacer></v-spacer>
        <v-container class="ma-lg-auto">
            <v-row>
                <v-card tile color="indigo lighten-5" class="ma-lg-auto">
                    <v-col cols="12" xs="12" md="6" lg="12">
                        <p class="black--text">Ser un guerrero no es una simple cuestión de desear ser uno. Es más bien
                            una lucha interminable
                            que continuará hasta el último momento de nuestras vidas. Nadie nace como guerrero,
                            exactamente
                            de la misma manera que nadie nace como un hombre promedio. Nos convertimos en uno u otro.
                        </p>
                        <p class="text-right black--text">—Kokoro por Natsume Sosek, 1914, Japan</p>
                    </v-col>
                    <v-col cols="12">
                        <Login/>
                    </v-col>
                </v-card>
            </v-row>

        </v-container>

    </div>
</template>

<script>
    import Login from "../components/Login";

    export default {
        name: 'home',
        components: {Login},

    }
</script>
