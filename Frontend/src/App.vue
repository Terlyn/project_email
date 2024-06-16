<template>
  <div>
    <Header />
    <div>
      <Body :Emails="emails"/>
    </div>
    <Footer />
  </div>
</template>

<script>
import axios from 'axios';
import Body from "../src/components/Body.vue";
import Header from "../src/components/Header.vue";
import Footer from "../src/components/Footer.vue";

export default {
  name: 'App',
  components: {
    Header,
    Footer,
    Body
  },
  data() {
    return {
      emails: []
    }
  },
  mounted() {
    this.fetchEmails();
  },
  methods: {
    async fetchEmails() {
      try {
        const response = await axios.get('http://localhost:8081/emails');
        this.emails = response.data;
      } catch (error) {
        console.error("There was an error fetching the emails:", error);
      }
    }
  }
}
</script>

<style scoped>
.container {
  max-width: 100%;
}
</style>
