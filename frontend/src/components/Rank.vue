<template>
  <v-container fluid>
      Ranking
    <v-card tile color="indigo lighten-5">
      <v-col v-for="(user, i) in users" :key="i">
        <!-- <v-divider v-if="i != 0" class="indigo darken-2"></v-divider> -->
        <v-avatar color="indigo lighten-1" size="25">
          {{i+1}}
        </v-avatar>
        {{user.nickname}}  {{user.score}}
      </v-col>
    </v-card>
  </v-container>
</template>

<script>
import axios from "axios";

export default {
  name: "Rank",
  data: () => ({
    users: []
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
      });
      // return this.users
    }
  },
  created() {
    this.usersList();
  }
};
</script>