<template>
    <v-container fluid>
        <div class="d-block pa-2 indigo darken-2 white--text">Selecione los planteamientos que considere
            verdaderos.
        </div>
        <v-card tile color="indigo lighten-5">
            <v-col
                    v-for="(question, i) in questions"
                    :key="i"
            >
                <v-divider v-if="i != 0" class="indigo darken-2"></v-divider>
                <v-checkbox
                        :label="question.statement"
                        color="deep-purple darken-2"
                        v-model="question.selected"
                        class="black--text"
                ></v-checkbox>
            </v-col>
        </v-card>
        <v-row justify-lg="center" class="ma-5 pa-5">
            <v-btn class="indigo darken-2 white--text font-weight-light" @click="userScore" href="rank">Enviar</v-btn>
        </v-row>
        <v-row align="center">            
                <div>
                    <v-alert
                            v-model="alert"
                            dismissible
                            color="cyan"
                            border="left"
                            elevation="2"
                            colored-border
                            icon="mdi-firework"
                    >
                        Su puntuacion fue de {{this.result}} / {{this.questions.length}}
                    </v-alert>
                </div>
        </v-row>
    </v-container>

</template>

<script>
    import axios from 'axios'

    export default {
        name: "QuestionsList",
        data: () => ({
            questions: [],
            questionSubjects: [
                "Sistema de ecuaciones",
                "Espacios vectoriales",
                "Dependencia lineal",
                "Base, subespacio generador y dimensión",
                "Aplicaciones lineales"
            ],
            result: 0,
            failed: [],
            alert: false,
            hide: false,
        }),
        methods: {
            questionList() {
                const path = "http://localhost:3000/api/questions";

                const options = {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('tokenData')}`,
                    }
                };
                axios.get(path, options)
                    .then(response => {
                        this.questions = response.data;
                    });
                return this.questions

            },
            userScore() {
                const path = "http://localhost:3000/api/users/score";
                const options = {
                    headers: {
                        'Authorization': `Bearer ${localStorage.getItem('tokenData')}`,
                    }
                };
                if (this.questions.length > 0) {
                    this.result = this.questions.filter(question => question.selected.toString() === question.result).length;
                    this.failed = this.questions.filter(question => question.selected.toString() !== question.result);
                    this.alert = true;
                    let data = {
                        nickname: localStorage.getItem('userAlias'),
                        score: this.result,
                        failed: this.failed
                    };
                    axios.post(path, data, options)
                }
            }
        },
        created() {
            this.questionList();
        },
    }
</script>

<style scoped>
    .black--text /deep/ label {
        color: black
    }
</style>