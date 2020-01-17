<template>
    <v-container>
        <QuestionsList />
    </v-container>
</template>

<script>
    import QuestionsList from "../components/QuestionsList";
    import Rank from "../components/Rank";
    export default {
        name: "Dashboard",
        components: {QuestionsList, Rank},
    }
</script>

<style scoped>

</style>