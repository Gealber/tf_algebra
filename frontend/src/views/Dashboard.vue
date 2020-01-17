<template>
    <v-container>
        <QuestionsList />
    </v-container>
</template>

<script>
    import QuestionsList from "../components/QuestionsList";    
    export default {
        name: "Dashboard",
        components: {QuestionsList},
    }
</script>

<style scoped>

</style>