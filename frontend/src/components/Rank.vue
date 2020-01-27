<template>
  <v-container fluid>
    Posición
    <v-card tile color="indigo lighten-5">
      <v-expansion-panels accordion>
        <v-expansion-panel v-for="(user, i) in users" :key="i" hide-actions>
          <v-expansion-panel-header>
            <v-row align="center" class="spacer" no-gutters>
              <v-col cols="4" sm="2" md="1">
                <v-avatar color="indigo lighten-1 white--text" size="25">{{i+1}}</v-avatar>
              </v-col>
              <v-col>{{user.nickname}} con {{user.score}} puntos falló las siguientes preguntas</v-col>
            </v-row>
          </v-expansion-panel-header>
          <v-expansion-panel-content>
            <v-divider class="light-blue darken-4"></v-divider>
            <v-list-item v-for="(q,i) in user.failed" :key="i">
                <v-list-item-content>
                  <v-list-item-title>{{q.statement}}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            <!-- <v-row v-for="(user, i) in users" :key="i">
              
            </v-row> -->
            <v-icon v-if="user.sucks" :color="user.color">mdi-thumb-up</v-icon>
            <v-icon v-else :color="user.color">mdi-thumb-down</v-icon>
          </v-expansion-panel-content>
        </v-expansion-panel>
      </v-expansion-panels>
    </v-card>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  name: "Rank",
  data: () => ({
    users: [],    
  }),
  methods: {
    usersList() {
      const path = "http://localhost:3000/api/users";
      axios.get(path).then(response => {
        this.users = response.data.sort((a, b) => {
          if (a.score < b.score) return 1;
          else if (a.score === b.score) return 0;
          else return -1;
        });
      this.users.map(user => {
        user.sucks = user.failed.length < 10
        user.color = user.failed.length < 10 ? "blue" : "red"
      });
      });
      // return this.users
    }
  },
  created() {
    this.usersList();
  }
};
</script>